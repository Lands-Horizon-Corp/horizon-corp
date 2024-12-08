package providers

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"go.uber.org/zap"
)

type DatabaseProvider struct {
	cfg    *config.AppConfig
	logger *zap.Logger
}

func NewDatabaseProvider(
	cfg *config.AppConfig,
	logger *zap.Logger,
) *DatabaseProvider {
	return &DatabaseProvider{
		cfg:    cfg,
		logger: logger,
	}
}
