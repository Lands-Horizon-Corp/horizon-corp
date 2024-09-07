package models

import "gorm.io/gorm"

type UserAttachments struct {
	gorm.Model
	UserID uint `gorm:"not null"`
	FileID uint `gorm:"not null"`
	User   User `gorm:"foreignKey:UserID"`
	File   File `gorm:"foreignKey:FileID"`
}
