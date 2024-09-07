package models

import "gorm.io/gorm"

type UserProfileImage struct {
	gorm.Model
	UserID uint `gorm:"not null;unique"`
	FileID uint `gorm:"not null;unique"`
	User   User `gorm:"foreignKey:UserID"`
	File   File `gorm:"foreignKey:FileID"`
}
