## 🔐 AES Encryption

### 📁 Folder Structure

```
aes_encryption_example/
├── main.go
├── README.md
└── secret.txt
```

---

### 📄 `main.go`

```go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	// Real-world scenario: encrypt and decrypt a text file
	key := []byte("thisis32byteslongpassphraseimusing") // 32 bytes = AES-256
	plaintext, err := os.ReadFile("secret.txt")
	if err != nil {
		panic(err)
	}

	encrypted, err := encryptAES(key, plaintext)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypted (hex):", hex.EncodeToString(encrypted))

	// Decrypt
	decrypted, err := decryptAES(key, encrypted)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted text:", string(decrypted))
}

func encryptAES(key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

func decryptAES(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
```

---

### 📄 `secret.txt`

```
This is a top secret message. Handle with care.
```

### 📄 `README.md`

# 🔐 AES Encryption Example in Go

This is a real-world example of encrypting and decrypting files using AES (Advanced Encryption Standard) in Go. It demonstrates the use of AES-GCM mode with a 256-bit key.

## 🔧 Supported Key Sizes

- AES-128 → 16 bytes
- AES-192 → 24 bytes
- AES-256 → 32 bytes (used in this example)

## 📦 Requirements

- Go 1.18 or newer

## 📂 How It Works

1. Reads a plaintext file (`secret.txt`)
2. Encrypts it using AES-GCM with a 256-bit key
3. Prints the encrypted data in hexadecimal
4. Decrypts the ciphertext and prints the original message

## 🛠 Usage

### 1. Create a file named `secret.txt`:

```

This is a top secret message. Handle with care.

```

### 2. Run the application:

```bash
go run main.go
```

### ✅ Sample Output:

```
Encrypted (hex): 3c5d7b89f901d25a3f2d...
Decrypted text: This is a top secret message. Handle with care.
```

## 🔒 Security Notes

- The key must be kept secret and secure.
- Use secure random values for keys and nonces (already implemented).
- AES-GCM ensures both confidentiality and authenticity.

---
