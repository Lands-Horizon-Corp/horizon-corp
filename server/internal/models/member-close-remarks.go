package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberCloseRemarks struct {
	gorm.Model
	MembersProfileID uint           `gorm:"unsigned" json:"members_profile_id"`
	Description      string         `gorm:"type:text" json:"description"`
	MembersProfile   *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

type MemberCloseRemarksResource struct {
	ID               uint                   `json:"id"`
	CreatedAt        string                 `json:"createdAt"`
	UpdatedAt        string                 `json:"updatedAt"`
	MembersProfileID uint                   `json:"membersProfileID"`
	Description      string                 `json:"description"`
	MembersProfile   *MemberProfileResource `json:"membersProfile,omitempty"`
}

func (m *ModelTransformer) MemberCloseRemarksToResource(remarks *MemberCloseRemarks) *MemberCloseRemarksResource {
	if remarks == nil {
		return nil
	}

	return &MemberCloseRemarksResource{
		ID:               remarks.ID,
		CreatedAt:        remarks.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        remarks.UpdatedAt.Format(time.RFC3339),
		MembersProfileID: remarks.MembersProfileID,
		Description:      remarks.Description,
		MembersProfile:   m.MemberProfileToResource(remarks.MembersProfile),
	}
}

func (m *ModelTransformer) MemberCloseRemarksToResourceList(remarksList []*MemberCloseRemarks) []*MemberCloseRemarksResource {
	if remarksList == nil {
		return nil
	}

	var remarksResources []*MemberCloseRemarksResource
	for _, remarks := range remarksList {
		remarksResources = append(remarksResources, m.MemberCloseRemarksToResource(remarks))
	}
	return remarksResources
}
