package auth_accounts

import (
	"errors"
	"fmt"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/Lands-Horizon-Corp/horizon-corp/server/middleware"
	"go.uber.org/zap"
)

const SignedInExpiration = time.Hour * 12

type AuthAccount struct {
	cfg         *config.AppConfig
	engine      *providers.EngineService
	middle      *middleware.Middleware
	otpProvider *providers.OTPService

	smsProvider   *providers.SMSService
	emailProvider *providers.EmailService

	tokenProvider *providers.TokenService
	logger        *providers.LoggerService
	modelResource *models.ModelResource

	helpers       *helpers.HelpersFunction
	cryptoHelpers *helpers.HelpersCryptography
}

func NewAuthAccount(
	cfg *config.AppConfig,
	engine *providers.EngineService,
	middle *middleware.Middleware,
	otpProvider *providers.OTPService,

	tokenProvider *providers.TokenService,
	smsProvider *providers.SMSService,
	emailProvider *providers.EmailService,

	logger *providers.LoggerService,
	modelResource *models.ModelResource,

	helpers *helpers.HelpersFunction,
	cryptoHelpers *helpers.HelpersCryptography,

) *AuthAccount {
	return &AuthAccount{
		cfg:           cfg,
		engine:        engine,
		middle:        middle,
		otpProvider:   otpProvider,
		tokenProvider: tokenProvider,
		smsProvider:   smsProvider,
		emailProvider: emailProvider,
		modelResource: modelResource,
		helpers:       helpers,
		cryptoHelpers: cryptoHelpers,
	}
}

func (at *AuthAccount) GenerateUserToken(id uint, accountType string, expiration time.Duration) (*string, error) {
	validTypes := []string{"Member", "Employee", "Admin", "Owner"}
	isValid := false
	for _, validType := range validTypes {
		if accountType == validType {
			isValid = true
			break
		}
	}

	if !isValid {
		return nil, errors.New("invalid account type")
	}
	claims := &providers.UserClaims{
		ID:          id,
		AccountType: accountType,
	}
	token, err := at.tokenProvider.GenerateToken(claims, expiration)
	if err != nil {
		at.logger.Error("Failed to generate admin token", zap.Error(err))
		return nil, err
	}
	return &token, nil
}

func (ap *AuthAccount) Create(user interface{}, accountType string) (uint, error) {
	if user == nil {
		return 0, errors.New("user cannot be nil")
	}

	switch accountType {
	case "Employee":
		employee, ok := user.(*models.Employee)
		if !ok {
			return 0, errors.New("invalid user type for Employee")
		}
		if err := ap.modelResource.EmployeeCreate(employee); err != nil {
			return 0, err
		}
		return employee.ID, nil

	case "Member":
		member, ok := user.(*models.Member)
		if !ok {
			return 0, errors.New("invalid user type for Member")
		}
		if err := ap.modelResource.MemberCreate(member); err != nil {
			return 0, err
		}
		return member.ID, nil

	case "Owner":
		owner, ok := user.(*models.Owner)
		if !ok {
			return 0, errors.New("invalid user type for Owner")
		}
		if err := ap.modelResource.OwnerCreate(owner); err != nil {
			return 0, err
		}
		return owner.ID, nil

	case "Admin":
		admin, ok := user.(*models.Admin)
		if !ok {
			return 0, errors.New("invalid user type for Admin")
		}
		if err := ap.modelResource.AdminCreate(admin); err != nil {
			return 0, err
		}
		return admin.ID, nil

	default:
		return 0, errors.New("unsupported account type")
	}
}

func (ap *AuthAccount) GetByID(accountType string, id uint) (interface{}, error) {
	switch accountType {
	case "Admin":
		preloads := []string{"Media", "Role", "Gender"}
		admin, err := ap.modelResource.AdminDB.FindByID(id, preloads...)
		return ap.modelResource.AdminToResource(admin), err
	case "Owner":
		preloads := []string{"Media", "Company", "Gender"}
		owner, err := ap.modelResource.OwnerDB.FindByID(id, preloads...)
		return ap.modelResource.OwnerToResource(owner), err
	case "Employee":
		preloads := []string{"Media", "Branch", "Role", "Gender"}
		owner, err := ap.modelResource.EmployeeDB.FindByID(id, preloads...)
		return ap.modelResource.EmployeeToResource(owner), err
	case "Member":
		preloads := []string{"Media", "Branch", "Role", "Gender"}
		owner, err := ap.modelResource.MemberDB.FindByID(id, preloads...)
		return ap.modelResource.MemberToResource(owner), err
	default:
		return nil, errors.New("invalid account type")
	}
}

