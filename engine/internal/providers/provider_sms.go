package providers

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"go.uber.org/zap"
)

type SMS interface {
	FormatSMS(req SMSRequest) (string, error)
	SendSMS(req SMSRequest) error
}

type SMSRequest struct {
	To   string             `json:"to"`
	Body string             `json:"body"`
	Vars *map[string]string `json:"vars"`
}

type snsClient interface {
	Publish(input *sns.PublishInput) (*sns.PublishOutput, error)
}

type SMSService struct {
	client  snsClient
	cfg     *config.AppConfig
	logger  Logger
	helpers *helpers.HelpersFunction
}

func NewSMSProvider(
	cfg *config.AppConfig,
	logger Logger,
	helpers *helpers.HelpersFunction,
) (SMS, error) {
	if cfg == nil {
		return nil, errors.New("configuration cannot be nil")
	}
	if cfg.AWSRegion == "" || cfg.AWSAccessKeyID == "" || cfg.AWSSecretAccessKey == "" {
		logger.Error("Incomplete AWS configuration for SMS service")
		return nil, errors.New("incomplete AWS configuration")
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(cfg.AWSRegion),
		Credentials: credentials.NewStaticCredentials(cfg.AWSAccessKeyID, cfg.AWSSecretAccessKey, ""),
	})
	if err != nil {
		logger.Error("Failed to create AWS session", zap.Error(err))
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	svc := sns.New(sess)
	logger.Info("AWS SNS session initialized successfully for SMS")
	return &SMSService{
		client: svc,
		logger: logger,
		cfg:    cfg,
	}, nil
}

func (s *SMSService) FormatSMS(req SMSRequest) (string, error) {
	tmpl, err := template.New("sms").Parse(req.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse SMS template: %w", err)
	}

	var data interface{}
	if req.Vars != nil {
		data = req.Vars
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute SMS template: %w", err)
	}

	return buf.String(), nil
}

func (s *SMSService) SendSMS(req SMSRequest) error {
	if req.To == "" {
		err := errors.New("recipient phone number is required")
		s.logger.Error("Missing recipient phone number", zap.Error(err))
		return err
	}

	if len(req.To) < 10 {
		err := fmt.Errorf("invalid phone number: %s", req.To)
		s.logger.Error("Invalid recipient phone number", zap.Error(err))
		return err
	}

	if req.Body == "" {
		err := errors.New("message body is required")
		s.logger.Error("Missing SMS body", zap.Error(err))
		return err
	}

	formattedMessage, err := s.FormatSMS(req)
	if err != nil {
		s.logger.Error("Failed to format SMS message", zap.Error(err))
		return err
	}

	params := &sns.PublishInput{
		Message:     aws.String(formattedMessage),
		PhoneNumber: aws.String(req.To),
	}

	s.logger.Info("Sending SMS via SNS",
		zap.String("to", req.To),
		zap.String("messagePreview", s.helpers.PreviewString(formattedMessage, 60)),
	)

	_, err = s.client.Publish(params)
	if err != nil {
		s.logger.Error("Failed to send SMS", zap.Error(err), zap.String("to", req.To))
		return fmt.Errorf("failed to send SMS to %s: %w", req.To, err)
	}

	s.logger.Info("SMS sent successfully",
		zap.String("to", req.To),
		zap.String("messagePreview", s.helpers.PreviewString(formattedMessage, 60)))
	return nil
}
