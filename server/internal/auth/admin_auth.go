package auth

import (
	"horizon/server/config"
	"horizon/server/internal/models"
	"time"

	"github.com/golang-jwt/jwt"
)

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
	claims := &UserClaims{
		ID:                admin.ID,
		Mode:              "admin",
		FirstName:         admin.FirstName,
		LastName:          admin.LastName,
		PermanentAddress:  admin.PermanentAddress,
		Description:       admin.Description,
		Birthdate:         admin.Birthdate,
		Email:             admin.Email,
		IsEmailVerified:   admin.IsEmailVerified,
		IsContactVerified: admin.IsContactVerified,
		StandardClaims: jwt.StandardClaims{
			Subject:   admin.FirstName + " " + admin.LastName,
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signed)
}
