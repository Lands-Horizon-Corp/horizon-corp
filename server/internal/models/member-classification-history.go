package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberClassificationHistory struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	MemberProfileID        uuid.UUID             `gorm:"unsigned" json:"member_profile_id"`
	MemberClassificationID uuid.UUID             `gorm:"type:bigint;unsigned;unsigned" json:"member_classification_id"`
	MemberProfile          *MemberProfile        `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	MemberClassification   *MemberClassification `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_classification"`
}

func (v *MemberClassificationHistory) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type MemberClassificationHistoryResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	MemberProfileID        uuid.UUID                     `json:"memberProfileID"`
	MemberClassificationID uuid.UUID                     `json:"memberClassificationID"`
	MemberProfile          *MemberProfileResource        `json:"memberProfile,omitempty"`
	MemberClassification   *MemberClassificationResource `json:"memberClassification,omitempty"`
}

func (m *ModelTransformer) MemberClassificationHistoryToResource(history *MemberClassificationHistory) *MemberClassificationHistoryResource {
	if history == nil {
		return nil
	}

	return &MemberClassificationHistoryResource{
		ID:        history.ID,
		CreatedAt: history.CreatedAt.Format(time.RFC3339),
		UpdatedAt: history.UpdatedAt.Format(time.RFC3339),
		DeletedAt: history.DeletedAt.Time.Format(time.RFC3339),

		MemberProfileID:        history.MemberProfileID,
		MemberClassificationID: history.MemberClassificationID,
		MemberProfile:          m.MemberProfileToResource(history.MemberProfile),
		MemberClassification:   m.MemberClassificationToResource(history.MemberClassification),
	}
}

func (m *ModelTransformer) MemberClassificationHistoryToResourceList(historyList []*MemberClassificationHistory) []*MemberClassificationHistoryResource {
	if historyList == nil {
		return nil
	}

	var historyResources []*MemberClassificationHistoryResource
	for _, history := range historyList {
		historyResources = append(historyResources, m.MemberClassificationHistoryToResource(history))
	}
	return historyResources
}

func (m *ModelRepository) MemberClassificationHistoryGetByID(id string, preloads ...string) (*MemberClassificationHistory, error) {
	repo := NewGenericRepository[MemberClassificationHistory](m.db.Client)
	return repo.GetByID(id, preloads...)
}

func (m *ModelRepository) MemberClassificationHistoryCreate(memberclassificationhistory *MemberClassificationHistory, preloads ...string) (*MemberClassificationHistory, error) {
	repo := NewGenericRepository[MemberClassificationHistory](m.db.Client)
	return repo.Create(memberclassificationhistory, preloads...)
}
func (m *ModelRepository) MemberClassificationHistoryUpdate(memberclassificationhistory *MemberClassificationHistory, preloads ...string) (*MemberClassificationHistory, error) {
	repo := NewGenericRepository[MemberClassificationHistory](m.db.Client)
	return repo.Update(memberclassificationhistory, preloads...)
}
func (m *ModelRepository) MemberClassificationHistoryUpdateByID(id string, column string, value interface{}, preloads ...string) (*MemberClassificationHistory, error) {
	repo := NewGenericRepository[MemberClassificationHistory](m.db.Client)
	return repo.UpdateByID(id, column, value, preloads...)
}
func (m *ModelRepository) MemberClassificationHistoryDeleteByID(id string) error {
	repo := NewGenericRepository[MemberClassificationHistory](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) MemberClassificationHistoryGetAll(preloads ...string) ([]*MemberClassificationHistory, error) {
	repo := NewGenericRepository[MemberClassificationHistory](m.db.Client)
	return repo.GetAll(preloads...)
}
