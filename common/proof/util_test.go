package proof

import (
	"context"
	"fmt"
	"github.com/consensys/gnark/test"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"testing"
)

func TestRecoverTransactionSignerAddress(t *testing.T) {
	assert := test.NewAssert(t)
	rpc, err := rpc.Dial("https://mainnet.infura.io/v3/18a7b9da0d3d462bb740c1485812b171")
	assert.NoError(err)
	ethClient := ethclient.NewClient(rpc)
	tx, _, err := ethClient.TransactionByHash(context.Background(), ethCommon.HexToHash("0x18f5d50299db75e46049bc44fc654b1e80f19f3dd33949adef467b43769c9a4d"))
	assert.NoError(err)
	addrBytes, err := RecoverTransactionSignerAddress(tx)
	assert.NoError(err)
	assert.Equal(fmt.Sprintf("%x", addrBytes), "0195b198088e464103e3840f52a1fa9ea81de84b")

}
