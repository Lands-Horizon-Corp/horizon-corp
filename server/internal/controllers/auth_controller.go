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

// Basic Authentication
func (c *AuthController) SignUp(ctx *gin.Context)  {}
func (c *AuthController) SignIn(ctx *gin.Context)  {}
func (c *AuthController) SignOut(ctx *gin.Context) {}

// Password
func (c *AuthController) ForgotPassword(ctx *gin.Context) {}
func (c *AuthController) ChangePassword(ctx *gin.Context) {}

// Sending to user
func (c *AuthController) SendEmailVerification(ctx *gin.Context) {}
func (c *AuthController) SendOTPVerification(ctx *gin.Context)   {}

// Sending Verification
func (c *AuthController) VerifyEmail(ctx *gin.Context) {}
func (c *AuthController) VerifyOTP(ctx *gin.Context)   {}

func AuthRoutes(router *gin.RouterGroup, controller *AuthController) {
	group := router.Group("/auth")
	{
		group.POST("/signup", controller.SignUp)
		group.POST("/signin", controller.SignIn)
		group.POST("/signout", controller.SignOut)
		group.POST("/forgot-password", controller.ForgotPassword)
		group.POST("/change-password", controller.ChangePassword)
		group.POST("/send-email-verification", controller.SendEmailVerification)
		group.POST("/send-otp-verification", controller.SendOTPVerification)
		group.POST("/verify-email", controller.VerifyEmail)
		group.POST("/verify-otp", controller.VerifyOTP)
	}
}
