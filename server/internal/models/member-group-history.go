package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberGroupHistory struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	MemberProfileID uuid.UUID      `gorm:"unsigned" json:"member_profile_id"`
	MemberGroupID   uuid.UUID      `gorm:"type:bigint;unsigned;unsigned" json:"member_group_id"`
	MemberProfile   *MemberProfile `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	MemberGroup     *MemberGroup   `gorm:"foreignKey:MemberGroupID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_group"`
}

func (v *MemberGroupHistory) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type MemberGroupHistoryResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	MemberProfileID uuid.UUID              `json:"memberProfileID"`
	MemberGroupID   uuid.UUID              `json:"memberGroupID"`
	MemberProfile   *MemberProfileResource `json:"memberProfile,omitempty"`
	MemberGroup     *MemberGroupResource   `json:"memberGroup,omitempty"`

	CompanyID uuid.UUID `gorm:"unsigned" json:"company_id"`
	Company   *Company  `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"company"`
}

func (m *ModelTransformer) MemberGroupHistoryToResource(history *MemberGroupHistory) *MemberGroupHistoryResource {
	if history == nil {
		return nil
	}

	return &MemberGroupHistoryResource{

		ID:        history.ID,
		CreatedAt: history.CreatedAt.Format(time.RFC3339),
		UpdatedAt: history.UpdatedAt.Format(time.RFC3339),
		DeletedAt: history.DeletedAt.Time.Format(time.RFC3339),

		MemberProfileID: history.MemberProfileID,
		MemberGroupID:   history.MemberGroupID,
		MemberProfile:   m.MemberProfileToResource(history.MemberProfile),
		MemberGroup:     m.MemberGroupToResource(history.MemberGroup),
	}
}

func (m *ModelTransformer) MemberGroupHistoryToResourceList(historyList []*MemberGroupHistory) []*MemberGroupHistoryResource {
	if historyList == nil {
		return nil
	}

	var historyResources []*MemberGroupHistoryResource
	for _, history := range historyList {
		historyResources = append(historyResources, m.MemberGroupHistoryToResource(history))
	}
	return historyResources
}

func (m *ModelRepository) MemberGroupHistoryGetByID(id string, preloads ...string) (*MemberGroupHistory, error) {
	repo := NewGenericRepository[MemberGroupHistory](m.db.Client)
	return repo.GetByID(id, preloads...)
}

func (m *ModelRepository) MemberGroupHistoryCreate(membergrouphistory *MemberGroupHistory, preloads ...string) (*MemberGroupHistory, error) {
	repo := NewGenericRepository[MemberGroupHistory](m.db.Client)
	return repo.Create(membergrouphistory, preloads...)
}
func (m *ModelRepository) MemberGroupHistoryUpdate(membergrouphistory *MemberGroupHistory, preloads ...string) (*MemberGroupHistory, error) {
	repo := NewGenericRepository[MemberGroupHistory](m.db.Client)
	return repo.Update(membergrouphistory, preloads...)
}
func (m *ModelRepository) MemberGroupHistoryUpdateByID(id string, value *MemberGroupHistory, preloads ...string) (*MemberGroupHistory, error) {
	repo := NewGenericRepository[MemberGroupHistory](m.db.Client)
	return repo.UpdateByID(id, value, preloads...)
}
func (m *ModelRepository) MemberGroupHistoryDeleteByID(id string) error {
	repo := NewGenericRepository[MemberGroupHistory](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) MemberGroupHistoryGetAll(preloads ...string) ([]*MemberGroupHistory, error) {
	repo := NewGenericRepository[MemberGroupHistory](m.db.Client)
	return repo.GetAll(preloads...)
}
