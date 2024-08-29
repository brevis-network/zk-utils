package utils

import (
	"github.com/iden3/go-iden3-crypto/poseidon"
	"math/big"
)

type PoseidonHasher struct {
	preimages []*big.Int
}

func NewPoseidon() *PoseidonHasher {
	return new(PoseidonHasher)
}

func (p *PoseidonHasher) Write(input *big.Int) {
	p.preimages = append(p.preimages, input)
}

func (p *PoseidonHasher) Reset() {
	p.preimages = []*big.Int{}
}

func (p *PoseidonHasher) Sum() (*big.Int, error) {
	return poseidon.Hash(p.preimages)
}
