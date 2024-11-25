package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

func main() {
	//
	// PBKDF2: The Classic Key Maker
	//
	fmt.Printf("===== PBKDF2: The Classic Key Maker\n=====\n")

	password := []byte("OpenSesaMe123")

	// let's add some randomness to our key-making process
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		panic("===== Oops, our randomness generator broke!!!")
	}

	// time to make our key
	iterations := 100000
	keyLength := 32
	key := pbkdf2.Key(password, salt, iterations, keyLength, sha256.New)

	// let's encode our new key and salt
	encodedKey := base64.RawStdEncoding.EncodeToString(key)
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)

	fmt.Printf("===== PBKDF2: our shiny new key: %s\n", encodedKey)
	fmt.Printf("===== PBKDF2: the salt we used : %s\n", encodedSalt)

}
