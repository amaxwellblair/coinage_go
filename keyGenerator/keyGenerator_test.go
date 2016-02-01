package keygen

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

    f, createFile := FindKeyFile(fakefile)
    if f == nil {
        t.Fatal("No file was actually created or found")
    }
    if createFile != true {
        t.Fatal("File was created yet it didn't say one was")
    }
}

func TestPrivateKey_ValidatingACorrectKeyIsBeingReturned(t *testing.T)  {
    privateKey := PrivateKey()
    output := privateKey.Validate()
    if output != nil {
        t.Fatal("The private key created was invalid")
    }
}

func TestGeneratePem_GeneratePemThenCheckIfValidPrivateKey(t *testing.T) {
    privateKey := PrivateKey()
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

func TestGeneratePublicPem_CheckIfPemIsValid(t *testing.T)  {
    privateKey := PrivateKey()
    publicKey := privateKey.PublicKey
    publicPem := GeneratePublicPem(&publicKey)
    pemBlock, _ := pem.Decode(publicPem)
    if pemBlock == nil {
        t.Fatal("The pem key didn't transfer back into a block")
    }
}

func TestGenerateKey_GeneratesAKeyInTheWalletDirectoryOrFindsOne(t *testing.T)  {
    privatePem := GenerateKey()
    pemBlock, _ := pem.Decode([]byte(privatePem))
    if pemBlock == nil {
        t.Fatal("The pem key read from the wallet didn't transfer back into a block")
    }
    _, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
    if err != nil {
        t.Fatal("The key that was read in from the wallet is invalid")
    }
}

func TestCreateClarkeKey_ShouldCreateAStructWithAllofTheFieldsFilled(t *testing.T) {
    ck := NewClarkeKey()
    pemBlock, _ := pem.Decode([]byte(ck.PrivatePem))
    if pemBlock == nil {
        t.Fatal("The Pem key uploaded into the Clarke Key Struct is invalid")
    }
    err := ck.PrivateKey.Validate()
    if err != nil {
        t.Fatal("No private key was created in the Clarke Key Struct")
    }

}


// Test<TypeName>_<MethodName>_<ShortDescription>()
