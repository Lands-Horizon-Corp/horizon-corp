package repositories

import (
	"errors"
	"fmt"
	"horizon/server/config"
	"horizon/server/internal/models"
)

type UserRepository struct {
	admin    *AdminRepository
	owner    *OwnerRepository
	employee *EmployeeRepository
	member   *MemberRepository
}

func NewUserRepository(
	admin *AdminRepository,
	owner *OwnerRepository,
	employee *EmployeeRepository,
	member *MemberRepository,
) *UserRepository {
	return &UserRepository{
		admin:    admin,
		owner:    owner,
		employee: employee,
		member:   member,
	}
}

func (r *UserRepository) Create(accountType string, account models.User) (models.User, error) {
	password, err := config.HashPassword(account.Password)
	if err != nil {
		return models.User{}, err
	}

	switch accountType {
	case "Admin":
		admin := ConvertUserToAdmin(account)
		admin.Password = password
		if err := r.admin.Create(&admin); err != nil {
			return models.User{}, err
		}
		return ConvertAdminToUser(admin), nil
	case "Owner":
		owner := ConvertUserToOwner(account)
		owner.Password = password
		if err := r.owner.Create(&owner); err != nil {
			return models.User{}, err
		}
		return ConvertOwnerToUser(owner), nil
	case "Employee":
		employee := ConvertUserToEmployee(account)
		employee.Password = password
		if err := r.employee.Create(&employee); err != nil {
			return models.User{}, err
		}
		return ConvertEmployeeToUser(employee), nil
	case "Member":
		member := ConvertUserToMember(account)
		member.Password = password
		if err := r.member.Create(&member); err != nil {
			return models.User{}, err
		}
		return ConvertMemberToUser(member), nil
	default:
		return models.User{}, errors.New("invalid account type")
	}
}

func (r *UserRepository) GetAll(accountType string) ([]models.User, error) {
	switch accountType {
	case "Admin":
		admins, err := r.admin.GetAll()
		return convertToUserSliceAdmin(admins), err
	case "Owner":
		owners, err := r.owner.GetAll()
		return convertToUserSliceOwner(owners), err
	case "Employee":
		employees, err := r.employee.GetAll()
		return convertToUserSliceEmployee(employees), err
	case "Member":
		members, err := r.member.GetAll()
		return convertToUserSliceMember(members), err
	default:
		return nil, errors.New("invalid account type")
	}
}

func (r *UserRepository) GetByID(accountType string, id uint) (models.User, error) {
	switch accountType {
	case "Admin":
		admin, err := r.admin.GetByID(id)
		return ConvertAdminToUser(admin), err
	case "Owner":
		owner, err := r.owner.GetByID(id)
		return ConvertOwnerToUser(owner), err
	case "Employee":
		employee, err := r.employee.GetByID(id)
		return ConvertEmployeeToUser(employee), err
	case "Member":
		member, err := r.member.GetByID(id)
		return ConvertMemberToUser(member), err
	default:
		return models.User{}, errors.New("invalid account type")
	}
}

func (r *UserRepository) Update(accountType string, id uint, account models.User) (models.User, error) {
	password, err := config.HashPassword(account.Password)
	if err != nil {
		return models.User{}, err
	}

	switch accountType {
	case "Admin":
		admin := ConvertUserToAdmin(account)
		admin.Model.ID = id
		admin.Password = password
		if err := r.admin.Update(id, &admin); err != nil {
			return models.User{}, err
		}
		return ConvertAdminToUser(admin), nil
	case "Owner":
		owner := ConvertUserToOwner(account)
		owner.Model.ID = id
		owner.Password = password
		if err := r.owner.Update(id, &owner); err != nil {
			return models.User{}, err
		}
		return ConvertOwnerToUser(owner), nil
	case "Employee":
		employee := ConvertUserToEmployee(account)
		employee.Model.ID = id
		employee.Password = password
		if err := r.employee.Update(id, &employee); err != nil {
			return models.User{}, err
		}
		return ConvertEmployeeToUser(employee), nil
	case "Member":
		member := ConvertUserToMember(account)
		member.Model.ID = id
		member.Password = password
		if err := r.member.Update(id, &member); err != nil {
			return models.User{}, err
		}
		return ConvertMemberToUser(member), nil
	default:
		return models.User{}, errors.New("invalid account type")
	}
}

func (r *UserRepository) Delete(accountType string, id uint) error {
	switch accountType {
	case "Admin":
		return r.admin.Delete(id)
	case "Owner":
		return r.owner.Delete(id)
	case "Employee":
		return r.employee.Delete(id)
	case "Member":
		return r.member.Delete(id)
	default:
		return errors.New("invalid account type")
	}
}

