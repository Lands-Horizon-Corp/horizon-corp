package server

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/server/middleware"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/server/routes"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"server",
	fx.Provide(
		middleware.NewMiddleware,
		routes.NewAPIRoutes,
	),
)
