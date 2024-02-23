package smt

import (
	"fmt"
	"math/big"

	"github.com/brevis-network/zk-utils/common/circuits"
	"github.com/brevis-network/zk-utils/common/utils"
	"github.com/consensys/gnark-crypto/ecc/bls12-377/fr"
	"github.com/consensys/gnark-crypto/ecc/bls12-377/fr/mimc"
	"github.com/consensys/gnark/frontend"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

type Hash [2]frontend.Variable

func EmptyHash() Hash {
	return Hash{0, 0}
}

func NewHash(bs []byte) [2]frontend.Variable {
	if len(bs) != 32 {
		panic(fmt.Errorf("invalid input len: want 32, got %d", len(bs)))
	}
	h := new(big.Int).SetBytes(bs[:])
	s := new(big.Int).Lsh(big.NewInt(1), 128)
	quo, rem := new(big.Int).QuoRem(h, s, new(big.Int))
	return [2]frontend.Variable{quo, rem}
}

func EncodeHeaders(headers []types.Header) (encoded [][]frontend.Variable, idxs []frontend.Variable, err error) {
	for _, header := range headers {
		ok := checkDynamicFieldLen(header)
		if !ok {
			return nil, nil, fmt.Errorf("dynamic field len check failed %+v", header)
		}
		headerRLP, encodeErr := rlp.EncodeToBytes(&header)
		if encodeErr != nil {
			err = encodeErr
			return
		}
		padded := circuits.Pad101Bytes(headerRLP)
		paddedBits := circuits.Bytes2BlockBits(padded)
		var fv []frontend.Variable
		for _, b := range paddedBits {
			fv = append(fv, b)
		}
		idxs = append(idxs, len(fv)/1088-1)
		zerosToPad := 5*1088 - len(paddedBits)
		for i := 0; i < zerosToPad; i++ {
			fv = append(fv, 0)
		}
		encoded = append(encoded, fv)
	}
	return
}

func checkDynamicFieldLen(h types.Header) bool {
	hasErr := 0
	hasErr += checkLen("Difficulty", len(h.Difficulty.Bytes()), 7)
	hasErr += checkLen("Number", len(h.Number.Bytes()), 8)
	hasErr += checkLen("GasLimit", getUintByteLen(h.GasLimit), 4)
	hasErr += checkLen("GasUsed", getUintByteLen(h.GasUsed), 4)
	hasErr += checkLen("Time", getUintByteLen(h.Time), 4)
	hasErr += checkLen("BaseFee", len(h.BaseFee.Bytes()), 7)
	return hasErr == 0
}

func getUintByteLen(x uint64) int {
	i := 0
	for x > 0 {
		x /= 256
		i++
	}
	return i
}

func checkLen(field string, actual, max int) int {
	if actual > max {
		fmt.Println(fmt.Errorf("field %s len actual %d > max %d", field, actual, max))
		return 1
	}
	return 0
}

func GenHeaders(count int, startBlockNum int, prevHash [32]byte) []types.Header {
	var hs []types.Header
	hs = append(hs, genHeader(prevHash, big.NewInt(int64(startBlockNum))))
	for i := 1; i < count; i++ {
		num := i + startBlockNum
		h := genHeader(hs[i-1].Hash(), big.NewInt(int64(num)))
		hs = append(hs, h)
	}
	return hs
}

func genHeader(parentHash [32]byte, number *big.Int) types.Header {
	n7 := new(big.Int).SetBytes([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	emptyHash := [32]byte{}
	return types.Header{
		ParentHash:      parentHash,
		UncleHash:       emptyHash,
		Coinbase:        [20]byte{},
		Root:            emptyHash,
		TxHash:          emptyHash,
		ReceiptHash:     emptyHash,
		Bloom:           [256]byte{},
		Difficulty:      n7,
		Number:          number,
		GasLimit:        0xffffffff,
		GasUsed:         0xffffffff,
		Time:            0xffffffff,
		Extra:           []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		MixDigest:       emptyHash,
		Nonce:           [8]byte{1, 1, 1, 1, 1, 1, 1, 1},
		BaseFee:         n7,
		WithdrawalsHash: (*common.Hash)(&emptyHash),
	}
}

func MimcRoot(leaves [][]byte) []byte {
	length := len(leaves)
	if length == 0 {
		panic("nil leaves")
	}
	if length&(length-1) != 0 {
		panic("length must be power of 2")
	}
	if length == 1 {
		return leaves[0]
	}
	h := mimc.NewMiMC()
	newLeaves := make([][]byte, length/2)
	for i := 0; i < len(leaves); i += 2 {
		h.Reset()
		_, err := h.Write(leaves[i])
		utils.CheckErrf(err, "error writing leaf to hasher")
		_, err = h.Write(leaves[i+1])
		utils.CheckErrf(err, "error writing leaf to hasher")
		hash := h.Sum(nil)
		newLeaves[i/2] = hash
	}
	return MimcRoot(newLeaves)
}

func MimcHashBytes32(in []byte) []byte {
	h := mimc.NewMiMC()
	el0 := new(fr.Element).SetBytes(padLeft(in[:16], 16)).Bytes()
	el1 := new(fr.Element).SetBytes(padLeft(in[16:], 16)).Bytes()
	h.Write(el0[:])
	h.Write(el1[:])
	return h.Sum(nil)
}

func padLeftVars(in []frontend.Variable, n int) []frontend.Variable {
	ret := make([]frontend.Variable, len(in)+n)
	for i := 0; i < n; i++ {
		ret[i] = 0
	}
	copy(ret[n:], in)
	return ret
}

func padLeft[T any](in []T, n int) []T {
	ret := make([]T, len(in)+n)
	copy(ret[n:], in)
	return ret
}
