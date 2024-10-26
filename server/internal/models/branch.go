package models

import "gorm.io/gorm"

type Branch struct {
	gorm.Model
	Name            string  `gorm:"type:varchar(255);not null" json:"name"`
	Address         string  `gorm:"type:varchar(500)" json:"address"`
	Longitude       float64 `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude        float64 `gorm:"type:decimal(10,7)" json:"latitude"`
	Email           string  `gorm:"type:varchar(255);unique;not null" json:"email"`
	ContactNumber   string  `gorm:"type:varchar(15);unique;not null" json:"contact_number"`
	IsAdminVerified bool    `gorm:"default:false" json:"is_admin_verified"`
	MediaID         *uint   `gorm:"type:bigint;unsigned" json:"media_id"`
	Media           *Media  `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media"`
	CompanyID       uint    `gorm:"type:bigint;unsigned;not null" json:"company_id"`
	Company         Company `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"company"`
}
