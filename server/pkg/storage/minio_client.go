package storage

import (
	"fmt"
	"horizon-server/config"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type FileClient struct {
	Client s3iface.S3API
}

func NewFileClient(cfg *config.Config) *FileClient {
	sess, err := session.NewSession(&aws.Config{
		Endpoint:         aws.String(cfg.Storage.Endpoint),
		Region:           aws.String(cfg.Storage.Region),
		Credentials:      credentials.NewStaticCredentials(cfg.Storage.AccessKey, cfg.Storage.SecretKey, ""),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create session: %v", err))
	}

	return &FileClient{
		Client: s3.New(sess),
	}
}

// CreateBucketIfNotExists checks if the bucket exists and creates it if not.
func (mc *FileClient) CreateBucketIfNotExists(bucketName string) error {

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

func (mc *FileClient) UploadFile(bucketName, key string, body io.Reader) error {
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

func (mc *FileClient) DeleteFile(bucketName, key string) error {
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

func (mc *FileClient) GeneratePresignedURL(bucketName, key string, expiration time.Duration) (string, error) {
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