func (ap *AuthAccount) Update(user interface{}, accountType string, preloads []string) (interface{}, error) {
	if user == nil {
		return nil, errors.New("user cannot be nil")
	}
	switch accountType {
	case "Employee":
		employee, ok := user.(*models.Employee)
		if !ok {
			return nil, errors.New("invalid user type for Employee")
		}
		if err := ap.modelResource.EmployeeUpdate(employee, preloads); err != nil {
			return nil, err
		}
		return ap.modelResource.EmployeeToResource(employee), nil

	case "Member":
		member, ok := user.(*models.Member)
		if !ok {
			return nil, errors.New("invalid user type for Member")
		}
		if err := ap.modelResource.MemberUpdate(member, preloads); err != nil {
			return nil, err
		}
		return ap.modelResource.MemberToResource(member), nil
	case "Owner":
		owner, ok := user.(*models.Owner)
		if !ok {
			return nil, errors.New("invalid user type for Owner")
		}
		if err := ap.modelResource.OwnerUpdate(owner, preloads); err != nil {
			return nil, err
		}
		return ap.modelResource.OwnerToResource(owner), nil

	case "Admin":
		admin, ok := user.(*models.Admin)
		if !ok {
			return nil, errors.New("invalid user type for Admin")
		}
		if err := ap.modelResource.AdminUpdate(admin, preloads); err != nil {
			return nil, err
		}
		return ap.modelResource.AdminToResource(admin), nil
	default:
		return nil, errors.New("unsupported account type")
	}
}

func (ap *AuthAccount) Delete(accountType string, id uint) error {
	switch accountType {
	case "Admin":
		return ap.modelResource.AdminDB.Delete(id)
	case "Owner":
		return ap.modelResource.OwnerDB.Delete(id)
	case "Employee":
		return ap.modelResource.EmployeeDB.Delete(id)
	case "Member":
		return ap.modelResource.MemberDB.Delete(id)
	default:
		return errors.New("invalid account type")
	}
}

func (ap *AuthAccount) FindByEmailUsernameOrContact(accountType, input string) (interface{}, error) {
	switch accountType {
	case "Admin":
		user, err := ap.modelResource.AdminFindByEmailUsernameOrContact(accountType)
		if err != nil {
			return nil, err
		}
		return ap.modelResource.AdminToResource(user), nil
	case "Owner":
		user, err := ap.modelResource.OwnerFindByEmailUsernameOrContact(accountType)
		if err != nil {
			return nil, err
		}
		return ap.modelResource.OwnerToResource(user), nil
	case "Employee":
		user, err := ap.modelResource.EmployeeFindByEmailUsernameOrContact(accountType)
		if err != nil {
			return nil, err
		}
		return ap.modelResource.EmployeeToResource(user), nil
	case "Member":
		user, err := ap.modelResource.MemberFindByEmailUsernameOrContact(accountType)
		if err != nil {
			return nil, err
		}
		return ap.modelResource.MemberToResource(user), nil
	default:
		return nil, errors.New("invalid account type")
	}
}

func (ap *AuthAccount) FindByEmailUsernameOrContactForPassword(accountType, input string) (uint, string, error) {
	switch accountType {
	case "Admin":
		user, err := ap.modelResource.AdminFindByEmailUsernameOrContact(input)
		if err != nil {
			return 0, "", err
		}
		return user.ID, user.Password, nil
	case "Owner":
		user, err := ap.modelResource.OwnerFindByEmailUsernameOrContact(input)
		if err != nil {
			return 0, "", err
		}
		return user.ID, user.Password, nil
	case "Employee":
		user, err := ap.modelResource.EmployeeFindByEmailUsernameOrContact(input)
		if err != nil {
			return 0, "", err
		}
		return user.ID, user.Password, nil
	case "Member":
		user, err := ap.modelResource.MemberFindByEmailUsernameOrContact(input)
		if err != nil {
			return 0, "", err
		}
		return user.ID, user.Password, nil
	default:
		return 0, "", errors.New("invalid account type")
	}
}

