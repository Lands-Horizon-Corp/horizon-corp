package auth

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

type UserClaims struct {
	ID                uint      `json:"id"`
	Mode              string    `json:"mode"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	MiddleName        string    `json:"middleName"`
	PermanentAddress  string    `json:"permanentAddress"`
	Description       string    `json:"description"`
	Birthdate         time.Time `json:"birthdate"`
	Email             string    `json:"email"`
	IsEmailVerified   bool      `json:"isEmailVerified"`
	IsContactVerified bool      `json:"isContactVerified"`
	ContactNumber     string    `json:"contactNumber"`
	MediaID           *uint     `json:"mediaID"`
	jwt.StandardClaims
}

type TokenService interface {
	GenerateToken(claims *UserClaims) (string, error)
	VerifyToken(tokenString string) (uint, error)
	StoreToken(tokenString string, userId uint) error
	DeleteToken(tokenString string) error
}

type tokenService struct {
	redisClient *redis.Client
	jwtSecret   string
	logger      *zap.Logger
}

func NewTokenService(redisClient *redis.Client, jwtSecret string, logger *zap.Logger) TokenService {
	return &tokenService{
		redisClient: redisClient,
		jwtSecret:   jwtSecret,
		logger:      logger,
	}
}

func (s *tokenService) GenerateToken(claims *UserClaims) (string, error) {
	claims.StandardClaims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	claims.StandardClaims.Issuer = "horizon-server"
	claims.StandardClaims.IssuedAt = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		s.logger.Error("Error signing token", zap.Error(err))
		return "", fmt.Errorf("error signing token: %w", err)
	}
	s.logger.Info("Token generated successfully", zap.String("token", tokenString))
	return tokenString, nil
}

func (s *tokenService) VerifyToken(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			s.logger.Warn("Unexpected signing method", zap.String("method", token.Header["alg"].(string)))
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		s.logger.Error("Error parsing token", zap.Error(err))
		return 0, fmt.Errorf("error parsing token: %w", err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		s.logger.Warn("Invalid token", zap.String("token", tokenString))
		return 0, fmt.Errorf("invalid token")
	}
	s.logger.Info("Token verified successfully", zap.Uint("userID", claims.ID))
	return claims.ID, nil
}

func (s *tokenService) StoreToken(tokenString string, userId uint) error {
	err := s.redisClient.Set(tokenString, userId, time.Hour*24).Err()
	if err != nil {
		s.logger.Error("Error storing token in Redis", zap.Error(err))
		return fmt.Errorf("error storing token in Redis: %w", err)
	}
	s.logger.Info("Token stored successfully in Redis", zap.String("token", tokenString), zap.Uint("userId", userId))
	return nil
}

func (s *tokenService) DeleteToken(tokenString string) error {
	err := s.redisClient.Del(tokenString).Err()
	if err != nil {
		s.logger.Error("Error deleting token from Redis", zap.Error(err))
		return fmt.Errorf("error deleting token from Redis: %w", err)
	}
	s.logger.Info("Token deleted successfully from Redis", zap.String("token", tokenString))
	return nil
}
