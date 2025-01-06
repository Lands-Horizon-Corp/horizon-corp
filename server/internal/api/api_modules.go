package api

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/middleware"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/routes"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"api",
	fx.Provide(
		middleware.NewMiddleware,
		routes.NewAPIRoutes,
	),
	fx.Invoke(
		NewAPIHandlerInvoke,
	),
)
