package feedback

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
)

type FeedbackService struct {
	controller *managers.Controller[models.Feedback, models.FeedbackRequest, models.FeedbackResource]
	db         *providers.DatabaseService
	engine     *providers.EngineService
	models     *models.ModelResource
}

func NewFeedbackService(
	db *providers.DatabaseService,
	engine *providers.EngineService,
	models *models.ModelResource,
) *FeedbackService {
	controller := managers.NewController(
		models.FeedbackDB,
		models.ValidateFeedbackRequest,
		models.FeedbackToResource,
		models.FeedbackToResourceList,
	)

	return &FeedbackService{
		controller: controller,
		db:         db,
		engine:     engine,
		models:     models,
	}
}

func (as *FeedbackService) RegisterRoutes() {
	routes := as.engine.Client.Group("/api/v1/feedback")
	{
		routes.POST("/", as.controller.Create)
		routes.GET("/", as.controller.GetAll)
		routes.GET("/:id", as.controller.GetByID)
		routes.PUT("/:id", as.controller.Update)
		routes.DELETE("/:id", as.controller.Delete)
	}
}
