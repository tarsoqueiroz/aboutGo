# Password Hashing and Key Derivation: Turning Secrets into Secure Gibberish, Go Crypto 8

> [`https://dev.to/rezmoss/password-hashing-and-key-derivation-turning-secrets-into-secure-gibberish-go-crypto-8-4gn`](https://dev.to/rezmoss/password-hashing-and-key-derivation-turning-secrets-into-secure-gibberish-go-crypto-8-4gn)

## Password Hashing: Making Passwords Unreadable (Even to Us!)

First up, let's talk about password hashing. It's like putting passwords through a cryptographic blender - what comes out looks nothing like what went in, and that's exactly what we want!

### Bcrypt: The Classic Password Smoothie

Bcrypt is like the classic smoothie of password hashing - tried, tested, and still delicious. Here's how to use it:

```sh
go get golang.org/x/crypto/bcrypt
```

> [`main.go`](./goCrypto8-a/main.go)

### Argon2: The Newer, Fancier Smoothie

Argon2 is like the newfangled smoothie with all the superfoods - it's designed to be extra resistant to modern password-cracking techniques. Here's how to use it:

```sh
go get golang.org/x/crypto/argon2
```

> [`main.go`](./goCrypto8-b/main.go)

## Key Derivation: Turning Passwords into Crypto Keys

Now, let's talk about key derivation. It's like turning a simple password into a complex key that can unlock our cryptographic treasures.

### PBKDF2: The Classic Key Maker

PBKDF2 is like an old, reliable key-cutting machine. It takes your password and turns it into a shiny new key. Here's how:

```sh
go get golang.org/x/crypto/pbkdf2
```

> [`main.go`](./goCrypto8-c/main.go)

### HKDF: The Key Factory

HKDF is like a magical key factory that can produce multiple keys from a single secret. It's perfect when you need several keys for different purposes. Here's how to use it:

```sh
go get golang.org/x/crypto/hkdf
```

> [`main.go`](./goCrypto8-d/main.go)

## The Golden Rules of Password Hashing and Key Derivation

Now that you're a master of turning secrets into secure gibberish, here are some golden rules to keep in mind:

1. **Use the right tool for the job:** For passwords, use bcrypt or Argon2. For key derivation, use PBKDF2 or HKDF.
1. **Salt to taste:** Always use a unique, random salt for each password or key. It's like adding a secret ingredient that makes each hash unique.
1. **Adjust your recipe:** Choose appropriate work factors (iterations, memory cost) based on your security needs and hardware capabilities. It's like adjusting the cooking time and temperature.
1. **Keep your recipe secret:** Securely generate and store your salts and other parameters. Don't let anyone peek at your secret ingredients!
1. **Never serve raw:** Never store plain text passwords or encryption keys. Always serve them well-hashed or derived.
1. **Timing is everything:** Use constant-time comparison functions when verifying passwords. It's like making sure you always take the same time to check a password, whether it's right or wrong.
1. **Keep up with the trends:** Regularly review and update your chosen algorithms and parameters. Cryptography is like fashion - what's secure today might not be tomorrow!

## That's all

...folks!!!
