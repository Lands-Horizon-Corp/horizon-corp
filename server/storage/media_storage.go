package storage

import (
	"fmt"
	"horizon/server/config"
	"io"
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
func CreateBucketIfNotExists(bucketName string) error {
	mc := GetFileClient()

	_, err := mc.Client.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		_, err = mc.Client.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
		})
		if err != nil {
			return fmt.Errorf("failed to create bucket %s: %v", bucketName, err)
		}
		err = mc.Client.WaitUntilBucketExists(&s3.HeadBucketInput{
			Bucket: aws.String(bucketName),
		})
		if err != nil {
			return fmt.Errorf("failed to wait for bucket %s to be created: %v", bucketName, err)
		}
	}
	return nil
}

// UploadFile uploads a file to the specified bucket and key.
func UploadFile(bucketName, key string, body io.Reader) error {
	if err := CreateBucketIfNotExists(bucketName); err != nil {
		return err
	}
	uploader := s3manager.NewUploaderWithClient(GetFileClient().Client)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   body,
	})
	return err
}

// DeleteFile deletes a file from the specified bucket and key.
func DeleteFile(bucketName, key string) error {
	_, err := GetFileClient().Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}

	return GetFileClient().Client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
}

// GeneratePresignedURL generates a presigned URL for the specified bucket, key, and expiration.
func GeneratePresignedURL(bucketName, key string, expiration time.Duration) (string, error) {
	req, _ := GetFileClient().Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	urlStr, err := req.Presign(expiration)
	if err != nil {
		return "", err
	}
	return urlStr, nil
}
