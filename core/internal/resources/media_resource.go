package resources

import (
	"horizon-core/internal/models"
	"time"
)

type MediaResource struct {
	ID           string    `json:"id"`
	URL          string    `json:"url"`
	FileName     string    `json:"file_name"`
	FileType     string    `json:"file_type"`
	FileSize     int       `json:"file_size"`
	UploadTime   time.Time `json:"upload_time"`
	Description  string    `json:"description"`
	BucketName   string    `json:"bucket_name"`
	TemporaryURL string    `json:"temporary_url"`
}

func NewMediaResource(media *models.Media) *MediaResource {
	if media == nil {
		return nil
	}

	return &MediaResource{
		ID:           media.ID,
		URL:          media.URL,
		FileName:     media.FileName,
		FileType:     media.FileType,
		FileSize:     media.FileSize,
		UploadTime:   media.UploadTime,
		Description:  media.Description,
		BucketName:   media.BucketName,
		TemporaryURL: media.TemporaryURL,
	}
}
