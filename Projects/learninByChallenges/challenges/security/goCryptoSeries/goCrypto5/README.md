# Public-Key Cryptography: The Digital Handshake, Go Crypto 5

> [`https://dev.to/rezmoss/public-key-cryptography-the-digital-handshake-go-crypto-5-167b`](https://dev.to/rezmoss/public-key-cryptography-the-digital-handshake-go-crypto-5-167b)

## RSA Key Generation: Your Digital Identity

Let's start by creating our RSA keys.

## RSA Encryption and Decryption: Passing Secret Notes

Now, let's use these keys to send a secret message.

## RSA Signing and Verification: Your Digital Signature

RSA isn't just for secret messages. It can also create digital signatures.

## Elliptic Curve Cryptography (ECC): The New Kid on the Block

Now, let's talk about ECC. It's like RSA's cooler, more efficient cousin. It offers similar security with smaller keys, which is great for mobile and IoT devices.

## ECDSA Key Generation: Your Elliptic Identity

Let's create our ECC keys.

## ECDSA Signing and Verification: Curvy Signatures

Now, let's sign something with our elliptic keys.

## Key Management: Keeping Your Digital Identity Safe

Now, let's talk about keeping these keys safe. It's like having a really important key to a really important door - you want to keep it secure!

## The Golden Rules of Public-Key Cryptography

Now that you're wielding these powerful crypto tools, here are some golden rules to keep in mind:

- **Size matters:** For RSA, go big or go home - at least 2048 bits. For ECC, 256 bits is the sweet spot.
- **Randomness is your friend:** Always use crypto/rand for key generation. Using weak randomness is like using "password123" as your key.
- **Rotate your keys:** Like changing your passwords, rotate your keys regularly.
- **Standard formats are standard for a reason:** Use PEM for storing and sending keys. It's like using a standard envelope - everyone knows how to handle it.
- **Padding is not just for furniture:** For RSA encryption, always use OAEP padding. It's like bubble wrap for your encrypted data.
- **Hash before you sign:** When signing large data, sign the hash, not the data itself. It's faster and just as secure.
- **Performance matters:** Public-key operations can be slow, especially RSA. Use them wisely.

## Coding

```sh
# initiating code
go mod init gocrypto5
# program
touch main.go
# importing modules
go get 
# running code
go run main.go 
```

## That's all

...folks!!!
