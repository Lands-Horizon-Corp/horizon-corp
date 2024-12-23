package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) EnforceHTTPS(c *gin.Context) {
	if c.Request.TLS == nil {
		c.AbortWithStatusJSON(http.StatusUpgradeRequired, gin.H{"error": "HTTPS is required"})
		return
	}
	c.Next()
}
