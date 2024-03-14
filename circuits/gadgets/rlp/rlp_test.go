package rlp

import (
	"fmt"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/test"
	"testing"
)

type lt struct {
	A frontend.Variable
	B frontend.Variable
}

func (c *lt) Define(api frontend.API) error {

	o := LessThan(api, c.A, c.B)
	api.AssertIsEqual(o, 1)

	//hbits := api.ToBinary(c.A, 4)[3]
	//hSubIsZero := api.IsZero(api.Sub(1, hbits))
	//lt2 := api.Sub(1, hSubIsZero)
	//api.AssertIsEqual(lt2, 1)

	return nil
}

func TestRlpLt(t *testing.T) {
	var circuit, assigment lt

	assigment.A = 7
	assigment.B = 8

	ccs, err := frontend.Compile(ecc.BLS12_377.ScalarField(), r1cs.NewBuilder, &circuit)

	asserts := test.NewAssert(t)
	asserts.NoError(err)

	fmt.Println("nb:", ccs.GetNbConstraints())
}
