// package service

package service

import (
	"horizon-core/internal/events"
	"horizon-core/internal/models"
	"horizon-core/internal/repository"
	"horizon-core/internal/resources"

	"gorm.io/gorm"
)

type OwnerService struct {
	*repository.ModelRepository[models.Owner]
}

func NewOwnerService(db *gorm.DB) *OwnerService {
	return &OwnerService{
		ModelRepository: repository.NewModelRepository[models.Owner](db),
	}
}

func (r *OwnerService) GetOwnerByID(id string) (resources.OwnerResource, error) {
	eagerLoads := []string{
		"ProfilePicture",
		"Roles.Permissions",
	}

	owner, err := r.FindByID(id, eagerLoads)
	if err != nil {
		return resources.OwnerResource{}, err
	}

	ownerResource := resources.NewOwnerResource(owner)
	return ownerResource, nil
}

func (r *OwnerService) ListOwners(req repository.ListRequest) (repository.ListResponse[resources.OwnerResource], error) {
	eagerLoads := []string{
		"ProfilePicture",
		"Roles.Permissions",
	}

	listResponse, err := r.List(req, eagerLoads)
	if err != nil {
		return repository.ListResponse[resources.OwnerResource]{}, err
	}

	ownerResources := make([]resources.OwnerResource, len(listResponse.Data))
	for i, owner := range listResponse.Data {
		ownerResources[i] = resources.NewOwnerResource(owner)
	}

	response := repository.ListResponse[resources.OwnerResource]{
		Data:       ownerResources,
		Pagination: listResponse.Pagination,
	}

	return response, nil
}

func (r *OwnerService) CreateOwner(ownerInput models.Owner) (resources.OwnerResource, error) {
	if err := r.Create(&ownerInput); err != nil {
		return resources.OwnerResource{}, err
	}
	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.OwnerCreated,
		Payload: ownerInput,
	})
	return r.GetOwnerByID(ownerInput.ID)
}

func (r *OwnerService) UpdateOwner(ownerInput models.Owner) (resources.OwnerResource, error) {
	if err := r.Update(&ownerInput); err != nil {
		return resources.OwnerResource{}, err
	}
	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.OwnerUpdated,
		Payload: ownerInput,
	})
	return r.GetOwnerByID(ownerInput.ID)
}

func (r *OwnerService) DeleteOwner(id string) error {
	owner, err := r.FindByID(id, nil)
	if err != nil {
		return err
	}
	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.OwnerDeleted,
		Payload: owner,
	})
	return r.Delete(&owner)
}
