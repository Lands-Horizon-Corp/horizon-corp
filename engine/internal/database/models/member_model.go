package models

import (
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model

	// Fields
	FirstName          string     `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName           string     `gorm:"type:varchar(255);not null" json:"last_name"`
	MiddleName         string     `gorm:"type:varchar(255)" json:"middle_name"`
	PermanentAddress   string     `gorm:"type:text" json:"permanent_address"`
	Description        string     `gorm:"type:text" json:"description"`
	BirthDate          time.Time  `gorm:"type:date;not null" json:"birth_date"`
	Username           string     `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email              string     `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password           string     `gorm:"type:varchar(255);not null" json:"password"`
	IsEmailVerified    bool       `gorm:"default:false" json:"is_email_verified"`
	IsContactVerified  bool       `gorm:"default:false" json:"is_contact_verified"`
	IsSkipVerification bool       `gorm:"default:false" json:"is_skip_verification"`
	ContactNumber      string     `gorm:"type:varchar(255);unique;not null" json:"contact_number"`
	Status             UserStatus `gorm:"type:varchar(255);default:'Pending'" json:"status"`

	// Relationship 0 to 1
	MediaID *uint  `gorm:"type:bigint;unsigned" json:"media_id"`
	Media   *Media `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media"`

	// Relationship 0 to 1
	BranchID *uint   `gorm:"type:bigint;unsigned" json:"branch_id"`
	Branch   *Branch `gorm:"foreignKey:BranchID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"branch"`

	// Relationship 0 to 1
	Longitude *float64 `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude  *float64 `gorm:"type:decimal(10,7)" json:"latitude"`

	// Relationship 0 to 1
	RoleID *uint `gorm:"type:bigint;unsigned" json:"role_id"`
	Role   *Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`

	// Relationship 0 to 1
	GenderID *uint   `gorm:"type:bigint;unsigned" json:"gender_id"`
	Gender   *Gender `gorm:"foreignKey:GenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"gender"`

	// Relationship 0 to many
	Footsteps []*Footstep `gorm:"foreignKey:MemberID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"footsteps,omitempty"`
}

type MemberResource struct {
	FirstName          string          `json:"firstName"`
	LastName           string          `json:"lastName"`
	MiddleName         string          `json:"middleName"`
	PermanentAddress   string          `json:"permanentAddress"`
	Description        string          `json:"description"`
	BirthDate          time.Time       `json:"birthDate"`
	Username           string          `json:"username"`
	Email              string          `json:"email"`
	Password           string          `json:"password"`
	IsEmailVerified    bool            `json:"isEmailVerified"`
	IsContactVerified  bool            `json:"isContactVerified"`
	IsSkipVerification bool            `json:"isSkipVerification"`
	ContactNumber      string          `json:"contactNumber"`
	Status             UserStatus      `json:"status"`
	MediaID            *uint           `json:"mediaID"`
	Media              *Media          `json:"media"`
	BranchID           *uint           `json:"branchID"`
	Branch             *BranchResource `json:"branch"`
	Longitude          *float64        `json:"longitude"`
	Latitude           *float64        `json:"latitude"`
	RoleID             *uint           `json:"roleID"`
	Role               *RoleResource   `json:"role"`
	GenderID           *uint           `json:"genderID"`
	Gender             *GenderResource `json:"gender"`

	Footsteps []*FootstepResource `json:"footsteps"`
}

type MemberModel struct {
	lc     *fx.Lifecycle
	db     *gorm.DB
	logger *zap.Logger
}

func NewMemberModel(
	lc *fx.Lifecycle,
	db *gorm.DB,
	logger *zap.Logger,
) *MemberModel {
	return &MemberModel{
		lc:     lc,
		db:     db,
		logger: logger,
	}
}

func (mm *MemberModel) SeedDatabase() {
}

func (mm *MemberModel) ToResource() {
}

func (mm *MemberModel) ToResourceList() {
}
