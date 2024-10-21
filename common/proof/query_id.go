package proof

import (
	"math/big"

	"github.com/brevis-network/zk-utils/common/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func MiMCHashReceiptCustomInputs(
	receiptInfo *SDKQueryProvingInfoForReceipt,
) ([]byte, error) {
	hasher := utils.NewPoseidonBn254()

	var bits []uint
	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(receiptInfo.BlockNumber), 8*4)...)
	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(receiptInfo.BlockBaseFee), 8*16)...)
	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(receiptInfo.MPTKey), 8*4)...)
	for _, field := range receiptInfo.LogExtractInfos {
		contractAddress, err := hexutil.Decode(field.ContractAddress)
		if err != nil {
			return nil, err
		}
		bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(contractAddress), 8*20)...)

		bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(field.LogIndex), 8*2)...)

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
		hasher.Write(new(big.Int).SetBytes(common.LeftPadBytes(v.Bytes(), 32)))
	}

	result, err := hasher.Sum()
	if err != nil {
		return nil, err
	}
	return result.Bytes(), nil
}

func MiMCHashStorageCustomInputs(
	data *SDKQueryProvingInfoForStorageSlot,
) ([]byte, error) {
	hasher := utils.NewPoseidonBn254()

	var bits []uint

	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(data.BlockNumber), 8*4)...)
	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(data.BlockBaseFee), 8*16)...)

	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(data.AccountAddress), 8*20)...)

	var slot = utils.ParseBytes32(GetPaddedSlotBytes(data.Slot), 248)
	bits = append(bits, utils.Byte32ToFrBits(slot, 248)...)

	storageValue, _ := new(big.Int).SetString(data.SlotValue, 10)

	var value = utils.ParseBytes32(storageValue.Bytes(), 248)
	bits = append(bits, utils.Byte32ToFrBits(value, 248)...)

	roundData := utils.PackBitsToInt(bits)
	for _, v := range roundData {
		hasher.Write(new(big.Int).SetBytes(common.LeftPadBytes(v.Bytes(), 32)))
	}

	result, err := hasher.Sum()
	if err != nil {
		return nil, err
	}
	return result.Bytes(), nil
}

func MiMCHashTxCustomInputs(
	tsInfo *SDKQueryProvingInfoForTransaction,
) ([]byte, error) {
	hasher := utils.NewPoseidonBn254()

	var bits []uint
	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(tsInfo.BlockNumber), 8*4)...)
	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(tsInfo.BlockBaseFee), 8*16)...)
	bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(tsInfo.MPTKey), 8*4)...)

	// bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(tsInfo.ExtraInfo.ChainId), 8*4)...)
	// bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(tsInfo.ExtraInfo.Nonce), 8*4)...)
	// bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(tsInfo.ExtraInfo.MaxPriorityFeePerGas), 8*8)...)
	// bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(tsInfo.ExtraInfo.MaxFeePerGas), 8*8)...)
	// bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(tsInfo.ExtraInfo.GasLimit), 8*4)...)
	// bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(tsInfo.ExtraInfo.From), 8*20)...)
	// bits = append(bits, utils.DecomposeBits(utils.Var2BigInt(tsInfo.ExtraInfo.To), 8*20)...)

	// value, err := hexutil.Decode(tsInfo.ExtraInfo.Value)
	// if err != nil {
	// 	return nil, err
	// }
	// var value32Byte = utils.ParseBytes32(value, 248)
	// bits = append(bits, utils.Byte32ToFrBits(value32Byte, 248)...)

	leafHash := utils.Hex2Bytes(tsInfo.LeafHash)
	var leafHashByte = utils.ParseBytes32(leafHash, 248)
	bits = append(bits, utils.Byte32ToFrBits(leafHashByte, 248)...)

	roundData := utils.PackBitsToInt(bits)
	for _, v := range roundData {
		hasher.Write(new(big.Int).SetBytes(common.LeftPadBytes(v.Bytes(), 32)))
	}

	result, err := hasher.Sum()
	if err != nil {
		return nil, err
	}
	return result.Bytes(), nil
}

// 0x01 ===> 0x0000000000000000000000000000000000000000000000000000000000000001
func GetPaddedSlotBytes(slot string) []byte {
	slotBytes := utils.Hex2Bytes(slot)
	padded32SlotBytes := [32]byte{}
	copy(padded32SlotBytes[32-len(slotBytes):32], slotBytes)
	return padded32SlotBytes[:]
}
