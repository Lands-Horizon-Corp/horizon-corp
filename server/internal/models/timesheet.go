package models

import (
	"time"

	"gorm.io/gorm"
)

type Timesheet struct {
	gorm.Model
	EmployeeID uint       `gorm:"type:bigint;unsigned;unsigned;index" json:"employee_id"`
	Employee   *Employee  `gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"employee"`
	TimeIn     *time.Time `gorm:"type:datetime(3)" json:"time_in"`
	TimeOut    *time.Time `gorm:"type:datetime(3)" json:"time_out"`

	// Relationship 0 to 1
	MediaInID *uint  `gorm:"type:bigint;unsigned" json:"media_in_id"`
	MediaIn   *Media `gorm:"foreignKey:MediaInID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media_in"`

	// Relationship 0 to 1
	MediaOutID *uint  `gorm:"type:bigint;unsigned" json:"media_out_id"`
	MediaOut   *Media `gorm:"foreignKey:MediaOutID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media_out"`
}

type TimesheetResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	EmployeeID uint              `json:"employeeID"`
	Employee   *EmployeeResource `json:"employee"`
	TimeIn     *time.Time        `json:"timeIn"`
	TimeOut    *time.Time        `json:"timeOut"`
	MediaInID  *uint             `json:"mediaInID"`
	MediaIn    *MediaResource    `json:"mediaIn"`
	MediaOutID *uint             `json:"mediaOutID"`
	MediaOut   *MediaResource    `json:"mediaOut"`
}

func (m *ModelTransformer) TimesheetToResource(timesheet *Timesheet) *TimesheetResource {
	if timesheet == nil {
		return nil
	}

	return &TimesheetResource{
		ID:         timesheet.ID,
		CreatedAt:  timesheet.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  timesheet.UpdatedAt.Format(time.RFC3339),
		EmployeeID: timesheet.EmployeeID,
		Employee:   m.EmployeeToResource(timesheet.Employee),
		TimeIn:     timesheet.TimeIn,
		TimeOut:    timesheet.TimeOut,
		MediaInID:  timesheet.MediaInID,
		MediaIn:    m.MediaToResource(timesheet.MediaIn),
		MediaOutID: timesheet.MediaOutID,
		MediaOut:   m.MediaToResource(timesheet.MediaOut),
	}
}

func (m *ModelTransformer) TimesheetToResourceList(timesheetList []*Timesheet) []*TimesheetResource {
	if timesheetList == nil {
		return nil
	}

	var timesheetResources []*TimesheetResource
	for _, timesheet := range timesheetList {
		timesheetResources = append(timesheetResources, m.TimesheetToResource(timesheet))
	}
	return timesheetResources
}
