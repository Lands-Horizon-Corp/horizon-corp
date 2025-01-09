package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberCenterHistory struct {
	gorm.Model
	MemberProfileID uint           `gorm:"not null" json:"member_profile_id"`
	MemberCenterID  uint           `gorm:"type:bigint;unsigned;not null" json:"member_center_id"`
	MemberProfile   *MemberProfile `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	MemberCenter    *MemberCenter  `gorm:"foreignKey:MemberCenterID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_center"`
}

type MemberCenterHistoryResource struct {
	ID              uint                   `json:"id"`
	CreatedAt       string                 `json:"createdAt"`
	UpdatedAt       string                 `json:"updatedAt"`
	MemberProfileID uint                   `json:"memberProfileID"`
	MemberCenterID  uint                   `json:"memberCenterID"`
	MemberProfile   *MemberProfileResource `json:"memberProfile,omitempty"`
	MemberCenter    *MemberCenterResource  `json:"memberCenter,omitempty"`
}

func (m *ModelTransformer) MemberCenterHistoryToResource(history *MemberCenterHistory) *MemberCenterHistoryResource {
	if history == nil {
		return nil
	}

	return &MemberCenterHistoryResource{
		ID:              history.ID,
		CreatedAt:       history.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       history.UpdatedAt.Format(time.RFC3339),
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
