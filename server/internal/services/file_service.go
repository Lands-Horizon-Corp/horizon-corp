package services

import (
	"horizon-server/pkg/storage"
	"io"
	"time"
)

type FileService struct {
	MinioClient *storage.MinioClient
}

func NewFileService(minioClient *storage.MinioClient) *FileService {
	return &FileService{MinioClient: minioClient}
}

func (s *FileService) UploadFile(bucketName, key string, body io.Reader) error {
	return s.MinioClient.UploadFile(bucketName, key, body)
}

func (s *FileService) DeleteFile(bucketName, key string) error {
	return s.MinioClient.DeleteFile(bucketName, key)
}

func (s *FileService) GeneratePresignedURL(bucketName, key string, expiration time.Duration) (string, error) {
	return s.MinioClient.GeneratePresignedURL(bucketName, key, expiration)
}
