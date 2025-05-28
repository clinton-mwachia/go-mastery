## 🔐 AES-GCM Encryption

This example demonstrates how to use AES-GCM (Galois/Counter Mode) encryption in Go to securely encrypt and decrypt sensitive data such as passwords.

## 📦 Features

- AES-256 encryption with integrity verification
- Uses Go's standard `crypto/aes` and `crypto/cipher` packages
- Hex encoding for easy storage (e.g., in databases)
- Real-world usage: encrypting user passwords before storing them

## 🚀 Getting Started

### Requirements

- Go 1.18+

### Run the Example

```bash
go run aes-gcm.go
```

### Example Output

```
Encrypted Password (Hex): 7758bbbea9b2d1e5bbe...
Decrypted Password: SuperSecret123!
```

## 🔒 Security Notes

- Always use a securely generated key and store it in a secure place (e.g., env variables, Vault).
- Never reuse nonces with the same key.
- AES-GCM not only encrypts the data but also ensures its integrity.

## 📚 Learn More

- [Go AES-GCM Docs](https://pkg.go.dev/crypto/cipher#NewGCM)

---

✅ **Recommended For:**

- Secure password storage
- Encrypting personal data in files or databases
- Any application that needs both encryption and authentication

---
