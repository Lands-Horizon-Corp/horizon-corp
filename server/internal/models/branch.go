package models

import "gorm.io/gorm"

type Branch struct {
	gorm.Model
	Name            string  `gorm:"size:255;not null" json:"name"`
	Address         string  `gorm:"size:500" json:"address"`
	Longitude       float64 `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude        float64 `gorm:"type:decimal(10,7)" json:"latitude"`
	Email           string  `gorm:"size:255;not null;unique" json:"email"`
	ContactNumber   string  `gorm:"size:15;not null;unique" json:"contact_number"`
	IsAdminVerified bool    `gorm:"default:false" json:"is_admin_verified"`
	MediaID         *uint   `gorm:"default:NULL" json:"media_id"`
	Media           Media   `gorm:"foreignKey:MediaID" json:"media"`
	CompanyID       uint    `gorm:"not null" json:"company_id"`
	Company         Company `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"company"`

	Employees []Employee `gorm:"foreignKey:BranchID" json:"employees"`
	Members   []Member   `gorm:"foreignKey:BranchID" json:"members"`
}
