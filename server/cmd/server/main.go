package main

import (
	"horizon-server/config"
	"horizon-server/internal/db"
	"horizon-server/internal/handlers"
	"horizon-server/internal/repositories"
	"horizon-server/internal/routes"
	"horizon-server/internal/services"
	"horizon-server/pkg/logger"
	"horizon-server/pkg/storage"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	_ "horizon-server/cmd/server/docs"
)

func main() {
	fx.New(
		fx.Provide(
			config.LoadConfig,
			logger.NewLogger,
			storage.NewMinioClient,
			db.NewDatabase,
			repositories.NewUserRepository,
			services.NewUserService,
			handlers.NewUserHandler,
			routes.SetupRouter,
		),
		fx.Invoke(startServer),
		fx.WithLogger(func(log *log.Logger) fxevent.Logger {
			return &fxevent.ConsoleLogger{W: log.Writer()}
		}),
	).Run()
}

func startServer(router *gin.Engine, cfg *config.Config) {
	router.Run(":" + cfg.App.Port)
}
