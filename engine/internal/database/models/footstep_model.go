package models

import (
	"gorm.io/gorm"
)

type Footstep struct {
	gorm.Model

	// Fields
	AdminID     *uint64 `gorm:"index" json:"admin_id,omitempty"`
	EmployeeID  *uint64 `gorm:"index" json:"employee_id,omitempty"`
	OwnerID     *uint64 `gorm:"index" json:"owner_id,omitempty"`
	MemberID    *uint64 `gorm:"index" json:"member_id,omitempty"`
	AccountType string  `gorm:"type:varchar(11);not null" json:"account_type"`
	Description string  `gorm:"type:varchar(1000)" json:"description,omitempty"`
	Activity    string  `gorm:"type:varchar(255);not null" json:"activity"`

	// Relationship 1 to 1
	Admin    Admin    `gorm:"foreignKey:AdminID;constraint:OnDelete:SET NULL,OnUpdate:CASCADE" json:"admin,omitempty"`
	Employee Employee `gorm:"foreignKey:EmployeeID;constraint:OnDelete:SET NULL,OnUpdate:CASCADE" json:"employee,omitempty"`
	Owner    Owner    `gorm:"foreignKey:OwnerID;constraint:OnDelete:SET NULL,OnUpdate:CASCADE" json:"owner,omitempty"`
	Member   Member   `gorm:"foreignKey:MemberID;constraint:OnDelete:SET NULL,OnUpdate:CASCADE" json:"member,omitempty"`
}
