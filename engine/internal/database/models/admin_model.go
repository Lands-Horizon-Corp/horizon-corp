package models

import (
	"errors"
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model

	// Fields
	FirstName          string     `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName           string     `gorm:"type:varchar(255);not null" json:"last_name"`
	MiddleName         string     `gorm:"type:varchar(255)" json:"middle_name"`
	PermanentAddress   string     `gorm:"type:text" json:"permanent_address"`
	Description        string     `gorm:"type:text" json:"description"`
	BirthDate          time.Time  `gorm:"type:date" json:"birth_date"`
	Username           string     `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email              string     `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password           string     `gorm:"type:varchar(255);not null" json:"password"`
	ContactNumber      string     `gorm:"type:varchar(15);unique;not null" json:"contact_number"`
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
	m.logger.Info("Seeding Admin")
	// Create user
	return nil
}
