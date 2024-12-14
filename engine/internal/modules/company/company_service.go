package company

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
)

type CompanyService struct {
	controller *managers.Controller[models.Company, models.CompanyRequest, models.CompanyResource]
	db         *providers.DatabaseService
	engine     *providers.EngineService
	models     *models.ModelResource
}

func NewCompanyService(
	db *providers.DatabaseService,
	engine *providers.EngineService,
	models *models.ModelResource,
) *CompanyService {
	controller := managers.NewController(
		models.CompanyDB,
		models.ValidateCompanyRequest,
		models.CompanyToResource,
		models.CompanyToResourceList,
	)

	return &CompanyService{
		controller: controller,
		db:         db,
		engine:     engine,
		models:     models,
	}
}

func (as *CompanyService) RegisterRoutes() {
	routes := as.engine.Client.Group("/api/v1/company")
	{
		routes.POST("/", as.controller.Create)
		routes.GET("/", as.controller.GetAll)
		routes.GET("/:id", as.controller.GetByID)
		routes.PUT("/:id", as.controller.Update)
		routes.DELETE("/:id", as.controller.Delete)
	}
}
