package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberProfile struct {
	gorm.Model

	Description          string `gorm:"type:text" json:"description"`
	Notes                string `gorm:"type:text" json:"notes"`
	ContactNumber        string `gorm:"type:varchar(255);unsigned" json:"contact_number"`
	OldferenceID         string `gorm:"type:varchar(255)" json:"old_reference_id"`
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

	// Verified By
	VerifiedByEmployeeID *uint     `gorm:"type:bigint;unsigned;index" json:"verified_by_employee_id"`
	VerifiedByEmployee   *Employee `gorm:"foreignKey:VerifiedByEmployeeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"verified_by_employee"`

	// one-to-one Relationships
	MediaID                       *uint `gorm:"type:bigint;unsigned;index" json:"media_id"`
	MemberClassificationID        *uint `gorm:"type:bigint;unsigned;index" json:"member_classification_id"`
	MemberGenderID                *uint `gorm:"type:bigint;unsigned;index" json:"member_gender_id"`
	MemberEducationalAttainmentID *uint `gorm:"type:bigint;unsigned;index" json:"member_educational_attainment_id"`

	MemberClassification *MemberClassification `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_classification"`
	MemberGender         *MemberGender         `gorm:"foreignKey:MemberGenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_gender"`

	MemberCenterID *uint         `gorm:"type:bigint;unsigned;index" json:"member_center_id"`
	MemberCenter   *MemberCenter `gorm:"foreignKey:MemberCenterID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_center"`

	// zero-to-many Relationships
	MemberDescription             []*MemberDescription             `gorm:"foreignKey:MembersProfileID" json:"member_description"`
	MemberRecruits                []*MemberRecruits                `gorm:"foreignKey:MembersProfileID" json:"member_recruits"`
	MemberContactNumberReferences []*MemberContactNumberReferences `gorm:"foreignKey:MembersProfileID" json:"member_contact_number_references"`
	MemberWallets                 []*MemberWallet                  `gorm:"foreignKey:MembersProfileID" json:"member_wallets"`
	MemberIncome                  []*MemberIncome                  `gorm:"foreignKey:MembersProfileID" json:"member_income"`
	MemberExpenses                []*MemberExpenses                `gorm:"foreignKey:MembersProfileID" json:"member_expenses"`
	MemberCloseRemarks            []*MemberCloseRemarks            `gorm:"foreignKey:MembersProfileID" json:"member_close_remarks"`
	MemberJointAccounts           []*MemberJointAccounts           `gorm:"foreignKey:MembersProfileID" json:"member_joint_accounts"`
	MemberRelativeAccounts        []*MemberRelativeAccounts        `gorm:"foreignKey:MembersProfileID" json:"member_relative_accounts"`
	MemberAddress                 []*MemberAddress                 `gorm:"foreignKey:MembersProfileID" json:"member_address"`
	MemberGovernmentBenefits      []*MemberGovernmentBenefits      `gorm:"foreignKey:MembersProfileID" json:"member_government_benefits"`
	MemberMutualFundsHistory      []*MemberMutualFundsHistory      `gorm:"foreignKey:MembersProfileID" json:"member_mutual_funds_history"`
	MemberAssets                  []*MemberAssets                  `gorm:"foreignKey:MembersProfileID" json:"member_assets"`

	BranchID *uint   `gorm:"type:bigint;unsigned;index" json:"branch_id"`
	Branch   *Branch `gorm:"foreignKey:BranchID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"branch"`
}

