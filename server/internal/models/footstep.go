package models

import (
	"gorm.io/gorm"
)

type Footstep struct {
	gorm.Model
	UserID      uint   `gorm:"not null" json:"user_id"`
	User        User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	AccountType string `gorm:"type:varchar(50);not null" json:"account_type"`
	Description string `gorm:"size:1000" json:"description"`
	Activity    string `gorm:"size:255;not null" json:"activity"`
}
