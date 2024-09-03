package storage

import (
	"horizon-server/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type MinioClient struct {
	Client *s3.S3
}

func NewMinioClient(cfg *config.Config) *MinioClient {
	sess, _ := session.NewSession(&aws.Config{
		Endpoint:    aws.String(cfg.Storage.Endpoint),
		Region:      aws.String(cfg.Storage.Region),
		Credentials: credentials.NewStaticCredentials(cfg.Storage.AccessKey, cfg.Storage.SecretKey, ""),
	})

	return &MinioClient{
		Client: s3.New(sess),
	}
}
