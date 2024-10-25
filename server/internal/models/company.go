package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name            string  `gorm:"size:255;not null" json:"name"`
	Description     string  `json:"description,omitempty"`
	Address         string  `gorm:"size:500" json:"address"`
	Longitude       float64 `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude        float64 `gorm:"type:decimal(10,7)" json:"latitude"`
	ContactNumber   string  `gorm:"size:255;not null;unique" json:"contact_number"`
	OwnerID         uint    `gorm:"not null" json:"owner_id"`
	MediaID         *uint   `gorm:"default:NULL" json:"media_id"`
	IsAdminVerified bool    `gorm:"default:false" json:"is_admin_verified"`

	Media    Media    `gorm:"foreignKey:MediaID" json:"media"`
	Owner    Owner    `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"owner"`
	Branches []Branch `gorm:"foreignKey:CompanyID" json:"branches"`
}
