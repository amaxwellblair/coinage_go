package transactions

import(
    "testing"
    "github.com/amaxwellblair/coinage_go/keyGenerator"
    // "crypto/sha256"
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
    input := trans.CreateInput("sourcehash", 0, "signature")

    inputBytes := input.inputInBytes()

    if string(inputBytes) != "sourcehash0signature" {
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

func TestTransaction_StampTransaction_MakeSureTheTransactionIsHashedAndTimeStamped(t *testing.T) {
    // trans := New()
    // trans.CreateInput("sourcehash", 0, "signature")
    // trans.CreateOutput(5, "Publickey")
    // trans.StampTransaction()
    // if trans.TimeStamp <= 1400000000000 {
    //     t.Fatal("TimeStamp wasn't properly set")
    // }
    //
    // inputs := trans.Inputs[0].SourceHash + trans.Inputs[0].SourceIndex trans.Input[0].Signature
    // outputs := trans.Outputs[0].Amount + trans.Outputs[0].Address + trans.TimeStamp
    //
    // hashString := inputs + outputs
    // shaNew := sha256.New()
    //
    // var buffer []byte
    // shaNew.Write(buffer)
    //
    // if trans.Hash == "hash"{
    //     t.Fatal("The hash isn't working properly")
    // }
}
