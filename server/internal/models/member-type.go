package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberType struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name        string `gorm:"size:255;unsigned"`
	Description string `gorm:"size:500"`
	Prefix      string `gorm:"size:100"`

	MembersProfile *MemberProfile       `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
	History        []*MemberTypeHistory `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

func (v *MemberType) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type MemberTypeResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

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

		ID:        memberType.ID,
		CreatedAt: memberType.CreatedAt.Format(time.RFC3339),
		UpdatedAt: memberType.UpdatedAt.Format(time.RFC3339),
		DeletedAt: memberType.DeletedAt.Time.Format(time.RFC3339),

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
