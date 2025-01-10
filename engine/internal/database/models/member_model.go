package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model

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

	// Relationship 0 to 1
	MediaID *uint  `gorm:"type:bigint;unsigned" json:"media_id"`
	Media   *Media `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media"`

	// Relationship 0 to 1
	BranchID *uint   `gorm:"type:bigint;unsigned" json:"branch_id"`
	Branch   *Branch `gorm:"foreignKey:BranchID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"branch"`

	// Relationship 0 to 1
	Longitude *float64 `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude  *float64 `gorm:"type:decimal(10,7)" json:"latitude"`

	// Relationship 0 to 1
	RoleID *uint `gorm:"type:bigint;unsigned" json:"role_id"`
	Role   *Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`

	// Relationship 0 to 1
	GenderID *uint   `gorm:"type:bigint;unsigned" json:"gender_id"`
	Gender   *Gender `gorm:"foreignKey:GenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"gender"`

	// Relationship 0 to many
	Footsteps []*Footstep `gorm:"foreignKey:MemberID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"footsteps,omitempty"`
}

type MemberResource struct {
	AccountType string `json:"accountType"`

	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	FirstName          string          `json:"firstName"`
	LastName           string          `json:"lastName"`
	MiddleName         string          `json:"middleName"`
	PermanentAddress   string          `json:"permanentAddress"`
	Description        string          `json:"description"`
	BirthDate          time.Time       `json:"birthDate"`
	Username           string          `json:"username"`
	Email              string          `json:"email"`
	IsEmailVerified    bool            `json:"isEmailVerified"`
	IsContactVerified  bool            `json:"isContactVerified"`
	IsSkipVerification bool            `json:"isSkipVerification"`
	ContactNumber      string          `json:"contactNumber"`
	Status             UserStatus      `json:"status"`
	MediaID            *uint           `json:"mediaID"`
	Media              *MediaResource  `json:"media"`
	BranchID           *uint           `json:"branchID"`
	Branch             *BranchResource `json:"branch"`
	Longitude          *float64        `json:"longitude"`
	Latitude           *float64        `json:"latitude"`
	RoleID             *uint           `json:"roleID"`
	Role               *RoleResource   `json:"role"`
	GenderID           *uint           `json:"genderID"`
	Gender             *GenderResource `json:"gender"`

	Footsteps []*FootstepResource `json:"footsteps"`
}

func (m *ModelResource) MemberToResource(member *Member) *MemberResource {
	if member == nil {
		return nil
	}
	return &MemberResource{

		AccountType: "Member",

		ID:        member.ID,
		CreatedAt: member.CreatedAt.Format(time.RFC3339),
		UpdatedAt: member.UpdatedAt.Format(time.RFC3339),

		FirstName:          member.FirstName,
		LastName:           member.LastName,
		MiddleName:         member.MiddleName,
		PermanentAddress:   member.PermanentAddress,
		Description:        member.Description,
		BirthDate:          member.BirthDate,
		Username:           member.Username,
		Email:              member.Email,
		IsEmailVerified:    member.IsEmailVerified,
		IsContactVerified:  member.IsContactVerified,
		IsSkipVerification: member.IsSkipVerification,
		ContactNumber:      member.ContactNumber,
		Status:             member.Status,
		MediaID:            member.MediaID,
		Media:              m.MediaToResource(member.Media),
		BranchID:           member.BranchID,
		Branch:             m.BranchToResource(member.Branch),
		Longitude:          member.Longitude,
		Latitude:           member.Latitude,
		RoleID:             member.RoleID,
		Role:               m.RoleToResource(member.Role),
		GenderID:           member.GenderID,
		Gender:             m.GenderToResource(member.Gender),
		Footsteps:          m.FootstepToResourceList(member.Footsteps),
	}
}

type MemberRequest struct {
	FirstName          string    `json:"firstName" validate:"required,max=255"`
	LastName           string    `json:"lastName" validate:"required,max=255"`
	MiddleName         string    `json:"middleName,omitempty" validate:"max=255"`
	PermanentAddress   string    `json:"permanentAddress,omitempty"`
	Description        string    `json:"description,omitempty"`
	BirthDate          time.Time `json:"birthDate" validate:"required"`
	Username           string    `json:"username" validate:"required,max=255"`
	Email              string    `json:"email" validate:"required,email,max=255"`
	Password           string    `json:"password" validate:"required,min=8,max=255"`
	IsEmailVerified    bool      `json:"isEmailVerified"`
	IsContactVerified  bool      `json:"isContactVerified"`
	IsSkipVerification bool      `json:"isSkipVerification"`
	ContactNumber      string    `json:"contactNumber" validate:"required,max=255"`
	Status             string    `json:"status" validate:"required,oneof=Pending Active Inactive"`
	MediaID            *uint     `json:"mediaID,omitempty"`
	BranchID           *uint     `json:"branchID,omitempty"`
	Longitude          *float64  `json:"longitude" validate:"omitempty,longitude"`
	Latitude           *float64  `json:"latitude" validate:"omitempty,latitude"`
	RoleID             *uint     `json:"roleID,omitempty"`
	GenderID           *uint     `json:"genderID,omitempty"`
}

