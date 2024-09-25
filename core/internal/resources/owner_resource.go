package resources

import (
	"horizon-core/internal/models"
	"time"
)

type OwnerResource struct {
	ID                 string         `json:"id"`
	Email              string         `json:"email"`
	Username           string         `json:"username"`
	FirstName          string         `json:"first_name"`
	LastName           string         `json:"last_name"`
	ContactNumber      string         `json:"contact_number,omitempty"`
	PermanentAddress   string         `json:"permanent_address,omitempty"`
	Description        string         `json:"description,omitempty"`
	Birthdate          time.Time      `json:"birthdate"`
	ValidEmail         bool           `json:"valid_email"`
	ValidContactNumber bool           `json:"valid_contact_number"`
	ProfilePicture     *MediaResource `json:"profile_picture,omitempty"`
	Roles              []RoleResource `json:"roles,omitempty"`
}

func NewOwnerResource(owner models.Owner) OwnerResource {
	profilePicture := NewMediaResource(owner.ProfilePicture)

	roles := make([]RoleResource, len(owner.Roles))
	for i, role := range owner.Roles {
		roles[i] = NewRoleResource(role)
	}

	return OwnerResource{
		ID:                 owner.ID,
		Email:              owner.Email,
		Username:           owner.Username,
		FirstName:          owner.FirstName,
		LastName:           owner.LastName,
		ContactNumber:      owner.ContactNumber,
		PermanentAddress:   owner.PermanentAddress,
		Description:        owner.Description,
		Birthdate:          owner.Birthdate,
		ValidEmail:         owner.ValidEmail,
		ValidContactNumber: owner.ValidContactNumber,
		ProfilePicture:     profilePicture,
		Roles:              roles,
	}
}
