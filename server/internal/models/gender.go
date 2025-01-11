package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Gender struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Fields
	Name        string `gorm:"type:varchar(255);unique;unsigned" json:"name"`
	Description string `gorm:"type:text" json:"description"`

	// Relationship 0 to many
	Employees []*Employee `gorm:"foreignKey:GenderID" json:"employees"`
	Members   []*Member   `gorm:"foreignKey:GenderID" json:"members"`
	Owners    []*Owner    `gorm:"foreignKey:GenderID" json:"owners"`
	Admins    []*Admin    `gorm:"foreignKey:GenderID" json:"admins"`
}

type GenderResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	Name        string              `json:"name"`
	Description string              `json:"description"`
	Employees   []*EmployeeResource `json:"employees"`
	Members     []*MemberResource   `json:"members"`
	Owners      []*OwnerResource    `json:"owners"`
	Admins      []*AdminResource    `json:"admins"`
}

func (m *ModelTransformer) GenderToResource(gender *Gender) *GenderResource {
	if gender == nil {
		return nil
	}

	return &GenderResource{

		ID:        gender.ID,
		CreatedAt: gender.CreatedAt.Format(time.RFC3339),
		UpdatedAt: gender.UpdatedAt.Format(time.RFC3339),
		DeletedAt: gender.DeletedAt.Time.Format(time.RFC3339),

		Name:        gender.Name,
		Description: gender.Description,
		Employees:   m.EmployeeToResourceList(gender.Employees),
		Members:     m.MemberToResourceList(gender.Members),
		Owners:      m.OwnerToResourceList(gender.Owners),
		Admins:      m.AdminToResourceList(gender.Admins),
	}
}

func (m *ModelTransformer) GenderToResourceList(genderList []*Gender) []*GenderResource {
	if genderList == nil {
		return nil
	}

	var genderResources []*GenderResource
	for _, gender := range genderList {
		genderResources = append(genderResources, m.GenderToResource(gender))
	}
	return genderResources
}
