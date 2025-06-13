# Go ECDSA Signature

This repository demonstrates ECDSA-based digital signing and verification using pure Go (standard library only).

## 🔐 Overview

ECDSA (Elliptic Curve Digital Signature Algorithm) allows you to:

1. Generate an EC key pair (`elliptic.P256`)
2. Sign messages with a private key using `SignASN1`
3. Verify signatures using the public key with `VerifyASN1`

This is **not** public-key encryption—ECDSA provides authentication via signatures, not confidentiality.

Refer to Go’s `crypto/ecdsa` package for full documentation.

## 📥 Getting Started

### Requirements

- Go 1.18 or higher

## 🎯 What the Code Does

1. **Key Pair Generation**
   Uses `ecdsa.GenerateKey(elliptic.P256(), rand.Reader)` to create a 256-bit EC key pair.

2. **Signing**

   - Hashes the message using SHA-256
   - Signs with `ecdsa.SignASN1()` producing a DER-encoded signature

3. **Verifying**

   - Recomputes SHA-256 of the message
   - Validates signature with `ecdsa.VerifyASN1()`

## 📌 Code Overview

| Function          | Responsibility                 |
| ----------------- | ------------------------------ |
| `generateKeyPair` | Create ECDSA private key       |
| `signMessage`     | Hash and sign a message        |
| `verifySignature` | Confirm signature authenticity |

## 🧠 Why Use ECDSA?

- **Security**: Comparable to RSA-3072 but with smaller key sizes
- **Efficiency**: Faster operations due to the use of elliptic curves.

## ⚠️ Notes & Security

- Use SHA-256 or stronger hashes
- Signatures are **non-deterministic** (each signature uses fresh randomness)
- Keep the private key confidential—exposure leads to signature forgery

## 📚 References

- Official Go [crypto/ecdsa](https://pkg.go.dev/crypto/ecdsa)
