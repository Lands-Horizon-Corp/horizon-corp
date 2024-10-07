package models

import (
	"gorm.io/gorm"
)

type ErrorDetails struct {
	gorm.Model
	Message  string  `gorm:"type:varchar(255);not null"`
	Name     string  `gorm:"type:varchar(255);not null"`
	Stack    *string `gorm:"type:text"`
	Response *string `gorm:"type:text"`
	Status   *int    `gorm:"type:int"`
}
