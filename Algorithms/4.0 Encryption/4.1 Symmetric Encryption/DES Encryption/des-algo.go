package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	// Simulating encryption of API key stored in config file
	key := []byte("8bytesk!") // DES requires a 8-byte key

	apiKey := "my-secret-api-key-123"

	encryptedKey, err := encryptDES([]byte(apiKey), key)
	if err != nil {
		log.Fatal("Encryption failed:", err)
	}
	fmt.Println("Encrypted API Key:", encryptedKey)

	decryptedKey, err := decryptDES(encryptedKey, key)
	if err != nil {
		log.Fatal("Decryption failed:", err)
	}
	fmt.Println("Decrypted API Key:", decryptedKey)
}

// pad to block size
func pad(src []byte, blockSize int) []byte {
	padLen := blockSize - len(src)%blockSize
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(src, padding...)
}

// remove padding
func unpad(src []byte) []byte {
	length := len(src)
	unpadLen := int(src[length-1])
	return src[:(length - unpadLen)]
}

// Encrypts plaintext using DES in CBC mode
func encryptDES(plaintext, key []byte) (string, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}

	paddedText := pad(plaintext, block.BlockSize())

	iv := key // in production, use a random IV
	mode := cipher.NewCBCEncrypter(block, iv)

	encrypted := make([]byte, len(paddedText))
	mode.CryptBlocks(encrypted, paddedText)

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// Decrypts ciphertext using DES in CBC mode
func decryptDES(ciphertextBase64 string, key []byte) (string, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		return "", err
	}

	iv := key // in production, use the same IV used during encryption
	mode := cipher.NewCBCDecrypter(block, iv)

	decrypted := make([]byte, len(ciphertext))
	mode.CryptBlocks(decrypted, ciphertext)

	return string(unpad(decrypted)), nil
}
