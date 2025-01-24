package controllers

import (
	"net/http"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type EmployeeController struct {
	repository  *models.ModelRepository
	transformer *models.ModelTransformer
	footstep    *handlers.FootstepHandler
	currentUser *handlers.CurrentUser
}

func NewEmployeeController(
	repository *models.ModelRepository,
	transformer *models.ModelTransformer,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
) *EmployeeController {
	return &EmployeeController{
		repository:  repository,
		transformer: transformer,
		footstep:    footstep,
		currentUser: currentUser,
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
}

func (c *EmployeeController) Store(ctx *gin.Context) {
	var req EmployeeStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if validator.New().Struct(req) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": req})
		return
	}
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
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create employee", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, c.transformer.EmployeeToResource(employee))
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

func (c *EmployeeController) ForgotPassword(ctx *gin.Context) {

}
