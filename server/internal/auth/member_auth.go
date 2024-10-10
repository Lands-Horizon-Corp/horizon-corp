package auth

import (
	"horizon/server/config"
	"horizon/server/internal/models"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateMemberJWT(member models.Member) (string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return "", err
	}
	signed, err := config.Decrypt(cfg.AppMemberToken, cfg.AppToken)
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &UserClaims{
		Mode:              "member",
		ID:                member.ID,
		FirstName:         member.FirstName,
		LastName:          member.LastName,
		PermanentAddress:  member.PermanentAddress,
		Description:       member.Description,
		Birthdate:         member.Birthdate,
		Email:             member.Email,
		IsEmailVerified:   member.IsEmailVerified,
		IsContactVerified: member.IsContactVerified,
		ContactNumber:     member.ContactNumber,
		MediaID:           member.MediaID,
		StandardClaims: jwt.StandardClaims{
			Subject:   member.FirstName + " " + member.LastName,
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signed)
}
