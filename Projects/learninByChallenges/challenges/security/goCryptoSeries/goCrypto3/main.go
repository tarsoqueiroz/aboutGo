package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"

	"fmt"

	"golang.org/x/crypto/sha3"
)

func main() {
	// SHA256
	sha256Data := []byte("Go crypto rocks!")
	sha256Hash := sha256.Sum256(sha256Data)

	fmt.Printf("===== sha256\nData: %s\nHash: %x\n\n", sha256Data, sha256Hash)

	// SHA512
	sha512Data := []byte("Go crypto rocks!")
	sha512Hash := sha512.Sum512(sha512Data)

	fmt.Printf("===== sha512\nData: %s\nHash: %x\n\n", sha512Data, sha512Hash)

	// SHA3 256
	sha3256Data := []byte("Go crypto rocks!")
	sha3256Hash := sha3.Sum256(sha3256Data)

	fmt.Printf("===== sha3.256\nData: %s\nHash: %x\n\n", sha3256Data, sha3256Hash)

	// SHA3 512
	sha3512Data := []byte("Go crypto rocks!")
	sha3512Hash := sha3.Sum512(sha3512Data)

	fmt.Printf("===== sha3.512\nData: %s\nHash: %x\n\n", sha3512Data, sha3512Hash)

	// HMAC
	hmacKey := []byte("super-secret-key")
	hmacMessage := []byte("The dog nhac-nhac my pernation")
	hmacMac := hmac.New(sha256.New, hmacKey)
	hmacMac.Write(hmacMessage)
	hmacSignature := hmacMac.Sum(nil)

	fmt.Printf("===== hmac\nKey : %s\nMesg: %s\nHMAC: %X\nSign: %x\n\n", hmacKey, hmacMessage, hmacMac, hmacSignature)

	// CMAC
	// cmacKey := []byte("16-byte-long-key")
	// cmacMessage := []byte("The dog nhac-nhac my pernation")
	// cmacCipher, _ := aes.NewCipher(cmacKey)
	// cmacMac, _ := cmac.New(cmacCipher)
	// cmacMac.Write(cmacMessage)
	// cmacSignature := cmacMac.Sum(nil)

	// fmt.Printf("===== cmac\nKey : %s\nMesg: %s\nCMAC: %X\nSign: %x\n\n", cmacKey, cmacMessage, cmacMac, cmacSignature)
	fmt.Printf("===== cmac\nNOT IMPLEMENTED:  go: module golang.org/x/crypto@upgrade found (v0.28.0), but does not contain package golang.org/x/crypto/cmac\n\n")

	// Roll a 32-bit die
	var randNumber int32
	binary.Read(rand.Reader, binary.BigEndian, &randNumber)

	fmt.Printf("===== rand\nNumb: %d\n\n", randNumber)

	// Generate nSize random bytes
	nSize := 16
	randBytes := make([]byte, nSize)
	_, err := rand.Read(randBytes)
	if err != nil {
		panic("===== rand\nRand: Oops, the universe broke!!!")
	}

	fmt.Printf("===== rand\nByte: %x\n\n", randBytes)
}
