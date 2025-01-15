package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberEducationalAttainmentHistory struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	MemberProfileID               uuid.UUID                    `gorm:"unsigned" json:"member_profile_id"`
	MemberEducationalAttainmentID uuid.UUID                    `gorm:"type:bigint;unsigned;unsigned" json:"member_educational_attainment_id"`
	MemberProfile                 *MemberProfile               `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	MemberEducationalAttainment   *MemberEducationalAttainment `gorm:"foreignKey:MemberEducationalAttainmentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_educational_attainment"`
	CompanyID                     uuid.UUID                    `gorm:"unsigned" json:"company_id"`
	Company                       *Company                     `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"company"`
}

func (v *MemberEducationalAttainmentHistory) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type MemberEducationalAttainmentHistoryResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	MemberProfileID               uuid.UUID                            `json:"memberProfileID"`
	MemberEducationalAttainmentID uuid.UUID                            `json:"memberEducationalAttainmentID"`
	MemberProfile                 *MemberProfileResource               `json:"memberProfile,omitempty"`
	MemberEducationalAttainment   *MemberEducationalAttainmentResource `json:"memberEducationalAttainment,omitempty"`
}

func (m *ModelTransformer) MemberEducationalAttainmentHistoryToResource(history *MemberEducationalAttainmentHistory) *MemberEducationalAttainmentHistoryResource {
	if history == nil {
		return nil
	}

	return &MemberEducationalAttainmentHistoryResource{

		ID:        history.ID,
		CreatedAt: history.CreatedAt.Format(time.RFC3339),
		UpdatedAt: history.UpdatedAt.Format(time.RFC3339),
		DeletedAt: history.DeletedAt.Time.Format(time.RFC3339),

		MemberProfileID:               history.MemberProfileID,
		MemberEducationalAttainmentID: history.MemberEducationalAttainmentID,
		MemberProfile:                 m.MemberProfileToResource(history.MemberProfile),
		MemberEducationalAttainment:   m.MemberEducationalAttainmentToResource(history.MemberEducationalAttainment),
	}
}

func (m *ModelTransformer) MemberEducationalAttainmentHistoryToResourceList(historyList []*MemberEducationalAttainmentHistory) []*MemberEducationalAttainmentHistoryResource {
	if historyList == nil {
		return nil
	}

	var historyResources []*MemberEducationalAttainmentHistoryResource
	for _, history := range historyList {
		historyResources = append(historyResources, m.MemberEducationalAttainmentHistoryToResource(history))
	}
	return historyResources
}

func (m *ModelRepository) MemberEducationalAttainmentHistoryGetByID(id string, preloads ...string) (*MemberEducationalAttainmentHistory, error) {
	repo := NewGenericRepository[MemberEducationalAttainmentHistory](m.db.Client)
	return repo.GetByID(id, preloads...)
}

func (m *ModelRepository) MemberEducationalAttainmentHistoryCreate(membereducationalattainmenthistory *MemberEducationalAttainmentHistory, preloads ...string) (*MemberEducationalAttainmentHistory, error) {
	repo := NewGenericRepository[MemberEducationalAttainmentHistory](m.db.Client)
	return repo.Create(membereducationalattainmenthistory, preloads...)
}
func (m *ModelRepository) MemberEducationalAttainmentHistoryUpdate(membereducationalattainmenthistory *MemberEducationalAttainmentHistory, preloads ...string) (*MemberEducationalAttainmentHistory, error) {
	repo := NewGenericRepository[MemberEducationalAttainmentHistory](m.db.Client)
	return repo.Update(membereducationalattainmenthistory, preloads...)
}
func (m *ModelRepository) MemberEducationalAttainmentHistoryUpdateByID(id string, column string, value interface{}, preloads ...string) (*MemberEducationalAttainmentHistory, error) {
	repo := NewGenericRepository[MemberEducationalAttainmentHistory](m.db.Client)
	return repo.UpdateByID(id, column, value, preloads...)
}
func (m *ModelRepository) MemberEducationalAttainmentHistoryDeleteByID(id string) error {
	repo := NewGenericRepository[MemberEducationalAttainmentHistory](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) MemberEducationalAttainmentHistoryGetAll(preloads ...string) ([]*MemberEducationalAttainmentHistory, error) {
	repo := NewGenericRepository[MemberEducationalAttainmentHistory](m.db.Client)
	return repo.GetAll(preloads...)
}
