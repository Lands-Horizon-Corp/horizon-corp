package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name            string  `gorm:"type:varchar(255);not null" json:"name"`
	Description     string  `gorm:"type:text" json:"description"`
	Address         string  `gorm:"type:varchar(500)" json:"address"`
	Longitude       float64 `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude        float64 `gorm:"type:decimal(10,7)" json:"latitude"`
	ContactNumber   string  `gorm:"type:varchar(255);unique;not null" json:"contact_number"`
	OwnerID         *uint   `gorm:"type:bigint;unsigned" json:"owner_id"`
	Owner           *Owner  `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"owner"`
	MediaID         *uint   `gorm:"type:bigint;unsigned" json:"media_id"`
	Media           *Media  `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media"`
	IsAdminVerified bool    `gorm:"default:false" json:"is_admin_verified"`

	Branches []*Branch `gorm:"foreignKey:CompanyID" json:"branches"`
}
