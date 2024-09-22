package main

import (
	"horizon-core/config"
	"horizon-core/database"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {

	fx.New(
		fx.Provide(
			// Dependencies
			config.PorvideConfig,
			config.ProvideLogger,
			database.ProvideDatabase,

			// Repositories
		),
		fx.Invoke(func(router *gin.Engine, cfg *config.Config) {
			router.Run(":" + cfg.App.Port)
		}),
		fx.Invoke(func(logger *zap.Logger) {
			logger.Info("Application starting up")
		}),
	).Run()
}
