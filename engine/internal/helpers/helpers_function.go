package helpers

import (
	"errors"
	"fmt"
	"net/mail"
	"path/filepath"
	"strings"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/google/uuid"
)

type HelpersFunction struct {
	cfg *config.AppConfig
}

func NewHelperFunction(
	cfg *config.AppConfig,
) *HelpersFunction {
	return &HelpersFunction{
		cfg: cfg,
	}
}

func (hf *HelpersFunction) UniqueFileName(originalName string) string {
	if originalName == "" {
		originalName = "file"
	}

	fileExt := filepath.Ext(originalName)
	fileBase := strings.TrimSuffix(originalName, fileExt)
	if fileBase == "" {
		fileBase = "file"
	}

	randomID := uuid.New().String()
	timestamp := time.Now().Unix()
	uniqueName := fmt.Sprintf("%s_%d_%s%s", fileBase, timestamp, randomID, fileExt)

	return uniqueName
}

func (hf *HelpersFunction) ValidateEmail(emailAddr string) error {
	_, err := mail.ParseAddress(emailAddr)
	if err != nil {
		return errors.New("invalid email address")
	}

	return nil
}

func (hf *HelpersFunction) PreviewString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
