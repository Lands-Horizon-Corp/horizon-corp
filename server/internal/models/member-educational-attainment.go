package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberEducationalAttainment struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name        string                                `gorm:"size:255;unsigned"`
	Description string                                `gorm:"size:500"`
	History     []*MemberEducationalAttainmentHistory `gorm:"foreignKey:MemberEducationalAttainmentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

func (v *MemberEducationalAttainment) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type MemberEducationalAttainmentResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	Name        string                                        `json:"name"`
	Description string                                        `json:"description"`
	History     []*MemberEducationalAttainmentHistoryResource `json:"history,omitempty"`
}

func (m *ModelTransformer) MemberEducationalAttainmentToResource(attainment *MemberEducationalAttainment) *MemberEducationalAttainmentResource {
	if attainment == nil {
		return nil
	}

	return &MemberEducationalAttainmentResource{

		ID:        attainment.ID,
		CreatedAt: attainment.CreatedAt.Format(time.RFC3339),
		UpdatedAt: attainment.UpdatedAt.Format(time.RFC3339),
		DeletedAt: attainment.DeletedAt.Time.Format(time.RFC3339),

		Name:        attainment.Name,
		Description: attainment.Description,
		History:     m.MemberEducationalAttainmentHistoryToResourceList(attainment.History),
	}
}

func (m *ModelTransformer) MemberEducationalAttainmentToResourceList(attainmentList []*MemberEducationalAttainment) []*MemberEducationalAttainmentResource {
	if attainmentList == nil {
		return nil
	}

	var attainmentResources []*MemberEducationalAttainmentResource
	for _, attainment := range attainmentList {
		attainmentResources = append(attainmentResources, m.MemberEducationalAttainmentToResource(attainment))
	}
	return attainmentResources
}
