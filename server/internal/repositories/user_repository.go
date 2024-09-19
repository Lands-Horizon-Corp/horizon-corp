package repositories

import (
	"horizon-server/internal/models"
	"horizon-server/pkg/helpers"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id uint) error
	GetByID(id uint) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Query(filters []helpers.Filter, pagination *helpers.Pagination) ([]models.User, error)

	DB() *gorm.DB
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&user).Error
}

func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail finds a user by their email.
func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) DB() *gorm.DB {
	return r.db
}

func (r *userRepository) Query(filters []helpers.Filter, pagination *helpers.Pagination) ([]models.User, error) {
	var users []models.User
	db := r.db.Model(&models.User{})
	err := helpers.Paginate(db, &users, filters, pagination)
	if err != nil {
		return nil, err
	}
	return users, nil
}
