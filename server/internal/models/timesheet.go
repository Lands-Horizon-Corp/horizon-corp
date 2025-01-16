package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Timesheet struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

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

func (v *Timesheet) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type TimesheetResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

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

		ID:        timesheet.ID,
		CreatedAt: timesheet.CreatedAt.Format(time.RFC3339),
		UpdatedAt: timesheet.UpdatedAt.Format(time.RFC3339),
		DeletedAt: timesheet.DeletedAt.Time.Format(time.RFC3339),

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

func (m *ModelRepository) TimesheetGetByID(id string, preloads ...string) (*Timesheet, error) {
	repo := NewGenericRepository[Timesheet](m.db.Client)
	return repo.GetByID(id, preloads...)
}

func (m *ModelRepository) TimesheetCreate(timesheet *Timesheet, preloads ...string) (*Timesheet, error) {
	repo := NewGenericRepository[Timesheet](m.db.Client)
	return repo.Create(timesheet, preloads...)
}
func (m *ModelRepository) TimesheetUpdate(timesheet *Timesheet, preloads ...string) (*Timesheet, error) {
	repo := NewGenericRepository[Timesheet](m.db.Client)
	return repo.Update(timesheet, preloads...)
}
func (m *ModelRepository) TimesheetUpdateByID(id string, column string, value interface{}, preloads ...string) (*Timesheet, error) {
	repo := NewGenericRepository[Timesheet](m.db.Client)
	return repo.UpdateByID(id, column, value, preloads...)
}
func (m *ModelRepository) TimesheetDeleteByID(id string) error {
	repo := NewGenericRepository[Timesheet](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) TimesheetGetAll(preloads ...string) ([]*Timesheet, error) {
	repo := NewGenericRepository[Timesheet](m.db.Client)
	return repo.GetAll(preloads...)
}
