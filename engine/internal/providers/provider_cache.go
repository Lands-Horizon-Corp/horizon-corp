package providers

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

type CacheProvider struct {
	client *redis.Client
	cfg    *config.AppConfig
	logger *zap.Logger
}

func NewCacheProvider(
	client *redis.Client,
	cfg *config.AppConfig,
	logger *zap.Logger,
) *CacheProvider {
	return &CacheProvider{
		client: client,
		cfg:    cfg,
		logger: logger,
	}
}
