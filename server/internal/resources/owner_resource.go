package resources

import (
	"horizon/server/internal/models"
	"time"
)

type OwnerResource struct {
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
	ContactNumber     string             `json:"contactNumber"`
	IsEmailVerified   bool               `json:"isEmailVerified"`
	IsContactVerified bool               `json:"isContactVerified"`
	Status            models.OwnerStatus `json:"status"`
	Media             *MediaResource     `json:"media,omitempty"`
	Companies         []CompanyResource  `json:"companies,omitempty"`
	GenderID          *uint              `json:"genderId,omitempty"`
	Gender            *GenderResource    `json:"gender,omitempty"`
	CreatedAt         string             `json:"createdAt"`
	UpdatedAt         string             `json:"updatedAt"`
}

func ToResourceOwner(owner models.Owner) OwnerResource {
	var mediaResource *MediaResource
	if owner.Media != nil {
		mediaRes := ToResourceMedia(*owner.Media)
		mediaResource = &mediaRes
	}

	var genderResource *GenderResource
	if owner.Gender != nil {
		genderRes := ToResourceGender(*owner.Gender)
		genderResource = &genderRes
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
		GenderID:          owner.GenderID,
		Gender:            genderResource,
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
