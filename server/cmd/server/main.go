package main

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/transformers"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		config.Module,
		helpers.Module,
		providers.Module,
		database.Module,
		models.Module,
		api.Module,
		transformers.Module,
		fx.Invoke(
			providers.NewTerminalService,
		),
	)
	app.Run()
}
