package transactions

import (
    "testing"
    "os"
    "crypto/rsa"
    "crypto/x509"
    "crypto/rand"
    "encoding/pem"
    // "crypto/rsa"
)

func TestFindKeyFile_CheckReturnsAFile(t *testing.T)  {
    fakefile := "bacon.md"
    defer os.Remove(fakefile)

    f := FindKeyFile(fakefile)
    if f == nil {
        t.Fatal("No file was actually created or found")
    }
}

func TestGeneratePrivateKey_ValidatingACorrectKeyIsBeingReturned(t *testing.T)  {
    privateKey := GeneratePrivateKey()
    output := privateKey.Validate()
    if output != nil {
        t.Fatal("The private key created was invalid")
    }
}

func TestGeneratePem_GeneratePemThenCheckIfValidPrivateKey(t *testing.T) {
    privateKey := GeneratePrivateKey()
    privatePem := GeneratePrivatePem(privateKey)
    pemBlock, _ := pem.Decode(privatePem)
    if pemBlock == nil {
        t.Fatal("The pem key didn't transfer back into a block")
    }
    validKey, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
    if err != nil {
        t.Fatal("The parser failed to create a key")
    }
    output := validKey.Validate()
    if output != nil {
        t.Fatal("The key that was created is not valid")
    }

    snippet := "Hello, world!"
    publicKey := privateKey.PublicKey
    publicEncrypt, _ := rsa.EncryptPKCS1v15(rand.Reader, &publicKey, []byte(snippet))

    privateDecrypt, err := rsa.DecryptPKCS1v15(rand.Reader, validKey, publicEncrypt)
    if err != nil {
        t.Fatal("Something went wrong in the decryption of publicEncrypt: ", err)
    }

    if string(privateDecrypt) != snippet {
        t.Fatal("Something didn't work in: ", privateDecrypt)
    }
}




// Test<TypeName>_<MethodName>_<ShortDescription>()