func (ap *AuthAccount) FindByEmailUsernameOrContactForID(accountType, input string) (uint, error) {
	switch accountType {
	case "Admin":
		user, err := ap.modelResource.AdminFindByEmailUsernameOrContact(input)
		if err != nil {
			return 0, err
		}
		return user.ID, nil
	case "Owner":
		user, err := ap.modelResource.OwnerFindByEmailUsernameOrContact(input)
		if err != nil {
			return 0, err
		}
		return user.ID, nil
	case "Employee":
		user, err := ap.modelResource.EmployeeFindByEmailUsernameOrContact(input)
		if err != nil {
			return 0, err
		}
		return user.ID, nil
	case "Member":
		user, err := ap.modelResource.MemberFindByEmailUsernameOrContact(input)
		if err != nil {
			return 0, err
		}
		return user.ID, nil
	default:
		return 0, errors.New("invalid account type")
	}
}

func (ap *AuthAccount) FindByEmailUsernameOrContactForName(accountType, input string) (string, error) {
	switch accountType {
	case "Admin":
		user, err := ap.modelResource.AdminFindByEmailUsernameOrContact(input)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s %s", user.FirstName, user.LastName), nil
	case "Owner":
		user, err := ap.modelResource.OwnerFindByEmailUsernameOrContact(input)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s %s", user.FirstName, user.LastName), nil
	case "Employee":
		user, err := ap.modelResource.EmployeeFindByEmailUsernameOrContact(input)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s %s", user.FirstName, user.LastName), nil
	case "Member":
		user, err := ap.modelResource.MemberFindByEmailUsernameOrContact(input)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s %s", user.FirstName, user.LastName), nil
	default:
		return "", errors.New("invalid account type")
	}
}

func (ap *AuthAccount) UpdatePassword(accountType string, userID uint, password string) error {
	switch accountType {
	case "Admin":
		return ap.modelResource.AdminUpdatePassword(userID, password)
	case "Owner":
		return ap.modelResource.OwnerUpdatePassword(userID, password)
	case "Employee":
		return ap.modelResource.EmployeeUpdatePassword(userID, password)
	case "Member":
		return ap.modelResource.MemberUpdatePassword(userID, password)
	default:
		return fmt.Errorf("invalid account type")
	}
}

func (ap *AuthAccount) UpdateVerification(accountType string, userID uint, verificationType string, value bool) (interface{}, error) {
	var preloads []string
	var err error
	var resource interface{}

	switch accountType {
	case "Admin":
		preloads = []string{"Media", "Role", "Gender"}
		admin := &models.Admin{}
		switch verificationType {
		case "email":
			admin.IsEmailVerified = value
		case "contact":
			admin.IsContactVerified = value
		case "skip":
			admin.IsSkipVerification = value
		default:
			return nil, fmt.Errorf("invalid verification type")
		}
		err = ap.modelResource.AdminUpdate(admin, preloads)
		resource = ap.modelResource.AdminToResource(admin)

	case "Owner":
		preloads = []string{"Media", "Company", "Gender"}
		owner := &models.Owner{}
		switch verificationType {
		case "email":
			owner.IsEmailVerified = value
		case "contact":
			owner.IsContactVerified = value
		case "skip":
			owner.IsSkipVerification = value
		default:
			return nil, fmt.Errorf("invalid verification type")
		}
		err = ap.modelResource.OwnerUpdate(owner, preloads)
		resource = ap.modelResource.OwnerToResource(owner)

	case "Employee":
		preloads = []string{"Media", "Branch", "Role", "Gender"}
		employee := &models.Employee{}
		switch verificationType {
		case "email":
			employee.IsEmailVerified = value
		case "contact":
			employee.IsContactVerified = value
		case "skip":
			employee.IsSkipVerification = value
		default:
			return nil, fmt.Errorf("invalid verification type")
		}
		err = ap.modelResource.EmployeeUpdate(employee, preloads)
		resource = ap.modelResource.EmployeeToResource(employee)

	case "Member":
		preloads = []string{"Media", "Branch", "Role", "Gender"}
		member := &models.Member{}
		switch verificationType {
		case "email":
			member.IsEmailVerified = value
		case "contact":
			member.IsContactVerified = value
		case "skip":
			member.IsSkipVerification = value
		default:
			return nil, fmt.Errorf("invalid verification type")
		}
		err = ap.modelResource.MemberUpdate(member, preloads)
		resource = ap.modelResource.MemberToResource(member)

	default:
		return nil, fmt.Errorf("invalid account type")
	}

	if err != nil {
		return nil, err
	}

	return resource, nil
}
