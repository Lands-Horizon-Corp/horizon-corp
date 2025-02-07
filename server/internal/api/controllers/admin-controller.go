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

type AdminController struct {
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

func NewAdminController(
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
) *AdminController {
	return &AdminController{
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

	admin, err := as.currentUser.Admin(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	updated, err := as.repository.AdminChangePassword(
		admin.ID.String(),
		req.OldPassword,
		req.NewPassword,
		as.helpers.GetPreload(ctx)...,
	)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to update password"})
		return
	}
	_, err = as.footstep.Create(ctx, "Admin", "ChangePassword", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}

	ctx.JSON(http.StatusOK, as.transformer.AdminToResource(updated))
}

type AdminForgotPasswordRequest struct {
	Key             string `json:"key" validate:"required,max=255"`
	EmailTemplate   string `json:"emailTemplate"`
	ContactTemplate string `json:"contactTemplate"`
}

func (c *AdminController) ForgotPassword(ctx *gin.Context) {
	link := c.ForgotPasswordResetLink(ctx)
	ctx.JSON(http.StatusBadRequest, gin.H{"link": link})
}

func (c *AdminController) ForgotPasswordResetLink(ctx *gin.Context) *string {
	var req AdminForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return nil
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return nil
	}

	user, err := c.repository.AdminSearch(req.Key)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return nil
	}

	token, err := c.tokenProvider.GenerateUserToken(providers.UserClaims{
		ID:          user.ID.String(),
		AccountType: "Admin",
		UserStatus:  user.Status,
	}, time.Minute*10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate reset token"})
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
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send SMS for password reset"})
			return nil
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
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email for password reset"})
			return nil
		}
		return &resetLink
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key type"})
		return nil
	}
}

