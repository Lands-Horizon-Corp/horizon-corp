package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberCenter struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name        string                 `gorm:"size:255;unsigned"`
	Description string                 `gorm:"size:500"`
	History     []*MemberCenterHistory `gorm:"foreignKey:MemberCenterID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

type MemberCenterResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	Name        string                         `json:"name"`
	Description string                         `json:"description"`
	History     []*MemberCenterHistoryResource `json:"history,omitempty"`
}

func (m *ModelTransformer) MemberCenterToResource(center *MemberCenter) *MemberCenterResource {
	if center == nil {
		return nil
	}

	return &MemberCenterResource{

		ID:        center.ID,
		CreatedAt: center.CreatedAt.Format(time.RFC3339),
		UpdatedAt: center.UpdatedAt.Format(time.RFC3339),
		DeletedAt: center.DeletedAt.Time.Format(time.RFC3339),

		Name:        center.Name,
		Description: center.Description,
		History:     m.MemberCenterHistoryToResourceList(center.History),
	}
}

func (m *ModelTransformer) MemberCenterToResourceList(centerList []*MemberCenter) []*MemberCenterResource {
	if centerList == nil {
		return nil
	}

	var centerResources []*MemberCenterResource
	for _, center := range centerList {
		centerResources = append(centerResources, m.MemberCenterToResource(center))
	}
	return centerResources
}
