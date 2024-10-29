package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func main() {
	//
	// RSA: The Granddaddy of Public-Key Crypto
	//
	fmt.Printf("===== Rivest-Shamir-Adleman (RSA) Public-Key Crypto\n=====\n")
	// Creating RSA keys
	fmt.Printf("===== RSA : Creating keys\n")

	// let's make a {keySize}-bit key (it's like choosing a really long pass)
	keySize := 2048
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		panic("===== Oops, our key generator is feeling shy today!!!")
	}

	publicKey := &privateKey.PublicKey

	fmt.Printf("-->> Key size: %d <<--\n", keySize)
	fmt.Printf("-->> Private key <<--\n%v\n", privateKey)
	fmt.Printf("-->> Public key<<--\n%v\n", publicKey)
	fmt.Printf("=====\n\n")

	// Passing secret notes
	fmt.Printf("===== Encryption & Decryption\n")

	secretMessage := []byte("RSA is like a magic envelope!")

	// encryption: sealing our magic envelope
	cipherText, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		secretMessage,
		nil,
	)
	if err != nil {
		panic("===== Oops, our magic envelope got stuck!!!")
	}

	fmt.Printf("-->> Message <<--\n%s\n", secretMessage)
	fmt.Printf("-->> Encrypted Message <<--\n%x\n", cipherText)

	// decryption: opening our magic envelope
	plainText, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey,
		cipherText,
		nil,
	)
	if err != nil {
		panic("===== Oops, we can't pém our own magic envelope!!!")
	}

	fmt.Printf("-->> Decrypted Message <<--\n%s\n", plainText)
	fmt.Printf("=====\n\n")

	// your Digital Signature
	fmt.Printf("===== Signing & Verification\n")

	myOpenMessage := []byte("I solemnly swear that I am up to no good.")
	hashMyOpenMessage := sha256.Sum256(myOpenMessage)

	// signing: like singing a digital contract
	mySignature, err := rsa.SignPKCS1v15(
		rand.Reader,
		privateKey,
		crypto.SHA256,
		hashMyOpenMessage[:],
	)
	if err != nil {
		panic("===== Oops, our digital pen ran out of ink!!!")
	}

	fmt.Printf("-->> My Open Message <<--\n%s\n", myOpenMessage)
	fmt.Printf("-->> My Signature <<--\n%x\n", mySignature)

	// verification: checking if the signature is genuine
	err = rsa.VerifyPKCS1v15(
		publicKey,
		crypto.SHA256,
		hashMyOpenMessage[:],
		mySignature,
	)
	if err != nil {
		fmt.Printf("===== Oops, this signature looks fishy!!!\n\n")
	} else {
		fmt.Printf("===== Uh-oh, signature checks out. Mischief managed!!!\n\n")
	}

	//
	// Elliptic Curve Cryptography (ECC): The New Kid on the Block
	//
	fmt.Printf("===== Elliptic Curve Cryptography (ECC)\n=====\n")
	// Creating ECC keys
	fmt.Printf("===== ECC : Creating keys\n")

	privEccKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic("===== Oops, our elliptic curve generator took a wrong turn!!!")
	}

	pubEccKey := &privEccKey.PublicKey

	fmt.Printf("-->> Private key <<--\n%v\n", privEccKey)
	fmt.Printf("-->> Public key<<--\n%v\n", pubEccKey)
	fmt.Printf("=====\n\n")

	//  ECDSA Signing & Verification
	fmt.Printf("===== Signing & Verification\n")

	myEllipticMessage := []byte("Elliptic curves are mathematically delicious!")
	hashMyEllipticMessage := sha256.Sum256(myEllipticMessage)

	// signing: like singing with a very curvy pen
	rEcc, sEcc, err := ecdsa.Sign(rand.Reader, privEccKey, hashMyEllipticMessage[:])
	if err != nil {
		panic("===== Oops, our curvy signature got a bit too curvy!!!")
	}

	fmt.Printf("-->> My Elliptic Message <<--\n%s\n", myEllipticMessage)
	fmt.Printf("-->> My Elliptic Signature <<--\nr = %x\ns = %x\n", rEcc, sEcc)

	// verification: checking if our curvy signature is legit
	validEcc := ecdsa.Verify(pubEccKey, hashMyEllipticMessage[:], rEcc, sEcc)

	fmt.Printf("-->> Is my elliptic signature valid? %v\n", validEcc)

	//
	// Key Management: Keeping Your Digital Identity Safe
	//
	fmt.Printf("===== Key Management\n=====\n")
	// encoding private key
	fmt.Printf("===== Encodig private key\n")

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privateKeyBytes,
		})

	fmt.Printf("-->> Key in its special envelop <<--\n%s\n", privateKeyPEM)

	// decoding private key
	blockDecoded, _ := pem.Decode(privateKeyPEM)
	decodedPrivateKey, err := x509.ParsePKCS1PrivateKey(blockDecoded.Bytes)
	if err != nil {
		panic("===== Oops, we forgot how to open our own envelope!!!")
	}

	fmt.Printf("-->> Decoded Private key <<--\n%v\n", decodedPrivateKey)
	fmt.Printf("-->> Private key <<--\n%v\n", privateKey)
}
