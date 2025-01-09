package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberOccupation struct {
	gorm.Model
	Name        string                     `gorm:"size:255;not null"`
	Description string                     `gorm:"size:500"`
	History     []*MemberOccupationHistory `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

type MemberOccupationResource struct {
	ID          uint                               `json:"id"`
	CreatedAt   string                             `json:"createdAt"`
	UpdatedAt   string                             `json:"updatedAt"`
	Name        string                             `json:"name"`
	Description string                             `json:"description"`
	History     []*MemberOccupationHistoryResource `json:"history,omitempty"`
}

func (m *ModelTransformer) MemberOccupationToResource(occupation *MemberOccupation) *MemberOccupationResource {
	if occupation == nil {
		return nil
	}

	return &MemberOccupationResource{
		ID:          occupation.ID,
		CreatedAt:   occupation.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   occupation.UpdatedAt.Format(time.RFC3339),
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
