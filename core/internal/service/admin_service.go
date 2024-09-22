// package service

package service

import (
	"horizon-core/internal/events"
	"horizon-core/internal/models"
	"horizon-core/internal/repository"
	"horizon-core/internal/resources"

	"gorm.io/gorm"
)

type AdminService struct {
	*repository.ModelRepository[models.Admin]
}

func NewAdminService(db *gorm.DB) *AdminService {
	return &AdminService{
		ModelRepository: repository.NewModelRepository[models.Admin](db),
	}
}

func (r *AdminService) GetAdminByID(id string) (resources.AdminResource, error) {
	eagerLoads := []string{
		"ProfilePicture",
		"Roles.Permissions",
	}

	admin, err := r.FindByID(id, eagerLoads)
	if err != nil {
		return resources.AdminResource{}, err
	}

	adminResource := resources.NewAdminResource(admin)
	return adminResource, nil
}

func (r *AdminService) ListAdmins(req repository.ListRequest) (repository.ListResponse[resources.AdminResource], error) {

	eagerLoads := []string{
		"ProfilePicture",
		"Roles.Permissions",
	}

	listResponse, err := r.List(req, eagerLoads)
	if err != nil {
		return repository.ListResponse[resources.AdminResource]{}, err
	}

	adminResources := make([]resources.AdminResource, len(listResponse.Data))
	for i, admin := range listResponse.Data {
		adminResources[i] = resources.NewAdminResource(admin)
	}

	response := repository.ListResponse[resources.AdminResource]{
		Data:       adminResources,
		Pagination: listResponse.Pagination,
	}

	return response, nil
}

func (r *AdminService) CreateAdmin(adminInput models.Admin) (resources.AdminResource, error) {

	if err := r.Create(&adminInput); err != nil {
		return resources.AdminResource{}, err
	}

	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.AdminCreated,
		Payload: adminInput,
	})

	return r.GetAdminByID(adminInput.ID)
}

func (r *AdminService) UpdateAdmin(adminInput models.Admin) (resources.AdminResource, error) {
	if err := r.Update(&adminInput); err != nil {
		return resources.AdminResource{}, err
	}
	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.AdminUpdated,
		Payload: adminInput,
	})
	return r.GetAdminByID(adminInput.ID)
}

func (r *AdminService) DeleteAdmin(id string) error {
	admin, err := r.FindByID(id, nil)
	if err != nil {
		return err
	}
	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.AdminDeleted,
		Payload: admin,
	})

	return r.Delete(&admin)
}
