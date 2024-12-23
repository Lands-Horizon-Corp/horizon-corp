package middleware

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
)

var suspicious = []string{
	".env", "env.", ".yaml", ".yml", ".ini", ".config", ".conf", ".xml",
	"dockerfile", "Dockerfile", ".git", ".htaccess", ".htpasswd", "backup",
	"secret", "credential", "password", "private", "key", "token", "dump",
	"database", "db", "logs", "debug",
}

type Middleware struct {
	cfg           *config.AppConfig
	logger        *providers.LoggerService
	cache         *providers.CacheService
	tokenProvider *providers.TokenService
	suspicious    []string
}

func NewMiddleware(
	cfg *config.AppConfig,
	logger *providers.LoggerService,
	cache *providers.CacheService,
	tokenProvider *providers.TokenService,
) *Middleware {
	return &Middleware{
		cfg:           cfg,
		logger:        logger,
		cache:         cache,
		suspicious:    suspicious,
		tokenProvider: tokenProvider,
	}
}
