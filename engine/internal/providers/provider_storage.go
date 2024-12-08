package providers

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"go.uber.org/zap"
)

type StorageProvider struct {
	cfg    *config.AppConfig
	logger *zap.Logger
}

func NewStorageProvider(
	cfg *config.AppConfig,
	logger *zap.Logger,
) *StorageProvider {
	return &StorageProvider{
		cfg:    cfg,
		logger: logger,
	}
}
