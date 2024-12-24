package models

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator"
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

type EmployeeRequest struct {
	FirstName          string    `json:"firstName" validate:"required,max=255"`
	LastName           string    `json:"lastName" validate:"required,max=255"`
	MiddleName         string    `json:"middleName,omitempty" validate:"max=255"`
	PermanentAddress   string    `json:"permanentAddress,omitempty"`
	Description        string    `json:"description,omitempty"`
	BirthDate          time.Time `json:"birthDate" validate:"required"`
	Username           string    `json:"username" validate:"required,max=255"`
	Email              string    `json:"email" validate:"required,email,max=255"`
	IsEmailVerified    bool      `json:"isEmailVerified"`
	IsContactVerified  bool      `json:"isContactVerified"`
	IsSkipVerification bool      `json:"isSkipVerification"`
	ContactNumber      string    `json:"contactNumber" validate:"required,max=15"`
	Status             string    `json:"status" validate:"required,oneof=Pending Active Inactive"`
	Longitude          *float64  `json:"longitude" validate:"omitempty,longitude"`
	Latitude           *float64  `json:"latitude" validate:"omitempty,latitude"`
	MediaID            *uint     `json:"mediaID,omitempty"`
	BranchID           *uint     `json:"branchID,omitempty"`
	RoleID             *uint     `json:"roleID,omitempty"`
	GenderID           *uint     `json:"genderID,omitempty"`
}

func (m *ModelResource) EmployeeToResource(employee *Employee) *EmployeeResource {
	if employee == nil {
		return nil
	}

	return &EmployeeResource{
		AccountType: "Employee",

		ID:        employee.ID,
		CreatedAt: employee.CreatedAt.Format(time.RFC3339),
		UpdatedAt: employee.UpdatedAt.Format(time.RFC3339),

		FirstName:        employee.FirstName,
		LastName:         employee.LastName,
		MiddleName:       employee.MiddleName,
		PermanentAddress: employee.PermanentAddress,
		Description:      employee.Description,
		BirthDate:        employee.BirthDate,
		Username:         employee.Username,
		Email:            employee.Email,

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

func (m *ModelResource) EmployeeToResourceList(employees []*Employee) []*EmployeeResource {
	if employees == nil {
		return nil
	}
	var employeeResources []*EmployeeResource
	for _, employee := range employees {
		employeeResources = append(employeeResources, m.EmployeeToResource(employee))
	}
	return employeeResources
}

// EmployeeToRecord converts a slice of Employee pointers into CSV records and headers.
func (m *ModelResource) EmployeeToRecord(employees []*Employee) ([][]string, []string) {
	// Convert Employee structs to EmployeeResource structs
	resource := m.EmployeeToResourceList(employees)
	records := make([][]string, 0, len(resource))

	for _, employee := range resource {
		// Basic Fields
		id := strconv.Itoa(int(employee.ID))
		firstName := sanitizeCSVField(employee.FirstName)
		lastName := sanitizeCSVField(employee.LastName)
		middleName := sanitizeCSVField(employee.MiddleName)
		permanentAddress := sanitizeCSVField(employee.PermanentAddress)
		description := sanitizeCSVField(employee.Description)
		birthDate := employee.BirthDate.Format("2006-01-02") // Format as YYYY-MM-DD
		username := sanitizeCSVField(employee.Username)
		email := sanitizeCSVField(employee.Email)
		contactNumber := sanitizeCSVField(employee.ContactNumber)
		isEmailVerified := strconv.FormatBool(employee.IsEmailVerified)
		isContactVerified := strconv.FormatBool(employee.IsContactVerified)
		isSkipVerification := strconv.FormatBool(employee.IsSkipVerification)
		status := sanitizeCSVField(string(employee.Status))

		// Longitude and Latitude
		longitude := "N/A"
		if employee.Longitude != nil {
			longitude = fmt.Sprintf("%.6f", *employee.Longitude)
		}
		latitude := "N/A"
		if employee.Latitude != nil {
			latitude = fmt.Sprintf("%.6f", *employee.Latitude)
		}

		createdAt := sanitizeCSVField(employee.CreatedAt)
		updatedAt := sanitizeCSVField(employee.UpdatedAt)

		// Handle Media
		mediaURL := "N/A"
		if employee.Media != nil {
			mediaURL = sanitizeCSVField(employee.Media.URL)
		}

		// Handle Branch
		branchName := "N/A"
		if employee.Branch != nil {
			branchName = sanitizeCSVField(employee.Branch.Name)
		}

		// Handle Role
		roleName := "N/A"
		if employee.Role != nil {
			roleName = sanitizeCSVField(employee.Role.Name) // Assuming Role has a Name field
		}

		// Handle Gender
		genderName := "N/A"
		if employee.Gender != nil {
			genderName = sanitizeCSVField(employee.Gender.Name) // Assuming Gender has a Name field
		}

		// Handle Timesheets
		timesheets := "N/A"
		if len(employee.Timesheets) > 0 {
			tsEntries := make([]string, 0, len(employee.Timesheets))
			for _, ts := range employee.Timesheets {
				timeIn := "N/A"
				if ts.TimeIn != nil {
					timeIn = ts.TimeIn.Format(time.RFC3339)
				}
				timeOut := "N/A"
				if ts.TimeOut != nil {
					timeOut = ts.TimeOut.Format(time.RFC3339)
				}
				// Combine TimeIn and TimeOut
				tsEntry := fmt.Sprintf("%s - %s", timeIn, timeOut)
				tsEntries = append(tsEntries, sanitizeCSVField(tsEntry))
			}
			timesheets = strings.Join(tsEntries, "; ")
		}

		// Handle Footsteps
		footsteps := "N/A"
		if len(employee.Footsteps) > 0 {
			fsDescriptions := make([]string, 0, len(employee.Footsteps))
			for _, fs := range employee.Footsteps {
				// Assuming FootstepResource has a Description field
				fsDescriptions = append(fsDescriptions, sanitizeCSVField(fs.Description))
			}
			footsteps = strings.Join(fsDescriptions, "; ")
		}

		// Assemble the record
		record := []string{
			id,
			firstName,
			lastName,
			middleName,
			permanentAddress,
			description,
			birthDate,
			username,
			email,
			contactNumber,
			isEmailVerified,
			isContactVerified,
			isSkipVerification,
			status,
			longitude,
			latitude,
			mediaURL,
			branchName,
			roleName,
			genderName,
			timesheets,
			footsteps,
			createdAt,
			updatedAt,
		}
		records = append(records, record)
	}

	headers := []string{
		"ID",
		"First Name",
		"Last Name",
		"Middle Name",
		"Permanent Address",
		"Description",
		"Birth Date",
		"Username",
		"Email",
		"Contact Number",
		"Is Email Verified",
		"Is Contact Verified",
		"Is Skip Verification",
		"Status",
		"Longitude",
		"Latitude",
		"Media URL",
		"Branch Name",
		"Role",
		"Gender",
		"Timesheets",
		"Footsteps",
		"Created At",
		"Updated At",
	}

	return records, headers
}

func (m *ModelResource) ValidateEmployeeRequest(req *EmployeeRequest) error {
	validate := validator.New()
	validate.RegisterValidation("longitude", func(fl validator.FieldLevel) bool {
		lon := fl.Field().Float()
		return lon >= -180 && lon <= 180
	})
	validate.RegisterValidation("latitude", func(fl validator.FieldLevel) bool {
		lat := fl.Field().Float()
		return lat >= -90 && lat <= 90
	})
	err := validate.Struct(req)
	if err != nil {
		return m.helpers.FormatValidationError(err)
	}
	return nil
}

func (r *ModelResource) EmployeeGetByEmail(email string) (*Employee, error) {
	var employee Employee
	err := r.EmployeeDB.DB.Client.Preload("Media").Where("email = ?", email).First(&employee).Error
	return &employee, err
}

func (r *ModelResource) EmployeeGetByUsername(username string) (*Employee, error) {
	var employee Employee
	err := r.EmployeeDB.DB.Client.Preload("Media").Where("username = ?", username).First(&employee).Error
	return &employee, err
}

func (r *ModelResource) EmployeeGetByContactNumber(contactNumber string) (*Employee, error) {
	var employee Employee
	err := r.EmployeeDB.DB.Client.Preload("Media").Where("contact_number = ?", contactNumber).First(&employee).Error
	return &employee, err
}

func (r *ModelResource) EmployeeFindByEmailUsernameOrContact(input string) (*Employee, error) {
	switch r.helpers.GetKeyType(input) {
	case "contact":
		return r.EmployeeGetByContactNumber(input)
	case "email":
		return r.EmployeeGetByEmail(input)
	default:
		return r.EmployeeGetByUsername(input)
	}
}

func (r *ModelResource) EmployeeUpdatePassword(id uint, password string) error {
	newPassword, err := r.cryptoHelpers.HashPassword(password)
	if err != nil {
		return err
	}
	updated := &Employee{Password: newPassword}
	_, err = r.EmployeeDB.UpdateColumns(id, *updated, []string{"Media"})
	return err
}

func (r *ModelResource) EmployeeCreate(user *Employee) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}
	if user.Password == "" {
		return errors.New("password cannot be empty")
	}
	hashedPassword, err := r.cryptoHelpers.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return r.EmployeeDB.Create(user)
}

