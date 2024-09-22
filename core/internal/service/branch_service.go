// package service

package service

import (
	"horizon-core/internal/events"
	"horizon-core/internal/models"
	"horizon-core/internal/repository"
	"horizon-core/internal/resources"

	"gorm.io/gorm"
)

type BranchService struct {
	*repository.ModelRepository[models.Branch]
}

func NewBranchService(db *gorm.DB) *BranchService {
	return &BranchService{
		ModelRepository: repository.NewModelRepository[models.Branch](db),
	}
}

func (r *BranchService) GetBranchByID(id string) (resources.BranchResource, error) {
	eagerLoads := []string{
		"ProfilePicture",
		"Company.Owner.ProfilePicture",
	}

	branch, err := r.FindByID(id, eagerLoads)
	if err != nil {
		return resources.BranchResource{}, err
	}

	branchResource := resources.NewBranchResource(branch)
	return branchResource, nil
}

func (r *BranchService) ListBranches(req repository.ListRequest) (repository.ListResponse[resources.BranchResource], error) {
	eagerLoads := []string{
		"ProfilePicture",
		"Company.Owner.ProfilePicture",
	}

	listResponse, err := r.List(req, eagerLoads)
	if err != nil {
		return repository.ListResponse[resources.BranchResource]{}, err
	}

	branchResources := make([]resources.BranchResource, len(listResponse.Data))
	for i, branch := range listResponse.Data {
		branchResources[i] = resources.NewBranchResource(branch)
	}

	response := repository.ListResponse[resources.BranchResource]{
		Data:       branchResources,
		Pagination: listResponse.Pagination,
	}

	return response, nil
}

func (r *BranchService) CreateBranch(branchInput models.Branch) (resources.BranchResource, error) {
	if err := r.Create(&branchInput); err != nil {
		return resources.BranchResource{}, err
	}

	return r.GetBranchByID(branchInput.ID)
}

func (r *BranchService) UpdateBranch(branchInput models.Branch) (resources.BranchResource, error) {
	if err := r.Update(&branchInput); err != nil {
		return resources.BranchResource{}, err
	}

	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.BranchUpdated,
		Payload: branchInput,
	})

	return r.GetBranchByID(branchInput.ID)
}

func (r *BranchService) DeleteBranch(id string) error {
	branch, err := r.FindByID(id, nil)
	if err != nil {
		return err
	}

	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.BranchDeleted,
		Payload: branch,
	})
	return r.Delete(&branch)
}
