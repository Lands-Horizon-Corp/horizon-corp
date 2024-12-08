package models

import (
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
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
	FirstName          string               `json:"firstName"`
	LastName           string               `json:"lastName"`
	MiddleName         string               `json:"middleName"`
	PermanentAddress   string               `json:"permanentAddress"`
	Description        string               `json:"description"`
	BirthDate          time.Time            `json:"birthDate"`
	Username           string               `json:"username"`
	Email              string               `json:"email"`
	Password           string               `json:"password"`
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

type (
	EmployeeResourceProvider interface {
		ToResource(employee *Employee) *EmployeeResource
		ToResourceList(employee []*Employee) []*EmployeeResource
	}
)

type EmployeeModel struct {
	lc             *fx.Lifecycle
	db             *gorm.DB
	logger         *zap.Logger
	mediaModel     MediaResourceProvider
	branchModel    BranchResourceProvider
	roleModel      RoleResourceProvider
	genderModel    GenderResourceProvider
	timesheetModel TimesheetResourceProvider
	footstepModel  FootstepResourceProvider
}

func NewEmployeeModel(
	lc *fx.Lifecycle,
	db *gorm.DB,
	logger *zap.Logger,
	mediaModel MediaResourceProvider,
	branchModel BranchResourceProvider,
	roleModel RoleResourceProvider,
	genderModel GenderResourceProvider,
	timesheetModel TimesheetResourceProvider,
	footstepModel FootstepResourceProvider,
) *EmployeeModel {
	return &EmployeeModel{
		lc:             lc,
		db:             db,
		logger:         logger,
		mediaModel:     mediaModel,
		branchModel:    branchModel,
		roleModel:      roleModel,
		genderModel:    genderModel,
		timesheetModel: timesheetModel,
		footstepModel:  footstepModel,
	}
}

func (em *EmployeeModel) SeedDatabase() {
}

func (em *EmployeeModel) ToResource(employee *Employee) *EmployeeResource {
	if employee == nil {
		return nil
	}

	return &EmployeeResource{
		FirstName:          employee.FirstName,
		LastName:           employee.LastName,
		MiddleName:         employee.MiddleName,
		PermanentAddress:   employee.PermanentAddress,
		Description:        employee.Description,
		BirthDate:          employee.BirthDate,
		Username:           employee.Username,
		Email:              employee.Email,
		Password:           employee.Password,
		IsEmailVerified:    employee.IsEmailVerified,
		IsContactVerified:  employee.IsContactVerified,
		IsSkipVerification: employee.IsSkipVerification,
		ContactNumber:      employee.ContactNumber,
		Status:             employee.Status,
		Longitude:          employee.Longitude,
		Latitude:           employee.Latitude,
		MediaID:            employee.MediaID,
		Media:              em.mediaModel.ToResource(employee.Media),
		BranchID:           employee.BranchID,
		Branch:             em.branchModel.ToResource(employee.Branch),
		RoleID:             employee.RoleID,
		Role:               em.roleModel.ToResource(employee.Role),
		GenderID:           employee.GenderID,
		Gender:             em.genderModel.ToResource(employee.Gender),
		Timesheets:         em.timesheetModel.ToResourceList(employee.Timesheets),
		Footsteps:          em.footstepModel.ToResourceList(employee.Footsteps),
	}
}

func (em *EmployeeModel) ToResourceList(employees []*Employee) []*EmployeeResource {
	if employees == nil {
		return nil
	}
	var employeeResources []*EmployeeResource
	for _, employee := range employees {
		employeeResources = append(employeeResources, em.ToResource(employee))
	}
	return employeeResources
}
