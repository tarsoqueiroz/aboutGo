package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your Secure Tunnel TLS: Welcome to our Secure Server (HTTPS Server)")

}

func main() {
	//
	//  Your Secure Tunnel TLS: Setting Up a Secure Server (HTTPS Server)
	//
	fmt.Printf("===== Your Secure Tunnel TLS: Setting Up a Secure Server (HTTPS Server)\n=====\n")

	http.HandleFunc("/", handler)

	// let's load our digital passport and secret key
	cert, err := tls.LoadX509KeyPair("digital_passport.pem", "digital_passport.key")
	if err != nil {
		panic("===== Oops, we can't find our passport or secret key!!!")
	}

	// now, let's set up our secure tunnel
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	// time to open our secure office
	server := &http.Server{
		Addr:      ":443",
		TLSConfig: tlsConfig,
	}

	// let's start welcoming visitors!
	fmt.Printf("===== Your Secure Tunnel TLS: our secure office is open at https://localhost:10443 \n=====\n")
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		panic("===== Oops, we couldn't open our office!!!")
	}

}
