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

type OwnerController struct {
	repository  *models.ModelRepository
	transformer *models.ModelTransformer
	footstep    *handlers.FootstepHandler
	currentUser *handlers.CurrentUser
	otpService  *providers.OTPService
}

func NewOwnerController(
	repository *models.ModelRepository,
	transformer *models.ModelTransformer,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
	otpService *providers.OTPService,
) *OwnerController {
	return &OwnerController{
		repository:  repository,
		transformer: transformer,
		footstep:    footstep,
		currentUser: currentUser,
		otpService:  otpService,
	}
}

// GET :/api/v1/owner/
// For owner only find all owner
func (c *OwnerController) Index(ctx *gin.Context) {

}

// GET :/api/v1/owner/id
// for owner only and self for owner
func (c *OwnerController) Show(ctx *gin.Context) {

}

// POST: /api/v1/owner/
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
	GenderID *uuid.UUID `json:"genderId" validate:"omitempty,uuid"`
	RoleID   *uuid.UUID `json:"roleId" validate:"omitempty,uuid"`

	EmailTemplate   string `json:"emailTemplate" validate:"required"`
	ContactTemplate string `json:"contactTemplate" validate:"required"`
}

func (c *OwnerController) Store(ctx *gin.Context) {
	c.Create(ctx)
}

// PUT: /api/v1/owner/:id
func (c *OwnerController) Update(ctx *gin.Context) {

}

// DELETE: /
// For owner only delete owner but if no branch and company and no employee or members
func (c *OwnerController) Destroy(ctx *gin.Context) {

}

// POST: /forgot-password
// owner: only  for email, phone number, and actual link
// Public: ownly  for email, and phone number
func (c *OwnerController) ForgotPassword(ctx *gin.Context) {

}

func (c *OwnerController) Create(ctx *gin.Context) *models.Owner {
	var req OwnerStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return nil
	}
	if validator.New().Struct(req) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": req})
		return nil
	}

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
		// Automatic not allowed
		Status: providers.NotAllowedStatus,

		MediaID:  req.MediaID,
		RoleID:   req.RoleID,
		GenderID: req.GenderID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create owner", "details": err.Error()})
		return nil
	}

	if err := c.otpService.SendEmailOTP(providers.OTPMessage{
		AccountType: "owner", ID: owner.ID.String(),
		MediumType: "owner",
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
		AccountType: "owner", ID: owner.ID.String(),
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
	ctx.JSON(http.StatusCreated, c.transformer.OwnerToResource(owner))
	return owner
}
