package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberCenterHistory struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	MemberProfileID uuid.UUID      `gorm:"unsigned" json:"member_profile_id"`
	MemberCenterID  uuid.UUID      `gorm:"type:bigint;unsigned;unsigned" json:"member_center_id"`
	MemberProfile   *MemberProfile `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	MemberCenter    *MemberCenter  `gorm:"foreignKey:MemberCenterID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_center"`

	CompanyID uuid.UUID `gorm:"unsigned" json:"company_id"`
	Company   *Company  `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"company"`
}

func (v *MemberCenterHistory) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type MemberCenterHistoryResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	MemberProfileID uuid.UUID              `json:"memberProfileID"`
	MemberCenterID  uuid.UUID              `json:"memberCenterID"`
	MemberProfile   *MemberProfileResource `json:"memberProfile,omitempty"`
	MemberCenter    *MemberCenterResource  `json:"memberCenter,omitempty"`
}

func (m *ModelTransformer) MemberCenterHistoryToResource(history *MemberCenterHistory) *MemberCenterHistoryResource {
	if history == nil {
		return nil
	}

	return &MemberCenterHistoryResource{

		ID:        history.ID,
		CreatedAt: history.CreatedAt.Format(time.RFC3339),
		UpdatedAt: history.UpdatedAt.Format(time.RFC3339),
		DeletedAt: history.DeletedAt.Time.Format(time.RFC3339),

		MemberProfileID: history.MemberProfileID,
		MemberCenterID:  history.MemberCenterID,
		MemberProfile:   m.MemberProfileToResource(history.MemberProfile),
		MemberCenter:    m.MemberCenterToResource(history.MemberCenter),
	}
}

func (m *ModelTransformer) MemberCenterHistoryToResourceList(historyList []*MemberCenterHistory) []*MemberCenterHistoryResource {
	if historyList == nil {
		return nil
	}

	var historyResources []*MemberCenterHistoryResource
	for _, history := range historyList {
		historyResources = append(historyResources, m.MemberCenterHistoryToResource(history))
	}
	return historyResources
}

func (m *ModelRepository) MemberCenterHistoryGetByID(id string, preloads ...string) (*MemberCenterHistory, error) {
	repo := NewGenericRepository[MemberCenterHistory](m.db.Client)
	return repo.GetByID(id, preloads...)
}

func (m *ModelRepository) MemberCenterHistoryCreate(membercenterhistory *MemberCenterHistory, preloads ...string) (*MemberCenterHistory, error) {
	repo := NewGenericRepository[MemberCenterHistory](m.db.Client)
	return repo.Create(membercenterhistory, preloads...)
}
func (m *ModelRepository) MemberCenterHistoryUpdate(membercenterhistory *MemberCenterHistory, preloads ...string) (*MemberCenterHistory, error) {
	repo := NewGenericRepository[MemberCenterHistory](m.db.Client)
	return repo.Update(membercenterhistory, preloads...)
}
func (m *ModelRepository) MemberCenterHistoryUpdateByID(id string, value *MemberCenterHistory, preloads ...string) (*MemberCenterHistory, error) {
	repo := NewGenericRepository[MemberCenterHistory](m.db.Client)
	return repo.UpdateByID(id, value, preloads...)
}
func (m *ModelRepository) MemberCenterHistoryDeleteByID(id string) error {
	repo := NewGenericRepository[MemberCenterHistory](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) MemberCenterHistoryGetAll(preloads ...string) ([]*MemberCenterHistory, error) {
	repo := NewGenericRepository[MemberCenterHistory](m.db.Client)
	return repo.GetAll(preloads...)
}
