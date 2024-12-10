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

	adminRoutes := engine.Client.Group("/admins")
	{
		adminRoutes.POST("/", controller.Create)
		adminRoutes.GET("/", controller.GetAll)
		adminRoutes.GET("/:id", controller.GetByID)
		adminRoutes.PUT("/:id", controller.Update)
		adminRoutes.DELETE("/:id", controller.Delete)
	}

	return &AdminService{
		controller: controller,
		db:         db,
		engine:     engine,
		models:     models,
	}
}

func (as *AdminService) RegisterRoutes() {
	adminRoutes := as.engine.Client.Group("/admins")
	{
		adminRoutes.POST("/", as.controller.Create)
		adminRoutes.GET("/", as.controller.GetAll)
		adminRoutes.GET("/:id", as.controller.GetByID)
		adminRoutes.PUT("/:id", as.controller.Update)
		adminRoutes.DELETE("/:id", as.controller.Delete)
	}
}
