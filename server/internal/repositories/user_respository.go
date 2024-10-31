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

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	hashedPassword, err := config.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	switch user.AccountType {
	case "Admin":
		admin := ConvertUserToAdmin(user)
		if err := r.admin.Create(admin); err != nil {
			return nil, err
		}
		return ConvertAdminToUser(admin), nil
	case "Owner":
		owner := ConvertUserToOwner(user)
		if err := r.owner.Create(owner); err != nil {
			return nil, err
		}
		return ConvertOwnerToUser(owner), nil
	case "Employee":
		employee := ConvertUserToEmployee(user)
		if err := r.employee.Create(employee); err != nil {
			return nil, err
		}
		return ConvertEmployeeToUser(employee), nil
	case "Member":
		member := ConvertUserToMember(user)
		if err := r.member.Create(member); err != nil {
			return nil, err
		}
		return ConvertMemberToUser(member), nil
	default:
		return nil, errors.New("invalid account type")
	}
}

func (r *UserRepository) GetAll(accountType string) ([]*models.User, error) {
	preloads := []string{"Media", "Branch", "Role", "Gender"}
	switch accountType {
	case "Admin":
		admins, err := r.admin.GetAll(preloads)
		return convertToUserSliceAdmin(admins), err
	case "Owner":
		owners, err := r.owner.GetAll(preloads)
		return convertToUserSliceOwner(owners), err
	case "Employee":
		employees, err := r.employee.GetAll(preloads)
		return convertToUserSliceEmployee(employees), err
	case "Member":
		members, err := r.member.GetAll(preloads)
		return convertToUserSliceMember(members), err
	default:
		return nil, errors.New("invalid account type")
	}
}

func (r *UserRepository) GetByID(accountType string, id uint) (*models.User, error) {
	preloads := []string{"Media", "Branch", "Role", "Gender"}
	switch accountType {
	case "Admin":

		admin, err := r.admin.GetByID(id, preloads)
		if err != nil {
			return nil, err
		}
		return ConvertAdminToUser(admin), nil
	case "Owner":
		owner, err := r.owner.GetByID(id, preloads)
		if err != nil {
			return nil, err
		}
		return ConvertOwnerToUser(owner), nil
	case "Employee":
		employee, err := r.employee.GetByID(id, preloads)
		if err != nil {
			return nil, err
		}
		return ConvertEmployeeToUser(employee), nil
	case "Member":
		member, err := r.member.GetByID(id, preloads)
		if err != nil {
			return nil, err
		}
		return ConvertMemberToUser(member), nil
	default:
		return nil, errors.New("invalid account type")
	}
}

