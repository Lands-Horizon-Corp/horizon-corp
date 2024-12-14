package providers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

type UserClaims struct {
	ID          uint   `json:"id"`
	AccountType string `json:"accountType"`
	jwt.StandardClaims
}

type TokenService struct {
	cfg    *config.AppConfig
	cache  *CacheService
	logger *LoggerService
}

// NewTokenProvider returns a new instance of TokenService.
func NewTokenProvider(
	cfg *config.AppConfig,
	cache *CacheService,
	logger *LoggerService,
) *TokenService {
	return &TokenService{
		cfg:    cfg,
		cache:  cache,
		logger: logger,
	}
}

// GenerateToken generates a new JWT with the provided claims and expiration.
func (s *TokenService) GenerateToken(claims *UserClaims, expiration time.Duration) (string, error) {
	if expiration == 0 {
		expiration = 12 * time.Hour
	}

	// Set registered claims
	claims.StandardClaims = jwt.StandardClaims{
		Issuer:   "horizon-server",
		Subject:  fmt.Sprintf("%d", claims.ID),
		IssuedAt: time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.cfg.AppToken))
	if err != nil {
		s.logger.Error("Error signing token", zap.Error(err))
		return "", fmt.Errorf("error signing token: %w", err)
	}

	// Don't log the entire token. If needed, log a substring or hash.
	s.logger.Info("Token generated successfully")

	return tokenString, nil
}

// VerifyToken verifies and parses the provided JWT token string.
func (s *TokenService) VerifyToken(tokenString string) (*UserClaims, error) {
	if tokenString == "" {
		s.logger.Warn("Empty token provided")
		return nil, fmt.Errorf("empty token")
	}

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			s.logger.Warn("Unexpected signing method", zap.String("method", token.Header["alg"].(string)))
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.cfg.AppToken), nil
	})

	if err != nil {
		var verr *jwt.ValidationError
		if ok := errors.As(err, &verr); ok {
			if verr.Errors&jwt.ValidationErrorExpired != 0 {
				s.logger.Warn("Token expired")
				return nil, fmt.Errorf("token expired: %w", err)
			}
		}
		s.logger.Error("Error parsing token", zap.Error(err))
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		s.logger.Warn("Invalid token")
		return nil, fmt.Errorf("invalid token")
	}

	if claims.Issuer != "horizon-server" {
		s.logger.Warn("Invalid issuer", zap.String("issuer", claims.Issuer))
		return nil, fmt.Errorf("invalid issuer")
	}

	s.logger.Info("Token verified successfully", zap.Uint("userID", claims.ID))
	return claims, nil
}

// StoreToken stores a token-related value in the cache.
// Consider storing JTI (unique token ID) rather than the entire token string.
func (s *TokenService) StoreToken(tokenString string, userId uint) error {
	expiration := 24 * time.Hour
	if err := s.cache.Set(tokenString, userId, expiration); err != nil {
		s.logger.Error("Error storing token in Redis", zap.Error(err))
		return fmt.Errorf("error storing token in Redis: %w", err)
	}
	s.logger.Info("Token reference stored successfully in Redis", zap.Uint("userId", userId))
	return nil
}

func (s *TokenService) DeleteToken(tokenString string) error {
	if err := s.cache.Delete(tokenString); err != nil {
		s.logger.Error("Error deleting token from Redis", zap.Error(err))
		return fmt.Errorf("error deleting token from Redis: %w", err)
	}
	s.logger.Info("Token reference deleted successfully from Redis")
	return nil
}

func (s *TokenService) ClearTokenCookie(ctx *gin.Context) {
	cookie, err := ctx.Cookie(s.cfg.AppTokenName)
	if err == nil && cookie != "" {
		if delErr := s.DeleteToken(cookie); delErr != nil {
			s.logger.Error("Failed to delete token from Redis", zap.Error(delErr))
		}
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     s.cfg.AppTokenName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteNoneMode,
	})
}
