package main

import (
	"context"
	"horizon/server/config"
	"horizon/server/database"
	"horizon/server/internal"
	"horizon/server/internal/auth"
	"horizon/server/internal/controllers"
	"horizon/server/internal/middleware"
	"horizon/server/internal/repositories"
	"horizon/server/internal/routes"
	"horizon/server/logger"
	"horizon/server/services"
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
			// Dependencies
			config.LoadConfig,
			logger.NewLogger,
			database.NewDatabaseService,
			database.NewCacheService,

			// Services
			services.NewEmailService,
			services.NewSMSService,
			services.NewOTPService,

			// Authentication
			auth.NewAdminAuthService,
			auth.NewEmployeeAuthService,
			auth.NewMemberAuthService,
			auth.NewOwnerAuthService,
			auth.NewTokenService,

			// Middleware
			middleware.NewAuthMiddleware,

			// Authentication
			repositories.NewAdminRepository,
			repositories.NewEmployeeRepository,
			repositories.NewOwnerRepository,
			repositories.NewMemberRepository,
			controllers.NewAuthController,

			// User
			auth.NewUserAuthService,
			repositories.NewUserRepository,
			controllers.NewUserController,

			// Roles
			repositories.NewRoleRepository,
			controllers.NewRoleController,

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

			// Timesheets
			repositories.NewTimesheetRepository,
			controllers.NewTimesheetController,

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
