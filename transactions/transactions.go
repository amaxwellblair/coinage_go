package transactions

import (
    // "fmt"
    "os"
    "crypto/rsa"
    "crypto/rand"
    "crypto/x509"
    "encoding/pem"
    "bytes"
)

func FindKeyFile(fileExtension string)(*os.File) {
    keyFile, err := os.Open(fileExtension)
    if err != nil {
        keyFile, err = os.Create(fileExtension)
        if err != nil {
            fileErr(err)
        }
    }
    return keyFile
}


func GeneratePrivateKey() (*rsa.PrivateKey) {
    privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
    return privateKey
}

func GeneratePrivatePem(privateKey *rsa.PrivateKey) ([]byte) {
    var bites []byte
    derFormat := x509.MarshalPKCS1PrivateKey(privateKey)
    pemKey := bytes.NewBuffer(bites)
    pemBlock := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: derFormat}
    err := pem.Encode(pemKey, pemBlock)
    if err != nil {
        fileErr(err)
    }
    return pemKey.Bytes()
}

func fileErr(e error)  {
    panic(e)
}
