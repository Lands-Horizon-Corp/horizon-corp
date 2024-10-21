package auth

import (
	"errors"
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"

	"go.uber.org/zap"
)

type UserAuthService struct {
	userRepo *repositories.UserRepository

	adminAuthService    *AdminAuthService
	employeeAuthService *EmployeeAuthService
	memberAuthService   *MemberAuthService
	ownerAuthService    *OwnerAuthService

	logger *zap.Logger
}

// NewUserAuthService initializes a new UserAuthService
func NewUserAuthService(
	userRepo *repositories.UserRepository,
	adminAuthService *AdminAuthService,
	employeeAuthService *EmployeeAuthService,
	memberAuthService *MemberAuthService,
	ownerAuthService *OwnerAuthService,
	logger *zap.Logger,
) *UserAuthService {
	return &UserAuthService{
		userRepo:            userRepo,
		adminAuthService:    adminAuthService,
		employeeAuthService: employeeAuthService,
		memberAuthService:   memberAuthService,
		ownerAuthService:    ownerAuthService,
		logger:              logger,
	}
}

// GenerateUserToken generates a token for the specified user based on their account type
func (s *UserAuthService) GenerateUserToken(user models.User, accountType string) (string, error) {
	var token string
	var err error
	switch accountType {
	case "Admin":
		token, err = s.adminAuthService.GenerateAdminToken(repositories.ConvertUserToAdmin(user))
	case "Employee":
		token, err = s.employeeAuthService.GenerateEmployeeToken(repositories.ConvertUserToEmployee(user))
	case "Member":
		token, err = s.memberAuthService.GenerateMemberToken(repositories.ConvertUserToMember(user))
	case "Owner":
		token, err = s.ownerAuthService.GenerateOwnerToken(repositories.ConvertUserToOwner(user))
	default:
		s.logger.Error("Invalid account type for token generation", zap.String("accountType", accountType))
		return "", errors.New("invalid account type")
	}

	if err != nil {
		s.logger.Error("Failed to generate user token", zap.Error(err))
		return "", err
	}

	return token, nil
}
