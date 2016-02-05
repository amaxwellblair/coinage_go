package transactions

import(
    "testing"
    "github.com/amaxwellblair/coinage_go/keyGenerator"
    "crypto/sha256"
    "crypto/rsa"
    "crypto/rand"
    "crypto"
    "encoding/hex"
    // "bytes"
)


func TestTransaction_CreateOutputs_BasicExistenceCheck(t *testing.T) {
    trans := New()
    ck := keygen.NewClarkeKey()

    output := trans.CreateOutput(5, ck.PublicPem)

    if output.Address != ck.PublicPem {
        t.Fatal("The output address is not returned")
    } else if output.Amount != 5 {
        t.Fatal("The output amount is not returned")
    }
}

func TestTransaction_CreateInputs_BasicExistenceCheck(t *testing.T) {
    trans := New()

    input := trans.CreateInput("sourcehash", 0, "signature")

    if input.SourceHash != "sourcehash" {
        t.Fatal("Source Hash doesn't work")
    }
    if input.SourceIndex != 0 {
        t.Fatal("Source Index doesn't work")
    }
    if input.Signature != "signature" {
        t.Fatal("Signature doesn't work")
    }
}

func TestCreateTime_MakeSureTimeCreationWorks(t *testing.T) {
    time := CreateTime()

    if time <= 1400000000000 {
        t.Fatal("Create time didn't work properly")
    }
}

func TestInput_inputInBytes_AllOftheInputsTogetherInAByteSlice(t *testing.T) {
    trans := New()
    var signature string
    input := trans.CreateInput("sourcehash", 0,  signature)

    inputBytes := input.inputInBytes()

    if string(inputBytes) != "sourcehash0" {
        t.Fatal("Input to bytes doesn't work too well: ", inputBytes)
    }
}

func TestOutput_outputInBytes_AllOftheOutputsTogetherInAByteSlice(t *testing.T) {
    trans := New()
    output := trans.CreateOutput(5, "address")

    outputBytes := output.outputInBytes()

    if string(outputBytes) != "5address" {
        t.Fatal("Output to bytes doesn't work too well: ", outputBytes)
    }
}

func TestTransaction_outputsString_AllOutputsAreInAStringTogether(t *testing.T) {
    trans := New()
    trans.CreateOutput(5, "address")
    trans.CreateOutput(10, "otheraddress")

    outputString := string(trans.outputsInBytes())
    if outputString != "5address10otheraddress" {
        t.Fatal("Taking all of the outputs and combining them into a string doesn't work: ", outputString)
    }
}

func TestTransaction_inputsString_AllInputsAreInAStringTogether(t *testing.T) {
    trans := New()
    var signature string
    trans.CreateInput("sourcehash", 0, signature)
    trans.CreateInput("otheraddress", 10, "signature")

    inputString := string(trans.inputsInBytes())
    if inputString != "sourcehash0otheraddress10signature" {
        t.Fatal("Taking all of the outputs and combining them into a string doesn't work: ", inputString)
    }
}

func TestTransaction_SignTransaction_DidThePrivateKeyProperlySign(t *testing.T) {
    trans := New()
    trans.CreateOutput(5, "address")
    trans.CreateInput("sourcehash", 0,  "signature")
    hashBytes := append(trans.inputsInBytesNoSig(), trans.outputsInBytes()...)
    shaNew := sha256.New()
    shaNew.Write(hashBytes)
    hashedIO := shaNew.Sum(nil)
    ck := keygen.NewClarkeKey()
    testSignature, err := rsa.SignPKCS1v15(rand.Reader, ck.PrivateKey, crypto.SHA256, hashedIO)
    if err != nil{
        t.Fatal("The test hashing / signing doesn't work...")
    }

    trans.SignInput(ck.PrivateKey, 0)

    if trans.Inputs[0].Signature != hex.EncodeToString(testSignature) {
        t.Fatal("The signatures don't match up: ",  trans.Inputs[0].Signature)
    }
}

func TestTransaction_HashEntireTransaction_HashesAllOfTheInputsAndOutputsAndTimeStamp(t *testing.T) {
    trans := New()
    ck := keygen.NewClarkeKey()
    trans.CreateOutput(5, "address")
    trans.CreateInput("sourcehash", 0,  "signature")
    trans.SignInput(ck.PrivateKey, 0)
    trans.CreateTimeStamp()
    trans.HashEntireTransaction()

    if trans.Hash == "" {
        t.Fatal("The full transaction has been not hashed")
    }
}

func TestJsonConvert_ConvertATransactionIntoJson (t *testing.T) {
    trans := New()
    ck := keygen.NewClarkeKey()
    trans.CreateOutput(5, "address")
    trans.CreateInput("sourcehash", 0,  "signature")
    trans.SignInput(ck.PrivateKey, 0)
    trans.CreateTimeStamp()
    trans.HashEntireTransaction()
    jsonTransaction := string(JsonConvert(trans))
    //just to read output
    if jsonTransaction == "" {
        t.Fatal("The full transaction is not in JSON: ", jsonTransaction )
    }
}
