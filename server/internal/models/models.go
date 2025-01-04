package models

import "github.com/golang-jwt/jwt"

type UserStatus string

const (
	PendingStatus    UserStatus = "Pending"
	VerifiedStatus   UserStatus = "Verified"
	NotAllowedStatus UserStatus = "Not allowed"
)

type UserAccountType string

const (
	AdminType    UserAccountType = "Admin"
	EmployeeType UserAccountType = "Employee"
	MemberType   UserAccountType = "Member"
	OwnerType    UserAccountType = "Owner"
)

type UserClaims struct {
	ID          uint            `json:"id"`
	AccountType UserAccountType `json:"accountType"`
	jwt.StandardClaims
}
