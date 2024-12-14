package media

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
)

type MediaService struct {
	controller *managers.Controller[models.Media, models.MediaRequest, models.MediaResource]
	db         *providers.DatabaseService
	engine     *providers.EngineService
	models     *models.ModelResource
}

func NewMediaService(
	db *providers.DatabaseService,
	engine *providers.EngineService,
	models *models.ModelResource,
) *MediaService {
	controller := managers.NewController(
		models.MediaDB,
		models.ValidateMediaRequest,
		models.MediaToResource,
		models.MediaToResourceList,
	)

	return &MediaService{
		controller: controller,
		db:         db,
		engine:     engine,
		models:     models,
	}
}

func (as *MediaService) RegisterRoutes() {
	routes := as.engine.Client.Group("/api/v1/media")
	{
		routes.POST("/", as.controller.Create)
		routes.GET("/", as.controller.GetAll)
		routes.GET("/:id", as.controller.GetByID)
		routes.PUT("/:id", as.controller.Update)
		routes.DELETE("/:id", as.controller.Delete)
	}
}
