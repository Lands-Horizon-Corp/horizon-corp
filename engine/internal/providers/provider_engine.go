package providers

import (
	"context"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

// EngineService holds the Gin engine and HTTP server
type EngineService struct {
	Client *gin.Engine
	server *http.Server
}

// SuspiciousPaths contains patterns of paths that should be blocked
var SuspiciousPaths = []string{
	".env", "env.", ".yaml", ".yml", ".ini", ".config", ".conf", ".xml",
	"dockerfile", "Dockerfile", ".git", ".htaccess", ".htpasswd", "backup",
	"secret", "credential", "password", "private", "key", "token", "dump",
	"database", "db", "logs", "debug",
}

// Helper function to check for suspicious patterns
func containsSuspiciousPattern(path string) bool {
	lowerPath := strings.ToLower(path)
	for _, pattern := range SuspiciousPaths {
		if strings.Contains(lowerPath, pattern) {
			return true
		}
	}
	return false
}

// Helper function to get client IP, considering proxy headers
func getClientIP(c *gin.Context) string {
	if forwarded := c.GetHeader("X-Forwarded-For"); forwarded != "" {
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0])
	}
	return c.ClientIP()
}

// Retry logic for cache operations
func setWithRetry(cache *CacheService, key string, value string, duration time.Duration, retries int, logger *LoggerService) error {
	var err error
	for i := 0; i < retries; i++ {
		err = cache.Set(key, value, duration)
		if err == nil {
			return nil
		}
		logger.Warn("Cache set failed, retrying...", zap.String("key", key), zap.Int("attempt", i+1), zap.Error(err))
		time.Sleep(time.Second * 2)
	}
	return err
}

// Alert function for suspicious activity
func alertAdmin(logger *LoggerService, ip string, path string) {
	logger.Warn("Admin Alert: Suspicious activity detected", zap.String("ip", ip), zap.String("path", path))
}

// Middleware to detect and block suspicious access attempts
func DetectSuspiciousAccessMiddleware(cache *CacheService, logger *LoggerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := getClientIP(c)
		requestPath := c.Request.URL.Path

		if containsSuspiciousPattern(requestPath) {
			logger.Warn("Suspicious file access attempt detected", zap.String("ip", clientIP), zap.String("path", requestPath))
			alertAdmin(logger, clientIP, requestPath)

			// Block the IP for 24 hours
			err := setWithRetry(cache, "blocked_ip:"+clientIP, "blocked", 24*time.Hour, 3, logger)
			if err != nil {
				logger.Error("Failed to block IP in cache", zap.String("ip", clientIP), zap.Error(err))
			}

			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Suspicious activity detected. Your IP has been blocked for 24 hours."})
			return
		}

		c.Next()
	}
}

// Middleware to block requests from blocked IPs
func BlockIPMiddleware(cache *CacheService, logger *LoggerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := getClientIP(c)

		exists, err := cache.Exists("blocked_ip:" + clientIP)
		if err != nil {
			logger.Error("Failed to check if IP is blocked", zap.String("ip", clientIP), zap.Error(err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		if exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Your IP is blocked for 24 hours due to suspicious activity"})
			return
		}

		c.Next()
	}
}

// Per-IP rate limiting middleware
func RateLimiterMiddleware(cache *CacheService, logger *LoggerService, r rate.Limit, b int) gin.HandlerFunc {
	limiters := make(map[string]*rate.Limiter)
	mu := sync.Mutex{}

	getLimiter := func(ip string) *rate.Limiter {
		mu.Lock()
		defer mu.Unlock()
		if limiter, exists := limiters[ip]; exists {
			return limiter
		}
		limiter := rate.NewLimiter(r, b)
		limiters[ip] = limiter
		return limiter
	}

	return func(c *gin.Context) {
		clientIP := getClientIP(c)
		limiter := getLimiter(clientIP)

		if !limiter.Allow() {
			logger.Warn("Too many requests, blocking IP", zap.String("ip", clientIP))

			err := setWithRetry(cache, "blocked_ip:"+clientIP, "blocked", 24*time.Hour, 3, logger)
			if err != nil {
				logger.Error("Failed to block IP in cache", zap.String("ip", clientIP), zap.Error(err))
			}

			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests. Your IP has been blocked for 24 hours."})
			return
		}

		c.Next()
	}
}

// Enforce HTTPS middleware
func enforceHTTPS(c *gin.Context) {
	if c.Request.TLS == nil {
		c.AbortWithStatusJSON(http.StatusUpgradeRequired, gin.H{"error": "HTTPS is required"})
		return
	}
	c.Next()
}

func NewEngineProvider(
	lc fx.Lifecycle,
	cfg *config.AppConfig,
	logger *LoggerService,
	cache *CacheService,
) (*EngineService, error) {
	router := gin.Default()

	router.Use(BlockIPMiddleware(cache, logger))
	router.Use(DetectSuspiciousAccessMiddleware(cache, logger))

	if cfg.AppEnv == "production" || cfg.AppEnv == "staging" {
		router.Use(enforceHTTPS)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://0.0.0.0",
			"http://0.0.0.0:80",
			"http://0.0.0.0:3000",
			"http://0.0.0.0:3001",
			"http://0.0.0.0:4173",
			"http://0.0.0.0:8080",

			// Client Docker
			"http://client",
			"http://client:80",
			"http://client:3000",
			"http://client:3001",
			"http://client:4173",
			"http://client:8080",

			// Localhost
			"http://localhost:",
			"http://localhost:80",
			"http://localhost:3000",
			"http://localhost:3001",
			"http://localhost:4173",
			"http://localhost:8080 ",
		},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(secure.New(secure.Config{
		STSSeconds:           31536000,
		STSIncludeSubdomains: true,
		FrameDeny:            true,
		ContentTypeNosniff:   true,
		BrowserXssFilter:     true,
		SSLProxyHeaders:      map[string]string{"X-Forwarded-Proto": "https"},
	}))

	router.Use(RateLimiterMiddleware(cache, logger, 20, 20))

	router.GET("/favicon.ico", func(c *gin.Context) {
		resp, err := http.Get("https://s3.ap-southeast-2.amazonaws.com/horizon.assets/ecoop-logo.png")
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		c.Header("Content-Type", resp.Header.Get("Content-Type"))

		_, err = io.Copy(c.Writer, resp.Body)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
	})

	// Example route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the secure Gin server!"})
	})

	// HTTP Server Configuration
	server := &http.Server{
		Addr:         ":" + cfg.AppPort,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("Starting Gin server on port http://localhost:" + cfg.AppPort)
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Error("Failed to start server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down Gin server...")
			ctxShutdown, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			if err := server.Shutdown(ctxShutdown); err != nil {
				logger.Error("Server forced to shutdown", zap.Error(err))
			}
			return nil
		},
	})

	return &EngineService{
		Client: router,
		server: server,
	}, nil
}
