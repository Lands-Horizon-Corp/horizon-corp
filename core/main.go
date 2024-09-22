package main

import (
	"horizon-core/config"
	"horizon-core/database"

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
		fx.Invoke(func(logger *zap.Logger) {
			logger.Info("Application starting up")
		}),
	).Run()
}
