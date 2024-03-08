package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/brevis-network/zk-utils/circuits/gadgets/keccak"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/plonk"
	cs "github.com/consensys/gnark/constraint/bn254"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/frontend/cs/scs"
	"github.com/consensys/gnark/test/unsafekzg"
)

var rounds = flag.Int("rounds", 1, "number of rounds to run")
var mode = flag.String("mode", "groth16", "the proving system to use")

// usage: go run main.go -mode=groth16 -rounds=64
func main() {
	flag.Parse()
	switch *mode {
	case "groth16":
		benchKeccakGroth16(*rounds)
	case "plonk":
		benchKeccakPlonk(*rounds)
	default:
		panic("mode not found")
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func benchKeccakPlonk(rounds int) {
	circuit := newKeccakCircuit(rounds)
	assign := newKeccakCircuit(rounds)
	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), scs.NewBuilder, circuit)
	check(err)
	r1cs := ccs.(*cs.SparseR1CS)
	canonical, lagrange, err := unsafekzg.NewSRS(r1cs)
	pk, _, err := plonk.Setup(r1cs, canonical, lagrange)
	check(err)
	w, err := frontend.NewWitness(assign, ecc.BN254.ScalarField())
	proof, err := plonk.Prove(ccs, pk, w)
	check(err)

	buf := bytes.NewBuffer(nil)
	n, err := proof.WriteRawTo(buf)
	check(err)
	fmt.Printf("proof size %d\n", n)
}

func benchKeccakGroth16(rounds int) {
	circuit := newKeccakCircuit(rounds)
	assign := newKeccakCircuit(rounds)
	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, circuit)
	check(err)

	pk, vk, err := groth16.Setup(ccs)

	w, err := frontend.NewWitness(assign, ecc.BN254.ScalarField())
	proof, err := groth16.Prove(ccs, pk, w)
	check(err)

	wp, err := w.Public()
	check(err)

	buf := bytes.NewBuffer(nil)
	n, err := proof.WriteRawTo(buf)
	check(err)
	fmt.Printf("proof size %d\n", n)

	err = groth16.Verify(proof, vk, wp)
}

func getTestData(rounds int) []frontend.Variable {
	data := make([]frontend.Variable, rounds*1088)
	for i := range data {
		data[i] = 0
	}
	return data
}

type KeccakCircuit struct {
	Data   []frontend.Variable
	rounds int
}

func newKeccakCircuit(rounds int) *KeccakCircuit {
	return &KeccakCircuit{
		Data:   getTestData(rounds),
		rounds: rounds,
	}
}

func (c *KeccakCircuit) Define(api frontend.API) error {
	keccak.Keccak256Bits(api, c.rounds, c.rounds-1, c.Data)
	return nil
}
