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

type MinioClient struct {
	Client s3iface.S3API
}

func NewMinioClient(cfg *config.Config) *MinioClient {
	sess, err := session.NewSession(&aws.Config{
		Endpoint:         aws.String(cfg.Storage.Endpoint),
		Region:           aws.String(cfg.Storage.Region),
		Credentials:      credentials.NewStaticCredentials(cfg.Storage.AccessKey, cfg.Storage.SecretKey, ""),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create session: %v", err))
	}

	return &MinioClient{
		Client: s3.New(sess),
	}
}

func (mc *MinioClient) UploadFile(bucketName, key string, body io.Reader) error {
	uploader := s3manager.NewUploaderWithClient(mc.Client)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   body,
	})
	return err
}

func (mc *MinioClient) DeleteFile(bucketName, key string) error {
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

func (mc *MinioClient) GeneratePresignedURL(bucketName, key string, expiration time.Duration) (string, error) {
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
