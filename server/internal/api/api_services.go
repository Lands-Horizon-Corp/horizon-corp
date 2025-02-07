package api

import (
	"context"
	"net/http"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/middleware"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func registerMiddleware(router *gin.Engine, middle *middleware.Middleware, cfg *config.AppConfig) {
	router.Use(middle.BlockIPMiddleware())
	router.Use(middle.DetectSuspiciousAccessMiddleware())
	if cfg.AppEnv == "production" || cfg.AppEnv == "staging" {
		router.Use(middle.EnforceHTTPS)
	}
	router.Use(middle.Config())
	router.Use(middle.Secure())
	router.Use(middle.RateLimiterMiddleware(20, 20))
}

func startServer(port string, engineService *providers.EngineService, logger *providers.LoggerService) {
	logger.Info("Starting Gin server on port: " + port)
	if err := engineService.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("Failed to start server", zap.Error(err))
	}
}

func stopServer(ctx context.Context, engineService *providers.EngineService, logger *providers.LoggerService) error {
	logger.Info("Shutting down Gin server...")
	ctxShutdown, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := engineService.Server.Shutdown(ctxShutdown); err != nil {
		logger.Error("Server forced to shutdown", zap.Error(err))
		return err
	}
	return nil
}

func runServer(lc fx.Lifecycle, cfg *config.AppConfig, engineService *providers.EngineService, logger *providers.LoggerService) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go startServer(cfg.AppPort, engineService, logger)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return stopServer(ctx, engineService, logger)
		},
	})
}
