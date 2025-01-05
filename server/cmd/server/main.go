package main

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/tags"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		config.Module,
		helpers.Module,
		providers.Module,
		database.Module,
		tags.Module,
		models.Module,
		fx.Invoke(
			// 	modules.NewModuleServiceProvider,
			providers.NewTerminalService,
		),
	)
	app.Run()
}
