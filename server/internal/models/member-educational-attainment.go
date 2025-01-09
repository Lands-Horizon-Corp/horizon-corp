package models

import "gorm.io/gorm"

type MemberEducationalAttainment struct {
	gorm.Model
	Name        string                                `gorm:"size:255;not null"`
	Description string                                `gorm:"size:500"`
	History     []*MemberEducationalAttainmentHistory `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

type MemberEducationalAttainmentResource struct {
	ID          uint                                          `json:"id"`
	CreatedAt   string                                        `json:"createdAt"`
	UpdatedAt   string                                        `json:"updatedAt"`
	Name        string                                        `json:"name"`
	Description string                                        `json:"description"`
	History     []*MemberEducationalAttainmentHistoryResource `json:"history,omitempty"`
}
