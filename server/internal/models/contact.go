package models

import (
	"time"

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
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	ContactNumber string `json:"contactNumber"`
	Description   string `json:"description"`
}

func (m *ModelTransformer) ContactToResource(contact *Contact) *ContactResource {
	if contact == nil {
		return nil
	}

	return &ContactResource{
		ID:            contact.ID,
		CreatedAt:     contact.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     contact.UpdatedAt.Format(time.RFC3339),
		FirstName:     contact.FirstName,
		LastName:      contact.LastName,
		Email:         contact.Email,
		ContactNumber: contact.ContactNumber,
		Description:   contact.Description,
	}
}

func (m *ModelTransformer) ContactToResourceList(contactList []*Contact) []*ContactResource {
	if contactList == nil {
		return nil
	}

	var contactResources []*ContactResource
	for _, contact := range contactList {
		contactResources = append(contactResources, m.ContactToResource(contact))
	}
	return contactResources
}
