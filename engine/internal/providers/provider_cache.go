package providers

import (
	"errors"
	"fmt"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func NewCacheProvider(
	cfg *config.AppConfig,
	logger *zap.Logger,
) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.CacheURL,
		Password: cfg.CachePassword,
		DB:       cfg.CacheDB,
	})

	maxRetries := 3
	retryDelay := 2 * time.Second

	for i := 1; i <= maxRetries; i++ {
		if err := client.Ping().Err(); err != nil {
			logger.Warn(fmt.Sprintf("Attempt %d: Could not connect to Redis", i), zap.Error(err))
			time.Sleep(retryDelay)
		} else {
			logger.Info("Successfully connected to Redis")
			return client, nil
		}
	}

	logger.Error("Could not connect to Redis after multiple attempts")
	return nil, errors.New("could not connect to Redis after multiple attempts")
}
