package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	FirstName   string `json:"first_name" gorm:"unique;not null"`
	LastName    string `json:"last_name" gorm:"unique;not null"`
	Email       string `json:"email" gorm:"unique;not null"`
	Description string `json:"description,omitempty"`
}
