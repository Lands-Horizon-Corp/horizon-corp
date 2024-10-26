package resources

import (
	"horizon/server/internal/models"
	"time"
)

type OwnerResource struct {
	AccountType       string             `json:"account_type"`
	ID                uint               `json:"id"`
	FirstName         string             `json:"first_name"`
	LastName          string             `json:"last_name"`
	MiddleName        string             `json:"middle_name"`
	PermanentAddress  string             `json:"permanent_address"`
	Description       string             `json:"description"`
	Birthdate         time.Time          `json:"birthdate"`
	Username          string             `json:"username"`
	Email             string             `json:"email"`
	ContactNumber     string             `json:"contact_number"`
	IsEmailVerified   bool               `json:"is_email_verified"`
	IsContactVerified bool               `json:"is_contact_verified"`
	Status            models.OwnerStatus `json:"status"`
	Media             *MediaResource     `json:"media,omitempty"`
	Companies         []CompanyResource  `json:"companies,omitempty"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToResourceOwner(owner models.Owner) OwnerResource {
	var mediaResource *MediaResource
	if owner.Media != nil {
		mediaRes := ToResourceMedia(*owner.Media)
		mediaResource = &mediaRes
	}
	return OwnerResource{
		AccountType:       "Owner",
		ID:                owner.ID,
		FirstName:         owner.FirstName,
		LastName:          owner.LastName,
		MiddleName:        owner.MiddleName,
		PermanentAddress:  owner.PermanentAddress,
		Description:       owner.Description,
		Birthdate:         owner.Birthdate,
		Username:          owner.Username,
		Email:             owner.Email,
		ContactNumber:     owner.ContactNumber,
		IsEmailVerified:   owner.IsEmailVerified,
		IsContactVerified: owner.IsContactVerified,
		Status:            owner.Status,
		Media:             mediaResource,
		Companies:         ToResourceListCompanies(owner.Companies),
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
