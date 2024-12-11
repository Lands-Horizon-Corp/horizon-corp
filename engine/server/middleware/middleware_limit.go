package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

func (m *Middleware) RateLimiterMiddleware(r rate.Limit, b int) gin.HandlerFunc {
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
		clientIP := m.getClientIP(c)
		limiter := getLimiter(clientIP)
		if !limiter.Allow() {
			m.logger.Warn("Too many requests, blocking IP", zap.String("ip", clientIP))
			err := m.setWithRetry("blocked_ip:"+clientIP, "blocked", 24*time.Hour, 3)
			if err != nil {
				m.logger.Error("Failed to block IP in cache", zap.String("ip", clientIP), zap.Error(err))
			}
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests. Your IP has been blocked for 24 hours."})
			return
		}
		c.Next()
	}
}
