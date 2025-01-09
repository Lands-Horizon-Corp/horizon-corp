package models

import "gorm.io/gorm"

type MemberEducationalAttainmentHistory struct {
	gorm.Model
	MemberProfileID               uint                         `gorm:"not null" json:"member_profile_id"`
	MemberEducationalAttainmentID uint                         `gorm:"type:bigint;unsigned;not null" json:"member_educational_attainment_id"`
	MemberProfile                 *MemberProfile               `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	MemberEducationalAttainment   *MemberEducationalAttainment `gorm:"foreignKey:MemberEducationalAttainmentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_educational_attainment"`
}
