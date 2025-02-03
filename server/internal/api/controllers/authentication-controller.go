package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type AuthController struct {
	authHandler   *handlers.AuthHandler
	currentUser   *handlers.CurrentUser
	tokenProvider *providers.TokenService
	cfg           *config.AppConfig
	helpers       *helpers.HelpersFunction
}

func NewAuthController(
	authHandler *handlers.AuthHandler,
	currentUser *handlers.CurrentUser,
	tokenProvider *providers.TokenService,
	cfg *config.AppConfig,
	helpers *helpers.HelpersFunction,
) *AuthController {
	return &AuthController{
		authHandler:   authHandler,
		currentUser:   currentUser,
		tokenProvider: tokenProvider,
		cfg:           cfg,
		helpers:       helpers,
	}
}

// SignUp handles user registration.
// Endpoint: POST /api/v1/auth/signup
type SignUpRequest struct {
	AccountType      string    `json:"accountType" validate:"required,max=10"`
	FirstName        string    `json:"firstName" validate:"required,min=2,max=255"`
	LastName         string    `json:"lastName" validate:"required,min=2,max=255"`
	MiddleName       string    `json:"middleName" validate:"omitempty,min=2,max=255"`
	PermanentAddress string    `json:"permanentAddress" validate:"required"`
	Description      string    `json:"description" validate:"omitempty"`
	BirthDate        time.Time `json:"birthDate" validate:"required,datetime=2006-01-02"`
	Username         string    `json:"username" validate:"required,alphanum,min=3,max=255"`
	Email            string    `json:"email" validate:"required,email"`
	Password         string    `json:"password" validate:"required,min=8,max=255"`
	ConfirmPassword  string    `json:"confirmPassword" validate:"required,eqfield=Password"`
	ContactNumber    string    `json:"contactNumber" validate:"required,e164"`

	MediaID  *uuid.UUID `json:"mediaId" validate:"omitempty,uuid"`
	RoleID   *uuid.UUID `json:"roleId" validate:"omitempty,uuid"`
	GenderID *uuid.UUID `json:"genderId" validate:"omitempty,uuid"`

	EmailTemplate   string `json:"emailTemplate" validate:"required"`
	ContactTemplate string `json:"contactTemplate" validate:"required"`
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

	user, err := as.authHandler.Create(
		ctx,
		req.AccountType,
		req.FirstName,
		req.LastName,
		req.MiddleName,
		req.PermanentAddress,
		req.Description,
		req.BirthDate,
		req.Username,
		req.Email,
		req.Password,
		req.ContactNumber,
		req.MediaID,
		req.RoleID,
		req.GenderID,
		req.EmailTemplate,
		req.ContactTemplate,
	)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userMap, ok := user.(map[string]interface{})
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user data format"})
		return
	}
	userID, ok := userMap["id"].(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}
	userStatus, ok := userMap["status"].(providers.UserStatus)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user Status"})
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, token, err := as.authHandler.SignIn(ctx, req.AccountType, req.Key, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	ctx.JSON(http.StatusOK, user)
}

// ForgotPassword handles password reset requests by sending a reset link.
// Endpoint: POST /api/v1/auth/forgot-password

type ForgotPasswordRequest struct {
	AccountType     string `json:"accountType" validate:"required,max=10"`
	Key             string `json:"key" validate:"required,max=255"`
	EmailTemplate   string `json:"emailTemplate"`
	ContactTemplate string `json:"contactTemplate"`
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
	link, err := as.authHandler.ForgotPasswordResetLink(ctx, req.AccountType, req.Key, req.EmailTemplate, req.ContactTemplate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"link": link})
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
	user, err := as.authHandler.ForceChangePassword(ctx, claims.AccountType, claims.ID, req.ConfirmPassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	as.tokenProvider.DeleteToken(req.ResetID)
	ctx.JSON(http.StatusOK, user)
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

type NewPasswordRequest struct {
	PreviousPassword string `json:"previousPassword" validate:"required,min=8,max=255"`
	NewPassword      string `json:"newPassword" validate:"required,min=8,max=255"`
	ConfirmPassword  string `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=NewPassword"`
}

func (as AuthController) NewPassword(ctx *gin.Context) {
	var req NewPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}
	user, _, err := as.currentUser.Claims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := as.authHandler.NewPassword(ctx, user.AccountType, req.PreviousPassword, req.ConfirmPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, updatedUser)
}

func (as AuthController) SkipVerification(ctx *gin.Context) {
	user, _, err := as.currentUser.Claims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	updatedUser, err := as.authHandler.SkipVerification(ctx, user.AccountType)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, updatedUser)
}

type SendEmailVerificationRequest struct {
	EmailTemplate string `json:"emailTemplate" validate:"required"`
}

func (as AuthController) SendEmailVerification(ctx *gin.Context) {
	var req SendEmailVerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, _, err := as.currentUser.Claims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if err := as.authHandler.SendEmailVerification(ctx, user.AccountType, req.EmailTemplate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Email verification sent successfully. Please check your inbox or spam folder"})
}

type VerifyEmailRequest struct {
	Otp string `json:"otp" validate:"required,len=6"`
}

func (as AuthController) VerifyEmail(ctx *gin.Context) {
	var req VerifyEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, _, err := as.currentUser.Claims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	updatedUser, err := as.authHandler.VerifyEmail(ctx, user.AccountType, req.Otp)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, updatedUser)
}

type SendContactNumberVerificationRequest struct {
	ContactTemplate string `json:"contactTemplate" validate:"required"`
}

func (as AuthController) SendContactNumberVerification(ctx *gin.Context) {
	var req SendContactNumberVerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, _, err := as.currentUser.Claims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if err := as.authHandler.SendContactNumberVerification(ctx, user.AccountType, req.ContactTemplate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Contact number verification OTP sent successfully"})
}

type VerifyContactNumberRequest struct {
	Otp string `json:"otp" validate:"required,len=6"`
}

func (as AuthController) VerifyContactNumber(ctx *gin.Context) {
	var req VerifyContactNumberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, _, err := as.currentUser.Claims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	updatedUser, err := as.authHandler.VerifyContactNumber(ctx, user.AccountType, req.Otp)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedUser)

}
