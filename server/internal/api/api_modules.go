package api

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/broadcast"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/controllers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/middleware"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"api",

	fx.Provide(
		// Broadcast
		broadcast.NewNotificationBroadcast,

		// Controllers
		controllers.NewAdminController,
		controllers.NewAuthController,
		controllers.NewBranchController,
		controllers.NewCompanyController,
		controllers.NewContactController,
		controllers.NewController,
		controllers.NewEmployeeController,
		controllers.NewFeedbackController,
		controllers.NewFootstepController,
		controllers.NewGenderController,
		controllers.NewMediaController,
		controllers.NewMemberController,
		controllers.NewOwnerController,
		controllers.NewProfileController,
		controllers.NewQRScannerController,
		controllers.NewTimesheetController,

		// Handlers
		middleware.NewMiddleware,
		handlers.NewCurrentUser,
		handlers.NewFootstepHandler,
	),
	fx.Invoke(
		NewAPIHandlerInvoke,
	),
)
