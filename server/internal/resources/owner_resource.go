package resources

import (
	"horizon/server/internal/models"
	"time"
)

type OwnerResource struct {
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
	Status            models.OwnerStatus `json:"status"`
	CreatedAt         string             `json:"createdAt"`
	UpdatedAt         string             `json:"updatedAt"`
}

func ToResourceOwner(owner models.Owner) OwnerResource {
	var mediaResource *MediaResource

	if owner.MediaID != nil {
		media := ToResourceMedia(owner.Media)
		mediaResource = &media
	}

	return OwnerResource{
		ID:                owner.ID,
		FirstName:         owner.FirstName,
		LastName:          owner.LastName,
		MiddleName:        owner.MiddleName,
		PermanentAddress:  owner.PermanentAddress,
		Description:       owner.Description,
		Birthdate:         owner.Birthdate,
		Username:          owner.Username,
		Email:             owner.Email,
		IsEmailVerified:   owner.IsEmailVerified,
		IsContactVerified: owner.IsContactVerified,
		ContactNumber:     owner.ContactNumber,
		Media:             mediaResource, // Set the media resource if exists
		Status:            owner.Status,
		CreatedAt:         owner.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         owner.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResourceListOwners(owners []models.Owner) []OwnerResource {
	var resources []OwnerResource
	for _, owner := range owners {
		resources = append(resources, ToResourceOwner(owner))
	}
	return resources
}
