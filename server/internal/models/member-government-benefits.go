package models

import "gorm.io/gorm"

type MemberGovernmentBenefits struct {
	gorm.Model
	MembersProfileID uint           `gorm:"not null" json:"members_profile_id"`
	Country          string         `gorm:"type:varchar(255);not null" json:"country"`
	Name             string         `gorm:"type:varchar(255);not null" json:"name"`
	Description      string         `gorm:"type:text" json:"description"`
	Value            float64        `gorm:"type:decimal(10,2)" json:"value"`
	FrontMediaID     *uint          `gorm:"type:bigint;unsigned" json:"front_media_id"`
	BackMediaID      *uint          `gorm:"type:bigint;unsigned" json:"back_media_id"`
	MembersProfile   *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
	FrontMedia       *Media         `gorm:"foreignKey:FrontMediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"front_media,omitempty"`
	BackMedia        *Media         `gorm:"foreignKey:BackMediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"back_media,omitempty"`
}

type MemberGovernmentBenefitsResource struct {
	ID               uint                   `json:"id"`
	CreatedAt        string                 `json:"createdAt"`
	UpdatedAt        string                 `json:"updatedAt"`
	MembersProfileID uint                   `json:"membersProfileID"`
	Country          string                 `json:"country"`
	Name             string                 `json:"name"`
	Description      string                 `json:"description"`
	Value            float64                `json:"value"`
	FrontMediaID     *uint                  `json:"frontMediaID,omitempty"`
	BackMediaID      *uint                  `json:"backMediaID,omitempty"`
	MembersProfile   *MemberProfileResource `json:"membersProfile,omitempty"`
	FrontMedia       *MediaResource         `json:"frontMedia,omitempty"`
	BackMedia        *MediaResource         `json:"backMedia,omitempty"`
}
