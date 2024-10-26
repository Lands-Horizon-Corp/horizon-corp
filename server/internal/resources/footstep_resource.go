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

// ToResourceFootstep converts a Footstep model to a *FootstepResource.
func ToResourceFootstep(footstep *models.Footstep) *FootstepResource {
	if footstep == nil {
		return nil
	}

	resource := &FootstepResource{
		ID:          footstep.ID,
		Description: footstep.Description,
		Activity:    footstep.Activity,
		AccountType: footstep.AccountType,
		CreatedAt:   footstep.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   footstep.UpdatedAt.Format(time.RFC3339),
	}

	// Populate the relevant field based on AccountType
	switch footstep.AccountType {
	case "Admin":
		if footstep.Admin != nil {
			resource.Admin = ToResourceAdmin(footstep.Admin)
		}
	case "Employee":
		if footstep.Employee != nil {
			resource.Employee = ToResourceEmployee(footstep.Employee)
		}
	case "Owner":
		if footstep.Owner != nil {
			resource.Owner = ToResourceOwner(footstep.Owner)
		}
	case "Member":
		if footstep.Member != nil {
			resource.Member = ToResourceMember(footstep.Member)
		}
	}

	return resource
}

// ToResourceListFootsteps converts a slice of Footstep models to a slice of *FootstepResource.
func ToResourceListFootsteps(footsteps []*models.Footstep) []*FootstepResource {
	resources := make([]*FootstepResource, len(footsteps))
	for i, footstep := range footsteps {
		resources[i] = ToResourceFootstep(footstep)
	}
	return resources
}
