package transformers

import (
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
)

type MemberResource struct {
	AccountType string `json:"accountType"`

	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	FirstName          string            `json:"firstName"`
	LastName           string            `json:"lastName"`
	MiddleName         string            `json:"middleName"`
	PermanentAddress   string            `json:"permanentAddress"`
	Description        string            `json:"description"`
	BirthDate          time.Time         `json:"birthDate"`
	Username           string            `json:"username"`
	Email              string            `json:"email"`
	IsEmailVerified    bool              `json:"isEmailVerified"`
	IsContactVerified  bool              `json:"isContactVerified"`
	IsSkipVerification bool              `json:"isSkipVerification"`
	ContactNumber      string            `json:"contactNumber"`
	Status             models.UserStatus `json:"status"`
	MediaID            *uint             `json:"mediaID"`
	Media              *MediaResource    `json:"media"`
	BranchID           *uint             `json:"branchID"`
	Branch             *BranchResource   `json:"branch"`
	Longitude          *float64          `json:"longitude"`
	Latitude           *float64          `json:"latitude"`
	RoleID             *uint             `json:"roleID"`
	Role               *RoleResource     `json:"role"`
	GenderID           *uint             `json:"genderID"`
	Gender             *GenderResource   `json:"gender"`

	Footsteps []*FootstepResource `json:"footsteps"`
}
