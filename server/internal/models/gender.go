package models

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
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

type GenderRepository struct {
	*managers.Repository[Gender]
}

func NewGenderRepository(db *gorm.DB) *GenderRepository {
	return &GenderRepository{
		Repository: managers.NewRepository[Gender](db),
	}
}
