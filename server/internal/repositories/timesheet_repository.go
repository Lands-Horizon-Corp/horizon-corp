package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type TimesheetRepository struct {
	DB *gorm.DB
}

func NewTimesheetRepository(db *gorm.DB) *TimesheetRepository {
	return &TimesheetRepository{DB: db}
}

// Create adds a new timesheet record to the database.
func (r *TimesheetRepository) Create(timesheet *models.Timesheet) error {
	err := r.DB.Create(timesheet).Error
	return handleDBError(err)
}

// GetAll retrieves all timesheets with eager loaded Employee and Media associations.
func (r *TimesheetRepository) GetAll() ([]models.Timesheet, error) {
	var timesheets []models.Timesheet
	// Eager load the Employee and Media associations for both MediaIn and MediaOut
	err := r.DB.Preload("Employee").Preload("MediaIn").Preload("MediaOut").Find(&timesheets).Error
	return timesheets, handleDBError(err)
}

// GetByID retrieves a timesheet by its ID with eager loading.
func (r *TimesheetRepository) GetByID(id uint) (models.Timesheet, error) {
	var timesheet models.Timesheet
	// Eager load the Employee and Media associations
	err := r.DB.Preload("Employee").Preload("MediaIn").Preload("MediaOut").First(&timesheet, id).Error
	return timesheet, handleDBError(err)
}

// Update modifies an existing timesheet
func (r *TimesheetRepository) Update(id uint, timesheet *models.Timesheet) error {
	timesheet.ID = id
	err := r.DB.Save(timesheet).Error
	return handleDBError(err)
}

func (r *TimesheetRepository) Delete(id uint) error {
	err := r.DB.Delete(&models.Timesheet{}, id).Error
	return handleDBError(err)
}

func (r *TimesheetRepository) GetByEmployeeID(employeeID uint) ([]models.Timesheet, error) {
	var timesheets []models.Timesheet
	err := r.DB.Preload("Employee").Preload("MediaIn").Preload("MediaOut").
		Where("employee_id = ?", employeeID).Find(&timesheets).Error
	return timesheets, handleDBError(err)
}

func (r *TimesheetRepository) GetLatestForEmployee(employeeID uint) (models.Timesheet, error) {
	var timesheet models.Timesheet
	err := r.DB.Preload("Employee").Preload("MediaIn").Preload("MediaOut").
		Where("employee_id = ?", employeeID).Order("created_at desc").First(&timesheet).Error
	return timesheet, handleDBError(err)
}
