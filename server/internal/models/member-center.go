package models

import "gorm.io/gorm"

type MemberCenter struct {
	gorm.Model
	Name        string                 `gorm:"size:255;not null"`
	Description string                 `gorm:"size:500"`
	History     []*MemberCenterHistory `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

type MemberCenterResource struct {
	ID          uint                           `json:"id"`
	CreatedAt   string                         `json:"createdAt"`
	UpdatedAt   string                         `json:"updatedAt"`
	Name        string                         `json:"name"`
	Description string                         `json:"description"`
	History     []*MemberCenterHistoryResource `json:"history,omitempty"`
}
