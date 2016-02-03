package transactions

import(
    // "github.com/amaxwellblair/coinage_go/keyGenerator"
    "time"
    // "fmt"
    "reflect"
    "strconv"
    "errors"
)

type Transaction struct {
    Inputs []Input
    Outputs []Output
    TimeStamp float64
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

func CreateTime() float64 {
    epochTime, err := time.Parse(time.RFC822, "01 Jan 70 00:00 MST")
    if err != nil {
        panic(err)
    }
    passedTime := time.Since(epochTime)
    totalSeconds := passedTime.Seconds()
    return totalSeconds * 1000
}

func (t *Transaction) StampTransaction() {

}

func (t *Transaction) fieldsInBytes() {

}

func (i *Input) inputInBytes() ([]byte) {
    var inputBytes []byte
    value := reflect.ValueOf(*i)
    for i := 0; i < value.NumField(); i++ {
        intermediary := value.Field(i).Interface()
        switch intermediary.(type) {
        case string:
            inputBytes = append(inputBytes, []byte(intermediary.(string))...)
        case int:
            intermediateString := strconv.Itoa(intermediary.(int))
            inputBytes = append(inputBytes, []byte(intermediateString)...)
        default:
            panic(errors.New("Encountered an input type that wasn't recognized"))
        }
    }
    return inputBytes
}

func (o *Output) outputInBytes() ([]byte) {
    var outputBytes []byte
    value := reflect.ValueOf(*o)
    for o := 0; o < value.NumField(); o++ {
        intermediary := value.Field(o).Interface()
        switch intermediary.(type) {
        case string:
            outputBytes = append(outputBytes, []byte(intermediary.(string))...)
        case int:
            intermediateString := strconv.Itoa(intermediary.(int))
            outputBytes = append(outputBytes, []byte(intermediateString)...)
        default:
            panic(errors.New("Encountered an output type that wasn't recognized"))
        }
    }
    return outputBytes
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
