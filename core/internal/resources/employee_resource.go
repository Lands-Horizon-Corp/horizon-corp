package resources

import (
	"horizon-core/internal/models"
	"time"
)

type EmployeeResource struct {
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

func NewEmployeeResource(employee models.Employee) EmployeeResource {
	profilePicture := NewMediaResource(employee.ProfilePicture)
	branch := NewBranchResource(employee.Branch)

	roles := make([]RoleResource, len(employee.Roles))
	for i, role := range employee.Roles {
		roles[i] = NewRoleResource(role)
	}

	return EmployeeResource{
		ID:                 employee.ID,
		Email:              employee.Email,
		Username:           employee.Username,
		FirstName:          employee.FirstName,
		LastName:           employee.LastName,
		ContactNumber:      employee.ContactNumber,
		PermanentAddress:   employee.PermanentAddress,
		Description:        employee.Description,
		Birthdate:          employee.Birthdate,
		ValidEmail:         employee.ValidEmail,
		ValidContactNumber: employee.ValidContactNumber,
		Branch:             &branch,
		ProfilePicture:     profilePicture,
		Roles:              roles,
	}
}
