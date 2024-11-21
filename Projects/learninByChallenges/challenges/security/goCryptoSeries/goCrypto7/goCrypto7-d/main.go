package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	//
	//  Your Secure Tunnel TLS: Creating a Secure Client
	//
	fmt.Printf("===== Your Secure Tunnel TLS: Creating a Secure Client\n=====\n")

	// let's load the passport of the server we want to visit
	certPool := x509.NewCertPool()
	pem, err := os.ReadFile("digital_passport.pem")
	if err != nil {
		panic("===== Oops, we can't find the server's passport!!!")
	}
	if !certPool.AppendCertsFromPEM(pem) {
		panic("===== Oops, this doesn't look like a valid passport!!!")
	}

	// now, let's prepare for our secure journey
	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}

	// time to create our secure transport
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// let's visit the secure server!
	resp, err := client.Get("https://10.15.18.44")
	// resp, err := client.Get("https://google.com")
	if err != nil {
		panic("===== Oops, our secure journey failed!!!")
	}
	defer resp.Body.Close()

	// What did the server say?
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("===== Oops, we couldn't understand the server's message!!!")
	}
	fmt.Printf("===== Your Secure Tunnel TLS: The server says: %s\n=====\n", body)

}
