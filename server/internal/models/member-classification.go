package models

import "gorm.io/gorm"

type MemberClassification struct {
	gorm.Model
	Name        string                         `gorm:"size:255;not null"`
	Description string                         `gorm:"size:500"`
	History     []*MemberClassificationHistory `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}
