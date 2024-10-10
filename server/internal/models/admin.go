package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	FirstName         string    `gorm:"size:255;not null" json:"first_name"`
	LastName          string    `gorm:"size:255;not null" json:"last_name"`
	PermanentAddress  string    `gorm:"size:500" json:"permanent_address"`
	Description       string    `gorm:"size:1000" json:"description"`
	Birthdate         time.Time `gorm:"not null" json:"birthdate"`
	Username          string    `gorm:"size:255;not null;unique" json:"username"`
	Email             string    `gorm:"size:255;not null;unique" json:"email"`
	Password          string    `gorm:"size:255;not null" json:"password"`
	MediaID           uint      `gorm:"not null" json:"media_id"`
	IsEmailVerified   bool      `gorm:"default:false" json:"is_email_verified"`
	IsContactVerified bool      `gorm:"default:false" json:"is_contact_verified"`
	ContactNumber     string    `json:"contac_number,omitempty"`
	Media             Media     `gorm:"foreignKey:MediaID" json:"media"`
}
