package resources

import (
	"time"
)

type SignInResource struct {
	Email      string `json:"email"`
	Token      string `json:"token"`
	LoggedInAt string `json:"loggedInAt"`
}

func ToResourceSignIn(email, token string) SignInResource {
	return SignInResource{
		Email:      email,
		Token:      token,
		LoggedInAt: time.Now().Format(time.RFC3339),
	}
}
