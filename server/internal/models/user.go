package models

type User struct {
	AccountType       string `json:"accountType"`
	ID                uint   `json:"id"`
	Email             string `json:"email"`
	ContactNumber     string `json:"contactNumber"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	IsEmailVerified   bool   `json:"isEmailVerified"`
	IsContactVerified bool   `json:"isContactVerified"`
}
