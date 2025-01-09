package transformers

import "time"

type MemberResource struct {
	ID                 uint               `json:"id"`
	CreatedAt          string             `json:"createdAt"`
	UpdatedAt          string             `json:"updatedAt"`
	FirstName          string             `json:"firstName"`
	LastName           string             `json:"lastName"`
	MiddleName         string             `json:"middleName,omitempty"`
	PermanentAddress   string             `json:"permanentAddress,omitempty"`
	Description        string             `json:"description,omitempty"`
	BirthDate          time.Time          `json:"birthDate"`
	Username           string             `json:"username"`
	Email              string             `json:"email"`
	IsEmailVerified    bool               `json:"isEmailVerified"`
	IsContactVerified  bool               `json:"isContactVerified"`
	IsSkipVerification bool               `json:"isSkipVerification"`
	ContactNumber      string             `json:"contactNumber"`
	Status             string             `json:"status"`
	Media              *MediaResource     `json:"media,omitempty"`
	Branch             *BranchResource    `json:"branch,omitempty"`
	Longitude          *float64           `json:"longitude,omitempty"`
	Latitude           *float64           `json:"latitude,omitempty"`
	Role               *RoleResource      `json:"role,omitempty"`
	Gender             *GenderResource    `json:"gender,omitempty"`
	Footsteps          []FootstepResource `json:"footsteps,omitempty"`
}
