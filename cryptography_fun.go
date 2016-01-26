package main

import (
    "fmt"
    "crypto/rsa"
    "crypto/rand"
    // "reflect"
    // "crypto/sha256"
)

func main() {
    // Had some issues with relfection and pointers - figured it out though!
    // publicKey := new(rsa.PublicKey)
    // fakeKey := privateKey.Public()
    // fakeType := reflect.TypeOf(fakeKey)
    // fakeValue := reflect.ValueOf(fakeKey)
    // fmt.Println("Reflection of fake key type: ", fakeType)
    // fmt.Println("Reflection of fake key value:", fakeValue)
    privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
    publicKey := privateKey.Public().(*rsa.PublicKey)

    basicPubEncrypt, _ := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte("Hello world!"))
    fmt.Println("RSA encryption using a public key: ", string(basicPubEncrypt))

    basicPrivateDecrypt, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, basicPubEncrypt)
    fmt.Println("RSA decryption using a private key: ", string(basicPrivateDecrypt))



}
