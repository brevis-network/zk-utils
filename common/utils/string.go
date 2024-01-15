package utils

import (
	"math/big"
	"strconv"
)

func Str2BigInt(str string) *big.Int {
	b := new(big.Int)
	b.SetString(str, 10)
	return b
}

func Str2Int(str string) int {
	ret, _ := strconv.Atoi(str)
	return ret
}

func Str2Uint64(str string) uint64 {
	return uint64(Str2Int(str))
}
