package proof

import (
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	if transaction.Type() == types.LegacyTxType {
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
	} else if transaction.Type() == types.DynamicFeeTxType {
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
	} else {
		return nil, fmt.Errorf("unsupported transaction type")
	}
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

func RecoverTransactionSignerAddress(transaction *ethtypes.Transaction) ([]byte, error) {
	unsignedTransactionRlp, err := PrepareUnsignedTransactionRlp(transaction)
	if err != nil {
		log.Errorf("Failed to get unsigned transaction rlp %s\n", err.Error())
		return nil, err
	}

	transactionType := transaction.Type()
	var unsignedTxBytes []byte
	if transactionType != types.LegacyTxType {
		unsignedTxBytes = append(unsignedTxBytes, transactionType)
	}
	unsignedTxBytes = append(unsignedTxBytes, unsignedTransactionRlp[:]...)

	preImage := crypto.Keccak256(unsignedTxBytes)

	var sigBytes []byte
	v, r, s := transaction.RawSignatureValues()

	if transactionType == types.LegacyTxType {
		actualV := v.Uint64() - 35 - transaction.ChainId().Uint64()*2
		v = new(big.Int).SetUint64(actualV)
	}

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
