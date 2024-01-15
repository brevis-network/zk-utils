package utils

import (
	"math/big"

	"github.com/consensys/gnark/frontend"
)

type Bytes32 struct {
	Val [2]frontend.Variable
}

// PackBitsToFr fill bits to maximum size of fr length.
func PackBitsToFr(api frontend.API, bits []frontend.Variable) []frontend.Variable {
	bitSize := api.Compiler().FieldBitLen() - 1
	var r []frontend.Variable
	for i := 0; i < len(bits); i += bitSize {
		end := i + bitSize
		if end > len(bits) {
			end = len(bits)
		}
		z := api.FromBinary(bits[i:end]...)
		r = append(r, z)
	}
	return r
}

func Flip[T any](in []T) []T {
	res := make([]T, len(in))
	copy(res, in)
	for i := 0; i < len(in)/2; i++ {
		tmp := res[i]
		res[i] = res[len(res)-1-i]
		res[len(res)-1-i] = tmp
	}
	return res
}

// FlipByGroups flips the order of the groups of groupSize. e.g. [1,2,3,4,5,6] with groupSize 2 is flipped to [5,6,3,4,1,2]
func FlipByGroups[T any](in []T, groupSize int) []T {
	res := make([]T, len(in))
	copy(res, in)
	for i := 0; i < len(res)/groupSize/2; i++ {
		for j := 0; j < groupSize; j++ {
			a := i*groupSize + j
			b := len(res) - (i+1)*groupSize + j
			res[a], res[b] = res[b], res[a]
		}
	}
	return res
}

// BitsToFrBigInt convert bits to ecc fr, if the result exceed the modulus, return the mod result
func BitsToFrBigInt(b []uint, frModulus *big.Int) *big.Int {
	var res = big.NewInt(0)
	lens := len(b)
	for i := 0; i < lens; i++ {
		if b[i] != 0 {
			var decimal = big.NewInt(1)
			decimal.Lsh(decimal, uint(lens-i-1))
			res.Add(res, decimal)
			if res.Cmp(frModulus) > 0 {
				res.Mod(res, frModulus)
			}
		}
	}
	return res
}

// DecodeByteToBinary decode a byte to the binary uint array
func DecodeByteToBinary(byteValue byte) []uint {
	binaryIntArray := make([]uint, 8) // Create an empty integer array of size 8
	for i := 0; i < 8; i++ {
		if (byteValue & (1 << uint(i))) != 0 {
			binaryIntArray[i] = 1
		} else {
			binaryIntArray[i] = 0
		}
	}

	for i, j := 0, len(binaryIntArray)-1; i < j; i, j = i+1, j-1 {
		binaryIntArray[i], binaryIntArray[j] = binaryIntArray[j], binaryIntArray[i]
	}

	return binaryIntArray[:]
}

func Byte32ToBits(api frontend.API, b32 Bytes32, frSize int) []frontend.Variable {
	var bits []frontend.Variable
	bits = append(bits, api.ToBinary(b32.Val[0], frSize)...)
	bits = append(bits, api.ToBinary(b32.Val[1], 32*8-frSize)...)
	return bits
}
