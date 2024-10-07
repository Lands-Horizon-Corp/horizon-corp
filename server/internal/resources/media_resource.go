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
	Key         string `json:"key"`
	BucketName  string `json:"bucket_name"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DownloadURL string `json:"download_url"`
}

func ToResourceMedia(media models.Media) (MediaResource, error) {
	expiration := 20 * time.Minute

	temporaryURL, err := storage.GeneratePresignedURL(media.BucketName, media.StorageKey, expiration)
	if err != nil {
		return MediaResource{}, err
	}

	return MediaResource{
		ID:          media.ID,
		FileName:    media.FileName,
		FileSize:    media.FileSize,
		FileType:    media.FileType,
		StorageKey:  media.StorageKey,
		Key:         media.Key,
		BucketName:  media.BucketName,
		CreatedAt:   media.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   media.UpdatedAt.Format(time.RFC3339),
		DownloadURL: temporaryURL,
	}, nil
}

func ToResourceListMedia(mediaList []models.Media) ([]MediaResource, error) {
	var resources []MediaResource
	for _, media := range mediaList {
		resource, err := ToResourceMedia(media)
		if err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}
	return resources, nil
}
