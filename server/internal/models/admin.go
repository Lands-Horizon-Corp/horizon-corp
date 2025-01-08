package models

import (
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model

	// Fields
	FirstName          string     `json:"first_name" gorm:"type:varchar(255);not null"`
	LastName           string     `json:"last_name" gorm:"type:varchar(255);not null"`
	MiddleName         string     `json:"middle_name" gorm:"type:varchar(255)"`
	PermanentAddress   string     `json:"permanent_address" gorm:"type:text"`
	Description        string     `json:"description" gorm:"type:text"`
	BirthDate          time.Time  `json:"birth_date" gorm:"type:date"`
	Username           string     `json:"username" gorm:"type:varchar(255);unique;not null"`
	Email              string     `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password           string     `json:"password" gorm:"type:varchar(255);not null"`
	ContactNumber      string     `json:"contact_number" gorm:"type:varchar(15);unique;not null"`
	IsEmailVerified    bool       `json:"is_email_verified" gorm:"default:false"`
	IsContactVerified  bool       `json:"is_contact_verified" gorm:"default:false"`
	IsSkipVerification bool       `json:"is_skip_verification" gorm:"default:false"`
	Status             UserStatus `json:"status" gorm:"type:varchar(11);default:'Pending'"`

	// Relationship 0 to 1
	MediaID *uint  `json:"media_id" gorm:"type:bigint;unsigned"`
	Media   *Media `json:"media" gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// Relationship 0 to 1
	RoleID *uint `json:"role_id" gorm:"type:bigint;unsigned"`
	Role   *Role `json:"role" gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	// Relationship 0 to 1
	GenderID *uint   `json:"gender_id" gorm:"type:bigint;unsigned"`
	Gender   *Gender `json:"gender" gorm:"foreignKey:GenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	// Relationship 0 to many
	Footsteps []*Footstep `json:"footsteps,omitempty" gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type AdminResource struct {
	AccountType string `json:"accountType"`
	ID          uint   `json:"id"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`

	FirstName          string              `json:"firstName"`
	LastName           string              `json:"lastName"`
	MiddleName         string              `json:"middleName"`
	PermanentAddress   string              `json:"permanentAddress"`
	Description        string              `json:"description"`
	BirthDate          time.Time           `json:"birthDate"`
	Username           string              `json:"username"`
	Email              string              `json:"email"`
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

type AdminRepository struct {
	*Repository[Admin]
}

func NewAdminRepository(db *providers.DatabaseService) *AdminRepository {
	return &AdminRepository{
		Repository: NewRepository[Admin](db),
	}
}

func (r *AdminRepository) DatabaseSeeder() {}
