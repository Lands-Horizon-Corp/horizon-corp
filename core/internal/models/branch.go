package models

import "gorm.io/gorm"

type Branch struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	gorm.Model
}
