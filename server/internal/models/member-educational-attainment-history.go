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
