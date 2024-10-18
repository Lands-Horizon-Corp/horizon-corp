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
	// Repositories
	adminRepo    *repositories.AdminRepository
	employeeRepo *repositories.EmployeeRepository
	ownerRepo    *repositories.OwnerRepository
	memberRepo   *repositories.MemberRepository
	mediaRepo    *repositories.MediaRepository

	// Auth Service
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
	// Repositories
	adminRepo *repositories.AdminRepository,
	employeeRepo *repositories.EmployeeRepository,
	ownerRepo *repositories.OwnerRepository,
	memberRepo *repositories.MemberRepository,
	mediaRepo *repositories.MediaRepository,

	// Auth service
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
		// Repositories
		adminRepo:    adminRepo,
		employeeRepo: employeeRepo,
		ownerRepo:    ownerRepo,
		memberRepo:   memberRepo,
		mediaRepo:    mediaRepo,

		// Auth service
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
			var media *models.Media
			if member.MediaID != nil {
				m, err := c.mediaRepo.GetByID(*member.MediaID)
				if err == nil {
					media = &m // create a pointer to the media
				}
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
				Status:            string(member.Status),
				Media:             media,
			}

		case "Employee":
			employee, err := c.employeeRepo.GetByID(claims.ID)
			if err != nil {
				c.tokenService.DeleteToken(cookie.Value)
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
				ctx.Abort()
				return
			}
			var media *models.Media
			if employee.MediaID != nil {
				m, err := c.mediaRepo.GetByID(*employee.MediaID)
				if err == nil {
					media = &m // create a pointer to the media
				}
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
				Status:            string(employee.Status),
				Media:             media,
			}

		case "Admin":
			admin, err := c.adminRepo.GetByID(claims.ID)
			if err != nil {
				c.tokenService.DeleteToken(cookie.Value)
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
				ctx.Abort()
				return
			}
			var media *models.Media
			if admin.MediaID != nil {
				m, err := c.mediaRepo.GetByID(*admin.MediaID)
				if err == nil {
					media = &m // create a pointer to the media
				}
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
				Status:            string(admin.Status),
				Media:             media,
			}

		case "Owner":
			owner, err := c.ownerRepo.GetByID(claims.ID)
			if err != nil {
				c.tokenService.DeleteToken(cookie.Value)
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Owner not found"})
				ctx.Abort()
				return
			}
			var media *models.Media
			if owner.MediaID != nil {
				m, err := c.mediaRepo.GetByID(*owner.MediaID)
				if err == nil {
					media = &m
				}
			}
			user = models.User{
				AccountType:       "Owner",
				ID:                owner.ID,
				Email:             owner.Email,
				ContactNumber:     owner.ContactNumber,
				FirstName:         owner.FirstName,
				LastName:          owner.LastName,
				MiddleName:        owner.MiddleName,
				IsEmailVerified:   owner.IsEmailVerified,
				IsContactVerified: owner.IsContactVerified,
				Status:            string(owner.Status),
				Media:             media,
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
