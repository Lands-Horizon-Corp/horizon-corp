package resources

import (
	"time"
)

type VerifyContactResource struct {
	ContactNumber string `json:"contactNumber"`
	VerifiedAt    string `json:"verifiedAt"`
	Message       string `json:"message"`
}

func ToResourceVerifyContact(contactNumber, message string) VerifyContactResource {
	return VerifyContactResource{
		ContactNumber: contactNumber,
		VerifiedAt:    time.Now().Format(time.RFC3339),
		Message:       message,
	}
}
