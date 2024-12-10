package managers

import (
	"errors"
	"fmt"

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
	if err := r.db.Client.Create(entity).Error; err != nil {
		return fmt.Errorf("failed to create entity: %w", err)
	}
	return nil
}

func (r *Repository[T]) FindByID(id uint, preloads ...string) (*T, error) {
	var entity T
	query := r.db.Client
	query = r.applyPreloads(query, preloads)

	result := query.First(&entity, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("entity with ID %d not found: %w", id, result.Error)
	}
	return &entity, result.Error
}

func (r *Repository[T]) FindAll(preloads ...string) ([]T, error) {
	var entities []T
	query := r.db.Client
	query = r.applyPreloads(query, preloads)

	result := query.Find(&entities)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve entities: %w", result.Error)
	}
	return entities, nil
}

func (r *Repository[T]) Update(entity *T) error {
	if err := r.db.Client.Save(entity).Error; err != nil {
		return fmt.Errorf("failed to update entity: %w", err)
	}
	return nil
}

func (r *Repository[T]) Delete(id uint) error {
	if err := r.db.Client.Delete(new(T), id).Error; err != nil {
		return fmt.Errorf("failed to delete entity with ID %d: %w", id, err)
	}
	return nil
}

// applyPreloads accepts a slice of preloads ([]string) for multiple preloads
func (r *Repository[T]) applyPreloads(query *gorm.DB, preloads []string) *gorm.DB {
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	return query
}
