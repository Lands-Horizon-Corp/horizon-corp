package member

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
)

type MemberService struct {
	controller *managers.Controller[models.Member, models.MemberRequest, models.MemberResource]
	db         *providers.DatabaseService
	engine     *providers.EngineService
	models     *models.ModelResource
}

func NewMemberService(
	db *providers.DatabaseService,
	engine *providers.EngineService,
	models *models.ModelResource,
) *MemberService {
	controller := managers.NewController(
		models.MemberDB,
		models.ValidateMemberRequest,
		models.MemberToResource,
		models.MemberToResourceList,
	)

	return &MemberService{
		controller: controller,
		db:         db,
		engine:     engine,
		models:     models,
	}
}

func (as *MemberService) RegisterRoutes() {
	routes := as.engine.Client.Group("/api/v1/member")
	{
		routes.POST("/", as.controller.Create)
		routes.GET("/", as.controller.GetAll)
		routes.GET("/:id", as.controller.GetByID)
		routes.PUT("/:id", as.controller.Update)
		routes.DELETE("/:id", as.controller.Delete)
	}
}
