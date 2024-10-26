package resources

import (
	"horizon/server/internal/models"
	"time"
)

type MemberResource struct {
	AccountType       string              `json:"account_type"`
	ID                uint                `json:"id"`
	FirstName         string              `json:"first_name"`
	LastName          string              `json:"last_name"`
	MiddleName        string              `json:"middle_name"`
	PermanentAddress  string              `json:"permanent_address"`
	Description       string              `json:"description"`
	Birthdate         time.Time           `json:"birthdate"`
	Username          string              `json:"username"`
	Email             string              `json:"email"`
	IsEmailVerified   bool                `json:"is_email_verified"`
	IsContactVerified bool                `json:"is_contact_verified"`
	Status            models.MemberStatus `json:"status"`
	ContactNumber     string              `json:"contact_number"`
	Longitude         *float64            `json:"longitude"`
	Latitude          *float64            `json:"latitude"`

	// Related entities
	Media  *MediaResource  `json:"media,omitempty"`
	Branch *BranchResource `json:"branch,omitempty"`
	Role   *RoleResource   `json:"role,omitempty"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToResourceMember(member models.Member) MemberResource {
	// Handle Media conversion
	var mediaResource *MediaResource
	if member.Media != nil {
		mediaRes := ToResourceMedia(*member.Media)
		mediaResource = &mediaRes
	}

	// Handle Branch conversion
	var branchResource *BranchResource
	if member.Branch != nil {
		branchRes := ToResourceBranch(*member.Branch)
		branchResource = &branchRes
	}

	// Handle Role conversion
	var roleResource *RoleResource
	if member.Role != nil {
		roleRes := ToResourceRole(*member.Role)
		roleResource = &roleRes
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
		Status:            member.Status,
		ContactNumber:     member.ContactNumber,
		Longitude:         member.Longitude,
		Latitude:          member.Latitude,
		Media:             mediaResource,
		Branch:            branchResource,
		Role:              roleResource,
		CreatedAt:         member.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         member.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResourceListMembers(members []models.Member) []MemberResource {
	resourceList := make([]MemberResource, len(members))
	for i, member := range members {
		resourceList[i] = ToResourceMember(member)
	}
	return resourceList
}
