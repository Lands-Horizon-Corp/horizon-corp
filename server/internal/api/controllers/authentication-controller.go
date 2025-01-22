package controllers

import (
	"net/http"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type AuthController struct {
	repository    *models.ModelRepository
	transformer   *models.ModelTransformer
	footstep      *handlers.FootstepHandler
	currentUser   *handlers.CurrentUser
	tokenProvider *providers.TokenService
	cfg           *config.AppConfig

	adminController    *AdminController
	ownerController    *OwnerController
	employeeController *EmployeeController
	memberController   *MemberController
}

func NewAuthController(
	repository *models.ModelRepository,
	transformer *models.ModelTransformer,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
	tokenProvider *providers.TokenService,
	cfg *config.AppConfig,

	adminController *AdminController,
	ownerController *OwnerController,
	employeeController *EmployeeController,
	memberController *MemberController,
) *AuthController {
	return &AuthController{
		repository:    repository,
		transformer:   transformer,
		footstep:      footstep,
		currentUser:   currentUser,
		tokenProvider: tokenProvider,
		cfg:           cfg,

		adminController:    adminController,
		ownerController:    ownerController,
		employeeController: employeeController,
		memberController:   memberController,
	}
}

// SignUp handles user registration.
// Endpoint: POST /api/v1/auth/signup
type SignUpRequest struct {
	AccountType string `json:"accountType" validate:"required,max=10"`
}

func (as AuthController) SignUp(ctx *gin.Context) {
	var req SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if validator.New().Struct(req) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": req})
		return
	}
	switch req.AccountType {
	case "Member":
		as.memberController.Store(ctx)
	case "Admin":
		as.adminController.Store(ctx)
	case "Owner":
		as.ownerController.Store(ctx)
	case "Employee":
		as.employeeController.Store(ctx)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
}

// SignIn handles user login.
// Endpoint: POST /api/v1/auth/signin
type SignInRequest struct {
	Key         string `json:"key" validate:"required,max=255"`
	Password    string `json:"password" validate:"required,min=8,max=255"`
	AccountType string `json:"accountType" validate:"required"`
}

func (as AuthController) SignIn(ctx *gin.Context) {
	var req SignInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data", "details": err.Error()})
		return
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": err.Error()})
		return
	}
	preloads := ctx.QueryArray("preload")
	var rawUser interface{}
	var transformedUser interface{}
	var id uuid.UUID
	var err error
	switch req.AccountType {
	case "Member":
		rawUser, err = as.repository.MemberSignIn(req.Key, req.Password, preloads...)
		if err == nil {
			id = rawUser.(*models.Member).ID
			transformedUser = as.transformer.MemberToResource(rawUser.(*models.Member))
		}
	case "Admin":
		rawUser, err = as.repository.AdminSignIn(req.Key, req.Password, preloads...)
		if err == nil {
			id = rawUser.(*models.Admin).ID
			transformedUser = as.transformer.AdminToResource(rawUser.(*models.Admin))
		}
	case "Owner":
		rawUser, err = as.repository.OwnerSignIn(req.Key, req.Password, preloads...)
		if err == nil {
			id = rawUser.(*models.Owner).ID
			transformedUser = as.transformer.OwnerToResource(rawUser.(*models.Owner))
		}
	case "Employee":
		rawUser, err = as.repository.EmployeeSignIn(req.Key, req.Password, preloads...)
		if err == nil {
			id = rawUser.(*models.Employee).ID
			transformedUser = as.transformer.EmployeeToResource(rawUser.(*models.Employee))
		}
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid account type"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authentication failed", "details": err.Error()})
		return
	}
	token, err := as.tokenProvider.GenerateUserToken(providers.UserClaims{
		ID:          id.String(),
		AccountType: req.AccountType,
	}, 24*time.Hour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token", "details": err.Error()})
		return
	}
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     as.cfg.AppTokenName,
		Value:    *token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})
	ctx.JSON(http.StatusOK, transformedUser)
}

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
