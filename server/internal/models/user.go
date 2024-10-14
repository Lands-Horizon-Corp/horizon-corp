package models

type User struct {
	ID            uint   `json:"id"`
	Email         string `json:"email"`
	ContactNumber string `json:"contactNumber"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
}
