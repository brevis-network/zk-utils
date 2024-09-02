package test

import (
	poseidoncircuit "github.com/brevis-network/zk-utils/circuits/gadgets/poseidon"
	"github.com/brevis-network/zk-utils/common/utils"
	"github.com/celer-network/goutils/log"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/scs"
	"github.com/consensys/gnark/test"
	"math/big"
	"testing"
)

func TestPoseidonHash(t *testing.T) {
	assert := test.NewAssert(t)
	poseidonHasher := utils.NewPoseidonBn254()

	var preimage []frontend.Variable

	for i := 0; i < 8; i++ {
		poseidonHasher.Write(new(big.Int).SetUint64(uint64(i)))
		preimage = append(preimage, new(big.Int).SetUint64(uint64(i)))
	}
	currentPoseidonHash, err := poseidonHasher.Sum()
	assert.NoError(err)
	log.Infof("currentPoseidonHash: %x", currentPoseidonHash)

	circuit := &PoseidonTestCircuit{
		Preimage: preimage,
		HashRes:  currentPoseidonHash,
	}
	assigment := &PoseidonTestCircuit{
		Preimage: preimage,
		HashRes:  currentPoseidonHash,
	}

	err = test.IsSolved(circuit, assigment, ecc.BN254.ScalarField())
	assert.NoError(err)

	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), scs.NewBuilder, circuit)
	assert.NoError(err)
	log.Infof("ccs: %d", ccs.GetNbConstraints())

}

type PoseidonTestCircuit struct {
	Preimage []frontend.Variable
	HashRes  frontend.Variable
}

func (c *PoseidonTestCircuit) Define(api frontend.API) error {
	poseidon, err := poseidoncircuit.NewBn254PoseidonCircuit(api)
	if err != nil {
		return err
	}
	for i := 0; i < len(c.Preimage); i++ {
		poseidon.Write(c.Preimage[i])
	}
	currentPoseidonHash := poseidon.Sum()
	api.AssertIsEqual(currentPoseidonHash, c.HashRes)

	poseidon.Reset()

	for i := 0; i < len(c.Preimage); i++ {
		poseidon.Write(c.Preimage[i])
	}
	currentPoseidonHash = poseidon.Sum()
	api.AssertIsEqual(currentPoseidonHash, c.HashRes)

	return nil
}
