package routes

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/admin"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
)

type APIRoutes struct {
	logger        *providers.LoggerService
	cache         *providers.CacheService
	engineService *providers.EngineService
	router        *gin.Engine
	adminService  *admin.AdminService
}

func NewAPIRoutes(
	logger *providers.LoggerService,
	cache *providers.CacheService,
	engineService *providers.EngineService,

	adminService *admin.AdminService,
) *APIRoutes {
	return &APIRoutes{
		logger:        logger,
		cache:         cache,
		engineService: engineService,
		router:        engineService.Client,
		adminService:  adminService,
	}
}

func (ar *APIRoutes) API() {
	ar.adminService.RegisterRoutes()
}
