# Cryptographic Building Blocks: The LEGO of Security, Go Crypto 3

> [`https://dev.to/rezmoss/cryptographic-building-blocks-the-lego-of-security-go-crypto-3-4l4k`](https://dev.to/rezmoss/cryptographic-building-blocks-the-lego-of-security-go-crypto-3-4l4k)

## Hash Functions: Your Digital Fingerprint Maker

First up, we've got hash functions. These are like those machines that squish pennies into oval shapes at tourist attractions. No matter what you put in, you always get a fixed-size output. But unlike those penny machines, a good hash function creates a unique "fingerprint" for every input.

- **SHA-256 and SHA-224:** The workhorses of the crypto world. Here's how you'd use SHA-256
- **SHA-512 and friends:** When you need extra security (or just like big numbers)
- **SHA-3:** The new kid on the block, resistant to some theoretical attacks on SHA-2

## Message Authentication Codes (MACs): Your Digital Seal

They don't keep the contents secret, but they do prove who sent the message and that it hasn't been tampered with.

- **HMAC (Hash-based Message Authentication Code):** The Swiss Army knife of MACs.
- **CMAC (Cipher-based Message Authentication Code):** Not in the standard library, but available if you need it.

## Random Number Generation: Your Digital Dice

It's crucial for generating keys, nonces, and anywhere else you need unpredictability.

> **Remember,** always use `crypto/rand`, not `math/rand`. Using `math/rand` for crypto is like using a toy safe to store your valuables!

## Best Practices: The Golden Rules of Crypto

Here are some golden rules to keep in mind:

- **Stick with the strong stuff:** For hashing, SHA-256 or better. It's like choosing a good lock - go for the best you can afford.
- **Keep your secrets secret:** MAC keys are like the key to your house. Don't leave them lying around!
- **Garbage in, garbage out:** When generating keys or other secret material, use high-quality random input. No birthdays or "password123" allowed!
- **Check your work:** Always check for and handle errors, especially with random number generation. It's like double-checking that you locked the door.
- **Know your limits:** Be aware of performance implications. Crypto operations can be heavy, especially on large data or in high-traffic scenarios. Profile your code!

## Coding

```sh
# initiating code
go mod init gocrypto3
# program
touch main.go
# importing modules
go get crypto/sha256
go get crypto/sha512
go get golang.org/x/crypto/sha3
go get crypto/hmac
go get crypto/aes
go get golang.org/x/crypto/cmac
go get crypto/rand
go get encoding/binary
# running code
go run main.go 
```

## That's all

...folks!!!
