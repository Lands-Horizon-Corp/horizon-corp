package resources

import (
	"horizon/server/internal/models"
	"time"
)

type OwnerResource struct {
	AccountType       string              `json:"accountType"`
	ID                uint                `json:"id"`
	FirstName         string              `json:"firstName"`
	LastName          string              `json:"lastName"`
	MiddleName        string              `json:"middleName"`
	PermanentAddress  string              `json:"permanentAddress"`
	Description       string              `json:"description"`
	BirthDate         time.Time           `json:"birthDate"`
	Username          string              `json:"username"`
	Email             string              `json:"email"`
	ContactNumber     string              `json:"contactNumber"`
	IsEmailVerified   bool                `json:"isEmailVerified"`
	IsContactVerified bool                `json:"isContactVerified"`
	Status            models.OwnerStatus  `json:"status"`
	Media             *MediaResource      `json:"media,omitempty"`
	Companies         []*CompanyResource  `json:"companies,omitempty"` // Updated to slice of pointers
	GenderID          *uint               `json:"genderId,omitempty"`
	Gender            *GenderResource     `json:"gender,omitempty"`
	Footsteps         []*FootstepResource `json:"footsteps,omitempty"` // Updated to slice of pointers

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Convert models.Owner to *OwnerResource
func ToResourceOwner(owner *models.Owner) *OwnerResource {
	if owner == nil {
		return nil
	}

	// Convert Media
	var mediaResource *MediaResource
	if owner.Media != nil {
		mediaResource = ToResourceMedia(owner.Media)
	}

	// Convert Gender
	var genderResource *GenderResource
	if owner.Gender != nil {
		genderResource = ToResourceGender(owner.Gender)
	}

	return &OwnerResource{
		AccountType:       "Owner",
		ID:                owner.ID,
		FirstName:         owner.FirstName,
		LastName:          owner.LastName,
		MiddleName:        owner.MiddleName,
		PermanentAddress:  owner.PermanentAddress,
		Description:       owner.Description,
		BirthDate:         owner.BirthDate,
		Username:          owner.Username,
		Email:             owner.Email,
		ContactNumber:     owner.ContactNumber,
		IsEmailVerified:   owner.IsEmailVerified,
		IsContactVerified: owner.IsContactVerified,
		Status:            owner.Status,
		Media:             mediaResource,
		GenderID:          owner.GenderID,
		Gender:            genderResource,
		Companies:         ToResourceListCompanies(owner.Companies),
		Footsteps:         ToResourceListFootsteps(owner.Footsteps),
		CreatedAt:         owner.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         owner.UpdatedAt.Format(time.RFC3339),
	}
}

// Convert []*models.Owner to []*OwnerResource
func ToResourceListOwners(owners []*models.Owner) []*OwnerResource {
	var resources []*OwnerResource
	for _, owner := range owners {
		resources = append(resources, ToResourceOwner(owner))
	}
	return resources
}
