package resources

import (
	"horizon/server/internal/models"
	"time"
)

type BranchResource struct {
	ID              uint               `json:"id"`
	Name            string             `json:"name"`
	Address         string             `json:"address,omitempty"`
	Longitude       float64            `json:"longitude,omitempty"`
	Latitude        float64            `json:"latitude,omitempty"`
	Email           string             `json:"email"`
	ContactNumber   string             `json:"contactNumber"`
	IsAdminVerified bool               `json:"isAdminVerified"`
	Media           *MediaResource     `json:"media,omitempty"`
	Company         CompanyResource    `json:"company,omitempty"`
	Employees       []EmployeeResource `json:"employees,omitempty"`
	Members         []MemberResource   `json:"members,omitempty"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func ToResourceBranch(branch models.Branch) BranchResource {
	var mediaResource *MediaResource
	if branch.Media != nil {
		mediaRes := ToResourceMedia(*branch.Media)
		mediaResource = &mediaRes
	}

	return BranchResource{
		ID:              branch.ID,
		Name:            branch.Name,
		Address:         branch.Address,
		Longitude:       branch.Longitude,
		Latitude:        branch.Latitude,
		Email:           branch.Email,
		ContactNumber:   branch.ContactNumber,
		IsAdminVerified: branch.IsAdminVerified,
		Media:           mediaResource,
		Company:         ToResourceCompany(branch.Company),
		Employees:       ToResourceListEmployees(branch.Employees),
		Members:         ToResourceListMembers(branch.Members),

		CreatedAt: branch.CreatedAt.Format(time.RFC3339),
		UpdatedAt: branch.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResourceListBranch(branches []models.Branch) []BranchResource {
	var resources []BranchResource
	for _, branch := range branches {
		resources = append(resources, ToResourceBranch(branch))
	}
	return resources
}
