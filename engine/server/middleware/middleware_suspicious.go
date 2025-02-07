package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (m *Middleware) containsSuspiciousPattern(path string) bool {
	lowerPath := strings.ToLower(path)
	for _, pattern := range m.suspicious {
		if strings.Contains(lowerPath, pattern) {
			return true
		}
	}
	return false
}

func (m *Middleware) DetectSuspiciousAccessMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := m.getClientIP(c)
		requestPath := c.Request.URL.Path
		if m.containsSuspiciousPattern(requestPath) {
			m.logger.Warn("Suspicious file access attempt detected", zap.String("ip", clientIP), zap.String("path", requestPath))
			m.alertAdmin(clientIP, requestPath)
			err := m.setWithRetry("blocked_ip:"+clientIP, "blocked", 24*time.Hour, 3)
			if err != nil {
				m.logger.Error("Failed to block IP in cache", zap.String("ip", clientIP), zap.Error(err))
			}
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Suspicious activity detected. Your IP has been blocked for 24 hours."})
			return
		}
		c.Next()
	}
}
