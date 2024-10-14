package middleware

import (
	"errors"
	"horizon/server/config"
	"horizon/server/internal/auth"
	"horizon/server/internal/repositories"
	"horizon/server/internal/resources"
	"horizon/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {

	// Database
	adminRepo    *repositories.AdminRepository
	employeeRepo *repositories.EmployeeRepository
	ownerRepo    *repositories.OwnerRepository
	memberRepo   *repositories.MemberRepository

	// Auth
	adminAuthService    *auth.AdminAuthService
	employeeAuthService *auth.EmployeeAuthService
	memberAuthService   *auth.MemberAuthService
	ownerAuthService    *auth.OwnerAuthService

	// Services
	otpService   *services.OTPService
	cfg          *config.AppConfig
	tokenService auth.TokenService
}

func NewAuthMiddleware(
	// Database
	adminRepo *repositories.AdminRepository,
	employeeRepo *repositories.EmployeeRepository,
	ownerRepo *repositories.OwnerRepository,
	memberRepo *repositories.MemberRepository,

	// Auth
	adminAuthService *auth.AdminAuthService,
	employeeAuthService *auth.EmployeeAuthService,
	memberAuthService *auth.MemberAuthService,
	ownerAuthService *auth.OwnerAuthService,

	// Services
	otpService *services.OTPService,
	cfg *config.AppConfig,
	tokenService auth.TokenService,

) *AuthMiddleware {
	return &AuthMiddleware{
		// Database
		adminRepo:    adminRepo,
		employeeRepo: employeeRepo,
		ownerRepo:    ownerRepo,
		memberRepo:   memberRepo,

		// Auth
		adminAuthService:    adminAuthService,
		employeeAuthService: employeeAuthService,
		memberAuthService:   memberAuthService,
		ownerAuthService:    ownerAuthService,

		// Services
		otpService:   otpService,
		cfg:          cfg,
		tokenService: tokenService,
	}
}

func (c *AuthMiddleware) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Retrieve the cookie
		cookie, err := ctx.Request.Cookie(c.cfg.AppTokenName)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: cookie not found"})
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			ctx.Abort()
			return
		}

		// Verify the token
		claims, err := c.tokenService.VerifyToken(cookie.Value)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: token verification failed"})
			ctx.Abort()
			return
		}

		// Fetch user based on account type and set in context
		var user interface{}
		switch claims.AccountType {
		case "Member":
			member, err := c.memberRepo.GetByID(claims.ID)
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
				ctx.Abort()
				return
			}
			user = resources.ToResourceMember(member)

		case "Employee":
			employee, err := c.employeeRepo.GetByID(claims.ID)
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
				ctx.Abort()
				return
			}
			user = resources.ToResourceEmployee(employee)

		case "Admin":
			admin, err := c.adminRepo.GetByID(claims.ID)
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
				ctx.Abort()
				return
			}
			user = resources.ToResourceAdmin(admin)

		case "Owner":
			owner, err := c.ownerRepo.GetByID(claims.ID)
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Owner not found"})
				ctx.Abort()
				return
			}
			user = resources.ToResourceOwner(owner)

		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account type"})
			ctx.Abort()
			return
		}

		// Store the user in context for further use in handlers
		ctx.Set("current-user", user)
		ctx.Next() // Continue to the next handler
	}
}
