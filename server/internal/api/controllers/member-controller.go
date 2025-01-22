package controllers

import (
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
	repository  *models.ModelRepository
	footstep    *handlers.FootstepHandler
	currentUser *handlers.CurrentUser
}

func NewMemberController(
	repository *models.ModelRepository,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
) *MemberController {
	return &MemberController{
		repository:  repository,
		footstep:    footstep,
		currentUser: currentUser,
	}
}

func (c *MemberController) Index(ctx *gin.Context) {}

func (c *MemberController) Show(ctx *gin.Context) {}

type MemberStoreRequest struct {
	FirstName        string             `json:"firstName" validate:"required,max=255"`
	LastName         string             `json:"lastName" validate:"required,max=255"`
	MiddleName       string             `json:"middleName" validate:"max=255"`
	Username         string             `json:"username" validate:"max=255"`
	Email            string             `json:"email" validate:"required,email,max=255"`
	Password         string             `json:"password" validate:"required,min=8,max=255"`
	ConfirmPassword  string             `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=Password"`
	BirthDate        time.Time          `json:"birthDate" validate:"required"`
	ContactNumber    string             `json:"contactNumber" validate:"required,max=15"`
	PermanentAddress string             `json:"permanentAddress" validate:"required,max=500"`
	EmailTemplate    string             `json:"emailTemplate" validate:"required"`
	ContactTemplate  string             `json:"contactTemplate" validate:"required"`
	Media            *MediaStoreRequest `json:"media" validate:"omitempty"`
}

func (c *MemberController) Store(ctx *gin.Context) {

}

func (c *MemberController) Update(ctx *gin.Context) {}

func (c *MemberController) Destroy(ctx *gin.Context) {}

func (c *MemberController) ForgotPassword(ctx *gin.Context) {}
