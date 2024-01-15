package utils

import (
	"fmt"
	"github.com/celer-network/goutils/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"reflect"
	"strings"
)

type ContractCodec struct {
	abi *abi.ABI
}

func NewContractCodec(abistr string) *ContractCodec {
	contractABI, err := abi.JSON(strings.NewReader(abistr))
	if err != nil {
		log.Panicf("cannot new contract call builder, %s", err.Error())
	}
	return &ContractCodec{abi: &contractABI}
}

func (c *ContractCodec) ABI() *abi.ABI {
	return c.abi
}

func (c *ContractCodec) Selector(methodName string) []byte {
	return c.abi.Methods[methodName].ID
}

func (c *ContractCodec) EncodeCalldata(methodName string, args ...interface{}) ([]byte, error) {
	method, ok := c.abi.Methods[methodName]
	if !ok {
		return nil, fmt.Errorf("method %s not fonud in abi", methodName)
	}
	inputs, err := method.Inputs.Pack(args...)
	if err != nil {
		return nil, err
	}
	return append(method.ID, inputs...), nil
}

func (c *ContractCodec) DecodeCalldata(methodName string, data []byte, ptrs ...interface{}) error {
	method, ok := c.abi.Methods[methodName]
	if !ok {
		return fmt.Errorf("method %s not fonud in abi", methodName)
	}
	return decodeArgs(method.Inputs, data, ptrs...)
}

func (c *ContractCodec) DecodeReturnData(methodName string, data []byte, ptrs ...interface{}) error {
	method, ok := c.abi.Methods[methodName]
	if !ok {
		return fmt.Errorf("method %s not fonud in abi", methodName)
	}
	return decodeArgs(method.Outputs, data, ptrs...)
}

func (c *ContractCodec) EncodeReturnData(methodName string, ifaces ...interface{}) ([]byte, error) {
	method, ok := c.abi.Methods[methodName]
	if !ok {
		return nil, fmt.Errorf("method %s not fonud in abi", methodName)
	}
	return method.Outputs.Pack(ifaces...)
}

func decodeArgs(args abi.Arguments, data []byte, ptrs ...interface{}) error {
	unpacked, err := args.Unpack(data)
	if err != nil {
		return err
	}
	if len(unpacked) != len(ptrs) {
		return fmt.Errorf(
			"length of the unpacked output (%d) does not equal to length of output interfaces (%d)",
			len(unpacked), len(ptrs))
	}
	for i := 0; i < len(ptrs); i++ {
		val := reflect.ValueOf(ptrs[i])
		if val.Kind() != reflect.Ptr {
			return fmt.Errorf("reference [%d] not pointer value", i)
		}
		unpackedVal := reflect.ValueOf(unpacked[i])
		if !unpackedVal.IsValid() {
			continue
		}
		if unpackedVal.Kind() == reflect.Ptr {
			val.Elem().Set(unpackedVal.Elem())
		} else {
			val.Elem().Set(unpackedVal)
		}
	}
	return nil
}
