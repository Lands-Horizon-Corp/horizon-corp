package models

import "gorm.io/gorm"

type Company struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`

	// Relation to Owner
	OwnerID string `gorm:"type:uuid"`
	Owner   Owner  `gorm:"foreignKey:OwnerID"` // BelongsTo relation to Owner

	// One-to-many relationship with Branches
	Branches []Branch `gorm:"foreignKey:CompanyID"`

	gorm.Model
}
