package storage

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

func UniqueFileName(originalName string) string {
	fileExt := filepath.Ext(originalName)
	fileBase := strings.TrimSuffix(originalName, fileExt)

	randomID := uuid.New().String()
	timestamp := time.Now().Unix()

	return fmt.Sprintf("%s_%d_%s%s", fileBase, timestamp, randomID, fileExt)
}
