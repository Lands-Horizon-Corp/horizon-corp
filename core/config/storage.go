package config

import (
	"fmt"

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

// MediaClient wraps the S3 client interface
type MediaClient struct {
	Client s3iface.S3API
}

var (
	mediaClientInstance *MediaClient
	mediaClientOnce     sync.Once
	mediaClientInitErr  error
)

// InitMediaClient initializes the singleton MediaClient instance.
// It should be called once during application startup.
func InitMediaClient() error {
	cfg := GetConfig()
	mediaClientOnce.Do(func() {
		sess, err := session.NewSession(&aws.Config{
			Endpoint:         aws.String(cfg.Storage.Endpoint),
			Region:           aws.String(cfg.Storage.Region),
			Credentials:      credentials.NewStaticCredentials(cfg.Storage.AccessKey, cfg.Storage.SecretKey, ""),
			S3ForcePathStyle: aws.Bool(true),
		})
		if err != nil {
			mediaClientInitErr = fmt.Errorf("failed to create session: %v", err)
			return
		}

		mediaClientInstance = &MediaClient{
			Client: s3.New(sess),
		}
	})

	return mediaClientInitErr
}

// GetMediaClient returns the singleton instance of MediaClient.
// It returns nil if the client has not been initialized.
func GetMediaClient() *MediaClient {
	return mediaClientInstance
}

// CreateBucketIfNotExists checks if the bucket exists and creates it if not.
func (mc *MediaClient) CreateBucketIfNotExists(bucketName string) error {
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
func (mc *MediaClient) UploadFile(bucketName, key string, body io.Reader) error {
	if err := mc.CreateBucketIfNotExists(bucketName); err != nil {
		return err
	}
	uploader := s3manager.NewUploaderWithClient(mc.Client)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   body,
	})
	return err
}

// DeleteFile deletes a file from the specified bucket and key.
func (mc *MediaClient) DeleteFile(bucketName, key string) error {
	_, err := mc.Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}

	return mc.Client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
}

// GeneratePresignedURL generates a presigned URL for accessing the file.
func (mc *MediaClient) GeneratePresignedURL(bucketName, key string, expiration time.Duration) (string, error) {
	req, _ := mc.Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	urlStr, err := req.Presign(expiration)
	if err != nil {
		return "", err
	}
	return urlStr, nil
}
