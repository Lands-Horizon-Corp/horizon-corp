package handlers

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
	"github.com/rotisserie/eris"
)

type UserInfo struct {
	IPAddress      string `json:"ipAddress"`
	UserAgent      string `json:"userAgent"`
	Referer        string `json:"referer"`
	Location       string `json:"location"`
	AcceptLanguage string `json:"acceptLanguage"`
}

type CurrentUser struct {
	tokenProvider *providers.TokenService
	repository    *models.ModelRepository
	helpers       *helpers.HelpersFunction
}

func NewCurrentUser(
	tokenProvider *providers.TokenService,
	repository *models.ModelRepository,
	helpers *helpers.HelpersFunction,
) *CurrentUser {
	return &CurrentUser{
		tokenProvider: tokenProvider,
		repository:    repository,
		helpers:       helpers,
	}
}

// Fetch Claims and User Info
func (c *CurrentUser) Claims(ctx *gin.Context) (*providers.UserClaims, *UserInfo, error) {
	claims, exists := ctx.Get("claims")
	if !exists {
		c.tokenProvider.ClearTokenCookie(ctx)
		return nil, nil, eris.New("claims not found in context")
	}
	userClaims, ok := claims.(*providers.UserClaims)
	if !ok {
		c.tokenProvider.ClearTokenCookie(ctx)
		return nil, nil, eris.New("failed to cast claims to *auth.UserClaims")
	}

	userInfo := &UserInfo{
		IPAddress:      ctx.ClientIP(),
		UserAgent:      ctx.GetHeader("User-Agent"),
		Referer:        ctx.GetHeader("Referer"),
		Location:       ctx.GetHeader("X-Location"),
		AcceptLanguage: ctx.GetHeader("Accept-Language"),
	}

	return userClaims, userInfo, nil
}

// Admin Handler
func (c *CurrentUser) Admin(ctx *gin.Context) (*models.Admin, error) {
	claims, _, err := c.Claims(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "unauthorized")
	}
	if claims.AccountType != "Admin" {
		return nil, eris.New("account type mismatch: not an admin")
	}
	admin, err := c.repository.AdminGetByID(claims.ID)
	if err != nil {
		return nil, eris.Wrap(err, "admin not found")
	}
	return admin, nil
}

// Employee Handler
func (c *CurrentUser) Employee(ctx *gin.Context) (*models.Employee, error) {
	claims, _, err := c.Claims(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "unauthorized")
	}

	if claims.AccountType != "Employee" {
		return nil, eris.New("account type mismatch: not an employee")
	}

	employee, err := c.repository.EmployeeGetByID(claims.ID)
	if err != nil {
		return nil, eris.Wrap(err, "employee not found")
	}
	return employee, nil
}

func (c *CurrentUser) Owner(ctx *gin.Context) (*models.Owner, error) {
	claims, _, err := c.Claims(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "unauthorized")
	}
	if claims.AccountType != "Owner" {
		return nil, eris.New("account type mismatch: not an owner")
	}
	owner, err := c.repository.OwnerGetByID(claims.ID)
	if err != nil {
		return nil, eris.Wrap(err, "owner not found")
	}

	return owner, nil
}

// Member Handler
func (c *CurrentUser) Member(ctx *gin.Context) (*models.Member, error) {
	claims, _, err := c.Claims(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "unauthorized")
	}
	if claims.AccountType != "Member" {
		return nil, eris.New("account type mismatch: not a member")
	}
	member, err := c.repository.MemberGetByID(claims.ID)
	if err != nil {
		return nil, eris.Wrap(err, "member not found")
	}

	return member, nil
}
