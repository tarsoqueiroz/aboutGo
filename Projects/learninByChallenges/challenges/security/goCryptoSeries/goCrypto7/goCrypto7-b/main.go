package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {
	//
	// X.509 Certificates: Your Digital Passport
	//
	fmt.Printf("===== X.509 Certificates: Creating Your Own Digital Passport (Self-Signed Certificate)\n=====\n")

	// Let's create our secret key
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic("===== Oops, our key generator is feeling shy!!!")
	}

	// Now, let's fill out our passport application
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Gopher's Cryptographic Adventures"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour * 24 * 180), // valid for 180 days

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	// Time to create our passport
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		panic("===== Oops, the passport printer is jammed!!!")
	}

	// Let's save our new passport
	certOut, err := os.Create("digital_passport.pem")
	if err != nil {
		panic("===== Oops, we can't save our new passport!!!")
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	// And let's keep our secret key safe
	keyOut, err := os.Create("digital_passport.key")
	if err != nil {
		panic("===== Oops, we can't save our new secret key!!!")
	}
	x509EncodedPrivate, _ := x509.MarshalECPrivateKey(privateKey)
	pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: x509EncodedPrivate})
	keyOut.Close()

	fmt.Printf("===== X.509 Certificates: Congratulations! You've got a new digital passport!!!\n=====\n")

}
