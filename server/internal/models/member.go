package models

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model

	FirstName         string    `gorm:"type:varchar(255);not null;index" json:"first_name"`
	LastName          string    `gorm:"type:varchar(255);not null;index" json:"last_name"`
	MiddleName        string    `gorm:"type:varchar(255)" json:"middle_name"`
	PermanentAddress  string    `gorm:"type:text" json:"permanent_address"`
	Description       string    `gorm:"type:text" json:"description"`
	BirthDate         time.Time `gorm:"type:date;not null" json:"birth_date"`
	Username          string    `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email             string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password          string    `gorm:"type:varchar(255);not null" json:"password"`
	IsEmailVerified   bool      `gorm:"default:false" json:"is_email_verified"`
	IsContactVerified bool      `gorm:"default:false" json:"is_contact_verified"`
	ContactNumber     string    `gorm:"type:varchar(255);unique;not null" json:"contact_number"`
	Status            string    `gorm:"type:varchar(50);default:'Pending'" json:"status"`

	// Nullable Foreign Keys
	MediaID  *uint   `gorm:"type:bigint;unsigned;index" json:"media_id"`
	Media    *Media  `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media"`
	BranchID *uint   `gorm:"type:bigint;unsigned;index" json:"branch_id"`
	Branch   *Branch `gorm:"foreignKey:BranchID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"branch"`
	RoleID   *uint   `gorm:"type:bigint;unsigned;index" json:"role_id"`
	Role     *Role   `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`

	GenderID *uint   `gorm:"type:bigint;unsigned;index" json:"gender_id"`
	Gender   *Gender `gorm:"foreignKey:GenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"gender"`

	// Location
	Longitude *float64 `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude  *float64 `gorm:"type:decimal(10,7)" json:"latitude"`

	// Relationships
	Footsteps []*Footstep `gorm:"foreignKey:MemberID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"footsteps,omitempty"`
}
