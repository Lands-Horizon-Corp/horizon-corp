package middleware

import (
	"context"
	"fmt"
	"horizon/server/config"
	"horizon/server/internal/auth"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type middleware string

const (
	AuthorizationHeader middleware = "Authorization"
	BearerPrefix        middleware = "Bearer "
	ClaimsKey           middleware = "claims"
)

func JWTMiddleware(mode string, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(string(AuthorizationHeader))
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, string(BearerPrefix)) {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, string(BearerPrefix))

		cfg, err := config.LoadConfig()
		if err != nil {
			http.Error(w, "Could not load config", http.StatusInternalServerError)
			return
		}

		modeToken, err := getTokenMode(mode)
		if err != nil {
			http.Error(w, "Could not load mode token", http.StatusInternalServerError)
			return
		}

		signed, err := config.Decrypt(modeToken, cfg.AppToken)
		if err != nil {
			http.Error(w, "Could not decrypt token", http.StatusInternalServerError)
			return
		}

		claims := &auth.UserClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return signed, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if claims.ExpiresAt < time.Now().Unix() {
			http.Error(w, "Token has expired", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, ClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func getTokenMode(mode string) ([]byte, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	switch mode {
	case "admin":
		return cfg.AppAdminToken, nil
	case "owner":
		return cfg.AppOwnerToken, nil
	case "member":
		return cfg.AppMemberToken, nil
	case "employee":
		return cfg.AppEmployeeToken, nil
	default:
		return nil, fmt.Errorf("unrecognized mode: %s", mode)
	}
}
