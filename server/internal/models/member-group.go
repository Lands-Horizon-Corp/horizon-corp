package models

import "gorm.io/gorm"

type MemberGroup struct {
	gorm.Model
	Name        string                `gorm:"size:255;not null"`
	Description string                `gorm:"size:500"`
	History     []*MemberGroupHistory `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

type MemberGroupResource struct {
	ID          uint                          `json:"id"`
	CreatedAt   string                        `json:"createdAt"`
	UpdatedAt   string                        `json:"updatedAt"`
	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	History     []*MemberGroupHistoryResource `json:"history,omitempty"`
}
