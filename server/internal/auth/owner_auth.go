package auth

import (
	"horizon/server/config"
	"horizon/server/internal/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type OwnerClaims struct {
	ID   uint   `json:"id"`
	Mode string `json:"mode"`
	jwt.StandardClaims
}

func GenerateOwnerJWT(owner models.Owner) (string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return "", err
	}
	signed, err := config.Decrypt(cfg.AppOwnerToken, cfg.AppToken)
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &OwnerClaims{
		Mode: "owner",
		ID:   owner.ID,
		StandardClaims: jwt.StandardClaims{
			Subject:   owner.FirstName + " " + owner.LastName,
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signed)
}
