package block

import(
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/amaxwellblair/coinage_go/transactions"
)

func TestCanCreateABlock_SimpleValidationTest(t *testing.T) {
    b := New()
    assert.Equal(t, b.BlockHeader.ParentHash, "", "Block wasn't properly created")
}

func TestBlock_InsertTransaction_TestIfBlockInsertsTransactions(t *testing.T) {
    b := New()
    trans := transactions.New()
    trans.CreateTimeStamp()
    b.InsertTransaction(trans)
    if (b.Transactions[0].TimeStamp < 1400000000000){
        t.Fatal("Transaction wasn't properly inserted")
    }
}

func TestBlock_HashTransactions_CheckABlockHashesAllOfItsTransactions(t *testing.T)  {
    b := New()
    trans := transactions.New()
    b.InsertTransaction(trans)
    b.HashTransactions()
    assert.NotEqual(t, "", b.BlockHeader.TransactionsHash, "Nothing was hashed in the transcation hash: ", b.BlockHeader.TransactionsHash)
}

func TestBlock_ConcatFields_WillCheckIfTheFieldsAreConcatenatedProperly(t *testing.T) {
    b := New()
    trans := transactions.New()
    b.BlockHeader.ParentHash = "test"
    b.InsertTransaction(trans)
    b.HashTransactions()
    output := b.ConcatFields()
    assert.Equal(t, "testa7247ba502750d814253e700298ddbf205b6f17d27d3565f47d9e4fa500271ff00", output, "The concatenate function doesn't work properly: ", output)
}
