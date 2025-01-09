package models

import "gorm.io/gorm"

type MemberType struct {
	gorm.Model
	Name        string               `gorm:"size:255;not null"`
	Description string               `gorm:"size:500"`
	Prefix      string               `gorm:"size:100"`
	History     []*MemberTypeHistory `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}
