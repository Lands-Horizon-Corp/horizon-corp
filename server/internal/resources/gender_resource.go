package resources

import (
	"horizon/server/internal/models"
	"time"
)

type GenderResource struct {
	ID          uint               `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Employees   []EmployeeResource `json:"employees,omitempty"`
	Members     []MemberResource   `json:"members,omitempty"`
	Owners      []OwnerResource    `json:"owners,omitempty"`
	Admins      []AdminResource    `json:"admins,omitempty"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func ToResourceGender(gender models.Gender) GenderResource {
	return GenderResource{
		ID:          gender.ID,
		Name:        gender.Name,
		Description: gender.Description,

		Employees: ToResourceListEmployees(gender.Employees),
		Members:   ToResourceListMembers(gender.Members),
		Owners:    ToResourceListOwners(gender.Owners),
		Admins:    ToResourceListAdmins(gender.Admins),

		CreatedAt: gender.CreatedAt.Format(time.RFC3339),
		UpdatedAt: gender.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResourceListGender(genderList []models.Gender) []GenderResource {
	resourceList := make([]GenderResource, len(genderList))
	for i, gender := range genderList {
		resourceList[i] = ToResourceGender(gender)
	}
	return resourceList
}
