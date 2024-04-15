package keccak

import (
	"fmt"
	"github.com/brevis-network/zk-utils/common/utils"
	"github.com/celer-network/goutils/log"
	"golang.org/x/crypto/sha3"
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/test"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestKeccakBits(t *testing.T) {
	maxRounds := 2
	// 1 round
	roundIndex := 0
	data, _ := hexutil.Decode("0xff00000000000000000000000000000000000000000000000000000000000010ff")
	hash, _ := hexutil.Decode("0x746cc57064795780b008312042c24f949ad9dc0ee2dce9f4828f5a8869ccecca")

	// 2 rounds
	// roundIndex := 1
	// data, _ := hexutil.Decode("0xff00000000000000000000000000000000000000000000000000000000000010ff00000000000000000000000000000000000000000000000000000000000010ff00000000000000000000000000000000000000000000000000000000000010dceb0ddf468b489ddb3ea6a3ef6ec613df11711daeb7d7d390d1148f95054df8dceb0ddf468b489ddb3ea6a3ef6ec613df11711daeb7d7d390d1148f95054df8")
	// hash, _ := hexutil.Decode("0x7f55bf028f17dea0c32680c64fd54365e4ded6f3eecec3f31a214e0a5d4025be")

	padded := Pad101Bytes(data)
	paddedBits := Bytes2BlockBits(padded)
	out := Bytes2Bits(hash)
	if len(out) != 256 {
		panic(fmt.Sprintf("out len %d", len(out)))
	}
	var out256 [256]frontend.Variable
	for i, v := range out {
		out256[i] = frontend.Variable(v)
	}
	var dataBits []frontend.Variable
	// convert int array to frontend.Variable array
	for _, b := range paddedBits {
		dataBits = append(dataBits, b)
	}
	// fill the rest with 0s
	zerosToPad := maxRounds*1088 - len(dataBits)
	for i := 0; i < zerosToPad; i++ {
		dataBits = append(dataBits, 0)
	}
	w := &Keccak256BitsCircuit{
		MaxRounds:  maxRounds,
		RoundIndex: roundIndex,
		Out:        out256,
		Data:       dataBits,
	}
	circuit := &Keccak256BitsCircuit{
		MaxRounds:  maxRounds,
		RoundIndex: roundIndex,
		Out:        out256,
		Data:       dataBits,
	}
	err := test.IsSolved(circuit, w, ecc.BN254.ScalarField())
	check(err)

	cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, circuit)
	check(err)
	fmt.Println("constraints", cs.GetNbConstraints())
}

type Keccak256BitsCircuit struct {
	MaxRounds  int
	Data       []frontend.Variable
	RoundIndex frontend.Variable      `gnark:",public"`
	Out        [256]frontend.Variable `gnark:",public"`
}

func (c *Keccak256BitsCircuit) Define(api frontend.API) error {
	out := Keccak256Bits(api, c.MaxRounds, c.RoundIndex, c.Data)
	for i := 0; i < 256; i++ {
		api.AssertIsEqual(out[i], c.Out[i])
	}
	return nil
}

type Keccak256BitsInCircuit struct {
	Preimage frontend.Variable
	ImageBE  [2]frontend.Variable
}

func (c *Keccak256BitsInCircuit) Define(api frontend.API) error {
	preimageBitsLE := api.ToBinary(c.Preimage, 64)             // pure LE: aka bytes are LE, bits in bytes are LE
	preimageBitsMixed := utils.FlipByGroups(preimageBitsLE, 8) // mixed endianness: bytes are BE, bits in bytes are LE

	padded := PadBits101(api, preimageBitsMixed, 1)

	out := Keccak256Bits(api, 1, 0, padded)      // mixed endianness
	imageBitsLE := utils.FlipByGroups(out[:], 8) // pure LE

	// recompose to two limbs with the pure LE bits
	lo := api.FromBinary(imageBitsLE[:128]...)
	hi := api.FromBinary(imageBitsLE[128:]...)

	api.AssertIsEqual(c.ImageBE[0], hi)
	api.AssertIsEqual(c.ImageBE[1], lo)

	log.Infof("image %x%x\n", hi, lo)
	return nil
}

func TestKeccakBitsInCircuit(t *testing.T) {
	in, ok := new(big.Int).SetString("74133B7E75160A80", 16)
	if !ok {
		panic("failed to parse string")
	}

	hash := sha3.NewLegacyKeccak256()
	hash.Write(in.Bytes())
	expectedOut := new(big.Int).SetBytes(hash.Sum(nil))

	// decompose the expected output hash into two limbs
	one := big.NewInt(1)
	hi, lo := new(big.Int).DivMod(expectedOut, new(big.Int).Lsh(one, 128), new(big.Int))

	w := &Keccak256BitsInCircuit{
		Preimage: in,
		ImageBE:  [2]frontend.Variable{hi, lo},
	}
	err := test.IsSolved(&Keccak256BitsInCircuit{}, w, ecc.BN254.ScalarField())
	check(err)
}
