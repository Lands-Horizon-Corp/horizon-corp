package database

import (
	"context"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"go.uber.org/fx"
)

type DatabaseMigration struct {
	lc     fx.Lifecycle
	db     *providers.DatabaseService
	logger *providers.LoggerService
}

func NewDatabaseMigration(
	lc fx.Lifecycle,
	db *providers.DatabaseService,
	logger *providers.LoggerService,
) *DatabaseMigration {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Running database migrations...")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Migration lifecycle stopped.")
			return nil
		},
	})
	return &DatabaseMigration{
		lc:     lc,
		db:     db,
		logger: logger,
	}
}

// type DatabaseMigration struct {
// 	lc       fx.Lifecycle
// 	db       *providers.DatabaseService
// 	logger   *providers.LoggerService
// 	resource *models.ModelResource
// }

// func NewDatabaseMigration(
// 	lc fx.Lifecycle,
// 	db *providers.DatabaseService,
// 	logger *providers.LoggerService,
// 	resource *models.ModelResource,
// ) *DatabaseMigration {

// 	migration := &DatabaseMigration{
// 		lc:       lc,
// 		db:       db,
// 		logger:   logger,
// 		resource: resource,
// 	}

// 	lc.Append(fx.Hook{
// 		OnStart: func(ctx context.Context) error {
// 			logger.Info("Running database migrations...")
// 			if err := migrateModels(db.Client, resource.Models, logger); err != nil {
// 				return err
// 			}
// 			if err := seedModels(db.Client, resource.Models, logger); err != nil {
// 				return err
// 			}
// 			logger.Info("Database migrations and seeding completed successfully.")
// 			return nil
// 		},
// 		OnStop: func(ctx context.Context) error {
// 			logger.Info("Migration lifecycle stopped.")
// 			return nil
// 		},
// 	})

// 	return migration
// }

// // migrateModels migrates all models
// func migrateModels(db *gorm.DB, models []models.MigrateItem, logger *providers.LoggerService) error {
// 	for _, item := range models {
// 		logger.Info("Migrating model", zap.String("model", item.ModelName))
// 		if err := db.AutoMigrate(item.Model); err != nil {
// 			logger.Error("Migration failed", zap.String("model", item.ModelName), zap.Error(err))
// 			return err
// 		}
// 	}
// 	return nil
// }

// // seedModels runs all seeders for the models
// func seedModels(db *gorm.DB, models []models.MigrateItem, logger *providers.LoggerService) error {
// 	for _, item := range models {
// 		shouldSeed, err := shouldSeedModel(db, item.Model, logger, item.ModelName)
// 		if err != nil {
// 			logger.Error("Seeder check failed", zap.String("model", item.ModelName), zap.Error(err))
// 			return err
// 		}

// 		if shouldSeed {
// 			logger.Info("Seeding model", zap.String("model", item.ModelName))
// 			if err := item.Seeder(); err != nil {
// 				logger.Error("Seeding failed", zap.String("model", item.ModelName), zap.Error(err))
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

// func shouldSeedModel(db *gorm.DB, model interface{}, logger *providers.LoggerService, modelName string) (bool, error) {
// 	if !db.Migrator().HasTable(model) {
// 		logger.Info(modelName + " table does not exist, skipping seeding.")
// 		return false, nil
// 	}
// 	var count int64
// 	if err := db.Model(model).Count(&count).Error; err != nil {
// 		return false, err
// 	}
// 	if count > 0 {
// 		logger.Info(modelName + " data already exists, skipping seeding.")
// 		return false, nil
// 	}
// 	return true, nil
// }
