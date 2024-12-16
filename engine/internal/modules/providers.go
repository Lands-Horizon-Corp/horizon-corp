package modules

import (
	"context"
	"net/http"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/server/middleware"
	"github.com/Lands-Horizon-Corp/horizon-corp/server/routes"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModuleServiceProvider(
	lc fx.Lifecycle,
	cfg *config.AppConfig,
	logger *providers.LoggerService,
	cache *providers.CacheService,
	engineService *providers.EngineService,
	middle *middleware.Middleware,
	apiRoutes *routes.APIRoutes,
) {
	router := engineService.Client
	// Common Middlewares
	// router.Use(middle.BlockIPMiddleware())
	// router.Use(middle.DetectSuspiciousAccessMiddleware())
	// if cfg.AppEnv == "production" || cfg.AppEnv == "staging" {
	// 	router.Use(middle.EnforceHTTPS)
	// }
	router.Use(middle.Config())
	// router.Use(middle.Secure())
	// router.Use(middle.RateLimiterMiddleware(20, 20))

	apiRoutes.APITestRoute()
	apiRoutes.API()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("Starting Gin server on port: " + cfg.AppPort)
				if err := engineService.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Error("Failed to start server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down Gin server...")
			ctxShutdown, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			if err := engineService.Server.Shutdown(ctxShutdown); err != nil {
				logger.Error("Server forced to shutdown", zap.Error(err))
			}
			return nil
		},
	})

}
