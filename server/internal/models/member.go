package models

import (
	"errors"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/google/uuid"
	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

type Member struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	FirstName          string               `gorm:"type:varchar(255);unsigned;index" json:"first_name"`
	LastName           string               `gorm:"type:varchar(255);unsigned;index" json:"last_name"`
	MiddleName         string               `gorm:"type:varchar(255)" json:"middle_name"`
	PermanentAddress   string               `gorm:"type:text" json:"permanent_address"`
	Description        string               `gorm:"type:text" json:"description"`
	BirthDate          time.Time            `gorm:"type:date;unsigned" json:"birth_date"`
	Username           string               `gorm:"type:varchar(255);unique;unsigned" json:"username"`
	Email              string               `gorm:"type:varchar(255);unique;unsigned" json:"email"`
	Password           string               `gorm:"type:varchar(255);unsigned" json:"password"`
	IsEmailVerified    bool                 `gorm:"default:false" json:"is_email_verified"`
	IsContactVerified  bool                 `gorm:"default:false" json:"is_contact_verified"`
	IsSkipVerification bool                 `gorm:"default:false" json:"is_skip_verification"`
	ContactNumber      string               `gorm:"type:varchar(255);unique;unsigned" json:"contact_number"`
	Status             providers.UserStatus `gorm:"type:varchar(50);default:'Pending'" json:"status"`

	// Nullable Foreign Keys
	MediaID *uuid.UUID `gorm:"type:bigint;unsigned;index" json:"media_id"`
	Media   *Media     `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media"`

	RoleID *uuid.UUID `gorm:"type:bigint;unsigned;index" json:"role_id"`
	Role   *Role      `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`

	GenderID *uuid.UUID `gorm:"type:bigint;unsigned;index" json:"gender_id"`
	Gender   *Gender    `gorm:"foreignKey:GenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"gender"`

	// Location
	Longitude *float64 `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude  *float64 `gorm:"type:decimal(10,7)" json:"latitude"`

	// Relationships
	Footsteps     []*Footstep    `gorm:"foreignKey:MemberID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"footsteps,omitempty"`
	MemberProfile *MemberProfile `gorm:"foreignKey:MemberID" json:"member_profile"`
}

func (v *Member) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type MemberResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName,omitempty"`
	FullName   string `json:"fullName"`

	PermanentAddress   string                 `json:"permanentAddress,omitempty"`
	Description        string                 `json:"description,omitempty"`
	BirthDate          time.Time              `json:"birthDate"`
	Username           string                 `json:"username"`
	Email              string                 `json:"email"`
	IsEmailVerified    bool                   `json:"isEmailVerified"`
	IsContactVerified  bool                   `json:"isContactVerified"`
	IsSkipVerification bool                   `json:"isSkipVerification"`
	ContactNumber      string                 `json:"contactNumber"`
	Status             providers.UserStatus   `json:"status"`
	Media              *MediaResource         `json:"media,omitempty"`
	Longitude          *float64               `json:"longitude,omitempty"`
	Latitude           *float64               `json:"latitude,omitempty"`
	Role               *RoleResource          `json:"role,omitempty"`
	Gender             *GenderResource        `json:"gender,omitempty"`
	Footsteps          []*FootstepResource    `json:"footsteps,omitempty"`
	MemberProfile      *MemberProfileResource `json:"memberProfile,omitempty"`
}

func (m *ModelTransformer) MemberToResource(member *Member) *MemberResource {
	if member == nil {
		return nil
	}

	return &MemberResource{

		ID:        member.ID,
		CreatedAt: member.CreatedAt.Format(time.RFC3339),
		UpdatedAt: member.UpdatedAt.Format(time.RFC3339),
		DeletedAt: member.DeletedAt.Time.Format(time.RFC3339),

		FirstName:  member.FirstName,
		LastName:   member.LastName,
		MiddleName: member.MiddleName,
		FullName:   member.FirstName + " " + member.MiddleName + " " + member.LastName,

		PermanentAddress:  member.PermanentAddress,
		Description:       member.Description,
		BirthDate:         member.BirthDate,
		Username:          member.Username,
		Email:             member.Email,
		IsEmailVerified:   member.IsEmailVerified,
		IsContactVerified: member.IsContactVerified,
		ContactNumber:     member.ContactNumber,
		Status:            member.Status,
		Media:             m.MediaToResource(member.Media),

		Longitude:     member.Longitude,
		Latitude:      member.Latitude,
		Role:          m.RoleToResource(member.Role),
		Gender:        m.GenderToResource(member.Gender),
		Footsteps:     m.FootstepToResourceList(member.Footsteps),
		MemberProfile: m.MemberProfileToResource(member.MemberProfile),
	}
}
func (m *ModelTransformer) MemberToResourceList(memberList []*Member) []*MemberResource {
	if memberList == nil {
		return nil
	}

	var memberResources []*MemberResource
	for _, member := range memberList {
		memberResources = append(memberResources, m.MemberToResource(member))
	}
	return memberResources
}

