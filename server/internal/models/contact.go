package models

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model

	// Fields
	FirstName     string `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName      string `gorm:"type:varchar(255);not null" json:"last_name"`
	Email         string `gorm:"type:varchar(255);unique;not null" json:"email"`
	ContactNumber string `gorm:"type:varchar(15);not null" json:"contact_number"`
	Description   string `gorm:"type:text" json:"description"`
}

type ContactRepository struct {
	*managers.Repository[Contact]
}

func NewContactRepository(db *gorm.DB) *ContactRepository {
	return &ContactRepository{
		Repository: managers.NewRepository[Contact](db),
	}
}
