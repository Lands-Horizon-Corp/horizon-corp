package providers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"go.uber.org/zap"
)

type Media struct {
	FileName   string
	FileSize   int64
	FileType   string
	StorageKey string
	URL        string
	BucketName string
}

type StorageProvider struct {
	Client  s3iface.S3API
	cfg     *config.AppConfig
	logger  *LoggerService
	helpers *helpers.HelpersFunction
}

func NewStorageProvider(
	cfg *config.AppConfig,
	logger *LoggerService,
	helpers *helpers.HelpersFunction,
) (*StorageProvider, error) {

	logger.Debug("Initializing AWS S3 session",
		zap.String("endpoint", cfg.StorageEndpoint),
		zap.String("region", cfg.StorageRegion),
		zap.String("bucketName", cfg.StorageBucketName),
	)

	sess, err := session.NewSession(&aws.Config{
		Endpoint:         aws.String(cfg.StorageEndpoint),
		Region:           aws.String(cfg.StorageRegion),
		Credentials:      credentials.NewStaticCredentials(cfg.StorageAccessKey, cfg.StorageSecretKey, ""),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		logger.Error("Failed to create AWS session", zap.Error(err))
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	logger.Info("Successfully initialized AWS S3 session.")

	return &StorageProvider{
		Client:  s3.New(sess),
		cfg:     cfg,
		logger:  logger,
		helpers: helpers,
	}, nil
}

func (fc *StorageProvider) CreateBucketIfNotExists() error {
	fc.logger.Debug("Checking if S3 bucket exists",
		zap.String("bucketName", fc.cfg.StorageBucketName),
	)

	_, err := fc.Client.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(fc.cfg.StorageBucketName),
	})
	if err == nil {
		fc.logger.Debug("Bucket already exists",
			zap.String("bucketName", fc.cfg.StorageBucketName),
		)
		return nil
	}

	fc.logger.Info("Bucket does not exist, creating bucket",
		zap.String("bucketName", fc.cfg.StorageBucketName),
	)

	_, err = fc.Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(fc.cfg.StorageBucketName),
	})
	if err != nil {
		fc.logger.Error("Failed to create bucket",
			zap.String("bucketName", fc.cfg.StorageBucketName),
			zap.Error(err),
		)
		return fmt.Errorf("failed to create bucket %s: %w", fc.cfg.StorageBucketName, err)
	}

	if waitErr := fc.Client.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(fc.cfg.StorageBucketName),
	}); waitErr != nil {
		fc.logger.Error("Failed waiting for bucket creation",
			zap.String("bucketName", fc.cfg.StorageBucketName),
			zap.Error(waitErr),
		)
		return fmt.Errorf("failed waiting for bucket %s to be created: %w", fc.cfg.StorageBucketName, waitErr)
	}

	return nil
}

func (fc *StorageProvider) UploadToS3(bucketName, key string, body io.Reader) (*s3manager.UploadOutput, error) {
	fc.logger.Info("Uploading file to S3",
		zap.String("bucketName", bucketName),
		zap.String("key", key),
	)

	uploader := s3manager.NewUploaderWithClient(fc.Client)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   body,
	})
	if err != nil {
		fc.logger.Error("Unable to upload file to S3",
			zap.String("bucketName", bucketName),
			zap.String("key", key),
			zap.Error(err),
		)
		return nil, fmt.Errorf("unable to upload file to s3: %w", err)
	}

	return result, nil
}

func (fc *StorageProvider) UploadFile(fileHeader *multipart.FileHeader) (*Media, error) {

	file, err := fileHeader.Open()
	if err != nil {
		fc.logger.Error("Unable to open multipart file",
			zap.String("originalFileName", fileHeader.Filename),
			zap.Error(err),
		)
		return nil, fmt.Errorf("unable to open file: %w", err)
	}
	defer file.Close()

	key := fc.helpers.UniqueFileName(fileHeader.Filename)
	fc.logger.Debug("Generated unique key for file",
		zap.String("key", key),
	)

	if err = fc.CreateBucketIfNotExists(); err != nil {
		return nil, err
	}

	uploader := s3manager.NewUploaderWithClient(fc.Client)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(fc.cfg.StorageBucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		fc.logger.Error("Unable to upload multipart file to S3",
			zap.String("originalFileName", fileHeader.Filename),
			zap.String("key", key),
			zap.Error(err),
		)
		return nil, fmt.Errorf("unable to upload file: %w", err)
	}

	return &Media{
		FileName:   fileHeader.Filename,
		FileSize:   fileHeader.Size,
		FileType:   fileHeader.Header.Get("Content-Type"),
		StorageKey: key,
		URL:        result.Location,
		BucketName: fc.cfg.StorageBucketName,
	}, nil
}

