package gender

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
)

type GenderService struct {
	controller *managers.Controller[models.Gender, models.GenderRequest, models.GenderResource]
	db         *providers.DatabaseService
	engine     *providers.EngineService
	models     *models.ModelResource
}

func NewGenderService(
	db *providers.DatabaseService,
	engine *providers.EngineService,
	models *models.ModelResource,
) *GenderService {
	controller := managers.NewController(
		models.GenderDB,
		models.ValidateGenderRequest,
		models.GenderToResource,
		models.GenderToResourceList,
	)

	return &GenderService{
		controller: controller,
		db:         db,
		engine:     engine,
		models:     models,
	}
}

func (as *GenderService) RegisterRoutes() {
	routes := as.engine.Client.Group("/api/v1/gender")
	{
		routes.POST("/", as.controller.Create)
		routes.GET("/", as.controller.GetAll)
		routes.GET("/:id", as.controller.GetByID)
		routes.PUT("/:id", as.controller.Update)
		routes.DELETE("/:id", as.controller.Delete)
	}
}
