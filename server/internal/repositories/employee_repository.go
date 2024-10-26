package repositories

import (
	"horizon/server/config"
	"horizon/server/helpers"
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	*Repository[models.Employee]
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{
		Repository: NewRepository[models.Employee](db),
	}
}

func (r *EmployeeRepository) GetByEmail(email string) (*models.Employee, error) {
	var employee models.Employee
	err := r.DB.Preload("Media").Where("email = ?", email).First(&employee).Error
	return &employee, handleDBError(err)
}

func (r *EmployeeRepository) GetByContactNumber(contactNumber string) (*models.Employee, error) {
	var employee models.Employee
	err := r.DB.Preload("Media").Where("contact_number = ?", contactNumber).First(&employee).Error
	return &employee, handleDBError(err)
}

func (r *EmployeeRepository) GetByUsername(username string) (*models.Employee, error) {
	var employee models.Employee
	err := r.DB.Preload("Media").Where("username = ?", username).First(&employee).Error
	return &employee, handleDBError(err)
}

func (r *EmployeeRepository) FindByEmailUsernameOrContact(input string) (*models.Employee, error) {
	switch helpers.GetKeyType(input) {
	case "contact":
		return r.GetByContactNumber(input)
	case "email":
		return r.GetByEmail(input)
	default:
		return r.GetByUsername(input)
	}
}

func (r *EmployeeRepository) UpdatePassword(id uint, password string) error {
	newPassword, err := config.HashPassword(password)
	if err != nil {
		return err
	}
	updated := &models.Employee{Password: newPassword}
	_, err = r.Repository.UpdateColumns(id, *updated)
	return err
}
