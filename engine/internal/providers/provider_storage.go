package providers

import (
	"fmt"
	"mime/multipart"
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
	logger  *zap.Logger
	helpers *helpers.HelpersFunction
}

func NewStorageProvider(
	cfg *config.AppConfig,
	logger *zap.Logger,
	helpers *helpers.HelpersFunction,
) (*StorageProvider, error) {
	sess, err := session.NewSession(&aws.Config{
		Endpoint:         aws.String(cfg.StorageEndpoint),
		Region:           aws.String(cfg.StorageRegion),
		Credentials:      credentials.NewStaticCredentials(cfg.StorageAccessKey, cfg.StorageSecretKey, ""),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %v", err)
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
	_, err := fc.Client.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(fc.cfg.StorageBucketName),
	})
	if err != nil {
		_, err = fc.Client.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(fc.cfg.StorageBucketName),
		})
		if err != nil {
			return fmt.Errorf("failed to create bucket %s: %v", fc.cfg.StorageBucketName, err)
		}
		err = fc.Client.WaitUntilBucketExists(&s3.HeadBucketInput{
			Bucket: aws.String(fc.cfg.StorageBucketName),
		})
		if err != nil {
			return fmt.Errorf("failed to wait for bucket %s to be created: %v", fc.cfg.StorageBucketName, err)
		}
	}
	return nil
}

func (fc *StorageProvider) UploadFile(fileHeader *multipart.FileHeader) (*Media, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	key := fc.helpers.UniqueFileName(fileHeader.Filename)

	if err := fc.CreateBucketIfNotExists(); err != nil {
		return nil, err
	}

	uploader := s3manager.NewUploaderWithClient(fc.Client)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(fc.cfg.StorageBucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to upload file: %v", err)
	}

	media := &Media{
		FileName:   fileHeader.Filename,
		FileSize:   fileHeader.Size,
		FileType:   fileHeader.Header.Get("Content-Type"),
		StorageKey: key,
		URL:        result.Location,
		BucketName: fc.cfg.StorageBucketName,
	}

	return media, nil
}

func (fc *StorageProvider) DeleteFile(key string) error {
	_, err := fc.Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(fc.cfg.StorageBucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}

	return fc.Client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(fc.cfg.StorageBucketName),
		Key:    aws.String(key),
	})
}

func (fc *StorageProvider) GeneratePresignedURL(key string) (string, error) {
	expiration := 20 * time.Minute
	req, _ := fc.Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(fc.cfg.StorageBucketName),
		Key:    aws.String(key),
	})
	urlStr, err := req.Presign(expiration)
	if err != nil {
		return "", err
	}
	return urlStr, nil
}
