## Secure Message Encryption with ChaCha20-Poly1305 in Go

This Go project demonstrates how to securely encrypt and decrypt text messages using the `ChaCha20-Poly1305` encryption algorithm.

## 🔐 Features

- Generates a secure random 256-bit key
- Encrypts a real-world text message using AEAD ChaCha20-Poly1305
- Saves the encrypted message to a file
- Decrypts and displays the original message
- Uses base64 encoding for key storage

## 📦 Requirements

- Go 1.17 or newer
- Internet (for initial `golang.org/x/crypto` fetch)

### 1. Initialize and Generate Key

```bash
go run chacha20-poly1305.go init
```

This creates a `key.txt` file with a base64-encoded encryption key.

### 2. Encrypt a Message

```bash
go run chacha20-poly1305.go encrypt
```

This encrypts a hardcoded message and saves it as `message.enc`.

### 3. Decrypt the Message

```bash
go run chacha20-poly1305.go decrypt
```

This reads the encrypted file, decrypts it, and prints the original message.

## 🧪 Example Output

```
$ go run main.go init
Encryption key generated and saved.

$ go run main.go encrypt
Message encrypted and saved to file.

$ go run main.go decrypt
Decrypted message:
Confidential message: The launch is at 0900 hours.
```

## 🛡️ Security Notes

- This example uses hardcoded messages for simplicity. In production, read user input or accept messages as function arguments.
- Ensure keys and encrypted files are stored securely.
- Never reuse nonces with the same key.

---
