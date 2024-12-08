package helpers

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type HelpersFunction struct {
	cfg *config.AppConfig
}

func NewHelperFunction(
	cfg *config.AppConfig,
	logger *zap.Logger,
) *HelpersFunction {
	return &HelpersFunction{
		cfg: cfg,
	}
}

func (hf *HelpersFunction) UniqueFileName(originalName string) string {
	fileExt := filepath.Ext(originalName)
	fileBase := strings.TrimSuffix(originalName, fileExt)

	randomID := uuid.New().String()
	timestamp := time.Now().Unix()

	return fmt.Sprintf("%s_%d_%s%s", fileBase, timestamp, randomID, fileExt)
}
