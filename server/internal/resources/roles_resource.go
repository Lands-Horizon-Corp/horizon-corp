package resources

import (
	"horizon/server/internal/models"
	"time"
)

type RoleResource struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ApiKey      string `json:"apiKey"`
	Color       string `json:"color"`

	// Permissions
	ReadRole   bool `json:"readRole"`
	WriteRole  bool `json:"writeRole"`
	UpdateRole bool `json:"updateRole"`
	DeleteRole bool `json:"deleteRole"`

	ReadErrorDetails   bool `json:"readErrorDetails"`
	WriteErrorDetails  bool `json:"writeErrorDetails"`
	UpdateErrorDetails bool `json:"updateErrorDetails"`
	DeleteErrorDetails bool `json:"deleteErrorDetails"`

	ReadGender   bool `json:"readGender"`
	WriteGender  bool `json:"writeGender"`
	UpdateGender bool `json:"updateGender"`
	DeleteGender bool `json:"deleteGender"`

	Admins    []*AdminResource    `json:"admins,omitempty"`    // Updated to slice of pointers
	Employees []*EmployeeResource `json:"employees,omitempty"` // Updated to slice of pointers
	Members   []*MemberResource   `json:"members,omitempty"`   // Updated to slice of pointers

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Convert models.Role to *RoleResource
func ToResourceRole(role *models.Role) *RoleResource {
	if role == nil {
		return nil
	}

	return &RoleResource{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		ApiKey:      role.ApiKey,
		Color:       role.Color,

		ReadRole:   role.ReadRole,
		WriteRole:  role.WriteRole,
		UpdateRole: role.UpdateRole,
		DeleteRole: role.DeleteRole,

		ReadErrorDetails:   role.ReadErrorDetails,
		WriteErrorDetails:  role.WriteErrorDetails,
		UpdateErrorDetails: role.UpdateErrorDetails,
		DeleteErrorDetails: role.DeleteErrorDetails,

		ReadGender:   role.ReadGender,
		WriteGender:  role.WriteGender,
		UpdateGender: role.UpdateGender,
		DeleteGender: role.DeleteGender,

		Admins:    ToResourceListAdmins(role.Admins),
		Employees: ToResourceListEmployees(role.Employees),
		Members:   ToResourceListMembers(role.Members),

		CreatedAt: role.CreatedAt.Format(time.RFC3339),
		UpdatedAt: role.UpdatedAt.Format(time.RFC3339),
	}
}

// Convert []*models.Role to []*RoleResource
func ToResourceListRoles(roles []*models.Role) []*RoleResource {
	resourceList := make([]*RoleResource, len(roles))
	for i, role := range roles {
		resourceList[i] = ToResourceRole(role)
	}
	return resourceList
}
