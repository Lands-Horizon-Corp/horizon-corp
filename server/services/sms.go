package services

import (
	"bytes"
	"fmt"
	"horizon/server/config"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"go.uber.org/zap"
)

type SMSRequest struct {
	To   string             `json:"to"`
	Body string             `json:"body"`
	Vars *map[string]string `json:"vars"`
}

type SMSService struct {
	logger *zap.Logger
	cfg    *config.AppConfig
	snsSvc *sns.SNS
}

func NewSMSService(logger *zap.Logger, cfg *config.AppConfig) (*SMSService, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(cfg.AWSRegion),
		Credentials: credentials.NewStaticCredentials(cfg.AWSAccessKeyID, cfg.AWSSecretAccessKey, ""),
	})
	if err != nil {
		logger.Error("Failed to create AWS session", zap.Error(err))
		return nil, err
	}

	svc := sns.New(sess)

	return &SMSService{
		logger: logger,
		cfg:    cfg,
		snsSvc: svc,
	}, nil
}

func (s *SMSService) FormatSMS(req SMSRequest) (string, error) {
	var buf bytes.Buffer
	formattedMessage := req.Body

	if req.Vars != nil {
		for key, value := range *req.Vars {
			placeholder := fmt.Sprintf("{{.%s}}", key)
			formattedMessage = strings.ReplaceAll(formattedMessage, placeholder, value)
		}
	}
	buf.WriteString(formattedMessage)

	return buf.String(), nil
}

func (s *SMSService) SendSMS(req SMSRequest) error {
	formattedMessage, err := s.FormatSMS(req)
	if err != nil {
		s.logger.Error("Failed to format SMS message", zap.Error(err))
		return err
	}

	params := &sns.PublishInput{
		Message:     aws.String(formattedMessage),
		PhoneNumber: aws.String(req.To),
	}

	_, err = s.snsSvc.Publish(params)
	if err != nil {
		s.logger.Error("Failed to send SMS", zap.Error(err), zap.String("to", req.To))
		return err
	}

	s.logger.Info("SMS sent successfully", zap.String("to", req.To), zap.String("message", formattedMessage))
	return nil
}
