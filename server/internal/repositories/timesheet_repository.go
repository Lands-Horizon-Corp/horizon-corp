package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type TimesheetRepository struct {
	*Repository[models.Timesheet]
}

func NewTimesheetRepository(db *gorm.DB) *TimesheetRepository {
	return &TimesheetRepository{
		Repository: NewRepository[models.Timesheet](db),
	}
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
