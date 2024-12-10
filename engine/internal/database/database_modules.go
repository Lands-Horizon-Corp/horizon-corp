package database

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"database",
	fx.Provide(models.NewModelResource),
	fx.Invoke(NewDatabaseMigration),
)
