package managers

import (
	"errors"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	db *providers.DatabaseService
}

func NewRepository[T any](db *providers.DatabaseService) *Repository[T] {
	return &Repository[T]{db: db}
}

func (r *Repository[T]) Create(entity *T) error {
	return r.db.Client.Create(entity).Error
}

func (r *Repository[T]) FindByID(id uint) (*T, error) {
	var entity T
	result := r.db.Client.First(&entity, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &entity, result.Error
}

func (r *Repository[T]) FindAll() ([]T, error) {
	var entities []T
	result := r.db.Client.Find(&entities)
	return entities, result.Error
}

func (r *Repository[T]) Update(entity *T) error {
	return r.db.Client.Save(entity).Error
}

func (r *Repository[T]) Delete(id uint) error {
	return r.db.Client.Delete(new(T), id).Error
}
