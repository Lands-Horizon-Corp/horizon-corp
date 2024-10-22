package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type MediaRepository struct {
	DB *gorm.DB
}

func NewMediaRepository(db *gorm.DB) *MediaRepository {
	return &MediaRepository{DB: db}
}

func (r *MediaRepository) Create(media *models.Media) error {
	err := r.DB.Create(media).Error
	return handleDBError(err)
}

func (r *MediaRepository) GetAll() ([]models.Media, error) {
	var medias []models.Media
	err := r.DB.Find(&medias).Error
	return medias, handleDBError(err)
}

func (r *MediaRepository) GetByID(id uint) (models.Media, error) {
	var media models.Media
	err := r.DB.First(&media, id).Error
	return media, handleDBError(err)
}

func (r *MediaRepository) Update(id uint, media *models.Media) error {
	media.ID = id
	err := r.DB.Save(media).Error
	return handleDBError(err)
}

func (r *MediaRepository) Delete(id uint) error {
	err := r.DB.Delete(&models.Media{}, id).Error
	return handleDBError(err)
}
