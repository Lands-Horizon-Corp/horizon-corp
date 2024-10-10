package auth

import (
	"horizon/server/config"
	"horizon/server/internal/models"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateEmployeeJWT(employee models.Employee) (string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return "", err
	}
	signed, err := config.Decrypt(cfg.AppEmployeeToken, cfg.AppToken)
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &UserClaims{
		ID:                employee.ID,
		Mode:              "employee",
		FirstName:         employee.FirstName,
		LastName:          employee.LastName,
		PermanentAddress:  employee.PermanentAddress,
		Description:       employee.Description,
		Birthdate:         employee.Birthdate,
		Email:             employee.Email,
		IsEmailVerified:   employee.IsEmailVerified,
		IsContactVerified: employee.IsContactVerified,
		ContactNumber:     employee.ContactNumber,
		MediaID:           employee.MediaID,
		StandardClaims: jwt.StandardClaims{
			Subject:   employee.FirstName + " " + employee.LastName,
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signed)
}
