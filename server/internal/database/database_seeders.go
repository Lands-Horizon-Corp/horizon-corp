package database

import (
	"context"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"go.uber.org/fx"
)

type DatabaseSeeder struct {
	lc     fx.Lifecycle
	db     *providers.DatabaseService
	logger *providers.LoggerService
}

func NewDatabaseSeeder(
	lc fx.Lifecycle,
	db *providers.DatabaseService,
	logger *providers.LoggerService,
) *DatabaseSeeder {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Running database seeders...")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Seeder lifecycle stopped.")
			return nil
		},
	})
	return &DatabaseSeeder{
		lc:     lc,
		db:     db,
		logger: logger,
	}
}
