package models

import "gorm.io/gorm"

type MemberContactNumberReferences struct {
	gorm.Model
	Name          string `gorm:"type:varchar(255);not null" json:"name"`
	Description   string `gorm:"type:text" json:"description"`
	ContactNumber string `gorm:"type:varchar(255);not null" json:"contact_number"`
}

type MemberContactNumberReferencesResource struct {
	ID            uint   `json:"id"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	ContactNumber string `json:"contactNumber"`
}
