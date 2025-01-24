package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
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
	helpers       *helpers.HelpersFunction

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
	helpers *helpers.HelpersFunction,

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
		helpers:       helpers,

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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var (
		userID     string
		userStatus providers.UserStatus
	)
	switch req.AccountType {
	case "Member":
		member := as.memberController.Create(ctx)
		userID = member.ID.String()
		userStatus = member.Status
	case "Admin":
		admin := as.adminController.Create(ctx)
		userID = admin.ID.String()
		userStatus = admin.Status
	case "Owner":
		owner := as.ownerController.Create(ctx)
		userID = owner.ID.String()
		userStatus = owner.Status
	case "Employee":
		employee := as.employeeController.Create(ctx)
		userID = employee.ID.String()
		userStatus = employee.Status
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
		return
	}
	token, err := as.tokenProvider.GenerateUserToken(
		providers.UserClaims{
			ID:          userID,
			AccountType: req.AccountType,
			UserStatus:  userStatus,
		},
		time.Hour*24,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": err.Error()})
		return
	}
	preloads := as.helpers.GetPreload(ctx)
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
	Key         string `json:"key" validate:"required,max=255"`
	AccountType string `json:"accountType" validate:"required,max=10"`
}

func (as AuthController) ForgotPassword(ctx *gin.Context) {
	var req ForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": err.Error()})
		return
	}
	switch req.AccountType {
	case "Member":
		as.memberController.ForgotPassword(ctx)
	case "Admin":
		as.adminController.ForgotPasswordResetLink(ctx)
	case "Owner":
		as.ownerController.ForgotPasswordResetLink(ctx)
	case "Employee":
		as.employeeController.ForgotPasswordResetLink(ctx)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid account type"})
		return
	}
}

// ChangePassword handles changing the user's password.
// Endpoint: POST /api/v1/auth/change-password
type ChangePasswordRequest struct {
	ResetID         string `json:"resetId" validate:"required,min=8,max=255"`
	NewPassword     string `json:"newPassword" validate:"required,min=8,max=255"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=NewPassword"`
}

func (as AuthController) ChangePassword(ctx *gin.Context) {
	var req ChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("ChangePassword: JSON binding error: %v", err)})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	claims, err := as.tokenProvider.VerifyToken(req.ResetID)
	if err != nil {
		as.tokenProvider.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("ChangePassword: Token verification error: %v", err)})
		return
	}
	switch claims.AccountType {
	case "Member":
		as.repository.MemberForceChangePassword(claims.ID, req.NewPassword)
	case "Admin":
		as.repository.AdminForceChangePassword(claims.ID, req.NewPassword)
	case "Owner":
		as.repository.OwnerForceChangePassword(claims.ID, req.NewPassword)
	case "Employee":
		as.repository.EmployeeForceChangePassword(claims.ID, req.NewPassword)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
	as.tokenProvider.DeleteToken(req.ResetID)
}

func (as AuthController) VerifyResetLink(ctx *gin.Context) {
	_, err := as.tokenProvider.VerifyToken(ctx.Param("id"))
	if err != nil {
		as.tokenProvider.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("VerifyResetLink: Token verification error: %v", err)})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Link verified successfully."})
}

func (as AuthController) SignOut(ctx *gin.Context) {
	as.tokenProvider.ClearTokenCookie(ctx)
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully signed out"})
}

func (as AuthController) CurrentUser(ctx *gin.Context) {
	user, err := as.currentUser.GenericUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (as AuthController) NewPassword(ctx *gin.Context) {
	user, _, err := as.currentUser.Claims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	switch user.AccountType {
	case "Member":
		as.memberController.NewPassword(ctx)
	case "Admin":
		as.adminController.NewPassword(ctx)
	case "Owner":
		as.ownerController.NewPassword(ctx)
	case "Employee":
		as.employeeController.NewPassword(ctx)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
}

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
