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
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/fx"
)

func main() {
	ctx := context.Background()
	app := fx.New(
		fx.Provide(
			config.LoadConfig,
			logger.NewLogger,
			database.NewDB,

			// Roles
			repositories.NewRolesRepository,
			controllers.NewRolesController,

			// Error Details
			repositories.NewErrorDetailRepository,
			controllers.NewErrorDetailController,

			// Gender
			repositories.NewGenderRepository,
			controllers.NewGenderController,

			// Contact
			repositories.NewContactsRepository,
			controllers.NewContactsController,

			// Feedback
			repositories.NewFeedbackRepository,
			controllers.NewFeedbackController,

			// Media
			repositories.NewMediaRepository,
			controllers.NewMediaController,

			// Router
			routes.ProvideAPI,
		),
		fx.Invoke(internal.StartServer),
	)

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig

	if err := app.Stop(ctx); err != nil {
		log.Fatal(err)
	}
}
