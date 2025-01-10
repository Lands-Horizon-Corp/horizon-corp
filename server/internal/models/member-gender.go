package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberGender struct {
	gorm.Model
	Name        string                 `gorm:"size:255;unsigned"`
	Description string                 `gorm:"size:500"`
	History     []*MemberGenderHistory `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

type MemberGenderResource struct {
	ID          uint                           `json:"id"`
	CreatedAt   string                         `json:"createdAt"`
	UpdatedAt   string                         `json:"updatedAt"`
	Name        string                         `json:"name"`
	Description string                         `json:"description"`
	History     []*MemberGenderHistoryResource `json:"history,omitempty"`
}

func (m *ModelTransformer) MemberGenderToResource(gender *MemberGender) *MemberGenderResource {
	if gender == nil {
		return nil
	}

	return &MemberGenderResource{
		ID:          gender.ID,
		CreatedAt:   gender.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   gender.UpdatedAt.Format(time.RFC3339),
		Name:        gender.Name,
		Description: gender.Description,
		History:     m.MemberGenderHistoryToResourceList(gender.History),
	}
}

func (m *ModelTransformer) MemberGenderToResourceList(genderList []*MemberGender) []*MemberGenderResource {
	if genderList == nil {
		return nil
	}

	var genderResources []*MemberGenderResource
	for _, gender := range genderList {
		genderResources = append(genderResources, m.MemberGenderToResource(gender))
	}
	return genderResources
}