func (c *AdminController) Create(ctx *gin.Context) *models.Admin {
	var req AdminStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return nil
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return nil
	}

	preloads := c.helpers.GetPreload(ctx)
	admin, err := c.repository.AdminCreate(&models.Admin{
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
		MediaID:            req.MediaID,
		RoleID:             req.RoleID,
		GenderID:           req.GenderID,
	}, preloads...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create admin", "details": err.Error()})
		return nil
	}

	if err := c.otpService.SendEmailOTP(providers.OTPMessage{
		AccountType: "Admin",
		ID:          admin.ID.String(),
		MediumType:  "email",
		Reference:   "email-verification",
	}, providers.EmailRequest{
		To:      req.Email,
		Subject: "ECOOP: Email Verification",
		Body:    req.EmailTemplate,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email verification"})
		return nil
	}

	if err := c.otpService.SendContactNumberOTP(providers.OTPMessage{
		AccountType: "Admin",
		ID:          admin.ID.String(),
		MediumType:  "sms",
		Reference:   "sms-verification",
	}, providers.SMSRequest{
		To:   req.ContactNumber,
		Body: req.ContactTemplate,
		Vars: &map[string]string{
			"name": fmt.Sprintf("%s %s", req.FirstName, req.LastName),
		},
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send contact number verification"})
		return nil
	}

	ctx.JSON(http.StatusCreated, c.transformer.AdminToResource(admin))
	return admin
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	admin, err := c.currentUser.Admin(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	updatedAdmin, err := c.repository.AdminChangePassword(
		admin.ID.String(), req.OldPassword, req.NewPassword,
		c.helpers.GetPreload(ctx)...,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	_, err = c.footstep.Create(ctx, "Admin", "NewPassword", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}

	ctx.JSON(http.StatusOK, c.transformer.AdminToResource(updatedAdmin))
}

func (c *AdminController) SkipVerification(ctx *gin.Context) {
	admin, err := c.currentUser.Admin(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	updatedAdmin, err := c.repository.AdminUpdateByID(admin.ID.String(), &models.Admin{
		IsSkipVerification: true,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update admin details"})
		return
	}
	_, err = c.footstep.Create(ctx, "Admin", "SkipVerification", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.AdminToResource(updatedAdmin))
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

	admin, err := c.currentUser.Admin(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	otpMessage := providers.OTPMessage{
		AccountType: "Admin",
		ID:          admin.ID.String(),
		MediumType:  providers.Email,
		Reference:   "send-email-verification",
	}
	emailRequest := providers.EmailRequest{
		To:      admin.Email,
		Subject: "ECOOP: Email Verification",
		Body:    req.EmailTemplate,
	}

	if err := c.otpService.SendEmailOTP(otpMessage, emailRequest); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email verification"})
		return
	}
	_, err = c.footstep.Create(ctx, "Admin", "SendEmailVerification", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
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

	admin, err := c.currentUser.Admin(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	otpMessage := providers.OTPMessage{
		AccountType: "Admin",
		ID:          admin.ID.String(),
		MediumType:  providers.Email,
		Reference:   "send-email-verification",
	}

	isValid, err := c.otpService.ValidateOTP(otpMessage, req.Otp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate OTP"})
		return
	}
	if !isValid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
		return
	}

	updatedAdmin, err := c.repository.AdminUpdateByID(admin.ID.String(), &models.Admin{
		IsEmailVerified: true,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update admin details"})
		return
	}
	_, err = c.footstep.Create(ctx, "Admin", "VerifyEmail", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}

	ctx.JSON(http.StatusOK, c.transformer.AdminToResource(updatedAdmin))
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

	admin, err := c.currentUser.Admin(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	otpMessage := providers.OTPMessage{
		AccountType: "Admin",
		ID:          admin.ID.String(),
		MediumType:  providers.Email,
		Reference:   "send-contact-number-verification",
	}
	contactReq := providers.SMSRequest{
		To:   admin.ContactNumber,
		Body: req.ContactTemplate,
		Vars: &map[string]string{
			"name": fmt.Sprintf("%s %s", admin.FirstName, admin.LastName),
		},
	}

	if err := c.otpService.SendContactNumberOTP(otpMessage, contactReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification OTP"})
		return
	}
	_, err = c.footstep.Create(ctx, "Admin", "SendContactNumberVerification", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
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

	admin, err := c.currentUser.Admin(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	otpMessage := providers.OTPMessage{
		AccountType: "Admin",
		ID:          admin.ID.String(),
		MediumType:  providers.Email,
		Reference:   "send-contact-number-verification",
	}

	isValid, err := c.otpService.ValidateOTP(otpMessage, req.Otp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "OTP validation error"})
		return
	}
	if !isValid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
		return
	}

	updatedAdmin, err := c.repository.AdminUpdateByID(admin.ID.String(), &models.Admin{
		IsContactVerified: true,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update admin details"})
		return
	}
	_, err = c.footstep.Create(ctx, "Admin", "VerifyContactNumber", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}

	ctx.JSON(http.StatusOK, c.transformer.AdminToResource(updatedAdmin))
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

	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}
	admin, err := c.currentUser.Admin(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	updatedAdmin, err := c.repository.AdminUpdateByID(admin.ID.String(), &models.Admin{
		MediaID: req.ID,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update admin details"})
		return
	}
	_, err = c.footstep.Create(ctx, "Admin", "ProfilePicture", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.AdminToResource(updatedAdmin))
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

	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}
	admin, err := c.currentUser.Admin(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	updatedAdmin, err := c.repository.AdminUpdateByID(admin.ID.String(), &models.Admin{
		BirthDate:        req.BirthDate,
		MiddleName:       req.MiddleName,
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		Description:      req.Description,
		PermanentAddress: req.PermanentAddress,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update admin details"})
		return
	}
	_, err = c.footstep.Create(ctx, "Admin", "ProfileAccountSetting", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.AdminToResource(updatedAdmin))
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
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}
	admin, err := c.currentUser.Admin(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	if !c.repository.AdminVerifyPassword(admin.ID.String(), req.Password) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong password"})
		return
	}
	updatedAdmin, err := c.repository.AdminUpdateByID(admin.ID.String(), &models.Admin{
		Email:           req.Email,
		IsEmailVerified: false,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update admin details"})
		return
	}
	_, err = c.footstep.Create(ctx, "Admin", "ProfileChangeEmail", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.AdminToResource(updatedAdmin))
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
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}
	admin, err := c.currentUser.Admin(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	if !c.repository.AdminVerifyPassword(admin.ID.String(), req.Password) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong password"})
		return
	}
	updatedAdmin, err := c.repository.AdminUpdateByID(admin.ID.String(), &models.Admin{
		ContactNumber:     req.ContactNumber,
		IsContactVerified: false,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update admin details"})
		return
	}
	_, err = c.footstep.Create(ctx, "Admin", "ProfileChangeContactNumber", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.AdminToResource(updatedAdmin))

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
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}
	admin, err := c.currentUser.Admin(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	if !c.repository.AdminVerifyPassword(admin.ID.String(), req.Password) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong password"})
		return
	}
	updatedAdmin, err := c.repository.AdminUpdateByID(admin.ID.String(), &models.Admin{
		Username: req.Username,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update admin details"})
		return
	}
	_, err = c.footstep.Create(ctx, "Admin", "ProfileChangeUsername", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.AdminToResource(updatedAdmin))

}
