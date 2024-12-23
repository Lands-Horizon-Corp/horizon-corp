package providers

import (
	"bytes"
	"errors"
	"fmt"
	"net/smtp"
	"text/template"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/microcosm-cc/bluemonday"
	"go.uber.org/zap"
)

type MailSender interface {
	SendEmail(req EmailRequest) error
}

type EmailService struct {
	cfg     *config.AppConfig
	logger  *LoggerService
	helpers *helpers.HelpersFunction
	Client  MailSender
}

type EmailRequest struct {
	To      string             `json:"to"`
	Subject string             `json:"subject"`
	Body    string             `json:"body"`
	Vars    *map[string]string `json:"vars,omitempty"`
}

func NewEmailProvider(
	cfg *config.AppConfig,
	logger *LoggerService,
	helpers *helpers.HelpersFunction,
) *EmailService {
	var sender MailSender
	if cfg.EmailHost == "mailhog" {
		sender = &MailHogSender{cfg: cfg, logger: logger}
	} else {
		sender = &AwsSesSender{cfg: cfg, logger: logger}
	}

	return &EmailService{
		cfg:     cfg,
		logger:  logger,
		helpers: helpers,
		Client:  sender,
	}
}

func (es *EmailService) FormatEmail(body string, vars *map[string]string) (string, error) {
	tmpl, err := template.New("email").Parse(body)
	if err != nil {
		return "", fmt.Errorf("failed to parse email template: %w", err)
	}

	var data interface{}
	if vars != nil {
		data = vars
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute email template: %w", err)
	}
	return buf.String(), nil
}

func (es *EmailService) SendEmail(req EmailRequest) error {
	es.logger.Info("Preparing to send email",
		zap.String("to", req.To),
		zap.String("subject", req.Subject),
	)

	// Basic validation
	if req.To == "" {
		err := errors.New("recipient email address is required")
		es.logger.Error("Missing recipient email address", zap.Error(err))
		return err
	}

	if req.Subject == "" {
		err := errors.New("subject is required")
		es.logger.Error("Missing email subject", zap.Error(err))
		return err
	}

	if err := es.helpers.ValidateEmail(req.To); err != nil {
		es.logger.Error("Invalid recipient email address", zap.Error(err), zap.String("to", req.To))
		return err
	}

	formattedBody, err := es.FormatEmail(req.Body, req.Vars)
	if err != nil {
		es.logger.Error("Failed to format email body", zap.Error(err))
		return err
	}

	p := bluemonday.UGCPolicy()
	sanitizedBody := p.Sanitize(formattedBody)
	req.Body = sanitizedBody

	// Use the configured mail sender
	if err := es.Client.SendEmail(req); err != nil {
		es.logger.Error("Failed to send email", zap.Error(err))
		return err
	}

	es.logger.Info("Email sent successfully",
		zap.String("to", req.To),
		zap.String("subject", req.Subject))
	return nil
}

type MailHogSender struct {
	cfg    *config.AppConfig
	logger *LoggerService
}

func (m *MailHogSender) SendEmail(req EmailRequest) error {
	var msg bytes.Buffer
	fmt.Fprintf(&msg, "From: %s\r\n", m.cfg.EmailUser)
	fmt.Fprintf(&msg, "To: %s\r\n", req.To)
	fmt.Fprintf(&msg, "Subject: %s\r\n", req.Subject)
	msg.WriteString("MIME-Version: 1.0\r\n")
	msg.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n")
	msg.WriteString(req.Body)

	smtpAddr := fmt.Sprintf("%s:%s", m.cfg.EmailHost, m.cfg.EmailPort)
	if err := smtp.SendMail(smtpAddr, nil, m.cfg.EmailUser, []string{req.To}, msg.Bytes()); err != nil {
		m.logger.Error("Failed to send email via MailHog",
			zap.String("smtpAddress", smtpAddr),
			zap.Error(err))
		return fmt.Errorf("failed to send email via MailHog: %w", err)
	}

	m.logger.Info("Email sent successfully via MailHog", zap.String("to", req.To))
	return nil
}

type AwsSesSender struct {
	cfg    *config.AppConfig
	logger *LoggerService
}

func (a *AwsSesSender) SendEmail(req EmailRequest) error {
	awsConfig := &aws.Config{
		Region: aws.String(a.cfg.AWSRegion),
	}
	if a.cfg.AWSAccessKeyID != "" && a.cfg.AWSSecretAccessKey != "" {
		awsConfig.Credentials = credentials.NewStaticCredentials(
			a.cfg.AWSAccessKeyID,
			a.cfg.AWSSecretAccessKey,
			"",
		)
	}

	sess, err := session.NewSession(awsConfig)
	if err != nil {
		a.logger.Error("Failed to create AWS session", zap.Error(err))
		return fmt.Errorf("failed to create AWS session: %w", err)
	}

	svc := ses.New(sess)
	input := &ses.SendEmailInput{
		Source: aws.String(a.cfg.EmailUser),
		Destination: &ses.Destination{
			ToAddresses: aws.StringSlice([]string{req.To}),
		},
		Message: &ses.Message{
			Subject: &ses.Content{
				Data: aws.String(req.Subject),
			},
			Body: &ses.Body{
				Html: &ses.Content{
					Data: aws.String(req.Body),
				},
			},
		},
	}

	_, err = svc.SendEmail(input)
	if err != nil {
		a.logger.Error("Failed to send email via AWS SES", zap.Error(err))
		return fmt.Errorf("failed to send email via AWS SES: %w", err)
	}

	a.logger.Info("Email sent successfully via AWS SES",
		zap.String("to", req.To),
		zap.String("subject", req.Subject))
	return nil
}
