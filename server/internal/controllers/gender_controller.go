package controllers

import (
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"
	"horizon/server/internal/requests"
	"horizon/server/internal/resources"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GenderController struct {
	*BaseController[models.Gender, *requests.GenderRequest, resources.GenderResource]
}

func NewGenderController(db *gorm.DB) *GenderController {
	repo := repositories.NewRepository[models.Gender](db)
	baseController := NewBaseController[models.Gender, *requests.GenderRequest, resources.GenderResource](
		db,
		ToGenderModel,
		resources.ToResourceGender,
		resources.ToResourceListGender,
		UpdateGenderModel,
	)
	baseController.Repo = repo
	return &GenderController{
		BaseController: baseController,
	}
}
func ToGenderModel(req *requests.GenderRequest) *models.Gender {
	return &models.Gender{
		Name:        req.Name,
		Description: req.Description,
	}
}
func UpdateGenderModel(model *models.Gender, req *requests.GenderRequest) {
	model.Name = req.Name
	model.Description = req.Description
}

func GenderRoutes(router *gin.RouterGroup, controller *GenderController) {
	group := router.Group("/gender")
	{
		group.POST("", controller.Create)
		group.GET("", controller.GetAll)
		group.GET("/:id", controller.GetByID)
		group.PUT("/:id", controller.Update)
		group.DELETE("/:id", controller.Delete)
	}
}
