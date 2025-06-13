package main

import (
	"crypto/ecdh"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	curve := ecdh.P256()

	aliceKey, _ := curve.GenerateKey(rand.Reader)
	bobKey, _ := curve.GenerateKey(rand.Reader)

	alicePub := aliceKey.PublicKey()
	bobPub := bobKey.PublicKey()

	secretA, _ := aliceKey.ECDH(bobPub)
	secretB, _ := bobKey.ECDH(alicePub)

	fmt.Printf("Shared secrets equal? %t\n",
		sha256.Sum256(secretA) == sha256.Sum256(secretB))
	fmt.Printf("Shared secret SHA‑256: %x\n", sha256.Sum256(secretA))
}
