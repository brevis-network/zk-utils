package keccak

import (
	"fmt"

	"github.com/brevis-network/zk-utils/circuits/gadgets/mux"

	"github.com/consensys/gnark/frontend"
)

func Keccak256BitsOptimized(api frontend.API, minRounds, maxRounds int, roundIndex frontend.Variable, data []frontend.Variable) (out [256]frontend.Variable) {
	if len(data) > maxRounds*1088 {
		panic("len(data) > maxRounds * 1088")
	}
	if len(data)%1088 != 0 {
		panic(fmt.Sprintf("invalid data length %d", len(data)))
	}
	var states [][1600]frontend.Variable
	// initial state
	states = append(states, newEmptyState())
	for i := 0; i < maxRounds; i++ {
		r := getRoundBits(data, i)
		s := absorbBits(api, states[i], r)
		states = append(states, s)
	}

	if maxRounds == 1 {
		copy(out[:], states[1][:256])
		return
	}
	roundIndex = api.Sub(roundIndex, minRounds-1)
	selected := mux.Multiplex(api, roundIndex, 256, maxRounds-minRounds+1, transpose2(states[minRounds:]))
	copy(out[:], selected[:256])
	return
}
