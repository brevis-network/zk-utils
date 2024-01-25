package smt

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"testing"

	"github.com/celer-network/goutils/log"
	"github.com/consensys/gnark-crypto/ecc"
	bls12377MiMC "github.com/consensys/gnark-crypto/ecc/bls12-377/fr/mimc"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
	"github.com/consensys/gnark/test"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/consensys/gnark-crypto/accumulator/merkletree"
)

type MerkleRootWithMiMCHashCircuit struct {
	LeafValueTree    []frontend.Variable
	LeafValueIndexes []frontend.Variable
	Root             frontend.Variable
}

func (c *MerkleRootWithMiMCHashCircuit) Define(api frontend.API) error {
	miMCHash, err := mimc.NewMiMC(api)
	if err != nil {
		return err
	}

	out := CalculateLeafValueHash(api, miMCHash, c.LeafValueTree, c.LeafValueIndexes)
	api.AssertIsEqual(out, c.Root)
	return nil
}

func TestLeafValueWithMiMCHash(t *testing.T) {
	assert := test.NewAssert(t)

	miMCHash := bls12377MiMC.NewMiMC()

	index := hexutil.MustDecode("0xa0")
	leafValue := hexutil.MustDecode("0x123456")

	indexInBits := convertByteToBits(index)
	bytes := []byte{0, 1, 2, 3, 4}
	branches := make([][]byte, len(indexInBits))

	for i := range indexInBits {
		randomLength := rand.Int() % 32
		branches[i] = make([]byte, randomLength)
		for j := 0; j < randomLength; j++ {
			branches[i][j] = bytes[rand.Int()%5]
		}
	}
	value := leafValue

	for i, index := range indexInBits {
		miMCHash.Reset()
		if index == 0 {
			miMCHash.Write(miMCBlockPad0(value, miMCHash.BlockSize()))
			miMCHash.Write(miMCBlockPad0(branches[i], miMCHash.BlockSize()))
		} else {
			miMCHash.Write(miMCBlockPad0(branches[i], miMCHash.BlockSize()))
			miMCHash.Write(miMCBlockPad0(value, miMCHash.BlockSize()))
		}
		value = miMCHash.Sum(nil)
	}

	log.Infof("Hash Result: %s\n", hexutil.Encode(value))

	var circuit, assignment MerkleRootWithMiMCHashCircuit
	indexFV := make([]frontend.Variable, len(indexInBits))

	for i := range indexInBits {
		indexFV[i] = indexInBits[i]
	}

	branchesFV := make([]frontend.Variable, len(indexInBits))

	for i := range indexInBits {
		branchesFV[i] = branches[i]
	}

	leafValueTree := make([]frontend.Variable, len(indexInBits)+1)
	leafValueTree[0] = leafValue
	for i := range indexInBits {
		leafValueTree[i+1] = branches[i]
	}

	assignment = MerkleRootWithMiMCHashCircuit{
		LeafValueTree:    leafValueTree,
		LeafValueIndexes: indexFV,
		Root:             value,
	}

	circuit = MerkleRootWithMiMCHashCircuit{
		LeafValueTree:    make([]frontend.Variable, len(indexInBits)+1),
		LeafValueIndexes: make([]frontend.Variable, len(indexInBits)),
		Root:             value,
	}
	err := test.IsSolved(&circuit, &assignment, ecc.BLS12_377.ScalarField())
	assert.NoError(err)

}

func convertByteToBits(bytes []byte) []uint8 {
	var result = make([]uint8, len(bytes)*8)

	for i, b := range bytes {
		result[i*8] = (b & 128) >> 7
		result[i*8+1] = (b & 64) >> 6
		result[i*8+2] = (b & 32) >> 5
		result[i*8+3] = (b & 16) >> 4
		result[i*8+4] = (b & 8) >> 3
		result[i*8+5] = (b & 4) >> 2
		result[i*8+6] = (b & 2) >> 1
		result[i*8+7] = b & 1
	}

	return result
}

func miMCBlockPad0(data []byte, blockSize int) []byte {
	var block = make([]byte, blockSize)
	for i := 0; i < blockSize; i++ {
		if i < blockSize-len(data) {
			block[i] = 0
		} else {
			block[i] = data[len(data)-blockSize+i]
		}
	}
	return block
}

func TestSMTProof(t *testing.T) {
	assert := test.NewAssert(t)

	hash := bls12377MiMC.NewMiMC()
	smt := merkletree.New(hash)
	err := smt.SetIndex(0)

	if err != nil {
		assert.NoError(err)
	}

	smtDepth := 2
	nbLeaves := 1 << smtDepth

	for i := 0; i < nbLeaves; i++ {
		randomValue := make([]byte, 32)
		for j := 0; j < 32; j++ {
			randomValue[j] = 0 // byte() rand.U % 256
		}
		smt.Push(randomValue)
	}
	smtRoot, proof, index, _ := smt.Prove()
	fmt.Printf("SMT root: %s\n", hexutil.Encode(smtRoot))
	for i, v := range proof {
		fmt.Printf("Proof%d: %s\n", i, hexutil.Encode(v))
	}
	fmt.Printf("SMT index: %d\n", index)

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, index)
	indexInBits := convertByteToBits(b)

	fmt.Printf("SMT index in bits: %v\n", indexInBits)

	var circuit, assignment MerkleRootWithMiMCHashCircuit
	indexFV := make([]frontend.Variable, len(indexInBits))

	for i := range indexInBits {
		indexFV[i] = indexInBits[i]
	}

	branchesFV := make([]frontend.Variable, len(proof))

	for i := range proof {
		branchesFV[i] = proof[i]
	}

	zero32 := [32]byte{}

	leafValueTree := make([]frontend.Variable, len(proof)+1)
	leafValueTree[0] = zero32
	for i := range proof {
		leafValueTree[i+1] = proof[i]
	}

	assignment = MerkleRootWithMiMCHashCircuit{
		LeafValueTree:    leafValueTree,
		LeafValueIndexes: indexFV,
		Root:             smtRoot,
	}

	fmt.Printf("Assignment: %v", assignment)

	circuit = MerkleRootWithMiMCHashCircuit{
		LeafValueTree:    make([]frontend.Variable, len(proof)+1),
		LeafValueIndexes: make([]frontend.Variable, len(proof)),
		Root:             0,
	}
	err = test.IsSolved(&circuit, &assignment, ecc.BLS12_377.ScalarField())
	assert.NoError(err)

	assert.Equal(1, 2)
}
