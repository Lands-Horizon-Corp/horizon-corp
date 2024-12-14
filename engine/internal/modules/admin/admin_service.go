package admin

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
)

type AdminService struct {
	controller *managers.Controller[models.Admin, models.AdminRequest, models.AdminResource]
	db         *providers.DatabaseService
	engine     *providers.EngineService
	models     *models.ModelResource
}

func NewAdminService(
	db *providers.DatabaseService,
	engine *providers.EngineService,
	models *models.ModelResource,
) *AdminService {
	controller := managers.NewController(
		models.AdminDB,
		models.ValidateAdminRequest,
		models.AdminToResource,
		models.AdminToResourceList,
	)

	return &AdminService{
		controller: controller,
		db:         db,
		engine:     engine,
		models:     models,
	}
}

func (as *AdminService) RegisterRoutes() {
	routes := as.engine.Client.Group("/api/v1//admin")
	{
		routes.POST("/", as.controller.Create)
		routes.GET("/", as.controller.GetAll)
		routes.GET("/:id", as.controller.GetByID)
		routes.PUT("/:id", as.controller.Update)
		routes.DELETE("/:id", as.controller.Delete)
	}
}
