package controllers

import (
	"net/http"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	repository    *models.ModelRepository
	transformer   *models.ModelTransformer
	footstep      *handlers.FootstepHandler
	currentUser   *handlers.CurrentUser
	tokenProvider *providers.TokenService
	cfg           *config.AppConfig
	helpers       *helpers.HelpersFunction

	adminController    *AdminController
	ownerController    *OwnerController
	employeeController *EmployeeController
	memberController   *MemberController
}

func NewProfileController(
	repository *models.ModelRepository,
	transformer *models.ModelTransformer,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
	tokenProvider *providers.TokenService,
	cfg *config.AppConfig,
	helpers *helpers.HelpersFunction,

	adminController *AdminController,
	ownerController *OwnerController,
	employeeController *EmployeeController,
	memberController *MemberController,
) *ProfileController {
	return &ProfileController{
		repository:    repository,
		transformer:   transformer,
		footstep:      footstep,
		currentUser:   currentUser,
		tokenProvider: tokenProvider,
		cfg:           cfg,
		helpers:       helpers,

		adminController:    adminController,
		ownerController:    ownerController,
		employeeController: employeeController,
		memberController:   memberController,
	}
}

func (as *ProfileController) ProfilePicture(ctx *gin.Context) {}

type AccountSettingRequest struct {
	BirthDate        time.Time `json:"birthDate" validate:"required"`
	FirstName        string    `json:"firstName" validate:"required,max=255"`
	MiddleName       string    `json:"middleName" validate:"required,max=255"`
	LastName         string    `json:"lastName" validate:"required,max=255"`
	Description      string    `json:"description" validate:"max=2048"`
	PermanentAddress string    `json:"permanentAddress" validate:"required,max=500"`
}

func (as *ProfileController) ProfileAccountSetting(ctx *gin.Context) {
	user, _, err := as.currentUser.Claims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	switch user.AccountType {
	case "Member":
		as.memberController.ProfileAccountSetting(ctx)
	case "Admin":
		as.adminController.ProfileAccountSetting(ctx)
	case "Owner":
		as.ownerController.ProfileAccountSetting(ctx)
	case "Employee":
		as.employeeController.ProfileAccountSetting(ctx)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
}

func (as *ProfileController) ProfileChangeEmail(ctx *gin.Context) {
	user, _, err := as.currentUser.Claims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	switch user.AccountType {
	case "Member":
		as.memberController.ProfileChangeEmail(ctx)
	case "Admin":
		as.adminController.ProfileChangeEmail(ctx)
	case "Owner":
		as.ownerController.ProfileChangeEmail(ctx)
	case "Employee":
		as.employeeController.ProfileChangeEmail(ctx)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
}

func (as *ProfileController) ProfileChangeContactNumber(ctx *gin.Context) {
	user, _, err := as.currentUser.Claims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	switch user.AccountType {
	case "Member":
		as.memberController.ProfileChangeContactNumber(ctx)
	case "Admin":
		as.adminController.ProfileChangeContactNumber(ctx)
	case "Owner":
		as.ownerController.ProfileChangeContactNumber(ctx)
	case "Employee":
		as.employeeController.ProfileChangeContactNumber(ctx)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
}

func (as *ProfileController) ProfileChangePassword(ctx *gin.Context) {
	user, _, err := as.currentUser.Claims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	switch user.AccountType {
	case "Member":
		as.memberController.ChangePassword(ctx)
	case "Admin":
		as.adminController.ChangePassword(ctx)
	case "Owner":
		as.ownerController.ChangePassword(ctx)
	case "Employee":
		as.employeeController.ChangePassword(ctx)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
}

func (as *ProfileController) ProfileChangeUsername(ctx *gin.Context) {

	user, _, err := as.currentUser.Claims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	switch user.AccountType {
	case "Member":
		as.memberController.ProfileChangeUsername(ctx)
	case "Admin":
		as.adminController.ProfileChangeUsername(ctx)
	case "Owner":
		as.ownerController.ProfileChangeUsername(ctx)
	case "Employee":
		as.employeeController.ProfileChangeUsername(ctx)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}

}
