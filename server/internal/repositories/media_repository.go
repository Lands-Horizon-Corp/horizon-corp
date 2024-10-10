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
	return r.DB.Create(media).Error
}

func (r *MediaRepository) GetAll() ([]models.Media, error) {
	var medias []models.Media
	err := r.DB.Find(&medias).Error
	return medias, err
}

func (r *MediaRepository) GetByID(id uint) (models.Media, error) {
	var media models.Media
	err := r.DB.First(&media, id).Error
	return media, err
}

func (r *MediaRepository) Update(id uint, media *models.Media) error {
	media.ID = id
	return r.DB.Save(media).Error
}

func (r *MediaRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Media{}, id).Error
}
