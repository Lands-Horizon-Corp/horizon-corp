package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberWallet struct {
	gorm.Model
	MembersProfileID uint      `gorm:"not null" json:"members_profile_id"`
	Debit            float64   `gorm:"type:decimal(12,2);default:0" json:"debit"`
	Credit           float64   `gorm:"type:decimal(12,2);default:0" json:"credit"`
	Date             time.Time `gorm:"type:date" json:"date"`
	Description      string    `gorm:"type:text" json:"description"`

	MembersProfile *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}
