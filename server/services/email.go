package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"horizon/server/config"
	"net/http"

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
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func NewEmailService(logger *zap.Logger, cfg *config.AppConfig) *EmailService {
	return &EmailService{logger: logger, cfg: cfg}
}

func (es *EmailService) SendEmail(req EmailRequest) error {
	es.logger.Info("Sending email", zap.String("to", req.To), zap.String("subject", req.Subject))

	if es.cfg.EmailHost == "localhost" {
		return es.sendEmailWithMailHog(req.To, req.Subject, req.Body)
	}
	return es.sendEmailWithAWS(req.To, req.Subject, req.Body)
}

func (es *EmailService) sendEmailWithMailHog(to, subject, body string) error {
	msg := map[string]interface{}{
		"to":      to,
		"subject": subject,
		"body":    body,
	}
	msgJSON, err := json.Marshal(msg)
	if err != nil {
		es.logger.Error("Failed to marshal email message", zap.Error(err))
		return err
	}
	mailHogURL := fmt.Sprintf("http://%s:%s/api/v2/send", es.cfg.EmailHost, es.cfg.EmailPort)
	resp, err := http.Post(mailHogURL, "application/json", bytes.NewBuffer(msgJSON))
	if err != nil {
		es.logger.Error("Failed to send email via MailHog", zap.Error(err))
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		es.logger.Error("Failed to send email via MailHog", zap.String("status", resp.Status))
		return fmt.Errorf("failed to send email via MailHog: %s", resp.Status)
	}
	return nil
}

func (es *EmailService) sendEmailWithAWS(to, subject, body string) error {
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
			ToAddresses: aws.StringSlice([]string{to}),
		},
		Message: &ses.Message{
			Subject: &ses.Content{
				Data: aws.String(subject),
			},
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(body),
				},
			},
		},
	}

	_, err = svc.SendEmail(input)
	if err != nil {
		es.logger.Error("Failed to send email via AWS SES", zap.Error(err))
		return err
	}

	es.logger.Info("Email sent successfully", zap.String("to", to))
	return nil
}
