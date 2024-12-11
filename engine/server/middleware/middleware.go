package middleware

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
)

var suspicious = []string{
	".env", "env.", ".yaml", ".yml", ".ini", ".config", ".conf", ".xml",
	"dockerfile", "Dockerfile", ".git", ".htaccess", ".htpasswd", "backup",
	"secret", "credential", "password", "private", "key", "token", "dump",
	"database", "db", "logs", "debug",
}

type Middleware struct {
	logger     *providers.LoggerService
	cache      *providers.CacheService
	suspicious []string
}

func NewMiddleware(
	logger *providers.LoggerService,
	cache *providers.CacheService,
) *Middleware {
	return &Middleware{
		logger:     logger,
		cache:      cache,
		suspicious: suspicious,
	}
}
