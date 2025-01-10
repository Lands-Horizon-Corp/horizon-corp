package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberTypeHistory struct {
	gorm.Model
	MemberProfileID uint           `gorm:"unsigned" json:"member_profile_id"`
	MemberTypeID    uint           `gorm:"type:bigint;unsigned;unsigned" json:"member_type_id"`
	MemberProfile   *MemberProfile `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	MemberType      *MemberType    `gorm:"foreignKey:MemberTypeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_type"`
}

type MemberTypeHistoryResource struct {
	ID              uint                   `json:"id"`
	CreatedAt       string                 `json:"createdAt"`
	UpdatedAt       string                 `json:"updatedAt"`
	MemberProfileID uint                   `json:"memberProfileID"`
	MemberTypeID    uint                   `json:"memberTypeID"`
	MemberProfile   *MemberProfileResource `json:"memberProfile,omitempty"`
	MemberType      *MemberTypeResource    `json:"memberType,omitempty"`
}

func (m *ModelTransformer) MemberTypeHistoryToResource(history *MemberTypeHistory) *MemberTypeHistoryResource {
	if history == nil {
		return nil
	}

	return &MemberTypeHistoryResource{
		ID:              history.ID,
		CreatedAt:       history.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       history.UpdatedAt.Format(time.RFC3339),
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
