package resources

import "horizon/server/internal/models"

// GenderResource represents the structure of the gender resource for API responses.
type GenderResource struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ToResource converts a gender model to a gender resource.
func ToResource(gender models.Gender) GenderResource {
	return GenderResource{
		ID:          gender.ID,
		Name:        gender.Name,
		Description: gender.Description,
	}
}

func ToResourceList(genders []models.Gender) []GenderResource {
	var resources []GenderResource
	for _, gender := range genders {
		resources = append(resources, ToResource(gender))
	}
	return resources
}
