package eth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/celer-network/goutils/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
)

type IClient interface {
	GetSlot(block uint64) (uint64, bool)
	GetBlock(slot uint64) uint64
	GetProof(ctx context.Context, account common.Address, keys []string, blockNumber *big.Int) (*gethclient.AccountResult, error)
}

type Clients map[uint64]IClient

type Client struct {
	*ethclient.Client
	GethClient *gethclient.Client
	ChainId    uint64
}

var _ IClient = &Client{}

func (ec *Client) GetProof(ctx context.Context, account common.Address, keys []string, blockNumber *big.Int) (*gethclient.AccountResult, error) {
	return ec.GethClient.GetProof(ctx, account, keys, blockNumber)
}

func (ec *Client) GetSlot(block uint64) (uint64, bool) {
	if ec.ChainId != 1 && ec.ChainId != 5 {
		log.Errorf("GetSlot failed, only ethereum supported")
		return 0, false
	}
	prefix := ""
	if ec.ChainId == 5 {
		prefix = "goerli."
	}
	// get slot based on block, see: https://beaconcha.in/api/v1/docs/index.html#/Execution/get_api_v1_execution_block__blockNumber_
	url := fmt.Sprintf("https://%sbeaconcha.in/api/v1/execution/block/%d", prefix, block)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Err is", err)
		return 0, false
	}
	if res.StatusCode != 200 {
		log.Errorf("failed to curl slot by block: status code %d", res.StatusCode)
		return 0, false
	}
	defer res.Body.Close()
	resBody, _ := io.ReadAll(res.Body)
	var jsonRes map[string]interface{}    // declaring a map for key names as string and values as interface
	_ = json.Unmarshal(resBody, &jsonRes) // Unmarshalling
	dataList := jsonRes["data"].([]interface{})
	data := dataList[0].(map[string]interface{})
	if data["posConsensus"] == nil {
		return 0, false
	}
	posConsensus := data["posConsensus"].(map[string]interface{})
	slot := uint64(0)
	finilized := false

	slot = uint64(posConsensus["slot"].(float64))
	finilized = posConsensus["finalized"].(bool)
	return slot, finilized
}

func (ec *Client) GetBlock(slot uint64) uint64 {
	if ec.ChainId == 56 || ec.ChainId == 97 {
		return slot
	}
	if ec.ChainId != 1 && ec.ChainId != 5 {
		log.Errorf("GetSlot failed, only ethereum supported")
		return 0
	}
	prefix := ""
	if ec.ChainId == 5 {
		prefix = "goerli."
	}
	// get slot based on block, see: https://beaconcha.in/api/v1/docs/index.html#/Slot/get_api_v1_slot__slotOrHash_
	url := fmt.Sprintf("https://%sbeaconcha.in/api/v1/slot/%d", prefix, slot)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Err is", err)
		return 0
	}
	if res.StatusCode != 200 {
		log.Errorf("failed to curl block by slot: status code %d", res.StatusCode)
		return 0
	}
	defer res.Body.Close()
	resBody, _ := io.ReadAll(res.Body)
	var jsonRes map[string]interface{}    // declaring a map for key names as string and values as interface
	_ = json.Unmarshal(resBody, &jsonRes) // Unmarshalling
	data := jsonRes["data"].(map[string]interface{})
	if data == nil {
		return 0
	}
	if data["exec_block_number"] == nil {
		return 0
	}
	return uint64(data["exec_block_number"].(float64))
}
