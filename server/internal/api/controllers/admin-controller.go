package controllers

import (
	"net/http"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type AdminController struct {
	authHandler     *handlers.AuthHandler
	footstepHandler *handlers.FootstepHandler
	cfg             *config.AppConfig
}

func NewAdminController(
	authHandler *handlers.AuthHandler,
	footstepHandler *handlers.FootstepHandler,
	cfg *config.AppConfig,
) *AdminController {
	return &AdminController{
		authHandler:     authHandler,
		footstepHandler: footstepHandler,
		cfg:             cfg,
	}
}

// GET: /api/v1/admin
// Retrieve admins with optional filtering for pagination or no pagination. Results can be converted to records.
//
//	Admin: Can retrieve all admins if admin  status is verified
//	Employee: not allowed
//	Member: not allowed
func (c *AdminController) Index(ctx *gin.Context) {

}

// GET: /api/v1/admin/:id
//
//	Admin: if admin  status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed

func (c *AdminController) Show(ctx *gin.Context) {
	// if id or email or
}

// POST: /api/v1/admin
//
//	Admin: Can create admin and automatically verified and also if admin and the status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
//	public allowed
type AdminStoreRequest struct {
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

func (c *AdminController) Store(ctx *gin.Context) {
	c.Create(ctx)
}

// PUT: /api/v1/admin/:id
//
//	Admin: Can change status of admin but only if admin and the status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
func (c *AdminController) Update(ctx *gin.Context) {

}

// DELETE:/api/v1/admin/:id
// Verifiy admin
//
//	Admin: only if admin  status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
func (c *AdminController) Destroy(ctx *gin.Context) {

}

type AdminChangePasswordRequest struct {
	OldPassword     string `json:"old_password" validate:"required,min=8,max=255"`
	NewPassword     string `json:"new_password" validate:"required,min=8,max=255"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8,max=255"`
}

func (as AdminController) ChangePassword(ctx *gin.Context) {
	var req AdminChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	user, err := as.authHandler.ChangePassword(ctx, "Admin", req.OldPassword, req.NewPassword, req.ConfirmPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type AdminForgotPasswordRequest struct {
	Key             string `json:"key" validate:"required,max=255"`
	EmailTemplate   string `json:"emailTemplate"`
	ContactTemplate string `json:"contactTemplate"`
}

func (c *AdminController) ForgotPassword(ctx *gin.Context) {
	link, err := c.ForgotPasswordResetLink(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"link": link})
}

func (c *AdminController) ForgotPasswordResetLink(ctx *gin.Context) (string, error) {
	var req AdminForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return "", err
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return "", err
	}
	return c.authHandler.ForgotPasswordResetLink(ctx, "Admin", req.Key, req.EmailTemplate, req.ContactTemplate)
}

func (c *AdminController) Create(ctx *gin.Context) {
	var req AdminStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := c.authHandler.SignUp(ctx, "Admin",
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
		req.ContactTemplate)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type AdminNewPasswordRequest struct {
	PreviousPassword string `json:"previousPassword" validate:"required,min=8,max=255"`
	NewPassword      string `json:"newPassword" validate:"required,min=8,max=255"`
	ConfirmPassword  string `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=NewPassword"`
}

func (c *AdminController) NewPassword(ctx *gin.Context) {
	var req AdminChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}
	user, err := c.authHandler.NewPassword(ctx, "Admin", req.OldPassword, req.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *AdminController) SkipVerification(ctx *gin.Context) {
	user, err := c.authHandler.SkipVerification(ctx, "Admin")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type AdminSendEmailVerificationRequest struct {
	EmailTemplate string `json:"emailTemplate" validate:"required"`
}

func (c *AdminController) SendEmailVerification(ctx *gin.Context) {
	var req AdminSendEmailVerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	if err := c.authHandler.SendEmailVerification(ctx, "Admin", req.EmailTemplate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Email verification sent successfully. Please check your inbox or spam folder"})
}

type AdminVerifyEmailRequest struct {
	Otp string `json:"otp" validate:"required,len=6"`
}

func (c *AdminController) VerifyEmail(ctx *gin.Context) {
	var req AdminVerifyEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, err := c.authHandler.VerifyEmail(ctx, "Admin", req.Otp)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type AdminSendContactNumberVerificationRequest struct {
	ContactTemplate string `json:"contactTemplate" validate:"required"`
}

func (c *AdminController) SendContactNumberVerification(ctx *gin.Context) {
	var req AdminSendContactNumberVerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	if err := c.authHandler.SendContactNumberVerification(ctx, "Admin", req.ContactTemplate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Contact number verification OTP sent successfully"})
}

type AdminVerifyContactNumberRequest struct {
	Otp string `json:"otp" validate:"required,len=6"`
}

func (c *AdminController) VerifyContactNumber(ctx *gin.Context) {
	var req AdminVerifyContactNumberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, err := c.authHandler.VerifyContactNumber(ctx, "Admin", req.Otp)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *AdminController) ProfilePicture(ctx *gin.Context) {
	var req MediaStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, err := c.authHandler.ProfilePicture(ctx, "Admin", req.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type AdminAccountSettingRequest struct {
	BirthDate        time.Time `json:"birthDate" validate:"required"`
	FirstName        string    `json:"firstName" validate:"required,max=255"`
	MiddleName       string    `json:"middleName" validate:"required,max=255"`
	LastName         string    `json:"lastName" validate:"required,max=255"`
	Description      string    `json:"description" validate:"max=2048"`
	PermanentAddress string    `json:"permanentAddress" validate:"required,max=500"`
}

func (c *AdminController) ProfileAccountSetting(ctx *gin.Context) {
	var req AdminAccountSettingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, err := c.authHandler.ProfileAccountSetting(ctx, "Admin",
		req.BirthDate,
		req.FirstName,
		req.MiddleName,
		req.LastName,
		req.Description,
		req.PermanentAddress,
	)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type AdminChangeEmailRequest struct {
	Password string `json:"password" validate:"required,min=8,max=255"`
	Email    string `json:"email" validate:"required,email"`
}

func (c *AdminController) ProfileChangeEmail(ctx *gin.Context) {
	var req AdminChangeEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, err := c.authHandler.ProfileChangeEmail(ctx, "Admin", req.Password, req.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type AdminChangeContactNumberRequest struct {
	Password      string `json:"password" validate:"required,min=8,max=255"`
	ContactNumber string `json:"contactNumber" validate:"required,min=10,max=15"`
}

func (c *AdminController) ProfileChangeContactNumber(ctx *gin.Context) {
	var req AdminChangeContactNumberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	user, err := c.authHandler.ProfileChangeContactNumber(ctx, "Admin", req.Password, req.ContactNumber)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

type AdminChangeUsernameRequest struct {
	Password string `json:"password" validate:"required,min=8,max=255"`
	Username string `json:"username" validate:"required,min=5,max=255"`
}

func (c *AdminController) ProfileChangeUsername(ctx *gin.Context) {
	var req AdminChangeUsernameRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, err := c.authHandler.ProfileChangeContactNumber(ctx, "Admin", req.Password, req.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type AdminSignInRequest struct {
	Key      string `json:"key" validate:"required,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

func (as AdminController) SignIn(ctx *gin.Context) {
	var req AdminSignInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, token, err := as.authHandler.SignIn(ctx, "Admin", req.Key, req.Password)
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
