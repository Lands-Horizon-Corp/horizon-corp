package config

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
)

type TokenService interface {
	GenerateToken(userId uint) (string, error)
	VerifyToken(tokenString string) (uint, error)
	StoreToken(tokenString string, userId uint) error
	DeleteToken(tokenString string) error
}

type tokenService struct {
	redisClient *redis.Client
	jwtSecret   string
}

func NewTokenService(redisClient *redis.Client, jwtSecret string) TokenService {
	return &tokenService{
		redisClient: redisClient,
		jwtSecret:   jwtSecret,
	}
}

func (s *tokenService) GenerateToken(userId uint) (string, error) {
	claims := jwt.StandardClaims{
		Issuer:    "horizon-server",
		Subject:   fmt.Sprintf("%d", userId),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}
	return tokenString, nil
}

func (s *tokenService) VerifyToken(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return 0, fmt.Errorf("error parsing token: %w", err)
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}
	userId, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing user ID from token: %w", err)
	}
	return uint(userId), nil
}

func (s *tokenService) StoreToken(tokenString string, userId uint) error {
	err := s.redisClient.Set(tokenString, userId, time.Hour*24).Err()
	if err != nil {
		return fmt.Errorf("error storing token in Redis: %w", err)
	}
	return nil
}

func (s *tokenService) DeleteToken(tokenString string) error {
	err := s.redisClient.Del(tokenString).Err()
	if err != nil {
		return fmt.Errorf("error deleting token from Redis: %w", err)
	}
	return nil
}
