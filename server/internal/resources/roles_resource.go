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

	ApiKey    string `json:"apiKey,omitempty"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	ReadRole           bool `json:"readRole,omitempty"`
	WriteRole          bool `json:"writeRole,omitempty"`
	UpdateRole         bool `json:"updateRole,omitempty"`
	DeleteRole         bool `json:"deleteRole,omitempty"`
	ReadErrorDetails   bool `json:"readErrorDetails,omitempty"`
	WriteErrorDetails  bool `json:"writeErrorDetails,omitempty"`
	UpdateErrorDetails bool `json:"updateErrorDetails,omitempty"`
	DeleteErrorDetails bool `json:"deleteErrorDetails,omitempty"`
	ReadGender         bool `json:"readGender,omitempty"`
	WriteGender        bool `json:"writeGender,omitempty"`
	UpdateGender       bool `json:"updateGender,omitempty"`
	DeleteGender       bool `json:"deleteGender,omitempty"`
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
