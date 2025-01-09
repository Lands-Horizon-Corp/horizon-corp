package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberGroupHistory struct {
	gorm.Model
	MemberProfileID uint           `gorm:"not null" json:"member_profile_id"`
	MemberGroupID   uint           `gorm:"type:bigint;unsigned;not null" json:"member_group_id"`
	MemberProfile   *MemberProfile `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	MemberGroup     *MemberGroup   `gorm:"foreignKey:MemberGroupID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_group"`
}

type MemberGroupHistoryResource struct {
	ID              uint                   `json:"id"`
	CreatedAt       string                 `json:"createdAt"`
	UpdatedAt       string                 `json:"updatedAt"`
	MemberProfileID uint                   `json:"memberProfileID"`
	MemberGroupID   uint                   `json:"memberGroupID"`
	MemberProfile   *MemberProfileResource `json:"memberProfile,omitempty"`
	MemberGroup     *MemberGroupResource   `json:"memberGroup,omitempty"`
}

func (m *ModelTransformer) MemberGroupHistoryToResource(history *MemberGroupHistory) *MemberGroupHistoryResource {
	if history == nil {
		return nil
	}

	return &MemberGroupHistoryResource{
		ID:              history.ID,
		CreatedAt:       history.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       history.UpdatedAt.Format(time.RFC3339),
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
