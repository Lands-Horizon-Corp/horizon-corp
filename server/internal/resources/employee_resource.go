package resources

import (
	"horizon/server/internal/models"
	"time"
)

type EmployeeResource struct {
	ID                uint          `json:"id"`
	FirstName         string        `json:"firstName"`
	LastName          string        `json:"lastName"`
	PermanentAddress  string        `json:"permanentAddress"`
	Description       string        `json:"description"`
	Birthdate         string        `json:"birthdate"`
	Username          string        `json:"username"`
	Email             string        `json:"email"`
	IsEmailVerified   bool          `json:"isEmailVerified"`
	IsContactVerified bool          `json:"isContactVerified"`
	Media             MediaResource `json:"media"`
}

func ToResourceEmployee(employee models.Employee) EmployeeResource {
	return EmployeeResource{
		ID:                employee.ID,
		FirstName:         employee.FirstName,
		LastName:          employee.LastName,
		PermanentAddress:  employee.PermanentAddress,
		Description:       employee.Description,
		Birthdate:         employee.Birthdate.Format(time.RFC3339),
		Username:          employee.Username,
		Email:             employee.Email,
		IsEmailVerified:   employee.IsEmailVerified,
		IsContactVerified: employee.IsContactVerified,
		Media:             ToResourceMedia(employee.Media),
	}
}

func ToResourceListEmployee(employeeList []models.Employee) []EmployeeResource {
	var resources []EmployeeResource
	for _, employee := range employeeList {
		resources = append(resources, ToResourceEmployee(employee))
	}
	return resources
}
