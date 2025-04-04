package models

import (
	"errors"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/google/uuid"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

type Owner struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Fields
	FirstName          string               `gorm:"type:varchar(255);unsigned" json:"first_name"`
	LastName           string               `gorm:"type:varchar(255);unsigned" json:"last_name"`
	MiddleName         string               `gorm:"type:varchar(255)" json:"middle_name"`
	PermanentAddress   string               `gorm:"type:text" json:"permanent_address"`
	Description        string               `gorm:"type:text" json:"description"`
	BirthDate          time.Time            `gorm:"type:date" json:"birth_date"`
	Username           string               `gorm:"type:varchar(255);unique;unsigned" json:"username"`
	Email              string               `gorm:"type:varchar(255);unique;unsigned" json:"email"`
	Password           string               `gorm:"type:varchar(255);unsigned" json:"password"`
	ContactNumber      string               `gorm:"type:varchar(15);unique;unsigned" json:"contact_number"`
	IsEmailVerified    bool                 `gorm:"default:false" json:"is_email_verified"`
	IsContactVerified  bool                 `gorm:"default:false" json:"is_contact_verified"`
	IsSkipVerification bool                 `gorm:"default:false" json:"is_skip_verification"`
	Status             providers.UserStatus `gorm:"type:varchar(11);default:'Pending'" json:"status"`

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

func (v *Owner) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type OwnerResource struct {
	AccountType string `json:"accountType"`

	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName"`
	FullName   string `json:"fullName"`

	PermanentAddress   string               `json:"permanentAddress"`
	Description        string               `json:"description"`
	BirthDate          time.Time            `json:"birthDate"`
	Username           string               `json:"username"`
	Email              string               `json:"email"`
	Password           string               `json:"password"`
	ContactNumber      string               `json:"contactNumber"`
	IsEmailVerified    bool                 `json:"isEmailVerified"`
	IsContactVerified  bool                 `json:"isContactVerified"`
	IsSkipVerification bool                 `json:"isSkipVerification"`
	Status             providers.UserStatus `json:"status"`
	MediaID            *uuid.UUID           `json:"mediaID"`
	Media              *MediaResource       `json:"media"`
	GenderID           *uuid.UUID           `json:"genderID"`
	Gender             *GenderResource      `json:"gender"`
	RoleID             *uuid.UUID           `json:"roleID"`
	Role               *RoleResource        `json:"role"`
	Footsteps          []*FootstepResource  `json:"footsteps"`
	Companies          []*CompanyResource   `json:"companies"`
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

		FirstName:  owner.FirstName,
		LastName:   owner.LastName,
		MiddleName: owner.MiddleName,
		FullName:   owner.FirstName + " " + owner.MiddleName + " " + owner.LastName,

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

func (m *ModelRepository) OwnerGetByID(id string, preloads ...string) (*Owner, error) {
	repo := NewGenericRepository[Owner](m.db.Client)
	return repo.GetByID(id, preloads...)
}
func (m *ModelRepository) OwnerGetByUsername(username string, preloads ...string) (*Owner, error) {
	repo := NewGenericRepository[Owner](m.db.Client)
	return repo.GetByColumn("username", username, preloads...)
}
func (m *ModelRepository) OwnerGetByEmail(email string, preloads ...string) (*Owner, error) {
	repo := NewGenericRepository[Owner](m.db.Client)
	return repo.GetByColumn("email", email, preloads...)
}
func (m *ModelRepository) OwnerGetByContactNumber(contact_number string, preloads ...string) (*Owner, error) {
	repo := NewGenericRepository[Owner](m.db.Client)
	return repo.GetByColumn("contact_number", contact_number, preloads...)
}
func (m *ModelRepository) OwnerSearch(input string, preloads ...string) (*Owner, error) {
	switch m.helpers.GetKeyType(input) {
	case "id":
		return m.OwnerGetByID(input, preloads...)
	case "contact":
		return m.OwnerGetByContactNumber(input, preloads...)
	case "email":
		return m.OwnerGetByEmail(input, preloads...)
	default:
		return m.OwnerGetByUsername(input, preloads...)
	}
}
func (m *ModelRepository) OwnerCreate(owner *Owner, preloads ...string) (*Owner, error) {
	repo := NewGenericRepository[Owner](m.db.Client)
	newPassword, err := m.cryptoHelpers.HashPassword(owner.Password)
	if err != nil {
		return nil, err
	}
	owner.Password = newPassword
	return repo.Create(owner, preloads...)
}
func (m *ModelRepository) OwnerUpdate(owner *Owner, preloads ...string) (*Owner, error) {
	repo := NewGenericRepository[Owner](m.db.Client)
	return repo.Update(owner, preloads...)
}
func (m *ModelRepository) OwnerUpdateByID(id string, value *Owner, preloads ...string) (*Owner, error) {
	repo := NewGenericRepository[Owner](m.db.Client)
	return repo.UpdateByID(id, value, preloads...)
}
func (m *ModelRepository) OwnerDeleteByID(id string) error {
	repo := NewGenericRepository[Owner](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) OwnerGetAll(preloads ...string) ([]*Owner, error) {
	repo := NewGenericRepository[Owner](m.db.Client)
	return repo.GetAll(preloads...)
}

func (m *ModelRepository) OwnerSignIn(key string, password string, preloads ...string) (*Owner, error) {
	owner, err := m.OwnerSearch(key, preloads...)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, eris.New("user not found")
		}
		return nil, eris.Wrap(err, "failed to search for owner")
	}
	if !m.cryptoHelpers.VerifyPassword(owner.Password, password) {
		return nil, eris.New("invalid credentials")
	}
	return owner, nil
}

func (m *ModelRepository) OwnerChangePassword(ownerID string, currentPassword, newPassword string, preloads ...string) (*Owner, error) {
	owner, err := m.OwnerGetByID(ownerID, preloads...)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, eris.New("owner not found")
		}
		return nil, eris.Wrap(err, "failed to retrieve owner")
	}
	if !m.cryptoHelpers.VerifyPassword(owner.Password, currentPassword) {
		return nil, eris.New("invalid current password")
	}
	hashedPassword, err := m.cryptoHelpers.HashPassword(newPassword)
	if err != nil {
		return nil, eris.Wrap(err, "unable to hash new password")
	}
	owner.Password = hashedPassword
	repo := NewGenericRepository[Owner](m.db.Client)
	updatedOwner, err := repo.Update(owner, preloads...)
	if err != nil {
		return nil, eris.Wrap(err, "failed to update owner with new password")
	}
	return updatedOwner, nil
}

func (m *ModelRepository) OwnerForceChangePassword(ownerID string, newPassword string, preloads ...string) (*Owner, error) {
	owner, err := m.OwnerGetByID(ownerID, preloads...)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, eris.New("owner not found")
		}
		return nil, eris.Wrap(err, "failed to retrieve owner")
	}
	hashedPassword, err := m.cryptoHelpers.HashPassword(newPassword)
	if err != nil {
		return nil, eris.Wrap(err, "unable to hash new password")
	}
	owner.Password = hashedPassword
	repo := NewGenericRepository[Owner](m.db.Client)
	updatedOwner, err := repo.Update(owner, preloads...)
	if err != nil {
		return nil, eris.Wrap(err, "failed to update owner with new password")
	}
	return updatedOwner, nil
}

func (m *ModelRepository) OwnerVerifyPassword(key string, password string) bool {
	admin, err := m.OwnerSearch(key)
	if err != nil {
		return false
	}
	return m.cryptoHelpers.VerifyPassword(admin.Password, password)
}
