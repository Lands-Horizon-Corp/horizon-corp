package services

import (
	"fmt"
	"horizon-server/pkg/storage"
	"io"
	"net/url"
	"time"
)

type FileService struct {
	MinioClient *storage.MinioClient
}

func NewFileService(minioClient *storage.MinioClient) *FileService {
	return &FileService{MinioClient: minioClient}
}

func (s *FileService) UploadFileProgress(
	bucketName, key string, body io.Reader, fileName string, fileSize int64,
	progressCallback func(fileName string, fileSize int64, progressBytes int64, progressPercentage float64)) error {

	pr, pw := io.Pipe()
	tee := io.TeeReader(body, pw)
	go func() {
		defer pw.Close()
		buffer := make([]byte, 4096)
		var totalRead int64
		for {
			n, err := tee.Read(buffer)
			if n > 0 {
				totalRead += int64(n)
				progressPercentage := float64(totalRead) / float64(fileSize) * 100
				progressCallback(fileName, fileSize, totalRead, progressPercentage)
			}
			if err != nil {
				break
			}
		}
	}()
	return s.MinioClient.UploadFile(bucketName, key, pr)
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

func (s *FileService) GetPublicURL(bucketName, key string) string {
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, url.PathEscape(key))
}
