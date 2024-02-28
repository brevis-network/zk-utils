package test

import (
	"fmt"
	"github.com/brevis-network/zk-utils/circuits/gadgets/index-proof/core"
	"github.com/celer-network/goutils/log"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/test"
	"github.com/ethereum/go-ethereum/rlp"
	"strconv"
	"testing"
)

func Test_INDEX(t *testing.T) {
	var indexBuf []byte
	indexBuf = rlp.AppendUint64(indexBuf, uint64(100000))
	log.Infof("%x", indexBuf)

}

func Test_INDEX_CHECK(t *testing.T) {
	assert := test.NewAssert(t)

	for i := 0; i < 10000; i++ {
		var indexBuf []byte
		indexBuf = rlp.AppendUint64(indexBuf, uint64(i))
		input := GetHexArray(fmt.Sprintf("%x", indexBuf), 7)
		if len(input) != 7 {
			log.Fatalf("invalid input, index: %v", i)
		}
		log.Infof("%v %x %d", indexBuf, indexBuf, i)
		var witnessInput [7]frontend.Variable
		for x, y := range input {
			witnessInput[x] = y
		}
		witness := core.IndexCheckCircuit{
			Index:     i,
			RlpString: witnessInput,
		}
		err := test.IsSolved(&core.IndexCheckCircuit{}, &witness, ecc.BN254.ScalarField())
		assert.NoError(err)
	}
}

func GetHexArray(hexStr string, maxLen int) (res []frontend.Variable) {
	for i := 0; i < maxLen; i++ {
		if i < len(hexStr) {
			intValue, _ := strconv.ParseInt(string(hexStr[i]), 16, 64)
			res = append(res, intValue)
		} else {
			res = append(res, 0)
		}
	}
	return
}
