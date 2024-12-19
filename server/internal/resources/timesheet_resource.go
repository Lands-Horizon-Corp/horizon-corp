package resources

import (
	"horizon/server/internal/models"
	"time"
)

type TimesheetResource struct {
	ID         uint              `json:"id"`
	EmployeeID uint              `json:"employeeId"`
	Employee   *EmployeeResource `json:"employee,omitempty"`
	TimeIn     *time.Time        `json:"timeIn"`
	TimeOut    *time.Time        `json:"timeOut"`
	MediaInID  *uint             `json:"mediaInId"`
	MediaIn    *MediaResource    `json:"mediaIn,omitempty"`
	MediaOutID *uint             `json:"mediaOutId"`
	MediaOut   *MediaResource    `json:"mediaOut,omitempty"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Convert models.Timesheet to *TimesheetResource
func ToResourceTimesheet(timesheet *models.Timesheet) *TimesheetResource {
	if timesheet == nil {
		return nil
	}

	var employeeResource *EmployeeResource
	if timesheet.Employee != nil {
		employeeResource = ToResourceEmployee(timesheet.Employee)
	}

	var mediaInResource *MediaResource
	if timesheet.MediaIn != nil {
		mediaInResource = ToResourceMedia(timesheet.MediaIn)
	}

	var mediaOutResource *MediaResource
	if timesheet.MediaOut != nil {
		mediaOutResource = ToResourceMedia(timesheet.MediaOut)
	}

	return &TimesheetResource{
		ID:         timesheet.ID,
		EmployeeID: timesheet.EmployeeID,
		Employee:   employeeResource,
		TimeIn:     timesheet.TimeIn,
		TimeOut:    timesheet.TimeOut,
		MediaInID:  timesheet.MediaInID,
		MediaIn:    mediaInResource,
		MediaOutID: timesheet.MediaOutID,
		MediaOut:   mediaOutResource,
		CreatedAt:  timesheet.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  timesheet.UpdatedAt.Format(time.RFC3339),
	}
}

// Convert []*models.Timesheet to []*TimesheetResource
func ToResourceListTimesheets(timesheets []*models.Timesheet) []*TimesheetResource {
	resourceList := make([]*TimesheetResource, len(timesheets))
	for i, timesheet := range timesheets {
		resourceList[i] = ToResourceTimesheet(timesheet)
	}
	return resourceList
}
