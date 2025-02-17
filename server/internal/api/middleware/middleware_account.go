package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
)

func (m *Middleware) AccountTypeMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie(m.cfg.AppTokenName)
		if err != nil {
			m.handleUnauthorized(ctx, err)
			return
		}
		claims, err := m.tokenProvider.VerifyToken(cookie.Value)
		if err != nil {
			m.handleUnauthorized(ctx, err)
			return
		}
		if !contains(requiredRoles, claims.AccountType) {
			m.handleForbidden(ctx, requiredRoles)
			return
		}
		if claims.UserStatus != providers.VerifiedStatus {
			m.handleNotVerified(ctx)
			return
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}

func (m *Middleware) handleUnauthorized(ctx *gin.Context, err error) {
	m.tokenProvider.ClearTokenCookie(ctx)
	if errors.Is(err, http.ErrNoCookie) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: cookie not found"})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: token invalid or expired"})
	}
	ctx.Abort()
}

func (m *Middleware) handleForbidden(ctx *gin.Context, roles []string) {
	ctx.JSON(http.StatusForbidden, gin.H{
		"error": fmt.Sprintf("Forbidden: %s access required", strings.Join(roles, " or ")),
	})
	m.tokenProvider.ClearTokenCookie(ctx)
	ctx.Abort()
}

func (m *Middleware) handleNotVerified(ctx *gin.Context) {
	ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: user must be verified"})
	ctx.Abort()
}

func contains(requiredRoles []string, role string) bool {
	for _, r := range requiredRoles {
		if r == role {
			return true
		}
	}
	return false
}
