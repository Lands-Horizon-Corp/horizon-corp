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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost+2) // Increase cost if needed
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hashedPassword), nil
}

func VerifyPassword(hashedPassword, password string) bool {
	if hashedPassword == "" {
		return false
	}
	decodedHash, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword(decodedHash, []byte(password))
	return err == nil
}
func Encrypt(data []byte) ([]byte, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, fmt.Errorf("encryption failed")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("encryption failed")
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf("encryption failed")
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], data)

	return ciphertext, nil
}

func Decrypt(ciphertext []byte) ([]byte, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {

		return nil, fmt.Errorf("decryption failed")
	}

	block, err := aes.NewCipher(key)
	if err != nil {

		return nil, fmt.Errorf("decryption failed")
	}
	if len(ciphertext) < aes.BlockSize {

		return nil, fmt.Errorf("decryption failed")
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
