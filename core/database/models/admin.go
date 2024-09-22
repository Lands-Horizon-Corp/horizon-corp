package models

import (
	"time"
)

type Admin struct {
	ID               string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email            string    `gorm:"type:varchar(255);not null;unique"`
	Password         string    `gorm:"type:varchar(255);not null"`
	ProfilePicture   string    `gorm:"type:uuid;null"` // Foreign key reference to Media
	Username         string    `gorm:"type:varchar(255);not null"`
	FirstName        string    `gorm:"type:varchar(255);not null"`
	LastName         string    `gorm:"type:varchar(255);not null"`
	ContactNumber    string    `gorm:"type:varchar(50)"`
	PermanentAddress string    `gorm:"type:varchar(255)"`
	Description      string    `gorm:"type:uuid"` // Assuming uuid here for linking to description table or something similar
	Birthdate        time.Time `gorm:"type:date"`
}
