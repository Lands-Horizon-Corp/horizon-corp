package models

import (
	"github.com/go-playground/validator"
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

type ContactRequest struct {
	FirstName     string `json:"firstName" validate:"required,max=255"`
	LastName      string `json:"lastName" validate:"required,max=255"`
	Email         string `json:"email" validate:"required,email,max=255"`
	ContactNumber string `json:"contactNumber" validate:"required,max=15"`
	Description   string `json:"description,omitempty"`
}

func (m *ModelResource) ContactToResource(contact *Contact) *ContactResource {
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

func (m *ModelResource) ContactToResourceList(contacts []*Contact) []*ContactResource {
	if contacts == nil {
		return nil
	}
	var contactResources []*ContactResource
	for _, contact := range contacts {
		contactResources = append(contactResources, m.ContactToResource(contact))
	}
	return contactResources
}

func (m *ModelResource) ValidateContactRequest(req *ContactRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return m.helpers.FormatValidationError(err)
	}
	return nil
}

func (m *ModelResource) ContactSeeders() error {
	m.logger.Info("Seeding Contact")
	return nil
}
