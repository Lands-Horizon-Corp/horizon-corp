package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberMutualFundsHistory struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	MembersProfileID uuid.UUID      `gorm:"unsigned" json:"members_profile_id"`
	Description      string         `gorm:"type:text" json:"description"`
	Amount           float64        `gorm:"type:decimal(10,2);unsigned" json:"amount"`
	MembersProfile   *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

func (v *MemberMutualFundsHistory) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type MemberMutualFundsHistoryResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	MembersProfileID uuid.UUID              `json:"membersProfileID"`
	Description      string                 `json:"description"`
	Amount           float64                `json:"amount"`
	MembersProfile   *MemberProfileResource `json:"membersProfile,omitempty"`
}

func (m *ModelTransformer) MemberMutualFundsHistoryToResource(history *MemberMutualFundsHistory) *MemberMutualFundsHistoryResource {
	if history == nil {
		return nil
	}

	return &MemberMutualFundsHistoryResource{

		ID:        history.ID,
		CreatedAt: history.CreatedAt.Format(time.RFC3339),
		UpdatedAt: history.UpdatedAt.Format(time.RFC3339),
		DeletedAt: history.DeletedAt.Time.Format(time.RFC3339),

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
