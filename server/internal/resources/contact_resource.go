package resources

import (
	"horizon/server/internal/models"
	"time"
)

type ContactResource struct {
	ID          uint   `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func ToResourceContact(contact models.Contact) ContactResource {
	return ContactResource{
		ID:          contact.ID,
		FirstName:   contact.FirstName,
		LastName:    contact.LastName,
		Email:       contact.Email,
		Description: contact.Description,
		CreatedAt:   contact.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   contact.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResourceListContacts(contacts []models.Contact) []ContactResource {
	var resources []ContactResource
	for _, contact := range contacts {
		resources = append(resources, ToResourceContact(contact))
	}
	return resources
}
