package internal

import (
	"context"
	"horizon/server/config"
	"horizon/server/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func StartServer(lc fx.Lifecycle, router *gin.Engine, logger *zap.Logger, cfg *config.AppConfig, smsService services.SMSService) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := smsService.SendSMS(ctx, cfg.ContactNumber, "Hello from AWS SNS via ECOOP!")
				if err != nil {
					logger.Error("Error sending test SMS", zap.Error(err))
				} else {
					logger.Info("Test SMS sent successfully")
				}
				if err := router.Run(":" + cfg.AppPort); err != nil {
					logger.Error("Failed to start server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping the server...")
			return nil
		},
	})
}
