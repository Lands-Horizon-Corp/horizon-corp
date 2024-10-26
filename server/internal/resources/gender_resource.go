package resources

import (
	"horizon/server/internal/models"
	"time"
)

// GenderResource struct for API response
type GenderResource struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ToResourceGender converts a Gender model to a GenderResource
func ToResourceGender(gender models.Gender) GenderResource {
	return GenderResource{
		ID:          gender.ID,
		Name:        gender.Name,
		Description: gender.Description,

		CreatedAt: gender.CreatedAt.Format(time.RFC3339),
		UpdatedAt: gender.UpdatedAt.Format(time.RFC3339),
	}
}

// ToResourceListGender converts a slice of Gender models to a slice of GenderResource
func ToResourceListGender(genderList []models.Gender) []GenderResource {
	resourceList := make([]GenderResource, len(genderList))
	for i, gender := range genderList {
		resourceList[i] = ToResourceGender(gender)
	}
	return resourceList
}
