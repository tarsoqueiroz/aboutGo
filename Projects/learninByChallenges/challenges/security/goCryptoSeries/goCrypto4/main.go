package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/chacha20"
)

func main() {
	// AES
	keySize := 32
	keyAes := make([]byte, keySize)
	if _, err := rand.Read(keyAes); err != nil {
		panic("===== AES\nAES : Oops, the universe's randomness machine broke!!!")
	}

	blockAes, err := aes.NewCipher(keyAes)
	if err != nil {
		panic("===== AES\nAES : Oops, AES threw a tantrum!")
	}

	fmt.Printf("===== AES\nKey : %x\nBsiz: %d bytes\nBloc: %x\n\n", keyAes, blockAes.BlockSize(), blockAes)

	// ChaCha20
	chacha20Key := make([]byte, chacha20.KeySize)
	chacha20Nonce := make([]byte, chacha20.NonceSize)

	chacha20Cipher, err := chacha20.NewUnauthenticatedCipher(chacha20Key, chacha20Nonce)
	if err != nil {
		panic("===== chacha20\nXa20: Oops, ChaCha20 isn't feeling chatty today!!!")
	}

	chacha20SecretMessage := []byte("ChaCha20 is my new dance move!!!")
	chacha20Encrypted := make([]byte, len(chacha20SecretMessage))
	chacha20Cipher.XORKeyStream(chacha20Encrypted, chacha20SecretMessage)

	fmt.Printf("===== chacha20\nKey : %x\nNonc: %x\nMesg: %s\nEncr: %x\n\n", chacha20Key, chacha20Nonce, chacha20SecretMessage, chacha20Encrypted)

	// GCM
	// - key   = keyAes
	// - block = blockAes
	nonceGcm := make([]byte, 12)
	if _, err := rand.Read(nonceGcm); err != nil {
		panic("===== GCM\nGCM : Oops, Nonce generator is feeling noncommittal!!!")
	}
	aesGcm, err := cipher.NewGCM(blockAes)
	if err != nil {
		panic("===== GCM\nGCM : Oops, GCM mode is feeling moody!!!")
	}

	gcmSecretMessage := []byte("AES-GCM: Making encryption great again!")
	gcmEncrypted := aesGcm.Seal(nil, nonceGcm, gcmSecretMessage, nil)

	fmt.Printf("===== GCM\nKey : %x\nBloc: %x\nNonc: %x\nAES : %x\nMesg: %s\nEncr: %x\n", keyAes, blockAes, nonceGcm, aesGcm, gcmSecretMessage, gcmEncrypted)

	gcmDescrypted, err := aesGcm.Open(nil, nonceGcm, gcmEncrypted, nil)
	if err != nil {
		panic("===== GCM\nGCM : Oops, Decryption failed! Did someone tamper with our message???")
	}

	fmt.Printf("Decr: %s\n\n", gcmDescrypted)

	// CTR
	// - key   = keyAes
	// - block = blockAes
	ivAes := make([]byte, aes.BlockSize)
	if _, err := rand.Read(ivAes); err != nil {
		panic("===== CTR\nCTR : Oops, IV generator is feeling too independent!!!")
	}
	streamCTR := cipher.NewCTR(blockAes, ivAes)

	ctrSecretMessage := []byte("CTR mode: Turning blocks into streams since 1979!")
	ctrEncrypted := make([]byte, len(ctrSecretMessage))
	streamCTR.XORKeyStream(ctrEncrypted, ctrSecretMessage)

	fmt.Printf("===== CTR\nKey : %x\nBloc: %x\nIV  : %x\nStrm: %x\nMesg: %s\nEncr: %x\n", keyAes, blockAes, ivAes, streamCTR, ctrSecretMessage, ctrEncrypted)

	ctrDecrypted := make([]byte, len(ctrEncrypted))
	streamCTR = cipher.NewCTR(blockAes, ivAes)
	streamCTR.XORKeyStream(ctrDecrypted, ctrEncrypted)

	fmt.Printf("Strm: %x\nDecr: %s\n", streamCTR, ctrDecrypted)
}
