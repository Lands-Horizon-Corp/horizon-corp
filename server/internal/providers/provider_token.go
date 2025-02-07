package providers

import (
	"net/http"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rotisserie/eris"
	"go.uber.org/zap"
)

type UserStatus string

const (
	PendingStatus    UserStatus = "Pending"
	VerifiedStatus   UserStatus = "Verified"
	NotAllowedStatus UserStatus = "Not allowed"
)

type UserClaims struct {
	ID          string     `json:"id"`
	AccountType string     `json:"accountType"`
	UserStatus  UserStatus `json:"userStatus"`
	jwt.StandardClaims
}

type TokenService struct {
	cfg    *config.AppConfig
	cache  *CacheService
	logger *LoggerService
}

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

func (at *TokenService) GenerateUserToken(user UserClaims, expiration time.Duration) (*string, error) {
	validTypes := []string{"Member", "Employee", "Admin", "Owner"}
	isValid := false
	for _, validType := range validTypes {
		if user.AccountType == validType {
			isValid = true
			break
		}
	}
	if !isValid {
		return nil, eris.New("invalid account type")
	}
	claims := &UserClaims{
		ID:          user.ID,
		AccountType: user.AccountType,
	}

	token, err := at.GenerateToken(claims, expiration)
	if err != nil {
		wrappedErr := eris.Wrap(err, "failed to generate user token")
		at.logger.Error(wrappedErr.Error(), zap.Error(err))
		return nil, wrappedErr
	}
	return &token, nil
}

func (s *TokenService) GenerateToken(claims *UserClaims, expiration time.Duration) (string, error) {
	if expiration == 0 {
		expiration = 12 * time.Hour
	}

	claims.StandardClaims = jwt.StandardClaims{
		Issuer:   "horizon-server",
		Subject:  claims.ID,
		IssuedAt: time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.cfg.AppToken))
	if err != nil {
		s.logger.Error("Error signing token", zap.Error(err))
		return "", eris.Wrap(err, "error signing token")
	}

	s.logger.Info("Token generated successfully")
	return tokenString, nil
}

func (s *TokenService) VerifyToken(tokenString string) (*UserClaims, error) {
	if tokenString == "" {
		s.logger.Warn("Empty token provided")
		return nil, eris.New("empty token")
	}

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			s.logger.Warn("Unexpected signing method", zap.String("method", token.Header["alg"].(string)))
			return nil, eris.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.cfg.AppToken), nil
	})

	if err != nil {
		var verr *jwt.ValidationError
		if ok := eris.As(err, &verr); ok {
			if verr.Errors&jwt.ValidationErrorExpired != 0 {
				s.logger.Warn("Token expired")
				return nil, eris.Wrap(err, "token expired")
			}
		}
		s.logger.Error("Error parsing token", zap.Error(err))
		return nil, eris.Wrap(err, "error parsing token")
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		s.logger.Warn("Invalid token")
		return nil, eris.New("invalid token")
	}

	if claims.Issuer != "horizon-server" {
		s.logger.Warn("Invalid issuer", zap.String("issuer", claims.Issuer))
		return nil, eris.Errorf("invalid issuer: %s", claims.Issuer)
	}

	s.logger.Info("Token verified successfully", zap.String("userID", claims.ID))
	return claims, nil
}

func (s *TokenService) StoreToken(tokenString string, userId string) error {
	expiration := 24 * time.Hour
	if err := s.cache.Set(tokenString, userId, expiration); err != nil {
		s.logger.Error("Error storing token in Redis", zap.Error(err))
		return eris.Wrap(err, "error storing token in Redis")
	}
	s.logger.Info("Token reference stored successfully in Redis", zap.String("userId", userId))
	return nil
}

func (s *TokenService) DeleteToken(tokenString string) error {
	if err := s.cache.Delete(tokenString); err != nil {
		s.logger.Error("Error deleting token from Redis", zap.Error(err))
		return eris.Wrap(err, "error deleting token from Redis")
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
