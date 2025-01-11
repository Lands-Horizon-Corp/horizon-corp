package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Owner struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Fields
	FirstName          string     `gorm:"type:varchar(255);unsigned" json:"first_name"`
	LastName           string     `gorm:"type:varchar(255);unsigned" json:"last_name"`
	MiddleName         string     `gorm:"type:varchar(255)" json:"middle_name"`
	PermanentAddress   string     `gorm:"type:text" json:"permanent_address"`
	Description        string     `gorm:"type:text" json:"description"`
	BirthDate          time.Time  `gorm:"type:date" json:"birth_date"`
	Username           string     `gorm:"type:varchar(255);unique;unsigned" json:"username"`
	Email              string     `gorm:"type:varchar(255);unique;unsigned" json:"email"`
	Password           string     `gorm:"type:varchar(255);unsigned" json:"password"`
	ContactNumber      string     `gorm:"type:varchar(15);unique;unsigned" json:"contact_number"`
	IsEmailVerified    bool       `gorm:"default:false" json:"is_email_verified"`
	IsContactVerified  bool       `gorm:"default:false" json:"is_contact_verified"`
	IsSkipVerification bool       `gorm:"default:false" json:"is_skip_verification"`
	Status             UserStatus `gorm:"type:varchar(11);default:'Pending'" json:"status"`

	// Relationship 0 to 1
	MediaID *uuid.UUID `gorm:"type:bigint;unsigned" json:"media_id"`
	Media   *Media     `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"media"`

	// Relationship 0 to 1
	GenderID *uuid.UUID `gorm:"type:bigint;unsigned" json:"gender_id"`

	// Relationship 0 to 1
	Gender *Gender `gorm:"foreignKey:GenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"gender"`

	// Relationship 0 to many
	Footsteps []*Footstep `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"footsteps,omitempty"`
	Companies []*Company  `gorm:"foreignKey:OwnerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"companies"`

	// Relationship 0 to many
	RoleID *uuid.UUID `gorm:"type:bigint;unsigned" json:"role_id"`
	Role   *Role      `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`
}

type OwnerResource struct {
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
	Password           string              `json:"password"`
	ContactNumber      string              `json:"contactNumber"`
	IsEmailVerified    bool                `json:"isEmailVerified"`
	IsContactVerified  bool                `json:"isContactVerified"`
	IsSkipVerification bool                `json:"isSkipVerification"`
	Status             UserStatus          `json:"status"`
	MediaID            *uuid.UUID          `json:"mediaID"`
	Media              *MediaResource      `json:"media"`
	GenderID           *uuid.UUID          `json:"genderID"`
	Gender             *GenderResource     `json:"gender"`
	RoleID             *uuid.UUID          `json:"roleID"`
	Role               *RoleResource       `json:"role"`
	Footsteps          []*FootstepResource `json:"footsteps"`
	Companies          []*CompanyResource  `json:"companies"`
}

func (m *ModelTransformer) OwnerToResource(owner *Owner) *OwnerResource {
	if owner == nil {
		return nil
	}

	return &OwnerResource{
		AccountType: "Owner",

		ID:        owner.ID,
		CreatedAt: owner.CreatedAt.Format(time.RFC3339),
		UpdatedAt: owner.UpdatedAt.Format(time.RFC3339),
		DeletedAt: owner.DeletedAt.Time.Format(time.RFC3339),

		FirstName:          owner.FirstName,
		LastName:           owner.LastName,
		MiddleName:         owner.MiddleName,
		PermanentAddress:   owner.PermanentAddress,
		Description:        owner.Description,
		BirthDate:          owner.BirthDate,
		Username:           owner.Username,
		Email:              owner.Email,
		Password:           owner.Password,
		ContactNumber:      owner.ContactNumber,
		IsEmailVerified:    owner.IsEmailVerified,
		IsContactVerified:  owner.IsContactVerified,
		IsSkipVerification: owner.IsSkipVerification,
		Status:             owner.Status,
		MediaID:            owner.MediaID,
		Media:              m.MediaToResource(owner.Media),
		GenderID:           owner.GenderID,
		Gender:             m.GenderToResource(owner.Gender),
		RoleID:             owner.RoleID,
		Role:               m.RoleToResource(owner.Role),
		Footsteps:          m.FootstepToResourceList(owner.Footsteps),
		Companies:          m.CompanyToResourceList(owner.Companies),
	}
}

func (m *ModelTransformer) OwnerToResourceList(ownerList []*Owner) []*OwnerResource {
	if ownerList == nil {
		return nil
	}

	var ownerResources []*OwnerResource
	for _, owner := range ownerList {
		ownerResources = append(ownerResources, m.OwnerToResource(owner))
	}
	return ownerResources
}
