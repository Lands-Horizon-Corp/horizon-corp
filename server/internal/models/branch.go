package models

import (
	"time"

	"gorm.io/gorm"
)

type Branch struct {
	gorm.Model

	// Fields
	Name            string  `gorm:"type:varchar(255);unsigned" json:"name"`
	Address         string  `gorm:"type:varchar(500)" json:"address"`
	Longitude       float64 `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude        float64 `gorm:"type:decimal(10,7)" json:"latitude"`
	Email           string  `gorm:"type:varchar(255);unsigned" json:"email"`
	ContactNumber   string  `gorm:"type:varchar(15);unsigned" json:"contact_number"`
	IsAdminVerified bool    `gorm:"default:false" json:"is_admin_verified"`

	// Relationship 0 to 1
	MediaID *uint  `gorm:"type:bigint;unsigned" json:"media_id"`
	Media   *Media `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media"`

	// Relationship 0 to 1
	CompanyID *uint    `gorm:"type:bigint;unsigned;unsigned" json:"company_id"`
	Company   *Company `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"company"`

	// Relationship 0 to many
	Employees []*Employee `gorm:"foreignKey:BranchID" json:"employees"`

	// Relationship 0 to many
	Members []*Member `gorm:"foreignKey:BranchID" json:"members"`
}

type BranchResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name            string              `json:"name"`
	Address         string              `json:"address"`
	Longitude       float64             `json:"longitude"`
	Latitude        float64             `json:"latitude"`
	Email           string              `json:"email"`
	ContactNumber   string              `json:"contactNumber"`
	IsAdminVerified bool                `json:"isAdminVerified"`
	MediaID         *uint               `json:"mediaID"`
	Media           *MediaResource      `json:"media"`
	CompanyID       *uint               `json:"companyID"`
	Company         *CompanyResource    `json:"company"`
	Employees       []*EmployeeResource `json:"employees"`
	Members         []*MemberResource   `json:"members"`
}

func (m *ModelTransformer) BranchToResource(branch *Branch) *BranchResource {
	if branch == nil {
		return nil
	}

	return &BranchResource{
		ID:              branch.ID,
		CreatedAt:       branch.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       branch.UpdatedAt.Format(time.RFC3339),
		Name:            branch.Name,
		Address:         branch.Address,
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
		Members:         m.MemberToResourceList(branch.Members),
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
