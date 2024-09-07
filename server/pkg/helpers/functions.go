package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

func GetUintParam(c *gin.Context, paramName string) (uint, error) {
	paramValue := c.Param(paramName)
	if paramValue == "" {
		return 0, fmt.Errorf("parameter %s is missing", paramName)
	}
	id, err := strconv.ParseUint(paramValue, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid parameter %s: %v", paramName, err)
	}
	return uint(id), nil
}

func CreateSessionToken(userID uint) (string, error) {
	// Define the size of the token
	tokenBytes := make([]byte, 32) // 32 bytes will give 256 bits of entropy

	// Read random bytes into the slice
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", fmt.Errorf("error generating token: %v", err)
	}

	// Encode the bytes to a base64 string
	tokenString := base64.URLEncoding.EncodeToString(tokenBytes)

	// You might want to prepend or append user-specific data
	// tokenString = fmt.Sprintf("uid%d:%s", userID, tokenString)

	return tokenString, nil
}
