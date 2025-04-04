package models

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model

	// Fields
	FirstName          string     `gorm:"type:varchar(255);unsigned" json:"first_name"`
	LastName           string     `gorm:"type:varchar(255);unsigned" json:"last_name"`
	MiddleName         string     `gorm:"type:varchar(255)" json:"middle_name"`
	PermanentAddress   string     `gorm:"type:text" json:"permanent_address"`
	Description        string     `gorm:"type:text" json:"description"`
	BirthDate          time.Time  `gorm:"type:date" json:"birth_date"`
	Username           string     `gorm:"type:varchar(255);unique;unsigned" json:"username"`
	Email              string     `gorm:"type:varchar(255);unique;unsigned" json:"email"`
	Password           string     `gorm:"type:varchar(255);unsigned" json:"password"`
	ContactNumber      string     `gorm:"type:varchar(15);unique;unsigned" json:"contact_number"`
	IsEmailVerified    bool       `gorm:"default:false" json:"is_email_verified"`
	IsContactVerified  bool       `gorm:"default:false" json:"is_contact_verified"`
	IsSkipVerification bool       `gorm:"default:false" json:"is_skip_verification"`
	Status             UserStatus `gorm:"type:varchar(11);default:'Pending'" json:"status"`

	// Relationship 0 to 1
	MediaID *uint  `gorm:"type:bigint;unsigned" json:"media_id"`
	Media   *Media `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"media"`

	// Relationship 0 to 1
	RoleID *uint `gorm:"type:bigint;unsigned" json:"role_id"`
	Role   *Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`

	// Relationship 0 to 1
	GenderID *uint   `gorm:"type:bigint;unsigned" json:"gender_id"`
	Gender   *Gender `gorm:"foreignKey:GenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"gender"`

	// Relationship 0 to many
	Footsteps []*Footstep `gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"footsteps,omitempty"`
}

type AdminResource struct {
	AccountType string `json:"accountType"`
	ID          uint   `json:"id"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`

	FirstName          string              `json:"firstName"`
	LastName           string              `json:"lastName"`
	MiddleName         string              `json:"middleName"`
	PermanentAddress   string              `json:"permanentAddress"`
	Description        string              `json:"description"`
	BirthDate          time.Time           `json:"birthDate"`
	Username           string              `json:"username"`
	Email              string              `json:"email"`
	ContactNumber      string              `json:"contactNumber"`
	IsEmailVerified    bool                `json:"isEmailVerified"`
	IsContactVerified  bool                `json:"isContactVerified"`
	IsSkipVerification bool                `json:"isSkipVerification"`
	Status             UserStatus          `json:"status"`
	MediaID            *uint               `json:"mediaID"`
	Media              *MediaResource      `json:"media"`
	RoleID             *uint               `json:"roleID"`
	Role               *RoleResource       `json:"role"`
	GenderID           *uint               `json:"genderID"`
	Gender             *GenderResource     `json:"gender"`
	Footsteps          []*FootstepResource `json:"footsteps"`
}

type AdminRequest struct {
	FirstName          string    `json:"firstName" validate:"required,max=255"`
	LastName           string    `json:"lastName" validate:"required,max=255"`
	MiddleName         string    `json:"middleName,omitempty" validate:"max=255"`
	PermanentAddress   string    `json:"permanentAddress,omitempty"`
	Description        string    `json:"description,omitempty"`
	BirthDate          time.Time `json:"birthDate" validate:"required"`
	Username           string    `json:"username" validate:"required,max=255"`
	Email              string    `json:"email" validate:"required,email,max=255"`
	Password           string    `json:"password" validate:"required,min=8,max=255"`
	ContactNumber      string    `json:"contactNumber" validate:"required,max=15"`
	IsEmailVerified    bool      `json:"isEmailVerified"`
	IsContactVerified  bool      `json:"isContactVerified"`
	IsSkipVerification bool      `json:"isSkipVerification"`
	Status             string    `json:"status" validate:"required,oneof=Pending Active Inactive"`

	MediaID  *uint `json:"mediaID,omitempty"`
	RoleID   *uint `json:"roleID,omitempty"`
	GenderID *uint `json:"genderID,omitempty"`
}

func (m *ModelResource) AdminToResource(admin *Admin) *AdminResource {
	if admin == nil {
		return nil
	}
	return &AdminResource{
		AccountType: "Admin",
		ID:          admin.ID,
		CreatedAt:   admin.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   admin.UpdatedAt.Format(time.RFC3339),

		FirstName:          admin.FirstName,
		LastName:           admin.LastName,
		MiddleName:         admin.MiddleName,
		PermanentAddress:   admin.PermanentAddress,
		Description:        admin.Description,
		BirthDate:          admin.BirthDate,
		Username:           admin.Username,
		Email:              admin.Email,
		ContactNumber:      admin.ContactNumber,
		IsEmailVerified:    admin.IsEmailVerified,
		IsContactVerified:  admin.IsContactVerified,
		IsSkipVerification: admin.IsSkipVerification,
		Status:             admin.Status,
		MediaID:            admin.MediaID,
		Media:              m.MediaToResource(admin.Media),
		RoleID:             admin.RoleID,
		Role:               m.RoleToResource(admin.Role),
		GenderID:           admin.GenderID,
		Gender:             m.GenderToResource(admin.Gender),
		Footsteps:          m.FootstepToResourceList(admin.Footsteps),
	}
}

