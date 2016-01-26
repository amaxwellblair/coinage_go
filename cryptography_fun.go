package main

import (
    "crypto/rsa"
    "crypto/rand"
    "crypto/sha256"
)

func main() {
    privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
    publicKey := privateKey.Public
}
