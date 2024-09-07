package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user entity with personal and authentication details.
type User struct {
	gorm.Model
	FirstName        string    `gorm:"size:255;not null"`
	LastName         string    `gorm:"size:255;not null"`
	PermanentAddress string    `gorm:"size:500"`
	Description      string    `gorm:"size:1000"`
	Birthdate        time.Time `gorm:"not null"`
	Username         string    `gorm:"size:255;not null;unique"`
	Email            string    `gorm:"size:255;not null;unique"`
	Password         string    `gorm:"size:255;not null"`
}
