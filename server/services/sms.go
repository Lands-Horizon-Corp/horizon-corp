package services

import (
	"horizon/server/config"

	"go.uber.org/zap"
)

type SMSService struct {
	logger *zap.Logger
	cfg    *config.AppConfig
}

func NewSMSService(logger *zap.Logger, cfg *config.AppConfig) *SMSService {
	return &SMSService{logger: logger, cfg: cfg}
}

func (es *SMSService) SendSMS(to, message string) error {

	return nil
}
