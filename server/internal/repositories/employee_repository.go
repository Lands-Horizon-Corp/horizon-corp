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
	var employee []models.Employee
	err := r.DB.Find(&employee).Error
	return employee, handleDBError(err)
}

func (r *EmployeeRepository) GetByID(id uint) (models.Employee, error) {
	var employee models.Employee
	err := r.DB.First(&employee, id).Error
	return employee, handleDBError(err)
}

func (r *EmployeeRepository) Update(id uint, employee *models.Employee) error {
	employee.ID = id
	err := r.DB.Save(employee).Error
	return handleDBError(err)
}

func (r *EmployeeRepository) FindByEmailUsernameOrContact(input string) (models.Employee, error) {
	var employee models.Employee

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	phoneRegex := `^\+?[1-9]\d{1,14}$`

	isEmail, _ := regexp.MatchString(emailRegex, input)
	isPhone, _ := regexp.MatchString(phoneRegex, input)

	if isEmail {
		err := r.DB.Where("email = ?", input).First(&employee).Error
		return employee, handleDBError(err)
	} else if isPhone {
		err := r.DB.Where("contact_number = ?", input).First(&employee).Error
		return employee, handleDBError(err)
	} else {
		err := r.DB.Where("username = ?", input).First(&employee).Error
		return employee, handleDBError(err)
	}
}
