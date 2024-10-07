package resources

import "horizon/server/internal/models"

// GenderResource represents the structure of the gender resource for API responses.
type GenderResource struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// ToResource converts a gender model to a gender resource.
func ToResourceGender(gender models.Gender) GenderResource {
	return GenderResource{
		ID:          gender.ID,
		Name:        gender.Name,
		Description: gender.Description,
		CreatedAt:   gender.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   gender.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToResourceListGender(genders []models.Gender) []GenderResource {
	var resources []GenderResource
	for _, gender := range genders {
		resources = append(resources, ToResourceGender(gender))
	}
	return resources
}
