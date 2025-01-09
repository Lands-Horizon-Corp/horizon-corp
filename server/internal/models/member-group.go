package models

import "gorm.io/gorm"

type MemberGroup struct {
	gorm.Model
	Name        string                `gorm:"size:255;not null"`
	Description string                `gorm:"size:500"`
	History     []*MemberGroupHistory `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}
