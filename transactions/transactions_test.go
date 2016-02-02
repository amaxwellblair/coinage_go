package transactions

import(
    "testing"
    "github.com/amaxwellblair/coinage_go/keyGenerator"
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
    if time <= 0 {
        t.Fatal("Create time didn't work properly")
    }
}