func (r *UserRepository) FindByEmailUsernameOrContact(accountType string, input string) (models.User, error) {
	switch accountType {
	case "Admin":
		admin, err := r.admin.FindByEmailUsernameOrContact(input)
		return ConvertAdminToUser(admin), err
	case "Owner":
		owner, err := r.owner.FindByEmailUsernameOrContact(input)
		return ConvertOwnerToUser(owner), err
	case "Employee":
		employee, err := r.employee.FindByEmailUsernameOrContact(input)
		return ConvertEmployeeToUser(employee), err
	case "Member":
		member, err := r.member.FindByEmailUsernameOrContact(input)
		return ConvertMemberToUser(member), err
	default:
		return models.User{}, errors.New("invalid account type")
	}
}

func (r *UserRepository) UpdateColumns(accountType string, id uint, columns map[string]interface{}) (models.User, error) {
	switch accountType {
	case "Admin":
		admin, err := r.admin.UpdateColumns(id, columns)
		return ConvertAdminToUser(admin), err
	case "Owner":
		owner, err := r.owner.UpdateColumns(id, columns)
		return ConvertOwnerToUser(owner), err
	case "Employee":
		employee, err := r.employee.UpdateColumns(id, columns)
		return ConvertEmployeeToUser(employee), err
	case "Member":
		member, err := r.member.UpdateColumns(id, columns)
		return ConvertMemberToUser(member), err
	default:
		return models.User{}, errors.New("invalid account type")
	}
}

func (r *UserRepository) UpdatePassword(accountType string, userID uint, password string) error {
	newPassword, err := config.HashPassword(password)
	if err != nil {
		return err
	}
	switch accountType {
	case "Admin":
		_, err := r.admin.UpdateColumns(userID, map[string]interface{}{
			"password": newPassword,
		})
		return err
	case "Owner":
		_, err := r.owner.UpdateColumns(userID, map[string]interface{}{
			"password": newPassword,
		})
		return err
	case "Employee":
		_, err := r.employee.UpdateColumns(userID, map[string]interface{}{
			"password": newPassword,
		})
		return err
	case "Member":
		_, err := r.member.UpdateColumns(userID, map[string]interface{}{
			"password": newPassword,
		})
		return err
	default:
		return fmt.Errorf("invalid account type")
	}
}

// Conversion functions to transform specific models to models.User
func ConvertAdminToUser(admin models.Admin) models.User {
	return models.User{
		Model:              admin.Model,
		FirstName:          admin.FirstName,
		LastName:           admin.LastName,
		MiddleName:         admin.MiddleName,
		PermanentAddress:   admin.PermanentAddress,
		Description:        admin.Description,
		Birthdate:          admin.Birthdate,
		Username:           admin.Username,
		Email:              admin.Email,
		Password:           admin.Password,
		MediaID:            admin.MediaID,
		IsEmailVerified:    admin.IsEmailVerified,
		IsContactVerified:  admin.IsContactVerified,
		IsSkipVerification: admin.IsSkipVerification,
		ContactNumber:      admin.ContactNumber,
		Media:              admin.Media,
		Status:             string(admin.Status),
	}
}

func ConvertOwnerToUser(owner models.Owner) models.User {
	return models.User{
		Model:              owner.Model,
		FirstName:          owner.FirstName,
		LastName:           owner.LastName,
		MiddleName:         owner.MiddleName,
		PermanentAddress:   owner.PermanentAddress,
		Description:        owner.Description,
		Birthdate:          owner.Birthdate,
		Username:           owner.Username,
		Email:              owner.Email,
		Password:           owner.Password,
		MediaID:            owner.MediaID,
		IsEmailVerified:    owner.IsEmailVerified,
		IsContactVerified:  owner.IsContactVerified,
		IsSkipVerification: owner.IsSkipVerification,
		ContactNumber:      owner.ContactNumber,
		Media:              owner.Media,
		Status:             string(owner.Status),
	}
}

func ConvertEmployeeToUser(employee models.Employee) models.User {
	return models.User{
		Model:              employee.Model,
		FirstName:          employee.FirstName,
		LastName:           employee.LastName,
		MiddleName:         employee.MiddleName,
		PermanentAddress:   employee.PermanentAddress,
		Description:        employee.Description,
		Birthdate:          employee.Birthdate,
		Username:           employee.Username,
		Email:              employee.Email,
		Password:           employee.Password,
		MediaID:            employee.MediaID,
		IsEmailVerified:    employee.IsEmailVerified,
		IsContactVerified:  employee.IsContactVerified,
		IsSkipVerification: employee.IsSkipVerification,
		ContactNumber:      employee.ContactNumber,
		Media:              employee.Media,
		Status:             string(employee.Status),
	}
}

