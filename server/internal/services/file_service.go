package services

import (
	"fmt"
	"horizon-server/config"
	"horizon-server/internal/models"
	"horizon-server/internal/repositories"
	"horizon-server/pkg/storage"
	"io"
	"net/url"
	"time"
)

type FileService interface {
	UploadFile(file *models.File, body io.Reader) error
	GetPublicURL(file *models.File) string
	DeleteFile(id uint) error
	GetFileByID(id uint) (*models.File, error)
	DownloadFile(id uint) (string, error)
}

type fileService struct {
	repo       repositories.FileRepository
	fileClient *storage.FileClient
	config     *config.Config
}

func NewFileService(repo repositories.FileRepository, fileClient *storage.FileClient, config *config.Config) FileService {
	return &fileService{repo: repo, fileClient: fileClient, config: config}
}

func (f *fileService) UploadFile(file *models.File, body io.Reader) error {
	if err := f.fileClient.UploadFile(f.config.Storage.BucketName, file.FileName, body); err != nil {
		return err
	}
	return f.repo.Create(file)
}

func (f *fileService) DeleteFile(id uint) error {
	tx := f.repo.DB().Begin()
	if tx.Error != nil {
		return tx.Error
	}

	file, err := f.repo.GetByIDTx(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = f.fileClient.DeleteFile(f.config.Storage.BucketName, file.FileName)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = f.repo.DeleteTx(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (f *fileService) GetPublicURL(file *models.File) string {
	return fmt.Sprintf("%s/%s/%s", f.config.Storage.BucketName, f.config.Storage.Endpoint, url.PathEscape(file.FileName))
}

func (s *fileService) GetFileByID(id uint) (*models.File, error) {
	return s.repo.GetByID(id)
}

func (f *fileService) DownloadFile(id uint) (string, error) {
	file, err := f.repo.GetByID(id)
	if err != nil {
		return "", err
	}
	expiration := time.Minute * 15 // Set expiration time of the URL
	url, err := f.fileClient.GeneratePresignedURL(file.BucketName, file.FileName, expiration)
	if err != nil {
		return "", err
	}
	return url, nil
}
