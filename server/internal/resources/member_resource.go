package resources

import (
	"horizon/server/internal/models"
	"time"
)

type MemberResource struct {
	AccountType       string              `json:"accountType"`
	ID                uint                `json:"id"`
	FirstName         string              `json:"firstName"`
	LastName          string              `json:"lastName"`
	MiddleName        string              `json:"middleName"`
	PermanentAddress  string              `json:"permanentAddress"`
	Description       string              `json:"description"`
	Birthdate         time.Time           `json:"birthdate"`
	Username          string              `json:"username"`
	Email             string              `json:"email"`
	IsEmailVerified   bool                `json:"isEmailVerified"`
	IsContactVerified bool                `json:"isContactVerified"`
	ContactNumber     string              `json:"contactNumber"`
	Media             *MediaResource      `json:"media,omitempty"`
	Status            models.MemberStatus `json:"status"`
	CreatedAt         string              `json:"createdAt"`
	UpdatedAt         string              `json:"updatedAt"`
}

func ToResourceMember(member models.Member) MemberResource {
	var mediaResource *MediaResource
	if member.MediaID != nil {
		media := ToResourceMedia(member.Media)
		mediaResource = &media
	}
	return MemberResource{
		AccountType:       "Member",
		ID:                member.ID,
		FirstName:         member.FirstName,
		LastName:          member.LastName,
		MiddleName:        member.MiddleName,
		PermanentAddress:  member.PermanentAddress,
		Description:       member.Description,
		Birthdate:         member.Birthdate,
		Username:          member.Username,
		Email:             member.Email,
		IsEmailVerified:   member.IsEmailVerified,
		IsContactVerified: member.IsContactVerified,
		ContactNumber:     member.ContactNumber,
		Media:             mediaResource,
		Status:            member.Status,
		CreatedAt:         member.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         member.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResourceListMembers(members []models.Member) []MemberResource {
	var resources []MemberResource
	for _, member := range members {
		resources = append(resources, ToResourceMember(member))
	}
	return resources
}
