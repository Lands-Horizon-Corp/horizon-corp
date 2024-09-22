// package service

package service

import (
	"horizon-core/internal/events"
	"horizon-core/internal/models"
	"horizon-core/internal/repository"
	"horizon-core/internal/resources"

	"gorm.io/gorm"
)

type MediaService struct {
	*repository.ModelRepository[models.Media]
}

func NewMediaService(db *gorm.DB) *MediaService {
	return &MediaService{
		ModelRepository: repository.NewModelRepository[models.Media](db),
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

func (r *MediaService) CreateMedia(mediaInput models.Media) (resources.MediaResource, error) {
	if err := r.Create(&mediaInput); err != nil {
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

	dispatcher := events.GetDispatcher()
	dispatcher.Dispatch(events.Event{
		Type:    events.EmployeeDeleted,
		Payload: media,
	})

	return r.Delete(&media)
}
