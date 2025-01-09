package models

import "gorm.io/gorm"

type MemberClassificationHistory struct {
	gorm.Model
	MemberProfileID        uint                  `gorm:"not null" json:"member_profile_id"`
	MemberClassificationID uint                  `gorm:"type:bigint;unsigned;not null" json:"member_classification_id"`
	MemberProfile          *MemberProfile        `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	MemberClassification   *MemberClassification `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member_classification"`
}

type MemberClassificationHistoryResource struct {
	ID                     uint                          `json:"id"`
	CreatedAt              string                        `json:"createdAt"`
	UpdatedAt              string                        `json:"updatedAt"`
	MemberProfileID        uint                          `json:"memberProfileID"`
	MemberClassificationID uint                          `json:"memberClassificationID"`
	MemberProfile          *MemberProfileResource        `json:"memberProfile,omitempty"`
	MemberClassification   *MemberClassificationResource `json:"memberClassification,omitempty"`
}
