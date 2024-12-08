package providers

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"go.uber.org/zap"
)

type LoggerProvider struct {
	cfg    *config.AppConfig
	logger *zap.Logger
}

func NewLoggerProvider(
	cfg *config.AppConfig,
	logger *zap.Logger,
) *LoggerProvider {
	return &LoggerProvider{
		cfg:    cfg,
		logger: logger,
	}
}
