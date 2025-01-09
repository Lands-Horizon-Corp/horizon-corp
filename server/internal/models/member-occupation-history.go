package models

import "gorm.io/gorm"

type MemberOccupationHistory struct {
	gorm.Model
	MemberProfileID    uint              `gorm:"not null" json:"member_profile_id"`
	MemberOccupationID uint              `gorm:"type:bigint;unsigned;not null" json:"member_occupation_id"`
	MemberProfile      *MemberProfile    `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	MemberOccupation   *MemberOccupation `gorm:"foreignKey:MemberOccupationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_occupation"`
}

type MemberOccupationHistoryResource struct {
	ID                 uint                      `json:"id"`
	CreatedAt          string                    `json:"createdAt"`
	UpdatedAt          string                    `json:"updatedAt"`
	MemberProfileID    uint                      `json:"memberProfileID"`
	MemberOccupationID uint                      `json:"memberOccupationID"`
	MemberProfile      *MemberProfileResource    `json:"memberProfile,omitempty"`
	MemberOccupation   *MemberOccupationResource `json:"memberOccupation,omitempty"`
}
