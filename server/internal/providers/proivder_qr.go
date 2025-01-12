package providers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/rotisserie/eris"
	"golang.org/x/crypto/pbkdf2"
)

type QRProvider struct {
	logger *LoggerService
	cfg    *config.AppConfig
}

func NewQRProvider(
	logger *LoggerService,
	cfg *config.AppConfig,
) *QRProvider {
	return &QRProvider{
		logger: logger,
		cfg:    cfg,
	}
}

func (qr *QRProvider) Encode(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", eris.Wrap(err, "failed to marshal data")
	}
	encrypted, err := qr.encrypt(jsonData, string(qr.cfg.AppToken))
	if err != nil {
		return "", eris.Wrap(err, "failed to encrypt data")
	}
	return encrypted, nil
}

func (qr *QRProvider) Decode(encodedString string) (interface{}, error) {
	decryptedData, err := qr.decrypt(encodedString, string(qr.cfg.AppToken))
	if err != nil {
		return nil, eris.Wrap(err, "failed to decrypt data")
	}
	var result interface{}
	if err := json.Unmarshal(decryptedData, &result); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal data")
	}
	return result, nil
}

func (qr *QRProvider) encrypt(data []byte, appToken string) (string, error) {
	salt := make([]byte, 16) // Generate a 16-byte salt
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", eris.Wrap(err, "failed to generate salt")
	}

	// Derive a secure key
	key := deriveKey([]byte(appToken), salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", eris.Wrap(err, "failed to create cipher block")
	}

	nonce := make([]byte, 12) // GCM standard nonce size
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", eris.Wrap(err, "failed to generate nonce")
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", eris.Wrap(err, "failed to create GCM")
	}

	// Encrypt and include salt as part of the ciphertext
	ciphertext := aesGCM.Seal(append(salt, nonce...), nonce, data, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (qr *QRProvider) decrypt(encryptedData string, appToken string) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, eris.Wrap(err, "failed to decode base64 data")
	}

	// Extract salt and nonce
	if len(ciphertext) < 28 { // 16 bytes salt + 12 bytes nonce
		return nil, eris.New("ciphertext too short")
	}
	salt, nonce, ciphertext := ciphertext[:16], ciphertext[16:28], ciphertext[28:]

	// Derive the key using the salt
	key := deriveKey([]byte(appToken), salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create cipher block")
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create GCM")
	}

	// Decrypt
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to decrypt data")
	}

	return plaintext, nil
}

// deriveKey derives a secure key from a password and a salt.
func deriveKey(password, salt []byte) []byte {
	// PBKDF2 parameters:
	// - Password: The input password or token
	// - Salt: A cryptographic salt (unique per use case)
	// - Iterations: A high iteration count (e.g., 100,000 for security)
	// - Key length: Desired key length (32 bytes for AES-256)
	// - Hash function: SHA-256
	return pbkdf2.Key(password, salt, 100000, 32, sha256.New)
}
