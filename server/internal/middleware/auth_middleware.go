// middleware/auth_middleware.go
package middleware

import (
	"errors"
	"horizon/server/config"
	"horizon/server/internal/auth"
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"
	"horizon/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	adminRepo           *repositories.AdminRepository
	employeeRepo        *repositories.EmployeeRepository
	ownerRepo           *repositories.OwnerRepository
	memberRepo          *repositories.MemberRepository
	adminAuthService    *auth.AdminAuthService
	employeeAuthService *auth.EmployeeAuthService
	memberAuthService   *auth.MemberAuthService
	ownerAuthService    *auth.OwnerAuthService
	otpService          *services.OTPService
	cfg                 *config.AppConfig
	tokenService        auth.TokenService
}

func NewAuthMiddleware(
	adminRepo *repositories.AdminRepository,
	employeeRepo *repositories.EmployeeRepository,
	ownerRepo *repositories.OwnerRepository,
	memberRepo *repositories.MemberRepository,
	adminAuthService *auth.AdminAuthService,
	employeeAuthService *auth.EmployeeAuthService,
	memberAuthService *auth.MemberAuthService,
	ownerAuthService *auth.OwnerAuthService,
	otpService *services.OTPService,
	cfg *config.AppConfig,
	tokenService auth.TokenService,
) *AuthMiddleware {
	return &AuthMiddleware{
		adminRepo:           adminRepo,
		employeeRepo:        employeeRepo,
		ownerRepo:           ownerRepo,
		memberRepo:          memberRepo,
		adminAuthService:    adminAuthService,
		employeeAuthService: employeeAuthService,
		memberAuthService:   memberAuthService,
		ownerAuthService:    ownerAuthService,
		otpService:          otpService,
		cfg:                 cfg,
		tokenService:        tokenService,
	}
}

func (c *AuthMiddleware) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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

		claims, err := c.tokenService.VerifyToken(cookie.Value)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: token verification failed"})
			ctx.Abort()
			return
		}

		var user interface{}
		switch claims.AccountType {
		case "Member":
			member, err := c.memberRepo.GetByID(claims.ID)
			if err != nil {
				c.tokenService.DeleteToken(cookie.Value)
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
				ctx.Abort()
				return
			}
			user = models.User{
				AccountType:       "Member",
				ID:                member.ID,
				Email:             member.Email,
				ContactNumber:     member.ContactNumber,
				FirstName:         member.FirstName,
				LastName:          member.LastName,
				IsEmailVerified:   member.IsEmailVerified,
				IsContactVerified: member.IsContactVerified,
			}

		case "Employee":
			employee, err := c.employeeRepo.GetByID(claims.ID)
			if err != nil {
				c.tokenService.DeleteToken(cookie.Value)
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
				ctx.Abort()
				return
			}
			user = models.User{
				AccountType:       "Employee",
				ID:                employee.ID,
				Email:             employee.Email,
				ContactNumber:     employee.ContactNumber,
				FirstName:         employee.FirstName,
				LastName:          employee.LastName,
				IsEmailVerified:   employee.IsEmailVerified,
				IsContactVerified: employee.IsContactVerified,
			}

		case "Admin":
			admin, err := c.adminRepo.GetByID(claims.ID)
			if err != nil {
				c.tokenService.DeleteToken(cookie.Value)
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
				ctx.Abort()
				return
			}
			user = models.User{
				AccountType:       "Admin",
				ID:                admin.ID,
				Email:             admin.Email,
				ContactNumber:     admin.ContactNumber,
				FirstName:         admin.FirstName,
				LastName:          admin.LastName,
				IsEmailVerified:   admin.IsEmailVerified,
				IsContactVerified: admin.IsContactVerified,
			}

		case "Owner":
			owner, err := c.ownerRepo.GetByID(claims.ID)
			if err != nil {
				c.tokenService.DeleteToken(cookie.Value)
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Owner not found"})
				ctx.Abort()
				return
			}
			user = models.User{
				AccountType:       "Owner",
				ID:                owner.ID,
				Email:             owner.Email,
				ContactNumber:     owner.ContactNumber,
				FirstName:         owner.FirstName,
				LastName:          owner.LastName,
				IsEmailVerified:   owner.IsEmailVerified,
				IsContactVerified: owner.IsContactVerified,
			}

		default:
			c.tokenService.DeleteToken(cookie.Value)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account type"})
			ctx.Abort()
			return
		}

		ctx.Set("current-user", user)
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
