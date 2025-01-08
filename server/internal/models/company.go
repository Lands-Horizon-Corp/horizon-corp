package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model

	// Fields
	Name          string  `gorm:"type:varchar(255);not null" json:"name"`
	Description   string  `gorm:"type:text" json:"description"`
	Address       string  `gorm:"type:varchar(500)" json:"address"`
	Longitude     float64 `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude      float64 `gorm:"type:decimal(10,7)" json:"latitude"`
	ContactNumber string  `gorm:"type:varchar(255);not null" json:"contact_number"`

	// Relationship 0 to 1
	OwnerID *uint  `gorm:"type:bigint;unsigned" json:"owner_id"`
	Owner   *Owner `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"owner"`

	// Relationship 0 to 1
	MediaID         *uint  `gorm:"type:bigint;unsigned" json:"media_id"`
	Media           *Media `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media"`
	IsAdminVerified bool   `gorm:"default:false" json:"is_admin_verified"`

	// Relationship 0 to many
	Branches []*Branch `gorm:"foreignKey:CompanyID" json:"branches"`
}

type CompanyResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name            string            `json:"name"`
	Description     string            `json:"description"`
	Address         string            `json:"address"`
	Longitude       float64           `json:"longitude"`
	Latitude        float64           `json:"latitude"`
	ContactNumber   string            `json:"contactNumber"`
	OwnerID         *uint             `json:"ownerID"`
	Owner           *OwnerResource    `json:"owner"`
	MediaID         *uint             `json:"mediaID"`
	Media           *MediaResource    `json:"media"`
	IsAdminVerified bool              `json:"isAdminVerified"`
	Branches        []*BranchResource `json:"branches"`
}
