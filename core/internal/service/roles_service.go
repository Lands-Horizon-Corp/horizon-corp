// package service

package service

import (
	"horizon-core/internal/events"
	"horizon-core/internal/models"
	"horizon-core/internal/repository"
	"horizon-core/internal/resources"

	"gorm.io/gorm"
)

type RoleService struct {
	*repository.ModelRepository[models.Role]
}

func NewRoleService(db *gorm.DB) *RoleService {
	return &RoleService{
		ModelRepository: repository.NewModelRepository[models.Role](db),
	}
}

func (r *RoleService) GetRoleByID(id string) (resources.RoleResource, error) {
	eagerLoads := []string{
		"Permissions",
	}

	role, err := r.FindByID(id, eagerLoads)
	if err != nil {
		return resources.RoleResource{}, err
	}

	roleResource := resources.NewRoleResource(role)
	return roleResource, nil
}

func (r *RoleService) ListRoles(req repository.ListRequest) (repository.ListResponse[resources.RoleResource], error) {
	eagerLoads := []string{
		"Permissions",
	}

	listResponse, err := r.List(req, eagerLoads)
	if err != nil {
		return repository.ListResponse[resources.RoleResource]{}, err
	}

	roleResources := make([]resources.RoleResource, len(listResponse.Data))
	for i, role := range listResponse.Data {
		roleResources[i] = resources.NewRoleResource(role)
	}

	response := repository.ListResponse[resources.RoleResource]{
		Data:       roleResources,
		Pagination: listResponse.Pagination,
	}

	return response, nil
}

func (r *RoleService) CreateRole(roleInput models.Role) (resources.RoleResource, error) {
	if err := r.Create(&roleInput); err != nil {
		return resources.RoleResource{}, err
	}
	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.RoleCreated,
		Payload: roleInput,
	})
	return r.GetRoleByID(roleInput.ID)
}

func (r *RoleService) UpdateRole(roleInput models.Role) (resources.RoleResource, error) {
	if err := r.Update(&roleInput); err != nil {
		return resources.RoleResource{}, err
	}
	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.RoleUpdated,
		Payload: roleInput,
	})
	return r.GetRoleByID(roleInput.ID)
}

func (r *RoleService) DeleteRole(id string) error {
	role, err := r.FindByID(id, nil)
	if err != nil {
		return err
	}
	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.RoleDeleted,
		Payload: role,
	})
	return r.Delete(&role)
}
