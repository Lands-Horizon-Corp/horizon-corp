package resources

import (
	"time"
)

type VerifyEmailResource struct {
	Email      string `json:"email"`
	VerifiedAt string `json:"verifiedAt"`
	Message    string `json:"message"`
}

func ToResourceVerifyEmail(token, email, message string) VerifyEmailResource {
	return VerifyEmailResource{
		Email:      email,
		VerifiedAt: time.Now().Format(time.RFC3339),
	}
}
