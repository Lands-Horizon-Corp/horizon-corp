package auth

import (
	"horizon/server/config"
	"horizon/server/internal/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type EmployeeClaims struct {
	ID   uint   `json:"id"`
	Mode string `json:"mode"`
	jwt.StandardClaims
}

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
	claims := &EmployeeClaims{
		Mode: "employee",
		ID:   employee.ID,
		StandardClaims: jwt.StandardClaims{
			Subject:   employee.FirstName + " " + employee.LastName,
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signed)
}
