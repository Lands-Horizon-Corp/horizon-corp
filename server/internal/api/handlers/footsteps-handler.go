package handlers

import (
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rotisserie/eris"
)

// This will be a middleware

type FootstepHandler struct {
	currentUser *CurrentUser
}

func NewFootstepHandler(
	currentUser *CurrentUser,
) *FootstepHandler {
	return &FootstepHandler{
		currentUser: currentUser,
	}
}
func (f *FootstepHandler) Create(ctx *gin.Context, module, activity, description string, latitude, longitude *float64) (*models.Footstep, error) {
	claims, userInfo, err := f.currentUser.Claims(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "failed to retrieve current user claims")
	}

	var adminID, employeeID, ownerID, memberID *uuid.UUID
	var preloads []string

	switch claims.AccountType {
	case "Admin":
		admin, err := f.currentUser.Admin(ctx)
		if err != nil {
			return nil, eris.Wrap(err, "failed to retrieve admin user")
		}
		adminID = &admin.ID
		preloads = append(preloads, "Admin")
	case "Employee":
		employee, err := f.currentUser.Employee(ctx)
		if err != nil {
			return nil, eris.Wrap(err, "failed to retrieve employee user")
		}
		employeeID = &employee.ID
		preloads = append(preloads, "Employee")
	case "Owner":
		owner, err := f.currentUser.Owner(ctx)
		if err != nil {
			return nil, eris.Wrap(err, "failed to retrieve owner user")
		}
		ownerID = &owner.ID
		preloads = append(preloads, "Owner")
	case "Member":
		member, err := f.currentUser.Member(ctx)
		if err != nil {
			return nil, eris.Wrap(err, "failed to retrieve member user")
		}
		memberID = &member.ID
		preloads = append(preloads, "Member")
	default:
		return nil, eris.New("invalid account type for footstep creation")
	}

	footstep := &models.Footstep{
		ID:             uuid.New(),
		AccountType:    claims.AccountType,
		Module:         module,
		Activity:       activity,
		Description:    description,
		Latitude:       latitude,
		Longitude:      longitude,
		Timestamp:      time.Now(),
		IsDeleted:      false,
		IPAddress:      userInfo.IPAddress,
		UserAgent:      userInfo.UserAgent,
		Referer:        userInfo.Referer,
		Location:       userInfo.Location,
		AcceptLanguage: userInfo.AcceptLanguage,
		AdminID:        adminID,
		EmployeeID:     employeeID,
		OwnerID:        ownerID,
		MemberID:       memberID,
	}

	savedFootstep, err := f.currentUser.repository.FootstepCreate(footstep, preloads...)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create footstep")
	}

	return savedFootstep, nil
}
