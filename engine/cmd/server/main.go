package main

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"go.uber.org/fx"
)

func main() {
	var moduleOptions []fx.Option

	moduleOptions = append(moduleOptions, config.Module)
	moduleOptions = append(moduleOptions, helpers.Module)
	moduleOptions = append(moduleOptions, providers.Module)
	moduleOptions = append(moduleOptions, database.Module)

	app := fx.New(
		fx.Options(moduleOptions...),
	)
	app.Run()
}
