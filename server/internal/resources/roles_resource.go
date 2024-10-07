package resources

import (
	"horizon/server/internal/models"
)

// RolesResource represents the structure of the roles resource for API responses.
type RolesResource struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Color       string `json:"color"`

	ApiKey    string `json:"api_key,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`

	ReadRole           bool `json:"read_role,omitempty"`
	WriteRole          bool `json:"write_role,omitempty"`
	UpdateRole         bool `json:"update_role,omitempty"`
	DeleteRole         bool `json:"delete_role,omitempty"`
	ReadErrorDetails   bool `json:"read_error_details,omitempty"`
	WriteErrorDetails  bool `json:"write_error_details,omitempty"`
	UpdateErrorDetails bool `json:"update_error_details,omitempty"`
	DeleteErrorDetails bool `json:"delete_error_details,omitempty"`
	ReadGender         bool `json:"read_gender,omitempty"`
	WriteGender        bool `json:"write_gender,omitempty"`
	UpdateGender       bool `json:"update_gender,omitempty"`
	DeleteGender       bool `json:"delete_gender,omitempty"`
}

// ToResource converts a Roles model to a RolesResource.
func ToResourceRoles(roles models.Roles) RolesResource {
	return RolesResource{
		ID:          roles.ID,
		Name:        roles.Name,
		Description: roles.Description,
		Color:       roles.Color,

		ApiKey:             roles.ApiKey,
		ReadRole:           roles.ReadRole,
		WriteRole:          roles.WriteRole,
		UpdateRole:         roles.UpdateRole,
		DeleteRole:         roles.DeleteRole,
		ReadErrorDetails:   roles.ReadErrorDetails,
		WriteErrorDetails:  roles.WriteErrorDetails,
		UpdateErrorDetails: roles.UpdateErrorDetails,
		DeleteErrorDetails: roles.DeleteErrorDetails,
		ReadGender:         roles.ReadGender,
		WriteGender:        roles.WriteGender,
		UpdateGender:       roles.UpdateGender,
		DeleteGender:       roles.DeleteGender,
		CreatedAt:          roles.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:          roles.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

// ToResourceList converts a slice of Roles models to a slice of RolesResource.
func ToResourceListRoles(rolesList []models.Roles) []RolesResource {
	var resources []RolesResource
	for _, roles := range rolesList {
		resources = append(resources, ToResourceRoles(roles))
	}
	return resources
}
