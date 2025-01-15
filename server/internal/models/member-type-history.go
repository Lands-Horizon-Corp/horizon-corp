package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberTypeHistory struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	MemberProfileID uuid.UUID      `gorm:"unsigned" json:"member_profile_id"`
	MemberTypeID    uuid.UUID      `gorm:"type:bigint;unsigned;unsigned" json:"member_type_id"`
	MemberProfile   *MemberProfile `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	MemberType      *MemberType    `gorm:"foreignKey:MemberTypeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_type"`
}

func (v *MemberTypeHistory) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type MemberTypeHistoryResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	MemberProfileID uuid.UUID              `json:"memberProfileID"`
	MemberTypeID    uuid.UUID              `json:"memberTypeID"`
	MemberProfile   *MemberProfileResource `json:"memberProfile,omitempty"`
	MemberType      *MemberTypeResource    `json:"memberType,omitempty"`
}

func (m *ModelTransformer) MemberTypeHistoryToResource(history *MemberTypeHistory) *MemberTypeHistoryResource {
	if history == nil {
		return nil
	}

	return &MemberTypeHistoryResource{

		ID:        history.ID,
		CreatedAt: history.CreatedAt.Format(time.RFC3339),
		UpdatedAt: history.UpdatedAt.Format(time.RFC3339),
		DeletedAt: history.DeletedAt.Time.Format(time.RFC3339),

		MemberProfileID: history.MemberProfileID,
		MemberTypeID:    history.MemberTypeID,
		MemberProfile:   m.MemberProfileToResource(history.MemberProfile),
		MemberType:      m.MemberTypeToResource(history.MemberType),
	}
}

func (m *ModelTransformer) MemberTypeHistoryToResourceList(historyList []*MemberTypeHistory) []*MemberTypeHistoryResource {
	if historyList == nil {
		return nil
	}

	var historyResources []*MemberTypeHistoryResource
	for _, history := range historyList {
		historyResources = append(historyResources, m.MemberTypeHistoryToResource(history))
	}
	return historyResources
}

func (m *ModelRepository) MemberTypeHistoryGetByID(id string, preloads ...string) (*MemberTypeHistory, error) {
	repo := NewGenericRepository[MemberTypeHistory](m.db.Client)
	return repo.GetByID(id, preloads...)
}

func (m *ModelRepository) MemberTypeHistoryCreate(membertypehistory *MemberTypeHistory, preloads ...string) (*MemberTypeHistory, error) {
	repo := NewGenericRepository[MemberTypeHistory](m.db.Client)
	return repo.Create(membertypehistory, preloads...)
}
func (m *ModelRepository) MemberTypeHistoryUpdate(membertypehistory *MemberTypeHistory, preloads ...string) (*MemberTypeHistory, error) {
	repo := NewGenericRepository[MemberTypeHistory](m.db.Client)
	return repo.Update(membertypehistory, preloads...)
}
func (m *ModelRepository) MemberTypeHistoryUpdateByID(id string, column string, value interface{}, preloads ...string) (*MemberTypeHistory, error) {
	repo := NewGenericRepository[MemberTypeHistory](m.db.Client)
	return repo.UpdateByID(id, column, value, preloads...)
}
func (m *ModelRepository) MemberTypeHistoryDeleteByID(id string) error {
	repo := NewGenericRepository[MemberTypeHistory](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) MemberTypeHistoryGetAll(preloads ...string) ([]*MemberTypeHistory, error) {
	repo := NewGenericRepository[MemberTypeHistory](m.db.Client)
	return repo.GetAll(preloads...)
}
