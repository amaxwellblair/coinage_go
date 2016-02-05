package transactions

import(
    "time"
    // "fmt"
    "reflect"
    "strconv"
    "errors"
    "crypto/rsa"
    "crypto/sha256"
    "crypto/rand"
    "crypto"
    "encoding/json"
    "encoding/hex"
)

type Transaction struct {
    Inputs []Input
    Outputs []Output
    TimeStamp string
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

func JsonConvert(trans interface{}) []byte {
    jsonTransaction, err := json.Marshal(trans)
    if err != nil {
        panic(err)
    }
    return jsonTransaction
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

func hashThing(hashBytes []byte) ([]byte) {
    shaNew := sha256.New()
    shaNew.Write(hashBytes)
    return shaNew.Sum(nil)
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

func (i *Input) inputInBytesNoSig() ([]byte) {
    var inputBytes []byte
    value := reflect.ValueOf(*i)
    for i := 0; i < value.NumField()-1; i++ {
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

func (t *Transaction) outputsInBytes() ([]byte)  {
    var allBytes []byte
    for _, value := range t.Outputs {
        allBytes = append(allBytes, value.outputInBytes()...)
    }
    return allBytes
}

func (t *Transaction) inputsInBytes() ([]byte) {
    var allBytes []byte
    for _, value := range t.Inputs {
        allBytes = append(allBytes, value.inputInBytes()...)
    }
    return allBytes
}

func (t *Transaction) inputsInBytesNoSig() ([]byte) {
    var allBytes []byte
    for _, value := range t.Inputs {
        allBytes = append(allBytes, value.inputInBytesNoSig()...)
    }
    return allBytes
}

func New() (*Transaction) {
    return new(Transaction)
}

func (t *Transaction) SignInput(privKey *rsa.PrivateKey, input_index int) {
    hashBytes := append(t.inputsInBytesNoSig(), t.outputsInBytes()...)
    hashedIO := hashThing(hashBytes)
    signature, err := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, hashedIO)
    if err != nil {
        panic(err)
    }
    t.Inputs[input_index].Signature = hex.EncodeToString(signature)
}

func (t *Transaction) CreateOutput(amount int, address string) (Output) {
    t.Outputs = append(t.Outputs, Output{amount, address})
    return t.Outputs[len(t.Outputs)-1]
}

func (t *Transaction) CreateInput(sourcehash string, sourceindex int, signature string) (Input) {
    t.Inputs = append(t.Inputs, Input{sourcehash, sourceindex, signature})
    return t.Inputs[len(t.Inputs)-1]
}

func (t * Transaction) CreateTimeStamp() {
    t.TimeStamp = strconv.Itoa(int(CreateTime()))
}

func (t * Transaction) HashEntireTransaction() {
    byteTime := []byte(t.TimeStamp)
    transactionByte := append(t.inputsInBytes(), append(t.outputsInBytes(), byteTime...)...)
    t.Hash = hex.EncodeToString(hashThing(transactionByte))
}
