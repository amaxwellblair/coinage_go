package main

import (
    "fmt"
    "crypto/rsa"
    "crypto/rand"
    "crypto"
    "crypto/sha256"
    "crypto/x509"
    "encoding/pem"
    "bytes"
)

func main() {
    privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
    publicKey := privateKey.Public().(*rsa.PublicKey)

    snippet := []byte("Hello world!")

    basicPubEncrypt, _ := rsa.EncryptPKCS1v15(rand.Reader, publicKey, snippet)
    fmt.Println("RSA encryption using a public key: ", string(basicPubEncrypt))

    basicPrivateDecrypt, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, basicPubEncrypt)
    fmt.Println("RSA decryption using a private key: ", string(basicPrivateDecrypt))

    //Now we will explore how to hash things
    //you must write the snippet to the hash before you can hash the phrase
    hashSHA := sha256.New()
    hashSHA.Write(snippet)
    hashedMessage := hashSHA.Sum(nil)
    fmt.Println("SHA256 hash of the message 'Hello world!': ", string(hashedMessage))

    //Now we will explore how to sign things
    hashOfChoice := crypto.SHA256
    basicPrivateSign, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, hashOfChoice, hashedMessage)

    fmt.Println("RSA signature: ", basicPrivateSign)

    basicPublicVerify := rsa.VerifyPKCS1v15(publicKey, hashOfChoice, hashedMessage, basicPrivateSign)
    fmt.Println("RSA verify nil = true: ", basicPublicVerify)

    // First the PKCS1 key is placed in x509 ASN.1 DER format then given to the PEM encoding function
    // helpful resource https://golang.org/src/crypto/tls/generate_cert.go?m=text
    derFormat := x509.MarshalPKCS1PrivateKey(privateKey)
    privatePem := make([]byte, 3000)
    pemPrivateKey := bytes.NewBuffer(privatePem)
    pemBlock := &pem.Block{Type: "RSA Private Key", Bytes: derFormat}
    pem.Encode(pemPrivateKey, pemBlock)
    fmt.Println("RSA private key to PEM: ", pemPrivateKey)

}
