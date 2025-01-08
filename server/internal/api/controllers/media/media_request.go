package media

type MediaRequest struct {
	ID *uint `json:"id"`

	FileName   string `json:"fileName" validate:"required,max=255"`
	FileSize   int64  `json:"fileSize" validate:"required,min=1"`
	FileType   string `json:"fileType" validate:"required,max=50"`
	StorageKey string `json:"storageKey" validate:"required,max=255"`
	URL        string `json:"url" validate:"required,url,max=255"`
	Key        string `json:"key,omitempty" validate:"max=255"`
	BucketName string `json:"bucketName,omitempty" validate:"max=255"`
}
