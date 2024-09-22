// package service

package service

import (
	"fmt"
	"horizon-core/config"
	"horizon-core/helpers"
	"horizon-core/internal/events"
	"horizon-core/internal/models"
	"horizon-core/internal/repository"
	"horizon-core/internal/resources"
	"io"
	"mime/multipart"
	"net/url"
	"time"

	"gorm.io/gorm"
)

type MediaService struct {
	*repository.ModelRepository[models.Media]
	config  *config.Config
	storage *config.MediaClient
}

func NewMediaService(db *gorm.DB) *MediaService {
	storage := config.GetMediaClient()

	return &MediaService{
		ModelRepository: repository.NewModelRepository[models.Media](db),
		config:          config.GetConfig(),
		storage:         storage,
	}
}

func (r *MediaService) GetMediaByID(id string) (resources.MediaResource, error) {
	media, err := r.FindByID(id, nil)
	if err != nil {
		return resources.MediaResource{}, err
	}

	mediaResource := resources.NewMediaResource(&media)
	return *mediaResource, nil
}

func (r *MediaService) DownloadFile(id string) (string, error) {
	media, err := r.FindByID(id, nil)
	if err != nil {
		return "", err
	}
	expiration := time.Minute * 15
	url, err := r.storage.GeneratePresignedURL(media.BucketName, media.FileName, expiration)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (r *MediaService) ListMedia(req repository.ListRequest) (repository.ListResponse[resources.MediaResource], error) {
	listResponse, err := r.List(req, nil)
	if err != nil {
		return repository.ListResponse[resources.MediaResource]{}, err
	}

	mediaResources := make([]resources.MediaResource, len(listResponse.Data))
	for i, media := range listResponse.Data {
		mediaResource := resources.NewMediaResource(&media)
		mediaResources[i] = *mediaResource
	}

	response := repository.ListResponse[resources.MediaResource]{
		Data:       mediaResources,
		Pagination: listResponse.Pagination,
	}

	return response, nil
}

func (r *MediaService) CreateMedia(body io.Reader, header *multipart.FileHeader) (resources.MediaResource, error) {

	name := helpers.UniqueFileName(header.Filename)

	if err := r.storage.UploadFile(r.config.Storage.BucketName, name, body); err != nil {
		return resources.MediaResource{}, err
	}

	expiration := 15 * time.Minute
	tempURL, _ := r.storage.GeneratePresignedURL(r.config.Storage.BucketName, name, expiration)
	mediaInput := &models.Media{
		URL:          r.GetPublicURL(name),
		FileName:     name,
		FileType:     header.Header.Get("Content-Type"),
		FileSize:     header.Size,
		UploadTime:   time.Now(),
		Description:  "Media upload",
		BucketName:   r.config.Storage.BucketName,
		TemporaryURL: tempURL,
	}

	if err := r.Create(mediaInput); err != nil {
		return resources.MediaResource{}, err
	}
	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.MediaCreated,
		Payload: mediaInput,
	})
	return r.GetMediaByID(mediaInput.ID)
}

func (r *MediaService) UpdateMedia(mediaInput models.Media) (resources.MediaResource, error) {
	if err := r.Update(&mediaInput); err != nil {
		return resources.MediaResource{}, err
	}
	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.MediaUpdated,
		Payload: mediaInput,
	})
	return r.GetMediaByID(mediaInput.ID)
}

func (r *MediaService) DeleteMedia(id string) error {

	media, err := r.FindByID(id, nil)
	if err != nil {
		return err
	}

	if err := r.storage.DeleteFile(r.config.Storage.BucketName, media.FileName); err != nil {
		return err
	}

	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.EmployeeDeleted,
		Payload: media,
	})

	return r.Delete(&media)
}

func (f *MediaService) GetPublicURL(filName string) string {
	return fmt.Sprintf("%s/%s/%s", f.config.Storage.BucketName, f.config.Storage.Endpoint, url.PathEscape(filName))
}
