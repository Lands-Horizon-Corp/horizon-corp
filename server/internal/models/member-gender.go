package models

import "gorm.io/gorm"

type MemberGender struct {
	gorm.Model
	Name        string                 `gorm:"size:255;not null"`
	Description string                 `gorm:"size:500"`
	History     []*MemberGenderHistory `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}
