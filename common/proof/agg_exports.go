package proof

import (
	"github.com/celer-network/zk-utils/common/utils"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/witness"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type ZKSDKAggPubInputs struct {
	AggVkHash    string
	SMTRoot      [2]string
	AppVkHash    string
	OutputCommit [2]string
	QueryId      string
}

func (pubs *ZKSDKAggPubInputs) GenerateBatchProofAndPublicInputs(proof groth16.Proof,
	witness witness.Witness,
	vk groth16.VerifyingKey,
) string {
	a, b, c, commitment := utils.ExportProof(proof)
	cPub := utils.GenGroth16Bn254CommitmentPub(witness, vk, proof)
	var A [2]string
	for i := 0; i < 2; i++ {
		A[i] = hexutil.Encode(a[i].Bytes())
	}

	var B [2][2]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			B[i][j] = hexutil.Encode(b[i][j].Bytes())
		}
	}

	var C [2]string
	for i := 0; i < 2; i++ {
		C[i] = hexutil.Encode(c[i].Bytes())
	}

	var Commitment [2]string
	for i := 0; i < 2; i++ {
		Commitment[i] = hexutil.Encode(commitment[i].Bytes())
	}

	proofData := Groth16Proof{
		A:          A,
		B:          B,
		C:          C,
		Commitment: Commitment,
	}

	var result = ""
	result += proofData.A[0] + ","
	result += proofData.A[1] + ","
	result += proofData.B[0][0] + ","
	result += proofData.B[0][1] + ","
	result += proofData.B[1][0] + ","
	result += proofData.B[1][1] + ","
	result += proofData.C[0] + ","
	result += proofData.C[1] + ","
	result += proofData.Commitment[0] + ","
	result += proofData.Commitment[1] + ","
	result += pubs.QueryId + ","
	result += pubs.SMTRoot[0] + ","
	result += pubs.SMTRoot[1] + ","
	result += pubs.AggVkHash + ","
	result += pubs.OutputCommit[0] + ","
	result += pubs.OutputCommit[1] + ","
	result += pubs.AppVkHash + ","
	result += cPub + ","

	return result
}
