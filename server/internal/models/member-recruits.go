package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberRecruits struct {
	gorm.Model
	MembersProfileID          uint           `gorm:"not null" json:"members_profile_id"`
	MembersProfileRecruitedID uint           `gorm:"not null" json:"members_profile_recruited_id"`
	DateRecruited             time.Time      `gorm:"type:date" json:"date_recruited"`
	Description               string         `gorm:"type:text" json:"description"`
	Name                      string         `gorm:"type:varchar(255);not null" json:"name"`
	MembersProfile            *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
	MembersProfileRecruited   *MemberProfile `gorm:"foreignKey:MembersProfileRecruitedID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile_recruited"`
}
