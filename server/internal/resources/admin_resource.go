package resources

import (
	"horizon/server/internal/models"
	"time"
)

type AdminResource struct {
	ID                uint          `json:"id"`
	FirstName         string        `json:"firstName"`
	LastName          string        `json:"lastName"`
	PermanentAddress  string        `json:"permanentAddress"`
	Description       string        `json:"description"`
	Birthdate         string        `json:"birthdate"`
	Username          string        `json:"username"`
	Email             string        `json:"email"`
	IsEmailVerified   bool          `json:"isEmailVerified"`
	IsContactVerified bool          `json:"isContactVerified"`
	Media             MediaResource `json:"media"`
}

func ToResourceAdmin(admin models.Admin) AdminResource {
	return AdminResource{
		ID:                admin.ID,
		FirstName:         admin.FirstName,
		LastName:          admin.LastName,
		PermanentAddress:  admin.PermanentAddress,
		Description:       admin.Description,
		Birthdate:         admin.Birthdate.Format(time.RFC3339),
		Username:          admin.Username,
		Email:             admin.Email,
		IsEmailVerified:   admin.IsEmailVerified,
		IsContactVerified: admin.IsContactVerified,
		Media:             ToResourceMedia(admin.Media),
	}
}

func ToResourceListAdmin(adminList []models.Admin) []AdminResource {
	var resources []AdminResource
	for _, admin := range adminList {
		resources = append(resources, ToResourceAdmin(admin))
	}
	return resources
}
