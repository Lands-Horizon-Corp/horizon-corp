// package service

package service

import (
	"horizon-core/internal/events"
	"horizon-core/internal/models"
	"horizon-core/internal/repository"
	"horizon-core/internal/resources"

	"gorm.io/gorm"
)

type PermissionService struct {
	*repository.ModelRepository[models.Permission]
}

func NewPermissionService(db *gorm.DB) *PermissionService {
	return &PermissionService{
		ModelRepository: repository.NewModelRepository[models.Permission](db),
	}
}

func (r *PermissionService) GetPermissionByID(id string) (resources.PermissionResource, error) {
	eagerLoads := []string{
		"Role", // Preload the Role relationship if needed
	}

	permission, err := r.FindByID(id, eagerLoads)
	if err != nil {
		return resources.PermissionResource{}, err
	}

	permissionResource := resources.NewPermissionResource(permission)
	return permissionResource, nil
}

func (r *PermissionService) ListPermissions(req repository.ListRequest) (repository.ListResponse[resources.PermissionResource], error) {
	eagerLoads := []string{
		"Role", // Preload the Role relationship if needed
	}

	listResponse, err := r.List(req, eagerLoads)
	if err != nil {
		return repository.ListResponse[resources.PermissionResource]{}, err
	}

	permissionResources := make([]resources.PermissionResource, len(listResponse.Data))
	for i, permission := range listResponse.Data {
		permissionResources[i] = resources.NewPermissionResource(permission)
	}

	response := repository.ListResponse[resources.PermissionResource]{
		Data:       permissionResources,
		Pagination: listResponse.Pagination,
	}

	return response, nil
}

func (r *PermissionService) CreatePermission(permissionInput models.Permission) (resources.PermissionResource, error) {
	if err := r.Create(&permissionInput); err != nil {
		return resources.PermissionResource{}, err
	}

	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.PermissionCreated,
		Payload: permissionInput,
	})

	return r.GetPermissionByID(permissionInput.ID)
}

func (r *PermissionService) UpdatePermission(permissionInput models.Permission) (resources.PermissionResource, error) {
	if err := r.Update(&permissionInput); err != nil {
		return resources.PermissionResource{}, err
	}

	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.PermissionUpdated,
		Payload: permissionInput,
	})

	return r.GetPermissionByID(permissionInput.ID)
}

func (r *PermissionService) DeletePermission(id string) error {
	permission, err := r.FindByID(id, nil)
	if err != nil {
		return err
	}

	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.PermissionDeleted,
		Payload: permission,
	})

	return r.Delete(&permission)
}
