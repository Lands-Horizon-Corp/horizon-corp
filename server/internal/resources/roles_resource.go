package resources

import (
	"horizon/server/internal/models"
)

type RoleResource struct {
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

func ToResourceRole(roles models.Role) RoleResource {
	return RoleResource{
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

func ToResourceListRole(rolesList []models.Role) []RoleResource {
	var resources []RoleResource
	for _, roles := range rolesList {
		resources = append(resources, ToResourceRole(roles))
	}
	return resources
}
