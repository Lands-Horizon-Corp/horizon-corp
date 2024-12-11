package server

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/server/middleware"
	"github.com/Lands-Horizon-Corp/horizon-corp/server/routes"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"server",
	fx.Provide(
		middleware.NewMiddleware,
		routes.NewAPIRoutes,
	),
)