func (r *UserRepository) Update(user *models.User) error {
	if user.Password != "" {
		hashedPassword, err := config.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}
	preloads := []string{"Media", "Branch", "Role", "Gender"}
	switch user.AccountType {
	case "Admin":
		admin := ConvertUserToAdmin(user)
		return r.admin.Update(admin, preloads)
	case "Owner":
		owner := ConvertUserToOwner(user)
		return r.owner.Update(owner, preloads)
	case "Employee":
		employee := ConvertUserToEmployee(user)
		return r.employee.Update(employee, preloads)
	case "Member":
		member := ConvertUserToMember(user)
		return r.member.Update(member, preloads)
	default:
		return errors.New("invalid account type")
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

func (r *UserRepository) FindByEmailUsernameOrContact(accountType, input string) (*models.User, error) {
	switch accountType {
	case "Admin":
		admin, err := r.admin.FindByEmailUsernameOrContact(input)
		if err != nil {
			return nil, err
		}
		return ConvertAdminToUser(admin), nil
	case "Owner":
		owner, err := r.owner.FindByEmailUsernameOrContact(input)
		if err != nil {
			return nil, err
		}
		return ConvertOwnerToUser(owner), nil
	case "Employee":
		employee, err := r.employee.FindByEmailUsernameOrContact(input)
		if err != nil {
			return nil, err
		}
		return ConvertEmployeeToUser(employee), nil
	case "Member":
		member, err := r.member.FindByEmailUsernameOrContact(input)
		if err != nil {
			return nil, err
		}
		return ConvertMemberToUser(member), nil
	default:
		return nil, errors.New("invalid account type")
	}
}
func (r *UserRepository) UpdateColumns(id uint, user *models.User) (*models.User, error) {
	// Hash the password if it's provided
	if user.Password != "" {
		hashedPassword, err := config.HashPassword(user.Password)
		if err != nil {
			return nil, fmt.Errorf("error hashing password: %v", err)
		}
		user.Password = hashedPassword
	}

	preloads := []string{"Media", "Branch", "Role", "Gender"}
	switch user.AccountType {
	case "Admin":
		admin := ConvertUserToAdmin(user)
		updatedAdmin, err := r.admin.UpdateColumns(id, *admin, preloads)
		if err != nil {
			return nil, err
		}
		return ConvertAdminToUser(updatedAdmin), nil
	case "Owner":
		owner := ConvertUserToOwner(user)
		updatedOwner, err := r.owner.UpdateColumns(id, *owner, preloads)
		if err != nil {
			return nil, err
		}
		return ConvertOwnerToUser(updatedOwner), nil
	case "Employee":
		employee := ConvertUserToEmployee(user)
		updatedEmployee, err := r.employee.UpdateColumns(id, *employee, preloads)
		if err != nil {
			return nil, err
		}
		return ConvertEmployeeToUser(updatedEmployee), nil
	case "Member":
		member := ConvertUserToMember(user)
		updatedMember, err := r.member.UpdateColumns(id, *member, preloads)
		if err != nil {
			return nil, err
		}
		return ConvertMemberToUser(updatedMember), nil
	default:
		return nil, errors.New("invalid account type")
	}
}

func (r *UserRepository) UpdatePassword(accountType string, userID uint, password string) error {
	switch accountType {
	case "Admin":
		return r.admin.UpdatePassword(userID, password)
	case "Owner":
		return r.owner.UpdatePassword(userID, password)
	case "Employee":
		return r.employee.UpdatePassword(userID, password)
	case "Member":
		return r.member.UpdatePassword(userID, password)
	default:
		return fmt.Errorf("invalid account type")
	}
}

// ConvertUserToAdmin converts a User to an Admin model
func ConvertUserToAdmin(user *models.User) *models.Admin {
	return &models.Admin{
		Model:              user.Model,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		MiddleName:         user.MiddleName,
		PermanentAddress:   user.PermanentAddress,
		Description:        user.Description,
		BirthDate:          user.BirthDate,
		Username:           user.Username,
		Email:              user.Email,
		Password:           user.Password,
		ContactNumber:      user.ContactNumber,
		IsEmailVerified:    user.IsEmailVerified,
		IsContactVerified:  user.IsContactVerified,
		IsSkipVerification: user.IsSkipVerification,
		MediaID:            user.MediaID,
		Media:              user.Media,
		Status:             models.AdminStatus(user.Status),
	}
}

func ConvertUserToOwner(user *models.User) *models.Owner {
	return &models.Owner{
		Model:              user.Model,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		MiddleName:         user.MiddleName,
		PermanentAddress:   user.PermanentAddress,
		Description:        user.Description,
		BirthDate:          user.BirthDate,
		Username:           user.Username,
		Email:              user.Email,
		Password:           user.Password,
		ContactNumber:      user.ContactNumber,
		IsEmailVerified:    user.IsEmailVerified,
		IsContactVerified:  user.IsContactVerified,
		IsSkipVerification: user.IsSkipVerification,
		MediaID:            user.MediaID,
		Media:              user.Media,
		Status:             models.OwnerStatus(user.Status),
	}
}

func ConvertUserToEmployee(user *models.User) *models.Employee {
	return &models.Employee{
		Model:              user.Model,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		MiddleName:         user.MiddleName,
		PermanentAddress:   user.PermanentAddress,
		Description:        user.Description,
		BirthDate:          user.BirthDate,
		Username:           user.Username,
		Email:              user.Email,
		Password:           user.Password,
		ContactNumber:      user.ContactNumber,
		IsEmailVerified:    user.IsEmailVerified,
		IsContactVerified:  user.IsContactVerified,
		IsSkipVerification: user.IsSkipVerification,
		MediaID:            user.MediaID,
		Media:              user.Media,
		Status:             models.EmployeeStatus(user.Status),
	}
}

func ConvertUserToMember(user *models.User) *models.Member {
	return &models.Member{
		Model:              user.Model,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		MiddleName:         user.MiddleName,
		PermanentAddress:   user.PermanentAddress,
		Description:        user.Description,
		BirthDate:          user.BirthDate,
		Username:           user.Username,
		Email:              user.Email,
		Password:           user.Password,
		ContactNumber:      user.ContactNumber,
		IsEmailVerified:    user.IsEmailVerified,
		IsContactVerified:  user.IsContactVerified,
		IsSkipVerification: user.IsSkipVerification,
		MediaID:            user.MediaID,
		Media:              user.Media,
		Status:             models.MemberStatus(user.Status),
	}
}

func ConvertAdminToUser(admin *models.Admin) *models.User {
	return &models.User{
		AccountType:        "Admin",
		Model:              admin.Model,
		FirstName:          admin.FirstName,
		LastName:           admin.LastName,
		MiddleName:         admin.MiddleName,
		PermanentAddress:   admin.PermanentAddress,
		Description:        admin.Description,
		BirthDate:          admin.BirthDate,
		Username:           admin.Username,
		Email:              admin.Email,
		Password:           admin.Password,
		ContactNumber:      admin.ContactNumber,
		IsEmailVerified:    admin.IsEmailVerified,
		IsContactVerified:  admin.IsContactVerified,
		IsSkipVerification: admin.IsSkipVerification,
		MediaID:            admin.MediaID,
		Media:              admin.Media,
		Status:             string(admin.Status),
	}
}
func ConvertOwnerToUser(owner *models.Owner) *models.User {
	return &models.User{
		AccountType:        "Owner",
		Model:              owner.Model,
		FirstName:          owner.FirstName,
		LastName:           owner.LastName,
		MiddleName:         owner.MiddleName,
		PermanentAddress:   owner.PermanentAddress,
		Description:        owner.Description,
		BirthDate:          owner.BirthDate,
		Username:           owner.Username,
		Email:              owner.Email,
		Password:           owner.Password,
		ContactNumber:      owner.ContactNumber,
		IsEmailVerified:    owner.IsEmailVerified,
		IsContactVerified:  owner.IsContactVerified,
		IsSkipVerification: owner.IsSkipVerification,
		MediaID:            owner.MediaID,
		Media:              owner.Media,
		Status:             string(owner.Status),
	}
}

func ConvertEmployeeToUser(employee *models.Employee) *models.User {
	return &models.User{
		AccountType:        "Employee",
		Model:              employee.Model,
		FirstName:          employee.FirstName,
		LastName:           employee.LastName,
		MiddleName:         employee.MiddleName,
		PermanentAddress:   employee.PermanentAddress,
		Description:        employee.Description,
		BirthDate:          employee.BirthDate,
		Username:           employee.Username,
		Email:              employee.Email,
		Password:           employee.Password,
		ContactNumber:      employee.ContactNumber,
		IsEmailVerified:    employee.IsEmailVerified,
		IsContactVerified:  employee.IsContactVerified,
		IsSkipVerification: employee.IsSkipVerification,
		MediaID:            employee.MediaID,
		Media:              employee.Media,
		Status:             string(employee.Status),
	}
}

func ConvertMemberToUser(member *models.Member) *models.User {
	return &models.User{
		AccountType:        "Member",
		Model:              member.Model,
		FirstName:          member.FirstName,
		LastName:           member.LastName,
		MiddleName:         member.MiddleName,
		PermanentAddress:   member.PermanentAddress,
		Description:        member.Description,
		BirthDate:          member.BirthDate,
		Username:           member.Username,
		Email:              member.Email,
		Password:           member.Password,
		ContactNumber:      member.ContactNumber,
		IsEmailVerified:    member.IsEmailVerified,
		IsContactVerified:  member.IsContactVerified,
		IsSkipVerification: member.IsSkipVerification,
		MediaID:            member.MediaID,
		Media:              member.Media,
		Status:             string(member.Status),
	}
}

func convertToUserSliceAdmin(admins []*models.Admin) []*models.User {
	users := make([]*models.User, len(admins))
	for i, admin := range admins {
		users[i] = ConvertAdminToUser(admin)
	}
	return users
}

func convertToUserSliceOwner(owners []*models.Owner) []*models.User {
	users := make([]*models.User, len(owners))
	for i, owner := range owners {
		users[i] = ConvertOwnerToUser(owner)
	}
	return users
}
func convertToUserSliceEmployee(employees []*models.Employee) []*models.User {
	users := make([]*models.User, len(employees))
	for i, employee := range employees {
		users[i] = ConvertEmployeeToUser(employee)
	}
	return users
}
func convertToUserSliceMember(members []*models.Member) []*models.User {
	users := make([]*models.User, len(members))
	for i, member := range members {
		users[i] = ConvertMemberToUser(member)
	}
	return users
}
