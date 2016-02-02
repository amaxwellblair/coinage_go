package transactions

import(
    // "github.com/amaxwellblair/coinage_go/keyGenerator"
    "time"
)

type Transaction struct {
    Inputs []Input
    Outputs []Output
    TimeStamp int
    Hash string
}

type Input struct {
    SourceHash string
    SourceIndex int
    Signature string
}

type Output struct {
    Amount int
    Address string
}

func CreateTime() int {
    currentTime := time.Now()
    epochTime := time.Parse(time.RFC822, "01 Jan 70 00:00 MST")
}

func (t * Transaction) StampTransaction() {

}

func New() (*Transaction) {
    return new(Transaction)
}

func (t *Transaction) CreateOutput(amount int, address string) (Output) {
    t.Outputs = append(t.Outputs, Output{amount, address})
    return t.Outputs[len(t.Outputs)-1]
}

func (t *Transaction) CreateInput(sourcehash string, sourceindex int, signature string) (Input) {
    t.Inputs = append(t.Inputs, Input{sourcehash, sourceindex, signature})
    return t.Inputs[len(t.Inputs)-1]
}

// {Amount: amount, Address: address}
