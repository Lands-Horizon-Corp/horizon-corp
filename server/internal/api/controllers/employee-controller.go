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

type EmployeeController struct {
	repository    *models.ModelRepository
	transformer   *models.ModelTransformer
	footstep      *handlers.FootstepHandler
	currentUser   *handlers.CurrentUser
	otpService    *providers.OTPService
	tokenProvider *providers.TokenService
	cfg           *config.AppConfig
	helpers       *helpers.HelpersFunction
	smsProvder    *providers.SMSService
	emailProvider *providers.EmailService
}

func NewEmployeeController(
	repository *models.ModelRepository,
	transformer *models.ModelTransformer,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
	otpService *providers.OTPService,
	tokenProvider *providers.TokenService,
	cfg *config.AppConfig,
	helpers *helpers.HelpersFunction,
	smsProvder *providers.SMSService,
	emailProvider *providers.EmailService,
) *EmployeeController {
	return &EmployeeController{
		repository:    repository,
		transformer:   transformer,
		footstep:      footstep,
		currentUser:   currentUser,
		otpService:    otpService,
		tokenProvider: tokenProvider,
		cfg:           cfg,
		helpers:       helpers,
		smsProvder:    smsProvder,
		emailProvider: emailProvider,
	}
}

// GET :/
// Admin: Can get anything but must give company id
// Owner: All employee but must be only his company and its branches
// Employee: only employee on current branch
func (c *EmployeeController) Index(ctx *gin.Context) {

}

// GET :/id
// Admin: Can get anything but must give company id
// Owner: All employee but must be only his company and its branches
// Employee: only employee on current branch
func (c *EmployeeController) Show(ctx *gin.Context) {

}

// POST: /
// Admin: Can create employee but must assign company and branch
// Owner: can create but must assign branch
// Employee: not allowed
// Member: not allowed
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

	Longitude *float64 `json:"longitude" validate:"omitempty,numeric"`
	Latitude  *float64 `json:"latitude" validate:"omitempty,numeric"`

	MediaID  *uuid.UUID `json:"mediaId" validate:"omitempty,uuid"`
	BranchID *uuid.UUID `json:"branchId" validate:"omitempty,uuid"`
	RoleID   *uuid.UUID `json:"roleId" validate:"omitempty,uuid"`
	GenderID *uuid.UUID `json:"genderId" validate:"omitempty,uuid"`

	EmailTemplate   string `json:"emailTemplate" validate:"required"`
	ContactTemplate string `json:"contactTemplate" validate:"required"`
}

func (c *EmployeeController) Store(ctx *gin.Context) {
	c.Create(ctx)
}

// PUT: /
// Admin: Can create employee but must assign company and branch
// Owner: can create but must assign branch
// Employee: not allowed
// Member: not allowed
func (c *EmployeeController) Update(ctx *gin.Context) {

}

// DELETE: /:id
// Admin: allowed
// Owner: allowed
// Employee: not allowed
// Member: not allowed
func (c *EmployeeController) Destroy(ctx *gin.Context) {

}

type EmployeeChangePasswordRequest struct {
	OldPassword     string `json:"old_password" validate:"required,min=8,max=255"`
	NewPassword     string `json:"new_password" validate:"required,min=8,max=255"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8,max=255"`
}

func (as EmployeeController) ChangePassword(ctx *gin.Context) {
	var req *EmployeeChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: JSON binding error: %v", err)})
		return
	}
	employee, err := as.currentUser.Employee(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	updated, err := as.repository.EmployeeChangePassword(
		employee.ID.String(),
		req.OldPassword,
		req.NewPassword,
		as.helpers.GetPreload(ctx)...,
	)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusCreated, as.transformer.EmployeeToResource(updated))
}

func (c *EmployeeController) ForgotPassword(ctx *gin.Context) {
	link := c.ForgotPasswordResetLink(ctx)
	ctx.JSON(http.StatusBadRequest, gin.H{"link": link})
}

type EmployeeForgotPasswordRequest struct {
	Key             string `json:"key" validate:"required,max=255"`
	EmailTemplate   string `json:"emailTemplate"`
	ContactTemplate string `json:"contactTemplate"`
}

