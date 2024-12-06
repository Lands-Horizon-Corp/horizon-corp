package models

type UserStatus string

const (
	AdminPending    UserStatus = "Pending"
	AdminVerified   UserStatus = "Verified"
	AdminNotAllowed UserStatus = "Not allowed"
)
