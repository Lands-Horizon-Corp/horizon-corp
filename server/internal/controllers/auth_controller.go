package controllers

import (
	"horizon/server/internal/repositories"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	adminRepo    *repositories.AdminRepository
	employeeRepo *repositories.EmployeeRepository
	ownerRepo    *repositories.OwnerRepository
	memberRepo   *repositories.MemberRepository
}

func NewAuthController(adminRepo *repositories.AdminRepository,
	employeeRepo *repositories.EmployeeRepository,
	ownerRepo *repositories.OwnerRepository,
	memberRepo *repositories.MemberRepository) *AuthController {
	return &AuthController{
		adminRepo:    adminRepo,
		employeeRepo: employeeRepo,
		ownerRepo:    ownerRepo,
		memberRepo:   memberRepo,
	}
}

// User
func (c *AuthController) CurrentUser(ctx *gin.Context) {}

// Basic Authentication
func (c *AuthController) SignUp(ctx *gin.Context)  {}
func (c *AuthController) SignIn(ctx *gin.Context)  {}
func (c *AuthController) SignOut(ctx *gin.Context) {}

// Password
func (c *AuthController) ForgotPassword(ctx *gin.Context) {}
func (c *AuthController) ChangePassword(ctx *gin.Context) {}

// Email
func (c *AuthController) ChangeEmail(ctx *gin.Context)           {}
func (c *AuthController) SendEmailVerification(ctx *gin.Context) {}
func (c *AuthController) VerifyEmail(ctx *gin.Context)           {}

// Contact Number
func (c *AuthController) ChangeContactNumber(ctx *gin.Context)           {}
func (c *AuthController) SendContactNumberVerification(ctx *gin.Context) {}
func (c *AuthController) VerifyContactNumber(ctx *gin.Context)           {}

func AuthRoutes(router *gin.RouterGroup, controller *AuthController) {
	group := router.Group("/auth")
	{
		// Basic Authentication
		group.POST("/current-user", controller.CurrentUser)

		// Basic Authentication
		group.POST("/signup", controller.SignUp)
		group.POST("/signin", controller.SignIn)
		group.POST("/signout", controller.SignOut)

		// Password
		group.POST("/forgot-password", controller.ForgotPassword)
		group.POST("/change-password", controller.ChangePassword)

		// Email
		group.POST("/change-email", controller.ChangeEmail)
		group.POST("/send-email-verification", controller.SendEmailVerification)
		group.POST("/verify-email", controller.VerifyEmail)

		// Contact Number
		group.POST("/change-contact-number", controller.ChangeContactNumber)
		group.POST("/send-otp-verification", controller.SendContactNumberVerification)
		group.POST("/verify-otp", controller.VerifyContactNumber)
	}
}
