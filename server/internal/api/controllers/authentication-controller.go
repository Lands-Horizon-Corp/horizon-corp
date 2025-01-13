package controllers

import "github.com/gin-gonic/gin"

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

// SignUp handles user registration.
// Endpoint: POST /api/v1/auth/signup
func (as AuthController) SignUp(ctx *gin.Context) {}

// SignIn handles user login.
// Endpoint: POST /api/v1/auth/signin
func (as AuthController) SignIn(ctx *gin.Context) {}

// ForgotPassword handles password reset requests by sending a reset link.
// Endpoint: POST /api/v1/auth/forgot-password
func (as AuthController) ForgotPassword(ctx *gin.Context) {}

// ChangePassword handles changing the user's password.
// Endpoint: POST /api/v1/auth/change-password
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
func (as AuthController) NewPassword(ctx *gin.Context) {}

// SkipVerification allows skipping the verification process under certain conditions.
// Endpoint: POST /api/v1/auth/skip-verification (requires authentication)
func (as AuthController) SkipVerification(ctx *gin.Context) {}

// SendEmailVerification sends a verification email to the user.
// Endpoint: POST /api/v1/auth/send-email-verification (requires authentication)
func (as AuthController) SendEmailVerification(ctx *gin.Context) {}

// VerifyEmail verifies the user's email address using the provided token.
// Endpoint: POST /api/v1/auth/verify-email (requires authentication)
func (as AuthController) VerifyEmail(ctx *gin.Context) {}

// SendContactNumberVerification sends a verification code to the user's contact number.
// Endpoint: POST /api/v1/auth/send-contact-number-verification (requires authentication)
func (as AuthController) SendContactNumberVerification(ctx *gin.Context) {}

// VerifyContactNumber verifies the user's contact number using the provided code.
// Endpoint: POST /api/v1/auth/verify-contact-number (requires authentication)
func (as AuthController) VerifyContactNumber(ctx *gin.Context) {}
