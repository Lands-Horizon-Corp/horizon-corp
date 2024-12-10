package models

import (
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
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
	EmployeeID uint              `json:"employeeID"`
	Employee   *EmployeeResource `json:"employee"`
	TimeIn     *time.Time        `json:"timeIn"`
	TimeOut    *time.Time        `json:"timeOut"`
	MediaInID  *uint             `json:"mediaInID"`
	MediaIn    *MediaResource    `json:"mediaIn"`
	MediaOutID *uint             `json:"mediaOutID"`
	MediaOut   *MediaResource    `json:"mediaOut"`
}

type (
	TimesheetResourceProvider interface {
		SeedDatabase()
		ToResource(timesheet *Timesheet) *TimesheetResource
		ToResourceList(timesheet []*Timesheet) []*TimesheetResource
	}
)

type TimesheetModel struct {
	lc            *fx.Lifecycle
	db            *gorm.DB
	logger        *zap.Logger
	employeeModel EmployeeResourceProvider
	mediaModel    MediaResourceProvider
}

func NewTimesheetModel(
	lc *fx.Lifecycle,
	db *gorm.DB,
	logger *zap.Logger,
	employeeModel EmployeeResourceProvider,
	mediaModel MediaResourceProvider,
) *TimesheetModel {
	return &TimesheetModel{
		lc:            lc,
		db:            db,
		logger:        logger,
		employeeModel: employeeModel,
		mediaModel:    mediaModel,
	}
}

func (mm *TimesheetModel) SeedDatabase() {
}

func (mm *TimesheetModel) ToResource(timesheet *Timesheet) *TimesheetResource {
	if timesheet == nil {
		return nil
	}

	return &TimesheetResource{
		EmployeeID: timesheet.EmployeeID,
		Employee:   mm.employeeModel.ToResource(timesheet.Employee),
		TimeIn:     timesheet.TimeIn,
		TimeOut:    timesheet.TimeOut,
		MediaInID:  timesheet.MediaInID,
		MediaIn:    mm.mediaModel.ToResource(timesheet.MediaIn),
		MediaOutID: timesheet.MediaOutID,
		MediaOut:   mm.mediaModel.ToResource(timesheet.MediaOut),
	}
}

func (mm *TimesheetModel) ToResourceList(timesheets []*Timesheet) []*TimesheetResource {
	if timesheets == nil {
		return nil
	}
	var timesheetResources []*TimesheetResource
	for _, timesheet := range timesheets {
		timesheetResources = append(timesheetResources, mm.ToResource(timesheet))
	}
	return timesheetResources
}
