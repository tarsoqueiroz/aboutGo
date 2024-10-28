# Symmetric Encryption: The Secret Handshake of Cryptography, Go Crypto 4

> [`https://dev.to/rezmoss/symmetric-encryption-the-secret-handshake-of-cryptography-go-crypto-4-1d17`](https://dev.to/rezmoss/symmetric-encryption-the-secret-handshake-of-cryptography-go-crypto-4-1d17)

## AES: The Heavyweight Champion

AES is versatile, strong, and widely used.

## Stream Ciphers: The Flowing River of Encryption

Next, we have stream ciphers. These are like a never-ending stream of random-looking bits that we XOR with our data to encrypt it. Go gives us ChaCha20, a modern, speedy stream cipher.

## ChaCha20: The New Kid on the Block

ChaCha20 is great when you need speed, especially on platforms without AES hardware acceleration.

## Modes of Operation: Putting It All Together

Now, let's talk about modes of operation. These are like the rules of a game - they define how we use our ciphers to encrypt data securely.

## GCM (Galois/Counter Mode): The Swiss Army Knife

GCM provides both secrecy and integrity, which is why it's highly recommended for most use cases.

## CTR (Counter Mode): The Streamifier

CTR is useful when you need the flexibility of a stream cipher but want to stick with a block cipher algorithm.

## The Golden Rules of Symmetric Encryption

Now that you've got these shiny new encryption tools, here are some golden rules to keep in mind:

- **GCM is your friend:** For most cases, use AES-GCM. It's like a bodyguard for your data - it protects both the secrecy and the integrity.
- **Nonce is the spice of life:** Always use a unique nonce (number used once) for each encryption operation. It's like a unique identifier for each secret message.
- **Randomness is key:** Generate your keys using crypto/rand. Using weak keys is like using "password123" for your bank account.
- **CTR needs a buddy:** If you're using CTR mode, remember it doesn't protect integrity. Consider pairing it with a MAC if you need integrity protection.
- **Error handling is not optional:** Always handle errors, especially during key generation and initialization. Ignoring errors in crypto code is like ignoring the "Check Engine" light on your car.
- **Keep your secrets secret:** Never, ever hard-code keys in your source code. It's like hiding your house key under the welcome mat - the first place an attacker will look!

## Coding

```sh
# initiating code
go mod init gocrypto4
# program
touch main.go
# importing modules
go get crypto/aes
go get crypto/rand
# running code
go run main.go 
```

## That's all

...folks!!!
