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
	Birthdate         time.Time           `json:"birthDate"`
	Username          string              `json:"username"`
	Email             string              `json:"email"`
	IsEmailVerified   bool                `json:"isEmailVerified"`
	IsContactVerified bool                `json:"isContactVerified"`
	Status            models.MemberStatus `json:"status"`
	ContactNumber     string              `json:"contactNumber"`
	Longitude         *float64            `json:"longitude"`
	Latitude          *float64            `json:"latitude"`

	// Related entities
	Media     *MediaResource      `json:"media,omitempty"`
	Branch    *BranchResource     `json:"branch,omitempty"`
	Role      *RoleResource       `json:"role,omitempty"`
	GenderID  *uint               `json:"genderId,omitempty"`
	Gender    *GenderResource     `json:"gender,omitempty"`
	Footsteps []*FootstepResource `json:"footsteps,omitempty"` // Updated to slice of pointers

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Convert models.Member to *MemberResource
func ToResourceMember(member *models.Member) *MemberResource {
	if member == nil {
		return nil
	}

	// Convert Media
	var mediaResource *MediaResource
	if member.Media != nil {
		mediaResource = ToResourceMedia(member.Media)
	}

	// Convert Branch
	var branchResource *BranchResource
	if member.Branch != nil {
		branchResource = ToResourceBranch(member.Branch)
	}

	// Convert Role
	var roleResource *RoleResource
	if member.Role != nil {
		roleResource = ToResourceRole(member.Role)
	}

	// Convert Gender
	var genderResource *GenderResource
	if member.Gender != nil {
		genderResource = ToResourceGender(member.Gender)
	}

	return &MemberResource{
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
		GenderID:          member.GenderID,
		Gender:            genderResource,
		Footsteps:         ToResourceListFootsteps(member.Footsteps),
		CreatedAt:         member.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         member.UpdatedAt.Format(time.RFC3339),
	}
}

// Convert []*models.Member to []*MemberResource
func ToResourceListMembers(members []*models.Member) []*MemberResource {
	resourceList := make([]*MemberResource, len(members))
	for i, member := range members {
		resourceList[i] = ToResourceMember(member)
	}
	return resourceList
}
