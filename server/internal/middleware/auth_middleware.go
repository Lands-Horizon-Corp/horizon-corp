// middleware/auth_middleware.go
package middleware

import (
	"errors"
	"horizon/server/config"
	"horizon/server/internal/auth"
	"horizon/server/internal/repositories"
	"horizon/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	useRepo      *repositories.UserRepository
	otpService   *services.OTPService
	cfg          *config.AppConfig
	tokenService auth.TokenService
}

func NewAuthMiddleware(
	useRepo *repositories.UserRepository,
	otpService *services.OTPService,
	cfg *config.AppConfig,
	tokenService auth.TokenService,
) *AuthMiddleware {
	return &AuthMiddleware{
		useRepo:      useRepo,
		otpService:   otpService,
		cfg:          cfg,
		tokenService: tokenService,
	}
}

func (m *AuthMiddleware) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie(m.cfg.AppTokenName)
		if err != nil {
			m.tokenService.ClearTokenCookie(ctx)
			if errors.Is(err, http.ErrNoCookie) {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: cookie not found"})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
			ctx.Abort()
			return
		}

		claims, err := m.tokenService.VerifyToken(cookie.Value)
		if err != nil {
			m.tokenService.ClearTokenCookie(ctx)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: token verification failed"})
			ctx.Abort()
			return
		}

		user, err := m.useRepo.GetByID(claims.AccountType, claims.ID)
		user.AccountType = claims.AccountType
		if err != nil {
			m.tokenService.ClearTokenCookie(ctx)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: user not found"})
			ctx.Abort()
			return
		}

		ctx.Set("current-user", user)
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
