package repositories

import (
	"horizon/server/internal/models"
	"regexp"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) Create(employee *models.Employee) error {
	err := r.DB.Create(employee).Error
	return handleDBError(err)
}

func (r *EmployeeRepository) GetAll() ([]models.Employee, error) {
	var employees []models.Employee
	err := r.DB.Preload("Media").Find(&employees).Error
	return employees, handleDBError(err)
}

func (r *EmployeeRepository) GetByID(id uint) (models.Employee, error) {
	var employee models.Employee
	err := r.DB.Preload("Media").First(&employee, id).Error
	return employee, handleDBError(err)
}

func (r *EmployeeRepository) Update(id uint, employee *models.Employee) error {
	employee.ID = id
	err := r.DB.Save(employee).Error
	return handleDBError(err)
}

func (r *EmployeeRepository) Delete(id uint) error {
	err := r.DB.Delete(&models.Member{}, id).Error
	return handleDBError(err)
}

func (r *EmployeeRepository) FindByEmailUsernameOrContact(input string) (models.Employee, error) {
	var employee models.Employee

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	phoneRegex := `^\+?[1-9]\d{1,14}$`

	isEmail, _ := regexp.MatchString(emailRegex, input)
	isPhone, _ := regexp.MatchString(phoneRegex, input)

	if isEmail {
		err := r.DB.Preload("Media").Where("email = ?", input).First(&employee).Error
		return employee, handleDBError(err)
	} else if isPhone {
		err := r.DB.Preload("Media").Where("contact_number = ?", input).First(&employee).Error
		return employee, handleDBError(err)
	} else {
		err := r.DB.Preload("Media").Where("username = ?", input).First(&employee).Error
		return employee, handleDBError(err)
	}
}

func (r *EmployeeRepository) UpdateColumns(id uint, columns map[string]interface{}) (models.Employee, error) {
	var employee models.Employee
	if err := r.DB.Model(&employee).Where("id = ?", id).Updates(columns).Error; err != nil {
		return employee, handleDBError(err)
	}
	if err := r.DB.Preload("Media").First(&employee, id).Error; err != nil {
		return employee, handleDBError(err)
	}
	return employee, nil
}
