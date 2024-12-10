package helpers

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"golang.org/x/crypto/bcrypt"
)

const (
	ErrEmptyPassword       = "password cannot be empty"
	ErrEmptyHashedPassword = "hashed password cannot be empty"
	ErrInvalidHash         = "invalid hashed password format"
)

type HelpersCryptography struct {
	cfg *config.AppConfig
}

func NewHelpersCryptography(cfg *config.AppConfig) *HelpersCryptography {
	return &HelpersCryptography{
		cfg: cfg,
	}
}

func (hc *HelpersCryptography) HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New(ErrEmptyPassword)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost+2)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return base64.StdEncoding.EncodeToString(hashedPassword), nil
}

func (hc *HelpersCryptography) VerifyPassword(hashedPassword, password string) bool {
	if hashedPassword == "" {
		return false
	}
	decodedHash, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword(decodedHash, []byte(password))
	return err == nil
}
