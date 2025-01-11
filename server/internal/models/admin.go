package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Fields
	FirstName          string     `json:"first_name" gorm:"type:varchar(255);unsigned"`
	LastName           string     `json:"last_name" gorm:"type:varchar(255);unsigned"`
	MiddleName         string     `json:"middle_name" gorm:"type:varchar(255)"`
	PermanentAddress   string     `json:"permanent_address" gorm:"type:text"`
	Description        string     `json:"description" gorm:"type:text"`
	BirthDate          time.Time  `json:"birth_date" gorm:"type:date"`
	Username           string     `json:"username" gorm:"type:varchar(255);unique;unsigned"`
	Email              string     `json:"email" gorm:"type:varchar(255);unique;unsigned"`
	Password           string     `json:"password" gorm:"type:varchar(255);unsigned"`
	ContactNumber      string     `json:"contact_number" gorm:"type:varchar(15);unique;unsigned"`
	IsEmailVerified    bool       `json:"is_email_verified" gorm:"default:false"`
	IsContactVerified  bool       `json:"is_contact_verified" gorm:"default:false"`
	IsSkipVerification bool       `json:"is_skip_verification" gorm:"default:false"`
	Status             UserStatus `json:"status" gorm:"type:varchar(11);default:'Pending'"`

	// Relationship 0 to 1
	MediaID *uuid.UUID `json:"media_id" gorm:"type:bigint;unsigned"`
	Media   *Media     `json:"media" gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// Relationship 0 to 1
	RoleID *uuid.UUID `json:"role_id" gorm:"type:bigint;unsigned"`
	Role   *Role      `json:"role" gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	// Relationship 0 to 1
	GenderID *uuid.UUID `json:"gender_id" gorm:"type:bigint;unsigned"`
	Gender   *Gender    `json:"gender" gorm:"foreignKey:GenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	// Relationship 0 to many
	Footsteps []*Footstep `json:"footsteps,omitempty" gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type AdminResource struct {
	AccountType string `json:"accountType"`

	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

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
	MediaID            *uuid.UUID          `json:"mediaID"`
	Media              *MediaResource      `json:"media"`
	RoleID             *uuid.UUID          `json:"roleID"`
	Role               *RoleResource       `json:"role"`
	GenderID           *uuid.UUID          `json:"genderID"`
	Gender             *GenderResource     `json:"gender"`
	Footsteps          []*FootstepResource `json:"footsteps"`
}

func (m *ModelTransformer) AdminToResource(admin *Admin) *AdminResource {
	if admin == nil {
		return nil
	}

	return &AdminResource{
		AccountType: "Admin",

		ID:        admin.ID,
		CreatedAt: admin.CreatedAt.Format(time.RFC3339),
		UpdatedAt: admin.UpdatedAt.Format(time.RFC3339),
		DeletedAt: admin.DeletedAt.Time.Format(time.RFC3339),

		FirstName:          admin.FirstName,
		LastName:           admin.LastName,
		MiddleName:         admin.MiddleName,
		PermanentAddress:   admin.PermanentAddress,
		Description:        admin.Description,
		BirthDate:          admin.BirthDate,
		Username:           admin.Username,
		Email:              admin.Email,
		ContactNumber:      admin.ContactNumber,
		IsEmailVerified:    admin.IsEmailVerified,
		IsContactVerified:  admin.IsContactVerified,
		IsSkipVerification: admin.IsSkipVerification,
		Status:             admin.Status,
		MediaID:            admin.MediaID,
		Media:              m.MediaToResource(admin.Media),
		RoleID:             admin.RoleID,
		Role:               m.RoleToResource(admin.Role),
		GenderID:           admin.GenderID,
		Gender:             m.GenderToResource(admin.Gender),
		Footsteps:          m.FootstepToResourceList(admin.Footsteps),
	}
}

func (m *ModelTransformer) AdminToResourceList(adminList []*Admin) []*AdminResource {
	if adminList == nil {
		return nil
	}

	var adminResources []*AdminResource
	for _, admin := range adminList {
		adminResources = append(adminResources, m.AdminToResource(admin))
	}
	return adminResources
}
