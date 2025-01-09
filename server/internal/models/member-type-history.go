package models

import "gorm.io/gorm"

type MemberTypeHistory struct {
	gorm.Model
	MemberProfileID uint           `gorm:"not null" json:"member_profile_id"`
	MemberTypeID    uint           `gorm:"type:bigint;unsigned;not null" json:"member_type_id"`
	MemberProfile   *MemberProfile `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	MemberType      *MemberType    `gorm:"foreignKey:MemberTypeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_type"`
}
