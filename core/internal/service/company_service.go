package service

import (
	"horizon-core/internal/models"
	"horizon-core/internal/repository"
	"horizon-core/internal/resources"

	"gorm.io/gorm"
)

type CompanyService struct {
	*repository.ModelRepository[models.Company]
}

func NewCompanyService(db *gorm.DB) *CompanyService {
	return &CompanyService{
		ModelRepository: repository.NewModelRepository[models.Company](db),
	}
}

func (r *CompanyService) GetCompanyByID(id string) (resources.CompanyResource, error) {
	eagerLoads := []string{
		"Owner.ProfilePicture",
		"Branches.ProfilePicture",
	}

	company, err := r.FindByID(id, eagerLoads)
	if err != nil {
		return resources.CompanyResource{}, err
	}

	companyResource := resources.NewCompanyResource(company)
	return companyResource, nil
}

func (r *CompanyService) ListCompanies(req repository.ListRequest) (repository.ListResponse[resources.CompanyResource], error) {
	eagerLoads := []string{
		"Owner.ProfilePicture",
		"Branches.ProfilePicture",
	}

	listResponse, err := r.List(req, eagerLoads)
	if err != nil {
		return repository.ListResponse[resources.CompanyResource]{}, err
	}

	companyResources := make([]resources.CompanyResource, len(listResponse.Data))
	for i, company := range listResponse.Data {
		companyResources[i] = resources.NewCompanyResource(company)
	}

	response := repository.ListResponse[resources.CompanyResource]{
		Data:       companyResources,
		Pagination: listResponse.Pagination,
	}

	return response, nil
}

func (r *CompanyService) CreateCompany(companyInput models.Company) (resources.CompanyResource, error) {
	if err := r.Create(&companyInput); err != nil {
		return resources.CompanyResource{}, err
	}

	return r.GetCompanyByID(companyInput.ID)
}

func (r *CompanyService) UpdateCompany(companyInput models.Company) (resources.CompanyResource, error) {
	if err := r.Update(&companyInput); err != nil {
		return resources.CompanyResource{}, err
	}

	return r.GetCompanyByID(companyInput.ID)
}

func (r *CompanyService) DeleteCompany(id string) error {
	company, err := r.FindByID(id, nil)
	if err != nil {
		return err
	}

	return r.Delete(&company)
}
