package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberStatus string

const (
	MemberPending    MemberStatus = "Pending"
	MemberVerified   MemberStatus = "Verified"
	MemberNotAllowed MemberStatus = "Not allowed"
)

type Member struct {
	gorm.Model
	FirstName          string       `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName           string       `gorm:"type:varchar(255);not null" json:"last_name"`
	MiddleName         string       `gorm:"type:varchar(255)" json:"middle_name"`
	PermanentAddress   string       `gorm:"type:text" json:"permanent_address"`
	Description        string       `gorm:"type:text" json:"description"`
	BirthDate          time.Time    `gorm:"type:date;not null" json:"birth_date"`
	Username           string       `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email              string       `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password           string       `gorm:"type:varchar(255);not null" json:"password"`
	IsEmailVerified    bool         `gorm:"default:false" json:"is_email_verified"`
	IsContactVerified  bool         `gorm:"default:false" json:"is_contact_verified"`
	IsSkipVerification bool         `gorm:"default:false" json:"is_skip_verification"`
	ContactNumber      string       `gorm:"type:varchar(255);unique;not null" json:"contact_number"`
	Status             MemberStatus `gorm:"type:varchar(255);default:'Pending'" json:"status"`
	MediaID            *uint        `gorm:"type:bigint;unsigned" json:"media_id"`
	Media              *Media       `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media"`
	BranchID           *uint        `gorm:"type:bigint;unsigned" json:"branch_id"`
	Branch             *Branch      `gorm:"foreignKey:BranchID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"branch"`
	Longitude          *float64     `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude           *float64     `gorm:"type:decimal(10,7)" json:"latitude"`

	RoleID *uint `gorm:"type:bigint;unsigned" json:"role_id"`
	Role   *Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`

	GenderID  *uint       `gorm:"type:bigint;unsigned" json:"gender_id"`
	Gender    *Gender     `gorm:"foreignKey:GenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"gender"`
	Footsteps []*Footstep `gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"footsteps,omitempty"`
}
