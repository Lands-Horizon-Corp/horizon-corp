package controllers

import (
	"net/http"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type EmployeeController struct {
	authHandler     *handlers.AuthHandler
	footstepHandler *handlers.FootstepHandler
}

func NewEmployeeController(
	authHandler *handlers.AuthHandler,
	footstepHandler *handlers.FootstepHandler,
) *EmployeeController {
	return &EmployeeController{
		authHandler:     authHandler,
		footstepHandler: footstepHandler,
	}
}

// GET: /api/v1/employee
// Retrieve employees with optional filtering for pagination or no pagination. Results can be converted to records.
//
//	Employee: Can retrieve all employees if employee  status is verified
//	Employee: not allowed
//	Member: not allowed
func (c *EmployeeController) Index(ctx *gin.Context) {

}

// GET: /api/v1/employee/:id
//
//      Employee: if employee  status is verified
//      Employee: not allowed
//      Owner: not allowed
//      Member: not allowed

func (c *EmployeeController) Show(ctx *gin.Context) {
	// if id or email or
}

// POST: /api/v1/employee
//
//	Employee: Can create employee and automatically verified and also if employee and the status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
//	public allowed
type EmployeeStoreRequest struct {
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

func (c *EmployeeController) Store(ctx *gin.Context) {
	c.Create(ctx)
}

// PUT: /api/v1/employee/:id
//
//	Employee: Can change status of employee but only if employee and the status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
func (c *EmployeeController) Update(ctx *gin.Context) {

}

// DELETE:/api/v1/employee/:id
// Verifiy employee
//
//	Employee: only if employee  status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
func (c *EmployeeController) Destroy(ctx *gin.Context) {

}

type EmployeeChangePasswordRequest struct {
	OldPassword     string `json:"old_password" validate:"required,min=8,max=255"`
	NewPassword     string `json:"new_password" validate:"required,min=8,max=255"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8,max=255"`
}

func (as EmployeeController) ChangePassword(ctx *gin.Context) {
	var req EmployeeChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	user, err := as.authHandler.ChangePassword(ctx, "Employee", req.OldPassword, req.NewPassword, req.ConfirmPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type EmployeeForgotPasswordRequest struct {
	Key             string `json:"key" validate:"required,max=255"`
	EmailTemplate   string `json:"emailTemplate"`
	ContactTemplate string `json:"contactTemplate"`
}

func (c *EmployeeController) ForgotPassword(ctx *gin.Context) {
	link, err := c.ForgotPasswordResetLink(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"link": link})
}

func (c *EmployeeController) ForgotPasswordResetLink(ctx *gin.Context) (string, error) {
	var req EmployeeForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return "", err
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return "", err
	}
	return c.authHandler.ForgotPasswordResetLink(ctx, "Employee", req.Key, req.EmailTemplate, req.ContactTemplate)
}

func (c *EmployeeController) Create(ctx *gin.Context) {
	var req EmployeeStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := c.authHandler.Create(ctx, "Employee",
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

type EmployeeNewPasswordRequest struct {
	PreviousPassword string `json:"previousPassword" validate:"required,min=8,max=255"`
	NewPassword      string `json:"newPassword" validate:"required,min=8,max=255"`
	ConfirmPassword  string `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=NewPassword"`
}

func (c *EmployeeController) NewPassword(ctx *gin.Context) {
	var req EmployeeChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}
	user, err := c.authHandler.NewPassword(ctx, "Employee", req.OldPassword, req.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *EmployeeController) SkipVerification(ctx *gin.Context) {
	user, err := c.authHandler.SkipVerification(ctx, "Employee")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type EmployeeSendEmailVerificationRequest struct {
	EmailTemplate string `json:"emailTemplate" validate:"required"`
}

func (c *EmployeeController) SendEmailVerification(ctx *gin.Context) {
	var req EmployeeSendEmailVerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	if err := c.authHandler.SendEmailVerification(ctx, "Employee", req.EmailTemplate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Email verification sent successfully. Please check your inbox or spam folder"})
}

type EmployeeVerifyEmailRequest struct {
	Otp string `json:"otp" validate:"required,len=6"`
}

func (c *EmployeeController) VerifyEmail(ctx *gin.Context) {
	var req EmployeeVerifyEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, err := c.authHandler.VerifyEmail(ctx, "Employee", req.Otp)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type EmployeeSendContactNumberVerificationRequest struct {
	ContactTemplate string `json:"contactTemplate" validate:"required"`
}

func (c *EmployeeController) SendContactNumberVerification(ctx *gin.Context) {
	var req EmployeeSendContactNumberVerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	if err := c.authHandler.SendContactNumberVerification(ctx, "Employee", req.ContactTemplate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Contact number verification OTP sent successfully"})
}

type EmployeeVerifyContactNumberRequest struct {
	Otp string `json:"otp" validate:"required,len=6"`
}

func (c *EmployeeController) VerifyContactNumber(ctx *gin.Context) {
	var req EmployeeVerifyContactNumberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, err := c.authHandler.VerifyContactNumber(ctx, "Employee", req.Otp)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *EmployeeController) ProfilePicture(ctx *gin.Context) {
	var req MediaStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, err := c.authHandler.ProfilePicture(ctx, "Employee", req.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type EmployeeAccountSettingRequest struct {
	BirthDate        time.Time `json:"birthDate" validate:"required"`
	FirstName        string    `json:"firstName" validate:"required,max=255"`
	MiddleName       string    `json:"middleName" validate:"required,max=255"`
	LastName         string    `json:"lastName" validate:"required,max=255"`
	Description      string    `json:"description" validate:"max=2048"`
	PermanentAddress string    `json:"permanentAddress" validate:"required,max=500"`
}

func (c *EmployeeController) ProfileAccountSetting(ctx *gin.Context) {
	var req EmployeeAccountSettingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, err := c.authHandler.ProfileAccountSetting(ctx, "Employee",
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

type EmployeeChangeEmailRequest struct {
	Password string `json:"password" validate:"required,min=8,max=255"`
	Email    string `json:"email" validate:"required,email"`
}

func (c *EmployeeController) ProfileChangeEmail(ctx *gin.Context) {
	var req EmployeeChangeEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, err := c.authHandler.ProfileChangeEmail(ctx, "Employee", req.Password, req.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type EmployeeChangeContactNumberRequest struct {
	Password      string `json:"password" validate:"required,min=8,max=255"`
	ContactNumber string `json:"contactNumber" validate:"required,min=10,max=15"`
}

func (c *EmployeeController) ProfileChangeContactNumber(ctx *gin.Context) {
	var req EmployeeChangeContactNumberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	user, err := c.authHandler.ProfileChangeContactNumber(ctx, "Employee", req.Password, req.ContactNumber)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

type EmployeeChangeUsernameRequest struct {
	Password string `json:"password" validate:"required,min=8,max=255"`
	Username string `json:"username" validate:"required,min=5,max=255"`
}

func (c *EmployeeController) ProfileChangeUsername(ctx *gin.Context) {
	var req EmployeeChangeUsernameRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	user, err := c.authHandler.ProfileChangeContactNumber(ctx, "Employee", req.Password, req.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
