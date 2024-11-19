package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	//
	// X.509 Certificates: Your Digital Passport
	//
	fmt.Printf("===== X.509 Certificates: Your Digital Passport\n=====\n")

	// Let's read our digital passport
	certPEM, err := os.ReadFile("digital_passport.pem")
	if err != nil {
		panic("===== Oops, we lost our passport!!!")
	}

	// Decode the PEM block (it's like opening the passport)
	block, _ := pem.Decode(certPEM)
	if block == nil {
		panic("===== Oops, this doesn't look like a passport!!!")
	}

	// Parse the certificate (reading the passport details)
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic("===== Oops, we can't read this passport!!!")
	}

	// Let's see what's in our passport
	fmt.Printf("===== X.509 Certificates: Passport owner : %s\n", cert.Subject)
	fmt.Printf("===== X.509 Certificates: Passport issuer: %s\n", cert.Issuer)
	fmt.Printf("===== X.509 Certificates: Valid from     : %s\n", cert.NotBefore)
	fmt.Printf("===== X.509 Certificates: Valid until    : %s\n", cert.NotAfter)

}
