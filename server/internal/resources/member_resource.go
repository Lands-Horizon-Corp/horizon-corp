package resources

import (
	"horizon/server/internal/models"
	"time"
)

type MemberResource struct {
	ID                uint          `json:"id"`
	FirstName         string        `json:"firstName"`
	LastName          string        `json:"lastName"`
	PermanentAddress  string        `json:"permanentAddress"`
	Description       string        `json:"description"`
	Birthdate         string        `json:"birthdate"`
	Username          string        `json:"username"`
	Email             string        `json:"email"`
	IsEmailVerified   bool          `json:"isEmailVerified"`
	IsContactVerified bool          `json:"isContactVerified"`
	Media             MediaResource `json:"media"`
}

func ToResourceMember(member models.Member) MemberResource {
	return MemberResource{
		ID:                member.ID,
		FirstName:         member.FirstName,
		LastName:          member.LastName,
		PermanentAddress:  member.PermanentAddress,
		Description:       member.Description,
		Birthdate:         member.Birthdate.Format(time.RFC3339),
		Username:          member.Username,
		Email:             member.Email,
		IsEmailVerified:   member.IsEmailVerified,
		IsContactVerified: member.IsContactVerified,
		Media:             ToResourceMedia(member.Media),
	}
}

func ToResourceListMember(memberList []models.Member) []MemberResource {
	var resources []MemberResource
	for _, member := range memberList {
		resources = append(resources, ToResourceMember(member))
	}
	return resources
}
