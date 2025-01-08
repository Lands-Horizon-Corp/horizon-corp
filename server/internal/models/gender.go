package models

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
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

type GenderRepository struct {
	*Repository[Gender]
}

func NewGenderRepository(db *providers.DatabaseService) *GenderRepository {
	return &GenderRepository{
		Repository: NewRepository[Gender](db),
	}
}