func (m *ModelResource) MemberToResourceList(members []*Member) []*MemberResource {
	if members == nil {
		return nil
	}
	var memberResources []*MemberResource
	for _, member := range members {
		memberResources = append(memberResources, m.MemberToResource(member))
	}
	return memberResources
}

// MemberToRecord converts a slice of Member pointers into CSV records and headers.
func (m *ModelResource) MemberToRecord(members []*Member) ([][]string, []string) {
	// Convert Member structs to MemberResource structs
	resources := m.MemberToResourceList(members)
	records := make([][]string, 0, len(resources))

	for _, member := range resources {
		// Basic Fields
		id := strconv.Itoa(int(member.ID))
		accountType := sanitizeCSVField(member.AccountType)
		firstName := sanitizeCSVField(member.FirstName)
		lastName := sanitizeCSVField(member.LastName)
		middleName := sanitizeCSVField(member.MiddleName)
		permanentAddress := sanitizeCSVField(member.PermanentAddress)
		description := sanitizeCSVField(member.Description)
		birthDate := member.BirthDate.Format("2006-01-02") // Format as YYYY-MM-DD
		username := sanitizeCSVField(member.Username)
		email := sanitizeCSVField(member.Email)
		contactNumber := sanitizeCSVField(member.ContactNumber)
		isEmailVerified := strconv.FormatBool(member.IsEmailVerified)
		isContactVerified := strconv.FormatBool(member.IsContactVerified)
		isSkipVerification := strconv.FormatBool(member.IsSkipVerification)
		status := sanitizeCSVField(string(member.Status))

		// Longitude and Latitude
		longitude := "N/A"
		if member.Longitude != nil {
			longitude = fmt.Sprintf("%.6f", *member.Longitude)
		}
		latitude := "N/A"
		if member.Latitude != nil {
			latitude = fmt.Sprintf("%.6f", *member.Latitude)
		}

		createdAt := sanitizeCSVField(member.CreatedAt)
		updatedAt := sanitizeCSVField(member.UpdatedAt)

		// Handle Media
		mediaURL := "N/A"
		if member.Media != nil {
			mediaURL = sanitizeCSVField(member.Media.URL)
		}

		// Handle Branch
		branchName := "N/A"
		if member.Branch != nil {
			branchName = sanitizeCSVField(member.Branch.Name)
		}

		// Handle Role
		roleName := "N/A"
		if member.Role != nil {
			roleName = sanitizeCSVField(member.Role.Name) // Assuming Role has a Name field
		}

		// Handle Gender
		genderName := "N/A"
		if member.Gender != nil {
			genderName = sanitizeCSVField(member.Gender.Name) // Assuming Gender has a Name field
		}

		// Handle Footsteps
		footsteps := "N/A"
		if len(member.Footsteps) > 0 {
			fsEntries := make([]string, 0, len(member.Footsteps))
			for _, fs := range member.Footsteps {
				activity := sanitizeCSVField(fs.Activity)
				description := sanitizeCSVField(fs.Description)
				// Combine Activity and Description
				fsEntry := fmt.Sprintf("Activity: %s; Description: %s", activity, description)
				fsEntries = append(fsEntries, fsEntry)
			}
			footsteps = strings.Join(fsEntries, "; ")
		}

		// Assemble the record
		record := []string{
			id,
			accountType,
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
			footsteps,
			createdAt,
			updatedAt,
		}
		records = append(records, record)
	}

	headers := []string{
		"ID",
		"Account Type",
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
		"Footsteps",
		"Created At",
		"Updated At",
	}

	return records, headers
}

func (m *ModelResource) ValidateMemberRequest(req *MemberRequest) error {
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

func (r *ModelResource) MemberGetByEmail(email string) (*Member, error) {
	var member Member
	err := r.MemberDB.DB.Client.Preload("Media").Where("email = ?", email).First(&member).Error
	return &member, err
}

func (r *ModelResource) MemberGetByUsername(username string) (*Member, error) {
	var member Member
	err := r.MemberDB.DB.Client.Preload("Media").Where("username = ?", username).First(&member).Error
	return &member, err
}

func (r *ModelResource) MemberGetByContactNumber(contactNumber string) (*Member, error) {
	var member Member
	err := r.MemberDB.DB.Client.Preload("Media").Where("contact_number = ?", contactNumber).First(&member).Error
	return &member, err
}

func (r *ModelResource) MemberFindByEmailUsernameOrContact(input string) (*Member, error) {
	switch r.helpers.GetKeyType(input) {
	case "contact":
		return r.MemberGetByContactNumber(input)
	case "email":
		return r.MemberGetByEmail(input)
	default:
		return r.MemberGetByUsername(input)
	}
}
func (r *ModelResource) MemberUpdatePassword(id uint, password string) error {
	newPassword, err := r.cryptoHelpers.HashPassword(password)
	if err != nil {
		return err
	}
	updated := &Member{Password: newPassword}
	_, err = r.MemberDB.UpdateColumns(id, *updated, []string{"Media"})
	return err
}

func (r *ModelResource) MemberCreate(user *Member) error {
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
	return r.MemberDB.Create(user)
}

func (r *ModelResource) MemberUpdate(user *Member, preloads []string) error {
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
	return r.MemberDB.Update(user, preloads)
}

func (m *ModelResource) MemberSeeders() error {
	m.logger.Info("Seeding Member")
	return nil
}
