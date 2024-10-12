package internal

import (
	"context"
	"horizon/server/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func StartServer(lc fx.Lifecycle, router *gin.Engine, logger *zap.Logger, cfg *config.AppConfig) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := router.Run(":" + cfg.AppPort); err != nil {
					logger.Error("Failed to start server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping the server...")
			return nil
		},
	})
}
