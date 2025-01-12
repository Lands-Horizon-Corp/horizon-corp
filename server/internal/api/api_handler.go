package api

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/controllers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/middleware"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"go.uber.org/fx"
)

func NewAPIHandlerInvoke(
	lc fx.Lifecycle,
	cfg *config.AppConfig,
	logger *providers.LoggerService,
	cache *providers.CacheService,
	engineService *providers.EngineService,
	middle *middleware.Middleware,

	// Controllers
	adminController *controllers.AdminController,
	controller *controllers.Controller,
) {
	router := engineService.Client
	registerMiddleware(router, middle, cfg)

	router.GET("/", controller.Index)
	router.GET("/favicon.ico", controller.Favico)
	router.GET("/ping", controller.Ping)
	apiGroup := router.Group("/api/v1")
	{
		adminGroup := apiGroup.Group("/admin")
		{
			adminGroup.GET("", adminController.Index)
		}
	}

	runServer(lc, cfg, engineService, logger)
}
