package resources

import (
	"horizon/server/internal/models"
	"time"
)

type GenderResource struct {
	ID          uint                `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Employees   []*EmployeeResource `json:"employees,omitempty"` // Updated to slice of pointers
	Members     []*MemberResource   `json:"members,omitempty"`   // Updated to slice of pointers
	Owners      []*OwnerResource    `json:"owners,omitempty"`    // Updated to slice of pointers
	Admins      []*AdminResource    `json:"admins,omitempty"`    // Updated to slice of pointers

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Convert models.Gender to *GenderResource
func ToResourceGender(gender *models.Gender) *GenderResource {
	if gender == nil {
		return nil
	}

	return &GenderResource{
		ID:          gender.ID,
		Name:        gender.Name,
		Description: gender.Description,
		Employees:   ToResourceListEmployees(gender.Employees),
		Members:     ToResourceListMembers(gender.Members),
		Owners:      ToResourceListOwners(gender.Owners),
		Admins:      ToResourceListAdmins(gender.Admins),
		CreatedAt:   gender.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   gender.UpdatedAt.Format(time.RFC3339),
	}
}

// Convert []*models.Gender to []*GenderResource
func ToResourceListGender(genderList []*models.Gender) []*GenderResource {
	resourceList := make([]*GenderResource, len(genderList))
	for i, gender := range genderList {
		resourceList[i] = ToResourceGender(gender)
	}
	return resourceList
}
