package main

import (
	"context"
	"horizon/server/config"
	"horizon/server/database"
	"horizon/server/internal"
	"horizon/server/internal/controllers"
	"horizon/server/internal/repositories"
	"horizon/server/internal/routes"
	"horizon/server/logger"
	"log"

	"go.uber.org/fx"
)

func main() {
	ctx := context.Background()
	app := fx.New(
		fx.Provide(
			config.LoadConfig,
			logger.NewLogger,
			database.NewDB,

			// Gender
			repositories.NewGenderRepository,
			controllers.NewGenderController,

			// Router
			routes.ProvideAPI,
		),
		fx.Invoke(internal.StartServer),
	)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := app.Stop(ctx); err != nil {
			log.Fatal(err)
		}
	}()
}
