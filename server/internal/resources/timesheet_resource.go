package resources

import (
	"horizon/server/internal/models"
	"time"
)

type TimesheetResource struct {
	ID         uint              `json:"id"`
	EmployeeID uint              `json:"employee_id"`
	Employee   *EmployeeResource `json:"employee,omitempty"`
	TimeIn     *time.Time        `json:"time_in"`
	TimeOut    *time.Time        `json:"time_out"`
	MediaInID  *uint             `json:"media_in_id"`
	MediaIn    *MediaResource    `json:"media_in,omitempty"`
	MediaOutID *uint             `json:"media_out_id"`
	MediaOut   *MediaResource    `json:"media_out,omitempty"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToResourceTimesheet(timesheet models.Timesheet) TimesheetResource {
	var employeeResource *EmployeeResource
	if timesheet.EmployeeID != 0 {
		empRes := ToResourceEmployee(timesheet.Employee)
		employeeResource = &empRes
	}

	var mediaInResource *MediaResource
	if timesheet.MediaIn != nil {
		mediaInRes := ToResourceMedia(*timesheet.MediaIn)
		mediaInResource = &mediaInRes
	}

	var mediaOutResource *MediaResource
	if timesheet.MediaOut != nil {
		mediaOutRes := ToResourceMedia(*timesheet.MediaOut)
		mediaOutResource = &mediaOutRes
	}

	return TimesheetResource{
		ID:         timesheet.ID,
		EmployeeID: timesheet.EmployeeID,
		Employee:   employeeResource,
		TimeIn:     timesheet.TimeIn,
		TimeOut:    timesheet.TimeOut,
		MediaInID:  timesheet.MediaInID,
		MediaIn:    mediaInResource,
		MediaOutID: timesheet.MediaOutID,
		MediaOut:   mediaOutResource,

		CreatedAt: timesheet.CreatedAt.Format(time.RFC3339),
		UpdatedAt: timesheet.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResourceListTimesheets(timesheets []models.Timesheet) []TimesheetResource {
	resourceList := make([]TimesheetResource, len(timesheets))
	for i, timesheet := range timesheets {
		resourceList[i] = ToResourceTimesheet(timesheet)
	}
	return resourceList
}
