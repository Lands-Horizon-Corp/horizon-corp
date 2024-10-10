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

func (c *AuthController) SignIn(ctx *gin.Context)                 {}
func (c *AuthController) SignOut(ctx *gin.Context)                {}
func (c *AuthController) SignUp(ctx *gin.Context)                 {}
func (c *AuthController) ForgotPassword(ctx *gin.Context)         {}
func (c *AuthController) ChangePasswordPassword(ctx *gin.Context) {}
func (c *AuthController) SendEmailVerification(ctx *gin.Context)  {}
func (c *AuthController) SendOTPVerification(ctx *gin.Context)    {}
func (c *AuthController) VerifyEmail(ctx *gin.Context)            {}
func (c *AuthController) VerifyOTP(ctx *gin.Context)              {}
