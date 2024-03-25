package utils

import (
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	groth16_bn254 "github.com/consensys/gnark/backend/groth16/bn254"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/witness"
)

// only for bn254
func ExportProof(proof groth16.Proof) (a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, commitment [2]*big.Int, commitmentPok [2]*big.Int) {
	bn254Proof := proof.(*groth16_bn254.Proof)
	// proof.Ar, proof.Bs, proof.Krs
	a[0] = bn254Proof.Ar.X.BigInt(new(big.Int))
	a[1] = bn254Proof.Ar.Y.BigInt(new(big.Int))

	b[0][0] = bn254Proof.Bs.X.A1.BigInt(new(big.Int))
	b[0][1] = bn254Proof.Bs.X.A0.BigInt(new(big.Int))
	b[1][0] = bn254Proof.Bs.Y.A1.BigInt(new(big.Int))
	b[1][1] = bn254Proof.Bs.Y.A0.BigInt(new(big.Int))

	c[0] = bn254Proof.Krs.X.BigInt(new(big.Int))
	c[1] = bn254Proof.Krs.Y.BigInt(new(big.Int))

	commitment[0] = bn254Proof.Commitments[0].X.BigInt(new(big.Int))
	commitment[1] = bn254Proof.Commitments[0].Y.BigInt(new(big.Int))

	commitmentPok[0] = bn254Proof.CommitmentPok.X.BigInt(new(big.Int))
	commitmentPok[1] = bn254Proof.CommitmentPok.Y.BigInt(new(big.Int))
	return
}

func GenGroth16Bn254CommitmentPub(witness witness.Witness, vk groth16.VerifyingKey, proof groth16.Proof) string {
	// generate last input(hex string) of commitment
	w, ok := witness.Vector().(fr.Vector)
	if !ok {
		log.Errorf("witness.Vector().(fr.Vector) fail")
		return ""
	}

	lastInput, err := groth16_bn254.VerifyExportCommitPub(proof.(*groth16_bn254.Proof), vk.(*groth16_bn254.VerifyingKey), w)
	if err != nil {
		log.Errorln(err)
		return ""
	}
	var bs = lastInput.Bytes()
	return hexutil.Encode(bs[:])
}
