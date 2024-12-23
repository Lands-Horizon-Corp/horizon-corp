package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (m *Middleware) getClientIP(c *gin.Context) string {
	if forwarded := c.GetHeader("X-Forwarded-For"); forwarded != "" {
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0])
	}
	return c.ClientIP()
}

func (m *Middleware) BlockIPMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := m.getClientIP(c)

		exists, err := m.cache.Exists("blocked_ip:" + clientIP)
		if err != nil {
			m.logger.Error("Failed to check if IP is blocked", zap.String("ip", clientIP), zap.Error(err))
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
