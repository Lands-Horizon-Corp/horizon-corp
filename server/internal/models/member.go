package models

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model

	FirstName         string    `gorm:"type:varchar(255);unsigned;index" json:"first_name"`
	LastName          string    `gorm:"type:varchar(255);unsigned;index" json:"last_name"`
	MiddleName        string    `gorm:"type:varchar(255)" json:"middle_name"`
	PermanentAddress  string    `gorm:"type:text" json:"permanent_address"`
	Description       string    `gorm:"type:text" json:"description"`
	BirthDate         time.Time `gorm:"type:date;unsigned" json:"birth_date"`
	Username          string    `gorm:"type:varchar(255);unique;unsigned" json:"username"`
	Email             string    `gorm:"type:varchar(255);unique;unsigned" json:"email"`
	Password          string    `gorm:"type:varchar(255);unsigned" json:"password"`
	IsEmailVerified   bool      `gorm:"default:false" json:"is_email_verified"`
	IsContactVerified bool      `gorm:"default:false" json:"is_contact_verified"`
	ContactNumber     string    `gorm:"type:varchar(255);unique;unsigned" json:"contact_number"`
	Status            string    `gorm:"type:varchar(50);default:'Pending'" json:"status"`

	// Nullable Foreign Keys
	MediaID  *uint   `gorm:"type:bigint;unsigned;index" json:"media_id"`
	Media    *Media  `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media"`
	BranchID *uint   `gorm:"type:bigint;unsigned;index" json:"branch_id"`
	Branch   *Branch `gorm:"foreignKey:BranchID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"branch"`
	RoleID   *uint   `gorm:"type:bigint;unsigned;index" json:"role_id"`
	Role     *Role   `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`

	GenderID *uint   `gorm:"type:bigint;unsigned;index" json:"gender_id"`
	Gender   *Gender `gorm:"foreignKey:GenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"gender"`

	// Location
	Longitude *float64 `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude  *float64 `gorm:"type:decimal(10,7)" json:"latitude"`

	// Relationships
	Footsteps     []*Footstep    `gorm:"foreignKey:MemberID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"footsteps,omitempty"`
	MemberProfile *MemberProfile `gorm:"foreignKey:MemberID" json:"member_profile"`
}

type MemberResource struct {
	ID                 uint                   `json:"id"`
	CreatedAt          string                 `json:"createdAt"`
	UpdatedAt          string                 `json:"updatedAt"`
	FirstName          string                 `json:"firstName"`
	LastName           string                 `json:"lastName"`
	MiddleName         string                 `json:"middleName,omitempty"`
	PermanentAddress   string                 `json:"permanentAddress,omitempty"`
	Description        string                 `json:"description,omitempty"`
	BirthDate          time.Time              `json:"birthDate"`
	Username           string                 `json:"username"`
	Email              string                 `json:"email"`
	IsEmailVerified    bool                   `json:"isEmailVerified"`
	IsContactVerified  bool                   `json:"isContactVerified"`
	IsSkipVerification bool                   `json:"isSkipVerification"`
	ContactNumber      string                 `json:"contactNumber"`
	Status             string                 `json:"status"`
	Media              *MediaResource         `json:"media,omitempty"`
	Branch             *BranchResource        `json:"branch,omitempty"`
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
		ID:                member.ID,
		CreatedAt:         member.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         member.UpdatedAt.Format(time.RFC3339),
		FirstName:         member.FirstName,
		LastName:          member.LastName,
		MiddleName:        member.MiddleName,
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
		Branch:            m.BranchToResource(member.Branch),
		Longitude:         member.Longitude,
		Latitude:          member.Latitude,
		Role:              m.RoleToResource(member.Role),
		Gender:            m.GenderToResource(member.Gender),
		Footsteps:         m.FootstepToResourceList(member.Footsteps),
		MemberProfile:     m.MemberProfileToResource(member.MemberProfile),
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
