package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberAssets struct {
	gorm.Model
	MembersProfileID uint           `gorm:"not null" json:"members_profile_id"`
	EntryDate        time.Time      `gorm:"type:date;not null" json:"entry_date"`
	Description      string         `gorm:"type:text" json:"description"`
	Name             string         `gorm:"type:varchar(255);not null" json:"name"`
	MembersProfile   *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

type MemberAssetsResource struct {
	ID               uint                   `json:"id"`
	CreatedAt        string                 `json:"createdAt"`
	UpdatedAt        string                 `json:"updatedAt"`
	MembersProfileID uint                   `json:"membersProfileID"`
	EntryDate        string                 `json:"entryDate"`
	Description      string                 `json:"description"`
	Name             string                 `json:"name"`
	MembersProfile   *MemberProfileResource `json:"membersProfile,omitempty"`
}
