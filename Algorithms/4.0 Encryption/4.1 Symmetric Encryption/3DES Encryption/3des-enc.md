# Triple DES (3DES) Encryption

This is a simple example showing how to use 3DES (Triple Data Encryption Standard) in Go to encrypt and decrypt sensitive data such as credit card numbers.

## 🔐 Why 3DES?

While AES is now the preferred standard, 3DES is still used in legacy systems and some payment systems. This demo shows how to handle such scenarios using Go's `crypto/des` package.

## 📁 Structure

- `3des-enc.go`: Contains encryption, decryption, and padding logic.
- `3des-enc.md`: This documentation.

## ▶️ Run the Code

```bash
go run 3des-enc.go
```

### 🧪 Output

```text
Original Card Number: 4111111111111111
Encrypted: TbnICv0ydZdw0Xf/+WQZSgH7AihnDK40
Decrypted: 4111111111111111
```

## ⚠️ Security Notice

3DES is outdated and **not recommended for new systems**. Use **AES-GCM** or **ChaCha20-Poly1305** for modern applications. This project is meant for educational or legacy system purposes only.
