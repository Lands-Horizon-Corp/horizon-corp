// repository/model_repository.go
package repository

import (
	"gorm.io/gorm"
)

type ModelRepository[T any] struct {
	db *gorm.DB
}

func NewModelRepository[T any](db *gorm.DB) *ModelRepository[T] {
	return &ModelRepository[T]{db: db}
}

func (r *ModelRepository[T]) preload(db *gorm.DB, eagerLoads []string) *gorm.DB {
	for _, relation := range eagerLoads {
		db = db.Preload(relation)
	}
	return db
}
func (r *ModelRepository[T]) FindAll(eagerLoads []string) ([]T, error) {
	var entities []T
	query := r.preload(r.db, eagerLoads)
	if err := query.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *ModelRepository[T]) FindByID(id string, eagerLoads []string) (T, error) {
	var entity T
	query := r.preload(r.db, eagerLoads)
	if err := query.First(&entity, "id = ?", id).Error; err != nil {
		return entity, err
	}
	return entity, nil
}

func (r *ModelRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *ModelRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *ModelRepository[T]) Delete(entity *T) error {
	return r.db.Delete(entity).Error
}
