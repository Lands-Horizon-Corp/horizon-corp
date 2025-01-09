package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model

	// Fields
	FirstName          string     `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName           string     `gorm:"type:varchar(255);not null" json:"last_name"`
	MiddleName         string     `gorm:"type:varchar(255)" json:"middle_name"`
	PermanentAddress   string     `gorm:"type:text" json:"permanent_address"`
	Description        string     `gorm:"type:text" json:"description"`
	BirthDate          time.Time  `gorm:"type:date;not null" json:"birth_date"`
	Username           string     `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email              string     `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password           string     `gorm:"type:varchar(255);not null" json:"password"`
	IsEmailVerified    bool       `gorm:"default:false" json:"is_email_verified"`
	IsContactVerified  bool       `gorm:"default:false" json:"is_contact_verified"`
	IsSkipVerification bool       `gorm:"default:false" json:"is_skip_verification"`
	ContactNumber      string     `gorm:"type:varchar(255);unique;not null" json:"contact_number"`
	Status             UserStatus `gorm:"type:varchar(255);default:'Pending'" json:"status"`
	Longitude          *float64   `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude           *float64   `gorm:"type:decimal(10,7)" json:"latitude"`

	// Relationship 0 to 1
	MediaID *uint  `gorm:"type:bigint;unsigned" json:"media_id"`
	Media   *Media `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media"`

	// Relationship 0 to 1
	BranchID *uint   `gorm:"type:bigint;unsigned" json:"branch_id"`
	Branch   *Branch `gorm:"foreignKey:BranchID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"branch"`

	// Relationship 0 to 1
	RoleID *uint `gorm:"type:bigint;unsigned" json:"role_id"`
	Role   *Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`

	// Relationship 0 to 1
	GenderID *uint   `gorm:"type:bigint;unsigned" json:"gender_id"`
	Gender   *Gender `gorm:"foreignKey:GenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"gender"`

	// Relationship 0 to many
	Timesheets []*Timesheet `gorm:"foreignKey:EmployeeID" json:"timesheets"`
	Footsteps  []*Footstep  `gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"footsteps,omitempty"`
}

type EmployeeResource struct {
	AccountType string `json:"accountType"`

	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	FirstName          string               `json:"firstName"`
	LastName           string               `json:"lastName"`
	MiddleName         string               `json:"middleName"`
	PermanentAddress   string               `json:"permanentAddress"`
	Description        string               `json:"description"`
	BirthDate          time.Time            `json:"birthDate"`
	Username           string               `json:"username"`
	Email              string               `json:"email"`
	IsEmailVerified    bool                 `json:"isEmailVerified"`
	IsContactVerified  bool                 `json:"isContactVerified"`
	IsSkipVerification bool                 `json:"isSkipVerification"`
	ContactNumber      string               `json:"contactNumber"`
	Status             UserStatus           `json:"status"`
	Longitude          *float64             `json:"longitude"`
	Latitude           *float64             `json:"latitude"`
	MediaID            *uint                `json:"mediaID"`
	Media              *MediaResource       `json:"media"`
	BranchID           *uint                `json:"branchID"`
	Branch             *BranchResource      `json:"branch"`
	RoleID             *uint                `json:"roleID"`
	Role               *RoleResource        `json:"role"`
	GenderID           *uint                `json:"genderID"`
	Gender             *GenderResource      `json:"gender"`
	Timesheets         []*TimesheetResource `json:"timesheets"`
	Footsteps          []*FootstepResource  `json:"footsteps"`
}

func (m *ModelTransformer) EmployeeToResource(employee *Employee) *EmployeeResource {
	if employee == nil {
		return nil
	}

	return &EmployeeResource{
		AccountType:        "Employee", // Assuming the account type is always "Employee"
		ID:                 employee.ID,
		CreatedAt:          employee.CreatedAt.Format(time.RFC3339),
		UpdatedAt:          employee.UpdatedAt.Format(time.RFC3339),
		FirstName:          employee.FirstName,
		LastName:           employee.LastName,
		MiddleName:         employee.MiddleName,
		PermanentAddress:   employee.PermanentAddress,
		Description:        employee.Description,
		BirthDate:          employee.BirthDate,
		Username:           employee.Username,
		Email:              employee.Email,
		IsEmailVerified:    employee.IsEmailVerified,
		IsContactVerified:  employee.IsContactVerified,
		IsSkipVerification: employee.IsSkipVerification,
		ContactNumber:      employee.ContactNumber,
		Status:             employee.Status,
		Longitude:          employee.Longitude,
		Latitude:           employee.Latitude,
		MediaID:            employee.MediaID,
		Media:              m.MediaToResource(employee.Media),
		BranchID:           employee.BranchID,
		Branch:             m.BranchToResource(employee.Branch),
		RoleID:             employee.RoleID,
		Role:               m.RoleToResource(employee.Role),
		GenderID:           employee.GenderID,
		Gender:             m.GenderToResource(employee.Gender),
		Timesheets:         m.TimesheetToResourceList(employee.Timesheets),
		Footsteps:          m.FootstepToResourceList(employee.Footsteps),
	}
}

func (m *ModelTransformer) EmployeeToResourceList(employeeList []*Employee) []*EmployeeResource {
	if employeeList == nil {
		return nil
	}

	var employeeResources []*EmployeeResource
	for _, employee := range employeeList {
		employeeResources = append(employeeResources, m.EmployeeToResource(employee))
	}
	return employeeResources
}
