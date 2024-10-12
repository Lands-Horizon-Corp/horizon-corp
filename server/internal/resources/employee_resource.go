package resources

import (
	"horizon/server/internal/models"
	"time"
)

type EmployeeResource struct {
	ID                uint                  `json:"id"`
	FirstName         string                `json:"firstName"`
	LastName          string                `json:"lastName"`
	MiddleName        string                `json:"middleName"`
	PermanentAddress  string                `json:"permanentAddress"`
	Description       string                `json:"description"`
	Birthdate         time.Time             `json:"birthdate"`
	Username          string                `json:"username"`
	Email             string                `json:"email"`
	IsEmailVerified   bool                  `json:"isEmailVerified"`
	IsContactVerified bool                  `json:"isContactVerified"`
	ContactNumber     string                `json:"contactNumber"`
	Media             *MediaResource        `json:"media,omitempty"`
	Status            models.EmployeeStatus `json:"status"`
	CreatedAt         string                `json:"createdAt"`
	UpdatedAt         string                `json:"updatedAt"`
}

func ToResourceEmployee(employee models.Employee) EmployeeResource {
	var mediaResource *MediaResource

	if employee.MediaID != nil {
		media := ToResourceMedia(employee.Media)
		mediaResource = &media
	}

	return EmployeeResource{
		ID:                employee.ID,
		FirstName:         employee.FirstName,
		LastName:          employee.LastName,
		MiddleName:        employee.MiddleName,
		PermanentAddress:  employee.PermanentAddress,
		Description:       employee.Description,
		Birthdate:         employee.Birthdate,
		Username:          employee.Username,
		Email:             employee.Email,
		IsEmailVerified:   employee.IsEmailVerified,
		IsContactVerified: employee.IsContactVerified,
		ContactNumber:     employee.ContactNumber,
		Media:             mediaResource, // Set the media resource if exists
		Status:            employee.Status,
		CreatedAt:         employee.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         employee.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResourceListEmployees(employees []models.Employee) []EmployeeResource {
	var resources []EmployeeResource
	for _, employee := range employees {
		resources = append(resources, ToResourceEmployee(employee))
	}
	return resources
}
