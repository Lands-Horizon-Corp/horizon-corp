package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberOccupation struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name        string                     `gorm:"size:255;unsigned"`
	Description string                     `gorm:"size:500"`
	History     []*MemberOccupationHistory `gorm:"foreignKey:MemberOccupationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

type MemberOccupationResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	Name        string                             `json:"name"`
	Description string                             `json:"description"`
	History     []*MemberOccupationHistoryResource `json:"history,omitempty"`
}

func (m *ModelTransformer) MemberOccupationToResource(occupation *MemberOccupation) *MemberOccupationResource {
	if occupation == nil {
		return nil
	}

	return &MemberOccupationResource{

		ID:        occupation.ID,
		CreatedAt: occupation.CreatedAt.Format(time.RFC3339),
		UpdatedAt: occupation.UpdatedAt.Format(time.RFC3339),
		DeletedAt: occupation.DeletedAt.Time.Format(time.RFC3339),

		Name:        occupation.Name,
		Description: occupation.Description,
		History:     m.MemberOccupationHistoryToResourceList(occupation.History),
	}
}

func (m *ModelTransformer) MemberOccupationToResourceList(occupationList []*MemberOccupation) []*MemberOccupationResource {
	if occupationList == nil {
		return nil
	}

	var occupationResources []*MemberOccupationResource
	for _, occupation := range occupationList {
		occupationResources = append(occupationResources, m.MemberOccupationToResource(occupation))
	}
	return occupationResources
}
