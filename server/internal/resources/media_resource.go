package resources

import (
	"horizon/server/internal/models"
	"horizon/server/storage"
	"time"
)

type MediaResource struct {
	ID          uint   `json:"id"`
	FileName    string `json:"file_name"`
	FileSize    int64  `json:"file_size"`
	FileType    string `json:"file_type"`
	StorageKey  string `json:"storage_key"`
	URL         string `json:"url"`
	BucketName  string `json:"bucket_name"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DownloadURL string `json:"download_url"`
}

func ToResourceMedia(media models.Media) MediaResource {
	temporaryURL, err := storage.GeneratePresignedURL(media.StorageKey)
	if err != nil {
		return MediaResource{}
	}

	return MediaResource{
		ID:          media.ID,
		FileName:    media.FileName,
		FileSize:    media.FileSize,
		FileType:    media.FileType,
		StorageKey:  media.StorageKey,
		URL:         media.URL,
		BucketName:  media.BucketName,
		CreatedAt:   media.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   media.UpdatedAt.Format(time.RFC3339),
		DownloadURL: temporaryURL,
	}
}

func ToResourceListMedia(mediaList []models.Media) []MediaResource {
	var resources []MediaResource
	for _, media := range mediaList {
		resources = append(resources, ToResourceMedia(media))
	}
	return resources
}
