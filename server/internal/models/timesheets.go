package models

import (
	"time"

	"gorm.io/gorm"
)

type Timesheet struct {
	gorm.Model
	ID         uint       `gorm:"primaryKey;autoIncrement;not null"`
	EmployeeID uint       `gorm:"not null;index"`
	TimeIn     time.Time  `gorm:"type:datetime(3);not null"`
	TimeOut    *time.Time `gorm:"type:datetime(3)"`
	MediaInID  *uint      `gorm:"index"`
	MediaOutID *uint      `gorm:"index"`
	CreatedAt  time.Time  `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3)"`
	UpdatedAt  time.Time  `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)"`
	MediaIn    Media      `gorm:"foreignKey:MediaInID"`
	MediaOut   Media      `gorm:"foreignKey:MediaOutID"`
}
