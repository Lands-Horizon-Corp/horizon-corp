package main

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/Lands-Horizon-Corp/horizon-corp/server"
	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		config.Module,
		helpers.Module,
		providers.Module,
		database.Module,
		modules.Module,
		server.Module,
		fx.Invoke(
			database.NewDatabaseMigration,
			modules.NewModuleServiceProvider,
			providers.NewTerminalService,
		),
	)
	app.Run()
}
