package resources

import (
	"time"
)

type SignInResource struct {
	ID         uint   `json:"id"`
	Mode       string `json:"mode"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	LoggedInAt string `json:"loggedInAt"`
}

func ToResourceSignIn(id uint, mode, token string) SignInResource {
	return SignInResource{
		ID:         id,
		Email:      mode,
		Token:      token,
		LoggedInAt: time.Now().Format(time.RFC3339),
	}
}
