package models

import (
	"time"

	"gorm.io/gorm"
)

type Timesheet struct {
	gorm.Model
	EmployeeID uint       `gorm:"type:bigint;unsigned;not null;index" json:"employee_id"`
	Employee   *Employee  `gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"employee"`
	TimeIn     *time.Time `gorm:"type:datetime(3)" json:"time_in"`
	TimeOut    *time.Time `gorm:"type:datetime(3)" json:"time_out"`
	MediaInID  *uint      `gorm:"type:bigint;unsigned" json:"media_in_id"`
	MediaIn    *Media     `gorm:"foreignKey:MediaInID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media_in"`
	MediaOutID *uint      `gorm:"type:bigint;unsigned" json:"media_out_id"`
	MediaOut   *Media     `gorm:"foreignKey:MediaOutID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media_out"`
}
