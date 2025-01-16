package controllers

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	repository  *models.ModelRepository
	footstep    *handlers.FootstepHandler
	currentUser *handlers.CurrentUser
}

func NewAuthController(
	repository *models.ModelRepository,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
) *AuthController {
	return &AuthController{
		repository:  repository,
		footstep:    footstep,
		currentUser: currentUser,
	}
}

// SignUp handles user registration.
// Endpoint: POST /api/v1/auth/signup
type SignUpRequest struct {
	AccountType string `json:"accountType" validate:"required,max=10"`
}

func (as AuthController) SignUp(ctx *gin.Context) {

}

// SignIn handles user login.
// Endpoint: POST /api/v1/auth/signin
type SignInRequest struct {
	Key         string `json:"key" validate:"required,max=255"`
	Password    string `json:"password" validate:"required,min=8,max=255"`
	AccountType string `json:"accountType" validate:"required"`
}

func (as AuthController) SignIn(ctx *gin.Context) {}

// ForgotPassword handles password reset requests by sending a reset link.
// Endpoint: POST /api/v1/auth/forgot-password

type ForgotPasswordRequest struct {
	Key             string `json:"key" validate:"required,max=255"`
	AccountType     string `json:"accountType" validate:"required,max=10"`
	EmailTemplate   string `json:"emailTemplate"`
	ContactTemplate string `json:"contactTemplate"`
}

func (as AuthController) ForgotPassword(ctx *gin.Context) {}

// ChangePassword handles changing the user's password.
// Endpoint: POST /api/v1/auth/change-password
type ChangePasswordRequest struct {
	ResetID         string `json:"resetId" validate:"required,min=8,max=255"`
	NewPassword     string `json:"newPassword" validate:"required,min=8,max=255"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=NewPassword"`
}

func (as AuthController) ChangePassword(ctx *gin.Context) {}

// VerifyResetLink verifies the reset link for password recovery.
// Endpoint: GET /api/v1/auth/verify-reset-link/:id
func (as AuthController) VerifyResetLink(ctx *gin.Context) {}

// SignOut logs the user out of the system.
// Endpoint: POST /api/v1/auth/signout
func (as AuthController) SignOut(ctx *gin.Context) {}

// CurrentUser retrieves the currently logged-in user's details.
// Endpoint: GET /api/v1/auth/current-user (requires authentication)
func (as AuthController) CurrentUser(ctx *gin.Context) {}

// NewPassword sets a new password for the user after verification.
// Endpoint: POST /api/v1/auth/new-password (requires authentication)
type NewPasswordRequest struct {
	PreviousPassword string `json:"previousPassword" validate:"required,min=8,max=255"`
	NewPassword      string `json:"newPassword" validate:"required,min=8,max=255"`
	ConfirmPassword  string `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=NewPassword"`
}

func (as AuthController) NewPassword(ctx *gin.Context) {}

// SkipVerification allows skipping the verification process under certain conditions.
// Endpoint: POST /api/v1/auth/skip-verification (requires authentication)
func (as AuthController) SkipVerification(ctx *gin.Context) {}

// SendEmailVerification sends a verification email to the user.
// Endpoint: POST /api/v1/auth/send-email-verification (requires authentication)
type SendEmailVerificationRequest struct {
	EmailTemplate string `json:"emailTemplate" validate:"required"`
}

func (as AuthController) SendEmailVerification(ctx *gin.Context) {}

// VerifyEmail verifies the user's email address using the provided token.
// Endpoint: POST /api/v1/auth/verify-email (requires authentication)
type VerifyEmailRequest struct {
	Otp string `json:"otp" validate:"required,len=6"`
}

func (as AuthController) VerifyEmail(ctx *gin.Context) {}

// SendContactNumberVerification sends a verification code to the user's contact number.
// Endpoint: POST /api/v1/auth/send-contact-number-verification (requires authentication)

type SendContactNumberVerificationRequest struct {
	ContactTemplate string `json:"contactTemplate" validate:"required"`
}

func (as AuthController) SendContactNumberVerification(ctx *gin.Context) {}

// VerifyContactNumber verifies the user's contact number using the provided code.
// Endpoint: POST /api/v1/auth/verify-contact-number (requires authentication)

type VerifyContactNumberRequest struct {
	Otp string `json:"otp" validate:"required,len=6"`
}

func (as AuthController) VerifyContactNumber(ctx *gin.Context) {}
