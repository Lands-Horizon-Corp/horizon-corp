package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"math/big"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func GenerateSalt(length int) (string, error) {
	if length < 16 {
		return "", errors.New("salt length should be at least 16 bytes")
	}

	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	return base64.RawStdEncoding.EncodeToString(salt), nil
}

func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}
	salt, err := GenerateSalt(16)
	if err != nil {
		return "", err
	}
	saltedPassword := fmt.Sprintf("%s%s", password, salt)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost+2) // Increase cost by 2
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", base64.StdEncoding.EncodeToString(hashedPassword), salt), nil
}

func VerifyPassword(hashedPassword, password string) bool {
	if hashedPassword == "" {
		return false
	}
	parts := strings.Split(hashedPassword, ":")
	if len(parts) != 2 {
		return false
	}
	hashedPwd, salt := parts[0], parts[1]
	saltedPassword := fmt.Sprintf("%s%s", password, salt)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(saltedPassword))
	return err == nil
}

func Encrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], data)
	return ciphertext, nil
}

func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	return ciphertext, nil
}

func GenerateSecureRandom6DigitNumber() (int, error) {
	min := int64(100000)
	max := int64(999999)

	n, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
	if err != nil {
		return 0, err
	}
	return int(n.Int64() + min), nil
}
