package resources

import "horizon-core/internal/models"

type PermissionResource struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	Read              bool   `json:"read"`
	ReadDescription   string `json:"read_description"`
	Update            bool   `json:"update"`
	UpdateDescription string `json:"update_description"`
	Create            bool   `json:"create"`
	CreateDescription string `json:"create_description"`
}

func NewPermissionResource(permission models.Permission) PermissionResource {
	return PermissionResource{
		ID:                permission.ID,
		Name:              permission.Name,
		Description:       permission.Description,
		Read:              permission.Read,
		ReadDescription:   permission.ReadDescription,
		Update:            permission.Update,
		UpdateDescription: permission.UpdateDescription,
		Create:            permission.Create,
		CreateDescription: permission.CreateDescription,
	}
}
