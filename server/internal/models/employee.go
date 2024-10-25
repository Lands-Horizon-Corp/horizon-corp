package models

import (
	"time"

	"gorm.io/gorm"
)

type EmployeeStatus string

const (
	EmployeePending    EmployeeStatus = "Pending"
	EmployeeVerified   EmployeeStatus = "Verified"
	EmployeeNotAllowed EmployeeStatus = "Not allowed"
)

type Employee struct {
	gorm.Model
	FirstName          string         `gorm:"size:255;not null" json:"first_name"`
	LastName           string         `gorm:"size:255;not null" json:"last_name"`
	MiddleName         string         `gorm:"size:255" json:"middle_name"`
	PermanentAddress   string         `gorm:"size:500" json:"permanent_address"`
	Description        string         `gorm:"size:1000" json:"description"`
	Birthdate          time.Time      `gorm:"not null" json:"birthdate"`
	Username           string         `gorm:"size:255;not null;unique" json:"username"`
	Email              string         `gorm:"size:255;not null;unique" json:"email"`
	Password           string         `gorm:"size:255;not null" json:"password"`
	IsEmailVerified    bool           `gorm:"default:false" json:"is_email_verified"`
	IsContactVerified  bool           `gorm:"default:false" json:"is_contact_verified"`
	IsSkipVerification bool           `gorm:"default:false" json:"is_skip_verification"`
	ContactNumber      string         `gorm:"size:255;not null;unique" json:"contact_number"`
	Status             EmployeeStatus `gorm:"type:varchar(255);default:'Pending'" json:"status"`
	MediaID            *uint          `gorm:"default:NULL" json:"media_id"`
	Media              Media          `gorm:"foreignKey:MediaID" json:"media"`

	// Replace Company with Branch
	BranchID *uint  `gorm:"default:NULL" json:"branch_id"` // Nullable foreign key for Branch
	Branch   Branch `gorm:"foreignKey:BranchID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"branch"`

	Timesheets []Timesheet `gorm:"foreignKey:EmployeeID"`

	Longitude *float64 `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude  *float64 `gorm:"type:decimal(10,7)" json:"latitude"`
}
