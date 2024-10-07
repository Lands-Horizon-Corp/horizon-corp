package resources

import "horizon/server/internal/models"

// ContactResource represents the structure of the contact resource for API responses.
type ContactResource struct {
	ID          uint   `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// ToResource converts a contact model to a contact resource.
func ToResourceContacts(contact models.Contacts) ContactResource {
	return ContactResource{
		ID:          contact.ID,
		FirstName:   contact.FirstName,
		LastName:    contact.LastName,
		Email:       contact.Email,
		Description: contact.Description,
		CreatedAt:   contact.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   contact.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

// ToResourceList converts a slice of contact models to a slice of contact resources.
func ToResourceListContacts(contacts []models.Contacts) []ContactResource {
	var resources []ContactResource
	for _, contact := range contacts {
		resources = append(resources, ToResourceContacts(contact))
	}
	return resources
}
