package database

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Migrate(lc fx.Lifecycle, db *gorm.DB, logger *zap.Logger, models []interface{}) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Running migrations...")
			db.AutoMigrate(models...)
			return nil
		},
	})
}
