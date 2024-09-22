package main

import (
	"horizon-core/config"
	"horizon-core/database"
	"horizon-core/internal/handler"
	"horizon-core/internal/router"
	"horizon-core/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			// Dependencies
			config.ProvideLogger,
			database.ProvideDatabase,

			// File Storage
			config.GetMediaClient,

			// Services and Handlers
			service.NewAdminService,
			handler.NewAdminHandler,
			service.NewBranchService,
			handler.NewBranchHandler,
			service.NewCompanyService,
			handler.NewCompanyHandler,
			service.NewEmployeeService,
			handler.NewEmployeeHandler,
			service.NewMediaService,
			handler.NewMediaHandler,
			service.NewMemberService,
			handler.NewMemberHandler,
			service.NewOwnerService,
			handler.NewOwnerHandler,
			service.NewPermissionService,
			handler.NewPermissionHandler,
			service.NewRoleService,
			handler.NewRoleHandler,

			// API
			router.ProvideAPI,
		),
		fx.Invoke(
			func(router *gin.Engine, cfg *config.Config) {
				router.Run(":" + cfg.App.Port)
			},
			func(logger *zap.Logger) {
				logger.Info("Application starting up")
			},
		),
	).Run()
}
