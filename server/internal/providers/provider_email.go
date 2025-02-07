package providers

import (
	"bytes"
	"text/template"

	"net/smtp"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/microcosm-cc/bluemonday"
	"github.com/rotisserie/eris"
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
		return "", eris.Wrap(err, "failed to parse email template")
	}

	var data interface{}
	if vars != nil {
		data = vars
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", eris.Wrap(err, "failed to execute email template")
	}
	return buf.String(), nil
}

func (es *EmailService) SendEmail(req EmailRequest) error {
	es.logger.Info("Preparing to send email",
		zap.String("to", req.To),
		zap.String("subject", req.Subject),
	)

	if req.To == "" {
		err := eris.New("recipient email address is required")
		es.logger.Error("Missing recipient email address", zap.Error(err))
		return err
	}

	if req.Subject == "" {
		err := eris.New("subject is required")
		es.logger.Error("Missing email subject", zap.Error(err))
		return err
	}

	if err := es.helpers.ValidateEmail(req.To); err != nil {
		es.logger.Error("Invalid recipient email address", zap.Error(err), zap.String("to", req.To))
		return eris.Wrap(err, "invalid recipient email address")
	}

	formattedBody, err := es.FormatEmail(req.Body, req.Vars)
	if err != nil {
		es.logger.Error("Failed to format email body", zap.Error(err))
		return eris.Wrap(err, "failed to format email body")
	}

	p := bluemonday.UGCPolicy()
	sanitizedBody := p.Sanitize(formattedBody)
	req.Body = sanitizedBody

	if err := es.Client.SendEmail(req); err != nil {
		es.logger.Error("Failed to send email", zap.Error(err))
		return eris.Wrap(err, "failed to send email")
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
	msg.WriteString("From: " + m.cfg.EmailUser + "\r\n")
	msg.WriteString("To: " + req.To + "\r\n")
	msg.WriteString("Subject: " + req.Subject + "\r\n")
	msg.WriteString("MIME-Version: 1.0\r\n")
	msg.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n")
	msg.WriteString(req.Body)

	smtpAddr := m.cfg.EmailHost + ":" + m.cfg.EmailPort
	if err := smtp.SendMail(smtpAddr, nil, m.cfg.EmailUser, []string{req.To}, msg.Bytes()); err != nil {
		m.logger.Error("Failed to send email via MailHog",
			zap.String("smtpAddress", smtpAddr),
			zap.Error(err))
		return eris.Wrap(err, "failed to send email via MailHog")
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
		return eris.Wrap(err, "failed to create AWS session")
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
		return eris.Wrap(err, "failed to send email via AWS SES")
	}

	a.logger.Info("Email sent successfully via AWS SES",
		zap.String("to", req.To),
		zap.String("subject", req.Subject))
	return nil
}
