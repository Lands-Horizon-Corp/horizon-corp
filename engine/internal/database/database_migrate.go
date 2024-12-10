package database

import (
	"context"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Migrate(
	lc fx.Lifecycle,
	db *gorm.DB,
	logger *zap.Logger,
	// admin models.AdminResourceProvider,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Running database migrations...")

			// Directly pass the list of models to AutoMigrate
			if err := db.AutoMigrate(
				&models.Admin{},
				&models.Branch{},
				&models.Company{},
				&models.Contact{},
				&models.Employee{},
				&models.Feedback{},
				&models.Footstep{},
				&models.Gender{},
				&models.Media{},
				&models.Member{},
				&models.Owner{},
				&models.Role{},
				&models.Timesheet{},
			); err != nil {
				logger.Error("Migration failed", zap.Error(err))
				return err
			}

			// admin.SeedDatabase()
			logger.Info("Database migrations completed successfully.")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Migration lifecycle stopped.")
			return nil
		},
	})
}
