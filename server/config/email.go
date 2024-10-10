package config

import (
	"horizon/server/internal/models"
)

func SendVerificationEmail(email, token string) error {
	// Email sending logic here
	return nil
}

// helpers/token.go
func GenerateEmailVerificationToken(admin models.Admin) (string, error) {
	return "passed", nil
}

func ParseEmailVerificationToken(tokenString string) (uint, error) {
	// Similar to ParseJWT but with the email verification signing key
	return 1, nil
}
