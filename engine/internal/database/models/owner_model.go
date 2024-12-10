package models

import (
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Owner struct {
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
	GenderID *uint `gorm:"type:bigint;unsigned" json:"gender_id"`

	// Relationship 0 to 1
	Gender *Gender `gorm:"foreignKey:GenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"gender"`

	// Relationship 0 to many
	Footsteps []*Footstep `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"footsteps,omitempty"`
	Companies []*Company  `gorm:"foreignKey:OwnerID" json:"companies"`

	// Relationship 0 to many
	RoleID *uint `gorm:"type:bigint;unsigned" json:"role_id"`
	Role   *Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`
}

type OwnerResource struct {
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
	GenderID           *uint               `json:"genderID"`
	Gender             *GenderResource     `json:"gender"`
	RoleID             *uint               `json:"roleID"`
	Role               *RoleResource       `json:"role"`
	Footsteps          []*FootstepResource `json:"footsteps"`
	Companies          []*CompanyResource  `json:"companies"`
}

type (
	OwnerResourceProvider interface {
		SeedDatabase()
		ToResource(owner *Owner) *OwnerResource
		ToResourceList(owner []*Owner) []*OwnerResource
	}
)

type OwnerModel struct {
	lc            *fx.Lifecycle
	db            *gorm.DB
	logger        *zap.Logger
	mediaModel    MediaResourceProvider
	genderModel   GenderResourceProvider
	footstepModel FootstepResourceProvider
	companyModel  CompanyResourceProvider
	roleModel     RoleResourceProvider
}

func NewOwnerModel(
	lc *fx.Lifecycle,
	db *gorm.DB,
	logger *zap.Logger,
	mediaModel MediaResourceProvider,
	genderModel GenderResourceProvider,
	footstepModel FootstepResourceProvider,
	companyModel CompanyResourceProvider,
	roleModel RoleResourceProvider,
) *OwnerModel {
	return &OwnerModel{
		lc:            lc,
		db:            db,
		logger:        logger,
		mediaModel:    mediaModel,
		genderModel:   genderModel,
		footstepModel: footstepModel,
		companyModel:  companyModel,
		roleModel:     roleModel,
	}
}

func (mm *OwnerModel) SeedDatabase() {
}

func (mm *OwnerModel) ToResource(owner *Owner) *OwnerResource {
	if owner == nil {
		return nil
	}

	return &OwnerResource{
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
		Media:              mm.mediaModel.ToResource(owner.Media),
		GenderID:           owner.GenderID,
		Gender:             mm.genderModel.ToResource(owner.Gender),
		RoleID:             owner.RoleID,
		Role:               mm.roleModel.ToResource(owner.Role),
		Footsteps:          mm.footstepModel.ToResourceList(owner.Footsteps),
		Companies:          mm.companyModel.ToResourceList(owner.Companies),
	}
}

func (mm *OwnerModel) ToResourceList(owners []*Owner) []*OwnerResource {
	if owners == nil {
		return nil
	}
	var ownerResources []*OwnerResource
	for _, owner := range owners {
		ownerResources = append(ownerResources, mm.ToResource(owner))
	}
	return ownerResources
}
