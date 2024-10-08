package storage

import (
	"fmt"
	"horizon/server/config"
	"horizon/server/internal/models"
	"mime/multipart"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	fileClientInstance *FileClient
	once               sync.Once
)

// FileClient wraps the AWS S3 client.
type FileClient struct {
	Client s3iface.S3API
	cfg    *config.AppConfig
}

// initializeFileClient sets up the singleton instance of FileClient.
func initializeFileClient() *FileClient {
	cfg, _ := config.LoadConfig()
	sess, err := session.NewSession(&aws.Config{
		Endpoint:         aws.String(cfg.StorageEndpoint),
		Region:           aws.String(cfg.StorageRegion),
		Credentials:      credentials.NewStaticCredentials(cfg.StorageAccessKey, cfg.StorageSecretKey, ""),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create session: %v", err))
	}

	return &FileClient{
		Client: s3.New(sess),
		cfg:    cfg,
	}
}

// GetFileClient returns the singleton instance of FileClient.
func GetFileClient() *FileClient {
	once.Do(func() {
		fileClientInstance = initializeFileClient()
	})
	return fileClientInstance
}

// CreateBucketIfNotExists checks if the bucket exists and creates it if not.
func CreateBucketIfNotExists() error {
	mc := GetFileClient()

	_, err := mc.Client.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(mc.cfg.StorageBucketName),
	})
	if err != nil {
		_, err = mc.Client.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(mc.cfg.StorageBucketName),
		})
		if err != nil {
			return fmt.Errorf("failed to create bucket %s: %v", mc.cfg.StorageBucketName, err)
		}
		err = mc.Client.WaitUntilBucketExists(&s3.HeadBucketInput{
			Bucket: aws.String(mc.cfg.StorageBucketName),
		})
		if err != nil {
			return fmt.Errorf("failed to wait for bucket %s to be created: %v", mc.cfg.StorageBucketName, err)
		}
	}
	return nil
}

func UploadFile(fileHeader *multipart.FileHeader) (*models.Media, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	key := UniqueFileName(fileHeader.Filename)

	mc := GetFileClient()
	if err := CreateBucketIfNotExists(); err != nil {
		return nil, err
	}

	uploader := s3manager.NewUploaderWithClient(mc.Client)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(mc.cfg.StorageBucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to upload file: %v", err)
	}

	media := &models.Media{
		FileName:   fileHeader.Filename,
		FileSize:   fileHeader.Size,
		FileType:   fileHeader.Header.Get("Content-Type"),
		StorageKey: key,
		URL:        result.Location,
		BucketName: mc.cfg.StorageBucketName,
	}

	return media, nil
}

func DeleteFile(key string) error {
	mc := GetFileClient()
	_, err := mc.Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(mc.cfg.StorageBucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}

	return GetFileClient().Client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(mc.cfg.StorageBucketName),
		Key:    aws.String(key),
	})
}

func GeneratePresignedURL(key string) (string, error) {
	expiration := 20 * time.Minute
	mc := GetFileClient()
	req, _ := GetFileClient().Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(mc.cfg.StorageBucketName),
		Key:    aws.String(key),
	})
	urlStr, err := req.Presign(expiration)
	if err != nil {
		return "", err
	}
	return urlStr, nil
}
