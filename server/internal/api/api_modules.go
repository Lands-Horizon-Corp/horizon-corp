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
		controllers.NewAuthController,
		controllers.NewBranchController,
		controllers.NewCompanyController,
		controllers.NewContactController,
		controllers.NewController,
		controllers.NewFeedbackController,
		controllers.NewFootstepController,
		controllers.NewGenderController,
		controllers.NewMediaController,
		controllers.NewMemberController,
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
