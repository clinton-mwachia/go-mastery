package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"os"

	"golang.org/x/crypto/chacha20poly1305"
)

const keyFile = "key.txt"
const encryptedFile = "message.enc"

// generateKey creates a new ChaCha20-Poly1305 key.
func generateKey() ([]byte, error) {
	key := make([]byte, chacha20poly1305.KeySize)
	_, err := rand.Read(key)
	return key, err
}

// saveKey writes the key to a file (base64-encoded).
func saveKey(key []byte) error {
	return os.WriteFile(keyFile, []byte(base64.StdEncoding.EncodeToString(key)), 0600)
}

// loadKey reads the key from file (base64-decoded).
func loadKey() ([]byte, error) {
	encodedKey, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(string(encodedKey))
}

// encrypt encrypts a message using ChaCha20-Poly1305.
func encrypt(key []byte, plaintext string) ([]byte, error) {
	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, chacha20poly1305.NonceSize)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := aead.Seal(nonce, nonce, []byte(plaintext), nil)
	return ciphertext, nil

}

// decrypt decrypts a message using ChaCha20-Poly1305.
func decrypt(key, ciphertext []byte) (string, error) {
	if len(ciphertext) < chacha20poly1305.NonceSize {
		return "", errors.New("ciphertext too short")
	}

	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return "", err
	}

	nonce := ciphertext[:chacha20poly1305.NonceSize]
	data := ciphertext[chacha20poly1305.NonceSize:]
	plaintext, err := aead.Open(nil, nonce, data, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [init|encrypt|decrypt]")
		return
	}

	command := os.Args[1]

	switch command {
	case "init":
		key, err := generateKey()
		if err != nil {
			panic(err)
		}
		if err := saveKey(key); err != nil {
			panic(err)
		}
		fmt.Println("Encryption key generated and saved.")

	case "encrypt":
		key, err := loadKey()
		if err != nil {
			panic(err)
		}

		message := "Confidential message: The launch is at 0900 hours."
		ciphertext, err := encrypt(key, message)
		if err != nil {
			panic(err)
		}

		if err := os.WriteFile(encryptedFile, ciphertext, 0600); err != nil {
			panic(err)
		}

		fmt.Println("Message encrypted and saved to file.")

	case "decrypt":
		key, err := loadKey()
		if err != nil {
			panic(err)
		}

		ciphertext, err := os.ReadFile(encryptedFile)
		if err != nil {
			panic(err)
		}

		plaintext, err := decrypt(key, ciphertext)
		if err != nil {
			panic(err)
		}

		fmt.Println("Decrypted message:")
		fmt.Println(plaintext)

	default:
		fmt.Println("Invalid command. Use: init | encrypt | decrypt")
	}

}
