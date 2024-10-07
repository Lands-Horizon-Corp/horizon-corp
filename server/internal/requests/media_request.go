package requests

import "github.com/go-playground/validator"

// MediaRequest represents the request payload for creating or updating media.
type MediaRequest struct {
	FileName   string `json:"file_name" validate:"required,max=255"`
	FileSize   int64  `json:"file_size" validate:"required,min=1"`
	FileType   string `json:"file_type" validate:"required,max=50"`
	StorageKey string `json:"storage_key" validate:"required,max=255"`
	Key        string `json:"key" validate:"omitempty,max=255"`
	URL        string `json:"url" validate:"omitempty,max=255"`
	BucketName string `json:"bucket_name" validate:"omitempty,max=255"`
}

// Validate validates the MediaRequest fields.
func (r *MediaRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
