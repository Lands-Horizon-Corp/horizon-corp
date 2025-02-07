package role

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
)

type RoleService struct {
	controller *managers.Controller[models.Role, models.RoleRequest, models.RoleResource]
	db         *providers.DatabaseService
	engine     *providers.EngineService
	models     *models.ModelResource
}

func NewRoleService(
	db *providers.DatabaseService,
	engine *providers.EngineService,
	models *models.ModelResource,
) *RoleService {
	controller := managers.NewController(
		models.RoleDB,
		models.ValidateRoleRequest,
		models.RoleToResource,
		models.RoleToResourceList,
	)

	return &RoleService{
		controller: controller,
		db:         db,
		engine:     engine,
		models:     models,
	}
}

func (as *RoleService) RegisterRoutes() {
	routes := as.engine.Client.Group("/api/v1/role")
	{
		routes.POST("/", as.controller.Create)
		routes.GET("/", as.controller.GetAll)
		routes.GET("/:id", as.controller.GetByID)
		routes.PUT("/:id", as.controller.Update)
		routes.DELETE("/:id", as.controller.Delete)
	}
}
