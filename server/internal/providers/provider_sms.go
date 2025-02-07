package providers

import (
	"bytes"
	"text/template"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/rotisserie/eris"
	"go.uber.org/zap"
)

type SMSRequest struct {
	To   string             `json:"to"`
	Body string             `json:"body"`
	Vars *map[string]string `json:"vars"`
}

type snsClient interface {
	Publish(input *sns.PublishInput) (*sns.PublishOutput, error)
}

type SMSService struct {
	Client  snsClient
	cfg     *config.AppConfig
	logger  *LoggerService
	helpers *helpers.HelpersFunction
}

func NewSMSProvider(
	cfg *config.AppConfig,
	logger *LoggerService,
	helpers *helpers.HelpersFunction,
) (*SMSService, error) {
	if cfg == nil {
		return nil, eris.New("configuration cannot be nil")
	}
	if cfg.AWSRegion == "" || cfg.AWSAccessKeyID == "" || cfg.AWSSecretAccessKey == "" {
		logger.Error("Incomplete AWS configuration for SMS service")
		return nil, eris.New("incomplete AWS configuration")
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(cfg.AWSRegion),
		Credentials: credentials.NewStaticCredentials(cfg.AWSAccessKeyID, cfg.AWSSecretAccessKey, ""),
	})
	if err != nil {
		logger.Error("Failed to create AWS session", zap.Error(err))
		return nil, eris.Wrap(err, "failed to create AWS session")
	}

	svc := sns.New(sess)
	logger.Info("AWS SNS session initialized successfully for SMS")
	return &SMSService{
		Client: svc,
		logger: logger,
		cfg:    cfg,
	}, nil
}

func (s *SMSService) FormatSMS(req SMSRequest) (string, error) {
	req.Body = s.helpers.SanitizeBody(req.Body)
	req.Vars = s.helpers.SanitizeMessageVars(req.Vars)

	tmpl, err := template.New("sms").Parse(req.Body)
	if err != nil {
		return "", eris.Wrap(err, "failed to parse SMS template")
	}

	var data interface{}
	if req.Vars != nil {
		data = *req.Vars
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", eris.Wrap(err, "failed to execute SMS template")
	}

	return buf.String(), nil
}

func (s *SMSService) SendSMS(req SMSRequest) error {
	req.To = s.helpers.SanitizePhoneNumber(req.To)

	if req.To == "" {
		err := eris.New("recipient phone number is required")
		s.logger.Error("Missing recipient phone number", zap.Error(err))
		return err
	}

	if len(req.To) < 10 {
		err := eris.Errorf("invalid phone number: %s", req.To)
		s.logger.Error("Invalid recipient phone number", zap.Error(err))
		return err
	}

	if req.Body == "" {
		err := eris.New("message body is required")
		s.logger.Error("Missing SMS body", zap.Error(err))
		return err
	}

	formattedMessage, err := s.FormatSMS(req)
	if err != nil {
		s.logger.Error("Failed to format SMS message", zap.Error(err))
		return eris.Wrap(err, "failed to format SMS message")
	}

	params := &sns.PublishInput{
		Message:     aws.String(formattedMessage),
		PhoneNumber: aws.String(req.To),
	}

	s.logger.Info("Sending SMS via SNS",
		zap.String("to", req.To),
		zap.String("messagePreview", s.helpers.PreviewString(formattedMessage, 60)),
	)

	_, err = s.Client.Publish(params)
	if err != nil {
		s.logger.Error("Failed to send SMS", zap.Error(err), zap.String("to", req.To))
		return eris.Wrapf(err, "failed to send SMS to %s", req.To)
	}

	s.logger.Info("SMS sent successfully",
		zap.String("to", req.To),
		zap.String("messagePreview", s.helpers.PreviewString(formattedMessage, 60)))
	return nil
}
