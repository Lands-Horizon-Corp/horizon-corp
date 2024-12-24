package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers/filter"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Timesheet struct {
	gorm.Model
	EmployeeID uint       `gorm:"type:bigint;unsigned;not null;index" json:"employee_id"`
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

type TimeInRequest struct {
	TimeIn  time.Time    `json:"timeIn" validate:"required"`
	MediaIn MediaRequest `json:"mediaIn" validate:"required"`
}

type TimeOutRequest struct {
	TimeOut  time.Time    `json:"timeOut" validate:"required"`
	MediaOut MediaRequest `json:"mediaOut" validate:"required"`
}

func (m *ModelResource) TimesheetToResource(timesheet *Timesheet) *TimesheetResource {
	if timesheet == nil {
		return nil
	}
	return &TimesheetResource{
		ID:        timesheet.ID,
		CreatedAt: timesheet.CreatedAt.Format(time.RFC3339),
		UpdatedAt: timesheet.UpdatedAt.Format(time.RFC3339),

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

// TimesheetToRecord converts a slice of Timesheet pointers into CSV records and headers.
func (m *ModelResource) TimesheetToRecord(timesheets []*Timesheet) ([][]string, []string) {
	// Convert Timesheet structs to TimesheetResource structs
	resource := m.TimesheetToResourceList(timesheets)
	records := make([][]string, 0, len(resource))

	for _, timesheet := range resource {
		// Basic Fields
		id := strconv.Itoa(int(timesheet.ID))
		employeeID := strconv.Itoa(int(timesheet.EmployeeID))
		createdAt := sanitizeCSVField(timesheet.CreatedAt)
		updatedAt := sanitizeCSVField(timesheet.UpdatedAt)

		// Employee Details
		employeeName := "N/A"
		if timesheet.Employee != nil {
			employeeName = sanitizeCSVField(fmt.Sprintf("%s %s", timesheet.Employee.FirstName, timesheet.Employee.LastName))
		}

		// TimeIn and TimeOut
		timeIn := "N/A"
		if timesheet.TimeIn != nil {
			timeIn = timesheet.TimeIn.Format(time.RFC3339)
		}
		timeOut := "N/A"
		if timesheet.TimeOut != nil {
			timeOut = timesheet.TimeOut.Format(time.RFC3339)
		}

		// MediaIn Details
		mediaInURL := "N/A"
		if timesheet.MediaIn != nil {
			mediaInURL = sanitizeCSVField(timesheet.MediaIn.URL)
		}

		// MediaOut Details
		mediaOutURL := "N/A"
		if timesheet.MediaOut != nil {
			mediaOutURL = sanitizeCSVField(timesheet.MediaOut.URL)
		}

		// Assemble the record
		record := []string{
			id,
			employeeID,
			employeeName,
			timeIn,
			mediaInURL,
			timeOut,
			mediaOutURL,
			createdAt,
			updatedAt,
		}
		records = append(records, record)
	}

	headers := []string{
		"ID",
		"Employee ID",
		"Employee Name",
		"Time In",
		"Media In URL",
		"Time Out",
		"Media Out URL",
		"Created At",
		"Updated At",
	}

	return records, headers
}

func (m *ModelResource) ValidateTimeInRequest(req *TimeInRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return m.helpers.FormatValidationError(err)
	}
	return nil
}

func (m *ModelResource) ValidateTimeOutRequest(req *TimeOutRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return m.helpers.FormatValidationError(err)
	}
	return nil
}

func (m *ModelResource) TimesheetToResourceList(timesheets []*Timesheet) []*TimesheetResource {
	if timesheets == nil {
		return nil
	}
	var timesheetResources []*TimesheetResource
	for _, timesheet := range timesheets {
		timesheetResources = append(timesheetResources, m.TimesheetToResource(timesheet))
	}
	return timesheetResources
}

func (m *ModelResource) TimesheetCurrentForEmployee(employeeId uint) (*Timesheet, error) {
	var timesheet Timesheet
	err := m.db.Client.
		Preload("Employee").
		Preload("MediaIn").
		Preload("MediaOut").
		Where("employee_id = ? AND time_out IS NULL", employeeId).
		Order("created_at DESC").
		First(&timesheet).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.New("could not retrieve current timesheet")
	}
	return &timesheet, err
}

func (m *ModelResource) TimesheetFindallForEmployee(employeeId uint, filters string) (filter.FilterPages[Timesheet], error) {
	db := m.db.Client.Where("employee_id = ? AND time_out IS NULL", employeeId)
	return m.TimesheetDB.GetPaginatedResult(db, filters)
}

func (m *ModelResource) TimesheetSeeders() error {
	m.logger.Info("Seeding Timesheet")
	return nil
}
