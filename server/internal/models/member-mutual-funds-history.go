package models

import "gorm.io/gorm"

type MemberMutualFundsHistory struct {
	gorm.Model
	MembersProfileID uint           `gorm:"not null" json:"members_profile_id"`
	Description      string         `gorm:"type:text" json:"description"`
	Amount           float64        `gorm:"type:decimal(10,2);not null" json:"amount"`
	MembersProfile   *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}
