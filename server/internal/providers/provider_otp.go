package providers

import (
	"fmt"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"

	"github.com/rotisserie/eris"
	"go.uber.org/zap"
)

// OTPService handles OTP generation, storage, validation, and sending via email or SMS.
type OTPService struct {
	cfg          *config.AppConfig
	logger       *LoggerService
	helpers      *helpers.HelpersCryptography
	mail         *EmailService
	sms          *SMSService
	cacheService *CacheService
}

type MediumType string

const (
	SMS   MediumType = "sms"
	Email MediumType = "email"
)

type OTPMessage struct {
	AccountType string
	ID          string
	MediumType  MediumType
	Reference   string
}

// NewOTPProvider initializes a new instance of OTPService.
func NewOTPProvider(
	cfg *config.AppConfig,
	logger *LoggerService,
	helpers *helpers.HelpersCryptography,
	cacheService *CacheService,
	mail *EmailService,
	sms *SMSService,
) *OTPService {
	return &OTPService{
		cfg:          cfg,
		logger:       logger,
		helpers:      helpers,
		mail:         mail,
		sms:          sms,
		cacheService: cacheService,
	}
}

// cacheKey generates a cache key for storing OTPs.
func (os *OTPService) cacheKey(msg OTPMessage) string {
	return fmt.Sprintf("otp_%s_%s_%s_%s", msg.AccountType, msg.ID, msg.MediumType, msg.Reference)
}

// SendEmailOTP generates an OTP, stores it, and sends it via email.
func (os *OTPService) SendEmailOTP(msg OTPMessage, req EmailRequest) error {
	otp, err := os.generateAndStoreOTP(msg)
	if err != nil {
		return eris.Wrap(err, "failed to generate and store OTP for email")
	}

	if req.Vars == nil {
		req.Vars = &map[string]string{}
	}
	(*req.Vars)["otp"] = otp

	if err := os.mail.SendEmail(req); err != nil {
		os.logger.Error("Failed to send email OTP", zap.Error(err), zap.String("user_id", msg.ID))
		return eris.Wrap(err, "failed to send email OTP")
	}

	os.logger.Info("OTP sent via email", zap.String("user_id", msg.ID))
	return nil
}

// SendContactNumberOTP generates an OTP, stores it, and sends it via SMS.
func (os *OTPService) SendContactNumberOTP(msg OTPMessage, req SMSRequest) error {
	otp, err := os.generateAndStoreOTP(msg)
	if err != nil {
		return eris.Wrap(err, "failed to generate and store OTP for SMS")
	}

	if req.Vars == nil {
		req.Vars = &map[string]string{}
	}
	(*req.Vars)["otp"] = otp

	if err := os.sms.SendSMS(req); err != nil {
		os.logger.Error("Failed to send SMS OTP", zap.Error(err), zap.String("user_id", msg.ID))
		return eris.Wrap(err, "failed to send SMS OTP")
	}

	os.logger.Info("OTP sent via SMS", zap.String("user_id", msg.ID))
	return nil
}

// generateAndStoreOTP generates a secure OTP, hashes it if needed, and stores it in the cache.
func (os *OTPService) generateAndStoreOTP(msg OTPMessage) (string, error) {
	otp, err := os.helpers.GenerateSecureRandom6DigitNumber()
	if err != nil {
		os.logger.Error("Failed to generate OTP", zap.Error(err))
		return "", eris.Wrap(err, "failed to generate OTP")
	}

	otpStr := fmt.Sprintf("%06d", otp)
	key := os.cacheKey(msg)
	expiration := 10 * time.Minute

	storedOTP := otpStr
	if os.cfg.AppEnv == "development" {
		os.logger.Info(fmt.Sprintf("[DEBUG] OTP before hashing: %s", otpStr))
	}
	if os.isHashingEnabled() {
		storedOTP, err = os.helpers.HashPassword(otpStr)
		if err != nil {
			os.logger.Error("Failed to hash OTP", zap.Error(err))
			return "", eris.Wrap(err, "failed to hash OTP")
		}
	}
	if err := os.cacheService.Set(key, storedOTP, expiration); err != nil {
		os.logger.Error("Failed to store OTP in cache", zap.Error(err))
		return "", eris.Wrap(err, "failed to store OTP in cache")
	}
	return otpStr, nil
}

// ValidateOTP validates the provided OTP against the stored OTP in the cache.
func (os *OTPService) ValidateOTP(msg OTPMessage, providedOTP string) (bool, error) {
	if providedOTP == "" {
		os.logger.Warn("Provided OTP is empty", zap.String("user_id", msg.ID))
		return false, eris.New("provided OTP is invalid: empty input")
	}

	key := os.cacheKey(msg)
	storedOTP, err := os.cacheService.Get(key)
	if err != nil {
		if eris.Is(err, eris.New("cache miss: key does not exist")) {
			os.logger.Warn("OTP not found or expired", zap.String("user_id", msg.ID))
			return false, nil
		}
		os.logger.Error("Failed to retrieve OTP from cache", zap.Error(err))
		return false, eris.Wrap(err, "failed to retrieve OTP from cache")
	}

	if os.isHashingEnabled() {
		if !os.helpers.VerifyPassword(storedOTP, providedOTP) {
			os.logger.Warn("Provided OTP does not match stored OTP", zap.String("user_id", msg.ID))
			return false, nil
		}
	} else if providedOTP != storedOTP {
		os.logger.Warn("Provided OTP does not match stored OTP", zap.String("user_id", msg.ID))
		return false, nil
	}

	if err := os.cacheService.Delete(key); err != nil {
		os.logger.Error("Failed to delete OTP from cache after validation", zap.Error(err), zap.String("user_id", msg.ID))
		return true, eris.Wrap(err, "failed to delete OTP from cache after validation")
	}

	return true, nil
}

// isHashingEnabled checks if hashing is enabled based on the app environment.
func (os *OTPService) isHashingEnabled() bool {
	return os.cfg.AppEnv == "staging" || os.cfg.AppEnv == "production"
}
