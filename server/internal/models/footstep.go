package models

import (
	"gorm.io/gorm"
)

type Footstep struct {
	gorm.Model
	UserID      uint   `gorm:"type:bigint;unsigned;not null" json:"user_id"`
	AccountType string `gorm:"type:varchar(50);not null" json:"account_type"`
	Description string `gorm:"type:varchar(1000)" json:"description"`
	Activity    string `gorm:"type:varchar(255);not null" json:"activity"`
}
