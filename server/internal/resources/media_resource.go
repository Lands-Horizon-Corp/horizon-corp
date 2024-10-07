package resources

import (
	"horizon/server/internal/models"
	"horizon/server/storage"
	"time"
)

type MediaResource struct {
	ID          uint   `json:"id"`
	FileName    string `json:"fileName"`
	FileSize    int64  `json:"fileSize"`
	FileType    string `json:"fileType"`
	StorageKey  string `json:"storageKey"`
	URL         string `json:"url"`
	BucketName  string `json:"bucketName"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	DownloadURL string `json:"downloadURL"`
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
