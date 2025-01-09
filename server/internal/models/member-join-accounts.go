package models

import "gorm.io/gorm"

type MemberJointAccounts struct {
	gorm.Model
	MembersProfileID   uint           `gorm:"not null" json:"members_profile_id"`
	Description        string         `gorm:"type:text" json:"description"`
	FirstName          string         `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName           string         `gorm:"type:varchar(255);not null" json:"last_name"`
	MiddleName         string         `gorm:"type:varchar(255)" json:"middle_name"`
	FamilyRelationship string         `gorm:"type:varchar(255)" json:"family_relationship"`
	MembersProfile     *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}
