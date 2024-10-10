package resources

import (
	"time"
)

type SignUpResource struct {
	ID         uint   `json:"id"`
	Mode       string `json:"mode"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	LoggedInAt string `json:"isContactVerified"`
}

func ToResourceSignUp(id uint, mode, token string) SignUpResource {
	return SignUpResource{
		ID:         id,
		Email:      mode,
		Token:      token,
		LoggedInAt: time.Now().Format(time.RFC3339),
	}
}
