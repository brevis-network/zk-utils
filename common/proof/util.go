package proof

import (
	"fmt"
	"math/big"

	fr_bls12377 "github.com/consensys/gnark-crypto/ecc/bls12-377/fr"
	fr_761 "github.com/consensys/gnark-crypto/ecc/bw6-761/fr"
	"github.com/consensys/gnark/backend/groth16"
	groth16_bls12377 "github.com/consensys/gnark/backend/groth16/bls12-377"
	groth16_761 "github.com/consensys/gnark/backend/groth16/bw6-761"
	"github.com/consensys/gnark/backend/witness"

	"github.com/celer-network/goutils/log"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

type unsignedTransactionData struct {
	ChainID    *big.Int
	Nonce      uint64
	GasTipCap  *big.Int // a.k.a. maxPriorityFeePerGas
	GasFeeCap  *big.Int // a.k.a. maxFeePerGas
	Gas        uint64
	To         *common.Address `rlp:"nil"` // nil means contract creation
	Value      *big.Int
	Data       []byte
	AccessList ethtypes.AccessList
}

func PrepareUnsignedTransactionRlp(transaction *ethtypes.Transaction) ([]byte, error) {
	unsignedTransaction := &unsignedTransactionData{
		ChainID:    transaction.ChainId(),
		Nonce:      transaction.Nonce(),
		GasTipCap:  transaction.GasTipCap(),
		GasFeeCap:  transaction.GasFeeCap(),
		Gas:        transaction.Gas(),
		To:         transaction.To(),
		Value:      transaction.Value(),
		Data:       transaction.Data(),
		AccessList: transaction.AccessList(),
	}

	return rlp.EncodeToBytes(unsignedTransaction)
}

type unsignedLegacyTransactionData struct {
	Nonce    uint64
	GasPrice *big.Int
	Gas      uint64
	To       *common.Address `rlp:"nil"`
	Value    *big.Int
	Data     []byte
	ChainID  *big.Int
	Pad1     *big.Int
	Pad2     *big.Int
}

func PrepareUnsignedLegacyTransactionRlp(transaction *ethtypes.Transaction) ([]byte, error) {
	unsignedTransaction := &unsignedLegacyTransactionData{
		Nonce:    transaction.Nonce(),
		GasPrice: transaction.GasPrice(),
		Gas:      transaction.Gas(),
		To:       transaction.To(),
		Value:    transaction.Value(),
		Data:     transaction.Data(),
		ChainID:  transaction.ChainId(),
		Pad1:     big.NewInt(0),
		Pad2:     big.NewInt(0),
	}

	return rlp.EncodeToBytes(unsignedTransaction)
}

func RecoverTransactionSignerAddress(transaction *ethtypes.Transaction) ([]byte, error) {
	unsignedTransactionRlp, err := PrepareUnsignedTransactionRlp(transaction)
	if err != nil {
		log.Errorf("Failed to get unsigned transaction rlp %s\n", err.Error())
		return nil, err
	}

	transactionType := transaction.Type()
	var unsignedTxBytes []byte
	unsignedTxBytes = append(unsignedTxBytes, transactionType)
	unsignedTxBytes = append(unsignedTxBytes, unsignedTransactionRlp[:]...)

	preImage := crypto.Keccak256(unsignedTxBytes)

	var sigBytes []byte
	v, r, s := transaction.RawSignatureValues()
	var rBuf [32]byte
	var sBuf [32]byte
	sigBytes = append(sigBytes, r.FillBytes(rBuf[:])...)
	sigBytes = append(sigBytes, s.FillBytes(sBuf[:])...)
	if v.Int64() == 0 {
		sigBytes = append(sigBytes, 0)
	} else {
		sigBytes = append(sigBytes, v.Bytes()...)
	}

	uncompressedPublicKey, err := crypto.Ecrecover(preImage, sigBytes)
	if err != nil {
		log.Errorf("Failed to ecrecover %s\n", err.Error())
		return nil, err
	}
	publicKey, err := crypto.UnmarshalPubkey(uncompressedPublicKey)
	if err != nil {
		log.Errorf("Failed to unmarshal pubkey %s\n", err.Error())
		return nil, err
	}
	signer := crypto.PubkeyToAddress(*publicKey)
	return signer[:], nil
}

func Verify377ReturnCommitPub(proof groth16.Proof, vk groth16.VerifyingKey, pubW witness.Witness) (fr_bls12377.Element, error) {
	w, ok := pubW.Vector().(fr_bls12377.Vector)
	if !ok {
		return fr_bls12377.Element{}, fmt.Errorf("witness.Vector().(fr.Vector) fail")
	} else {
		return groth16_bls12377.VerifyExportCommitPub(proof.(*groth16_bls12377.Proof), vk.(*groth16_bls12377.VerifyingKey), w)
	}
}

func Verify761ReturnCommitPub(proof groth16.Proof, vk groth16.VerifyingKey, pubW witness.Witness) (fr_761.Element, error) {
	w, ok := pubW.Vector().(fr_761.Vector)
	if !ok {
		return fr_761.Element{}, fmt.Errorf("witness.Vector().(fr.Vector) fail")
	} else {
		return groth16_761.VerifyExportCommitPub(proof.(*groth16_761.Proof), vk.(*groth16_761.VerifyingKey), w)
	}
}
