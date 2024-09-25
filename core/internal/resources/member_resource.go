package resources

import (
	"horizon-core/internal/models"
	"time"
)

type MemberResource struct {
	ID                 string          `json:"id"`
	Email              string          `json:"email"`
	Username           string          `json:"username"`
	FirstName          string          `json:"first_name"`
	LastName           string          `json:"last_name"`
	ContactNumber      string          `json:"contact_number,omitempty"`
	PermanentAddress   string          `json:"permanent_address,omitempty"`
	Description        string          `json:"description,omitempty"`
	Birthdate          time.Time       `json:"birthdate"`
	ValidEmail         bool            `json:"valid_email"`
	ValidContactNumber bool            `json:"valid_contact_number"`
	Branch             *BranchResource `json:"branch,omitempty"`
	ProfilePicture     *MediaResource  `json:"profile_picture,omitempty"`
	Roles              []RoleResource  `json:"roles,omitempty"`
}

func NewMemberResource(member models.Member) MemberResource {
	profilePicture := NewMediaResource(member.ProfilePicture)

	var branch *BranchResource
	if member.Branch != nil {
		b := NewBranchResource(*member.Branch)
		branch = &b
	}

	roles := make([]RoleResource, len(member.Roles))
	for i, role := range member.Roles {
		roles[i] = NewRoleResource(role)
	}

	return MemberResource{
		ID:                 member.ID,
		Email:              member.Email,
		Username:           member.Username,
		FirstName:          member.FirstName,
		LastName:           member.LastName,
		ContactNumber:      member.ContactNumber,
		PermanentAddress:   member.PermanentAddress,
		Description:        member.Description,
		Birthdate:          member.Birthdate,
		ValidEmail:         member.ValidEmail,
		ValidContactNumber: member.ValidContactNumber,
		Branch:             branch,
		ProfilePicture:     profilePicture,
		Roles:              roles,
	}
}
