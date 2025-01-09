package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberType struct {
	gorm.Model
	Name        string               `gorm:"size:255;not null"`
	Description string               `gorm:"size:500"`
	Prefix      string               `gorm:"size:100"`
	History     []*MemberTypeHistory `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

type MemberTypeResource struct {
	ID          uint                         `json:"id"`
	CreatedAt   string                       `json:"createdAt"`
	UpdatedAt   string                       `json:"updatedAt"`
	Name        string                       `json:"name"`
	Description string                       `json:"description"`
	Prefix      string                       `json:"prefix"`
	History     []*MemberTypeHistoryResource `json:"history,omitempty"`
}

func (m *ModelTransformer) MemberTypeToResource(memberType *MemberType) *MemberTypeResource {
	if memberType == nil {
		return nil
	}

	return &MemberTypeResource{
		ID:          memberType.ID,
		CreatedAt:   memberType.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   memberType.UpdatedAt.Format(time.RFC3339),
		Name:        memberType.Name,
		Description: memberType.Description,
		Prefix:      memberType.Prefix,
		History:     m.MemberTypeHistoryToResourceList(memberType.History),
	}
}

func (m *ModelTransformer) MemberTypeToResourceList(memberTypeList []*MemberType) []*MemberTypeResource {
	if memberTypeList == nil {
		return nil
	}

	var memberTypeResources []*MemberTypeResource
	for _, memberType := range memberTypeList {
		memberTypeResources = append(memberTypeResources, m.MemberTypeToResource(memberType))
	}
	return memberTypeResources
}
