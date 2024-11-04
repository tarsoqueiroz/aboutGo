package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func main() {
	//
	// RSA Signatures: The Classic Autograph
	//
	fmt.Printf("===== RSA Signatures: The Classic Autograph\n=====\n")

	// Let's create our special signing pen (RSA key pair)
	fmt.Printf("===== RSA Signatures: Creating keys\n")
	keySize := 2048
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		panic("===== Oops, our pen ran out of ink!!!")
	}
	publicKey := &privateKey.PublicKey

	fmt.Printf("=====\n")

	// Our important message
	importantMessage := []byte("I solemnly swear that I am up to no good.")
	fmt.Printf("===== RSA Signatures: our important message: %x (%s)\n", importantMessage, importantMessage)

	// Let's create a fingerprint of our message
	messageHash := sha256.Sum256(importantMessage)

	// Time to sign!
	messageSignature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, messageHash[:])
	if err != nil {
		panic("===== Oops, our hand cramped while signing!")
	}

	fmt.Printf("===== RSA Signatures: our RSA signature    : %x\n", messageSignature)

	// Now, let's verify our signature
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, messageHash[:], messageSignature)
	if err != nil {
		fmt.Printf("===== RSA Signatures: verify RSA signature : Oops, someone forged our signature!\n")
	} else {
		fmt.Printf("===== RSA Signatures: verify RSA signature : Yeah, signature checks out. Mischief managed!\n")
	}

	fmt.Printf("=====\n\n")

	//
	// ECDSA Signatures: The Curvy Autograph
	//
	fmt.Printf("===== ECDSA Signatures: The Curvy Autograph\n=====\n")

	// Let's create our curvy signing pen
	privateEcdsaKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic("===== ECDSA Signatures: our curvy pen got a bit too curvy!")
	}
	publicEcdsaKey := &privateEcdsaKey.PublicKey

	// Our important message
	importantMessage = []byte("Elliptic curves are mathematically delicious!")
	fmt.Printf("===== ECDSA Signatures: our important message  : %x (%s)\n", importantMessage, importantMessage)

	// Create a fingerprint of our message
	messageHash = sha256.Sum256(importantMessage)

	// Time to sign with our curvy pen!
	r, s, err := ecdsa.Sign(rand.Reader, privateEcdsaKey, messageHash[:])
	if err != nil {
		panic("===== Oops, our hand slipped while signing these curves!")
	}

	ecdsaSignature := append(r.Bytes(), s.Bytes()...)
	fmt.Printf("===== ECDSA Signatures: our signature message  : %x\n", ecdsaSignature)

	// Let's verify our curvy signature
	r = new(big.Int).SetBytes(ecdsaSignature[:len(ecdsaSignature)/2])
	s = new(big.Int).SetBytes(ecdsaSignature[len(ecdsaSignature)/2:])
	valid := ecdsa.Verify(publicEcdsaKey, messageHash[:], r, s)
	fmt.Printf("===== ECDSA Signatures: Is our curvy sign valid? %v\n", valid)

	fmt.Printf("=====\n\n")

	//
	// Ed25519 Signatures: The Speed Demon Autograph
	//
	fmt.Printf("===== Ed25519 Signatures: The Speed Demon Autograph\n=====\n")

	// Let's create our speedy signing pen
	publicEd25519Key, privateEd25519Key, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic("===== Ed25519 Signatures: our speedy pen got a speeding ticket!")
	}

	// Our important message
	importantMessage = []byte("Speed is my middle name!")
	fmt.Printf("===== Ed25519 Signatures: our important message    : %x (%s)\n", importantMessage, importantMessage)

	// Time to sign at lightning speed!
	ed25519Signature := ed25519.Sign(privateEd25519Key, importantMessage)
	fmt.Printf("===== Ed25519 Signatures: our signature message    : %x\n", ed25519Signature)

	// Let's verify our speedy signature
	ed25519Valid := ed25519.Verify(publicEd25519Key, importantMessage, ed25519Signature)
	fmt.Printf("===== Ed25519 Signatures: Is our ed25519 sign valid? %v\n", ed25519Valid)

	fmt.Printf("=====\n\n")

}
