package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberJointAccounts struct {
	gorm.Model
	MembersProfileID   uint           `gorm:"not null" json:"members_profile_id"`
	Description        string         `gorm:"type:text" json:"description"`
	FirstName          string         `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName           string         `gorm:"type:varchar(255);not null" json:"last_name"`
	MiddleName         string         `gorm:"type:varchar(255)" json:"middle_name"`
	FamilyRelationship string         `gorm:"type:varchar(255)" json:"family_relationship"`
	MembersProfile     *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

type MemberJointAccountsResource struct {
	ID                 uint                   `json:"id"`
	CreatedAt          string                 `json:"createdAt"`
	UpdatedAt          string                 `json:"updatedAt"`
	MembersProfileID   uint                   `json:"membersProfileID"`
	Description        string                 `json:"description"`
	FirstName          string                 `json:"firstName"`
	LastName           string                 `json:"lastName"`
	MiddleName         string                 `json:"middleName,omitempty"`
	FamilyRelationship string                 `json:"familyRelationship,omitempty"`
	MembersProfile     *MemberProfileResource `json:"membersProfile,omitempty"`
}

func (m *ModelTransformer) MemberJointAccountsToResource(account *MemberJointAccounts) *MemberJointAccountsResource {
	if account == nil {
		return nil
	}

	return &MemberJointAccountsResource{
		ID:                 account.ID,
		CreatedAt:          account.CreatedAt.Format(time.RFC3339),
		UpdatedAt:          account.UpdatedAt.Format(time.RFC3339),
		MembersProfileID:   account.MembersProfileID,
		Description:        account.Description,
		FirstName:          account.FirstName,
		LastName:           account.LastName,
		MiddleName:         account.MiddleName,
		FamilyRelationship: account.FamilyRelationship,
		MembersProfile:     m.MemberProfileToResource(account.MembersProfile),
	}
}

func (m *ModelTransformer) MemberJointAccountsToResourceList(accountList []*MemberJointAccounts) []*MemberJointAccountsResource {
	if accountList == nil {
		return nil
	}

	var accountResources []*MemberJointAccountsResource
	for _, account := range accountList {
		accountResources = append(accountResources, m.MemberJointAccountsToResource(account))
	}
	return accountResources
}
