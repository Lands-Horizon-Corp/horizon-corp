package models

import "gorm.io/gorm"

type MemberGenderHistory struct {
	gorm.Model
	MemberProfileID uint           `gorm:"not null" json:"member_profile_id"`
	MemberGenderID  uint           `gorm:"type:bigint;unsigned;not null" json:"member_gender_id"`
	MemberProfile   *MemberProfile `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	MemberGender    *MemberGender  `gorm:"foreignKey:MemberGenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_gender"`
}

type MemberGenderHistoryResource struct {
	ID              uint                   `json:"id"`
	CreatedAt       string                 `json:"createdAt"`
	UpdatedAt       string                 `json:"updatedAt"`
	MemberProfileID uint                   `json:"memberProfileID"`
	MemberGenderID  uint                   `json:"memberGenderID"`
	MemberProfile   *MemberProfileResource `json:"memberProfile,omitempty"`
	MemberGender    *MemberGenderResource  `json:"memberGender,omitempty"`
}
