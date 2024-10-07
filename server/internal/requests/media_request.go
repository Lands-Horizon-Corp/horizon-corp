package requests

import "github.com/go-playground/validator"

// MediaRequest represents the request payload for creating or updating media.
type MediaRequest struct {
	FileName   string `json:"fileName" validate:"required,max=255"`
	FileSize   int64  `json:"fileSize" validate:"required,min=1"`
	FileType   string `json:"fileType" validate:"required,max=50"`
	StorageKey string `json:"storageKey" validate:"required,max=255"`
	Key        string `json:"key" validate:"omitempty,max=255"`
	URL        string `json:"url" validate:"omitempty,max=255"`
	BucketName string `json:"bucketName" validate:"omitempty,max=255"`
}

// Validate validates the MediaRequest fields.
func (r *MediaRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
