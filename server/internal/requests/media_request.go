package requests

import (
	"github.com/go-playground/validator"
)

type MediaRequest struct {
	ID         uint   `json:"id,omitempty"`
	FileName   string `json:"fileName" validate:"required,max=255"`
	FileSize   int64  `json:"fileSize" validate:"required,min=1"`
	FileType   string `json:"fileType" validate:"required,max=50"`
	StorageKey string `json:"storageKey" validate:"required,max=255"`
	URL        string `json:"url" validate:"omitempty,max=255"`
	BucketName string `json:"bucketName" validate:"omitempty,max=255"`
}

func (r *MediaRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return FormatValidationError(err)
	}
	return nil
}
