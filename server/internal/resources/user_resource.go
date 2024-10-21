package resources

import (
	"horizon/server/internal/models"
	"time"
)

type UserResource struct {
	AccountType       string         `json:"accountType"`
	ID                uint           `json:"id"`
	FirstName         string         `json:"firstName"`
	LastName          string         `json:"lastName"`
	MiddleName        string         `json:"middleName"`
	PermanentAddress  string         `json:"permanentAddress"`
	Description       string         `json:"description"`
	Birthdate         time.Time      `json:"birthdate"`
	Username          string         `json:"username"`
	Email             string         `json:"email"`
	IsEmailVerified   bool           `json:"isEmailVerified"`
	IsContactVerified bool           `json:"isContactVerified"`
	ContactNumber     string         `json:"contactNumber"`
	Media             *MediaResource `json:"media,omitempty"`
	Status            string         `json:"status"`
	CreatedAt         string         `json:"createdAt"`
	UpdatedAt         string         `json:"updatedAt"`
}

func ToResourceUser(user models.User, accountType string) UserResource {
	var mediaResource *MediaResource

	if user.MediaID != nil {
		media := ToResourceMedia(user.Media)
		mediaResource = &media
	}

	return UserResource{
		AccountType:       accountType,
		ID:                user.ID,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		MiddleName:        user.MiddleName,
		PermanentAddress:  user.PermanentAddress,
		Description:       user.Description,
		Birthdate:         user.Birthdate,
		Username:          user.Username,
		Email:             user.Email,
		IsEmailVerified:   user.IsEmailVerified,
		IsContactVerified: user.IsContactVerified,
		ContactNumber:     user.ContactNumber,
		Media:             mediaResource,
		Status:            user.Status,
		CreatedAt:         user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         user.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResourceListUsers(users []models.User, accountType string) []UserResource {
	var resources []UserResource
	for _, user := range users {
		resources = append(resources, ToResourceUser(user, accountType))
	}
	return resources
}
