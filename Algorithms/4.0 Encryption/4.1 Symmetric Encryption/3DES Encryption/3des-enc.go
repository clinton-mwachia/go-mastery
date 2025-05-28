package main

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"errors"
	"fmt"
)

func main() {
	cardNumber := "4111111111111111" // Sample Visa test card number
	fmt.Println("Original Card Number:", cardNumber)

	encrypted, err := Encrypt(cardNumber)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypted:", encrypted)

	decrypted, err := Decrypt(encrypted)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted:", decrypted)

}

// 24-byte key for 3DES
var key = []byte("examplekey12345678901234") // Must be exactly 24 bytes

// Encrypt encrypts plain text using 3DES
func Encrypt(plaintext string) (string, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}

	plaintextBytes := []byte(plaintext)
	plaintextBytes = pad(plaintextBytes, block.BlockSize())

	iv := key[:block.BlockSize()]
	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintextBytes))
	mode.CryptBlocks(ciphertext, plaintextBytes)

	return base64.StdEncoding.EncodeToString(ciphertext), nil

}

// Decrypt decrypts base64 encoded 3DES ciphertext
func Decrypt(encrypted string) (string, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	if len(ciphertext)%block.BlockSize() != 0 {
		return "", errors.New("ciphertext is not a multiple of block size")
	}

	iv := key[:block.BlockSize()]
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)
	plaintext = unpad(plaintext)

	return string(plaintext), nil

}

// pad pads plaintext using PKCS#5 padding
func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytesRepeat(byte(padding), padding)
	return append(src, padText...)
}

// unpad removes PKCS#5 padding
func unpad(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

// bytesRepeat returns a new byte slice consisting of count copies of b
func bytesRepeat(b byte, count int) []byte {
	result := make([]byte, count)
	for i := range result {
		result[i] = b
	}
	return result
}
