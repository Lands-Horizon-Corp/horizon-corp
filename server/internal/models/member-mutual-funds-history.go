package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberMutualFundsHistory struct {
	gorm.Model
	MembersProfileID uint           `gorm:"unsigned" json:"members_profile_id"`
	Description      string         `gorm:"type:text" json:"description"`
	Amount           float64        `gorm:"type:decimal(10,2);unsigned" json:"amount"`
	MembersProfile   *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

type MemberMutualFundsHistoryResource struct {
	ID               uint                   `json:"id"`
	CreatedAt        string                 `json:"createdAt"`
	UpdatedAt        string                 `json:"updatedAt"`
	MembersProfileID uint                   `json:"membersProfileID"`
	Description      string                 `json:"description"`
	Amount           float64                `json:"amount"`
	MembersProfile   *MemberProfileResource `json:"membersProfile,omitempty"`
}

func (m *ModelTransformer) MemberMutualFundsHistoryToResource(history *MemberMutualFundsHistory) *MemberMutualFundsHistoryResource {
	if history == nil {
		return nil
	}

	return &MemberMutualFundsHistoryResource{
		ID:               history.ID,
		CreatedAt:        history.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        history.UpdatedAt.Format(time.RFC3339),
		MembersProfileID: history.MembersProfileID,
		Description:      history.Description,
		Amount:           history.Amount,
		MembersProfile:   m.MemberProfileToResource(history.MembersProfile),
	}
}

func (m *ModelTransformer) MemberMutualFundsHistoryToResourceList(historyList []*MemberMutualFundsHistory) []*MemberMutualFundsHistoryResource {
	if historyList == nil {
		return nil
	}

	var historyResources []*MemberMutualFundsHistoryResource
	for _, history := range historyList {
		historyResources = append(historyResources, m.MemberMutualFundsHistoryToResource(history))
	}
	return historyResources
}
