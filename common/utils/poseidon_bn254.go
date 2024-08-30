package utils

import (
	"github.com/iden3/go-iden3-crypto/poseidon"
	"math/big"
)

type PoseidonBn254Hasher struct {
	preimages []*big.Int
}

func NewPoseidonBn254() *PoseidonBn254Hasher {
	return new(PoseidonBn254Hasher)
}

func (p *PoseidonBn254Hasher) Write(input *big.Int) {
	p.preimages = append(p.preimages, input)
}

func (p *PoseidonBn254Hasher) Reset() {
	p.preimages = []*big.Int{}
}

func (p *PoseidonBn254Hasher) Sum() (*big.Int, error) {
	return poseidon.Hash(p.preimages)
}
