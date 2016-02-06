package block

import(
        "github.com/amaxwellblair/coinage_go/transactions"
        "crypto/sha256"
        // "crypto/rand"
        // "crypto"
        // "encoding/json"
        "encoding/hex"
        "reflect"
        "strconv"
)

type Block struct {
    BlockHeader Header
    Transactions []transactions.Transaction
}

type Header struct {
    ParentHash string
    TransactionsHash string
    Target string
    TimeStamp int
    Nonce int
    Hash string
}

func New() *Block {
    return &Block{BlockHeader: Header{Nonce: 0}}
}

func (b *Block) InsertTransaction(t *transactions.Transaction)  {
    b.Transactions = append(b.Transactions, *t)
}

func (b *Block) HashTransactions()  {
    var allTransBytes []byte
    var allTransHash []byte
    for _, transaction := range b.Transactions {
        allTransBytes = append(allTransHash, []byte(transaction.JsonConvert())...)
    }
    hasher := sha256.New()
    hasher.Write(allTransBytes)
    allTransHash = hasher.Sum(nil)

    b.BlockHeader.TransactionsHash = hex.EncodeToString(allTransHash)
}

func (b *Block) ConcatFields() (string) {
    var allBlockBytes []byte

    headerValue := reflect.ValueOf(b.BlockHeader)
    for i := 0; i < headerValue.NumField(); i++ {
        intermediary := headerValue.Field(i).Interface()
        switch intermediary.(type) {
        case string:
            allBlockBytes = append(allBlockBytes, []byte(intermediary.(string))...)
        case int:
            allBlockBytes = append(allBlockBytes, []byte(strconv.Itoa(intermediary.(int)))...)
        }
    }
    return string(allBlockBytes)
}

// SHA256 Hash ( previous block hash + transactions hash + timestamp + difficulty target + nonce )
