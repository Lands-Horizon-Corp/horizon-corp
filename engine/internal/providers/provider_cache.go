package providers

import (
	"errors"
	"fmt"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/go-redis/redis"

	"go.uber.org/zap"
)

type CacheService struct {
	Client *redis.Client
	cfg    *config.AppConfig
	logger *LoggerService
}

const (
	defaultMaxRetries  = 3
	defaultRetryDelay  = 2 * time.Second
	cacheMissErrString = "cache miss: key does not exist"
)

func NewCacheProvider(
	cfg *config.AppConfig,
	logger *LoggerService,
) (*CacheService, error) {

	if cfg == nil {
		return nil, errors.New("configuration is nil")
	}

	if cfg.CacheURL == "" {
		logger.Error("Cache URL is missing from configuration.")
		return nil, errors.New("cache URL configuration missing")
	}

	maxRetries := cfg.DBMaxRetries
	if maxRetries <= 0 {
		maxRetries = defaultMaxRetries
	}
	retryDelay := cfg.DBRetryDelay
	if retryDelay <= 0 {
		retryDelay = defaultRetryDelay
	}

	logger.Info("Initializing Redis client",
		zap.String("address", cfg.CacheURL),
		zap.Int("db", cfg.CacheDB),
		zap.Int("maxRetries", maxRetries),
		zap.Duration("retryDelay", retryDelay),
	)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.CacheURL,
		Password: cfg.CachePassword,
		DB:       cfg.CacheDB,
	})

	var lastErr error
	for i := 1; i <= maxRetries; i++ {
		if err := redisClient.Ping().Err(); err != nil {
			logger.Warn("Could not connect to Redis",
				zap.Int("attempt", i),
				zap.Error(err),
			)
			lastErr = err
			time.Sleep(retryDelay)
		} else {
			logger.Info("Successfully connected to Redis",
				zap.String("address", cfg.CacheURL))
			return &CacheService{
				Client: redisClient,
				cfg:    cfg,
				logger: logger,
			}, nil
		}
	}
	logger.Error("Could not connect to Redis after multiple attempts",
		zap.Int("maxRetries", maxRetries),
		zap.Error(lastErr))
	return nil, fmt.Errorf("could not connect to Redis after %d attempts: %w", maxRetries, lastErr)
}

func (cs *CacheService) Delete(key string) error {
	err := cs.Client.Del(key).Err()
	if err != nil {
		cs.logger.Error("Failed to delete key from cache",
			zap.String("key", key),
			zap.Error(err))
	}
	return err
}

func (cs *CacheService) Exists(key string) (bool, error) {
	val, err := cs.Client.Exists(key).Result()
	if err != nil {
		cs.logger.Error("Failed to check if key exists in cache",
			zap.String("key", key),
			zap.Error(err))
		return false, err
	}
	return val > 0, nil
}

func (cs *CacheService) Set(key string, value interface{}, expiration time.Duration) error {
	err := cs.Client.Set(key, value, expiration).Err()
	if err != nil {
		cs.logger.Error("Failed to set key in cache",
			zap.String("key", key),
			zap.Error(err))
	}
	return err
}

func (cs *CacheService) Get(key string) (string, error) {
	val, err := cs.Client.Get(key).Result()
	if err == redis.Nil {
		cs.logger.Debug("Cache miss",
			zap.String("key", key))
		return "", errors.New(cacheMissErrString)
	} else if err != nil {
		cs.logger.Error("Failed to get key from cache",
			zap.String("key", key),
			zap.Error(err))
		return "", err
	}
	return val, nil
}

func (cs *CacheService) Increment(key string) (int64, error) {
	val, err := cs.Client.Incr(key).Result()
	if err != nil {
		cs.logger.Error("Failed to increment key in cache",
			zap.String("key", key),
			zap.Error(err))
	}
	return val, err
}

func (cs *CacheService) Decrement(key string) (int64, error) {
	val, err := cs.Client.Decr(key).Result()
	if err != nil {
		cs.logger.Error("Failed to decrement key in cache",
			zap.String("key", key),
			zap.Error(err))
	}
	return val, err
}

// Ping checks the connection to Redis and logs the result
func (cs *CacheService) Ping() error {
	err := cs.Client.Ping().Err()
	if err != nil {
		return err
	}
	return nil
}
