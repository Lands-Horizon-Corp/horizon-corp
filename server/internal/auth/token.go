package auth

import (
	"fmt"
	"horizon/server/config"
	"horizon/server/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

type UserClaims struct {
	ID          uint   `json:"id"`
	AccountType string `json:"accountType"`
	jwt.StandardClaims
}

type TokenService interface {
	GenerateToken(claims *UserClaims) (string, error)
	VerifyToken(tokenString string) (*UserClaims, error)
	StoreToken(tokenString string, userId uint) error
	DeleteToken(tokenString string) error
	ClearTokenCookie(ctx *gin.Context)
}

type tokenService struct {
	redisClient *database.CacheService
	cfg         *config.AppConfig
	logger      *zap.Logger
}

func NewTokenService(redisClient *database.CacheService, cfg *config.AppConfig, logger *zap.Logger) TokenService {
	return &tokenService{
		redisClient: redisClient,
		cfg:         cfg,
		logger:      logger,
	}
}

func (s *tokenService) GenerateToken(claims *UserClaims) (string, error) {
	claims.StandardClaims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	claims.StandardClaims.Issuer = "horizon-server"
	claims.StandardClaims.IssuedAt = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.cfg.AppToken))
	if err != nil {
		s.logger.Error("Error signing token", zap.Error(err))
		return "", fmt.Errorf("error signing token: %w", err)
	}
	s.logger.Info("Token generated successfully", zap.String("token", tokenString))
	return tokenString, nil
}

func (s *tokenService) VerifyToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			s.logger.Warn("Unexpected signing method", zap.String("method", token.Header["alg"].(string)))
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.cfg.AppToken), nil
	})
	if err != nil {
		s.logger.Error("Error parsing token", zap.Error(err))
		return nil, fmt.Errorf("error parsing token: %w", err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		s.logger.Warn("Invalid token", zap.String("token", tokenString))
		return nil, fmt.Errorf("invalid token")
	}
	s.logger.Info("Token verified successfully", zap.Uint("userID", claims.ID))
	return claims, nil
}

func (s *tokenService) StoreToken(tokenString string, userId uint) error {
	err := s.redisClient.Set(tokenString, userId, time.Hour*24)
	if err != nil {
		s.logger.Error("Error storing token in Redis", zap.Error(err))
		return fmt.Errorf("error storing token in Redis: %w", err)
	}
	s.logger.Info("Token stored successfully in Redis", zap.String("token", tokenString), zap.Uint("userId", userId))
	return nil
}

func (s *tokenService) DeleteToken(tokenString string) error {
	err := s.redisClient.Delete(tokenString)
	if err != nil {
		s.logger.Error("Error deleting token from Redis", zap.Error(err))
		return fmt.Errorf("error deleting token from Redis: %w", err)
	}
	s.logger.Info("Token deleted successfully from Redis", zap.String("token", tokenString))
	return nil
}

func (s *tokenService) ClearTokenCookie(ctx *gin.Context) {
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
