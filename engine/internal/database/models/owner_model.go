package models

import (
	"errors"
	"time"

	"github.com/go-playground/validator"
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

func (m *ModelResource) OwnerToResource(owner *Owner) *OwnerResource {
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
		Media:              m.MediaToResource(owner.Media),
		GenderID:           owner.GenderID,
		Gender:             m.GenderToResource(owner.Gender),
		RoleID:             owner.RoleID,
		Role:               m.RoleToResource(owner.Role),
		Footsteps:          m.FootstepToResourceList(owner.Footsteps),
		Companies:          m.CompanyToResourceList(owner.Companies),
	}
}

type OwnerRequest struct {
	FirstName          string    `json:"firstName" validate:"required,max=255"`
	LastName           string    `json:"lastName" validate:"required,max=255"`
	MiddleName         string    `json:"middleName,omitempty" validator:"max=255"`
	PermanentAddress   string    `json:"permanentAddress,omitempty"`
	Description        string    `json:"description,omitempty"`
	BirthDate          time.Time `json:"birthDate" validate:"required"`
	Username           string    `json:"username" validate:"required,max=255"`
	Email              string    `json:"email" validate:"required,email,max=255"`
	Password           string    `json:"password" validate:"required,min=8,max=255"`
	ContactNumber      string    `json:"contactNumber" validate:"required,max=15"`
	IsEmailVerified    bool      `json:"isEmailVerified"`
	IsContactVerified  bool      `json:"isContactVerified"`
	IsSkipVerification bool      `json:"isSkipVerification"`
	Status             string    `json:"status" validate:"required,oneof=Pending Active Inactive"`
	MediaID            *uint     `json:"mediaID,omitempty"`
	GenderID           *uint     `json:"genderID,omitempty"`
	RoleID             *uint     `json:"roleID,omitempty"`
}

func (m *ModelResource) OwnerToResourceList(owners []*Owner) []*OwnerResource {
	if owners == nil {
		return nil
	}
	var ownerResources []*OwnerResource
	for _, owner := range owners {
		ownerResources = append(ownerResources, m.OwnerToResource(owner))
	}
	return ownerResources
}

func (m *ModelResource) ValidateOwnerRequest(req *OwnerRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return m.helpers.FormatValidationError(err)
	}
	return nil
}

func (r *ModelResource) OwnerGetByEmail(email string) (*Owner, error) {
	var owner Owner
	err := r.OwnerDB.DB.Client.Preload("Media").Where("email = ?", email).First(&owner).Error
	return &owner, err
}

func (r *ModelResource) OwnerGetByUsername(username string) (*Owner, error) {
	var owner Owner
	err := r.OwnerDB.DB.Client.Preload("Media").Where("username = ?", username).First(&owner).Error
	return &owner, err
}

func (r *ModelResource) OwnerGetByContactNumber(contactNumber string) (*Owner, error) {
	var owner Owner
	err := r.OwnerDB.DB.Client.Preload("Media").Where("contact_number = ?", contactNumber).First(&owner).Error
	return &owner, err
}

func (r *ModelResource) OwnerFindByEmailUsernameOrContact(input string) (*Owner, error) {
	switch r.helpers.GetKeyType(input) {
	case "contact":
		return r.OwnerGetByContactNumber(input)
	case "email":
		return r.OwnerGetByEmail(input)
	default:
		return r.OwnerGetByUsername(input)
	}
}
func (r *ModelResource) OwnerUpdatePassword(id uint, password string) error {
	newPassword, err := r.cryptoHelpers.HashPassword(password)
	if err != nil {
		return err
	}
	updated := &Owner{Password: newPassword}
	_, err = r.OwnerDB.UpdateColumns(id, *updated, []string{"Media"})
	return err
}

func (r *ModelResource) OwnerCreate(user *Owner) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}
	if user.Password == "" {
		return errors.New("password cannot be empty")
	}
	hashedPassword, err := r.cryptoHelpers.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return r.OwnerDB.Create(user)
}

func (r *ModelResource) OwnerUpdate(user *Owner, preloads []string) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}
	if user.Password != "" {
		hashedPassword, err := r.cryptoHelpers.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}
	return r.OwnerDB.Update(user, preloads)
}
func (m *ModelResource) OwnerSeeders() error {
	m.logger.Info("Seeding Owner")
	return nil
}
