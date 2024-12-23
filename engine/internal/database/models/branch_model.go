package models

import (
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers/filter"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Branch struct {
	gorm.Model

	// Fields
	Name            string  `gorm:"type:varchar(255);not null" json:"name"`
	Address         string  `gorm:"type:varchar(500)" json:"address"`
	Longitude       float64 `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude        float64 `gorm:"type:decimal(10,7)" json:"latitude"`
	Email           string  `gorm:"type:varchar(255);not null" json:"email"`
	ContactNumber   string  `gorm:"type:varchar(15);not null" json:"contact_number"`
	IsAdminVerified bool    `gorm:"default:false" json:"is_admin_verified"`

	// Relationship 0 to 1
	MediaID *uint  `gorm:"type:bigint;unsigned" json:"media_id"`
	Media   *Media `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media"`

	// Relationship 0 to 1
	CompanyID *uint    `gorm:"type:bigint;unsigned;not null" json:"company_id"`
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

type BranchRequest struct {
	Name            string  `json:"name" validate:"required,max=255"`
	Address         string  `json:"address" validate:"max=500"`
	Longitude       float64 `json:"longitude" validate:"longitude"`
	Latitude        float64 `json:"latitude" validate:"latitude"`
	Email           string  `json:"email" validate:"required,email,max=255"`
	ContactNumber   string  `json:"contactNumber" validate:"required,max=15"`
	IsAdminVerified bool    `json:"isAdminVerified"`
	MediaID         *uint   `json:"mediaID,omitempty"`
	CompanyID       *uint   `json:"companyID" validate:"required"`
}

func (m *ModelResource) BranchToResource(branch *Branch) *BranchResource {
	if branch == nil {
		return nil
	}
	return &BranchResource{
		ID:        branch.ID,
		CreatedAt: branch.CreatedAt.Format(time.RFC3339),
		UpdatedAt: branch.UpdatedAt.Format(time.RFC3339),

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

func (m *ModelResource) BranchToResourceList(branch []*Branch) []*BranchResource {
	if branch == nil {
		return nil
	}
	var branchResources []*BranchResource
	for _, branch := range branch {
		branchResources = append(branchResources, m.BranchToResource(branch))
	}
	return branchResources
}

func (m *ModelResource) ValidateBranchRequest(req *BranchRequest) error {
	validate := validator.New()
	validate.RegisterValidation("longitude",
		func(fl validator.FieldLevel) bool {
			lon := fl.Field().Float()
			return lon >= -180 && lon <= 180
		})
	validate.RegisterValidation("latitude", func(fl validator.FieldLevel) bool {
		lat := fl.Field().Float()
		return lat >= -90 && lat <= 90
	})
	err := validate.Struct(req)
	if err != nil {
		return m.helpers.FormatValidationError(err)
	}
	return nil
}

func (m *ModelResource) BranchFilterForAdmin(filters string) (filter.FilterPages[Branch], error) {
	db := m.db.Client
	return m.BranchDB.GetPaginatedResult(db, filters)
}

func (m *ModelResource) BranchFilterForOwner(filters string, ownerId uint) (filter.FilterPages[Branch], error) {
	db := m.db.Client.Where("owner_id = ?", ownerId)
	return m.BranchDB.GetPaginatedResult(db, filters)
}

func (m *ModelResource) BranchSeeders() error {
	m.logger.Info("Seeding Branch")
	return nil
}
