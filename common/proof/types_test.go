package proof

import (
	"context"
	"math/big"
	"reflect"
	"testing"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/zk-utils/common/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestChunkProof_ProofData(t *testing.T) {
	type fields struct {
		ChunkRoot     string
		PrevHash      string
		EndHash       string
		StartBlockNum uint64
		EndBlockNum   uint64
		Proof         Groth16Proof
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "default",
			fields: fields{
				ChunkRoot:     "0x4bf40b6be2b3922282427d456f8ba76271cd296029a7255c5907b87fc729e88d",
				PrevHash:      "0x00000000000000000000000000000000000000000000000000000000000000ff",
				EndHash:       "0x000000000000000000000000000000000000000000000000000000000000017f",
				StartBlockNum: 256,
				EndBlockNum:   383,
				Proof: Groth16Proof{
					A: [2]string{"0x23daa084f7eea1dbb39fea65f11361a15f475a06ba24e457ec91f008a5398635", "0x046d53ea95f2468fedfb9fe9260f65a921b4da5f28fd31caf558a39ce82dd8d3"},
					B: [2][2]string{
						{"0x168b54d53dd8c3c73d0df973aff474f6ab94887964cde70940ff4db4f6b9129f", "0x25f9c5214b259b6b9351453d5812378dc96abb631d5aa72266b0aec8ae739e14"},
						{"0x07c1bec5a7d9fed0b3d21560248e1b6a7aad77c19ecb94b7fbb1d3ce24fb2f77", "0x13c75213428e00d60a3cb884892ace0599722676fea52ea0febcc627f3d06d19"},
					},
					C:          [2]string{"0x0958800af01bdb80d89660e94d0fff3f2677ea8a8f2286b80fa66aaadba3d4e6", "0x17a607eda80f2c1134380f043c83121aa85d66540fc54e55ef4aa3ebf037a7fe"},
					Commitment: [2]string{"0x00", "0x00"},
				},
			},
			want: utils.Hex2Bytes("0x23daa084f7eea1dbb39fea65f11361a15f475a06ba24e457ec91f008a5398635046d53ea95f2468fedfb9fe9260f65a921b4da5f28fd31caf558a39ce82dd8d3168b54d53dd8c3c73d0df973aff474f6ab94887964cde70940ff4db4f6b9129f25f9c5214b259b6b9351453d5812378dc96abb631d5aa72266b0aec8ae739e1407c1bec5a7d9fed0b3d21560248e1b6a7aad77c19ecb94b7fbb1d3ce24fb2f7713c75213428e00d60a3cb884892ace0599722676fea52ea0febcc627f3d06d190958800af01bdb80d89660e94d0fff3f2677ea8a8f2286b80fa66aaadba3d4e617a607eda80f2c1134380f043c83121aa85d66540fc54e55ef4aa3ebf037a7fe00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004bf40b6be2b3922282427d456f8ba7620000000000000000000000000000000071cd296029a7255c5907b87fc729e88d000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ff0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000017f0000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000017f"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ChunkProof{
				ChunkRoot:     tt.fields.ChunkRoot,
				PrevHash:      tt.fields.PrevHash,
				EndHash:       tt.fields.EndHash,
				StartBlockNum: tt.fields.StartBlockNum,
				EndBlockNum:   tt.fields.EndBlockNum,
				Proof:         tt.fields.Proof,
			}
			if got := p.EncodeABI(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeABI() = %x\nwant %x", got, tt.want)
			}
		})
	}
}

func TestReceiptProof_ProofData(t *testing.T) {
	ec, err := ethclient.Dial("https://ethereum.blockpi.network/v1/rpc/public")
	if err != nil {
		log.Fatalln(err)
	}

	bk, err := ec.BlockByNumber(context.Background(), new(big.Int).SetUint64(17490377)) // 21 transactions, with ext nodes
	if err != nil {
		log.Fatal(err)
	}

	var receipts types.Receipts
	for _, tx := range bk.Transactions() {
		rec, err := ec.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		receipts = append(receipts, rec)
	}

	nodes, indexBuff, _, err := GetReceiptProof(bk, receipts, 5)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("index buffer:%x", indexBuff)
	for _, b := range nodes {
		log.Infof("node: %x", b)
	}
}
