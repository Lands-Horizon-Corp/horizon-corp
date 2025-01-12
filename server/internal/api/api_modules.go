package api

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/controllers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/middleware"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"api",
	fx.Provide(
		controllers.NewAdminController,
		controllers.NewController,
	),
	fx.Provide(
		middleware.NewMiddleware,
		handlers.NewCurrentUser,
		handlers.NewFootstepHandler,
	),
	fx.Invoke(
		NewAPIHandlerInvoke,
	),
)
