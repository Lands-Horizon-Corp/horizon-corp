package database

import (
	"errors"
	"horizon/server/config"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

type CacheService struct {
	Client *redis.Client
}

func NewCacheService(cfg *config.AppConfig, logger *zap.Logger) (*CacheService, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.CacheURL,
		Password: cfg.CachePassword,
		DB:       cfg.CacheDB,
	})

	if err := client.Ping().Err(); err != nil {
		logger.Error("Could not connect to Redis", zap.Error(err))
		return nil, errors.New("could not connect to Redis")
	}

	return &CacheService{Client: client}, nil
}
