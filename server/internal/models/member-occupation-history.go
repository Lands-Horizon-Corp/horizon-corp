package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberOccupationHistory struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	MemberProfileID    uuid.UUID         `gorm:"unsigned" json:"member_profile_id"`
	MemberOccupationID uuid.UUID         `gorm:"type:bigint;unsigned;unsigned" json:"member_occupation_id"`
	MemberProfile      *MemberProfile    `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	MemberOccupation   *MemberOccupation `gorm:"foreignKey:MemberOccupationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_occupation"`
}

type MemberOccupationHistoryResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	MemberProfileID    uuid.UUID                 `json:"memberProfileID"`
	MemberOccupationID uuid.UUID                 `json:"memberOccupationID"`
	MemberProfile      *MemberProfileResource    `json:"memberProfile,omitempty"`
	MemberOccupation   *MemberOccupationResource `json:"memberOccupation,omitempty"`
}

func (m *ModelTransformer) MemberOccupationHistoryToResource(history *MemberOccupationHistory) *MemberOccupationHistoryResource {
	if history == nil {
		return nil
	}

	return &MemberOccupationHistoryResource{

		ID:        history.ID,
		CreatedAt: history.CreatedAt.Format(time.RFC3339),
		UpdatedAt: history.UpdatedAt.Format(time.RFC3339),
		DeletedAt: history.DeletedAt.Time.Format(time.RFC3339),

		MemberProfileID:    history.MemberProfileID,
		MemberOccupationID: history.MemberOccupationID,
		MemberProfile:      m.MemberProfileToResource(history.MemberProfile),
		MemberOccupation:   m.MemberOccupationToResource(history.MemberOccupation),
	}
}

func (m *ModelTransformer) MemberOccupationHistoryToResourceList(historyList []*MemberOccupationHistory) []*MemberOccupationHistoryResource {
	if historyList == nil {
		return nil
	}

	var historyResources []*MemberOccupationHistoryResource
	for _, history := range historyList {
		historyResources = append(historyResources, m.MemberOccupationHistoryToResource(history))
	}
	return historyResources
}
