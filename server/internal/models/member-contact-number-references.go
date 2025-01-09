package models

import "gorm.io/gorm"

type MemberContactNumberReferences struct {
	gorm.Model
	Name          string `gorm:"type:varchar(255);not null" json:"name"`
	Description   string `gorm:"type:text" json:"description"`
	ContactNumber string `gorm:"type:varchar(255);not null" json:"contact_number"`
}
