package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Owner struct {
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
	GenderID *uint `gorm:"type:bigint;unsigned" json:"gender_id"`

	// Relationship 0 to 1
	Gender *Gender `gorm:"foreignKey:GenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"gender"`

	// Relationship 0 to many
	Footsteps []*Footstep `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"footsteps,omitempty"`
	Companies []*Company  `gorm:"foreignKey:OwnerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"companies"`

	// Relationship 0 to many
	RoleID *uint `gorm:"type:bigint;unsigned" json:"role_id"`
	Role   *Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`
}

type OwnerResource struct {
	AccountType string `json:"accountType"`

	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	FirstName          string              `json:"firstName"`
	LastName           string              `json:"lastName"`
	MiddleName         string              `json:"middleName"`
	PermanentAddress   string              `json:"permanentAddress"`
	Description        string              `json:"description"`
	BirthDate          time.Time           `json:"birthDate"`
	Username           string              `json:"username"`
	Email              string              `json:"email"`
	Password           string              `json:"password"`
	ContactNumber      string              `json:"contactNumber"`
	IsEmailVerified    bool                `json:"isEmailVerified"`
	IsContactVerified  bool                `json:"isContactVerified"`
	IsSkipVerification bool                `json:"isSkipVerification"`
	Status             UserStatus          `json:"status"`
	MediaID            *uint               `json:"mediaID"`
	Media              *MediaResource      `json:"media"`
	GenderID           *uint               `json:"genderID"`
	Gender             *GenderResource     `json:"gender"`
	RoleID             *uint               `json:"roleID"`
	Role               *RoleResource       `json:"role"`
	Footsteps          []*FootstepResource `json:"footsteps"`
	Companies          []*CompanyResource  `json:"companies"`
}

func (m *ModelResource) OwnerToResource(owner *Owner) *OwnerResource {
	if owner == nil {
		return nil
	}

	return &OwnerResource{
		AccountType: "Owner",
		ID:          owner.ID,
		CreatedAt:   owner.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   owner.UpdatedAt.Format(time.RFC3339),

		FirstName:        owner.FirstName,
		LastName:         owner.LastName,
		MiddleName:       owner.MiddleName,
		PermanentAddress: owner.PermanentAddress,
		Description:      owner.Description,
		BirthDate:        owner.BirthDate,
		Username:         owner.Username,
		Email:            owner.Email,

		ContactNumber:      owner.ContactNumber,
		IsEmailVerified:    owner.IsEmailVerified,
		IsContactVerified:  owner.IsContactVerified,
		IsSkipVerification: owner.IsSkipVerification,
		Status:             owner.Status,
		MediaID:            owner.MediaID,
		Media:              m.MediaToResource(owner.Media),
		GenderID:           owner.GenderID,
		Gender:             m.GenderToResource(owner.Gender),
		RoleID:             owner.RoleID,
		Role:               m.RoleToResource(owner.Role),
		Footsteps:          m.FootstepToResourceList(owner.Footsteps),
		Companies:          m.CompanyToResourceList(owner.Companies),
	}
}

type OwnerRequest struct {
	FirstName        string    `json:"firstName" validate:"required,max=255"`
	LastName         string    `json:"lastName" validate:"required,max=255"`
	MiddleName       string    `json:"middleName,omitempty" validator:"max=255"`
	PermanentAddress string    `json:"permanentAddress,omitempty"`
	Description      string    `json:"description,omitempty"`
	BirthDate        time.Time `json:"birthDate" validate:"required"`
	Username         string    `json:"username" validate:"required,max=255"`
	Email            string    `json:"email" validate:"required,email,max=255"`

	ContactNumber      string `json:"contactNumber" validate:"required,max=15"`
	IsEmailVerified    bool   `json:"isEmailVerified"`
	IsContactVerified  bool   `json:"isContactVerified"`
	IsSkipVerification bool   `json:"isSkipVerification"`
	Status             string `json:"status" validate:"required,oneof=Pending Active Inactive"`
	MediaID            *uint  `json:"mediaID,omitempty"`
	GenderID           *uint  `json:"genderID,omitempty"`
	RoleID             *uint  `json:"roleID,omitempty"`
}

func (m *ModelResource) OwnerToResourceList(owners []*Owner) []*OwnerResource {
	if owners == nil {
		return nil
	}
	var ownerResources []*OwnerResource
	for _, owner := range owners {
		ownerResources = append(ownerResources, m.OwnerToResource(owner))
	}
	return ownerResources
}

func (m *ModelResource) ValidateOwnerRequest(req *OwnerRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return m.helpers.FormatValidationError(err)
	}
	return nil
}

func (r *ModelResource) OwnerGetByEmail(email string) (*Owner, error) {
	var owner Owner
	err := r.OwnerDB.DB.Client.Preload("Media").Where("email = ?", email).First(&owner).Error
	return &owner, err
}

func (r *ModelResource) OwnerGetByUsername(username string) (*Owner, error) {
	var owner Owner
	err := r.OwnerDB.DB.Client.Preload("Media").Where("username = ?", username).First(&owner).Error
	return &owner, err
}

func (r *ModelResource) OwnerGetByContactNumber(contactNumber string) (*Owner, error) {
	var owner Owner
	err := r.OwnerDB.DB.Client.Preload("Media").Where("contact_number = ?", contactNumber).First(&owner).Error
	return &owner, err
}

func (r *ModelResource) OwnerFindByEmailUsernameOrContact(input string) (*Owner, error) {
	switch r.helpers.GetKeyType(input) {
	case "contact":
		return r.OwnerGetByContactNumber(input)
	case "email":
		return r.OwnerGetByEmail(input)
	default:
		return r.OwnerGetByUsername(input)
	}
}
func (r *ModelResource) OwnerUpdatePassword(id uint, password string) error {
	newPassword, err := r.cryptoHelpers.HashPassword(password)
	if err != nil {
		return err
	}
	updated := &Owner{Password: newPassword}
	_, err = r.OwnerDB.UpdateColumns(id, *updated, []string{"Media"})
	return err
}

func (r *ModelResource) OwnerCreate(user *Owner) error {
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
	return r.OwnerDB.Create(user)
}

func (r *ModelResource) OwnerUpdate(user *Owner, preloads []string) error {
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
	return r.OwnerDB.Update(user, preloads)
}
func (m *ModelResource) OwnerSeeders() error {
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
	owners := []Owner{
		{
			FirstName:         "John",
			LastName:          "Doe",
			MiddleName:        "A.",
			PermanentAddress:  "123 Main St, Springfield",
			Description:       "A sample owner",
			BirthDate:         time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Username:          "johndoe",
			Email:             m.cfg.OwnerEmail,
			Password:          m.cfg.OwnerPassword,
			ContactNumber:     "1234567890",
			IsEmailVerified:   true,
			IsContactVerified: true,
			Status:            "Verified",
			MediaID:           &uploaded.ID,
		},
	}

	for _, owner := range owners {
		err := m.OwnerCreate(&owner)
		if err != nil {
			log.Printf("Error seeding owner %s: %v", owner.Email, err)
		} else {
			companies := []Company{
				{
					Name:            fmt.Sprintf("%s %s", "Horizon Tech Solutions", owner.FirstName),
					Description:     "A leading tech solutions provider.",
					Address:         "123 Innovation Drive, Silicon Valley",
					Longitude:       -122.4194,
					Latitude:        37.7749,
					ContactNumber:   "555-123-4567",
					IsAdminVerified: true,
					OwnerID:         &owner.ID,
					MediaID:         &uploaded.ID,
				},
				{
					Name:            fmt.Sprintf("%s %s", "Green Earth Enterprises", owner.FirstName),
					Description:     "Eco-friendly products and services.",
					Address:         "456 Green Lane, Eco City",
					Longitude:       -74.0060,
					Latitude:        40.7128,
					ContactNumber:   "555-987-6543",
					IsAdminVerified: false,
					OwnerID:         &owner.ID,
					MediaID:         &uploaded.ID,
				},
				{
					Name:            fmt.Sprintf("%s %s", "Sky High Airlines", owner.FirstName),
					Description:     "Premium airline services worldwide.",
					Address:         "789 Skyway Ave, Airport City",
					Longitude:       151.2093,
					Latitude:        -33.8688,
					ContactNumber:   "555-321-7890",
					IsAdminVerified: true,
					OwnerID:         &owner.ID,
					MediaID:         &uploaded.ID,
				},
			}

			for _, company := range companies {
				err := m.CompanyDB.Create(&company)
				if err != nil {
					log.Printf("Error seeding company %s: %v", company.Name, err)
				}
			}
			m.logger.Info("Companies seeded successfully")
		}
	}
	return nil
}
