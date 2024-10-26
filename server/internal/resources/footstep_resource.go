package resources

import (
	"horizon/server/internal/models"
	"time"
)

// FootstepResource represents the JSON structure for a Footstep model.
type FootstepResource struct {
	ID          uint   `json:"id"`
	Description string `json:"description,omitempty"`
	Activity    string `json:"activity"`
	AccountType string `json:"accountType"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`

	Admin    *AdminResource    `json:"admin,omitempty"`
	Employee *EmployeeResource `json:"employee,omitempty"`
	Owner    *OwnerResource    `json:"owner,omitempty"`
	Member   *MemberResource   `json:"member,omitempty"`
}

// ToResourceFootstep converts a Footstep model to a FootstepResource.
func ToResourceFootstep(footstep models.Footstep) FootstepResource {
	resource := FootstepResource{
		ID:          footstep.ID,
		Description: footstep.Description,
		Activity:    footstep.Activity,
		AccountType: footstep.AccountType,
		CreatedAt:   footstep.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   footstep.UpdatedAt.Format(time.RFC3339),
	}

	switch footstep.AccountType {
	case "Admin":
		if footstep.Admin != nil {
			adminResource := ToResourceAdmin(*footstep.Admin)
			resource.Admin = &adminResource
		}
	case "Employee":
		if footstep.Employee != nil {
			employeeResource := ToResourceEmployee(*footstep.Employee)
			resource.Employee = &employeeResource
		}
	case "Owner":
		if footstep.Owner != nil {
			ownerResource := ToResourceOwner(*footstep.Owner)
			resource.Owner = &ownerResource
		}
	case "Member":
		if footstep.Member != nil {
			memberResource := ToResourceMember(*footstep.Member)
			resource.Member = &memberResource
		}
	}

	return resource
}

// ToResourceListFootsteps converts a slice of Footstep models to a slice of FootstepResource.
func ToResourceListFootsteps(footsteps []models.Footstep) []FootstepResource {
	resources := make([]FootstepResource, len(footsteps))
	for i, footstep := range footsteps {
		resources[i] = ToResourceFootstep(footstep)
	}
	return resources
}
