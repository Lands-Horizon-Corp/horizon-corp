package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Fields
	FirstName          string     `gorm:"type:varchar(255);unsigned" json:"first_name"`
	LastName           string     `gorm:"type:varchar(255);unsigned" json:"last_name"`
	MiddleName         string     `gorm:"type:varchar(255)" json:"middle_name"`
	PermanentAddress   string     `gorm:"type:text" json:"permanent_address"`
	Description        string     `gorm:"type:text" json:"description"`
	BirthDate          time.Time  `gorm:"type:date;unsigned" json:"birth_date"`
	Username           string     `gorm:"type:varchar(255);unique;unsigned" json:"username"`
	Email              string     `gorm:"type:varchar(255);unique;unsigned" json:"email"`
	Password           string     `gorm:"type:varchar(255);unsigned" json:"password"`
	IsEmailVerified    bool       `gorm:"default:false" json:"is_email_verified"`
	IsContactVerified  bool       `gorm:"default:false" json:"is_contact_verified"`
	IsSkipVerification bool       `gorm:"default:false" json:"is_skip_verification"`
	ContactNumber      string     `gorm:"type:varchar(255);unique;unsigned" json:"contact_number"`
	Status             UserStatus `gorm:"type:varchar(255);default:'Pending'" json:"status"`
	Longitude          *float64   `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude           *float64   `gorm:"type:decimal(10,7)" json:"latitude"`

	// Relationship 0 to 1
	MediaID *uuid.UUID `gorm:"type:bigint;unsigned" json:"media_id"`
	Media   *Media     `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media"`

	// Relationship 0 to 1
	BranchID *uuid.UUID `gorm:"type:bigint;unsigned" json:"branch_id"`
	Branch   *Branch    `gorm:"foreignKey:BranchID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"branch"`

	// Relationship 0 to 1
	RoleID *uuid.UUID `gorm:"type:bigint;unsigned" json:"role_id"`
	Role   *Role      `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`

	// Relationship 0 to 1
	GenderID *uuid.UUID `gorm:"type:bigint;unsigned" json:"gender_id"`
	Gender   *Gender    `gorm:"foreignKey:GenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"gender"`

	// Relationship 0 to many
	Timesheets []*Timesheet `gorm:"foreignKey:EmployeeID" json:"timesheets"`
	Footsteps  []*Footstep  `gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"footsteps,omitempty"`

	// Reverse relationship
	VerifiedProfiles []*MemberProfile `gorm:"foreignKey:VerifiedByEmployeeID" json:"verified_profiles,omitempty"`
}

func (v *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type EmployeeResource struct {
	AccountType string `json:"accountType"`

	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	FirstName          string                   `json:"firstName"`
	LastName           string                   `json:"lastName"`
	MiddleName         string                   `json:"middleName"`
	PermanentAddress   string                   `json:"permanentAddress"`
	Description        string                   `json:"description"`
	BirthDate          time.Time                `json:"birthDate"`
	Username           string                   `json:"username"`
	Email              string                   `json:"email"`
	IsEmailVerified    bool                     `json:"isEmailVerified"`
	IsContactVerified  bool                     `json:"isContactVerified"`
	IsSkipVerification bool                     `json:"isSkipVerification"`
	ContactNumber      string                   `json:"contactNumber"`
	Status             UserStatus               `json:"status"`
	Longitude          *float64                 `json:"longitude"`
	Latitude           *float64                 `json:"latitude"`
	MediaID            *uuid.UUID               `json:"mediaID"`
	Media              *MediaResource           `json:"media"`
	BranchID           *uuid.UUID               `json:"branchID"`
	Branch             *BranchResource          `json:"branch"`
	RoleID             *uuid.UUID               `json:"roleID"`
	Role               *RoleResource            `json:"role"`
	GenderID           *uuid.UUID               `json:"genderID"`
	Gender             *GenderResource          `json:"gender"`
	Timesheets         []*TimesheetResource     `json:"timesheets"`
	Footsteps          []*FootstepResource      `json:"footsteps"`
	VerifiedProfiles   []*MemberProfileResource `json:"verifiedProfiles,omitempty"`
}

func (m *ModelTransformer) EmployeeToResource(employee *Employee) *EmployeeResource {
	if employee == nil {
		return nil
	}

	return &EmployeeResource{
		AccountType: "Employee",

		ID:        employee.ID,
		CreatedAt: employee.CreatedAt.Format(time.RFC3339),
		UpdatedAt: employee.UpdatedAt.Format(time.RFC3339),
		DeletedAt: employee.DeletedAt.Time.Format(time.RFC3339),

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
		VerifiedProfiles:   m.MemberProfileToResourceList(employee.VerifiedProfiles),
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
