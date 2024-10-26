package resources

import (
	"horizon/server/internal/models"
	"time"
)

type EmployeeResource struct {
	AccountType        string                `json:"accountType"`
	ID                 uint                  `json:"id"`
	FirstName          string                `json:"firstName"`
	LastName           string                `json:"lastName"`
	MiddleName         string                `json:"middleName,omitempty"`
	PermanentAddress   string                `json:"permanentAddress,omitempty"`
	Description        string                `json:"description,omitempty"`
	Birthdate          time.Time             `json:"birthdate"`
	Username           string                `json:"username"`
	Email              string                `json:"email"`
	ContactNumber      string                `json:"contactNumber"`
	IsEmailVerified    bool                  `json:"isEmailVerified"`
	IsContactVerified  bool                  `json:"isContactVerified"`
	IsSkipVerification bool                  `json:"isSkipVerification"`
	Status             models.EmployeeStatus `json:"status"`
	Media              *MediaResource        `json:"media,omitempty"`
	Branch             *BranchResource       `json:"branch,omitempty"`
	Longitude          *float64              `json:"longitude,omitempty"`
	Latitude           *float64              `json:"latitude,omitempty"`
	Timesheets         []*TimesheetResource  `json:"timesheets,omitempty"` // Updated to slice of pointers
	Role               *RoleResource         `json:"role,omitempty"`
	GenderID           *uint                 `json:"genderId,omitempty"`
	Gender             *GenderResource       `json:"gender,omitempty"`
	Footsteps          []*FootstepResource   `json:"footsteps,omitempty"` // Updated to slice of pointers

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Convert models.Employee to *EmployeeResource
func ToResourceEmployee(employee *models.Employee) *EmployeeResource {
	if employee == nil {
		return nil
	}

	// Convert Media
	var mediaResource *MediaResource
	if employee.Media != nil {
		mediaResource = ToResourceMedia(employee.Media)
	}

	// Convert Branch
	var branchResource *BranchResource
	if employee.Branch != nil {
		branchResource = ToResourceBranch(employee.Branch)
	}

	// Convert Role
	var roleResource *RoleResource
	if employee.Role != nil {
		roleResource = ToResourceRole(employee.Role)
	}

	// Convert Gender
	var genderResource *GenderResource
	if employee.Gender != nil {
		genderResource = ToResourceGender(employee.Gender)
	}

	return &EmployeeResource{
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
		Footsteps:          ToResourceListFootsteps(employee.Footsteps),
		Role:               roleResource,
		GenderID:           employee.GenderID,
		Gender:             genderResource,
		CreatedAt:          employee.CreatedAt.Format(time.RFC3339),
		UpdatedAt:          employee.UpdatedAt.Format(time.RFC3339),
	}
}

// Convert []models.Employee to []*EmployeeResource
func ToResourceListEmployees(employees []*models.Employee) []*EmployeeResource {
	var resources []*EmployeeResource
	for _, employee := range employees {
		resources = append(resources, ToResourceEmployee(employee))
	}
	return resources
}
