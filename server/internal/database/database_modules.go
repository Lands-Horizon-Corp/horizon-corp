package database

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"database",
	fx.Invoke(NewDatabaseMigration),
	fx.Invoke(NewDatabaseSeeder),
)
