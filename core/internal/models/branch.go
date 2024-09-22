package models

import "gorm.io/gorm"

type Branch struct {
	ID            string  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name          string  `gorm:"type:varchar(255);not null"`
	Email         string  `gorm:"type:varchar(255);not null"`
	Address       string  `gorm:"type:varchar(255);not null"`
	ContactNumber string  `gorm:"type:varchar(255);not null"`
	Approved      bool    `gorm:"default:false"`
	Description   string  `gorm:"type:text"`
	Theme         string  `gorm:"type:text;default:'default'"`
	CompanyID     string  `gorm:"type:uuid"`            // Foreign key for Company
	Company       Company `gorm:"foreignKey:CompanyID"` // BelongsTo relation to Company

	// Foreign key to Media table (optional)
	MediaID        *string `gorm:"type:uuid;null"`
	ProfilePicture Media   `gorm:"foreignKey:MediaID"`

	gorm.Model
}
