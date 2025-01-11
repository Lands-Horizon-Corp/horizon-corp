package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Fields
	Name          string  `gorm:"type:varchar(255);unsigned" json:"name"`
	Description   string  `gorm:"type:text" json:"description"`
	Address       string  `gorm:"type:varchar(500)" json:"address"`
	Longitude     float64 `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude      float64 `gorm:"type:decimal(10,7)" json:"latitude"`
	ContactNumber string  `gorm:"type:varchar(255);unsigned" json:"contact_number"`

	// Relationship 0 to 1
	OwnerID *uuid.UUID `gorm:"type:bigint;unsigned" json:"owner_id"`
	Owner   *Owner     `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"owner"`

	// Relationship 0 to 1
	MediaID         *uuid.UUID `gorm:"type:bigint;unsigned" json:"media_id"`
	Media           *Media     `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media"`
	IsAdminVerified bool       `gorm:"default:false" json:"is_admin_verified"`

	// Relationship 0 to many
	Branches []*Branch `gorm:"foreignKey:CompanyID" json:"branches"`
}

type CompanyResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	Name            string            `json:"name"`
	Description     string            `json:"description"`
	Address         string            `json:"address"`
	Longitude       float64           `json:"longitude"`
	Latitude        float64           `json:"latitude"`
	ContactNumber   string            `json:"contactNumber"`
	OwnerID         *uuid.UUID        `json:"ownerID"`
	Owner           *OwnerResource    `json:"owner"`
	MediaID         *uuid.UUID        `json:"mediaID"`
	Media           *MediaResource    `json:"media"`
	IsAdminVerified bool              `json:"isAdminVerified"`
	Branches        []*BranchResource `json:"branches"`
}

func (m *ModelTransformer) CompanyToResource(company *Company) *CompanyResource {
	if company == nil {
		return nil
	}

	return &CompanyResource{
		ID:        company.ID,
		CreatedAt: company.CreatedAt.Format(time.RFC3339),
		UpdatedAt: company.UpdatedAt.Format(time.RFC3339),
		DeletedAt: company.DeletedAt.Time.Format(time.RFC3339),

		Name:            company.Name,
		Description:     company.Description,
		Address:         company.Address,
		Longitude:       company.Longitude,
		Latitude:        company.Latitude,
		ContactNumber:   company.ContactNumber,
		OwnerID:         company.OwnerID,
		Owner:           m.OwnerToResource(company.Owner),
		MediaID:         company.MediaID,
		Media:           m.MediaToResource(company.Media),
		IsAdminVerified: company.IsAdminVerified,
		Branches:        m.BranchToResourceList(company.Branches),
	}
}

func (m *ModelTransformer) CompanyToResourceList(companyList []*Company) []*CompanyResource {
	if companyList == nil {
		return nil
	}

	var companyResources []*CompanyResource
	for _, company := range companyList {
		companyResources = append(companyResources, m.CompanyToResource(company))
	}
	return companyResources
}
