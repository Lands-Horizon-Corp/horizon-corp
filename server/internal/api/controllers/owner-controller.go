package controllers

import (
	"net/http"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type OwnerController struct {
	authHandler     *handlers.AuthHandler
	footstepHandler *handlers.FootstepHandler
}

func NewOwnerController(
	authHandler *handlers.AuthHandler,
	footstepHandler *handlers.FootstepHandler,
) *OwnerController {
	return &OwnerController{
		authHandler:     authHandler,
		footstepHandler: footstepHandler,
	}
}

// GET: /api/v1/owner
// Retrieve owners with optional filtering for pagination or no pagination. Results can be converted to records.
//
//	Owner: Can retrieve all owners if owner  status is verified
//	Employee: not allowed
//	Member: not allowed
func (c *OwnerController) Index(ctx *gin.Context) {

}

// GET: /api/v1/owner/:id
//
//      Owner: if owner  status is verified
//      Employee: not allowed
//      Owner: not allowed
//      Member: not allowed

func (c *OwnerController) Show(ctx *gin.Context) {
	// if id or email or
}

// POST: /api/v1/owner
//
//	Owner: Can create owner and automatically verified and also if owner and the status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
//	public allowed
type OwnerStoreRequest struct {
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

func (c *OwnerController) Store(ctx *gin.Context) {
	c.Create(ctx)
}

// PUT: /api/v1/owner/:id
//
//	Owner: Can change status of owner but only if owner and the status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
func (c *OwnerController) Update(ctx *gin.Context) {

}

// DELETE:/api/v1/owner/:id
// Verifiy owner
//
//	Owner: only if owner  status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
func (c *OwnerController) Destroy(ctx *gin.Context) {

}

type OwnerChangePasswordRequest struct {
	OldPassword     string `json:"old_password" validate:"required,min=8,max=255"`
	NewPassword     string `json:"new_password" validate:"required,min=8,max=255"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8,max=255"`
}

func (as OwnerController) ChangePassword(ctx *gin.Context) {
	var req OwnerChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	user, err := as.authHandler.ChangePassword(ctx, "Owner", req.OldPassword, req.NewPassword, req.ConfirmPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type OwnerForgotPasswordRequest struct {
	Key             string `json:"key" validate:"required,max=255"`
	EmailTemplate   string `json:"emailTemplate"`
	ContactTemplate string `json:"contactTemplate"`
}

func (c *OwnerController) ForgotPassword(ctx *gin.Context) {
	link, err := c.ForgotPasswordResetLink(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"link": link})
}

func (c *OwnerController) ForgotPasswordResetLink(ctx *gin.Context) (string, error) {
	var req OwnerForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return "", err
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return "", err
	}
	return c.authHandler.ForgotPasswordResetLink(ctx, "Owner", req.Key, req.EmailTemplate, req.ContactTemplate)
}

func (c *OwnerController) Create(ctx *gin.Context) {
	var req OwnerStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := c.authHandler.Create(ctx, "Owner",
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

type OwnerNewPasswordRequest struct {
	PreviousPassword string `json:"previousPassword" validate:"required,min=8,max=255"`
	NewPassword      string `json:"newPassword" validate:"required,min=8,max=255"`
	ConfirmPassword  string `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=NewPassword"`
}

func (c *OwnerController) NewPassword(ctx *gin.Context) {
	var req OwnerChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}
	user, err := c.authHandler.NewPassword(ctx, "Owner", req.OldPassword, req.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *OwnerController) SkipVerification(ctx *gin.Context) {
	user, err := c.authHandler.SkipVerification(ctx, "Owner")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type OwnerSendEmailVerificationRequest struct {
	EmailTemplate string `json:"emailTemplate" validate:"required"`
}

func (c *OwnerController) SendEmailVerification(ctx *gin.Context) {
	var req OwnerSendEmailVerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	if err := c.authHandler.SendEmailVerification(ctx, "Owner", req.EmailTemplate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Email verification sent successfully. Please check your inbox or spam folder"})
}

type OwnerVerifyEmailRequest struct {
	Otp string `json:"otp" validate:"required,len=6"`
}

func (c *OwnerController) VerifyEmail(ctx *gin.Context) {
	var req OwnerVerifyEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, err := c.authHandler.VerifyEmail(ctx, "Owner", req.Otp)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type OwnerSendContactNumberVerificationRequest struct {
	ContactTemplate string `json:"contactTemplate" validate:"required"`
}

func (c *OwnerController) SendContactNumberVerification(ctx *gin.Context) {
	var req OwnerSendContactNumberVerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	if err := c.authHandler.SendContactNumberVerification(ctx, "Owner", req.ContactTemplate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Contact number verification OTP sent successfully"})
}

type OwnerVerifyContactNumberRequest struct {
	Otp string `json:"otp" validate:"required,len=6"`
}

func (c *OwnerController) VerifyContactNumber(ctx *gin.Context) {
	var req OwnerVerifyContactNumberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, err := c.authHandler.VerifyContactNumber(ctx, "Owner", req.Otp)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *OwnerController) ProfilePicture(ctx *gin.Context) {
	var req MediaStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, err := c.authHandler.ProfilePicture(ctx, "Owner", req.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type OwnerAccountSettingRequest struct {
	BirthDate        time.Time `json:"birthDate" validate:"required"`
	FirstName        string    `json:"firstName" validate:"required,max=255"`
	MiddleName       string    `json:"middleName" validate:"required,max=255"`
	LastName         string    `json:"lastName" validate:"required,max=255"`
	Description      string    `json:"description" validate:"max=2048"`
	PermanentAddress string    `json:"permanentAddress" validate:"required,max=500"`
}

func (c *OwnerController) ProfileAccountSetting(ctx *gin.Context) {
	var req OwnerAccountSettingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, err := c.authHandler.ProfileAccountSetting(ctx, "Owner",
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

type OwnerChangeEmailRequest struct {
	Password string `json:"password" validate:"required,min=8,max=255"`
	Email    string `json:"email" validate:"required,email"`
}

func (c *OwnerController) ProfileChangeEmail(ctx *gin.Context) {
	var req OwnerChangeEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, err := c.authHandler.ProfileChangeEmail(ctx, "Owner", req.Password, req.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type OwnerChangeContactNumberRequest struct {
	Password      string `json:"password" validate:"required,min=8,max=255"`
	ContactNumber string `json:"contactNumber" validate:"required,min=10,max=15"`
}

func (c *OwnerController) ProfileChangeContactNumber(ctx *gin.Context) {
	var req OwnerChangeContactNumberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	user, err := c.authHandler.ProfileChangeContactNumber(ctx, "Owner", req.Password, req.ContactNumber)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

type OwnerChangeUsernameRequest struct {
	Password string `json:"password" validate:"required,min=8,max=255"`
	Username string `json:"username" validate:"required,min=5,max=255"`
}

func (c *OwnerController) ProfileChangeUsername(ctx *gin.Context) {
	var req OwnerChangeUsernameRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, err := c.authHandler.ProfileChangeContactNumber(ctx, "Owner", req.Password, req.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
