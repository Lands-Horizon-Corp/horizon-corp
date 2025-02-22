package models

import (
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model

	// Fields
	FirstName     string `gorm:"type:varchar(255);unsigned" json:"first_name"`
	LastName      string `gorm:"type:varchar(255);unsigned" json:"last_name"`
	Email         string `gorm:"type:varchar(255);unique;unsigned" json:"email"`
	ContactNumber string `gorm:"type:varchar(15);unsigned" json:"contact_number"`
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
		ID:        contact.ID,
		CreatedAt: contact.CreatedAt.Format(time.RFC3339),
		UpdatedAt: contact.UpdatedAt.Format(time.RFC3339),

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

// ContactToRecord converts a slice of Contact pointers into CSV records and headers.
func (m *ModelResource) ContactToRecord(contacts []*Contact) ([][]string, []string) {
	// Convert Contact structs to ContactResource structs
	resource := m.ContactToResourceList(contacts)
	records := make([][]string, 0, len(resource))

	for _, contact := range resource {
		id := strconv.Itoa(int(contact.ID))
		firstName := sanitizeCSVField(contact.FirstName)
		lastName := sanitizeCSVField(contact.LastName)
		email := sanitizeCSVField(contact.Email)
		contactNumber := sanitizeCSVField(contact.ContactNumber)
		description := sanitizeCSVField(contact.Description)
		createdAt := sanitizeCSVField(contact.CreatedAt)
		updatedAt := sanitizeCSVField(contact.UpdatedAt)

		record := []string{
			id,
			firstName,
			lastName,
			email,
			contactNumber,
			description,
			createdAt,
			updatedAt,
		}
		records = append(records, record)
	}

	headers := []string{
		"ID",
		"First Name",
		"Last Name",
		"Email",
		"Contact Number",
		"Description",
		"Created At",
		"Updated At",
	}

	return records, headers
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