type MemberProfileResource struct {
	ID                     uint                          `json:"id"`
	CreatedAt              string                        `json:"createdAt"`
	UpdatedAt              string                        `json:"updatedAt"`
	Description            string                        `json:"description"`
	Notes                  string                        `json:"notes"`
	ContactNumber          string                        `json:"contactNumber"`
	OldReferenceID         string                        `json:"oldReferenceID,omitempty"`
	Status                 string                        `json:"status"`
	PassbookNumber         string                        `json:"passbookNumber,omitempty"`
	IsClosed               bool                          `json:"isClosed"`
	Occupation             string                        `json:"occupation,omitempty"`
	BusinessAddress        string                        `json:"businessAddress,omitempty"`
	BusinessContact        string                        `json:"businessContact,omitempty"`
	TinNumber              string                        `json:"tinNumber,omitempty"`
	CivilStatus            string                        `json:"civilStatus"`
	SSSNumber              string                        `json:"sssNumber,omitempty"`
	PagibigNumber          string                        `json:"pagibigNumber,omitempty"`
	PhilhealthNumber       string                        `json:"philhealthNumber,omitempty"`
	IsMutualFundMember     bool                          `json:"isMutualFundMember"`
	IsMicroFinanceMember   bool                          `json:"isMicroFinanceMember"`
	MemberTypeID           *uint                         `json:"memberTypeID,omitempty"`
	MemberType             *MemberTypeResource           `json:"memberType,omitempty"`
	MemberID               *uint                         `json:"memberID,omitempty"`
	Member                 *MemberResource               `json:"member,omitempty"`
	MediaID                *uint                         `json:"mediaID,omitempty"`
	MemberClassificationID *uint                         `json:"memberClassificationID,omitempty"`
	MemberClassification   *MemberClassificationResource `json:"memberClassification,omitempty"`
	MemberGenderID         *uint                         `json:"memberGenderID,omitempty"`
	MemberGender           *MemberGenderResource         `json:"memberGender,omitempty"`

	VerifiedByEmployeeID *uint             `json:"verifiedByEmployeeID,omitempty"`
	VerifiedByEmployee   *EmployeeResource `json:"verifiedByEmployee,omitempty"`

	BranchID       *uint                 `json:"branchID,omitempty"`
	Branch         *BranchResource       `json:"branch,omitempty"`
	MemberCenterID *uint                 `json:"memberCenterID,omitempty"`
	MemberCenter   *MemberCenterResource `json:"memberCenter,omitempty"`

	MemberEducationalAttainmentID *uint `json:"memberEducationalAttainmentID,omitempty"`

	// Zero-to-Many Relationships
	MemberDescriptions            []*MemberDescriptionResource             `json:"memberDescriptions,omitempty"`
	MemberRecruits                []*MemberRecruitsResource                `json:"memberRecruits,omitempty"`
	MemberContactNumberReferences []*MemberContactNumberReferencesResource `json:"memberContactNumberReferences,omitempty"`
	MemberWallets                 []*MemberWalletResource                  `json:"memberWallets,omitempty"`
	MemberIncome                  []*MemberIncomeResource                  `json:"memberIncome,omitempty"`
	MemberExpenses                []*MemberExpensesResource                `json:"memberExpenses,omitempty"`
	MemberCloseRemarks            []*MemberCloseRemarksResource            `json:"memberCloseRemarks,omitempty"`
	MemberJointAccounts           []*MemberJointAccountsResource           `json:"memberJointAccounts,omitempty"`
	MemberRelativeAccounts        []*MemberRelativeAccountsResource        `json:"memberRelativeAccounts,omitempty"`
	MemberAddresses               []*MemberAddressResource                 `json:"memberAddresses,omitempty"`
	MemberGovernmentBenefits      []*MemberGovernmentBenefitsResource      `json:"memberGovernmentBenefits,omitempty"`
	MemberMutualFundsHistory      []*MemberMutualFundsHistoryResource      `json:"memberMutualFundsHistory,omitempty"`
	MemberAssets                  []*MemberAssetsResource                  `json:"memberAssets,omitempty"`
}

