package models

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
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

type ContactResource struct {
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	ContactNumber string `json:"contactNumber"`
	Description   string `json:"description"`
}

type ContactModel struct {
	lc     *fx.Lifecycle
	db     *gorm.DB
	logger *zap.Logger
}

func NewContactModel(
	lc *fx.Lifecycle,
	db *gorm.DB,
	logger *zap.Logger,
) *ContactModel {
	return &ContactModel{
		lc:     lc,
		db:     db,
		logger: logger,
	}
}

func (cm *ContactModel) SeedDatabase() {
}

func (cm *ContactModel) ToResource(contact *Contact) *ContactResource {
	if contact == nil {
		return nil
	}
	return &ContactResource{
		FirstName:     contact.FirstName,
		LastName:      contact.LastName,
		Email:         contact.Email,
		ContactNumber: contact.ContactNumber,
		Description:   contact.Description,
	}
}

func (cm *ContactModel) ToResourceList(contacts []*Contact) []*ContactResource {
	if contacts == nil {
		return nil
	}
	var contactResources []*ContactResource
	for _, contact := range contacts {
		contactResources = append(contactResources, cm.ToResource(contact))
	}
	return contactResources
}
