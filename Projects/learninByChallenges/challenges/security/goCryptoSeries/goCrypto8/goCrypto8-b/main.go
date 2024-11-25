package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

func main() {
	//
	// Argon2: The Newer, Fancier Smoothie
	//
	fmt.Printf("===== Argon2: The Newer, Fancier Smoothie\n=====\n")

	password := []byte("iLoveCryptoEvenMore456")

	// first, let's add some salt to our smoothie
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		panic("===== Oops, our salt shaker is empty!!!")
	}

	// now, let's blend our password
	timeCost := uint32(1)
	memoryCost := uint32(64 * 1024)
	threads := uint8(4)
	keyLength := uint32(32)

	hash := argon2.IDKey(password, salt, timeCost, memoryCost, threads, keyLength)

	// let's encode our smoothie and salt for storage
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)

	fmt.Printf("===== Argon2: our fancy password smoothie: %s\n", encodedHash)
	fmt.Printf("===== Argon2: our salt                   : %s\n", encodedSalt)

	// to verify, we'd need to decode the salt, reblend with the same recipe, and compare
}
