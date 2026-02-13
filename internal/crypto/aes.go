package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// EncryptResourceID encrypts a plaintext resource ID using AES-GCM.
// It returns the hex-encoded ciphertext.
// The key must be exactly 32 bytes (AES-256).
func EncryptResourceID(plaintext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// GCM (Galois/Counter Mode) is used for authenticated encryption.
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Create a random nonce (number used once).
	// Standard nonce size for GCM is 12 bytes.
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Seal encrypts and authenticates the plaintext.
	// The nonce is prepended to the ciphertext to allow decryption later.
	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)

	// Return as hex string for safe JSON transmission
	return hex.EncodeToString(ciphertext), nil
}

// GenerateRandomKey creates a secure 32-byte key for AES-256.
func GenerateRandomKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}