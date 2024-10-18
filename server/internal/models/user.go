package models

type User struct {
	AccountType       string `json:"accountType"`
	ID                uint   `json:"id"`
	Email             string `json:"email"`
	ContactNumber     string `json:"contactNumbeår"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	MiddleName        string `json:"middleName"`
	IsEmailVerified   bool   `json:"isEmailVerified"`
	IsContactVerified bool   `json:"isContactVerified"`
	Status            string `json:"status"`
	Media             *Media `json:"media"`
}
