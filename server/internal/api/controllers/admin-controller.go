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

type AdminController struct {
	repository  *models.ModelRepository
	transformer *models.ModelTransformer
	footstep    *handlers.FootstepHandler
	currentUser *handlers.CurrentUser
	otpService  *providers.OTPService
}

func NewAdminController(
	repository *models.ModelRepository,
	transformer *models.ModelTransformer,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
	otpService *providers.OTPService,
) *AdminController {
	return &AdminController{
		repository:  repository,
		transformer: transformer,
		footstep:    footstep,
		currentUser: currentUser,
		otpService:  otpService,
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
	// if Logged in or not
	//	 	Only Admin and verified
	//		do not set cookies
	// else
	// 		Log me in
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

func (c *AdminController) ForgotPassword(ctx *gin.Context) {

}

func (c *AdminController) Create(ctx *gin.Context) *models.Admin {
	var req AdminStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return nil
	}
	if validator.New().Struct(req) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": req})
		return nil
	}

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
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create admin", "details": err.Error()})
		return nil
	}
	if err := c.otpService.SendEmailOTP(providers.OTPMessage{
		AccountType: "admin", ID: admin.ID.String(),
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
		AccountType: "admin", ID: admin.ID.String(),
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
	ctx.JSON(http.StatusCreated, c.transformer.AdminToResource(admin))
	return admin
}
