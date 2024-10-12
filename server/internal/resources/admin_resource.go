package resources

import (
	"horizon/server/internal/models"
	"time"
)

type AdminResource struct {
	AccountType       string             `json:"accountType"`
	ID                uint               `json:"id"`
	FirstName         string             `json:"firstName"`
	LastName          string             `json:"lastName"`
	MiddleName        string             `json:"middleName"`
	PermanentAddress  string             `json:"permanentAddress"`
	Description       string             `json:"description"`
	Birthdate         time.Time          `json:"birthdate"`
	Username          string             `json:"username"`
	Email             string             `json:"email"`
	IsEmailVerified   bool               `json:"isEmailVerified"`
	IsContactVerified bool               `json:"isContactVerified"`
	ContactNumber     string             `json:"contactNumber"`
	Media             *MediaResource     `json:"media,omitempty"`
	Status            models.AdminStatus `json:"status"`
	CreatedAt         string             `json:"createdAt"`
	UpdatedAt         string             `json:"updatedAt"`
}

func ToResourceAdmin(admin models.Admin) AdminResource {
	var mediaResource *MediaResource

	if admin.MediaID != nil {
		media := ToResourceMedia(admin.Media)
		mediaResource = &media
	}

	return AdminResource{
		AccountType:       "Admin",
		ID:                admin.ID,
		FirstName:         admin.FirstName,
		LastName:          admin.LastName,
		MiddleName:        admin.MiddleName,
		PermanentAddress:  admin.PermanentAddress,
		Description:       admin.Description,
		Birthdate:         admin.Birthdate,
		Username:          admin.Username,
		Email:             admin.Email,
		IsEmailVerified:   admin.IsEmailVerified,
		IsContactVerified: admin.IsContactVerified,
		ContactNumber:     admin.ContactNumber,
		Media:             mediaResource, // Set the media resource if exists
		Status:            admin.Status,
		CreatedAt:         admin.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         admin.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResourceListAdmins(admins []models.Admin) []AdminResource {
	var resources []AdminResource
	for _, admin := range admins {
		resources = append(resources, ToResourceAdmin(admin))
	}
	return resources
}
