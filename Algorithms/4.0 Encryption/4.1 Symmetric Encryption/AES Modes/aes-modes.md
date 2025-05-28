## AES Encryption Modes in Go

This project demonstrates how to use AES in **CFB**, **OFB**, and **CTR** modes using the Go `crypto/cipher` standard library.

## 💡 Real-World Scenario

The example simulates a note-keeping app that encrypts user notes before saving them to disk. Depending on your need for speed or security:

- Use **CFB** when you need a stream cipher that's resistant to bit-flipping attacks.
- Use **OFB** for simple stream cipher use (less common).
- Use **CTR** when you want parallel encryption or high performance.

## 🔧 How It Works

- A fixed 32-byte AES-256 key is used.
- A random IV is generated for each mode.
- The original plaintext is encrypted and decrypted back using each mode.
- Results are printed to the terminal.

## ▶️ Run

```bash
go run main.go
```

## 🔐 Output Example

```
Original: This is a sensitive user note to be stored securely.

AES-CFB:
Encrypted: 8f93b6...
Decrypted: This is a sensitive user note to be stored securely.

AES-OFB:
Encrypted: 5ee1f2...
Decrypted: This is a sensitive user note to be stored securely.

AES-CTR:
Encrypted: a2c7bd...
Decrypted: This is a sensitive user note to be stored securely.
```

## 📁 Notes

- Always use a **unique IV per encryption**.
- Store the IV alongside the ciphertext for decryption.
- These modes **don't provide authentication** (consider AES-GCM for that).

---
