package models

import (
	"time"

	"gorm.io/gorm"
)

type OwnerStatus string

const (
	OwnerPending    OwnerStatus = "Pending"
	OwnerVerified   OwnerStatus = "Verified"
	OwnerNotAllowed OwnerStatus = "Not allowed"
)

type Owner struct {
	gorm.Model
	FirstName          string      `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName           string      `gorm:"type:varchar(255);not null" json:"last_name"`
	MiddleName         string      `gorm:"type:varchar(255)" json:"middle_name"`
	PermanentAddress   string      `gorm:"type:text" json:"permanent_address"`
	Description        string      `gorm:"type:text" json:"description"`
	Birthdate          time.Time   `gorm:"type:date" json:"birth_date"`
	Username           string      `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email              string      `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password           string      `gorm:"type:varchar(255);not null" json:"password"`
	ContactNumber      string      `gorm:"type:varchar(15);unique;not null" json:"contact_number"`
	IsEmailVerified    bool        `gorm:"default:false" json:"is_email_verified"`
	IsContactVerified  bool        `gorm:"default:false" json:"is_contact_verified"`
	IsSkipVerification bool        `gorm:"default:false" json:"is_skip_verification"`
	Status             OwnerStatus `gorm:"type:varchar(11);default:'Pending'" json:"status"`
	MediaID            *uint       `gorm:"type:bigint;unsigned" json:"media_id"`
	Media              *Media      `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"media"`

	Companies []*Company  `gorm:"foreignKey:OwnerID" json:"companies"`
	GenderID  *uint       `gorm:"type:bigint;unsigned" json:"gender_id"`
	Gender    *Gender     `gorm:"foreignKey:GenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"gender"`
	Footsteps []*Footstep `gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"footsteps,omitempty"`
}
