package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberClassificationHistory struct {
	gorm.Model
	MemberProfileID        uint                  `gorm:"unsigned" json:"member_profile_id"`
	MemberClassificationID uint                  `gorm:"type:bigint;unsigned;unsigned" json:"member_classification_id"`
	MemberProfile          *MemberProfile        `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	MemberClassification   *MemberClassification `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_classification"`
}

type MemberClassificationHistoryResource struct {
	ID                     uint                          `json:"id"`
	CreatedAt              string                        `json:"createdAt"`
	UpdatedAt              string                        `json:"updatedAt"`
	MemberProfileID        uint                          `json:"memberProfileID"`
	MemberClassificationID uint                          `json:"memberClassificationID"`
	MemberProfile          *MemberProfileResource        `json:"memberProfile,omitempty"`
	MemberClassification   *MemberClassificationResource `json:"memberClassification,omitempty"`
}

func (m *ModelTransformer) MemberClassificationHistoryToResource(history *MemberClassificationHistory) *MemberClassificationHistoryResource {
	if history == nil {
		return nil
	}

	return &MemberClassificationHistoryResource{
		ID:                     history.ID,
		CreatedAt:              history.CreatedAt.Format(time.RFC3339),
		UpdatedAt:              history.UpdatedAt.Format(time.RFC3339),
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
