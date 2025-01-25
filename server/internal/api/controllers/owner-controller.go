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

type OwnerController struct {
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

func NewOwnerController(
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
) *OwnerController {
	return &OwnerController{
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

	owner, err := as.currentUser.Owner(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	updated, err := as.repository.OwnerChangePassword(
		owner.ID.String(),
		req.OldPassword,
		req.NewPassword,
		as.helpers.GetPreload(ctx)...,
	)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to update password"})
		return
	}
	_, err = as.footstep.Create(ctx, "Owner", "ChangePassword", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}

	ctx.JSON(http.StatusOK, as.transformer.OwnerToResource(updated))
}

type OwnerForgotPasswordRequest struct {
	Key             string `json:"key" validate:"required,max=255"`
	EmailTemplate   string `json:"emailTemplate"`
	ContactTemplate string `json:"contactTemplate"`
}

func (c *OwnerController) ForgotPassword(ctx *gin.Context) {
	link := c.ForgotPasswordResetLink(ctx)
	ctx.JSON(http.StatusBadRequest, gin.H{"link": link})
}

func (c *OwnerController) ForgotPasswordResetLink(ctx *gin.Context) *string {
	var req OwnerForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return nil
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return nil
	}

	user, err := c.repository.OwnerSearch(req.Key)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return nil
	}

	token, err := c.tokenProvider.GenerateUserToken(providers.UserClaims{
		ID:          user.ID.String(),
		AccountType: "Owner",
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

func (c *OwnerController) Create(ctx *gin.Context) *models.Owner {
	var req OwnerStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return nil
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return nil
	}

	preloads := c.helpers.GetPreload(ctx)
	owner, err := c.repository.OwnerCreate(&models.Owner{
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create owner", "details": err.Error()})
		return nil
	}

	if err := c.otpService.SendEmailOTP(providers.OTPMessage{
		AccountType: "Owner",
		ID:          owner.ID.String(),
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
		AccountType: "Owner",
		ID:          owner.ID.String(),
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

	ctx.JSON(http.StatusCreated, c.transformer.OwnerToResource(owner))
	return owner
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	owner, err := c.currentUser.Owner(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	updatedOwner, err := c.repository.OwnerChangePassword(
		owner.ID.String(), req.OldPassword, req.NewPassword,
		c.helpers.GetPreload(ctx)...,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	_, err = c.footstep.Create(ctx, "Owner", "NewPassword", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}

	ctx.JSON(http.StatusOK, c.transformer.OwnerToResource(updatedOwner))
}

func (c *OwnerController) SkipVerification(ctx *gin.Context) {
	owner, err := c.currentUser.Owner(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	updatedOwner, err := c.repository.OwnerUpdateByID(owner.ID.String(), &models.Owner{
		IsSkipVerification: true,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update owner details"})
		return
	}
	_, err = c.footstep.Create(ctx, "Owner", "SkipVerification", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.OwnerToResource(updatedOwner))
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

	owner, err := c.currentUser.Owner(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	otpMessage := providers.OTPMessage{
		AccountType: "Owner",
		ID:          owner.ID.String(),
		MediumType:  providers.Email,
		Reference:   "send-email-verification",
	}
	emailRequest := providers.EmailRequest{
		To:      owner.Email,
		Subject: "ECOOP: Email Verification",
		Body:    req.EmailTemplate,
	}

	if err := c.otpService.SendEmailOTP(otpMessage, emailRequest); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email verification"})
		return
	}
	_, err = c.footstep.Create(ctx, "Owner", "SendEmailVerification", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
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

	owner, err := c.currentUser.Owner(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	otpMessage := providers.OTPMessage{
		AccountType: "Owner",
		ID:          owner.ID.String(),
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

	updatedOwner, err := c.repository.OwnerUpdateByID(owner.ID.String(), &models.Owner{
		IsEmailVerified: true,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update owner details"})
		return
	}
	_, err = c.footstep.Create(ctx, "Owner", "VerifyEmail", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}

	ctx.JSON(http.StatusOK, c.transformer.OwnerToResource(updatedOwner))
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

	owner, err := c.currentUser.Owner(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	otpMessage := providers.OTPMessage{
		AccountType: "Owner",
		ID:          owner.ID.String(),
		MediumType:  providers.Email,
		Reference:   "send-contact-number-verification",
	}
	contactReq := providers.SMSRequest{
		To:   owner.ContactNumber,
		Body: req.ContactTemplate,
		Vars: &map[string]string{
			"name": fmt.Sprintf("%s %s", owner.FirstName, owner.LastName),
		},
	}

	if err := c.otpService.SendContactNumberOTP(otpMessage, contactReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification OTP"})
		return
	}
	_, err = c.footstep.Create(ctx, "Owner", "SendContactNumberVerification", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
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

	owner, err := c.currentUser.Owner(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	otpMessage := providers.OTPMessage{
		AccountType: "Owner",
		ID:          owner.ID.String(),
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

	updatedOwner, err := c.repository.OwnerUpdateByID(owner.ID.String(), &models.Owner{
		IsContactVerified: true,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update owner details"})
		return
	}
	_, err = c.footstep.Create(ctx, "Owner", "VerifyContactNumber", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}

	ctx.JSON(http.StatusOK, c.transformer.OwnerToResource(updatedOwner))
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

	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}
	owner, err := c.currentUser.Owner(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	updatedOwner, err := c.repository.OwnerUpdateByID(owner.ID.String(), &models.Owner{
		MediaID: req.ID,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update owner details"})
		return
	}
	_, err = c.footstep.Create(ctx, "Owner", "ProfilePicture", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.OwnerToResource(updatedOwner))
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

	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}
	owner, err := c.currentUser.Owner(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	updatedOwner, err := c.repository.OwnerUpdateByID(owner.ID.String(), &models.Owner{
		BirthDate:        req.BirthDate,
		MiddleName:       req.MiddleName,
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		Description:      req.Description,
		PermanentAddress: req.PermanentAddress,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update owner details"})
		return
	}
	_, err = c.footstep.Create(ctx, "Owner", "ProfileAccountSetting", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.OwnerToResource(updatedOwner))
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
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}
	owner, err := c.currentUser.Owner(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	if !c.repository.OwnerVerifyPassword(owner.ID.String(), req.Password) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong password"})
		return
	}
	updatedOwner, err := c.repository.OwnerUpdateByID(owner.ID.String(), &models.Owner{
		Email:           req.Email,
		IsEmailVerified: false,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update owner details"})
		return
	}
	_, err = c.footstep.Create(ctx, "Owner", "ProfileChangeEmail", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.OwnerToResource(updatedOwner))
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
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}
	owner, err := c.currentUser.Owner(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	if !c.repository.OwnerVerifyPassword(owner.ID.String(), req.Password) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong password"})
		return
	}
	updatedOwner, err := c.repository.OwnerUpdateByID(owner.ID.String(), &models.Owner{
		ContactNumber:     req.ContactNumber,
		IsContactVerified: false,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update owner details"})
		return
	}
	_, err = c.footstep.Create(ctx, "Owner", "ProfileChangeContactNumber", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.OwnerToResource(updatedOwner))

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
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return
	}
	owner, err := c.currentUser.Owner(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	if !c.repository.OwnerVerifyPassword(owner.ID.String(), req.Password) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong password"})
		return
	}
	updatedOwner, err := c.repository.OwnerUpdateByID(owner.ID.String(), &models.Owner{
		Username: req.Username,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update owner details"})
		return
	}
	_, err = c.footstep.Create(ctx, "Owner", "ProfileChangeUsername", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.OwnerToResource(updatedOwner))

}
