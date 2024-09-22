// package service

package service

import (
	"horizon-core/config"
	"horizon-core/internal/events"
	"horizon-core/internal/models"
	"horizon-core/internal/repository"
	"horizon-core/internal/resources"

	"gorm.io/gorm"
)

type EmployeeService struct {
	*repository.ModelRepository[models.Employee]
	config *config.Config
}

func NewEmployeeService(db *gorm.DB) *EmployeeService {
	return &EmployeeService{
		ModelRepository: repository.NewModelRepository[models.Employee](db),
		config:          config.GetConfig(),
	}
}

func (r *EmployeeService) GetEmployeeByID(id string) (resources.EmployeeResource, error) {
	eagerLoads := []string{
		"ProfilePicture",
		"Roles.Permissions",
		"Branch.ProfilePicture",
	}

	employee, err := r.FindByID(id, eagerLoads)
	if err != nil {
		return resources.EmployeeResource{}, err
	}

	employeeResource := resources.NewEmployeeResource(employee)
	return employeeResource, nil
}

func (r *EmployeeService) ListEmployees(req repository.ListRequest) (repository.ListResponse[resources.EmployeeResource], error) {
	eagerLoads := []string{
		"ProfilePicture",
		"Roles.Permissions",
		"Branch.ProfilePicture",
	}

	listResponse, err := r.List(req, eagerLoads)
	if err != nil {
		return repository.ListResponse[resources.EmployeeResource]{}, err
	}

	employeeResources := make([]resources.EmployeeResource, len(listResponse.Data))
	for i, employee := range listResponse.Data {
		employeeResources[i] = resources.NewEmployeeResource(employee)
	}

	response := repository.ListResponse[resources.EmployeeResource]{
		Data:       employeeResources,
		Pagination: listResponse.Pagination,
	}

	return response, nil
}

func (r *EmployeeService) CreateEmployee(employeeInput models.Employee) (resources.EmployeeResource, error) {
	if err := r.Create(&employeeInput); err != nil {
		return resources.EmployeeResource{}, err
	}

	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.EmployeeCreated,
		Payload: employeeInput,
	})
	return r.GetEmployeeByID(employeeInput.ID)
}

func (r *EmployeeService) UpdateEmployee(employeeInput models.Employee) (resources.EmployeeResource, error) {
	if err := r.Update(&employeeInput); err != nil {
		return resources.EmployeeResource{}, err
	}

	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.EmployeeUpdated,
		Payload: employeeInput,
	})

	return r.GetEmployeeByID(employeeInput.ID)
}

func (r *EmployeeService) DeleteEmployee(id string) error {
	employee, err := r.FindByID(id, nil)
	if err != nil {
		return err
	}

	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.EmployeeDeleted,
		Payload: employee,
	})
	return r.Delete(&employee)
}
