package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberDescription struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	MembersProfileID uuid.UUID      `gorm:"unsigned" json:"members_profile_id"`
	Date             time.Time      `gorm:"type:date" json:"date"`
	Description      string         `gorm:"type:text" json:"description"`
	Name             string         `gorm:"type:varchar(255);unsigned" json:"name"`
	MembersProfile   *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

type MemberDescriptionResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	MembersProfileID uuid.UUID              `json:"membersProfileID"`
	Date             string                 `json:"date"`
	Description      string                 `json:"description"`
	Name             string                 `json:"name"`
	MembersProfile   *MemberProfileResource `json:"membersProfile,omitempty"`
}

func (m *ModelTransformer) MemberDescriptionToResource(description *MemberDescription) *MemberDescriptionResource {
	if description == nil {
		return nil
	}

	return &MemberDescriptionResource{

		ID:        description.ID,
		CreatedAt: description.CreatedAt.Format(time.RFC3339),
		UpdatedAt: description.UpdatedAt.Format(time.RFC3339),
		DeletedAt: description.DeletedAt.Time.Format(time.RFC3339),

		MembersProfileID: description.MembersProfileID,
		Date:             description.Date.Format("2006-01-02"),
		Description:      description.Description,
		Name:             description.Name,
		MembersProfile:   m.MemberProfileToResource(description.MembersProfile),
	}
}

func (m *ModelTransformer) MemberDescriptionToResourceList(descriptionList []*MemberDescription) []*MemberDescriptionResource {
	if descriptionList == nil {
		return nil
	}

	var descriptionResources []*MemberDescriptionResource
	for _, description := range descriptionList {
		descriptionResources = append(descriptionResources, m.MemberDescriptionToResource(description))
	}
	return descriptionResources
}
