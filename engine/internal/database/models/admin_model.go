package models

import (
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model

	// Fields
	FirstName          string     `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName           string     `gorm:"type:varchar(255);not null" json:"last_name"`
	MiddleName         string     `gorm:"type:varchar(255)" json:"middle_name"`
	PermanentAddress   string     `gorm:"type:text" json:"permanent_address"`
	Description        string     `gorm:"type:text" json:"description"`
	BirthDate          time.Time  `gorm:"type:date" json:"birth_date"`
	Username           string     `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email              string     `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password           string     `gorm:"type:varchar(255);not null" json:"password"`
	ContactNumber      string     `gorm:"type:varchar(15);unique;not null" json:"contact_number"`
	IsEmailVerified    bool       `gorm:"default:false" json:"is_email_verified"`
	IsContactVerified  bool       `gorm:"default:false" json:"is_contact_verified"`
	IsSkipVerification bool       `gorm:"default:false" json:"is_skip_verification"`
	Status             UserStatus `gorm:"type:varchar(11);default:'Pending'" json:"status"`

	// Relationship 0 to 1
	MediaID *uint  `gorm:"type:bigint;unsigned" json:"media_id"`
	Media   *Media `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"media"`

	// Relationship 0 to 1
	RoleID *uint `gorm:"type:bigint;unsigned" json:"role_id"`
	Role   *Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`

	// Relationship 0 to 1
	GenderID *uint   `gorm:"type:bigint;unsigned" json:"gender_id"`
	Gender   *Gender `gorm:"foreignKey:GenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"gender"`

	// Relationship 0 to many
	Footsteps []*Footstep `gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"footsteps,omitempty"`
}

type AdminResource struct {
	FirstName          string              `json:"firstName"`
	LastName           string              `json:"lastName"`
	MiddleName         string              `json:"middleName"`
	PermanentAddress   string              `json:"permanentAddress"`
	Description        string              `json:"description"`
	BirthDate          time.Time           `json:"birthDate"`
	Username           string              `json:"username"`
	Email              string              `json:"email"`
	Password           string              `json:"password"`
	ContactNumber      string              `json:"contactNumber"`
	IsEmailVerified    bool                `json:"isEmailVerified"`
	IsContactVerified  bool                `json:"isContactVerified"`
	IsSkipVerification bool                `json:"isSkipVerification"`
	Status             UserStatus          `json:"status"`
	MediaID            *uint               `json:"mediaID"`
	Media              *MediaResource      `json:"media"`
	RoleID             *uint               `json:"roleID"`
	Role               *RoleResource       `json:"role"`
	GenderID           *uint               `json:"genderID"`
	Gender             *GenderResource     `json:"gender"`
	Footsteps          []*FootstepResource `json:"footsteps"`
}

type AdminModel struct {
	lc     *fx.Lifecycle
	db     *gorm.DB
	logger *zap.Logger
}

func NewAdminModel(
	lc *fx.Lifecycle,
	db *gorm.DB,
	logger *zap.Logger,
) *AdminModel {
	return &AdminModel{
		lc:     lc,
		db:     db,
		logger: logger,
	}
}

func (am *AdminModel) SeedDatabase() {

}

func (am *AdminModel) ToResource() {

}

func (am *AdminModel) ToResourceList() {

}
