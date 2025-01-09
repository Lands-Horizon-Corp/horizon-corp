package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberClassification struct {
	gorm.Model
	Name        string                         `gorm:"size:255;not null"`
	Description string                         `gorm:"size:500"`
	History     []*MemberClassificationHistory `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

type MemberClassificationResource struct {
	ID          uint                                   `json:"id"`
	CreatedAt   string                                 `json:"createdAt"`
	UpdatedAt   string                                 `json:"updatedAt"`
	Name        string                                 `json:"name"`
	Description string                                 `json:"description"`
	History     []*MemberClassificationHistoryResource `json:"history,omitempty"`
}

func (m *ModelTransformer) MemberClassificationToResource(classification *MemberClassification) *MemberClassificationResource {
	if classification == nil {
		return nil
	}

	return &MemberClassificationResource{
		ID:          classification.ID,
		CreatedAt:   classification.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   classification.UpdatedAt.Format(time.RFC3339),
		Name:        classification.Name,
		Description: classification.Description,
		History:     m.MemberClassificationHistoryToResourceList(classification.History),
	}
}

func (m *ModelTransformer) MemberClassificationToResourceList(classificationList []*MemberClassification) []*MemberClassificationResource {
	if classificationList == nil {
		return nil
	}

	var classificationResources []*MemberClassificationResource
	for _, classification := range classificationList {
		classificationResources = append(classificationResources, m.MemberClassificationToResource(classification))
	}
	return classificationResources
}
