package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	AccountType        string    `json:"accountType"`
	FirstName          string    `json:"first_name"`
	LastName           string    `json:"last_name"`
	MiddleName         string    `json:"middle_name"`
	PermanentAddress   string    `json:"permanent_address"`
	Description        string    `json:"description"`
	Birthdate          time.Time `json:"birth_date"`
	Username           string    `json:"username"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	MediaID            *uint     `json:"media_id"`
	IsEmailVerified    bool      `json:"is_email_verified"`
	IsContactVerified  bool      `json:"is_contact_verified"`
	IsSkipVerification bool      `json:"is_skip_verification"`
	ContactNumber      string    `json:"contact_number"`
	Media              *Media    `json:"media"`
	Status             string    `json:"status"`
}
