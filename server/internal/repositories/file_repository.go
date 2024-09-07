package repositories

import (
	"horizon-server/internal/models"

	"github.com/jinzhu/gorm"
)

type FileRepository interface {
	Create(file *models.File) error
	Delete(id uint) error
	GetByID(id uint) (*models.File, error)
	DeleteTx(tx *gorm.DB, id uint) error
	GetByIDTx(tx *gorm.DB, id uint) (*models.File, error)
	DB() *gorm.DB
}

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return &fileRepository{db: db}
}

func (r *fileRepository) Create(file *models.File) error {
	return r.db.Create(file).Error
}

func (r *fileRepository) Delete(id uint) error {
	var file models.File
	if err := r.db.First(&file, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&file).Error
}

func (r *fileRepository) GetByID(id uint) (*models.File, error) {
	var file models.File
	if err := r.db.First(&file, id).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func (r *fileRepository) DeleteTx(tx *gorm.DB, id uint) error {
	var file models.File
	if err := tx.First(&file, id).Error; err != nil {
		return err
	}
	return tx.Delete(&file).Error
}

func (r *fileRepository) GetByIDTx(tx *gorm.DB, id uint) (*models.File, error) {
	var file models.File
	if err := tx.First(&file, id).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func (r *fileRepository) DB() *gorm.DB {
	return r.db
}
