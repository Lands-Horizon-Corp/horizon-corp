package resources

import "horizon-core/internal/models"

type BranchResource struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Email          string         `json:"email"`
	Address        string         `json:"address"`
	ContactNumber  string         `json:"contact_number"`
	Approved       bool           `json:"approved"`
	Description    string         `json:"description"`
	Theme          string         `json:"theme"`
	ProfilePicture *MediaResource `json:"profile_picture,omitempty"`
}

func NewBranchResource(branch models.Branch) BranchResource {
	profilePicture := NewMediaResource(branch.ProfilePicture)

	return BranchResource{
		ID:             branch.ID,
		Name:           branch.Name,
		Email:          branch.Email,
		Address:        branch.Address,
		ContactNumber:  branch.ContactNumber,
		Approved:       branch.Approved,
		Description:    branch.Description,
		Theme:          branch.Theme,
		ProfilePicture: profilePicture,
	}
}
