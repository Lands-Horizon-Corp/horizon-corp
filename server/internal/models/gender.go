package models

import "gorm.io/gorm"

type Gender struct {
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description,omitempty"`
	gorm.Model
}
