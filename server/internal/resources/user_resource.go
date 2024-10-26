package resources

import (
	"horizon/server/internal/models"
	"time"
)

type UserResource struct {
	AccountType        string         `json:"accountType"`
	ID                 uint           `json:"id"`
	FirstName          string         `json:"firstName"`
	LastName           string         `json:"lastName"`
	MiddleName         string         `json:"middleName"`
	PermanentAddress   string         `json:"permanentAddress"`
	Description        string         `json:"description"`
	Birthdate          time.Time      `json:"birthdate"`
	Username           string         `json:"username"`
	Email              string         `json:"email"`
	IsEmailVerified    bool           `json:"isEmailVerified"`
	IsContactVerified  bool           `json:"isContactVerified"`
	IsSkipVerification bool           `json:"isSkipVerification"`
	ContactNumber      string         `json:"contactNumber"`
	Media              *MediaResource `json:"media,omitempty"`
	Status             string         `json:"status"`
	CreatedAt          string         `json:"createdAt"`
	UpdatedAt          string         `json:"updatedAt"`
}

// Convert models.User to *UserResource
func ToResourceUser(user *models.User, accountType string) *UserResource {
	if user == nil {
		return nil
	}

	var mediaResource *MediaResource
	if user.Media != nil {
		mediaResource = ToResourceMedia(user.Media)
	}

	return &UserResource{
		AccountType:        accountType,
		ID:                 user.ID,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		MiddleName:         user.MiddleName,
		PermanentAddress:   user.PermanentAddress,
		Description:        user.Description,
		Birthdate:          user.Birthdate,
		Username:           user.Username,
		Email:              user.Email,
		IsEmailVerified:    user.IsEmailVerified,
		IsContactVerified:  user.IsContactVerified,
		IsSkipVerification: user.IsSkipVerification,
		ContactNumber:      user.ContactNumber,
		Media:              mediaResource,
		Status:             user.Status,
		CreatedAt:          user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:          user.UpdatedAt.Format(time.RFC3339),
	}
}

// Convert []*models.User to []*UserResource
func ToResourceListUsers(users []*models.User, accountType string) []*UserResource {
	var resources []*UserResource
	for _, user := range users {
		resources = append(resources, ToResourceUser(user, accountType))
	}
	return resources
}
