package employee

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
)

type EmployeeService struct {
	controller *managers.Controller[models.Employee, models.EmployeeRequest, models.EmployeeResource]
	db         *providers.DatabaseService
	engine     *providers.EngineService
	models     *models.ModelResource
}

func NewEmployeeService(
	db *providers.DatabaseService,
	engine *providers.EngineService,
	models *models.ModelResource,
) *EmployeeService {
	controller := managers.NewController(
		models.EmployeeDB,
		models.ValidateEmployeeRequest,
		models.EmployeeToResource,
		models.EmployeeToResourceList,
	)

	return &EmployeeService{
		controller: controller,
		db:         db,
		engine:     engine,
		models:     models,
	}
}

func (as *EmployeeService) RegisterRoutes() {
	routes := as.engine.Client.Group("/api/v1/employee")
	{
		routes.POST("/", as.controller.Create)
		routes.GET("/", as.controller.GetAll)
		routes.GET("/:id", as.controller.GetByID)
		routes.PUT("/:id", as.controller.Update)
		routes.DELETE("/:id", as.controller.Delete)
	}
}
