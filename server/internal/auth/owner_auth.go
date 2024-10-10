package auth

import (
	"horizon/server/config"
	"horizon/server/internal/models"
	"time"

	"github.com/golang-jwt/jwt"
)

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
	claims := &UserClaims{
		Mode:              "owner",
		ID:                owner.ID,
		FirstName:         owner.FirstName,
		LastName:          owner.LastName,
		PermanentAddress:  owner.PermanentAddress,
		Description:       owner.Description,
		Birthdate:         owner.Birthdate,
		Email:             owner.Email,
		IsEmailVerified:   owner.IsEmailVerified,
		IsContactVerified: owner.IsContactVerified,
		ContactNumber:     owner.ContactNumber,
		MediaID:           owner.MediaID,
		StandardClaims: jwt.StandardClaims{
			Subject:   owner.FirstName + " " + owner.LastName,
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signed)
}
