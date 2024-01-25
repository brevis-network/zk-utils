package smt

import (
	"github.com/celer-network/zk-utils/circuits/gadgets/merkle"
	"github.com/celer-network/zk-utils/circuits/gadgets/rlp"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
)

func CheckSMTInclusion(
	api frontend.API,
	mimcHash mimc.MiMC,
	maxDepth int,
	leafValueTree []frontend.Variable,
	leafValueBranchesIndex []frontend.Variable,
	indexes []frontend.Variable,
	branches []frontend.Variable,
	rootHash frontend.Variable) frontend.Variable {
	leafValue := CalculateLeafValueHash(api, mimcHash, leafValueTree, leafValueBranchesIndex)

	hashResult := merkle.MerkleRootBasedOnMiMCHash(
		api,
		mimcHash,
		leafValue,
		indexes,
		branches,
	)

	return rlp.Equal(api, rootHash, hashResult)
}

func CalculateLeafValueHash(
	api frontend.API,
	mimcHash mimc.MiMC,
	leafValueTree []frontend.Variable,
	leafValueBranchesIndex []frontend.Variable,
) frontend.Variable {
	api.AssertIsEqual(len(leafValueTree)-1, len(leafValueBranchesIndex))
	return merkle.MerkleRootBasedOnMiMCHash(api, mimcHash, leafValueTree[0], leafValueBranchesIndex, leafValueTree[1:])
}
