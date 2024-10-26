package models

import (
	"gorm.io/gorm"
)

type ErrorDetail struct {
	gorm.Model
	Message  string `gorm:"type:varchar(255);not null" json:"message"`
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Stack    string `gorm:"type:text" json:"stack"`
	Response string `gorm:"type:text" json:"response"`
	Status   int    `gorm:"type:int" json:"status"`
}
