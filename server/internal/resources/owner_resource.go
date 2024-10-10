package resources

import (
	"horizon/server/internal/models"
	"time"
)

type OwnerResource struct {
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

func ToResourceOwner(owner models.Owner) OwnerResource {
	return OwnerResource{
		ID:                owner.ID,
		FirstName:         owner.FirstName,
		LastName:          owner.LastName,
		PermanentAddress:  owner.PermanentAddress,
		Description:       owner.Description,
		Birthdate:         owner.Birthdate.Format(time.RFC3339),
		Username:          owner.Username,
		Email:             owner.Email,
		IsEmailVerified:   owner.IsEmailVerified,
		IsContactVerified: owner.IsContactVerified,
		Media:             ToResourceMedia(owner.Media),
	}
}

func ToResourceListOwner(ownerList []models.Owner) []OwnerResource {
	var resources []OwnerResource
	for _, owner := range ownerList {
		resources = append(resources, ToResourceOwner(owner))
	}
	return resources
}
