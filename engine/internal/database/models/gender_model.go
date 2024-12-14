package models

import (
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model

	// Fields
	Name        string `gorm:"type:varchar(255);unique;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`

	// Relationship 0 to many
	Employees []*Employee `gorm:"foreignKey:GenderID" json:"employees"`
	Members   []*Member   `gorm:"foreignKey:GenderID" json:"members"`
	Owners    []*Owner    `gorm:"foreignKey:GenderID" json:"owners"`
	Admins    []*Admin    `gorm:"foreignKey:GenderID" json:"admins"`
}

type GenderResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name        string              `json:"name"`
	Description string              `json:"description"`
	Employees   []*EmployeeResource `json:"employees"`
	Members     []*MemberResource   `json:"members"`
	Owners      []*OwnerResource    `json:"owners"`
	Admins      []*AdminResource    `json:"admins"`
}

type GenderRequest struct {
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description,omitempty" validate:"max=1000"`
}

func (m *ModelResource) GenderToResource(gender *Gender) *GenderResource {
	if gender == nil {
		return nil
	}
	return &GenderResource{
		ID:        gender.ID,
		CreatedAt: gender.CreatedAt.Format(time.RFC3339),
		UpdatedAt: gender.UpdatedAt.Format(time.RFC3339),

		Name:        gender.Name,
		Description: gender.Description,
		Employees:   m.EmployeeToResourceList(gender.Employees),
		Members:     m.MemberToResourceList(gender.Members),
		Owners:      m.OwnerToResourceList(gender.Owners),
		Admins:      m.AdminToResourceList(gender.Admins),
	}
}

func (m *ModelResource) GenderToResourceList(genders []*Gender) []*GenderResource {
	if genders == nil {
		return nil
	}

	var genderResources []*GenderResource
	for _, gender := range genders {
		genderResources = append(genderResources, m.GenderToResource(gender))
	}
	return genderResources
}

func (m *ModelResource) ValidateGenderRequest(req *GenderRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return m.helpers.FormatValidationError(err)
	}
	return nil
}

func (m *ModelResource) GenderSeeders() error {
	m.logger.Info("Seeding Gender")
	return nil
}