// UploadLocalFile uploads a file from the local file system to S3.
func (fc *StorageProvider) UploadLocalFile(localFilePath string) (*Media, error) {
	fc.logger.Info("Uploading local file",
		zap.String("localFilePath", localFilePath),
	)

	file, err := os.Open(localFilePath)
	if err != nil {
		fc.logger.Error("Unable to open local file",
			zap.String("localFilePath", localFilePath),
			zap.Error(err),
		)
		return nil, fmt.Errorf("unable to open local file %s: %w", localFilePath, err)
	}
	defer file.Close()

	fileName := path.Base(localFilePath)
	key := fc.helpers.UniqueFileName(fileName)

	fc.logger.Debug("Generated unique key for local file",
		zap.String("localFilePath", localFilePath),
		zap.String("key", key),
	)

	if err = fc.CreateBucketIfNotExists(); err != nil {
		return nil, err
	}

	result, err := fc.UploadToS3(fc.cfg.StorageBucketName, key, file)
	if err != nil {
		return nil, err
	}

	fileInfo, err := os.Stat(localFilePath)
	if err != nil {
		fc.logger.Error("Unable to stat local file",
			zap.String("localFilePath", localFilePath),
			zap.Error(err),
		)
		return nil, fmt.Errorf("unable to stat local file %s: %w", localFilePath, err)
	}

	return &Media{
		FileName:   fileName,
		FileSize:   fileInfo.Size(),
		FileType:   "application/octet-stream",
		StorageKey: key,
		URL:        result.Location,
		BucketName: fc.cfg.StorageBucketName,
	}, nil
}

// UploadFromURL downloads a file from a remote URL and uploads it to S3.
func (fc *StorageProvider) UploadFromURL(fileURL string) (*Media, error) {
	fc.logger.Info("Uploading file from URL",
		zap.String("url", fileURL),
	)

	resp, err := http.Get(fileURL)
	if err != nil {
		fc.logger.Error("Unable to fetch file from URL",
			zap.String("url", fileURL),
			zap.Error(err),
		)
		return nil, fmt.Errorf("unable to fetch file from URL %s: %w", fileURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fc.logger.Error("File download failed",
			zap.String("url", fileURL),
			zap.Int("statusCode", resp.StatusCode),
		)
		return nil, fmt.Errorf("file download failed with status %d", resp.StatusCode)
	}

	fileName := path.Base(fileURL)
	if idx := strings.Index(fileName, "?"); idx != -1 {
		fileName = fileName[:idx]
	}
	key := fc.helpers.UniqueFileName(fileName)

	if err = fc.CreateBucketIfNotExists(); err != nil {
		return nil, err
	}

	result, err := fc.UploadToS3(fc.cfg.StorageBucketName, key, resp.Body)
	if err != nil {
		return nil, err
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	contentLength := resp.ContentLength
	if contentLength < 0 {
		contentLength = 0
	}

	return &Media{
		FileName:   fileName,
		FileSize:   contentLength,
		FileType:   contentType,
		StorageKey: key,
		URL:        result.Location,
		BucketName: fc.cfg.StorageBucketName,
	}, nil
}

// DeleteFile deletes a file from the S3 bucket.
func (fc *StorageProvider) DeleteFile(key string) error {
	fc.logger.Info("Deleting file from S3",
		zap.String("bucketName", fc.cfg.StorageBucketName),
		zap.String("key", key),
	)

	_, err := fc.Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(fc.cfg.StorageBucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		fc.logger.Error("Failed to delete file from bucket",
			zap.String("bucketName", fc.cfg.StorageBucketName),
			zap.String("key", key),
			zap.Error(err),
		)
		return fmt.Errorf("failed to delete file %s from bucket %s: %w", key, fc.cfg.StorageBucketName, err)
	}

	if waitErr := fc.Client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(fc.cfg.StorageBucketName),
		Key:    aws.String(key),
	}); waitErr != nil {
		fc.logger.Error("Failed waiting for file deletion",
			zap.String("bucketName", fc.cfg.StorageBucketName),
			zap.String("key", key),
			zap.Error(waitErr),
		)
		return fmt.Errorf("failed waiting for file %s deletion: %w", key, waitErr)
	}

	return nil
}

// GeneratePresignedURL generates a presigned URL for a file in S3.
func (fc *StorageProvider) GeneratePresignedURL(key string) (string, error) {
	fc.logger.Info("Generating presigned URL",
		zap.String("bucketName", fc.cfg.StorageBucketName),
		zap.String("key", key),
	)

	expiration := 20 * time.Minute
	req, _ := fc.Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(fc.cfg.StorageBucketName),
		Key:    aws.String(key),
	})

	urlStr, err := req.Presign(expiration)
	if err != nil {
		fc.logger.Error("Unable to generate presigned URL",
			zap.String("bucketName", fc.cfg.StorageBucketName),
			zap.String("key", key),
			zap.Error(err),
		)
		return "", fmt.Errorf("unable to generate pre-signed URL for key %s: %w", key, err)
	}

	return urlStr, nil
}
