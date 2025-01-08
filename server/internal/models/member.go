package models

import (
	"time"

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

type MemberProfile struct {
	gorm.Model

	Description          string `gorm:"type:text" json:"description"`
	Notes                string `gorm:"type:text" json:"notes"`
	ContactNumber        string `gorm:"type:varchar(255);not null" json:"contact_number"`
	OldReferenceID       string `gorm:"type:varchar(255)" json:"old_reference_id"`
	Status               string `gorm:"type:enum('pending', 'cancelled', 'done');default:'pending'" json:"status"`
	PassbookNumber       string `gorm:"type:varchar(255)" json:"passbook_number"`
	IsClosed             bool   `gorm:"default:false" json:"is_closed"`
	Occupation           string `gorm:"type:varchar(255)" json:"occupation"`
	BusinessAddress      string `gorm:"type:text" json:"business_address"`
	BusinessContact      string `gorm:"type:varchar(255)" json:"business_contact"`
	TinNumber            string `gorm:"type:varchar(255)" json:"tin_number"`
	CivilStatus          string `gorm:"type:enum('single', 'married', 'widow');default:'single'" json:"civil_status"`
	SSSNumber            string `gorm:"type:varchar(255)" json:"sss_number"`
	PagibigNumber        string `gorm:"type:varchar(255)" json:"pagibig_number"`
	PhilhealthNumber     string `gorm:"type:varchar(255)" json:"philhealth_number"`
	IsMutualFundMember   bool   `gorm:"default:false" json:"is_mutual_fund_member"`
	IsMicroFinanceMember bool   `gorm:"default:false" json:"is_micro_finance_member"`
	IsDosriMember        bool   `gorm:"default:false" json:"is_dosri_member"`

	// Relationships
	MemberTypeID *uint       `gorm:"type:bigint;unsigned" json:"member_type_id"`
	MemberType   *MemberType `gorm:"foreignKey:MemberTypeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_type,omitempty"`

	MemberID *uint   `gorm:"type:bigint;unsigned" json:"member_id"`
	Member   *Member `gorm:"foreignKey:MemberID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member,omitempty"`

	MediaID *uint  `gorm:"type:bigint;unsigned" json:"media_id"`
	Media   *Media `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media,omitempty"`

	MemberClassificationID *uint                 `gorm:"type:bigint;unsigned" json:"member_classification_id"`
	MemberClassification   *MemberClassification `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_classification,omitempty"`

	MemberGenderID *uint         `gorm:"type:bigint;unsigned" json:"member_gender_id"`
	MemberGender   *MemberGender `gorm:"foreignKey:MemberGenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_gender,omitempty"`

	MemberEducationalAttainmentID *uint                        `gorm:"type:bigint;unsigned" json:"member_educational_attainment_id"`
	MemberEducationalAttainment   *MemberEducationalAttainment `gorm:"foreignKey:MemberEducationalAttainmentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_educational_attainment,omitempty"`

	MemberGroupID *uint        `gorm:"type:bigint;unsigned" json:"member_group_id"`
	MemberGroup   *MemberGroup `gorm:"foreignKey:MemberGroupID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_group,omitempty"`

	MemberCenterID *uint         `gorm:"type:bigint;unsigned" json:"member_center_id"`
	MemberCenter   *MemberCenter `gorm:"foreignKey:MemberCenterID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_center,omitempty"`

	MemberOccupationID *uint             `gorm:"type:bigint;unsigned" json:"member_occupation_id"`
	MemberOccupation   *MemberOccupation `gorm:"foreignKey:MemberOccupationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_occupation,omitempty"`
}

type MemberClassification struct {
	gorm.Model
	Name        string                         `gorm:"size:255;not null"`
	Description string                         `gorm:"size:500"`
	History     []*MemberClassificationHistory `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

type MemberGender struct {
	gorm.Model
	Name        string                 `gorm:"size:255;not null"`
	Description string                 `gorm:"size:500"`
	History     []*MemberGenderHistory `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

type MemberCenter struct {
	gorm.Model
	Name        string                 `gorm:"size:255;not null"`
	Description string                 `gorm:"size:500"`
	History     []*MemberCenterHistory `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

type MemberGroup struct {
	gorm.Model
	Name        string                `gorm:"size:255;not null"`
	Description string                `gorm:"size:500"`
	History     []*MemberGroupHistory `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

type MemberEducationalAttainment struct {
	gorm.Model
	Name        string                                `gorm:"size:255;not null"`
	Description string                                `gorm:"size:500"`
	History     []*MemberEducationalAttainmentHistory `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

type MemberOccupation struct {
	gorm.Model
	Name        string                     `gorm:"size:255;not null"`
	Description string                     `gorm:"size:500"`
	History     []*MemberOccupationHistory `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

type MemberType struct {
	gorm.Model
	Name        string               `gorm:"size:255;not null"`
	Description string               `gorm:"size:500"`
	Prefix      string               `gorm:"size:100"`
	History     []*MemberTypeHistory `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

// History Models
type MemberClassificationHistory struct {
	gorm.Model
	MemberClassificationID uint      `gorm:"not null"`
	MemberID               uint      `gorm:"not null"`
	CreatedAt              time.Time `gorm:"autoCreateTime"`
}

type MemberGenderHistory struct {
	gorm.Model
	MemberClassificationID uint      `gorm:"not null"`
	MemberID               uint      `gorm:"not null"`
	CreatedAt              time.Time `gorm:"autoCreateTime"`
}

type MemberCenterHistory struct {
	gorm.Model
	MemberClassificationID uint      `gorm:"not null"`
	MemberID               uint      `gorm:"not null"`
	CreatedAt              time.Time `gorm:"autoCreateTime"`
}

type MemberGroupHistory struct {
	gorm.Model
	MemberClassificationID uint      `gorm:"not null"`
	MemberID               uint      `gorm:"not null"`
	CreatedAt              time.Time `gorm:"autoCreateTime"`
}

type MemberEducationalAttainmentHistory struct {
	gorm.Model
	MemberClassificationID uint      `gorm:"not null"`
	MemberID               uint      `gorm:"not null"`
	CreatedAt              time.Time `gorm:"autoCreateTime"`
}

type MemberOccupationHistory struct {
	gorm.Model
	MemberClassificationID uint      `gorm:"not null"`
	MemberID               uint      `gorm:"not null"`
	CreatedAt              time.Time `gorm:"autoCreateTime"`
}

type MemberTypeHistory struct {
	gorm.Model
	MemberClassificationID uint      `gorm:"not null"`
	MemberID               uint      `gorm:"not null"`
	CreatedAt              time.Time `gorm:"autoCreateTime"`
}

// MemberDescription
type MemberDescription struct {
	gorm.Model
	MembersProfileID uint      `gorm:"not null" json:"members_profile_id"`
	Date             time.Time `gorm:"type:date" json:"date"`
	Description      string    `gorm:"type:text" json:"description"`
	Name             string    `gorm:"type:varchar(255);not null" json:"name"`

	MembersProfile *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

// MemberRecruits
type MemberRecruits struct {
	gorm.Model
	MembersProfileID          uint      `gorm:"not null" json:"members_profile_id"`
	MembersProfileRecruitedID uint      `gorm:"not null" json:"members_profile_recruited_id"`
	DateRecruited             time.Time `gorm:"type:date" json:"date_recruited"`
	Description               string    `gorm:"type:text" json:"description"`
	Name                      string    `gorm:"type:varchar(255);not null" json:"name"`

	MembersProfile          *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
	MembersProfileRecruited *MemberProfile `gorm:"foreignKey:MembersProfileRecruitedID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile_recruited"`
}

// ContactNumberReferences
type ContactNumberReferences struct {
	gorm.Model
	Name          string `gorm:"type:varchar(255);not null" json:"name"`
	Description   string `gorm:"type:text" json:"description"`
	ContactNumber string `gorm:"type:varchar(255);not null" json:"contact_number"`
}

type MemberWallet struct {
	gorm.Model
	MembersProfileID uint      `gorm:"not null" json:"members_profile_id"`
	Debit            float64   `gorm:"type:decimal(12,2);default:0" json:"debit"`  // Increased precision
	Credit           float64   `gorm:"type:decimal(12,2);default:0" json:"credit"` // Increased precision
	Date             time.Time `gorm:"type:date" json:"date"`
	Description      string    `gorm:"type:text" json:"description"`

	MembersProfile *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

// MemberIncome
type MemberIncome struct {
	gorm.Model
	MembersProfileID uint      `gorm:"not null" json:"members_profile_id"`
	Name             string    `gorm:"type:varchar(255);not null" json:"name"`
	Amount           float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	Date             time.Time `gorm:"type:date" json:"date"`
	Description      string    `gorm:"type:text" json:"description"`

	MembersProfile *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

// MemberExpenses
type MemberExpenses struct {
	gorm.Model
	MembersProfileID uint      `gorm:"not null" json:"members_profile_id"`
	Name             string    `gorm:"type:varchar(255);not null" json:"name"`
	Amount           float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	Date             time.Time `gorm:"type:date" json:"date"`
	Description      string    `gorm:"type:text" json:"description"`

	MembersProfile *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

// MemberCloseRemarks
type MemberCloseRemarks struct {
	gorm.Model
	MembersProfileID uint   `gorm:"not null" json:"members_profile_id"`
	Description      string `gorm:"type:text" json:"description"`

	MembersProfile *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

// MemberJointAccounts
type MemberJointAccounts struct {
	gorm.Model
	MembersProfileID   uint   `gorm:"not null" json:"members_profile_id"`
	Description        string `gorm:"type:text" json:"description"`
	FirstName          string `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName           string `gorm:"type:varchar(255);not null" json:"last_name"`
	MiddleName         string `gorm:"type:varchar(255)" json:"middle_name"`
	FamilyRelationship string `gorm:"type:varchar(255)" json:"family_relationship"`

	MembersProfile *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

// MemberRelativeAccounts
type MemberRelativeAccounts struct {
	gorm.Model
	MemberID           uint   `gorm:"not null" json:"member_id"`
	RelativeMemberID   uint   `gorm:"not null" json:"relative_member_id"`
	FamilyRelationship string `gorm:"type:varchar(255)" json:"family_relationship"`
	Description        string `gorm:"type:text" json:"description"`

	Member         *MemberProfile `gorm:"foreignKey:MemberID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member"`
	RelativeMember *MemberProfile `gorm:"foreignKey:RelativeMemberID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"relative_member"`
}

// MemberAddress
type MemberAddress struct {
	gorm.Model
	MembersProfileID uint   `gorm:"not null" json:"members_profile_id"`
	PostalCode       string `gorm:"type:varchar(20)" json:"postal_code"`
	Province         string `gorm:"type:varchar(255)" json:"province"`
	City             string `gorm:"type:varchar(255)" json:"city"`
	Barangay         string `gorm:"type:varchar(255)" json:"barangay"`
	Region           string `gorm:"type:varchar(255)" json:"region"`
	Label            string `gorm:"type:enum('work', 'home', 'province', 'business');default:'home'" json:"label"`

	MembersProfile *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

// MemberGovernmentBenefits
type MemberGovernmentBenefits struct {
	gorm.Model
	MembersProfileID uint    `gorm:"not null" json:"members_profile_id"`
	Country          string  `gorm:"type:varchar(255);not null" json:"country"`
	Name             string  `gorm:"type:varchar(255);not null" json:"name"`
	Description      string  `gorm:"type:text" json:"description"`
	Value            float64 `gorm:"type:decimal(10,2)" json:"value"`
	FrontMediaID     *uint   `gorm:"type:bigint;unsigned" json:"front_media_id"`
	BackMediaID      *uint   `gorm:"type:bigint;unsigned" json:"back_media_id"`

	MembersProfile *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
	FrontMedia     *Media         `gorm:"foreignKey:FrontMediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"front_media,omitempty"`
	BackMedia      *Media         `gorm:"foreignKey:BackMediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"back_media,omitempty"`
}

// MutualFundsHistory
type MutualFundsHistory struct {
	gorm.Model
	MembersProfileID uint    `gorm:"not null" json:"members_profile_id"`
	Description      string  `gorm:"type:text" json:"description"`
	Amount           float64 `gorm:"type:decimal(10,2);not null" json:"amount"`

	MembersProfile *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

// MemberAssets
type MemberAssets struct {
	gorm.Model
	MembersProfileID uint      `gorm:"not null" json:"members_profile_id"`
	EntryDate        time.Time `gorm:"type:date;not null" json:"entry_date"`
	Description      string    `gorm:"type:text" json:"description"`
	Name             string    `gorm:"type:varchar(255);not null" json:"name"`

	MembersProfile *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}
