package repositories

import (
	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func NewRepository[T any](db *gorm.DB) *Repository[T] {
	return &Repository[T]{DB: db}
}

func (r *Repository[T]) Create(entity *T) error {
	err := r.DB.Create(entity).Error
	return handleDBError(err)
}

func (r *Repository[T]) GetAll(preloads ...string) ([]*T, error) {
	var entities []*T
	db := r.DB
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	err := db.Find(&entities).Error
	return entities, handleDBError(err)
}

func (r *Repository[T]) GetByID(id uint, preloads ...string) (*T, error) {
	entity := new(T)
	db := r.DB
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	err := db.First(entity, id).Error
	return entity, handleDBError(err)
}

func (r *Repository[T]) Update(entity *T) error {
	err := r.DB.Save(entity).Error
	return handleDBError(err)
}

func (r *Repository[T]) Delete(id uint) error {
	entity := new(T)
	err := r.DB.Delete(entity, id).Error
	return handleDBError(err)
}

func (r *Repository[T]) UpdateColumns(id uint, updates T) (*T, error) {
	entity := new(T)
	if err := r.DB.Model(entity).Where("id = ?", id).Updates(updates).Error; err != nil {
		return nil, handleDBError(err)
	}
	if err := r.DB.First(entity, id).Error; err != nil {
		return nil, handleDBError(err)
	}
	return entity, nil
}
