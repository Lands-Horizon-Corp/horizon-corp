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

			// Error Details
			repositories.NewErrorDetailsRepository,
			controllers.NewErrorDetailsController,

			// Gender
			repositories.NewGenderRepository,
			controllers.NewGenderController,

			// Router
			routes.ProvideAPI,
		),
		fx.Invoke(internal.StartServer),
	)

	// Start the application
	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}

	// Create a channel to listen for OS signals (like SIGINT or SIGTERM)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-sig

	// Gracefully stop the application
	if err := app.Stop(ctx); err != nil {
		log.Fatal(err)
	}
}
