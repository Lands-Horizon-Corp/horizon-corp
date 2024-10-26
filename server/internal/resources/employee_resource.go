package resources

import (
	"horizon/server/internal/models"
	"time"
)

type EmployeeResource struct {
	AccountType        string                `json:"account_type"`
	ID                 uint                  `json:"id"`
	FirstName          string                `json:"first_name"`
	LastName           string                `json:"last_name"`
	MiddleName         string                `json:"middle_name,omitempty"`
	PermanentAddress   string                `json:"permanent_address,omitempty"`
	Description        string                `json:"description,omitempty"`
	Birthdate          time.Time             `json:"birthdate"`
	Username           string                `json:"username"`
	Email              string                `json:"email"`
	ContactNumber      string                `json:"contact_number"`
	IsEmailVerified    bool                  `json:"is_email_verified"`
	IsContactVerified  bool                  `json:"is_contact_verified"`
	IsSkipVerification bool                  `json:"is_skip_verification"`
	Status             models.EmployeeStatus `json:"status"`
	Media              *MediaResource        `json:"media,omitempty"`
	Branch             *BranchResource       `json:"branch,omitempty"`
	Longitude          *float64              `json:"longitude,omitempty"`
	Latitude           *float64              `json:"latitude,omitempty"`
	Timesheets         []TimesheetResource   `json:"timesheets,omitempty"`
	Role               *RoleResource         `json:"role,omitempty"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToResourceEmployee(employee models.Employee) EmployeeResource {
	var mediaResource *MediaResource
	if employee.Media != nil {
		mediaRes := ToResourceMedia(*employee.Media)
		mediaResource = &mediaRes
	}

	var branchResource *BranchResource
	if employee.Branch != nil {
		branchRes := ToResourceBranch(*employee.Branch)
		branchResource = &branchRes
	}

	var roleResource *RoleResource
	if employee.Role != nil {
		roleRes := ToResourceRole(*employee.Role)
		roleResource = &roleRes
	}

	return EmployeeResource{
		AccountType:        "Employee",
		ID:                 employee.ID,
		FirstName:          employee.FirstName,
		LastName:           employee.LastName,
		MiddleName:         employee.MiddleName,
		PermanentAddress:   employee.PermanentAddress,
		Description:        employee.Description,
		Birthdate:          employee.Birthdate,
		Username:           employee.Username,
		Email:              employee.Email,
		ContactNumber:      employee.ContactNumber,
		IsEmailVerified:    employee.IsEmailVerified,
		IsContactVerified:  employee.IsContactVerified,
		IsSkipVerification: employee.IsSkipVerification,
		Status:             employee.Status,
		Media:              mediaResource,
		Branch:             branchResource,
		Longitude:          employee.Longitude,
		Latitude:           employee.Latitude,
		Timesheets:         ToResourceListTimesheets(employee.Timesheets),
		Role:               roleResource,

		CreatedAt: employee.CreatedAt.Format(time.RFC3339),
		UpdatedAt: employee.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResourceListEmployees(employees []models.Employee) []EmployeeResource {
	var resources []EmployeeResource
	for _, employee := range employees {
		resources = append(resources, ToResourceEmployee(employee))
	}
	return resources
}
