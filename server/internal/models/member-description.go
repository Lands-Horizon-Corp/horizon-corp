package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberDescription struct {
	gorm.Model
	MembersProfileID uint           `gorm:"not null" json:"members_profile_id"`
	Date             time.Time      `gorm:"type:date" json:"date"`
	Description      string         `gorm:"type:text" json:"description"`
	Name             string         `gorm:"type:varchar(255);not null" json:"name"`
	MembersProfile   *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}