func (m *ModelResource) AdminToResourceList(admins []*Admin) []*AdminResource {
	if admins == nil {
		return nil
	}
	var adminResources []*AdminResource
	for _, admin := range admins {
		adminResources = append(adminResources, m.AdminToResource(admin))
	}
	return adminResources
}

func (m *ModelResource) AdminToRecord(admins []*Admin) ([][]string, []string) {
	// Convert Admin structs to AdminResource structs
	resource := m.AdminToResourceList(admins)
	records := make([][]string, 0, len(resource))

	for _, admin := range resource {
		id := strconv.Itoa(int(admin.ID))
		firstName := sanitizeCSVField(admin.FirstName)
		lastName := sanitizeCSVField(admin.LastName)
		middleName := sanitizeCSVField(admin.MiddleName)
		permanentAddress := sanitizeCSVField(admin.PermanentAddress)
		description := sanitizeCSVField(admin.Description)
		birthDate := admin.BirthDate.Format("2006-01-02") // Format as YYYY-MM-DD
		username := sanitizeCSVField(admin.Username)
		email := sanitizeCSVField(admin.Email)
		contactNumber := sanitizeCSVField(admin.ContactNumber)
		isEmailVerified := strconv.FormatBool(admin.IsEmailVerified)
		isContactVerified := strconv.FormatBool(admin.IsContactVerified)
		isSkipVerification := strconv.FormatBool(admin.IsSkipVerification)
		status := sanitizeCSVField(string(admin.Status))
		createdAt := sanitizeCSVField(admin.CreatedAt)
		updatedAt := sanitizeCSVField(admin.UpdatedAt)

		// Handle Media
		mediaURL := "N/A"
		if admin.Media != nil {
			mediaURL = sanitizeCSVField(admin.Media.URL)
		}

		// Handle Role
		roleName := "N/A"
		if admin.Role != nil {
			roleName = sanitizeCSVField(admin.Role.Name) // Assuming Role has a Name field
		}

		// Handle Gender
		genderName := "N/A"
		if admin.Gender != nil {
			genderName = sanitizeCSVField(admin.Gender.Name) // Assuming Gender has a Name field
		}

		// Handle Footsteps
		footsteps := "N/A"
		if len(admin.Footsteps) > 0 {
			footstepDescriptions := make([]string, 0, len(admin.Footsteps))
			for _, footstep := range admin.Footsteps {
				footstepDescriptions = append(footstepDescriptions, sanitizeCSVField(footstep.Description))
			}
			footsteps = strings.Join(footstepDescriptions, "; ")
		}

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
			roleName,
			genderName,
			mediaURL,
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
		"Role",
		"Gender",
		"Media URL",
		"Footsteps",
		"Created At",
		"Updated At",
	}

	return records, headers
}

func (m *ModelResource) ValidateAdminRequest(req *AdminRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return m.helpers.FormatValidationError(err)
	}
	return nil
}

func (r *ModelResource) AdminGetByEmail(email string) (*Admin, error) {
	var admin Admin
	err := r.AdminDB.DB.Client.Preload("Media").Where("email = ?", email).First(&admin).Error
	return &admin, err
}

func (r *ModelResource) AdminGetByUsername(username string) (*Admin, error) {
	var admin Admin
	err := r.AdminDB.DB.Client.Preload("Media").Where("username = ?", username).First(&admin).Error
	return &admin, err
}

func (r *ModelResource) AdminGetByContactNumber(contactNumber string) (*Admin, error) {
	var admin Admin
	err := r.AdminDB.DB.Client.Preload("Media").Where("contact_number = ?", contactNumber).First(&admin).Error
	return &admin, err
}

func (r *ModelResource) AdminFindByEmailUsernameOrContact(input string) (*Admin, error) {
	switch r.helpers.GetKeyType(input) {
	case "contact":
		return r.AdminGetByContactNumber(input)
	case "email":
		return r.AdminGetByEmail(input)
	default:
		return r.AdminGetByUsername(input)
	}
}
func (r *ModelResource) AdminUpdatePassword(id uint, password string) error {
	newPassword, err := r.cryptoHelpers.HashPassword(password)
	if err != nil {
		return err
	}
	updated := &Admin{Password: newPassword}
	_, err = r.AdminDB.UpdateColumns(id, *updated, []string{"Media"})
	return err
}
func (r *ModelResource) AdminCreate(user *Admin) error {
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
	return r.AdminDB.Create(user)
}

func (r *ModelResource) AdminUpdate(user *Admin, preloads []string) error {
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
	return r.AdminDB.Update(user, preloads)
}

func (m *ModelResource) AdminSeeders() error {
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

	admins := []Admin{
		{
			FirstName:         "Lands",
			LastName:          "Horizon",
			MiddleName:        "ecoop",
			PermanentAddress:  "123 Main St, Springfield",
			Description:       "Super administrator",
			BirthDate:         time.Date(1985, 7, 12, 0, 0, 0, 0, time.UTC),
			Username:          "alicej",
			Email:             m.cfg.AdminEmail,
			Password:          m.cfg.AdminPassword,
			ContactNumber:     "5551234567",
			IsEmailVerified:   true,
			IsContactVerified: true,
			Status:            "Verified",
			MediaID:           &uploaded.ID,
		},
	}

	for _, admin := range admins {
		err := m.AdminCreate(&admin)
		if err != nil {
			log.Printf("Error seeding admin %s: %v", admin.Email, err)
		} else {
			log.Printf("Admin %s seeded successfully", admin.Email)
		}
	}

	m.logger.Info("Admins seeded successfully")
	return nil
}
