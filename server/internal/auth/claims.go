package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	ID   uint   `json:"id"`
	Mode string `json:"mode"`

	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	PermanentAddress  string    `json:"permanent_address"`
	Description       string    `json:"description"`
	Birthdate         time.Time `json:"birthdate"`
	Email             string    `json:"email"`
	IsEmailVerified   bool      `json:"is_email_verified"`
	IsContactVerified bool      `json:"is_contact_verified"`
	jwt.StandardClaims
}
