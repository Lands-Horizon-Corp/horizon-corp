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

type MemberResource struct {
	ID                 uint               `json:"id"`
	CreatedAt          string             `json:"createdAt"`
	UpdatedAt          string             `json:"updatedAt"`
	FirstName          string             `json:"firstName"`
	LastName           string             `json:"lastName"`
	MiddleName         string             `json:"middleName,omitempty"`
	PermanentAddress   string             `json:"permanentAddress,omitempty"`
	Description        string             `json:"description,omitempty"`
	BirthDate          time.Time          `json:"birthDate"`
	Username           string             `json:"username"`
	Email              string             `json:"email"`
	IsEmailVerified    bool               `json:"isEmailVerified"`
	IsContactVerified  bool               `json:"isContactVerified"`
	IsSkipVerification bool               `json:"isSkipVerification"`
	ContactNumber      string             `json:"contactNumber"`
	Status             string             `json:"status"`
	Media              *MediaResource     `json:"media,omitempty"`
	Branch             *BranchResource    `json:"branch,omitempty"`
	Longitude          *float64           `json:"longitude,omitempty"`
	Latitude           *float64           `json:"latitude,omitempty"`
	Role               *RoleResource      `json:"role,omitempty"`
	Gender             *GenderResource    `json:"gender,omitempty"`
	Footsteps          []FootstepResource `json:"footsteps,omitempty"`
}
