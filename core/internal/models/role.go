package models

import "gorm.io/gorm"

type Role struct {
	ID          string       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string       `gorm:"type:varchar(255);not null"`
	Description string       `gorm:"type:text"`
	Permissions []Permission `gorm:"foreignKey:RoleID"`

	gorm.Model
}
