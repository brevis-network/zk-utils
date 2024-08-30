package poseidon

import (
	"github.com/consensys/gnark/frontend"
	poseidoncircuit "github.com/liyue201/gnark-circomlib/circuits"
)

// MiMC contains the params of the Mimc hash func and the curves on which it is implemented
type PoseidonCircuit struct {
	preimage []frontend.Variable // state storage. data is updated when Write() is called. Sum sums the data.
	api      frontend.API        // underlying constraint system
}

// NewMiMC returns a MiMC instance, that can be used in a gnark circuit
func NewBn254PoseidonCircuit(api frontend.API) (PoseidonCircuit, error) {
	return PoseidonCircuit{
		api: api,
	}, nil
}

func (p *PoseidonCircuit) Write(data frontend.Variable) {
	p.preimage = append(p.preimage, data)
}

func (p *PoseidonCircuit) Sum() frontend.Variable {
	return poseidoncircuit.Poseidon(p.api, p.preimage)
}
