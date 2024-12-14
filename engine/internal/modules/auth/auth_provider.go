package auth

import (
	"errors"
	"fmt"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/go-playground/validator"
)

type SignUpRequest struct {
	AccountType string `json:"accountType" validate:"required,max=10"`
}

type AuthProvider struct {
	cfg           *config.AppConfig
	cryptoHelpers *helpers.HelpersCryptography
	modelResource *models.ModelResource
}

func NewAuthProvider(
	cfg *config.AppConfig,
	cryptoHelpers *helpers.HelpersCryptography,
	modelResource *models.ModelResource,
) *AuthProvider {
	return &AuthProvider{
		cfg:           cfg,
		cryptoHelpers: cryptoHelpers,
		modelResource: modelResource,
	}
}

func (ap *AuthProvider) AccountTypeValidator(fl validator.FieldLevel) bool {
	accountType := fl.Field().String()
	validTypes := []string{"Member", "Employee", "Admin", "Owner"}
	for _, validType := range validTypes {
		if accountType == validType {
			return true
		}
	}
	return false
}

func (ap *AuthProvider) Validate(r SignUpRequest) error {
	validate := validator.New()
	err := validate.Struct(r)
	if err := validate.RegisterValidation("accountType", ap.AccountTypeValidator); err != nil {
		return fmt.Errorf("failed to register account type validator: %v", err)
	}
	if err != nil {
		return err
	}
	return nil
}

func (ap *AuthProvider) Create(user interface{}, accountType string) (uint, error) {
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

func (ap *AuthProvider) GetByID(accountType string, id uint) (interface{}, error) {
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

func (ap *AuthProvider) Update(user interface{}, accountType string, preloads []string) (interface{}, error) {
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

func (ap *AuthProvider) Delete(accountType string, id uint) error {
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

func (ap *AuthProvider) FindByEmailUsernameOrContact(accountType, input string) (interface{}, error) {
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

func (ap *AuthProvider) UpdatePassword(accountType string, userID uint, password string) error {
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

func (ap *AuthProvider) UpdateVerification(accountType string, userID uint, verificationType string, value bool) (interface{}, error) {
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
