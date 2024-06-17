package utils

type ChainSMTStartBlockInfo struct {
	BlockNumber  uint64
	InitialIndex uint64
	ChainId      uint64
}

func GetStartBlockInfo(chainId uint64) ChainSMTStartBlockInfo {
	blockNumber := uint64(0)
	if chainId == 42161 {
		blockNumber = 185536768
	} else if chainId == 10 {
		blockNumber = 119300224
	}

	return ChainSMTStartBlockInfo{
		BlockNumber:  blockNumber,
		InitialIndex: blockNumber / 128,
		ChainId:      chainId,
	}
}
