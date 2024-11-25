package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/hkdf"
)

func main() {
	//
	// HKDF: The Key Factory
	//
	fmt.Printf("===== HKDF: The Key Factory\n=====\n")

	secret := []byte("MySuperSecretValue")
	salt := []byte("SaltySalt")
	info := []byte("KeyForEncryption")

	fmt.Printf("===== HKDF: secret      : %s\n", secret)
	fmt.Printf("===== HKDF: salt        : %s\n", salt)
	fmt.Printf("===== HKDF: info        : %s\n", info)

	// let's start up our key factory
	keyFactory := hkdf.New(sha256.New, secret, salt, info)

	// now, let's produce two 32-byte keys
	key1 := make([]byte, 32)
	key2 := make([]byte, 32)

	if _, err := io.ReadFull(keyFactory, key1); err != nil {
		panic("===== Oops, our key factory 1 had a malfunction!!!")
	}
	if _, err := io.ReadFull(keyFactory, key2); err != nil {
		panic("===== Oops, our key factory 21 had a malfunction!!!")
	}

	fmt.Printf("===== HKDF: key1        : %x\n", key1)
	fmt.Printf("===== HKDF: key2        : %x\n", key2)

	// let's encode our new keys
	encodedKey1 := base64.RawStdEncoding.EncodeToString(key1)
	encodedKey2 := base64.RawStdEncoding.EncodeToString(key2)

	fmt.Printf("===== HKDF: encoded key1: %s\n", encodedKey1)
	fmt.Printf("===== HKDF: encoded key2: %s\n", encodedKey2)
}
