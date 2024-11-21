package resources

import (
	"horizon/server/internal/models"
	"time"
)

type CompanyResource struct {
	ID              uint              `json:"id"`
	Name            string            `json:"name"`
	Description     string            `json:"description,omitempty"`
	Address         string            `json:"address,omitempty"`
	Longitude       float64           `json:"longitude,omitempty"`
	Latitude        float64           `json:"latitude,omitempty"`
	ContactNumber   string            `json:"contactNumber"`
	IsAdminVerified bool              `json:"isAdminVerified"`
	Owner           *OwnerResource    `json:"owner,omitempty"`
	Media           *MediaResource    `json:"media,omitempty"`
	Branches        []*BranchResource `json:"branches,omitempty"` // Updated to slice of pointers

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Convert models.Company to *CompanyResource
func ToResourceCompany(company *models.Company) *CompanyResource {
	if company == nil {
		return nil
	}

	// Convert Owner
	var ownerResource *OwnerResource
	if company.Owner != nil {
		ownerResource = ToResourceOwner(company.Owner)
	}

	// Convert Media
	var mediaResource *MediaResource
	if company.Media != nil {
		mediaResource = ToResourceMedia(company.Media)
	}

	return &CompanyResource{
		ID:              company.ID,
		Name:            company.Name,
		Description:     company.Description,
		Address:         company.Address,
		Longitude:       company.Longitude,
		Latitude:        company.Latitude,
		ContactNumber:   company.ContactNumber,
		IsAdminVerified: company.IsAdminVerified,
		Owner:           ownerResource,
		Media:           mediaResource,
		Branches:        ToResourceListBranch(company.Branches),
		CreatedAt:       company.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       company.UpdatedAt.Format(time.RFC3339),
	}
}

// Convert []models.Company to []*CompanyResource
func ToResourceListCompanies(companies []*models.Company) []*CompanyResource {
	var resources []*CompanyResource
	for _, company := range companies {
		resources = append(resources, ToResourceCompany(company))
	}
	return resources
}