func ConvertMemberToUser(member models.Member) models.User {
	return models.User{
		Model:              member.Model,
		FirstName:          member.FirstName,
		LastName:           member.LastName,
		MiddleName:         member.MiddleName,
		PermanentAddress:   member.PermanentAddress,
		Description:        member.Description,
		Birthdate:          member.Birthdate,
		Username:           member.Username,
		Email:              member.Email,
		Password:           member.Password,
		MediaID:            member.MediaID,
		IsEmailVerified:    member.IsEmailVerified,
		IsContactVerified:  member.IsContactVerified,
		IsSkipVerification: member.IsSkipVerification,
		ContactNumber:      member.ContactNumber,
		Media:              member.Media,
		Status:             string(member.Status),
	}
}

// ConvertUserToAdmin converts a models.User to models.Admin
func ConvertUserToAdmin(user models.User) models.Admin {
	return models.Admin{
		Model:              user.Model,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		MiddleName:         user.MiddleName,
		PermanentAddress:   user.PermanentAddress,
		Description:        user.Description,
		Birthdate:          user.Birthdate,
		Username:           user.Username,
		Email:              user.Email,
		Password:           user.Password,
		MediaID:            user.MediaID,
		IsEmailVerified:    user.IsEmailVerified,
		IsContactVerified:  user.IsContactVerified,
		IsSkipVerification: user.IsSkipVerification,
		ContactNumber:      user.ContactNumber,
		Status:             models.AdminStatus(user.Status), // Assuming Status is an AdminStatus
	}
}

// ConvertUserToOwner converts a models.User to models.Owner
func ConvertUserToOwner(user models.User) models.Owner {
	return models.Owner{
		Model:              user.Model,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		MiddleName:         user.MiddleName,
		PermanentAddress:   user.PermanentAddress,
		Description:        user.Description,
		Birthdate:          user.Birthdate,
		Username:           user.Username,
		Email:              user.Email,
		Password:           user.Password,
		MediaID:            user.MediaID,
		IsEmailVerified:    user.IsEmailVerified,
		IsContactVerified:  user.IsContactVerified,
		IsSkipVerification: user.IsSkipVerification,
		ContactNumber:      user.ContactNumber,
		Status:             models.OwnerStatus(user.Status), // Assuming Status is an OwnerStatus
	}
}

// ConvertUserToEmployee converts a models.User to models.Employee
func ConvertUserToEmployee(user models.User) models.Employee {
	return models.Employee{
		Model:              user.Model,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		MiddleName:         user.MiddleName,
		PermanentAddress:   user.PermanentAddress,
		Description:        user.Description,
		Birthdate:          user.Birthdate,
		Username:           user.Username,
		Email:              user.Email,
		Password:           user.Password,
		MediaID:            user.MediaID,
		IsEmailVerified:    user.IsEmailVerified,
		IsContactVerified:  user.IsContactVerified,
		IsSkipVerification: user.IsSkipVerification,
		ContactNumber:      user.ContactNumber,
		Status:             models.EmployeeStatus(user.Status), // Assuming Status is an EmployeeStatus
	}
}

// ConvertUserToMember converts a models.User to models.Member
func ConvertUserToMember(user models.User) models.Member {
	return models.Member{
		Model:              user.Model,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		MiddleName:         user.MiddleName,
		PermanentAddress:   user.PermanentAddress,
		Description:        user.Description,
		Birthdate:          user.Birthdate,
		Username:           user.Username,
		Email:              user.Email,
		Password:           user.Password,
		MediaID:            user.MediaID,
		IsEmailVerified:    user.IsEmailVerified,
		IsContactVerified:  user.IsContactVerified,
		IsSkipVerification: user.IsSkipVerification,
		ContactNumber:      user.ContactNumber,
		Status:             models.MemberStatus(user.Status), // Assuming Status is a MemberStatus
	}
}

func convertToUserSliceAdmin(admins []models.Admin) []models.User {
	users := make([]models.User, len(admins))
	for i, admin := range admins {
		users[i] = ConvertAdminToUser(admin)
	}
	return users
}

func convertToUserSliceOwner(owners []models.Owner) []models.User {
	users := make([]models.User, len(owners))
	for i, owner := range owners {
		users[i] = ConvertOwnerToUser(owner)
	}
	return users
}

func convertToUserSliceEmployee(employees []models.Employee) []models.User {
	users := make([]models.User, len(employees))
	for i, employee := range employees {
		users[i] = ConvertEmployeeToUser(employee)
	}
	return users
}

func convertToUserSliceMember(members []models.Member) []models.User {
	users := make([]models.User, len(members))
	for i, member := range members {
		users[i] = ConvertMemberToUser(member)
	}
	return users
}
