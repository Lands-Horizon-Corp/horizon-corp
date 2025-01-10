package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberGroup struct {
	gorm.Model
	Name        string                `gorm:"size:255;unsigned"`
	Description string                `gorm:"size:500"`
	History     []*MemberGroupHistory `gorm:"foreignKey:MemberGroupID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

type MemberGroupResource struct {
	ID          uint                          `json:"id"`
	CreatedAt   string                        `json:"createdAt"`
	UpdatedAt   string                        `json:"updatedAt"`
	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	History     []*MemberGroupHistoryResource `json:"history,omitempty"`
}

func (m *ModelTransformer) MemberGroupToResource(group *MemberGroup) *MemberGroupResource {
	if group == nil {
		return nil
	}

	return &MemberGroupResource{
		ID:          group.ID,
		CreatedAt:   group.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   group.UpdatedAt.Format(time.RFC3339),
		Name:        group.Name,
		Description: group.Description,
		History:     m.MemberGroupHistoryToResourceList(group.History),
	}
}

func (m *ModelTransformer) MemberGroupToResourceList(groupList []*MemberGroup) []*MemberGroupResource {
	if groupList == nil {
		return nil
	}

	var groupResources []*MemberGroupResource
	for _, group := range groupList {
		groupResources = append(groupResources, m.MemberGroupToResource(group))
	}
	return groupResources
}
