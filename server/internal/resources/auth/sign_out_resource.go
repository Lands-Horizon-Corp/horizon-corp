package resources

import (
	"time"
)

type SignOutResource struct {
	Message     string `json:"message"`
	LoggedOutAt string `json:"loggedOutAt"`
}

func ToResourceSignOut(message string) SignOutResource {
	return SignOutResource{
		Message:     message,
		LoggedOutAt: time.Now().Format(time.RFC3339),
	}
}
