# Go RSA Encryption

A lightweight, pure-Go example demonstrating RSA public-key encryption and decryption using OAEP (SHA-256).

## 📘 Overview

This program shows how to:

1. Generate a 2048-bit RSA key pair.
2. Export keys to PEM format.
3. Encrypt plaintext using the public key (`EncryptOAEP`).
4. Decrypt ciphertext using the private key (`DecryptOAEP`).

## ⚙️ Requirements

- Go 1.16 or later.
- No external dependencies—uses only the Go standard library.

## Observe:

- PEM-encoded private and public keys.
- Hex-encoded ciphertext.
- Decrypted plaintext confirmation.

## 🧠 Explanation

- `rsa.GenerateKey(...)`: Creates RSA private key with corresponding public key.
- `pem.EncodeToMemory(...)`: Converts binary key data to human-readable PEM format.
- `rsa.EncryptOAEP(...)` and `rsa.DecryptOAEP(...)`: Use OAEP padding scheme with SHA-256 for encryption/decryption.

## 🔐 Security Notes

- **Key size**: 2048-bit RSA is considered secure for most use cases.
- **Padding**: OAEP is recommended over PKCS#1 v1.5 padding.
- **Data size limit**: RSA can only encrypt data smaller than `keySize–2*hashLen–2`. For larger data, use RSA to encrypt symmetric keys.

## 🚀 Extending the Example

- Implement _signing_ with `rsa.SignPSS` and `rsa.VerifyPSS` using `crypto/sha256`.
- PEM import/export to/from files using `crypto/x509` and `encoding/pem`.

## 📚 References

- Go [crypto/rsa](https://pkg.go.dev/crypto/rsa) documentation
