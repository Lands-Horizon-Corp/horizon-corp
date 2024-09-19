package services

import (
	"horizon-server/internal/models"
	"horizon-server/internal/repositories"
	"horizon-server/pkg/helpers"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *models.User) error
	Login(username, password string) (*models.User, error)
	ChangePassword(userID uint, oldPassword, newPassword string) error
	ForgotPassword(email string) error // Implement a reset logic via email
	UpdateUser(user *models.User) error
	AddAttachment(userID, fileID uint) error
	GetUserByID(id uint) (*models.User, error)
	ListUsers(filters []helpers.Filter, pagination *helpers.Pagination) ([]models.User, error)
}
type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repo.Create(user)
}

func (s *userService) Login(username, password string) (*models.User, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) ChangePassword(userID uint, oldPassword, newPassword string) error {
	user, err := s.repo.GetByID(userID)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return err
	}
	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedNewPassword)
	return s.repo.Update(user)
}

func (s *userService) ForgotPassword(email string) error {
	return nil
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.Update(user)
}

func (s *userService) AddAttachment(userID, fileID uint) error {
	return nil
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) ListUsers(filters []helpers.Filter, pagination *helpers.Pagination) ([]models.User, error) {
	return s.repo.Query(filters, pagination)
}
