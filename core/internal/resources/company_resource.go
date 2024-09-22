package resources

import "horizon-core/internal/models"

type CompanyResource struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Owner       *OwnerResource   `json:"owner,omitempty"`
	Branches    []BranchResource `json:"branches,omitempty"`
}

func NewCompanyResource(company models.Company) CompanyResource {
	owner := NewOwnerResource(company.Owner)

	branches := make([]BranchResource, len(company.Branches))
	for i, branch := range company.Branches {
		branches[i] = NewBranchResource(branch)
	}

	return CompanyResource{
		ID:          company.ID,
		Name:        company.Name,
		Description: company.Description,
		Owner:       &owner,
		Branches:    branches,
	}
}
