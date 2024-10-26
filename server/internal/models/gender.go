package models

import "gorm.io/gorm"

type Gender struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);unique;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
}
