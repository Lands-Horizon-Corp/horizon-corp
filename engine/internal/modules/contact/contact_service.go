package contact

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
)

type ContactService struct {
	controller *managers.Controller[models.Contact, models.ContactRequest, models.ContactResource]
	db         *providers.DatabaseService
	engine     *providers.EngineService
	models     *models.ModelResource
}

func NewContactService(
	db *providers.DatabaseService,
	engine *providers.EngineService,
	models *models.ModelResource,
) *ContactService {
	controller := managers.NewController(
		models.ContactDB,
		models.ValidateContactRequest,
		models.ContactToResource,
		models.ContactToResourceList,
	)

	return &ContactService{
		controller: controller,
		db:         db,
		engine:     engine,
		models:     models,
	}
}

func (as *ContactService) RegisterRoutes() {
	routes := as.engine.Client.Group("/api/v1/contact")
	{
		routes.POST("/", as.controller.Create)
		routes.GET("/", as.controller.GetAll)
		routes.GET("/:id", as.controller.GetByID)
		routes.PUT("/:id", as.controller.Update)
		routes.DELETE("/:id", as.controller.Delete)
	}
}
