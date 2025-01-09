package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberGenderHistory struct {
	gorm.Model
	MemberProfileID uint           `gorm:"not null" json:"member_profile_id"`
	MemberGenderID  uint           `gorm:"type:bigint;unsigned;not null" json:"member_gender_id"`
	MemberProfile   *MemberProfile `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	MemberGender    *MemberGender  `gorm:"foreignKey:MemberGenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_gender"`
}

type MemberGenderHistoryResource struct {
	ID              uint                   `json:"id"`
	CreatedAt       string                 `json:"createdAt"`
	UpdatedAt       string                 `json:"updatedAt"`
	MemberProfileID uint                   `json:"memberProfileID"`
	MemberGenderID  uint                   `json:"memberGenderID"`
	MemberProfile   *MemberProfileResource `json:"memberProfile,omitempty"`
	MemberGender    *MemberGenderResource  `json:"memberGender,omitempty"`
}

func (m *ModelTransformer) MemberGenderHistoryToResource(history *MemberGenderHistory) *MemberGenderHistoryResource {
	if history == nil {
		return nil
	}

	return &MemberGenderHistoryResource{
		ID:              history.ID,
		CreatedAt:       history.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       history.UpdatedAt.Format(time.RFC3339),
		MemberProfileID: history.MemberProfileID,
		MemberGenderID:  history.MemberGenderID,
		MemberProfile:   m.MemberProfileToResource(history.MemberProfile),
		MemberGender:    m.MemberGenderToResource(history.MemberGender),
	}
}

func (m *ModelTransformer) MemberGenderHistoryToResourceList(historyList []*MemberGenderHistory) []*MemberGenderHistoryResource {
	if historyList == nil {
		return nil
	}

	var historyResources []*MemberGenderHistoryResource
	for _, history := range historyList {
		historyResources = append(historyResources, m.MemberGenderHistoryToResource(history))
	}
	return historyResources
}
