package resources

import (
	"horizon/server/internal/models"
	"time"
)

type RoleResource struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ApiKey      string `json:"api_key"`
	Color       string `json:"color"`

	// Permissions
	ReadRole   bool `json:"read_role"`
	WriteRole  bool `json:"write_role"`
	UpdateRole bool `json:"update_role"`
	DeleteRole bool `json:"delete_role"`

	ReadErrorDetails   bool `json:"read_error_details"`
	WriteErrorDetails  bool `json:"write_error_details"`
	UpdateErrorDetails bool `json:"update_error_details"`
	DeleteErrorDetails bool `json:"delete_error_details"`

	ReadGender   bool `json:"read_gender"`
	WriteGender  bool `json:"write_gender"`
	UpdateGender bool `json:"update_gender"`
	DeleteGender bool `json:"delete_gender"`

	Admins    []AdminResource    `json:"admins,omitempty"`
	Employees []EmployeeResource `json:"employees,omitempty"`
	Members   []MemberResource   `json:"members,omitempty"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToResourceRole(role models.Role) RoleResource {

	return RoleResource{
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

func ToResourceListRoles(roles []models.Role) []RoleResource {
	resourceList := make([]RoleResource, len(roles))
	for i, role := range roles {
		resourceList[i] = ToResourceRole(role)
	}
	return resourceList
}
