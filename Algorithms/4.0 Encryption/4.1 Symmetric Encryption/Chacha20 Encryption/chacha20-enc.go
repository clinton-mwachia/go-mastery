package main

import (
	"crypto/rand"
	"fmt"
	"io"

	"golang.org/x/crypto/chacha20"
)

func main() {
	// Secret key (32 bytes for ChaCha20)
	key := make([]byte, chacha20.KeySize)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err)
	}

	// 12-byte nonce (unique per message)
	nonce := make([]byte, chacha20.NonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	// The plaintext message (e.g. user’s sensitive data)
	plaintext := []byte("Hello, this is a super secret message!")

	fmt.Printf("Original message: %s\n", plaintext)

	// Create a new ChaCha20 cipher stream for encryption
	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		panic(err)
	}

	// Encrypt plaintext (XOR with keystream)
	ciphertext := make([]byte, len(plaintext))
	cipher.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("Encrypted message (hex): %x\n", ciphertext)

	// To decrypt, create a new cipher stream with same key and nonce
	decipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		panic(err)
	}

	// Decrypt ciphertext (XOR with keystream)
	decrypted := make([]byte, len(ciphertext))
	decipher.XORKeyStream(decrypted, ciphertext)

	fmt.Printf("Decrypted message: %s\n", decrypted)
}
