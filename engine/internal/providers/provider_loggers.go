package providers

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"go.uber.org/zap"
)

func NewLoggerProvider(
	cfg *config.AppConfig,
) (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return logger, nil
}
