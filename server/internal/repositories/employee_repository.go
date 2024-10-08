package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) Create(employee *models.Employee) error {
	return r.DB.Create(employee).Error
}

func (r *EmployeeRepository) GetAll() ([]models.Employee, error) {
	var employee []models.Employee
	err := r.DB.Find(&employee).Error
	return employee, err
}

func (r *EmployeeRepository) GetByID(id uint) (models.Employee, error) {
	var employee models.Employee
	err := r.DB.First(&employee, id).Error
	return employee, err
}

func (r *EmployeeRepository) Update(id uint, employee *models.Employee) error {
	employee.ID = id
	return r.DB.Save(employee).Error
}

func (r *EmployeeRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Employee{}, id).Error
}
