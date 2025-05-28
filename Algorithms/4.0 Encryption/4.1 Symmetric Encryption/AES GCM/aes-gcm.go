package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	// Key must be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256
	key := []byte("thisis32bitlongpassphraseimusing") // 32 bytes = AES-256

	// Simulate encrypting user password before storing
	userPassword := "SuperSecret123!"
	encrypted, err := Encrypt(key, userPassword)
	if err != nil {
		panic(err)
	}

	fmt.Println("Encrypted Password (Hex):", encrypted)

	// Decrypt to use it later
	decrypted, err := Decrypt(key, encrypted)
	if err != nil {
		panic(err)
	}

	fmt.Println("Decrypted Password:", decrypted)
}

// Encrypt encrypts plaintext using AES-GCM with a randomly generated nonce.
func Encrypt(key []byte, plaintext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %w", err)
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	return hex.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the AES-GCM encrypted hex string.
func Decrypt(key []byte, encryptedHex string) (string, error) {
	ciphertext, err := hex.DecodeString(encryptedHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode hex: %w", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt: %w", err)
	}

	return string(plaintext), nil
}
