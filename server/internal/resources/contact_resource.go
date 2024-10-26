package resources

import (
	"horizon/server/internal/models"
	"time"
)

type ContactResource struct {
	ID            uint   `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	ContactNumber string `json:"contact_number"`
	Description   string `json:"description,omitempty"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToResourceContact(contact models.Contact) ContactResource {
	return ContactResource{
		ID:            contact.ID,
		FirstName:     contact.FirstName,
		LastName:      contact.LastName,
		Email:         contact.Email,
		ContactNumber: contact.ContactNumber,
		Description:   contact.Description,

		CreatedAt: contact.CreatedAt.Format(time.RFC3339),
		UpdatedAt: contact.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResourceListContacts(contacts []models.Contact) []ContactResource {
	var resources []ContactResource
	for _, contact := range contacts {
		resources = append(resources, ToResourceContact(contact))
	}
	return resources
}
