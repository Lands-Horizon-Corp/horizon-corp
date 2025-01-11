package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Branch struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Fields
	Name            string  `gorm:"type:varchar(255);unsigned" json:"name"`
	Address         string  `gorm:"type:varchar(500)" json:"address"`
	PostalCode      string  `gorm:"type:varchar(20)" json:"postal_code"`
	Province        string  `gorm:"type:varchar(255)" json:"province"`
	City            string  `gorm:"type:varchar(255)" json:"city"`
	Barangay        string  `gorm:"type:varchar(255)" json:"barangay"`
	Region          string  `gorm:"type:varchar(255)" json:"region"`
	Longitude       float64 `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude        float64 `gorm:"type:decimal(10,7)" json:"latitude"`
	Email           string  `gorm:"type:varchar(255);unsigned" json:"email"`
	ContactNumber   string  `gorm:"type:varchar(15);unsigned" json:"contact_number"`
	IsAdminVerified bool    `gorm:"default:false" json:"is_admin_verified"`

	// Relationship 0 to 1
	MediaID *uuid.UUID `gorm:"type:bigint;unsigned" json:"media_id"`
	Media   *Media     `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media"`

	// Relationship 0 to 1
	CompanyID *uuid.UUID `gorm:"type:bigint;unsigned;unsigned" json:"company_id"`
	Company   *Company   `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"company"`

	// Relationship 0 to many
	Employees []*Employee `gorm:"foreignKey:BranchID" json:"employees"`

	// Relationship 0 to many
	MemberProfile []*MemberProfile `gorm:"foreignKey:BranchID" json:"members_profile"`
}

type BranchResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	Name       string `json:"name"`
	Address    string `json:"address"`
	PostalCode string `json:"postalCode"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Barangay   string `json:"barangay"`
	Region     string `json:"region"`

	Longitude       float64                  `json:"longitude"`
	Latitude        float64                  `json:"latitude"`
	Email           string                   `json:"email"`
	ContactNumber   string                   `json:"contactNumber"`
	IsAdminVerified bool                     `json:"isAdminVerified"`
	MediaID         *uuid.UUID               `json:"mediaID"`
	Media           *MediaResource           `json:"media"`
	CompanyID       *uuid.UUID               `json:"companyID"`
	Company         *CompanyResource         `json:"company"`
	Employees       []*EmployeeResource      `json:"employees"`
	MemberProfile   []*MemberProfileResource `json:"members"`
}

func (m *ModelTransformer) BranchToResource(branch *Branch) *BranchResource {
	if branch == nil {
		return nil
	}

	return &BranchResource{

		ID:        branch.ID,
		CreatedAt: branch.CreatedAt.Format(time.RFC3339),
		UpdatedAt: branch.UpdatedAt.Format(time.RFC3339),
		DeletedAt: branch.DeletedAt.Time.Format(time.RFC3339),

		Name:       branch.Name,
		Address:    branch.Address,
		PostalCode: branch.PostalCode,
		Province:   branch.Province,
		City:       branch.City,
		Barangay:   branch.Barangay,
		Region:     branch.Region,

		Longitude:       branch.Longitude,
		Latitude:        branch.Latitude,
		Email:           branch.Email,
		ContactNumber:   branch.ContactNumber,
		IsAdminVerified: branch.IsAdminVerified,
		MediaID:         branch.MediaID,
		Media:           m.MediaToResource(branch.Media),
		CompanyID:       branch.CompanyID,
		Company:         m.CompanyToResource(branch.Company),
		Employees:       m.EmployeeToResourceList(branch.Employees),
		MemberProfile:   m.MemberProfileToResourceList(branch.MemberProfile),
	}
}

func (m *ModelTransformer) BranchToResourceList(branchList []*Branch) []*BranchResource {
	if branchList == nil {
		return nil
	}

	var branchResources []*BranchResource
	for _, branch := range branchList {
		branchResources = append(branchResources, m.BranchToResource(branch))
	}
	return branchResources
}
