package models

import (
	"errors"
	"time"

	"github.com/go-playground/validator"
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
	Media              *MediaResource  `json:"media"`
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

func (m *ModelResource) MemberToResource(member *Member) *MemberResource {
	if member == nil {
		return nil
	}
	return &MemberResource{
		FirstName:          member.FirstName,
		LastName:           member.LastName,
		MiddleName:         member.MiddleName,
		PermanentAddress:   member.PermanentAddress,
		Description:        member.Description,
		BirthDate:          member.BirthDate,
		Username:           member.Username,
		Email:              member.Email,
		Password:           member.Password,
		IsEmailVerified:    member.IsEmailVerified,
		IsContactVerified:  member.IsContactVerified,
		IsSkipVerification: member.IsSkipVerification,
		ContactNumber:      member.ContactNumber,
		Status:             member.Status,
		MediaID:            member.MediaID,
		Media:              m.MediaToResource(member.Media),
		BranchID:           member.BranchID,
		Branch:             m.BranchToResource(member.Branch),
		Longitude:          member.Longitude,
		Latitude:           member.Latitude,
		RoleID:             member.RoleID,
		Role:               m.RoleToResource(member.Role),
		GenderID:           member.GenderID,
		Gender:             m.GenderToResource(member.Gender),
		Footsteps:          m.FootstepToResourceList(member.Footsteps),
	}
}

type MemberRequest struct {
	FirstName          string    `json:"firstName" validate:"required,max=255"`
	LastName           string    `json:"lastName" validate:"required,max=255"`
	MiddleName         string    `json:"middleName,omitempty" validate:"max=255"`
	PermanentAddress   string    `json:"permanentAddress,omitempty"`
	Description        string    `json:"description,omitempty"`
	BirthDate          time.Time `json:"birthDate" validate:"required"`
	Username           string    `json:"username" validate:"required,max=255"`
	Email              string    `json:"email" validate:"required,email,max=255"`
	Password           string    `json:"password" validate:"required,min=8,max=255"`
	IsEmailVerified    bool      `json:"isEmailVerified"`
	IsContactVerified  bool      `json:"isContactVerified"`
	IsSkipVerification bool      `json:"isSkipVerification"`
	ContactNumber      string    `json:"contactNumber" validate:"required,max=255"`
	Status             string    `json:"status" validate:"required,oneof=Pending Active Inactive"`
	MediaID            *uint     `json:"mediaID,omitempty"`
	BranchID           *uint     `json:"branchID,omitempty"`
	Longitude          *float64  `json:"longitude" validate:"omitempty,longitude"`
	Latitude           *float64  `json:"latitude" validate:"omitempty,latitude"`
	RoleID             *uint     `json:"roleID,omitempty"`
	GenderID           *uint     `json:"genderID,omitempty"`
}

func (m *ModelResource) MemberToResourceList(members []*Member) []*MemberResource {
	if members == nil {
		return nil
	}
	var memberResources []*MemberResource
	for _, member := range members {
		memberResources = append(memberResources, m.MemberToResource(member))
	}
	return memberResources
}

func (m *ModelResource) ValidateMemberRequest(req *MemberRequest) error {
	validate := validator.New()
	validate.RegisterValidation("longitude", func(fl validator.FieldLevel) bool {
		lon := fl.Field().Float()
		return lon >= -180 && lon <= 180
	})
	validate.RegisterValidation("latitude", func(fl validator.FieldLevel) bool {
		lat := fl.Field().Float()
		return lat >= -90 && lat <= 90
	})
	err := validate.Struct(req)
	if err != nil {
		return m.helpers.FormatValidationError(err)
	}
	return nil
}

func (r *ModelResource) MemberGetByEmail(email string) (*Member, error) {
	var member Member
	err := r.MemberDB.DB.Client.Preload("Media").Where("email = ?", email).First(&member).Error
	return &member, err
}

func (r *ModelResource) MemberGetByUsername(username string) (*Member, error) {
	var member Member
	err := r.MemberDB.DB.Client.Preload("Media").Where("username = ?", username).First(&member).Error
	return &member, err
}

func (r *ModelResource) MemberGetByContactNumber(contactNumber string) (*Member, error) {
	var member Member
	err := r.MemberDB.DB.Client.Preload("Media").Where("contact_number = ?", contactNumber).First(&member).Error
	return &member, err
}

func (r *ModelResource) MemberFindByEmailUsernameOrContact(input string) (*Member, error) {
	switch r.helpers.GetKeyType(input) {
	case "contact":
		return r.MemberGetByContactNumber(input)
	case "email":
		return r.MemberGetByEmail(input)
	default:
		return r.MemberGetByUsername(input)
	}
}
func (r *ModelResource) MemberUpdatePassword(id uint, password string) error {
	newPassword, err := r.cryptoHelpers.HashPassword(password)
	if err != nil {
		return err
	}
	updated := &Member{Password: newPassword}
	_, err = r.MemberDB.UpdateColumns(id, *updated, []string{"Media"})
	return err
}

func (r *ModelResource) MemberCreate(user *Member) error {
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
	return r.MemberDB.Create(user)
}

func (r *ModelResource) MemberUpdate(user *Member, preloads []string) error {
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
	return r.MemberDB.Update(user, preloads)
}

func (m *ModelResource) MemberSeeders() error {
	m.logger.Info("Seeding Member")
	return nil
}
