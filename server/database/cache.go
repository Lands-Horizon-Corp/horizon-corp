package database

import (
	"errors"
	"fmt"
	"horizon/server/config"
	"time"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

// ErrCacheMiss is returned when a cache key does not exist.
var ErrCacheMiss = errors.New("cache miss: key does not exist")

// CacheService provides methods to interact with Redis.
type CacheService struct {
	Client *redis.Client
}

// NewCacheService constructs a new CacheService with an optional retry mechanism.
func NewCacheService(cfg *config.AppConfig, logger *zap.Logger) (*CacheService, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.CacheURL,
		Password: cfg.CachePassword,
		DB:       cfg.CacheDB,
	})

	maxRetries := 3
	retryDelay := 2 * time.Second

	// Retry mechanism for Redis connection
	for i := 1; i <= maxRetries; i++ {
		if err := client.Ping().Err(); err != nil {
			logger.Warn(fmt.Sprintf("Attempt %d: Could not connect to Redis", i), zap.Error(err))
			time.Sleep(retryDelay)
		} else {
			logger.Info("Successfully connected to Redis")
			return &CacheService{Client: client}, nil
		}
	}

	// Log and return error if all connection attempts fail
	logger.Error("Could not connect to Redis after multiple attempts")
	return nil, errors.New("could not connect to Redis after multiple attempts")
}

// Set stores a key-value pair in the cache with an optional expiration time.
func (cs *CacheService) Set(key string, value interface{}, expiration time.Duration) error {
	return cs.Client.Set(key, value, expiration).Err()
}

// Get retrieves a value from the cache by key.
func (cs *CacheService) Get(key string) (string, error) {
	val, err := cs.Client.Get(key).Result()
	if err == redis.Nil {
		return "", ErrCacheMiss
	} else if err != nil {
		return "", err
	}
	return val, nil
}

// Delete removes a key from the cache.
func (cs *CacheService) Delete(key string) error {
	return cs.Client.Del(key).Err()
}

// Exists checks if a key exists in the cache.
func (cs *CacheService) Exists(key string) (bool, error) {
	val, err := cs.Client.Exists(key).Result()
	if err != nil {
		return false, err
	}
	return val > 0, nil
}

// Increment increases the integer value of a key by one.
func (cs *CacheService) Increment(key string) (int64, error) {
	return cs.Client.Incr(key).Result()
}

// Decrement decreases the integer value of a key by one.
func (cs *CacheService) Decrement(key string) (int64, error) {
	return cs.Client.Decr(key).Result()
}

// StoreOTP is a specialized method for storing OTPs with an expiration time.
func (cs *CacheService) StoreOTP(key, otp string, expiration time.Duration) error {
	return cs.Set(key, otp, expiration)
}

// GetOTP retrieves an OTP value from the cache by key.
func (cs *CacheService) GetOTP(key string) (string, error) {
	return cs.Get(key)
}

// DeleteOTP removes an OTP key from the cache.
func (cs *CacheService) DeleteOTP(key string) error {
	return cs.Delete(key)
}
