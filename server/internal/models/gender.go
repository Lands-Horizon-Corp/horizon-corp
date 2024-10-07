package models

import "gorm.io/gorm"

type Gender struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description,omitempty"`
	gorm.Model
}
