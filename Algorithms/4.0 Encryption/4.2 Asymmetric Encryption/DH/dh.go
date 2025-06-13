package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Generate a random prime modulus (e.g., 256-bit)
func genPrime(bits int) (*big.Int, error) {
	return rand.Prime(rand.Reader, bits)
}

// Generate a private key in [1, p-1]
func genPrivate(p *big.Int) (*big.Int, error) {
	return rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(2)))
}

// Compute public key: g^priv mod p
func pubKey(g, priv, p *big.Int) *big.Int {
	return new(big.Int).Exp(g, priv, p)
}

// Compute shared secret: otherPub^priv mod p
func sharedSecret(otherPub, priv, p *big.Int) *big.Int {
	return new(big.Int).Exp(otherPub, priv, p)
}

func main() {
	p, _ := genPrime(256)
	g := big.NewInt(2)

	a, _ := genPrivate(p)
	A := pubKey(g, a, p)

	b, _ := genPrivate(p)
	B := pubKey(g, b, p)

	secretA := sharedSecret(B, a, p)
	secretB := sharedSecret(A, b, p)

	fmt.Printf("Shared secret equal? %t\n", secretA.Cmp(secretB) == 0)
	fmt.Printf("Secret: %x\n", secretA)
}
