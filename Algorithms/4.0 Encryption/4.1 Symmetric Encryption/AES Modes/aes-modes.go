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
	key := []byte("thisis32bitlongpassphraseimusing") // 32 bytes for AES-256
	plaintext := []byte("This is a sensitive user note to be stored securely.")

	fmt.Println("Original:", string(plaintext))

	// CFB Mode
	ciphertextCFB, ivCFB := encryptCFB(plaintext, key)
	decryptedCFB := decryptCFB(ciphertextCFB, key, ivCFB)
	fmt.Println("\nAES-CFB:")
	fmt.Println("Encrypted:", hex.EncodeToString(ciphertextCFB))
	fmt.Println("Decrypted:", string(decryptedCFB))

	// OFB Mode
	ciphertextOFB, ivOFB := encryptOFB(plaintext, key)
	decryptedOFB := decryptOFB(ciphertextOFB, key, ivOFB)
	fmt.Println("\nAES-OFB:")
	fmt.Println("Encrypted:", hex.EncodeToString(ciphertextOFB))
	fmt.Println("Decrypted:", string(decryptedOFB))

	// CTR Mode
	ciphertextCTR, ivCTR := encryptCTR(plaintext, key)
	decryptedCTR := decryptCTR(ciphertextCTR, key, ivCTR)
	fmt.Println("\nAES-CTR:")
	fmt.Println("Encrypted:", hex.EncodeToString(ciphertextCTR))
	fmt.Println("Decrypted:", string(decryptedCTR))
}

// Utility to generate random IV
func generateIV(blockSize int) []byte {
	iv := make([]byte, blockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	return iv
}

// CFB mode
func encryptCFB(plaintext, key []byte) ([]byte, []byte) {
	block, _ := aes.NewCipher(key)
	iv := generateIV(block.BlockSize())
	stream := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)
	return ciphertext, iv
}

func decryptCFB(ciphertext, key, iv []byte) []byte {
	block, _ := aes.NewCipher(key)
	stream := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)
	return plaintext
}

// OFB mode
func encryptOFB(plaintext, key []byte) ([]byte, []byte) {
	block, _ := aes.NewCipher(key)
	iv := generateIV(block.BlockSize())
	stream := cipher.NewOFB(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)
	return ciphertext, iv
}

func decryptOFB(ciphertext, key, iv []byte) []byte {
	block, _ := aes.NewCipher(key)
	stream := cipher.NewOFB(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)
	return plaintext
}

// CTR mode
func encryptCTR(plaintext, key []byte) ([]byte, []byte) {
	block, _ := aes.NewCipher(key)
	iv := generateIV(block.BlockSize())
	stream := cipher.NewCTR(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)
	return ciphertext, iv
}

func decryptCTR(ciphertext, key, iv []byte) []byte {
	block, _ := aes.NewCipher(key)
	stream := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)
	return plaintext
}
