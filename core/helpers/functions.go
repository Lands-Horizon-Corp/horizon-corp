package helpers

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func FileSizeStringToInt64(data string) int64 {
	if intValue, err := strconv.ParseInt(data, 10, 64); err == nil {
		return intValue
	}
	return 10 << 20
}

func UniqueFileName(originalName string) string {
	timestamp := time.Now().Unix()
	fileExt := filepath.Ext(originalName)
	fileBase := strings.TrimSuffix(originalName, fileExt)
	return fmt.Sprintf("%s_%d%s", fileBase, timestamp, fileExt)
}
