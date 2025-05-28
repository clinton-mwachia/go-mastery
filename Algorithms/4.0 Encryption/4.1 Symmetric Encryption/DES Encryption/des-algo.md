## 🔐 DES Encryption Example in Go

This is a simple example of using DES encryption in Go to secure sensitive configuration values like API keys.

> ⚠️ **DES is considered insecure for most modern applications.** This example is for educational and legacy compatibility use cases only. Use AES or ChaCha20-Poly1305 for secure systems.

## 💡 Use Case

In a desktop app, you might want to encrypt sensitive config values such as API keys before storing them to disk.

This example encrypts an API key using **DES in CBC mode** and then decrypts it back.

## 📂 Files

- `des-algo.go`: Contains the full Go code for encryption and decryption.
- `des-algo.md`: This file.

## 🔧 Requirements

- Go 1.16 or higher

## 🚀 Run

```bash
go run main.go
```

### 📤 Output

```
Encrypted API Key: <base64-encrypted-value>
Decrypted API Key: my-secret-api-key-123
```

## 🔒 Security Note

- Always use **AES-GCM** or **ChaCha20-Poly1305** for strong encryption.
- This example uses the **key as the IV**, which is insecure for real use. Use a **random IV** in production and store/transmit it securely.

## 📚 Learn More

- Go crypto documentation: [crypto](https://pkg.go.dev/crypto)
- Why DES is insecure: [Data Encryption Standard](https://en.wikipedia.org/wiki/Data_Encryption_Standard)

---
