package controllers

import (
	"horizon/server/helpers"
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"
	"horizon/server/internal/requests"
	"horizon/server/internal/resources"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ErrorDetailController struct {
	*BaseController[models.ErrorDetail, *requests.ErrorDetailRequest, resources.ErrorDetailResource]
}

func NewErrorDetailController(db *gorm.DB) *ErrorDetailController {
	repo := repositories.NewRepository[models.ErrorDetail](db)

	baseController := NewBaseController[models.ErrorDetail, *requests.ErrorDetailRequest, resources.ErrorDetailResource](
		db,
		ToErrorDetailModel,
		resources.ToResourceErrorDetail,
		resources.ToResourceListErrorDetail,
		UpdateErrorDetailModel,
	)

	// Set the repository
	baseController.Repo = repo

	return &ErrorDetailController{
		BaseController: baseController,
	}
}

func ToErrorDetailModel(req *requests.ErrorDetailRequest) *models.ErrorDetail {
	var stackJSON, responseJSON string

	if req.Stack != "" {
		stackJSON = helpers.JSONStringify(req.Stack)
	}
	if req.Response != "" {
		responseJSON = helpers.JSONStringify(req.Response)
	}

	return &models.ErrorDetail{
		Model:    gorm.Model{ID: req.ID},
		Message:  req.Message,
		Name:     req.Name,
		Stack:    stackJSON,
		Response: responseJSON,
		Status:   req.Status,
	}
}

// UpdateErrorDetailModel updates an ErrorDetail model with data from an ErrorDetailRequest
func UpdateErrorDetailModel(model *models.ErrorDetail, req *requests.ErrorDetailRequest) {
	var stackJSON, responseJSON string

	if req.Stack != "" {
		stackJSON = helpers.JSONStringify(req.Stack)
	}
	if req.Response != "" {
		responseJSON = helpers.JSONStringify(req.Response)
	}

	model.Message = req.Message
	model.Name = req.Name
	model.Stack = stackJSON
	model.Response = responseJSON
	model.Status = req.Status
}

func ErrorDetailRoutes(router *gin.RouterGroup, controller *ErrorDetailController) {
	group := router.Group("/error-details")
	{
		group.POST("", controller.Create)
		group.GET("", controller.GetAll)
		group.GET("/:id", controller.GetByID)
		group.PUT("/:id", controller.Update)
		group.DELETE("/:id", controller.Delete)
	}
}