func (m *ModelRepository) MemberGetByID(id string, preloads ...string) (*Member, error) {
	repo := NewGenericRepository[Member](m.db.Client)
	return repo.GetByID(id, preloads...)
}
func (m *ModelRepository) MemberGetByUsername(username string, preloads ...string) (*Member, error) {
	repo := NewGenericRepository[Member](m.db.Client)
	return repo.GetByColumn("username", username, preloads...)
}
func (m *ModelRepository) MemberGetByEmail(email string, preloads ...string) (*Member, error) {
	repo := NewGenericRepository[Member](m.db.Client)
	return repo.GetByColumn("email", email, preloads...)
}
func (m *ModelRepository) MemberGetByContactNumber(contact_number string, preloads ...string) (*Member, error) {
	repo := NewGenericRepository[Member](m.db.Client)
	return repo.GetByColumn("contact_number", contact_number, preloads...)
}
func (m *ModelRepository) MemberSearch(input string, preloads ...string) (*Member, error) {
	switch m.helpers.GetKeyType(input) {
	case "id":
		return m.MemberGetByID(input, preloads...)
	case "contact":
		return m.MemberGetByContactNumber(input, preloads...)
	case "email":
		return m.MemberGetByEmail(input, preloads...)
	default:
		return m.MemberGetByUsername(input, preloads...)
	}
}
func (m *ModelRepository) MemberCreate(member *Member, preloads ...string) (*Member, error) {
	repo := NewGenericRepository[Member](m.db.Client)
	newPassword, err := m.cryptoHelpers.HashPassword(member.Password)
	if err != nil {
		return nil, err
	}
	member.Password = newPassword
	member.Status = providers.VerifiedStatus
	return repo.Create(member, preloads...)
}
func (m *ModelRepository) MemberUpdate(member *Member, preloads ...string) (*Member, error) {
	repo := NewGenericRepository[Member](m.db.Client)
	return repo.Update(member, preloads...)
}
func (m *ModelRepository) MemberUpdateByID(id string, value *Member, preloads ...string) (*Member, error) {
	repo := NewGenericRepository[Member](m.db.Client)
	return repo.UpdateByID(id, value, preloads...)
}
func (m *ModelRepository) MemberDeleteByID(id string) error {
	repo := NewGenericRepository[Member](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) MemberGetAll(preloads ...string) ([]*Member, error) {
	repo := NewGenericRepository[Member](m.db.Client)
	return repo.GetAll(preloads...)
}

func (m *ModelRepository) MemberSignIn(key, password string, preload ...string) (*Member, error) {
	member, err := m.MemberSearch(key)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, eris.New("user not found")
		}
		return nil, eris.Wrap(err, "failed to search for member")
	}
	if !m.cryptoHelpers.VerifyPassword(member.Password, password) {
		return nil, eris.New("invalid credentials")
	}
	return member, nil
}

func (m *ModelRepository) MemberChangePassword(memberID string, currentPassword, newPassword string, preloads ...string) (*Member, error) {
	member, err := m.MemberGetByID(memberID, preloads...)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, eris.New("member not found")
		}
		return nil, eris.Wrap(err, "failed to retrieve member")
	}
	if !m.cryptoHelpers.VerifyPassword(member.Password, currentPassword) {
		return nil, eris.New("invalid current password")
	}
	hashedPassword, err := m.cryptoHelpers.HashPassword(newPassword)
	if err != nil {
		return nil, eris.Wrap(err, "unable to hash new password")
	}
	member.Password = hashedPassword
	repo := NewGenericRepository[Member](m.db.Client)
	updatedMember, err := repo.Update(member, preloads...)
	if err != nil {
		return nil, eris.Wrap(err, "failed to update member with new password")
	}
	return updatedMember, nil
}

func (m *ModelRepository) MemberForceChangePassword(memberID string, newPassword string, preloads ...string) (*Member, error) {
	member, err := m.MemberGetByID(memberID, preloads...)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, eris.New("member not found")
		}
		return nil, eris.Wrap(err, "failed to retrieve member")
	}
	hashedPassword, err := m.cryptoHelpers.HashPassword(newPassword)
	if err != nil {
		return nil, eris.Wrap(err, "unable to hash new password")
	}
	member.Password = hashedPassword
	repo := NewGenericRepository[Member](m.db.Client)
	updatedMember, err := repo.Update(member, preloads...)
	if err != nil {
		return nil, eris.Wrap(err, "failed to update member with new password")
	}
	return updatedMember, nil
}
