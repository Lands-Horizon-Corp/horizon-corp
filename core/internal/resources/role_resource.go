package resources

import "horizon-core/internal/models"

type RoleResource struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Permissions []PermissionResource `json:"permissions,omitempty"`
}

func NewRoleResource(role models.Role) RoleResource {
	permissions := make([]PermissionResource, len(role.Permissions))
	for i, perm := range role.Permissions {
		permissions[i] = NewPermissionResource(perm)
	}

	return RoleResource{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		Permissions: permissions,
	}
}
