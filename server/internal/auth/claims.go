package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	ID   uint   `json:"id"`
	Mode string `json:"mode"`

	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	PermanentAddress  string    `json:"permanentAddress"`
	Description       string    `json:"description"`
	Birthdate         time.Time `json:"birthdate"`
	Email             string    `json:"email"`
	IsEmailVerified   bool      `json:"isEmailVerified"`
	IsContactVerified bool      `json:"isContactVerified"`
	ContactNumber     string    `json:"contacNumber"`
	MediaID           uint      `json:"mediaID"`
	jwt.StandardClaims
}
