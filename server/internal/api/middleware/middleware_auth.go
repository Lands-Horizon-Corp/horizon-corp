package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie(m.cfg.AppTokenName)
		if err != nil {
			m.tokenProvider.ClearTokenCookie(ctx)
			if errors.Is(err, http.ErrNoCookie) {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: cookie not found"})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
			ctx.Abort()
			return
		}
		claims, err := m.tokenProvider.VerifyToken(cookie.Value)
		if err != nil {
			m.tokenProvider.ClearTokenCookie(ctx)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: token verification failed"})
			ctx.Abort()
			return
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
