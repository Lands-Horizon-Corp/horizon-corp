package services

import (
	"errors"
	"fmt"
	"horizon/server/config"
	"horizon/server/database"
	"time"

	"go.uber.org/zap"
)

type OTPService struct {
	logger       *zap.Logger
	cfg          *config.AppConfig
	emailService *EmailService
	cacheService *database.CacheService
	smsService   *SMSService
}

func NewOTPService(
	logger *zap.Logger,
	cfg *config.AppConfig,
	emailService *EmailService,
	cacheService *database.CacheService,
	smsService *SMSService,
) *OTPService {
	return &OTPService{
		logger:       logger,
		cfg:          cfg,
		emailService: emailService,
		cacheService: cacheService,
		smsService:   smsService,
	}
}

// cacheKey generates a unique cache key for storing OTPs by user ID.
func (es *OTPService) cacheKey(accountType string, id uint, mediumType string) string {
	return fmt.Sprintf("otp:%s:%d:%s", accountType, id, mediumType)
}

// generateOTP generates a secure, 6-digit OTP.
func (es *OTPService) generateOTP() (string, error) {
	otp, err := config.GenerateSecureRandom6DigitNumber()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", otp), nil
}

// generateAndStoreOTP generates an OTP, stores it in the cache, and includes retry handling.
func (es *OTPService) generateAndStoreOTP(accountType string, id uint, mediumType string) (string, error) {
	otpStr, err := es.generateOTP()
	if err != nil {
		return "", err
	}

	key := es.cacheKey(accountType, id, mediumType)
	expiration := 10 * time.Minute // Or make this configurable

	if err := es.cacheService.StoreOTP(key, otpStr, expiration); err != nil {
		es.logger.Error("Failed to store OTP in cache", zap.Error(err))
		return "", err
	}

	return otpStr, nil
}

// SendEmailOTP generates and sends an OTP via email.
func (es *OTPService) SendEmailOTP(accountType string, id uint, req EmailRequest) error {
	otpStr, err := es.generateAndStoreOTP(accountType, id, "email")
	if err != nil {
		return err
	}

	if req.Vars == nil {
		req.Vars = &map[string]string{}
	}

	(*req.Vars)["otp"] = otpStr

	if err := es.emailService.SendEmail(req); err != nil {
		es.logger.Error("Failed to send email", zap.Error(err))
		return err
	}

	es.logger.Info("OTP sent via email", zap.Uint("user_id", id))
	return nil
}

// SendEContactNumberOTP generates and sends an OTP via SMS.
func (es *OTPService) SendEContactNumberOTP(accountType string, id uint, req SMSRequest) error {
	otpStr, err := es.generateAndStoreOTP(accountType, id, "sms")
	if err != nil {
		return err
	}

	if req.Vars == nil {
		req.Vars = &map[string]string{}
	}

	(*req.Vars)["otp"] = otpStr

	if err := es.smsService.SendSMS(req); err != nil {
		es.logger.Error("Failed to send SMS", zap.Error(err))
		return err
	}

	es.logger.Info("OTP sent via SMS", zap.Uint("user_id", id))
	return nil
}

func (es *OTPService) ValidateOTP(accountType string, id uint, providedOTP string, mediumType string) (bool, error) {
	key := es.cacheKey(accountType, id, mediumType)

	storedOTP, err := es.cacheService.GetOTP(key)
	if err != nil {
		if errors.Is(err, database.ErrCacheMiss) {
			es.logger.Warn("OTP not found or expired", zap.Uint("user_id", id))
			return false, nil
		}
		es.logger.Error("Failed to retrieve OTP from cache", zap.Error(err))
		return false, err
	}

	if providedOTP != storedOTP {
		es.logger.Warn("Invalid OTP provided", zap.Uint("user_id", id))
		return false, nil
	}

	if err := es.cacheService.DeleteOTP(key); err != nil {
		es.logger.Error("Failed to delete OTP from cache after validation", zap.Error(err), zap.Uint("user_id", id))
	}

	es.logger.Info("OTP validated successfully", zap.Uint("user_id", id))
	return true, nil
}
