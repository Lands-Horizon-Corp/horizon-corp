package auth

import (
	"horizon/server/config"
	"horizon/server/internal/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type AdminClaims struct {
	ID   uint   `json:"id"`
	Mode string `json:"mode"`
	jwt.StandardClaims
}

func GenerateAdminJWT(admin models.Admin) (string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return "", err
	}
	signed, err := config.Decrypt(cfg.AppAdminToken, cfg.AppToken)
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &AdminClaims{
		Mode: "admin",
		ID:   admin.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signed)
}
