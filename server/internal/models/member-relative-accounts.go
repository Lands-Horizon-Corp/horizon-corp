package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberRelativeAccounts struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	MembersProfileID             uuid.UUID      `gorm:"unsigned" json:"members_profile_id"`
	RelativeProfileMemberID      uuid.UUID      `gorm:"unsigned" json:"relative_member_id"`
	FamilyRelationship           string         `gorm:"type:varchar(255)" json:"family_relationship"`
	Description                  string         `gorm:"type:text" json:"description"`
	MemberProfile                *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	RelativeProfileMemberProfile *MemberProfile `gorm:"foreignKey:RelativeProfileMemberID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"relative_member_profile"`
}

func (v *MemberRelativeAccounts) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type MemberRelativeAccountsResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	MembersProfileID             uuid.UUID              `json:"membersProfileID"`
	RelativeProfileMemberID      uuid.UUID              `json:"relativeProfileMemberID"`
	FamilyRelationship           string                 `json:"familyRelationship"`
	Description                  string                 `json:"description"`
	MemberProfile                *MemberProfileResource `json:"memberProfile,omitempty"`
	RelativeProfileMemberProfile *MemberProfileResource `json:"relativeProfileMemberProfile,omitempty"`
}

func (m *ModelTransformer) MemberRelativeAccountsToResource(account *MemberRelativeAccounts) *MemberRelativeAccountsResource {
	if account == nil {
		return nil
	}

	return &MemberRelativeAccountsResource{

		ID:        account.ID,
		CreatedAt: account.CreatedAt.Format(time.RFC3339),
		UpdatedAt: account.UpdatedAt.Format(time.RFC3339),
		DeletedAt: account.DeletedAt.Time.Format(time.RFC3339),

		MembersProfileID:             account.MembersProfileID,
		RelativeProfileMemberID:      account.RelativeProfileMemberID,
		FamilyRelationship:           account.FamilyRelationship,
		Description:                  account.Description,
		MemberProfile:                m.MemberProfileToResource(account.MemberProfile),
		RelativeProfileMemberProfile: m.MemberProfileToResource(account.RelativeProfileMemberProfile),
	}
}

func (m *ModelTransformer) MemberRelativeAccountsToResourceList(accountList []*MemberRelativeAccounts) []*MemberRelativeAccountsResource {
	if accountList == nil {
		return nil
	}

	var accountResources []*MemberRelativeAccountsResource
	for _, account := range accountList {
		accountResources = append(accountResources, m.MemberRelativeAccountsToResource(account))
	}
	return accountResources
}

func (m *ModelRepository) MemberRelativeAccountsGetByID(id string, preloads ...string) (*MemberRelativeAccounts, error) {
	repo := NewGenericRepository[MemberRelativeAccounts](m.db.Client)
	return repo.GetByID(id, preloads...)
}

func (m *ModelRepository) MemberRelativeAccountsCreate(memberrelativeaccounts *MemberRelativeAccounts, preloads ...string) (*MemberRelativeAccounts, error) {
	repo := NewGenericRepository[MemberRelativeAccounts](m.db.Client)
	return repo.Create(memberrelativeaccounts, preloads...)
}
func (m *ModelRepository) MemberRelativeAccountsUpdate(memberrelativeaccounts *MemberRelativeAccounts, preloads ...string) (*MemberRelativeAccounts, error) {
	repo := NewGenericRepository[MemberRelativeAccounts](m.db.Client)
	return repo.Update(memberrelativeaccounts, preloads...)
}
func (m *ModelRepository) MemberRelativeAccountsUpdateByID(id string, column string, value interface{}, preloads ...string) (*MemberRelativeAccounts, error) {
	repo := NewGenericRepository[MemberRelativeAccounts](m.db.Client)
	return repo.UpdateByID(id, column, value, preloads...)
}
func (m *ModelRepository) MemberRelativeAccountsDeleteByID(id string) error {
	repo := NewGenericRepository[MemberRelativeAccounts](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) MemberRelativeAccountsGetAll(preloads ...string) ([]*MemberRelativeAccounts, error) {
	repo := NewGenericRepository[MemberRelativeAccounts](m.db.Client)
	return repo.GetAll(preloads...)
}
