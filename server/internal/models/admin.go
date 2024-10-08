package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	FirstName        string    `gorm:"size:255;not null"`
	LastName         string    `gorm:"size:255;not null"`
	PermanentAddress string    `gorm:"size:500"`
	Description      string    `gorm:"size:1000"`
	Birthdate        time.Time `gorm:"not null"`
	Username         string    `gorm:"size:255;not null;unique"`
	Email            string    `gorm:"size:255;not null;unique"`
	Password         string    `gorm:"size:255;not null"`
	MediaID          uint      `gorm:"not null"`           // Foreign key reference to Media
	Media            Media     `gorm:"foreignKey:MediaID"` // Relationship
}
