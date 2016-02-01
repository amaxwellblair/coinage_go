// Package will find or generate a valid private key file
// for the given file extension

package keygen

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

type ClarkeKey struct{
    PrivateKey *rsa.PrivateKey
    PublicKey *rsa.PublicKey
    PrivatePem string
    PublicPem string
    FileExtension string
}

func NewClarkeKey() (*ClarkeKey) {
    var err error
    ck := new(ClarkeKey)
    ck.PrivatePem = GenerateKey()
    pemBlock, _ := pem.Decode([]byte(ck.PrivatePem))
    ck.PrivateKey, err = x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
    panicAtDisco(err)
    ck.PublicKey = &ck.PrivateKey.PublicKey
    ck.PublicPem = string(GeneratePublicPem(ck.PublicKey))
    return ck
}

func GenerateKey() string {
	keyFile, createFile := FindKeyFile("/Users/maxwell/.wallet/privatekey.pem")
	var pemPrivateKey []byte
	defer keyFile.Close()
	if createFile == true {
		fmt.Println("Key Generator has created a key under ~/.wallet/privatekey.pem")
		privateKey := PrivateKey()
		pemPrivateKey = GeneratePrivatePem(privateKey)
		keyFile.Write(pemPrivateKey)
	} else {
		fmt.Println("Your key has been successfully found in your wallet")
		fileInfo, _ := keyFile.Stat()
		pemPrivateKey = make([]byte, fileInfo.Size())
		keyFile.Read(pemPrivateKey)
	}
	return string(pemPrivateKey)
}

func FindKeyFile(fileExtension string) (*os.File, bool) {
	createdFile := false
	keyFile, err := os.Open(fileExtension)
	if err != nil {
		createdFile = true
		keyFile, err = os.Create(fileExtension)
		panicAtDisco(err)
	}
	return keyFile, createdFile
}

func PrivateKey() *rsa.PrivateKey {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	return privateKey
}

func GeneratePrivatePem(privateKey *rsa.PrivateKey) []byte {
	var bites []byte
	derFormat := x509.MarshalPKCS1PrivateKey(privateKey)
	pemKey := bytes.NewBuffer(bites)
	pemBlock := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: derFormat}
	err := pem.Encode(pemKey, pemBlock)
	panicAtDisco(err)
	return pemKey.Bytes()
}

func GeneratePublicPem(publicKey *rsa.PublicKey) []byte {
    var bites []byte
	derFormat, err := x509.MarshalPKIXPublicKey(publicKey)
    panicAtDisco(err)
	pemKey := bytes.NewBuffer(bites)
	pemBlock := &pem.Block{Type: "RSA PUBLICs KEY", Bytes: derFormat}
	err = pem.Encode(pemKey, pemBlock)
	panicAtDisco(err)
	return pemKey.Bytes()
}

func panicAtDisco(e error) {
    if e != nil {
        panic(e)
    }
}
