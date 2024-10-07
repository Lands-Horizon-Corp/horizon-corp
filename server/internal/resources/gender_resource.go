package resources

import (
	"horizon/server/internal/models"
	"time"
)

type GenderResource struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func ToResourceGender(gender models.Gender) GenderResource {
	return GenderResource{
		ID:          gender.ID,
		Name:        gender.Name,
		Description: gender.Description,
		CreatedAt:   gender.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   gender.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResourceListGender(genders []models.Gender) []GenderResource {
	var resources []GenderResource
	for _, gender := range genders {
		resources = append(resources, ToResourceGender(gender))
	}
	return resources
}
