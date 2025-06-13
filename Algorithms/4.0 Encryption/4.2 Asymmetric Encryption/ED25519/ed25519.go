package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// generateKeyPair generates a new Ed25519 public/private key pair.
func generateKeyPair() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return pub, priv, nil
}

// signMessage creates a signature for the message using the private key.
func signMessage(priv ed25519.PrivateKey, message []byte) []byte {
	return ed25519.Sign(priv, message)
}

// verifySignature checks the signature against the message and public key.
func verifySignature(pub ed25519.PublicKey, message, sig []byte) bool {
	return ed25519.Verify(pub, message, sig)
}

func main() {
	pub, priv, err := generateKeyPair()
	if err != nil {
		panic(err)
	}

	msg := []byte("Hello, Ed25519 in Go!")
	sig := signMessage(priv, msg)

	fmt.Printf("Message:   %s\n", msg)
	fmt.Printf("Signature: %s\n", hex.EncodeToString(sig))
	fmt.Printf("Verified?  %t\n", verifySignature(pub, msg, sig))
}
