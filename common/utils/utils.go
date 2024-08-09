package utils

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/celer-network/goutils/log"
	bls12377MiMC "github.com/consensys/gnark-crypto/ecc/bls12-377/fr/mimc"
)

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln("error ("+msg+"):", err)
	}
}

func CheckOk(ok bool, msg string) {
	if !ok {
		log.Fatalln("check ok failed (" + msg + ")")
	}
}

func Convert2DimensionalByteArrayToString(input [][]byte) string {
	data := make([]string, len(input))
	for i := range data {
		data[i] = hexutil.Encode(input[i])
	}
	return strings.Join(data, ",")
}

func CheckErrf(err error, msg string, args ...interface{}) {
	if err != nil {
		log.Fatalf("error (%s): %s", fmt.Sprintf(msg, args...), err.Error())
	}
}

func CheckOkf(ok bool, msg string, args ...interface{}) {
	if !ok {
		log.Fatalf("check ok failed (%s)", fmt.Sprintf(msg, args...))
	}
}

// Simple check if a string is in array or not. Note, not perform well for a big array
func Contains(array []string, target string) bool {
	for _, s := range array {
		if s == target {
			return true
		}
	}
	return false
}

func Partition[K comparable, V any](vs []V, keys []K, f func(K, V) bool) map[K][]V {
	ret := make(map[K][]V)
	for _, v := range vs {
		for _, k := range keys {
			if f(k, v) {
				ret[k] = append(ret[k], v)
			}
		}
	}
	return ret
}

func Recompose32ByteToNibbles(api frontend.API, trunk []frontend.Variable) [64]frontend.Variable {
	var trunkBits []frontend.Variable

	var truckSize = len(trunk)
	var remainder = 256 % truckSize
	if remainder != 0 {
		panic(fmt.Sprintf("the trunk length %d not allowed", truckSize))
	}
	var trunkPieceBits = 256 / len(trunk)

	for i := 0; i < len(trunk); i++ {
		bs := api.ToBinary(trunk[truckSize-i-1], trunkPieceBits)
		trunkBits = append(trunkBits, bs...)
	}

	var nibbles [64]frontend.Variable
	for i := 0; i < 64; i++ {
		nibbles[i] = api.FromBinary(trunkBits[i*4 : (i+1)*4]...)
	}
	for i, j := 0, 64-1; i < j; i, j = i+1, j-1 {
		nibbles[i], nibbles[j] = nibbles[j], nibbles[i]
	}
	return nibbles
}

func RecomposeSDKByte32ToNibble(api frontend.API, b32 Bytes32) [64]frontend.Variable {
	h := b32.Val[0]
	l := b32.Val[1]

	var bits []frontend.Variable
	hBits := api.ToBinary(h, 248)
	lBits := api.ToBinary(l, 8)
	bits = append(bits, hBits...)

	bits = append(bits, lBits...)

	var hex [64]frontend.Variable
	for i := 0; i < len(hex); i++ {
		hex[i] = api.FromBinary(bits[i*4 : (i+1)*4]...)
	}

	for i, j := 0, 64-1; i < j; i, j = i+1, j-1 {
		hex[i], hex[j] = hex[j], hex[i]
	}

	return hex
}

func RecomposeSDKByte32To2Fr(api frontend.API, b32 Bytes32) [2]frontend.Variable {
	h := b32.Val[1]
	l := b32.Val[0]

	var bits []frontend.Variable
	hBits := api.ToBinary(h, 8)
	lBits := api.ToBinary(l, 248)
	bits = append(bits, lBits...)
	bits = append(bits, hBits...)
	api.Println(bits)
	var frs [2]frontend.Variable

	for i := 0; i < len(frs); i++ {
		frs[i] = api.FromBinary(bits[i*128 : (i+1)*128]...)
		api.Println(frs[i])
	}

	frs[0], frs[1] = frs[1], frs[0]

	return frs

}

func Recompose6BytesToNibbles(api frontend.API, d frontend.Variable) [12]frontend.Variable {
	bits := api.ToBinary(d, 48)
	var nibbles [12]frontend.Variable
	for i := 0; i < 12; i++ {
		nibbles[i] = api.FromBinary(bits[i*4 : (i+1)*4]...)
	}
	for i, j := 0, 12-1; i < j; i, j = i+1, j-1 {
		nibbles[i], nibbles[j] = nibbles[j], nibbles[i]
	}
	return nibbles
}

func CreateOrOpenFile(path string) (*os.File, error) {
	f, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		f, err = os.Create(path)
		if err != nil {
			return nil, err
		}
	} else {
		if err != nil {
			return nil, err
		}
	}
	return f, nil
}

func GenerateBls12377Hash(input []byte) []byte {
	hash := bls12377MiMC.NewMiMC()

	var bits []uint
	for _, byteVal := range input {
		bits = append(bits, DecodeByteToBinary(byteVal)...)
	}

	var bitSize = 253
	var frDataArray []*big.Int

	var i int
	for i < len(bits) {
		if i+bitSize > len(bits) {
			frDataArray = append(frDataArray, BitsToFrBigInt(bits[i:], ecc.BLS12_377.ScalarField()))
		} else {
			frDataArray = append(frDataArray, BitsToFrBigInt(bits[i:i+bitSize], ecc.BLS12_377.ScalarField()))
		}
		i += bitSize
	}

	for _, frData := range frDataArray {
		var frByteBlock = make([]byte, 32)
		frData.FillBytes(frByteBlock)
		hash.Write(frByteBlock)
	}

	return hash.Sum(nil)
}

// MiMCBlockPad0 pad 0 into miMC block up to specific block size in Big-Endian
func MiMCBlockPad0(data []byte, blockSize int) []byte {
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

const BlockExtraDataHexMaxLength = 2000

func ExportBlockHeaderExtraData(blockRlp string) ([]byte, error) {
	rlpBytes := Hex2Bytes(blockRlp)
	var decodedBlockRlp [][]byte
	err := rlp.Decode(bytes.NewReader(rlpBytes), &decodedBlockRlp)
	if err != nil {
		return nil, fmt.Errorf("failed to decode block rlp: %s", err.Error())
	}

	if len(decodedBlockRlp) < 13 {
		return nil, fmt.Errorf("block rlp decoded with unexpected length %s", blockRlp)
	}

	return decodedBlockRlp[12], nil
}
