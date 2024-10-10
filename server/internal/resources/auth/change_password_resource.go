package resources

import (
	"time"
)

type ChangePasswordResource struct {
	Message   string `json:"message"`
	UpdatedAt string `json:"updatedAt"`
}

func ToResourceChangePassword(message string) ChangePasswordResource {
	return ChangePasswordResource{
		Message:   message,
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
}