func (r *ModelResource) EmployeeUpdate(user *Employee, preloads []string) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}
	if user.Password != "" {
		hashedPassword, err := r.cryptoHelpers.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}
	return r.EmployeeDB.Update(user, preloads)
}

func (m *ModelResource) EmployeeSeeders() error {
	media, err := m.storage.UploadFromURL("https://s3.ap-southeast-2.amazonaws.com/horizon.assets/ecoop-logo.png")
	if err != nil {
		log.Printf("Error uploading ecoop url")
		return err
	}
	uploaded := &Media{
		FileName:   media.FileName,
		FileSize:   media.FileSize,
		FileType:   media.FileType,
		StorageKey: media.StorageKey,
		URL:        media.URL,
		BucketName: media.BucketName,
	}
	if err := m.MediaDB.Create(uploaded); err != nil {
		log.Printf("Error saving ecoop url")
		return err
	}
	employees := []Employee{
		{
			FirstName:         "Lands",
			LastName:          "Horizon",
			MiddleName:        "ecoop",
			PermanentAddress:  "123 Main St, Springfield",
			Description:       "Super administrator",
			BirthDate:         time.Date(1985, 7, 12, 0, 0, 0, 0, time.UTC),
			Username:          "alicej",
			Email:             m.cfg.EmployeeEmail,
			Password:          m.cfg.EmployeePassword,
			ContactNumber:     "5551234567",
			IsEmailVerified:   true,
			IsContactVerified: true,
			Status:            "Verified",
			MediaID:           &uploaded.ID,
		},
	}
	for _, employee := range employees {
		err := m.EmployeeCreate(&employee)
		if err != nil {
			log.Printf("Error seeding employee %s: %v", employee.Email, err)
		} else {
			log.Printf("employee %s seeded successfully", employee.Email)
		}
	}

	return nil
}
