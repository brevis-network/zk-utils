package proof

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/celer-network/zk-utils/common/utils"
	"github.com/consensys/gnark-crypto/ecc/bls12-377/fr/mimc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/iden3/go-iden3-crypto/keccak256"
)

func MiMCHashReceiptCustomInputs(
	receiptInfo *SDKQueryProvingInfoForReceipt,
) ([]byte, error) {
	hasher := mimc.NewMiMC()

	var bits []uint
	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(receiptInfo.BlkNum), 8*4)...)
	for _, field := range receiptInfo.LogExtractInfos {
		contractAddress, err := hexutil.Decode(field.ContractAddress)
		if err != nil {
			return nil, err
		}
		bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(contractAddress), 8*20)...)

		logTopic0, err := hexutil.Decode(field.LogTopic0)
		if err != nil {
			return nil, err
		}
		bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(logTopic0[:6]), 8*6)...)

		valueFromTopic := 0
		if field.ValueFromTopic {
			valueFromTopic = 1
		}
		bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(valueFromTopic), 1)...)

		bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(field.ValueIndex), 7)...)

		value, err := hexutil.Decode(field.Value)
		if err != nil {
			return nil, err
		}

		var value32Byte = utils.ParseBytes32(value, 248)
		bits = append(bits, utils.Byte32ToFrBits(value32Byte, 248)...)
	}

	roundData := utils.PackBitsToInt(bits)
	for _, v := range roundData {
		hasher.Write(common.LeftPadBytes(v.Bytes(), 32))
	}

	return hasher.Sum(nil), nil
}

func MiMCHashStorageCustomInputs(
	data *SDKQueryProvingInfoForStorageSlot,
) ([]byte, error) {
	hasher := mimc.NewMiMC()

	var bits []uint

	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(data.BlockNumber), 8*4)...)
	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(data.AccountAddress), 8*20)...)

	slotKey := GetStorageMPTProofKey(data.StorageKey)
	var slot32Byte = utils.ParseBytes32(hexutil.MustDecode(slotKey), 248)
	bits = append(bits, utils.Byte32ToFrBits(slot32Byte, 248)...)

	storageValue, _ := new(big.Int).SetString(data.SlotValue, 10)

	var slotV32Byte = utils.ParseBytes32(storageValue.Bytes(), 248)
	bits = append(bits, utils.Byte32ToFrBits(slotV32Byte, 248)...)

	roundData := utils.PackBitsToInt(bits)
	for _, v := range roundData {
		hasher.Write(common.LeftPadBytes(v.Bytes(), 32))
	}

	return hasher.Sum(nil), nil
}

func MiMCHashTxCustomInputs(
	tsInfo *SDKQueryProvingInfoForTransaction,
) ([]byte, error) {
	hasher := mimc.NewMiMC()

	var bits []uint

	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(tsInfo.BlockNumber), 8*4)...)
	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(tsInfo.ExtraInfo.ChainId), 8*4)...)
	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(tsInfo.ExtraInfo.Nonce), 8*4)...)
	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(tsInfo.ExtraInfo.MaxPriorityFeePerGas), 8*8)...)
	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(tsInfo.ExtraInfo.MaxFeePerGas), 8*8)...)
	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(tsInfo.ExtraInfo.GasLimit), 8*4)...)
	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(tsInfo.ExtraInfo.From), 8*20)...)
	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(tsInfo.ExtraInfo.To), 8*20)...)

	value, err := hexutil.Decode(tsInfo.ExtraInfo.Value)
	if err != nil {
		return nil, err
	}
	var value32Byte = utils.ParseBytes32(value, 248)
	bits = append(bits, utils.Byte32ToFrBits(value32Byte, 248)...)

	roundData := utils.PackBitsToInt(bits)
	for _, v := range roundData {
		hasher.Write(common.LeftPadBytes(v.Bytes(), 32))
	}

	return hasher.Sum(nil), nil
}

func GetStorageMPTProofKey(storageKey string) string {
	key := strings.ReplaceAll(storageKey, "0x", "")
	storageKeyLength := len(key)
	for i := 0; i < 64-storageKeyLength; i++ {
		key = fmt.Sprintf("0%s", key)
	}
	key = fmt.Sprintf("0x%s", key)
	storageProofKey := hexutil.MustDecode(key)

	return hexutil.Encode(keccak256.Hash(storageProofKey))
}
