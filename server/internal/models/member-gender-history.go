package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberGenderHistory struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	MemberProfileID uuid.UUID      `gorm:"unsigned" json:"member_profile_id"`
	MemberGenderID  uuid.UUID      `gorm:"type:bigint;unsigned;unsigned" json:"member_gender_id"`
	MemberProfile   *MemberProfile `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	MemberGender    *MemberGender  `gorm:"foreignKey:MemberGenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_gender"`
	CompanyID       uuid.UUID      `gorm:"unsigned" json:"company_id"`
	Company         *Company       `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"company"`
}

func (v *MemberGenderHistory) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type MemberGenderHistoryResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	MemberProfileID uuid.UUID              `json:"memberProfileID"`
	MemberGenderID  uuid.UUID              `json:"memberGenderID"`
	MemberProfile   *MemberProfileResource `json:"memberProfile,omitempty"`
	MemberGender    *MemberGenderResource  `json:"memberGender,omitempty"`
}

func (m *ModelTransformer) MemberGenderHistoryToResource(history *MemberGenderHistory) *MemberGenderHistoryResource {
	if history == nil {
		return nil
	}

	return &MemberGenderHistoryResource{

		ID:        history.ID,
		CreatedAt: history.CreatedAt.Format(time.RFC3339),
		UpdatedAt: history.UpdatedAt.Format(time.RFC3339),
		DeletedAt: history.DeletedAt.Time.Format(time.RFC3339),

		MemberProfileID: history.MemberProfileID,
		MemberGenderID:  history.MemberGenderID,
		MemberProfile:   m.MemberProfileToResource(history.MemberProfile),
		MemberGender:    m.MemberGenderToResource(history.MemberGender),
	}
}

func (m *ModelTransformer) MemberGenderHistoryToResourceList(historyList []*MemberGenderHistory) []*MemberGenderHistoryResource {
	if historyList == nil {
		return nil
	}

	var historyResources []*MemberGenderHistoryResource
	for _, history := range historyList {
		historyResources = append(historyResources, m.MemberGenderHistoryToResource(history))
	}
	return historyResources
}

func (m *ModelRepository) MemberGenderHistoryGetByID(id string, preloads ...string) (*MemberGenderHistory, error) {
	repo := NewGenericRepository[MemberGenderHistory](m.db.Client)
	return repo.GetByID(id, preloads...)
}

func (m *ModelRepository) MemberGenderHistoryCreate(membergenderhistory *MemberGenderHistory, preloads ...string) (*MemberGenderHistory, error) {
	repo := NewGenericRepository[MemberGenderHistory](m.db.Client)
	return repo.Create(membergenderhistory, preloads...)
}
func (m *ModelRepository) MemberGenderHistoryUpdate(membergenderhistory *MemberGenderHistory, preloads ...string) (*MemberGenderHistory, error) {
	repo := NewGenericRepository[MemberGenderHistory](m.db.Client)
	return repo.Update(membergenderhistory, preloads...)
}
func (m *ModelRepository) MemberGenderHistoryUpdateByID(id string, column string, value interface{}, preloads ...string) (*MemberGenderHistory, error) {
	repo := NewGenericRepository[MemberGenderHistory](m.db.Client)
	return repo.UpdateByID(id, column, value, preloads...)
}
func (m *ModelRepository) MemberGenderHistoryDeleteByID(id string) error {
	repo := NewGenericRepository[MemberGenderHistory](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) MemberGenderHistoryGetAll(preloads ...string) ([]*MemberGenderHistory, error) {
	repo := NewGenericRepository[MemberGenderHistory](m.db.Client)
	return repo.GetAll(preloads...)
}
