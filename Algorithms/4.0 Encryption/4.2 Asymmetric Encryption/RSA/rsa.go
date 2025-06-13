package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// generateKeys produces a new RSA 2048-bit key pair.
func generateKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return priv, &priv.PublicKey, nil
}

// exportPrivateKey encodes a private key to PEM format.
func exportPrivateKey(priv *rsa.PrivateKey) []byte {
	privDER := x509.MarshalPKCS1PrivateKey(priv)
	return pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privDER})
}

// exportPublicKey encodes a public key to PEM format.
func exportPublicKey(pub *rsa.PublicKey) ([]byte, error) {
	pubDER, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return nil, err
	}
	return pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pubDER}), nil
}

// encryptOAEP encrypts plaintext with the recipient's public key.
func encryptOAEP(pub *rsa.PublicKey, msg []byte) ([]byte, error) {
	label := []byte("") // optional context
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, msg, label)
}

// decryptOAEP decrypts ciphertext using the private key.
func decryptOAEP(priv *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	label := []byte("")
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, ciphertext, label)
}

func main() {
	priv, pub, err := generateKeys()
	if err != nil {
		panic(err)
	}

	fmt.Println("Private Key PEM:\n", string(exportPrivateKey(priv)))
	pubPEM, _ := exportPublicKey(pub)
	fmt.Println("Public Key PEM:\n", string(pubPEM))

	plaintext := []byte("Hello, RSA with Go!")
	ciphertext, err := encryptOAEP(pub, plaintext)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	decrypted, err := decryptOAEP(priv, ciphertext)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}