func (m *ModelTransformer) MemberProfileToResource(profile *MemberProfile) *MemberProfileResource {
	if profile == nil {
		return nil
	}

	return &MemberProfileResource{
		ID:                            profile.ID,
		CreatedAt:                     profile.CreatedAt.Format(time.RFC3339),
		UpdatedAt:                     profile.UpdatedAt.Format(time.RFC3339),
		Description:                   profile.Description,
		Notes:                         profile.Notes,
		ContactNumber:                 profile.ContactNumber,
		OldReferenceID:                profile.OldferenceID,
		Status:                        profile.Status,
		PassbookNumber:                profile.PassbookNumber,
		IsClosed:                      profile.IsClosed,
		Occupation:                    profile.Occupation,
		BusinessAddress:               profile.BusinessAddress,
		BusinessContact:               profile.BusinessContact,
		TinNumber:                     profile.TinNumber,
		CivilStatus:                   profile.CivilStatus,
		SSSNumber:                     profile.SSSNumber,
		PagibigNumber:                 profile.PagibigNumber,
		PhilhealthNumber:              profile.PhilhealthNumber,
		IsMutualFundMember:            profile.IsMutualFundMember,
		IsMicroFinanceMember:          profile.IsMicroFinanceMember,
		MemberTypeID:                  profile.MemberTypeID,
		MemberType:                    m.MemberTypeToResource(profile.MemberType),
		MemberID:                      profile.MemberID,
		Member:                        m.MemberToResource(profile.Member),
		MediaID:                       profile.MediaID,
		MemberClassificationID:        profile.MemberClassificationID,
		MemberClassification:          m.MemberClassificationToResource(profile.MemberClassification),
		MemberGenderID:                profile.MemberGenderID,
		MemberGender:                  m.MemberGenderToResource(profile.MemberGender),
		MemberCenterID:                profile.MemberCenterID,
		MemberCenter:                  m.MemberCenterToResource(profile.MemberCenter),
		VerifiedByEmployeeID:          profile.VerifiedByEmployeeID,
		VerifiedByEmployee:            m.EmployeeToResource(profile.VerifiedByEmployee),
		MemberEducationalAttainmentID: profile.MemberEducationalAttainmentID,
		BranchID:                      profile.BranchID,
		Branch:                        m.BranchToResource(profile.Branch),

		// Zero-to-Many Relationships
		MemberDescriptions:            m.MemberDescriptionToResourceList(profile.MemberDescription),
		MemberRecruits:                m.MemberRecruitsToResourceList(profile.MemberRecruits),
		MemberContactNumberReferences: m.MemberContactNumberReferencesToResourceList(profile.MemberContactNumberReferences),
		MemberWallets:                 m.MemberWalletToResourceList(profile.MemberWallets),
		MemberIncome:                  m.MemberIncomeToResourceList(profile.MemberIncome),
		MemberExpenses:                m.MemberExpensesToResourceList(profile.MemberExpenses),
		MemberCloseRemarks:            m.MemberCloseRemarksToResourceList(profile.MemberCloseRemarks),
		MemberJointAccounts:           m.MemberJointAccountsToResourceList(profile.MemberJointAccounts),
		MemberRelativeAccounts:        m.MemberRelativeAccountsToResourceList(profile.MemberRelativeAccounts),
		MemberAddresses:               m.MemberAddressToResourceList(profile.MemberAddress),
		MemberGovernmentBenefits:      m.MemberGovernmentBenefitsToResourceList(profile.MemberGovernmentBenefits),
		MemberMutualFundsHistory:      m.MemberMutualFundsHistoryToResourceList(profile.MemberMutualFundsHistory),
		MemberAssets:                  m.MemberAssetsToResourceList(profile.MemberAssets),
	}
}

func (m *ModelTransformer) MemberProfileToResourceList(profileList []*MemberProfile) []*MemberProfileResource {
	if profileList == nil {
		return nil
	}

	var profileResources []*MemberProfileResource
	for _, profile := range profileList {
		profileResources = append(profileResources, m.MemberProfileToResource(profile))
	}
	return profileResources
}
