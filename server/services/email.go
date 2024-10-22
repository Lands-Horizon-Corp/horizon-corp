package services

import (
	"bytes"
	"fmt"
	"horizon/server/config"
	"net/smtp"
	"text/template"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"go.uber.org/zap"
)

type EmailService struct {
	logger *zap.Logger
	cfg    *config.AppConfig
}

type EmailRequest struct {
	To      string             `json:"to"`
	Subject string             `json:"subject"`
	Body    string             `json:"body"`
	Vars    *map[string]string `json:"vars"`
}

func NewEmailService(logger *zap.Logger, cfg *config.AppConfig) *EmailService {
	return &EmailService{logger: logger, cfg: cfg}
}
func (es *EmailService) FormatEmail(body string, vars *map[string]string) (string, error) {
	tmpl, err := template.New("email").Parse(body)
	if err != nil {
		return "", fmt.Errorf("failed to parse email template: %w", err)
	}

	var buf bytes.Buffer
	if vars != nil {
		if err := tmpl.Execute(&buf, vars); err != nil {
			return "", fmt.Errorf("failed to execute email template: %w", err)
		}
	} else {
		if err := tmpl.Execute(&buf, nil); err != nil {
			return "", fmt.Errorf("failed to execute email template without variables: %w", err)
		}
	}

	return buf.String(), nil
}

func (es *EmailService) SendEmail(req EmailRequest) error {
	es.logger.Info("Sending email", zap.String("to", req.To), zap.String("subject", req.Subject))
	formattedBody, err := es.FormatEmail(req.Body, req.Vars)
	req.Body = formattedBody
	if err != nil {
		es.logger.Error("Failed to format email body", zap.Error(err))
		return err
	}
	if es.cfg.EmailHost == "mailhog" {
		return es.sendEmailWithMailHog(req)
	}
	return es.sendEmailWithAWS(req)
}

func (es *EmailService) sendEmailWithMailHog(req EmailRequest) error {
	var msg bytes.Buffer
	msg.WriteString(fmt.Sprintf("From: %s\r\n", es.cfg.EmailUser))
	msg.WriteString(fmt.Sprintf("To: %s\r\n", req.To))
	msg.WriteString(fmt.Sprintf("Subject: %s\r\n", req.Subject))
	msg.WriteString("MIME-Version: 1.0\r\n")
	msg.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n")
	msg.WriteString("\r\n")
	msg.WriteString(req.Body)

	smtpAddr := fmt.Sprintf("%s:%s", es.cfg.EmailHost, es.cfg.EmailPort)
	err := smtp.SendMail(smtpAddr, nil, es.cfg.EmailUser, []string{req.To}, msg.Bytes())
	if err != nil {
		es.logger.Error("Failed to send email via MailHog", zap.Error(err))
		return err
	}

	es.logger.Info("Email sent successfully via MailHog")
	return nil
}

func (es *EmailService) sendEmailWithAWS(req EmailRequest) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(es.cfg.AWSRegion),
		Credentials: credentials.NewStaticCredentials(
			es.cfg.AWSAccessKeyID,
			es.cfg.AWSSecretAccessKey,
			"",
		),
	})
	if err != nil {
		es.logger.Error("Failed to create AWS session", zap.Error(err))
		return err
	}

	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Source: aws.String(es.cfg.EmailUser),
		Destination: &ses.Destination{
			ToAddresses: aws.StringSlice([]string{req.To}),
		},
		Message: &ses.Message{
			Subject: &ses.Content{
				Data: aws.String(req.Subject),
			},
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(req.Body),
				},
			},
		},
	}

	_, err = svc.SendEmail(input)
	if err != nil {
		es.logger.Error("Failed to send email via AWS SES", zap.Error(err))
		return err
	}

	es.logger.Info("Email sent successfully", zap.String("to", req.To))
	return nil
}
