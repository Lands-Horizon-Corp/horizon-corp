package routes

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
)

type APIRoutes struct {
	logger        *providers.LoggerService
	cache         *providers.CacheService
	engineService *providers.EngineService
	router        *gin.Engine
}

func NewAPIRoutes(
	logger *providers.LoggerService,
	cache *providers.CacheService,
	engineService *providers.EngineService,
) *APIRoutes {
	return &APIRoutes{
		logger:        logger,
		cache:         cache,
		engineService: engineService,
		router:        engineService.Client,
	}
}

func (ar *APIRoutes) Routes() {
	// Here routes
}
