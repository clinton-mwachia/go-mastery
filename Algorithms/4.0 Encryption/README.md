# 🔐 Golang Encryption Algorithms

This repository provides an overview of various encryption algorithms you can implement using Go (Golang). It covers both built-in capabilities and third-party packages, and is organized by category for easy navigation.

---

## 🔒 Symmetric Encryption (Secret Key)

Symmetric encryption uses the same key for both encryption and decryption.

| Algorithm                           | Description                             | Package                                |
| ----------------------------------- | --------------------------------------- | -------------------------------------- |
| **AES (AES-128, AES-192, AES-256)** | Industry standard block cipher          | `crypto/aes`                           |
| **DES**                             | Outdated block cipher (not recommended) | `crypto/des`                           |
| **3DES (Triple DES)**               | More secure than DES but still outdated | `crypto/des`                           |
| **ChaCha20**                        | Modern stream cipher, fast and secure   | `golang.org/x/crypto/chacha20`         |
| **ChaCha20-Poly1305**               | AEAD cipher with authentication         | `golang.org/x/crypto/chacha20poly1305` |
| **AES-GCM**                         | AES with Galois/Counter Mode (AEAD)     | `crypto/cipher`                        |
| **AES-CFB, AES-OFB, AES-CTR**       | AES in various stream/block modes       | `crypto/cipher`                        |

---

## 🔑 Asymmetric Encryption (Public/Private Key)

Used for secure key exchange, encryption, and digital signatures.

| Algorithm               | Description                            | Package                          |
| ----------------------- | -------------------------------------- | -------------------------------- |
| **RSA**                 | Widely used for encryption and signing | `crypto/rsa`                     |
| **ECDSA**               | Elliptic curve digital signatures      | `crypto/ecdsa`                   |
| **Ed25519**             | High-performance digital signatures    | `crypto/ed25519`                 |
| **Diffie-Hellman (DH)** | Key exchange protocol                  | Custom / third-party             |
| **ECDH**                | ECC-based DH key exchange              | `golang.org/x/crypto/curve25519` |

---

## 🔐 Key Derivation Functions (KDF)

Used to derive cryptographic keys from passwords.

| Algorithm  | Description                              | Package                      |
| ---------- | ---------------------------------------- | ---------------------------- |
| **PBKDF2** | Password-Based Key Derivation Function 2 | `golang.org/x/crypto/pbkdf2` |
| **bcrypt** | Secure password hashing                  | `golang.org/x/crypto/bcrypt` |
| **scrypt** | Memory-hard password hashing             | `golang.org/x/crypto/scrypt` |
| **Argon2** | Winner of password hashing competition   | `golang.org/x/crypto/argon2` |

---

## 📜 Hash Functions and HMAC

Hashing is used for data integrity and digital signatures.

| Algorithm                    | Description                        | Package                          |
| ---------------------------- | ---------------------------------- | -------------------------------- |
| **MD5**                      | Insecure for cryptographic use     | `crypto/md5`                     |
| **SHA-1**                    | Weak, avoid for secure apps        | `crypto/sha1`                    |
| **SHA-2 (SHA-256, SHA-512)** | Secure and widely used             | `crypto/sha256`, `crypto/sha512` |
| **SHA-3**                    | Modern hashing standard (FIPS 202) | `golang.org/x/crypto/sha3`       |
| **BLAKE2b / BLAKE2s**        | Fast and secure alternative        | `golang.org/x/crypto/blake2b`    |
| **HMAC**                     | Keyed message authentication code  | `crypto/hmac`                    |

---

## 🧱 AEAD (Authenticated Encryption with Associated Data)

Combines encryption and authentication in a single operation.

| Algorithm              | Description                        | Package                                |
| ---------------------- | ---------------------------------- | -------------------------------------- |
| **AES-GCM**            | Authenticated encryption using AES | `crypto/cipher`                        |
| **ChaCha20-Poly1305**  | Secure and fast AEAD cipher        | `golang.org/x/crypto/chacha20poly1305` |
| **XChaCha20-Poly1305** | Extended nonce version             | `golang.org/x/crypto/chacha20poly1305` |

---

## 📦 Advanced / Other Cryptographic Tools

| Algorithm / Tool              | Description                  | Package                             |
| ----------------------------- | ---------------------------- | ----------------------------------- |
| **OpenPGP**                   | PGP encryption (deprecated)  | `golang.org/x/crypto/openpgp`       |
| **Noise Protocol**            | Secure messaging protocol    | `github.com/flynn/noise`            |
| **TLS/SSL**                   | Secure network communication | `crypto/tls`                        |
| **Libsodium (NaCl bindings)** | Modern cryptographic library | `github.com/GoKillers/libsodium-go` |

---

## ✅ Recommendations

| Use Case             | Recommended Algorithms         |
| -------------------- | ------------------------------ |
| General encryption   | AES-GCM, ChaCha20-Poly1305     |
| Password storage     | Argon2id, bcrypt               |
| Digital signatures   | Ed25519                        |
| Secure communication | TLS, Noise Protocol, Libsodium |
