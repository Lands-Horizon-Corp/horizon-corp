package resources

import (
	"horizon/server/internal/models"
	"time"
)

type AdminResource struct {
	AccountType        string             `json:"accountType"`
	ID                 uint               `json:"id"`
	FirstName          string             `json:"firstName"`
	LastName           string             `json:"lastName"`
	MiddleName         string             `json:"middleName,omitempty"`
	PermanentAddress   string             `json:"permanentAddress,omitempty"`
	Description        string             `json:"description,omitempty"`
	Birthdate          time.Time          `json:"birthdate"`
	Username           string             `json:"username"`
	Email              string             `json:"email"`
	ContactNumber      string             `json:"contactNumber"`
	IsEmailVerified    bool               `json:"isEmailVerified"`
	IsContactVerified  bool               `json:"isContactVerified"`
	IsSkipVerification bool               `json:"isSkipVerification"`
	Status             models.AdminStatus `json:"status"`
	Media              *MediaResource     `json:"media,omitempty"`
	Role               *RoleResource      `json:"role,omitempty"`
	GenderID           *uint              `json:"genderId,omitempty"`
	Gender             *GenderResource    `json:"gender,omitempty"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func ToResourceAdmin(admin models.Admin) AdminResource {
	var mediaResource *MediaResource
	if admin.Media != nil {
		mediaRes := ToResourceMedia(*admin.Media)
		mediaResource = &mediaRes
	}

	var roleResource *RoleResource
	if admin.Role != nil {
		roleRes := ToResourceRole(*admin.Role)
		roleResource = &roleRes
	}
	var genderResource *GenderResource
	if admin.Gender != nil {
		genderRes := ToResourceGender(*admin.Gender)
		genderResource = &genderRes
	}

	return AdminResource{
		AccountType:        "Admin",
		ID:                 admin.ID,
		FirstName:          admin.FirstName,
		LastName:           admin.LastName,
		MiddleName:         admin.MiddleName,
		PermanentAddress:   admin.PermanentAddress,
		Description:        admin.Description,
		Birthdate:          admin.Birthdate,
		Username:           admin.Username,
		Email:              admin.Email,
		ContactNumber:      admin.ContactNumber,
		IsEmailVerified:    admin.IsEmailVerified,
		IsContactVerified:  admin.IsContactVerified,
		IsSkipVerification: admin.IsSkipVerification,
		Status:             admin.Status,
		Media:              mediaResource,
		Role:               roleResource,
		GenderID:           admin.GenderID,
		Gender:             genderResource,
		CreatedAt:          admin.CreatedAt.Format(time.RFC3339),
		UpdatedAt:          admin.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResourceListAdmins(admins []models.Admin) []AdminResource {
	var resources []AdminResource
	for _, admin := range admins {
		resources = append(resources, ToResourceAdmin(admin))
	}
	return resources
}
