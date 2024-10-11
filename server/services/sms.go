package services

import (
	"context"

	"horizon/server/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// SMSService defines the interface for sending SMS messages.
type SMSService interface {
	SendSMS(ctx context.Context, to, message string) error
}

// smsServiceImpl is the concrete implementation of SMSService.
type smsServiceImpl struct {
	logger *zap.Logger
	cfg    *config.AppConfig
	snsSvc *sns.SNS
}

// NewSMSService constructs a new SMSService.
func NewSMSService(logger *zap.Logger, cfg *config.AppConfig) (SMSService, error) {
	// Create AWS session with custom configuration
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(cfg.AWSRegion),
		Credentials: credentials.NewStaticCredentials(cfg.AWSAccessKeyID, cfg.AWSSecretAccessKey, ""),
	})
	if err != nil {
		logger.Error("Failed to create AWS session", zap.Error(err))
		return nil, err
	}

	svc := sns.New(sess)

	return &smsServiceImpl{
		logger: logger,
		cfg:    cfg,
		snsSvc: svc,
	}, nil
}

func (s *smsServiceImpl) SendSMS(ctx context.Context, to, message string) error {
	params := &sns.PublishInput{
		Message:     aws.String(message),
		PhoneNumber: aws.String(to),
	}

	_, err := s.snsSvc.Publish(params)
	if err != nil {
		s.logger.Error("Failed to send SMS", zap.Error(err), zap.String("to", to))
		return err
	}

	s.logger.Info("SMS sent successfully", zap.String("to", to))
	return nil
}

func SMSRun(lc fx.Lifecycle, smsService SMSService, logger *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := smsService.SendSMS(ctx, "+1234567890", "Hello from AWS SNS via UberFx!")
			if err != nil {
				logger.Error("Error sending test SMS", zap.Error(err))
			} else {
				logger.Info("Test SMS sent successfully")
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
