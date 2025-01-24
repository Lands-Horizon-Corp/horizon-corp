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

type MemberController struct {
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

func NewMemberController(
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
) *MemberController {
	return &MemberController{
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

// GET: /api/v1/member
// Retrieve members with optional filtering for pagination or no pagination. Results can be converted to records.
//
//	Member: Can retrieve all members if member  status is verified
//	Employee: not allowed
//	Member: not allowed
func (c *MemberController) Index(ctx *gin.Context) {

}

// GET: /api/v1/member/:id
//
//      Member: if member  status is verified
//      Employee: not allowed
//      Owner: not allowed
//      Member: not allowed

func (c *MemberController) Show(ctx *gin.Context) {
	// if id or email or
}

// POST: /api/v1/member
//
//	Member: Can create member and automatically verified and also if member and the status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
//	public allowed
type MemberStoreRequest struct {
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

func (c *MemberController) Store(ctx *gin.Context) {
	c.Create(ctx)
}

// PUT: /api/v1/member/:id
//
//	Member: Can change status of member but only if member and the status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
func (c *MemberController) Update(ctx *gin.Context) {

}

// DELETE:/api/v1/member/:id
// Verifiy member
//
//	Member: only if member  status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
func (c *MemberController) Destroy(ctx *gin.Context) {

}

type MemberChangePasswordRequest struct {
	OldPassword     string `json:"old_password" validate:"required,min=8,max=255"`
	NewPassword     string `json:"new_password" validate:"required,min=8,max=255"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8,max=255"`
}

func (as MemberController) ChangePassword(ctx *gin.Context) {
	var req MemberChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	member, err := as.currentUser.Member(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	updated, err := as.repository.MemberChangePassword(
		member.ID.String(),
		req.OldPassword,
		req.NewPassword,
		as.helpers.GetPreload(ctx)...,
	)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to update password"})
		return
	}

	ctx.JSON(http.StatusOK, as.transformer.MemberToResource(updated))
}

type MemberForgotPasswordRequest struct {
	Key             string `json:"key" validate:"required,max=255"`
	EmailTemplate   string `json:"emailTemplate"`
	ContactTemplate string `json:"contactTemplate"`
}

func (c *MemberController) ForgotPassword(ctx *gin.Context) {
	link := c.ForgotPasswordResetLink(ctx)
	ctx.JSON(http.StatusBadRequest, gin.H{"link": link})
}

func (c *MemberController) ForgotPasswordResetLink(ctx *gin.Context) *string {
	var req MemberForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return nil
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		return nil
	}

	user, err := c.repository.MemberSearch(req.Key)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return nil
	}

	token, err := c.tokenProvider.GenerateUserToken(providers.UserClaims{
		ID:          user.ID.String(),
		AccountType: "Member",
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

func (c *MemberController) Create(ctx *gin.Context) *models.Member {
	var req MemberStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return nil
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return nil
	}

	preloads := c.helpers.GetPreload(ctx)
	member, err := c.repository.MemberCreate(&models.Member{
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create member", "details": err.Error()})
		return nil
	}

	if err := c.otpService.SendEmailOTP(providers.OTPMessage{
		AccountType: "Member",
		ID:          member.ID.String(),
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
		AccountType: "Member",
		ID:          member.ID.String(),
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

	ctx.JSON(http.StatusCreated, c.transformer.MemberToResource(member))
	return member
}

type MemberNewPasswordRequest struct {
	PreviousPassword string `json:"previousPassword" validate:"required,min=8,max=255"`
	NewPassword      string `json:"newPassword" validate:"required,min=8,max=255"`
	ConfirmPassword  string `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=NewPassword"`
}

func (c *MemberController) NewPassword(ctx *gin.Context) {
	var req MemberChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	member, err := c.currentUser.Member(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	updatedMember, err := c.repository.MemberChangePassword(
		member.ID.String(), req.OldPassword, req.NewPassword,
		c.helpers.GetPreload(ctx)...,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	ctx.JSON(http.StatusOK, c.transformer.MemberToResource(updatedMember))
}

func (c *MemberController) SkipVerification(ctx *gin.Context) {
	member, err := c.currentUser.Member(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	updatedMember, err := c.repository.MemberUpdateByID(member.ID.String(), &models.Member{
		IsSkipVerification: true,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update member details"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.MemberToResource(updatedMember))
}

type MemberSendEmailVerificationRequest struct {
	EmailTemplate string `json:"emailTemplate" validate:"required"`
}

func (c *MemberController) SendEmailVerification(ctx *gin.Context) {
	var req MemberSendEmailVerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	member, err := c.currentUser.Member(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	otpMessage := providers.OTPMessage{
		AccountType: "Member",
		ID:          member.ID.String(),
		MediumType:  providers.Email,
		Reference:   "send-email-verification",
	}
	emailRequest := providers.EmailRequest{
		To:      member.Email,
		Subject: "ECOOP: Email Verification",
		Body:    req.EmailTemplate,
	}

	if err := c.otpService.SendEmailOTP(otpMessage, emailRequest); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email verification"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Email verification sent successfully. Please check your inbox or spam folder"})
}

type MemberVerifyEmailRequest struct {
	Otp string `json:"otp" validate:"required,len=6"`
}

func (c *MemberController) VerifyEmail(ctx *gin.Context) {
	var req MemberVerifyEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	member, err := c.currentUser.Member(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	otpMessage := providers.OTPMessage{
		AccountType: "Member",
		ID:          member.ID.String(),
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

	updatedMember, err := c.repository.MemberUpdateByID(member.ID.String(), &models.Member{
		IsEmailVerified: true,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update member details"})
		return
	}

	ctx.JSON(http.StatusOK, c.transformer.MemberToResource(updatedMember))
}

type MemberSendContactNumberVerificationRequest struct {
	ContactTemplate string `json:"contactTemplate" validate:"required"`
}

func (c *MemberController) SendContactNumberVerification(ctx *gin.Context) {
	var req MemberSendContactNumberVerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	member, err := c.currentUser.Member(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	otpMessage := providers.OTPMessage{
		AccountType: "Member",
		ID:          member.ID.String(),
		MediumType:  providers.Email,
		Reference:   "send-contact-number-verification",
	}
	contactReq := providers.SMSRequest{
		To:   member.ContactNumber,
		Body: req.ContactTemplate,
		Vars: &map[string]string{
			"name": fmt.Sprintf("%s %s", member.FirstName, member.LastName),
		},
	}

	if err := c.otpService.SendContactNumberOTP(otpMessage, contactReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification OTP"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Contact number verification OTP sent successfully"})
}

type MemberVerifyContactNumberRequest struct {
	Otp string `json:"otp" validate:"required,len=6"`
}

func (c *MemberController) VerifyContactNumber(ctx *gin.Context) {
	var req MemberVerifyContactNumberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	member, err := c.currentUser.Member(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	otpMessage := providers.OTPMessage{
		AccountType: "Member",
		ID:          member.ID.String(),
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

	updatedMember, err := c.repository.MemberUpdateByID(member.ID.String(), &models.Member{
		IsContactVerified: true,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update member details"})
		return
	}

	ctx.JSON(http.StatusOK, c.transformer.MemberToResource(updatedMember))
}

func (c *MemberController) ProfilePicture(ctx *gin.Context) {
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
	member, err := c.currentUser.Member(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	updatedMember, err := c.repository.MemberUpdateByID(member.ID.String(), &models.Member{
		MediaID: req.ID,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update member details"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.MemberToResource(updatedMember))
}

type MemberAccountSettingRequest struct {
	BirthDate        time.Time `json:"birthDate" validate:"required"`
	FirstName        string    `json:"firstName" validate:"required,max=255"`
	MiddleName       string    `json:"middleName" validate:"required,max=255"`
	LastName         string    `json:"lastName" validate:"required,max=255"`
	Description      string    `json:"description" validate:"max=2048"`
	PermanentAddress string    `json:"permanentAddress" validate:"required,max=500"`
}

func (c *MemberController) ProfileAccountSetting(ctx *gin.Context) {
	var req MemberAccountSettingRequest
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
	member, err := c.currentUser.Member(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	updatedMember, err := c.repository.MemberUpdateByID(member.ID.String(), &models.Member{
		BirthDate:        req.BirthDate,
		MiddleName:       req.MiddleName,
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		Description:      req.Description,
		PermanentAddress: req.PermanentAddress,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update member details"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.MemberToResource(updatedMember))
}

type MemberChangeEmailRequest struct {
	Password string `json:"password" validate:"required,min=8,max=255"`
	Email    string `json:"email" validate:"required,email"`
}

func (c *MemberController) ProfileChangeEmail(ctx *gin.Context) {
	var req MemberChangeEmailRequest
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
	member, err := c.currentUser.Member(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	if !c.repository.MemberVerifyPassword(member.ID.String(), req.Password) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong password"})
		return
	}
	updatedMember, err := c.repository.MemberUpdateByID(member.ID.String(), &models.Member{
		Email:           req.Email,
		IsEmailVerified: false,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update member details"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.MemberToResource(updatedMember))
}

type MemberChangeContactNumberRequest struct {
	Password      string `json:"password" validate:"required,min=8,max=255"`
	ContactNumber string `json:"contactNumber" validate:"required,min=10,max=15"`
}

func (c *MemberController) ProfileChangeContactNumber(ctx *gin.Context) {
	var req MemberChangeContactNumberRequest
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
	member, err := c.currentUser.Member(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	if !c.repository.MemberVerifyPassword(member.ID.String(), req.Password) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong password"})
		return
	}
	updatedMember, err := c.repository.MemberUpdateByID(member.ID.String(), &models.Member{
		ContactNumber:     req.ContactNumber,
		IsContactVerified: false,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update member details"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.MemberToResource(updatedMember))

}

type MemberChangeUsernameRequest struct {
	Password string `json:"password" validate:"required,min=8,max=255"`
	Username string `json:"username" validate:"required,min=5,max=255"`
}

func (c *MemberController) ProfileChangeUsername(ctx *gin.Context) {
	var req MemberChangeUsernameRequest
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
	member, err := c.currentUser.Member(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	if !c.repository.MemberVerifyPassword(member.ID.String(), req.Password) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong password"})
		return
	}
	updatedMember, err := c.repository.MemberUpdateByID(member.ID.String(), &models.Member{
		Username: req.Username,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update member details"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.MemberToResource(updatedMember))

}
