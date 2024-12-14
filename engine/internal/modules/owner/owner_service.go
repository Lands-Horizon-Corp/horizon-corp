package owner

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
)

type OwnerService struct {
	controller *managers.Controller[models.Owner, models.OwnerRequest, models.OwnerResource]
	db         *providers.DatabaseService
	engine     *providers.EngineService
	models     *models.ModelResource
}

func NewOwnerService(
	db *providers.DatabaseService,
	engine *providers.EngineService,
	models *models.ModelResource,
) *OwnerService {
	controller := managers.NewController(
		models.OwnerDB,
		models.ValidateOwnerRequest,
		models.OwnerToResource,
		models.OwnerToResourceList,
	)

	return &OwnerService{
		controller: controller,
		db:         db,
		engine:     engine,
		models:     models,
	}
}

func (as *OwnerService) RegisterRoutes() {
	routes := as.engine.Client.Group("/api/v1/owner")
	{
		routes.POST("/", as.controller.Create)
		routes.GET("/", as.controller.GetAll)
		routes.GET("/:id", as.controller.GetByID)
		routes.PUT("/:id", as.controller.Update)
		routes.DELETE("/:id", as.controller.Delete)
	}
}
