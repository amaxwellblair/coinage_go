package transactions

import(
    "testing"
    "github.com/amaxwellblair/coinage_go/keyGenerator"
)


func TestTransaction_CreateInputs_BasicExistenceCheck(t *testing.T)  {
    t := New()
    t.CreateInput(5, )
}
