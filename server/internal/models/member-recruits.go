package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberRecruits struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	MembersProfileID          uuid.UUID      `gorm:"unsigned" json:"members_profile_id"`
	MembersProfileRecruitedID uuid.UUID      `gorm:"unsigned" json:"members_profile_recruited_id"`
	DateRecruited             time.Time      `gorm:"type:date" json:"date_recruited"`
	Description               string         `gorm:"type:text" json:"description"`
	Name                      string         `gorm:"type:varchar(255);unsigned" json:"name"`
	MembersProfile            *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
	MembersProfileRecruited   *MemberProfile `gorm:"foreignKey:MembersProfileRecruitedID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile_recruited"`
}

type MemberRecruitsResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	MembersProfileID          uuid.UUID              `json:"membersProfileID"`
	MembersProfileRecruitedID uuid.UUID              `json:"membersProfileRecruitedID"`
	DateRecruited             string                 `json:"dateRecruited"`
	Description               string                 `json:"description"`
	Name                      string                 `json:"name"`
	MembersProfile            *MemberProfileResource `json:"membersProfile,omitempty"`
	MembersProfileRecruited   *MemberProfileResource `json:"membersProfileRecruited,omitempty"`
}

func (m *ModelTransformer) MemberRecruitsToResource(recruit *MemberRecruits) *MemberRecruitsResource {
	if recruit == nil {
		return nil
	}

	return &MemberRecruitsResource{
		ID:        recruit.ID,
		CreatedAt: recruit.CreatedAt.Format(time.RFC3339),
		UpdatedAt: recruit.UpdatedAt.Format(time.RFC3339),
		DeletedAt: recruit.DeletedAt.Time.Format(time.RFC3339),

		MembersProfileID:          recruit.MembersProfileID,
		MembersProfileRecruitedID: recruit.MembersProfileRecruitedID,
		DateRecruited:             recruit.DateRecruited.Format("2006-01-02"),
		Description:               recruit.Description,
		Name:                      recruit.Name,
		MembersProfile:            m.MemberProfileToResource(recruit.MembersProfile),
		MembersProfileRecruited:   m.MemberProfileToResource(recruit.MembersProfileRecruited),
	}
}

func (m *ModelTransformer) MemberRecruitsToResourceList(recruitList []*MemberRecruits) []*MemberRecruitsResource {
	if recruitList == nil {
		return nil
	}

	var recruitResources []*MemberRecruitsResource
	for _, recruit := range recruitList {
		recruitResources = append(recruitResources, m.MemberRecruitsToResource(recruit))
	}
	return recruitResources
}
