package resources

import (
	"horizon/server/internal/models"
	"time"
)

type TimesheetResource struct {
	ID         uint           `json:"id"`
	EmployeeID uint           `json:"employeeId"`
	TimeIn     *time.Time     `json:"timeIn"`
	TimeOut    *time.Time     `json:"timeOut"`
	MediaIn    *MediaResource `json:"mediaIn,omitempty"`
	MediaOut   *MediaResource `json:"mediaOut,omitempty"`
	CreatedAt  string         `json:"createdAt"`
	UpdatedAt  string         `json:"updatedAt"`
}

func ToResourceTimesheet(timesheet models.Timesheet) TimesheetResource {
	var mediaInResource, mediaOutResource *MediaResource

	if timesheet.MediaInID != nil {
		mediaIn := ToResourceMedia(timesheet.MediaIn)
		mediaInResource = &mediaIn
	}
	if timesheet.MediaOutID != nil {
		mediaOut := ToResourceMedia(timesheet.MediaOut)
		mediaOutResource = &mediaOut
	}

	return TimesheetResource{
		ID:         timesheet.ID,
		EmployeeID: timesheet.EmployeeID,
		TimeIn:     &timesheet.TimeIn,
		TimeOut:    timesheet.TimeOut,
		MediaIn:    mediaInResource,
		MediaOut:   mediaOutResource,
		CreatedAt:  timesheet.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  timesheet.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResourceListTimesheets(timesheets []models.Timesheet) []TimesheetResource {
	var resources []TimesheetResource
	for _, timesheet := range timesheets {
		resources = append(resources, ToResourceTimesheet(timesheet))
	}
	return resources
}
