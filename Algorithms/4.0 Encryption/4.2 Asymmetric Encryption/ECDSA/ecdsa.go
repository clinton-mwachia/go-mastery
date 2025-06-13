package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// generateKeyPair generates an ECDSA P‑256 private/public key pair.
func generateKeyPair() (*ecdsa.PrivateKey, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	return priv, nil
}

// signMessage hashes and signs the message using ECDSA (ASN.1 DER format).
func signMessage(priv *ecdsa.PrivateKey, message []byte) ([]byte, error) {
	hash := sha256.Sum256(message)
	return ecdsa.SignASN1(rand.Reader, priv, hash[:])
}

// verifySignature checks the signed message against the public key.
func verifySignature(pub *ecdsa.PublicKey, message, sig []byte) bool {
	hash := sha256.Sum256(message)
	return ecdsa.VerifyASN1(pub, hash[:], sig)
}

func main() {
	// 1. Generate ECDSA key pair
	priv, err := generateKeyPair()
	if err != nil {
		panic(err)
	}
	pub := &priv.PublicKey

	// 2. Sign a message
	msg := []byte("Hello, ECDSA in Go!")
	sig, err := signMessage(priv, msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Message: %s\nSignature: %s\n", msg, hex.EncodeToString(sig))

	// 3. Verify the signature
	valid := verifySignature(pub, msg, sig)
	fmt.Printf("Signature valid? %t\n", valid)
}