func (c *EmployeeController) ForgotPasswordResetLink(ctx *gin.Context) *string {
	var req EmployeeForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return nil
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": err.Error()})
		return nil
	}
	user, err := c.repository.EmployeeSearch(req.Key)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("SignIn: User not found: %v", err)})
		return nil
	}

	token, err := c.tokenProvider.GenerateUserToken(providers.UserClaims{
		ID:          user.ID.String(),
		AccountType: "Employee",
		UserStatus:  user.Status,
	}, time.Minute*10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return nil
	}

	resetLink := fmt.Sprintf("%s/auth/password-reset/%s", c.cfg.AppClientUrl, *token)
	keyType := c.helpers.GetKeyType(req.Key)
	switch keyType {
	case "contact":
		contactReq := providers.SMSRequest{
			To:   req.Key,
			Body: req.ContactTemplate,
			Vars: &map[string]string{
				"name":      fmt.Sprintf("%s %s", user.FirstName, user.LastName),
				"eventLink": resetLink,
			},
		}
		if err := c.smsProvder.SendSMS(contactReq); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("ForgotPassword: SMS sending error %v", err)})
		}
		return &resetLink
	case "email":
		emailReq := providers.EmailRequest{
			To:      req.Key,
			Subject: "ECOOP: Change Password Request",
			Body:    req.ContactTemplate,
			Vars: &map[string]string{
				"name":      fmt.Sprintf("%s %s", user.FirstName, user.LastName),
				"eventLink": resetLink,
			},
		}
		if err := c.emailProvider.SendEmail(emailReq); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("ForgotPassword: Email sending error: %v", err)})
			return nil
		}
		return &resetLink
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("ForgotPassword: Invalid key type: %s", keyType)})
		return nil
	}
}

func (c *EmployeeController) Create(ctx *gin.Context) *models.Employee {
	var req EmployeeStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return nil
	}
	if validator.New().Struct(req) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": req})
		return nil
	}
	preloads := c.helpers.GetPreload(ctx)
	employee, err := c.repository.EmployeeCreate(&models.Employee{
		FirstName:          req.FirstName,
		LastName:           req.LastName,
		MiddleName:         req.MiddleName,
		PermanentAddress:   req.PermanentAddress,
		Description:        req.Description,
		BirthDate:          req.BirthDate,
		Username:           req.Username,
		Email:              req.Email,
		Password:           req.Password,
		ContactNumber:      req.ContactNumber,
		IsEmailVerified:    false,
		IsContactVerified:  false,
		IsSkipVerification: false,
		Status:             providers.NotAllowedStatus,
		Longitude:          req.Longitude,
		Latitude:           req.Latitude,

		MediaID:  req.MediaID,
		BranchID: req.BranchID,
		RoleID:   req.RoleID,
		GenderID: req.GenderID,
	}, preloads...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create employee", "details": err.Error()})
		return nil
	}
	if err := c.otpService.SendEmailOTP(providers.OTPMessage{
		AccountType: "Employee", ID: employee.ID.String(),
		MediumType: "email",
		Reference:  "email-verification",
	}, providers.EmailRequest{
		To:      req.Email,
		Subject: "ECOOP: Email Verification",
		Body:    req.EmailTemplate,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return nil
	}
	if err := c.otpService.SendContactNumberOTP(providers.OTPMessage{
		AccountType: "Employee", ID: employee.ID.String(),
		MediumType: "sms",
		Reference:  "sms-verification",
	}, providers.SMSRequest{
		To:   req.ContactNumber,
		Body: req.ContactTemplate,
		Vars: &map[string]string{
			"name": fmt.Sprintf("%s %s", req.FirstName, req.LastName),
		},
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return nil
	}
	ctx.JSON(http.StatusCreated, c.transformer.EmployeeToResource(employee))
	return employee
}

type EmployeeNewPasswordRequest struct {
	PreviousPassword string `json:"previousPassword" validate:"required,min=8,max=255"`
	NewPassword      string `json:"newPassword" validate:"required,min=8,max=255"`
	ConfirmPassword  string `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=NewPassword"`
}

func (c *EmployeeController) NewPassword(ctx *gin.Context) {
	var req *EmployeeChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: JSON binding error: %v", err)})
		return
	}
	if validator.New().Struct(req) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": req})
		return
	}
	employee, err := c.currentUser.Employee(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	updatedEmployee, err := c.repository.EmployeeChangePassword(
		employee.ID.String(), req.OldPassword, req.NewPassword,
		c.helpers.GetPreload(ctx)...,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.EmployeeToResource(updatedEmployee))
}
