package branch

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
)

type BranchService struct {
	controller *managers.Controller[models.Branch, models.BranchRequest, models.BranchResource]
	db         *providers.DatabaseService
	engine     *providers.EngineService
	models     *models.ModelResource
}

func NewBranchService(
	db *providers.DatabaseService,
	engine *providers.EngineService,
	models *models.ModelResource,
) *BranchService {
	controller := managers.NewController(
		models.BranchDB,
		models.ValidateBranchRequest,
		models.BranchToResource,
		models.BranchToResourceList,
	)

	return &BranchService{
		controller: controller,
		db:         db,
		engine:     engine,
		models:     models,
	}
}

func (as *BranchService) RegisterRoutes() {
	routes := as.engine.Client.Group("/branch")
	{
		routes.POST("/", as.controller.Create)
		routes.GET("/", as.controller.GetAll)
		routes.GET("/:id", as.controller.GetByID)
		routes.PUT("/:id", as.controller.Update)
		routes.DELETE("/:id", as.controller.Delete)
	}
}
