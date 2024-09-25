package models

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	ID                 string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email              string    `gorm:"type:varchar(255);not null;unique"`
	Password           string    `gorm:"type:varchar(255);not null"`
	Username           string    `gorm:"type:varchar(255);not null"`
	FirstName          string    `gorm:"type:varchar(255);not null"`
	LastName           string    `gorm:"type:varchar(255);not null"`
	ContactNumber      string    `gorm:"type:varchar(50)"`
	PermanentAddress   string    `gorm:"type:varchar(255)"`
	Description        string    `gorm:"type:text"`
	Birthdate          time.Time `gorm:"type:date"`
	ValidEmail         bool      `gorm:"default:false"`
	ValidContactNumber bool      `gorm:"default:false"`

	// Foreign key to Branch (optional, Member may belong to a Branch)
	BranchID *string `gorm:"type:uuid;null"`
	Branch   *Branch `gorm:"foreignKey:BranchID;references:ID"`

	// Foreign key to Media table (optional)
	MediaID        *string `gorm:"type:uuid;null"`
	ProfilePicture *Media  `gorm:"foreignKey:MediaID;references:ID"`

	// Many-to-many relation with Role
	Roles []Role `gorm:"many2many:member_roles"`

	gorm.Model
}

type MemberRole struct {
	MemberID string `gorm:"type:uuid;primaryKey"`
	Member   Member `gorm:"foreignKey:MemberID;references:ID"`

	RoleID string `gorm:"type:uuid;primaryKey"`
	Role   Role   `gorm:"foreignKey:RoleID;references:ID"`

	AssignedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}