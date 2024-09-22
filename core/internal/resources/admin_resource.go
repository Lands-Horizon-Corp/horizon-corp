package resources

import (
	"horizon-core/internal/models"
	"time"
)

type AdminResource struct {
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

func NewAdminResource(admin models.Admin) AdminResource {
	profilePicture := NewMediaResource(admin.ProfilePicture)

	roles := make([]RoleResource, len(admin.Roles))
	for i, role := range admin.Roles {
		roles[i] = NewRoleResource(role)
	}

	return AdminResource{
		ID:                 admin.ID,
		Email:              admin.Email,
		Username:           admin.Username,
		FirstName:          admin.FirstName,
		LastName:           admin.LastName,
		ContactNumber:      admin.ContactNumber,
		PermanentAddress:   admin.PermanentAddress,
		Description:        admin.Description,
		Birthdate:          admin.Birthdate,
		ValidEmail:         admin.ValidEmail,
		ValidContactNumber: admin.ValidContactNumber,
		ProfilePicture:     profilePicture,
		Roles:              roles,
	}
}
