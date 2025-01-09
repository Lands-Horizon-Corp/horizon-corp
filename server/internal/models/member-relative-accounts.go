package models

import "gorm.io/gorm"

type MemberRelativeAccounts struct {
	gorm.Model
	MembersProfileID             uint           `gorm:"not null" json:"members_profile_id"`
	RelativeProfileMemberID      uint           `gorm:"not null" json:"relative_member_id"`
	FamilyRelationship           string         `gorm:"type:varchar(255)" json:"family_relationship"`
	Description                  string         `gorm:"type:text" json:"description"`
	MemberProfile                *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member_profile"`
	RelativeProfileMemberProfile *MemberProfile `gorm:"foreignKey:RelativeProfileMemberID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"relative_member_profile"`
}

type MemberRelativeAccountsResource struct {
	ID                           uint                   `json:"id"`
	CreatedAt                    string                 `json:"createdAt"`
	UpdatedAt                    string                 `json:"updatedAt"`
	MembersProfileID             uint                   `json:"membersProfileID"`
	RelativeProfileMemberID      uint                   `json:"relativeProfileMemberID"`
	FamilyRelationship           string                 `json:"familyRelationship"`
	Description                  string                 `json:"description"`
	MemberProfile                *MemberProfileResource `json:"memberProfile,omitempty"`
	RelativeProfileMemberProfile *MemberProfileResource `json:"relativeProfileMemberProfile,omitempty"`
}
