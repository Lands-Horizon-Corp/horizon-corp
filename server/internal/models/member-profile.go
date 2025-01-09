package models

import "gorm.io/gorm"

type MemberProfile struct {
	gorm.Model

	Description          string `gorm:"type:text" json:"description"`
	Notes                string `gorm:"type:text" json:"notes"`
	ContactNumber        string `gorm:"type:varchar(255);not null" json:"contact_number"`
	OldReferenceID       string `gorm:"type:varchar(255)" json:"old_reference_id"`
	Status               string `gorm:"type:varchar(50);default:'pending'" json:"status"`
	PassbookNumber       string `gorm:"type:varchar(255)" json:"passbook_number"`
	IsClosed             bool   `gorm:"default:false" json:"is_closed"`
	Occupation           string `gorm:"type:varchar(255)" json:"occupation"`
	BusinessAddress      string `gorm:"type:text" json:"business_address"`
	BusinessContact      string `gorm:"type:varchar(255)" json:"business_contact"`
	TinNumber            string `gorm:"type:varchar(255)" json:"tin_number"`
	CivilStatus          string `gorm:"type:varchar(50);default:'single'" json:"civil_status"`
	SSSNumber            string `gorm:"type:varchar(255)" json:"sss_number"`
	PagibigNumber        string `gorm:"type:varchar(255)" json:"pagibig_number"`
	PhilhealthNumber     string `gorm:"type:varchar(255)" json:"philhealth_number"`
	IsMutualFundMember   bool   `gorm:"default:false" json:"is_mutual_fund_member"`
	IsMicroFinanceMember bool   `gorm:"default:false" json:"is_micro_finance_member"`

	// Relationships
	MemberTypeID *uint       `gorm:"type:bigint;unsigned;index" json:"member_type_id"`
	MemberType   *MemberType `gorm:"foreignKey:MemberTypeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_type"`
	MemberID     *uint       `gorm:"type:bigint;unsigned;index" json:"member_id"`
	Member       *Member     `gorm:"foreignKey:MemberID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member"`

	// Many-to-One Relationships
	MediaID                       *uint `gorm:"type:bigint;unsigned;index" json:"media_id"`
	MemberClassificationID        *uint `gorm:"type:bigint;unsigned;index" json:"member_classification_id"`
	MemberGenderID                *uint `gorm:"type:bigint;unsigned;index" json:"member_gender_id"`
	MemberEducationalAttainmentID *uint `gorm:"type:bigint;unsigned;index" json:"member_educational_attainment_id"`

	MemberClassification *MemberClassification `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_classification"`
	MemberGender         *MemberGender         `gorm:"foreignKey:MemberGenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_gender"`
}

type MemberProfileResource struct {
	ID                            uint                          `json:"id"`
	CreatedAt                     string                        `json:"createdAt"`
	UpdatedAt                     string                        `json:"updatedAt"`
	Description                   string                        `json:"description"`
	Notes                         string                        `json:"notes"`
	ContactNumber                 string                        `json:"contactNumber"`
	OldReferenceID                string                        `json:"oldReferenceID,omitempty"`
	Status                        string                        `json:"status"`
	PassbookNumber                string                        `json:"passbookNumber,omitempty"`
	IsClosed                      bool                          `json:"isClosed"`
	Occupation                    string                        `json:"occupation,omitempty"`
	BusinessAddress               string                        `json:"businessAddress,omitempty"`
	BusinessContact               string                        `json:"businessContact,omitempty"`
	TinNumber                     string                        `json:"tinNumber,omitempty"`
	CivilStatus                   string                        `json:"civilStatus"`
	SSSNumber                     string                        `json:"sssNumber,omitempty"`
	PagibigNumber                 string                        `json:"pagibigNumber,omitempty"`
	PhilhealthNumber              string                        `json:"philhealthNumber,omitempty"`
	IsMutualFundMember            bool                          `json:"isMutualFundMember"`
	IsMicroFinanceMember          bool                          `json:"isMicroFinanceMember"`
	MemberTypeID                  *uint                         `json:"memberTypeID,omitempty"`
	MemberType                    *MemberTypeResource           `json:"memberType,omitempty"`
	MemberID                      *uint                         `json:"memberID,omitempty"`
	Member                        *MemberResource               `json:"member,omitempty"`
	MediaID                       *uint                         `json:"mediaID,omitempty"`
	MemberClassificationID        *uint                         `json:"memberClassificationID,omitempty"`
	MemberClassification          *MemberClassificationResource `json:"memberClassification,omitempty"`
	MemberGenderID                *uint                         `json:"memberGenderID,omitempty"`
	MemberGender                  *MemberGenderResource         `json:"memberGender,omitempty"`
	MemberEducationalAttainmentID *uint                         `json:"memberEducationalAttainmentID,omitempty"`
}
