package proof

import (
	"context"
	"fmt"
	"testing"

	"github.com/consensys/gnark/test"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func TestRecoverTransactionSignerAddress(t *testing.T) {
	assert := test.NewAssert(t)
	rpc, err := rpc.Dial("https://mainnet.infura.io/v3/18a7b9da0d3d462bb740c1485812b171")
	assert.NoError(err)
	ethClient := ethclient.NewClient(rpc)
	tx, _, err := ethClient.TransactionByHash(context.Background(), ethCommon.HexToHash("0x8d97af67ecbac859107af5901ecda0d1f6d437fd3e02db7caa87702c1513b83a"))
	assert.NoError(err)
	addrBytes, err := RecoverTransactionSignerAddress(tx)
	assert.NoError(err)
	assert.Equal(fmt.Sprintf("%x", addrBytes), "c0471f8a0ca9ac3e234448344284dca8e93e0b58")
}

func TestRecoverTransactionSignerAddressForType0(t *testing.T) {
	assert := test.NewAssert(t)
	rpc, err := rpc.Dial("https://mainnet.infura.io/v3/18a7b9da0d3d462bb740c1485812b171")
	assert.NoError(err)
	ethClient := ethclient.NewClient(rpc)
	tx, _, err := ethClient.TransactionByHash(context.Background(), ethCommon.HexToHash("0xab7cded2ac591a069a509b2bcf84fa0101b409feb0dee555c661b076e76c6953"))
	assert.NoError(err)
	addrBytes, err := RecoverTransactionSignerAddress(tx)
	assert.NoError(err)
	assert.Equal(fmt.Sprintf("%x", addrBytes), "fCAAf32bDE89A81C71a8C2Fb2C7Bc8358C9E3696")
}
