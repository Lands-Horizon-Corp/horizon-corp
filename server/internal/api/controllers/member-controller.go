package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type MemberController struct {
	repository  *models.ModelRepository
	transformer *models.ModelTransformer
	footstep    *handlers.FootstepHandler
	currentUser *handlers.CurrentUser
	otpService  *providers.OTPService
}

func NewMemberController(
	repository *models.ModelRepository,
	transformer *models.ModelTransformer,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
	otpService *providers.OTPService,
) *MemberController {
	return &MemberController{
		repository:  repository,
		transformer: transformer,
		footstep:    footstep,
		currentUser: currentUser,
		otpService:  otpService,
	}
}

func (c *MemberController) Index(ctx *gin.Context) {}

func (c *MemberController) Show(ctx *gin.Context) {}

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

	Longitude *float64   `json:"longitude" validate:"omitempty,numeric"`
	Latitude  *float64   `json:"latitude" validate:"omitempty,numeric"`
	MediaID   *uuid.UUID `json:"mediaId" validate:"omitempty,uuid"`
	RoleID    *uuid.UUID `json:"roleId" validate:"omitempty,uuid"`
	GenderID  *uuid.UUID `json:"genderId" validate:"omitempty,uuid"`

	EmailTemplate   string `json:"emailTemplate" validate:"required"`
	ContactTemplate string `json:"contactTemplate" validate:"required"`
}

func (c *MemberController) Store(ctx *gin.Context) {
	c.Create(ctx)
}

func (c *MemberController) Update(ctx *gin.Context) {}

func (c *MemberController) Destroy(ctx *gin.Context) {}

func (c *MemberController) ForgotPassword(ctx *gin.Context) {}

func (c *MemberController) Create(ctx *gin.Context) *models.Member {
	var req MemberStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return nil
	}
	if validator.New().Struct(req) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": req})
		return nil
	}

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
		Status:             providers.VerifiedStatus,

		Longitude: req.Longitude,
		Latitude:  req.Latitude,
		MediaID:   req.MediaID,
		RoleID:    req.RoleID,
		GenderID:  req.GenderID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create member", "details": err.Error()})
		return nil
	}
	if err := c.otpService.SendEmailOTP(providers.OTPMessage{
		AccountType: "member", ID: member.ID.String(),
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
		AccountType: "member", ID: member.ID.String(),
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
	ctx.JSON(http.StatusCreated, c.transformer.MemberToResource(member))
	return member
}
