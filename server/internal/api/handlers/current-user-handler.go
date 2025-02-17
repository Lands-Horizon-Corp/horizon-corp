package handlers

import (
	"strconv"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
	"github.com/rotisserie/eris"
)

type UserInfo struct {
	IPAddress      string  `json:"ipAddress"`
	UserAgent      string  `json:"userAgent"`
	Referer        string  `json:"referer"`
	Location       string  `json:"location"`
	AcceptLanguage string  `json:"acceptLanguage"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitudde"`
}

type CurrentUser struct {
	tokenProvider *providers.TokenService
	transformer   *models.ModelTransformer
	repository    *models.ModelRepository
	helpers       *helpers.HelpersFunction
}

func NewCurrentUser(
	tokenProvider *providers.TokenService,
	transformer *models.ModelTransformer,
	repository *models.ModelRepository,
	helpers *helpers.HelpersFunction,
) *CurrentUser {
	return &CurrentUser{
		tokenProvider: tokenProvider,
		transformer:   transformer,
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
	latStr := ctx.GetHeader("X-Latitude")
	lonStr := ctx.GetHeader("X-Longitude")

	var latitude, longitude *float64
	if latStr != "" {
		if lat, err := strconv.ParseFloat(latStr, 64); err == nil {
			latitude = &lat
		}
	}
	if lonStr != "" {
		if lon, err := strconv.ParseFloat(lonStr, 64); err == nil {
			longitude = &lon
		}
	}

	var lat, lon float64
	if latitude != nil {
		lat = *latitude
	}
	if longitude != nil {
		lon = *longitude
	}
	userInfo := &UserInfo{
		IPAddress:      ctx.ClientIP(),
		UserAgent:      ctx.GetHeader("User-Agent"),
		Referer:        ctx.GetHeader("Referer"),
		Location:       ctx.GetHeader("X-Location"),
		AcceptLanguage: ctx.GetHeader("Accept-Language"),
		Latitude:       lat,
		Longitude:      lon,
	}

	return userClaims, userInfo, nil
}

func (c *CurrentUser) GenericUser(ctx *gin.Context) (interface{}, error) {
	// Get claims
	claims, _, err := c.Claims(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "unauthorized")
	}

	// Switch on AccountType to figure out what user type to fetch
	preloads := c.helpers.GetPreload(ctx)
	switch claims.AccountType {
	case "Admin":
		admin, err := c.repository.AdminGetByID(claims.ID, preloads...)
		if err != nil {
			c.tokenProvider.ClearTokenCookie(ctx)
			return nil, eris.Wrap(err, "admin not found")
		}
		return c.transformer.AdminToResource(admin), nil

	case "Employee":
		employee, err := c.repository.EmployeeGetByID(claims.ID, preloads...)
		if err != nil {
			c.tokenProvider.ClearTokenCookie(ctx)
			return nil, eris.Wrap(err, "employee not found")
		}
		return c.transformer.EmployeeToResource(employee), nil

	case "Owner":
		owner, err := c.repository.OwnerGetByID(claims.ID, preloads...)
		if err != nil {
			c.tokenProvider.ClearTokenCookie(ctx)
			return nil, eris.Wrap(err, "owner not found")
		}
		return c.transformer.OwnerToResource(owner), nil

	case "Member":
		member, err := c.repository.MemberGetByID(claims.ID, preloads...)
		if err != nil {
			c.tokenProvider.ClearTokenCookie(ctx)
			return nil, eris.Wrap(err, "member not found")
		}
		return c.transformer.MemberToResource(member), nil

	default:
		return nil, eris.New("unknown account type")
	}
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
		c.tokenProvider.ClearTokenCookie(ctx)
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
		c.tokenProvider.ClearTokenCookie(ctx)
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
		c.tokenProvider.ClearTokenCookie(ctx)
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
		c.tokenProvider.ClearTokenCookie(ctx)
		return nil, eris.Wrap(err, "member not found")
	}

	return member, nil
}
