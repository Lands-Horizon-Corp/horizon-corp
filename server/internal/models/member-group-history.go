package models

import "gorm.io/gorm"

type MemberGroupHistory struct {
	gorm.Model
	MemberProfileID uint           `gorm:"not null" json:"member_profile_id"`
	MemberGroupID   uint           `gorm:"type:bigint;unsigned;not null" json:"member_group_id"`
	MemberProfile   *MemberProfile `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	MemberGroup     *MemberGroup   `gorm:"foreignKey:MemberGroupID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_group"`
}
