package footstep

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
)

type FootstepService struct {
	controller *managers.Controller[models.Footstep, models.FootstepRequest, models.FootstepResource]
	db         *providers.DatabaseService
	engine     *providers.EngineService
	models     *models.ModelResource
}

func NewFootstepService(
	db *providers.DatabaseService,
	engine *providers.EngineService,
	models *models.ModelResource,
) *FootstepService {
	controller := managers.NewController(
		models.FootstepDB,
		models.ValidateFootstepRequest,
		models.FootstepToResource,
		models.FootstepToResourceList,
	)

	return &FootstepService{
		controller: controller,
		db:         db,
		engine:     engine,
		models:     models,
	}
}

func (as *FootstepService) RegisterRoutes() {
	routes := as.engine.Client.Group("/api/v1/footstep")
	{
		routes.POST("/", as.controller.Create)
		routes.GET("/", as.controller.GetAll)
		routes.GET("/:id", as.controller.GetByID)
		routes.PUT("/:id", as.controller.Update)
		routes.DELETE("/:id", as.controller.Delete)
	}
}
