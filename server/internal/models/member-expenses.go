package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberExpenses struct {
	gorm.Model
	MembersProfileID uint           `gorm:"not null" json:"members_profile_id"`
	Name             string         `gorm:"type:varchar(255);not null" json:"name"`
	Amount           float64        `gorm:"type:decimal(10,2);not null" json:"amount"`
	Date             time.Time      `gorm:"type:date" json:"date"`
	Description      string         `gorm:"type:text" json:"description"`
	MembersProfile   *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}
