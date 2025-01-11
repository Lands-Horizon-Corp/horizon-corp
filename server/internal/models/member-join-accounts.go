package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberJointAccounts struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	MembersProfileID   uuid.UUID      `gorm:"unsigned" json:"members_profile_id"`
	Description        string         `gorm:"type:text" json:"description"`
	FirstName          string         `gorm:"type:varchar(255);unsigned" json:"first_name"`
	LastName           string         `gorm:"type:varchar(255);unsigned" json:"last_name"`
	MiddleName         string         `gorm:"type:varchar(255)" json:"middle_name"`
	FamilyRelationship string         `gorm:"type:varchar(255)" json:"family_relationship"`
	MembersProfile     *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

type MemberJointAccountsResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	MembersProfileID   uuid.UUID              `json:"membersProfileID"`
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

		ID:        account.ID,
		CreatedAt: account.CreatedAt.Format(time.RFC3339),
		UpdatedAt: account.UpdatedAt.Format(time.RFC3339),
		DeletedAt: account.DeletedAt.Time.Format(time.RFC3339),

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
