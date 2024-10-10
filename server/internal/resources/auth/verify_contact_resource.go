package resources

import (
	"time"
)

type VerifyContactResource struct {
	Token         string `json:"token"`
	ContactNumber string `json:"contactNumber"`
	VerifiedAt    string `json:"verifiedAt"`
	Message       string `json:"message"`
}

func ToResourceVerifyContact(token, contactNumber, message string) VerifyContactResource {
	return VerifyContactResource{
		Token:         token,
		ContactNumber: contactNumber,
		VerifiedAt:    time.Now().Format(time.RFC3339),
		Message:       message,
	}
}
