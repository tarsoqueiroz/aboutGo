# Cracking Open Go's Crypto Toolbox, Go Crypto 2

- [`https://dev.to/rezmoss/cracking-open-gos-crypto-toolbox-go-crypto-2-50cn`](https://dev.to/rezmoss/cracking-open-gos-crypto-toolbox-go-crypto-2-50cn)

## The Master Plan: Go's Crypto Philosophy

A set of guiding principles that make this package a joy to use:

- **Keep It Simple, Smarty (KISS):** They designed the API to be so straightforward that you don't need a Ph.D. in cryptography to use it. It's like the LEGO of crypto - simple blocks that you can easily put together.
- **Safety First:** The package is like that friend who always reminds you to wear your seatbelt. It's got your back, implementing secure defaults and trying its best to stop you from making those facepalm-worthy crypto mistakes.
- **Speed Demon:** Nobody likes waiting around, especially in the digital age. That's why the crypto package is optimized for speed, with many operations coded in assembly language for different architectures. It's like having a sports car engine in your crypto toolkit.
- **Play Well With Others:** While it comes with a ton of built-in goodies, the package is designed to play nice with custom implementations. It's like a potluck dinner - bring your own crypto dish if you want!
- **By the Book:** The implementations in this package follow the rules. They adhere to widely accepted cryptographic standards and best practices. It's like having a strict but fair referee in your code.

## What's in the Box? The Structure of Go's Crypto Package

- **crypto:** This is the main package, the outer doll if you will. It's got the common crypto constants and interfaces that the other packages use.
- **crypto/aes:** Need to keep secrets? This package implements the AES encryption algorithm, perfect for symmetric encryption.
- **crypto/cipher:** This is your Swiss Army knife for encryption. It's got block cipher modes, AEAD ciphers, and stream ciphers.
- **crypto/ecdsa and crypto/ed25519:** These are your go-to packages for digital signatures. ECDSA is like the classic rock of digital signatures, while Ed25519 is the new kid on the block.
- **crypto/elliptic:** This package deals with elliptic curves. It's like the geometry class of cryptography.
- **crypto/hmac:** Want to make sure your message hasn't been tampered with? HMAC's got your back.
- **crypto/md5:** The old-timer of hash functions. But remember, it's here for compatibility, not for new projects!
- **crypto/rand:** This is your cryptographic dice roller. When you need random numbers that are really, really random, this is your guy.
- **crypto/rc4:** Another oldie but goodie. It's a stream cipher, but like MD5, it's not recommended for new systems.
- **crypto/rsa:** The granddaddy of public-key cryptography. Great for both encryption and digital signatures.
- **crypto/sha1, crypto/sha256, crypto/sha512:** The SHA family of hash functions. They're like siblings - similar, but each with their own strengths.
- **crypto/subtle:** This package is all about timing attack prevention. It's like a ninja, working in the shadows to keep your operations secure.
- **crypto/tls:** Implementing secure connections? This package has got you covered with TLS 1.2 and 1.3 support.
- **crypto/x509:** Dealing with certificates? This package helps you navigate the world of X.509 public key infrastructure.
