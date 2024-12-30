package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"golang.org/x/crypto/bcrypt"
)

const (
	ErrEmptyPassword       = "password cannot be empty"
	ErrEmptyHashedPassword = "hashed password cannot be empty"
	ErrInvalidHash         = "invalid hashed password format"
)

// HelpersCryptography handles cryptographic operations.
type HelpersCryptography struct {
	cfg *config.AppConfig
}

// NewHelpersCryptography creates a new instance of HelpersCryptography.
func NewHelpersCryptography(cfg *config.AppConfig) *HelpersCryptography {
	return &HelpersCryptography{cfg: cfg}
}

// HashPassword hashes the given password using bcrypt.
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

// VerifyPassword compares a hashed password with a plain password.
func (hc *HelpersCryptography) VerifyPassword(hashedPassword, password string) bool {
	if hashedPassword == "" || password == "" {
		return false
	}

	decodedHash, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return false
	}

	return bcrypt.CompareHashAndPassword(decodedHash, []byte(password)) == nil
}

// GenerateSecureRandom6DigitNumber generates a secure random 6-digit number.
func (hc *HelpersCryptography) GenerateSecureRandom6DigitNumber() (int, error) {
	const min = 100000
	const max = 999999

	n, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
	if err != nil {
		return 0, fmt.Errorf("failed to generate random number: %w", err)
	}

	return int(n.Int64() + min), nil
}