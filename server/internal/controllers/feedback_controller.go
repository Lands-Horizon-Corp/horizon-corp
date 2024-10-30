package controllers

import (
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"
	"horizon/server/internal/requests"
	"horizon/server/internal/resources"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FeedbackController struct {
	*BaseController[models.Feedback, *requests.FeedbackRequest, resources.FeedbackResource]
}

func NewFeedbackController(db *gorm.DB) *FeedbackController {
	repo := repositories.NewRepository[models.Feedback](db)

	baseController := NewBaseController[models.Feedback, *requests.FeedbackRequest, resources.FeedbackResource](
		db,
		ToFeedbackModel,
		resources.ToResourceFeedback,
		resources.ToResourceListFeedback,
		UpdateFeedbackModel,
	)

	baseController.Repo = repo

	return &FeedbackController{
		BaseController: baseController,
	}
}

func ToFeedbackModel(req *requests.FeedbackRequest) *models.Feedback {
	return &models.Feedback{
		Model: gorm.Model{
			ID: req.ID,
		},
		Email:        req.Email,
		Description:  req.Description,
		FeedbackType: req.FeedbackType,
	}
}

func UpdateFeedbackModel(model *models.Feedback, req *requests.FeedbackRequest) {
	model.Email = req.Email
	model.Description = req.Description
	model.FeedbackType = req.FeedbackType
}

func FeedbackRoutes(router *gin.RouterGroup, controller *FeedbackController) {
	group := router.Group("/feedback")
	{
		group.POST("", controller.Create)
		group.GET("", controller.GetAll)
		group.GET("/:id", controller.GetByID)
		group.PUT("/:id", controller.Update)
		group.DELETE("/:id", controller.Delete)
	}
}
