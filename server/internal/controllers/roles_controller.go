package controllers

import (
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"
	"horizon/server/internal/requests"
	"horizon/server/internal/resources"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleController struct {
	*BaseController[models.Role, *requests.RoleRequest, resources.RoleResource]
}

func NewRoleController(db *gorm.DB) *RoleController {
	repo := repositories.NewRepository[models.Role](db)

	baseController := NewBaseController[models.Role, *requests.RoleRequest, resources.RoleResource](
		db,
		ToRoleModel,
		resources.ToResourceRole,
		resources.ToResourceListRoles,
		UpdateRoleModel,
	)

	baseController.Repo = repo

	return &RoleController{
		BaseController: baseController,
	}
}

func ToRoleModel(req *requests.RoleRequest) *models.Role {
	return &models.Role{
		Model: gorm.Model{
			ID: req.ID,
		},
		Name:               req.Name,
		Description:        req.Description,
		Color:              req.Color,
		ApiKey:             req.ApiKey,
		ReadRole:           req.ReadRole,
		WriteRole:          req.WriteRole,
		UpdateRole:         req.UpdateRole,
		DeleteRole:         req.DeleteRole,
		ReadErrorDetails:   req.ReadErrorDetails,
		WriteErrorDetails:  req.WriteErrorDetails,
		UpdateErrorDetails: req.UpdateErrorDetails,
		DeleteErrorDetails: req.DeleteErrorDetails,
		ReadGender:         req.ReadGender,
		WriteGender:        req.WriteGender,
		UpdateGender:       req.UpdateGender,
		DeleteGender:       req.DeleteGender,
	}
}

func UpdateRoleModel(model *models.Role, req *requests.RoleRequest) {
	model.Name = req.Name
	model.Description = req.Description
	model.Color = req.Color
	model.ApiKey = req.ApiKey
	model.ReadRole = req.ReadRole
	model.WriteRole = req.WriteRole
	model.UpdateRole = req.UpdateRole
	model.DeleteRole = req.DeleteRole
	model.ReadErrorDetails = req.ReadErrorDetails
	model.WriteErrorDetails = req.WriteErrorDetails
	model.UpdateErrorDetails = req.UpdateErrorDetails
	model.DeleteErrorDetails = req.DeleteErrorDetails
	model.ReadGender = req.ReadGender
	model.WriteGender = req.WriteGender
	model.UpdateGender = req.UpdateGender
	model.DeleteGender = req.DeleteGender
}

func RoleRoutes(router *gin.RouterGroup, controller *RoleController) {
	group := router.Group("/role")
	{
		group.POST("", controller.Create)
		group.GET("", controller.GetAll)
		group.GET("/:id", controller.GetByID)
		group.PUT("/:id", controller.Update)
		group.DELETE("/:id", controller.Delete)
	}
}
