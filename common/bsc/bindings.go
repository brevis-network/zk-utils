// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bsc

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CrossChainChannelInit is an auto generated low-level Go binding around an user-defined struct.
type CrossChainChannelInit struct {
	ChannelId uint8
	Sequence  uint64
}

// MerkleRootData is an auto generated low-level Go binding around an user-defined struct.
type MerkleRootData struct {
	Hash []byte
}

// PoALightClientBNBHeaderInfo is an auto generated low-level Go binding around an user-defined struct.
type PoALightClientBNBHeaderInfo struct {
	ParentHash       [32]byte
	Sha3Uncles       [32]byte
	Miner            common.Address
	StateRoot        [32]byte
	TransactionsRoot [32]byte
	ReceiptsRoot     [32]byte
	LogsBloom        []byte
	Difficulty       *big.Int
	Number           *big.Int
	GasLimit         uint64
	GasUsed          uint64
	Timestamp        uint64
	ExtraData        []byte
	MixHash          [32]byte
	Nonce            [8]byte
}

// TimestampData is an auto generated low-level Go binding around an user-defined struct.
type TimestampData struct {
	Seconds int64
	Nanos   int32
}

// BSCValidatorSetMetaData contains all meta data concerning the BSCValidatorSet contract.
var BSCValidatorSetMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"FailedWithReasonStr\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"UnexpectedPackage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ValidatorSetUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EPOCH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_FAIL_CHECK_VALIDATORS\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_LEN_OF_VAL_MISMATCH\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_UNKNOWN_PACKAGE_TYPE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXPIRE_TIME_SECOND_GAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_NUM_OF_CABINETS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NUM_OF_VALIDATORS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORS_UPDATE_MESSAGE_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentValidatorSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"feeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BBCFeeAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"votingPower\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"incoming\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"currentValidatorSetMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expireTimeSecondGap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleAckPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleFailAckPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleSynPackage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"responsePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_system\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_initValidatorSetBytes\",\"type\":\"bytes\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"valAddress\",\"type\":\"address\"}],\"name\":\"isCurrentValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461001657611cab908161001c8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c80631182b8751461011257806355614fcc1461010d5780635667515a146101085780635d77156c146101035780636969a25c146100fe57806381650b62146100f9578063831d65d1146100d6578063853230aa146100f457806386249882146100ef578063a0dc2758146100ea578063ad3c9da6146100e5578063b8cf4ef1146100e0578063c0d91eaf146100db578063c8509d81146100d6578063e086c7b1146100d15763eda5868c146100cc57600080fd5b610836565b61081a565b6103c0565b61065e565b610533565b6104f5565b6104d9565b6104bb565b61049e565b6103a4565b610316565b6102a5565b610289565b610249565b6101b3565b604060031982011261016e5760043560ff8116810361016e579160243567ffffffffffffffff9283821161016e578060238301121561016e57816004013593841161016e576024848301011161016e576024019190565b600080fd5b919082519283825260005b84811061019f575050826000602080949584010152601f8019910116010190565b60208183018101518483018201520161017e565b3461016e576101c136610117565b915060ff60005416156101f3576101ef916101db91610b28565b604051918291602083526020830190610173565b0390f35b60405162461bcd60e51b815260206004820152600f60248201527f6e6f7420696e697469616c697a656400000000000000000000000000000000006044820152606490fd5b6001600160a01b0381160361016e57565b3461016e57602036600319011261016e576001600160a01b0360043561026e81610238565b16600052600460205260206040600020541515604051908152f35b3461016e57600036600319011261016e57602060405160008152f35b3461016e57600036600319011261016e57602060405160678152f35b634e487b7160e01b600052603260045260246000fd5b60015481101561031157600160005260021b7fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60190600090565b6102c1565b3461016e57602036600319011261016e5760043560015481101561016e5761033d906102d7565b50805460018201546002830154600390930154604080516001600160a01b03948516815292841660208401529284169282019290925260a083811c67ffffffffffffffff16606083015260e09390931c60ff1615156080820152918201528060c081010390f35b3461016e57600036600319011261016e57602060405160668152f35b3461016e5760046103d036610117565b9092916001600160a01b0391602083600354166040519586809263248bedb360e11b82525afa8015610499577faa5ba621c8b3d7d05bb9e51a7506108251d4d5dbe542ca66fc7bb52aacb02b659561043a606095849360ff9860009161046b575b50163314610a30565b60405195869416845260406020850152816040850152848401376000828201840152601f01601f19168101030190a1005b61048c915060203d8111610492575b61048481836105da565b810190610a0c565b38610431565b503d61047a565b610a24565b3461016e57600036600319011261016e5760206040516103e88152f35b3461016e57600036600319011261016e576020600254604051908152f35b3461016e57600036600319011261016e57602060405160c88152f35b3461016e57602036600319011261016e576001600160a01b0360043561051a81610238565b1660005260046020526020604060002054604051908152f35b3461016e57600036600319011261016e57602060405160158152f35b634e487b7160e01b600052604160045260246000fd5b60c0810190811067ffffffffffffffff82111761058157604052565b61054f565b6040810190811067ffffffffffffffff82111761058157604052565b6020810190811067ffffffffffffffff82111761058157604052565b6060810190811067ffffffffffffffff82111761058157604052565b90601f8019910116810190811067ffffffffffffffff82111761058157604052565b6040519061060982610586565b565b67ffffffffffffffff811161058157601f01601f191660200190565b9291926106338261060b565b9161064160405193846105da565b82948184528183011161016e578281602093846000960137010152565b3461016e57604036600319011261016e5760043561067b81610238565b60243567ffffffffffffffff811161016e573660238201121561016e576106ac903690602481600401359101610627565b60ff600054166107d5576106bf90611215565b156107865760005b602082018051805183101561073e5791610737916106f06106ea83602096610889565b51610999565b61073161071861070a84610703816109c7565b9451610889565b51516001600160a01b031690565b6001600160a01b03166000526004602052604060002090565b55610868565b90506106c7565b6107726001600160a01b03866107556103e8600255565b166001600160a01b03166001600160a01b03196003541617600355565b610784600160ff196000541617600055565b005b60405162461bcd60e51b815260206004820152602160248201527f6661696c656420746f20706172736520696e69742076616c696461746f7253656044820152601d60fa1b6064820152608490fd5b60405162461bcd60e51b815260206004820152601360248201527f616c726561647920696e697469616c697a6564000000000000000000000000006044820152606490fd5b3461016e57600036600319011261016e57602060405160298152f35b3461016e57600036600319011261016e57602060405160658152f35b634e487b7160e01b600052601160045260246000fd5b60001981146108775760010190565b610852565b8051156103115760200190565b80518210156103115760209160051b010190565b634e487b7160e01b600052600060045260246000fd5b919061099457805182546001600160a01b0319166001600160a01b0391821617835560039160a0916109048160208401511660018701906001600160a01b03166001600160a01b0319825416179055565b61092e600286019160408401511682906001600160a01b03166001600160a01b0319825416179055565b606082015181547fffffff000000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff000000000000000000000000000000000000000060ff60e01b6080870151151560e01b1693871b169116171790550151910155565b61089d565b6001549068010000000000000000821015610581576109c182600161060994016001556102d7565b906108b3565b906001820180921161087757565b906037820180921161087757565b906080820180921161087757565b9060c0820180921161087757565b9190820180921161087757565b9081602091031261016e5751610a2181610238565b90565b6040513d6000823e3d90fd5b15610a3757565b60405162461bcd60e51b815260206004820152601860248201527f6e6f742063726f737320636861696e20636f6e747261637400000000000000006044820152606490fd5b9081602091031261016e575163ffffffff8116810361016e5790565b60405190610aa5826105a2565b60008252565b604051610ab7816105a2565b60008152906000368137565b60405190610ad082610586565b6001825260203681840137565b60405190610aea82610586565b60208083523683820137565b90610b008261060b565b610b0d60405191826105da565b8281528092610b1e601f199161060b565b0190602036910137565b90610b4a610b3e6003546001600160a01b031690565b6001600160a01b031690565b906040519263248bedb360e11b84526020938481600481875afa918215610499578593610b906001600160a01b03610b9c95610b9795600091610caa5750163314610a30565b3691610627565b611215565b92909215610c5b5750508160ff610bb4835160ff1690565b16610c5157610bc4910151610cc1565b905b600481610bde610b3e6003546001600160a01b031690565b6040516355a8ddcb60e11b815292839182905afa91821561049957600092610c24575b505063ffffffff828116911603610c1b5750610a21610aab565b610a2190611433565b610c439250803d10610c4a575b610c3b81836105da565b810190610a7c565b3880610c01565b503d610c31565b5050606590610bc6565b604051630bee7a6760e01b81529250829060049082905afa90811561049957610a2192600092610c8d575b5050611433565b610ca39250803d10610c4a57610c3b81836105da565b3880610c86565b61048c9150883d8a116104925761048481836105da565b610cca81610fae565b9015610ead5750600180549082519060005b838110610e1b575050808211610dfb575b80821015610df35750905b60005b828110610d7e575050506040517fdf380451f7286b929282db56cae856ef8d5d940e01a022844d086bbedc51f755600080a1602081600481610d48610b3e6003546001600160a01b031690565b6355a8ddcb60e11b82525afa90811561049957600091610d66575090565b610a21915060203d8111610c4a57610c3b81836105da565b80610daf610dab610d92610db89486610889565b51610da5610d9f856102d7565b50611040565b906110eb565b1590565b610dbd57610868565b610cfb565b610dc6816109c7565b610dd661071861070a8487610889565b55610dee610de48285610889565b516109c1836102d7565b610868565b905090610cf8565b805b828110610e0a5750610ced565b610e1690610dee6110a0565b610dfd565b81610e28610d9f836102d7565b9060005b858110610e68575b5090610e469291610e4b575b50610868565b610cdc565b610e61610718600092516001600160a01b031690565b5538610e40565b82516001600160a01b03166001600160a01b03610e8b610b3e61070a858d610889565b911614610ea057610e9b90610868565b610e2c565b5060009050610e46610e34565b7f9e61be46205d080ba82669bf51f14a527ab993061b4555c481f877592f92949f9150610ee890604051918291602083526020830190610173565b0390a1606690565b60405190610efd826105be565b602982527f746865206c696d697400000000000000000000000000000000000000000000006040837f746865206e756d626572206f662076616c696461746f7273206578636565642060208201520152565b60405190610f5c826105be565b602b82527f616c696461746f725365740000000000000000000000000000000000000000006040837f6475706c696361746520636f6e73656e7375732061646472657373206f66207660208201520152565b6029815111611034576000905b80518210156110275760005b828110610fdf575090610fd990610868565b90610fbb565b610fec61070a8484610889565b6001600160a01b03611004610b3e61070a8587610889565b9116146110195761101490610868565b610fc7565b505050600090610a21610f4f565b5050600190610a21610a98565b50600090610a21610ef0565b9060405161104d81610565565b60a06003829460ff6001600160a01b038083541686528060018401541660208701526002830154908116604087015267ffffffffffffffff81861c16606087015260e01c16151560808501520154910152565b60015480156110d557600019016110b6816102d7565b6109945760038160008093558260018201558260028201550155600155565b634e487b7160e01b600052603160045260246000fd5b906001600160a01b0391828151168383511614928361114a575b83611133575b508261111657505090565b60609081015191015167ffffffffffffffff908116911614919050565b60408083015190840151821691161492503861110b565b809350602082015116836020840151161492611105565b6040519061116e82610586565b6060602083600081520152565b67ffffffffffffffff81116105815760051b60200190565b604051906111a082610565565b8160a06000918281528260208201528260408201528260608201528260808201520152565b906111cf8261117b565b6111dc60405191826105da565b82815280926111ed601f199161117b565b019060005b8281106111fe57505050565b602090611209611193565b828285010152016111f2565b9061121e611161565b5061123861123361122d611161565b9361151a565b611540565b6000805b611245836114ff565b15611321578061127c5761127790610dee61127061126a611265876114b6565b6116fd565b60ff1690565b60ff168752565b61123c565b9060019081830361131a5750611299611294846114b6565b61158e565b6112a381516111c5565b60208701908152906000835b6112c1575b5050506112779091610868565b8151811015611315576112dd6112d78284610889565b51611327565b156113095781610e4085949361130193889751906112fb8383610889565b52610889565b9091926112af565b50505050505050600090565b6112b4565b9392505050565b50919050565b61132f611193565b5061134161133b611193565b91611540565b90600091825b611350826114ff565b1561142e57806113845761137f90610dee61137261136d856114b6565b611688565b6001600160a01b03168552565b611347565b60018181036113b4575061137f90610dee6113a4610b3e61136d866114b6565b6001600160a01b03166020860152565b600282036113e0575061137f90610dee6113d061136d856114b6565b6001600160a01b03166040860152565b909360038503611425575061137f9061141f61140e611401611265866114b6565b67ffffffffffffffff1690565b67ffffffffffffffff166060860152565b93610868565b94935050509190565b505091565b9060405161144081610586565b6001815260005b60208082101561146257906060602092828501015201611447565b5050909161147863ffffffff61147d9216611bb6565b611961565b9080511561031157610a219160208201526114978161087c565b506119ab565b604051906114aa82610586565b60006020838281520152565b6114be61149d565b506114c8816114ff565b1561016e5760200180516114db81611854565b91828201908183116108775752604051916114f583610586565b8252602082015290565b60208151910151602082015191518201809211610877571090565b61152261149d565b5060208151916040519261153584610586565b835201602082015290565b6000602060405161155081610586565b61155861149d565b8152015261156581611661565b1561016e576020810151611578816118f0565b810180911161087757604051916114f583610586565b61159781611661565b1561016e576115a5816117a6565b6115ae8161117b565b916115bc60405193846105da565b818352601f196115cb8361117b565b0160005b81811061164a5750506115f06020809201516115ea816118f0565b906109ff565b6000905b838210611602575050505090565b61163e8161161261164493611854565b9061161b6105fc565b828152818782015261162d868a610889565b526116388589610889565b506109ff565b91610868565b906115f4565b60209061165561149d565b828288010152016115cf565b80511561168257602060c09101515160001a1061167d57600190565b600090565b50600090565b601581510361016e576116a26001600160a01b03916116fd565b1690565b60bf1981019190821161087757565b607f1981019190821161087757565b602003906020821161087757565b60f61981019190821161087757565b60b61981019190821161087757565b9190820391821161087757565b8051801515908161179a575b501561016e57602081019061171e82516118f0565b80825110611755576117348161173b93516116f0565b92516109ff565b519060208110611749575090565b6020036101000a900490565b60405162461bcd60e51b815260206004820152601a60248201527f6c656e677468206973206c657373207468616e206f66667365740000000000006044820152606490fd5b60219150111538611709565b80511561168257600090602081019081516117c0816118f0565b8101809111610877579151905181018091116108775791905b8281106117e65750905090565b6117ef81611854565b8101809111610877576118029091610868565b906117d9565b1561180f57565b60405162461bcd60e51b815260206004820152601160248201527f6164646974696f6e206f766572666c6f770000000000000000000000000000006044820152606490fd5b805160001a90608082101561186a575050600190565b60b88210156118855750611880610a21916116b5565b6109c7565b60c08210156118b45760010151602082900360b7016101000a900490810160b5190190610a2190821015611808565b60f88210156118ca5750611880610a21916116a6565b60010151602082900360f7016101000a900490810160f5190190610a2190821015611808565b5160001a60808110156119035750600090565b60b88110801561193a575b156119195750600190565b60c081101561192e57611880610a21916116e1565b611880610a21916116d2565b5060c0811015801561190e575060f8811061190e565b908151811015610311570160200190565b8051600181149081611986575b50610a215780611981610a219251611a94565b611a1b565b905015610311576020810151607f60f81b6001600160f81b031990911611153861196e565b90815115611a1157815115610311576020820151600180805b6119db575b5050610a219192506119818151611b63565b90918451831015611a0b576119fe611a04916119f78588610889565b5190611a1b565b92610868565b90806119c4565b916119c9565b9050610a21610aab565b6040519181518084526020808501918501928184019282808701915b858110611a845750505080518093875182018852940193828086019201905b828110611a75575050505090603f91601f199351011501011660405290565b81518152908301908301611a56565b8251815291810191849101611a37565b68010000000000000000811015611b1e57611aad610ac3565b906037811115611aff5790611ac4610a2192611bb6565b90611aed611add61126a611ad885516109e3565b6109d5565b60f81b6001600160f81b03191690565b60001a611af98261087c565b53611a1b565b611add61126a611b0e926109e3565b60001a611b1a8261087c565b5390565b60405162461bcd60e51b815260206004820152600e60248201527f696e70757420746f6f206c6f6e670000000000000000000000000000000000006044820152606490fd5b68010000000000000000811015611b1e57611b7c610ac3565b906037811115611ba75790611b93610a2192611bb6565b90611aed611add61126a611ad885516109f1565b611add61126a611b0e926109f1565b611bbe610add565b60208082018390529190600067ffffffffffffffff198316611c4b57506018905b838210611c06575b50611bf39192506116c4565b611bfc81610af6565b9181830152815290565b90611c32611c25611c178385611950565b516001600160f81b03191690565b6001600160f81b03191690565b611c4557611c3f90610868565b90611bdf565b90611be7565b6fffffffffffffffffffffffffffffffff198316611c6f57506010925b9290611bdf565b92611c6856fea26469706673582212201ffbb583a28709d8d2d66417801d6eb7ab9b51979fc5b3ce42d559c4844d8cbf64736f6c63430008120033",
}

// BSCValidatorSetABI is the input ABI used to generate the binding from.
// Deprecated: Use BSCValidatorSetMetaData.ABI instead.
var BSCValidatorSetABI = BSCValidatorSetMetaData.ABI

// BSCValidatorSetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BSCValidatorSetMetaData.Bin instead.
var BSCValidatorSetBin = BSCValidatorSetMetaData.Bin

// DeployBSCValidatorSet deploys a new Ethereum contract, binding an instance of BSCValidatorSet to it.
func DeployBSCValidatorSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BSCValidatorSet, error) {
	parsed, err := BSCValidatorSetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BSCValidatorSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BSCValidatorSet{BSCValidatorSetCaller: BSCValidatorSetCaller{contract: contract}, BSCValidatorSetTransactor: BSCValidatorSetTransactor{contract: contract}, BSCValidatorSetFilterer: BSCValidatorSetFilterer{contract: contract}}, nil
}

// BSCValidatorSet is an auto generated Go binding around an Ethereum contract.
type BSCValidatorSet struct {
	BSCValidatorSetCaller     // Read-only binding to the contract
	BSCValidatorSetTransactor // Write-only binding to the contract
	BSCValidatorSetFilterer   // Log filterer for contract events
}

// BSCValidatorSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type BSCValidatorSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCValidatorSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BSCValidatorSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCValidatorSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BSCValidatorSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCValidatorSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BSCValidatorSetSession struct {
	Contract     *BSCValidatorSet  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BSCValidatorSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BSCValidatorSetCallerSession struct {
	Contract *BSCValidatorSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BSCValidatorSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BSCValidatorSetTransactorSession struct {
	Contract     *BSCValidatorSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BSCValidatorSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type BSCValidatorSetRaw struct {
	Contract *BSCValidatorSet // Generic contract binding to access the raw methods on
}

// BSCValidatorSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BSCValidatorSetCallerRaw struct {
	Contract *BSCValidatorSetCaller // Generic read-only contract binding to access the raw methods on
}

// BSCValidatorSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BSCValidatorSetTransactorRaw struct {
	Contract *BSCValidatorSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBSCValidatorSet creates a new instance of BSCValidatorSet, bound to a specific deployed contract.
func NewBSCValidatorSet(address common.Address, backend bind.ContractBackend) (*BSCValidatorSet, error) {
	contract, err := bindBSCValidatorSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSet{BSCValidatorSetCaller: BSCValidatorSetCaller{contract: contract}, BSCValidatorSetTransactor: BSCValidatorSetTransactor{contract: contract}, BSCValidatorSetFilterer: BSCValidatorSetFilterer{contract: contract}}, nil
}

// NewBSCValidatorSetCaller creates a new read-only instance of BSCValidatorSet, bound to a specific deployed contract.
func NewBSCValidatorSetCaller(address common.Address, caller bind.ContractCaller) (*BSCValidatorSetCaller, error) {
	contract, err := bindBSCValidatorSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetCaller{contract: contract}, nil
}

// NewBSCValidatorSetTransactor creates a new write-only instance of BSCValidatorSet, bound to a specific deployed contract.
func NewBSCValidatorSetTransactor(address common.Address, transactor bind.ContractTransactor) (*BSCValidatorSetTransactor, error) {
	contract, err := bindBSCValidatorSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetTransactor{contract: contract}, nil
}

// NewBSCValidatorSetFilterer creates a new log filterer instance of BSCValidatorSet, bound to a specific deployed contract.
func NewBSCValidatorSetFilterer(address common.Address, filterer bind.ContractFilterer) (*BSCValidatorSetFilterer, error) {
	contract, err := bindBSCValidatorSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetFilterer{contract: contract}, nil
}

// bindBSCValidatorSet binds a generic wrapper to an already deployed contract.
func bindBSCValidatorSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BSCValidatorSetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BSCValidatorSet *BSCValidatorSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BSCValidatorSet.Contract.BSCValidatorSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BSCValidatorSet *BSCValidatorSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.BSCValidatorSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BSCValidatorSet *BSCValidatorSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.BSCValidatorSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BSCValidatorSet *BSCValidatorSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BSCValidatorSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BSCValidatorSet *BSCValidatorSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BSCValidatorSet *BSCValidatorSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.contract.Transact(opts, method, params...)
}

// EPOCH is a free data retrieval call binding the contract method 0xa0dc2758.
//
// Solidity: function EPOCH() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) EPOCH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "EPOCH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EPOCH is a free data retrieval call binding the contract method 0xa0dc2758.
//
// Solidity: function EPOCH() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) EPOCH() (*big.Int, error) {
	return _BSCValidatorSet.Contract.EPOCH(&_BSCValidatorSet.CallOpts)
}

// EPOCH is a free data retrieval call binding the contract method 0xa0dc2758.
//
// Solidity: function EPOCH() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) EPOCH() (*big.Int, error) {
	return _BSCValidatorSet.Contract.EPOCH(&_BSCValidatorSet.CallOpts)
}

// ERRORFAILCHECKVALIDATORS is a free data retrieval call binding the contract method 0x81650b62.
//
// Solidity: function ERROR_FAIL_CHECK_VALIDATORS() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCaller) ERRORFAILCHECKVALIDATORS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "ERROR_FAIL_CHECK_VALIDATORS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORFAILCHECKVALIDATORS is a free data retrieval call binding the contract method 0x81650b62.
//
// Solidity: function ERROR_FAIL_CHECK_VALIDATORS() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetSession) ERRORFAILCHECKVALIDATORS() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORFAILCHECKVALIDATORS(&_BSCValidatorSet.CallOpts)
}

// ERRORFAILCHECKVALIDATORS is a free data retrieval call binding the contract method 0x81650b62.
//
// Solidity: function ERROR_FAIL_CHECK_VALIDATORS() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) ERRORFAILCHECKVALIDATORS() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORFAILCHECKVALIDATORS(&_BSCValidatorSet.CallOpts)
}

// ERRORLENOFVALMISMATCH is a free data retrieval call binding the contract method 0x5d77156c.
//
// Solidity: function ERROR_LEN_OF_VAL_MISMATCH() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCaller) ERRORLENOFVALMISMATCH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "ERROR_LEN_OF_VAL_MISMATCH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORLENOFVALMISMATCH is a free data retrieval call binding the contract method 0x5d77156c.
//
// Solidity: function ERROR_LEN_OF_VAL_MISMATCH() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetSession) ERRORLENOFVALMISMATCH() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORLENOFVALMISMATCH(&_BSCValidatorSet.CallOpts)
}

// ERRORLENOFVALMISMATCH is a free data retrieval call binding the contract method 0x5d77156c.
//
// Solidity: function ERROR_LEN_OF_VAL_MISMATCH() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) ERRORLENOFVALMISMATCH() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORLENOFVALMISMATCH(&_BSCValidatorSet.CallOpts)
}

// ERRORUNKNOWNPACKAGETYPE is a free data retrieval call binding the contract method 0xeda5868c.
//
// Solidity: function ERROR_UNKNOWN_PACKAGE_TYPE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCaller) ERRORUNKNOWNPACKAGETYPE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "ERROR_UNKNOWN_PACKAGE_TYPE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORUNKNOWNPACKAGETYPE is a free data retrieval call binding the contract method 0xeda5868c.
//
// Solidity: function ERROR_UNKNOWN_PACKAGE_TYPE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetSession) ERRORUNKNOWNPACKAGETYPE() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORUNKNOWNPACKAGETYPE(&_BSCValidatorSet.CallOpts)
}

// ERRORUNKNOWNPACKAGETYPE is a free data retrieval call binding the contract method 0xeda5868c.
//
// Solidity: function ERROR_UNKNOWN_PACKAGE_TYPE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) ERRORUNKNOWNPACKAGETYPE() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORUNKNOWNPACKAGETYPE(&_BSCValidatorSet.CallOpts)
}

// EXPIRETIMESECONDGAP is a free data retrieval call binding the contract method 0x853230aa.
//
// Solidity: function EXPIRE_TIME_SECOND_GAP() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) EXPIRETIMESECONDGAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "EXPIRE_TIME_SECOND_GAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXPIRETIMESECONDGAP is a free data retrieval call binding the contract method 0x853230aa.
//
// Solidity: function EXPIRE_TIME_SECOND_GAP() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) EXPIRETIMESECONDGAP() (*big.Int, error) {
	return _BSCValidatorSet.Contract.EXPIRETIMESECONDGAP(&_BSCValidatorSet.CallOpts)
}

// EXPIRETIMESECONDGAP is a free data retrieval call binding the contract method 0x853230aa.
//
// Solidity: function EXPIRE_TIME_SECOND_GAP() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) EXPIRETIMESECONDGAP() (*big.Int, error) {
	return _BSCValidatorSet.Contract.EXPIRETIMESECONDGAP(&_BSCValidatorSet.CallOpts)
}

// INITNUMOFCABINETS is a free data retrieval call binding the contract method 0xb8cf4ef1.
//
// Solidity: function INIT_NUM_OF_CABINETS() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) INITNUMOFCABINETS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "INIT_NUM_OF_CABINETS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITNUMOFCABINETS is a free data retrieval call binding the contract method 0xb8cf4ef1.
//
// Solidity: function INIT_NUM_OF_CABINETS() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) INITNUMOFCABINETS() (*big.Int, error) {
	return _BSCValidatorSet.Contract.INITNUMOFCABINETS(&_BSCValidatorSet.CallOpts)
}

// INITNUMOFCABINETS is a free data retrieval call binding the contract method 0xb8cf4ef1.
//
// Solidity: function INIT_NUM_OF_CABINETS() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) INITNUMOFCABINETS() (*big.Int, error) {
	return _BSCValidatorSet.Contract.INITNUMOFCABINETS(&_BSCValidatorSet.CallOpts)
}

// MAXNUMOFVALIDATORS is a free data retrieval call binding the contract method 0xe086c7b1.
//
// Solidity: function MAX_NUM_OF_VALIDATORS() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) MAXNUMOFVALIDATORS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "MAX_NUM_OF_VALIDATORS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXNUMOFVALIDATORS is a free data retrieval call binding the contract method 0xe086c7b1.
//
// Solidity: function MAX_NUM_OF_VALIDATORS() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) MAXNUMOFVALIDATORS() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MAXNUMOFVALIDATORS(&_BSCValidatorSet.CallOpts)
}

// MAXNUMOFVALIDATORS is a free data retrieval call binding the contract method 0xe086c7b1.
//
// Solidity: function MAX_NUM_OF_VALIDATORS() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) MAXNUMOFVALIDATORS() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MAXNUMOFVALIDATORS(&_BSCValidatorSet.CallOpts)
}

// VALIDATORSUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x5667515a.
//
// Solidity: function VALIDATORS_UPDATE_MESSAGE_TYPE() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCaller) VALIDATORSUPDATEMESSAGETYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "VALIDATORS_UPDATE_MESSAGE_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VALIDATORSUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x5667515a.
//
// Solidity: function VALIDATORS_UPDATE_MESSAGE_TYPE() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetSession) VALIDATORSUPDATEMESSAGETYPE() (uint8, error) {
	return _BSCValidatorSet.Contract.VALIDATORSUPDATEMESSAGETYPE(&_BSCValidatorSet.CallOpts)
}

// VALIDATORSUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x5667515a.
//
// Solidity: function VALIDATORS_UPDATE_MESSAGE_TYPE() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) VALIDATORSUPDATEMESSAGETYPE() (uint8, error) {
	return _BSCValidatorSet.Contract.VALIDATORSUPDATEMESSAGETYPE(&_BSCValidatorSet.CallOpts)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address consensusAddress, address feeAddress, address BBCFeeAddress, uint64 votingPower, bool jailed, uint256 incoming)
func (_BSCValidatorSet *BSCValidatorSetCaller) CurrentValidatorSet(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	BBCFeeAddress    common.Address
	VotingPower      uint64
	Jailed           bool
	Incoming         *big.Int
}, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "currentValidatorSet", arg0)

	outstruct := new(struct {
		ConsensusAddress common.Address
		FeeAddress       common.Address
		BBCFeeAddress    common.Address
		VotingPower      uint64
		Jailed           bool
		Incoming         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConsensusAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.FeeAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BBCFeeAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.VotingPower = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.Jailed = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Incoming = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address consensusAddress, address feeAddress, address BBCFeeAddress, uint64 votingPower, bool jailed, uint256 incoming)
func (_BSCValidatorSet *BSCValidatorSetSession) CurrentValidatorSet(arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	BBCFeeAddress    common.Address
	VotingPower      uint64
	Jailed           bool
	Incoming         *big.Int
}, error) {
	return _BSCValidatorSet.Contract.CurrentValidatorSet(&_BSCValidatorSet.CallOpts, arg0)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address consensusAddress, address feeAddress, address BBCFeeAddress, uint64 votingPower, bool jailed, uint256 incoming)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) CurrentValidatorSet(arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	BBCFeeAddress    common.Address
	VotingPower      uint64
	Jailed           bool
	Incoming         *big.Int
}, error) {
	return _BSCValidatorSet.Contract.CurrentValidatorSet(&_BSCValidatorSet.CallOpts, arg0)
}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) CurrentValidatorSetMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "currentValidatorSetMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) CurrentValidatorSetMap(arg0 common.Address) (*big.Int, error) {
	return _BSCValidatorSet.Contract.CurrentValidatorSetMap(&_BSCValidatorSet.CallOpts, arg0)
}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) CurrentValidatorSetMap(arg0 common.Address) (*big.Int, error) {
	return _BSCValidatorSet.Contract.CurrentValidatorSetMap(&_BSCValidatorSet.CallOpts, arg0)
}

// ExpireTimeSecondGap is a free data retrieval call binding the contract method 0x86249882.
//
// Solidity: function expireTimeSecondGap() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) ExpireTimeSecondGap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "expireTimeSecondGap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpireTimeSecondGap is a free data retrieval call binding the contract method 0x86249882.
//
// Solidity: function expireTimeSecondGap() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) ExpireTimeSecondGap() (*big.Int, error) {
	return _BSCValidatorSet.Contract.ExpireTimeSecondGap(&_BSCValidatorSet.CallOpts)
}

// ExpireTimeSecondGap is a free data retrieval call binding the contract method 0x86249882.
//
// Solidity: function expireTimeSecondGap() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) ExpireTimeSecondGap() (*big.Int, error) {
	return _BSCValidatorSet.Contract.ExpireTimeSecondGap(&_BSCValidatorSet.CallOpts)
}

// IsCurrentValidator is a free data retrieval call binding the contract method 0x55614fcc.
//
// Solidity: function isCurrentValidator(address valAddress) view returns(bool)
func (_BSCValidatorSet *BSCValidatorSetCaller) IsCurrentValidator(opts *bind.CallOpts, valAddress common.Address) (bool, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "isCurrentValidator", valAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurrentValidator is a free data retrieval call binding the contract method 0x55614fcc.
//
// Solidity: function isCurrentValidator(address valAddress) view returns(bool)
func (_BSCValidatorSet *BSCValidatorSetSession) IsCurrentValidator(valAddress common.Address) (bool, error) {
	return _BSCValidatorSet.Contract.IsCurrentValidator(&_BSCValidatorSet.CallOpts, valAddress)
}

// IsCurrentValidator is a free data retrieval call binding the contract method 0x55614fcc.
//
// Solidity: function isCurrentValidator(address valAddress) view returns(bool)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) IsCurrentValidator(valAddress common.Address) (bool, error) {
	return _BSCValidatorSet.Contract.IsCurrentValidator(&_BSCValidatorSet.CallOpts, valAddress)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) HandleAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "handleAckPackage", channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.HandleAckPackage(&_BSCValidatorSet.TransactOpts, channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.HandleAckPackage(&_BSCValidatorSet.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) HandleFailAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "handleFailAckPackage", channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.HandleFailAckPackage(&_BSCValidatorSet.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.HandleFailAckPackage(&_BSCValidatorSet.TransactOpts, channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_BSCValidatorSet *BSCValidatorSetTransactor) HandleSynPackage(opts *bind.TransactOpts, arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "handleSynPackage", arg0, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_BSCValidatorSet *BSCValidatorSetSession) HandleSynPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.HandleSynPackage(&_BSCValidatorSet.TransactOpts, arg0, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) HandleSynPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.HandleSynPackage(&_BSCValidatorSet.TransactOpts, arg0, msgBytes)
}

// Init is a paid mutator transaction binding the contract method 0xc0d91eaf.
//
// Solidity: function init(address _system, bytes _initValidatorSetBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) Init(opts *bind.TransactOpts, _system common.Address, _initValidatorSetBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "init", _system, _initValidatorSetBytes)
}

// Init is a paid mutator transaction binding the contract method 0xc0d91eaf.
//
// Solidity: function init(address _system, bytes _initValidatorSetBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetSession) Init(_system common.Address, _initValidatorSetBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.Init(&_BSCValidatorSet.TransactOpts, _system, _initValidatorSetBytes)
}

// Init is a paid mutator transaction binding the contract method 0xc0d91eaf.
//
// Solidity: function init(address _system, bytes _initValidatorSetBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) Init(_system common.Address, _initValidatorSetBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.Init(&_BSCValidatorSet.TransactOpts, _system, _initValidatorSetBytes)
}

// BSCValidatorSetFailedWithReasonStrIterator is returned from FilterFailedWithReasonStr and is used to iterate over the raw logs and unpacked data for FailedWithReasonStr events raised by the BSCValidatorSet contract.
type BSCValidatorSetFailedWithReasonStrIterator struct {
	Event *BSCValidatorSetFailedWithReasonStr // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetFailedWithReasonStrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetFailedWithReasonStr)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetFailedWithReasonStr)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetFailedWithReasonStrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetFailedWithReasonStrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetFailedWithReasonStr represents a FailedWithReasonStr event raised by the BSCValidatorSet contract.
type BSCValidatorSetFailedWithReasonStr struct {
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFailedWithReasonStr is a free log retrieval operation binding the contract event 0x9e61be46205d080ba82669bf51f14a527ab993061b4555c481f877592f92949f.
//
// Solidity: event FailedWithReasonStr(string message)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterFailedWithReasonStr(opts *bind.FilterOpts) (*BSCValidatorSetFailedWithReasonStrIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "FailedWithReasonStr")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetFailedWithReasonStrIterator{contract: _BSCValidatorSet.contract, event: "FailedWithReasonStr", logs: logs, sub: sub}, nil
}

// WatchFailedWithReasonStr is a free log subscription operation binding the contract event 0x9e61be46205d080ba82669bf51f14a527ab993061b4555c481f877592f92949f.
//
// Solidity: event FailedWithReasonStr(string message)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchFailedWithReasonStr(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetFailedWithReasonStr) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "FailedWithReasonStr")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetFailedWithReasonStr)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "FailedWithReasonStr", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFailedWithReasonStr is a log parse operation binding the contract event 0x9e61be46205d080ba82669bf51f14a527ab993061b4555c481f877592f92949f.
//
// Solidity: event FailedWithReasonStr(string message)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseFailedWithReasonStr(log types.Log) (*BSCValidatorSetFailedWithReasonStr, error) {
	event := new(BSCValidatorSetFailedWithReasonStr)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "FailedWithReasonStr", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetUnexpectedPackageIterator is returned from FilterUnexpectedPackage and is used to iterate over the raw logs and unpacked data for UnexpectedPackage events raised by the BSCValidatorSet contract.
type BSCValidatorSetUnexpectedPackageIterator struct {
	Event *BSCValidatorSetUnexpectedPackage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetUnexpectedPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetUnexpectedPackage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetUnexpectedPackage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetUnexpectedPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetUnexpectedPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetUnexpectedPackage represents a UnexpectedPackage event raised by the BSCValidatorSet contract.
type BSCValidatorSetUnexpectedPackage struct {
	ChannelId uint8
	MsgBytes  []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnexpectedPackage is a free log retrieval operation binding the contract event 0xaa5ba621c8b3d7d05bb9e51a7506108251d4d5dbe542ca66fc7bb52aacb02b65.
//
// Solidity: event UnexpectedPackage(uint8 channelId, bytes msgBytes)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterUnexpectedPackage(opts *bind.FilterOpts) (*BSCValidatorSetUnexpectedPackageIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "UnexpectedPackage")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetUnexpectedPackageIterator{contract: _BSCValidatorSet.contract, event: "UnexpectedPackage", logs: logs, sub: sub}, nil
}

// WatchUnexpectedPackage is a free log subscription operation binding the contract event 0xaa5ba621c8b3d7d05bb9e51a7506108251d4d5dbe542ca66fc7bb52aacb02b65.
//
// Solidity: event UnexpectedPackage(uint8 channelId, bytes msgBytes)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchUnexpectedPackage(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetUnexpectedPackage) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "UnexpectedPackage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetUnexpectedPackage)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "UnexpectedPackage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnexpectedPackage is a log parse operation binding the contract event 0xaa5ba621c8b3d7d05bb9e51a7506108251d4d5dbe542ca66fc7bb52aacb02b65.
//
// Solidity: event UnexpectedPackage(uint8 channelId, bytes msgBytes)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseUnexpectedPackage(log types.Log) (*BSCValidatorSetUnexpectedPackage, error) {
	event := new(BSCValidatorSetUnexpectedPackage)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "UnexpectedPackage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetValidatorSetUpdatedIterator is returned from FilterValidatorSetUpdated and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdated events raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorSetUpdatedIterator struct {
	Event *BSCValidatorSetValidatorSetUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetValidatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetValidatorSetUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetValidatorSetUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetValidatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetValidatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetValidatorSetUpdated represents a ValidatorSetUpdated event raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorSetUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdated is a free log retrieval operation binding the contract event 0xdf380451f7286b929282db56cae856ef8d5d940e01a022844d086bbedc51f755.
//
// Solidity: event ValidatorSetUpdated()
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterValidatorSetUpdated(opts *bind.FilterOpts) (*BSCValidatorSetValidatorSetUpdatedIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "ValidatorSetUpdated")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetValidatorSetUpdatedIterator{contract: _BSCValidatorSet.contract, event: "ValidatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdated is a free log subscription operation binding the contract event 0xdf380451f7286b929282db56cae856ef8d5d940e01a022844d086bbedc51f755.
//
// Solidity: event ValidatorSetUpdated()
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchValidatorSetUpdated(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetValidatorSetUpdated) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "ValidatorSetUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetValidatorSetUpdated)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorSetUpdated is a log parse operation binding the contract event 0xdf380451f7286b929282db56cae856ef8d5d940e01a022844d086bbedc51f755.
//
// Solidity: event ValidatorSetUpdated()
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseValidatorSetUpdated(log types.Log) (*BSCValidatorSetValidatorSetUpdated, error) {
	event := new(BSCValidatorSetValidatorSetUpdated)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BatchEntryMetaData contains all meta data concerning the BatchEntry contract.
var BatchEntryMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea264697066735822122085bacec60c2122486c43e0d778a11f918f454b14c9d84db06491d2e2976fe85d64736f6c63430008120033",
}

// BatchEntryABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchEntryMetaData.ABI instead.
var BatchEntryABI = BatchEntryMetaData.ABI

// BatchEntryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BatchEntryMetaData.Bin instead.
var BatchEntryBin = BatchEntryMetaData.Bin

// DeployBatchEntry deploys a new Ethereum contract, binding an instance of BatchEntry to it.
func DeployBatchEntry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BatchEntry, error) {
	parsed, err := BatchEntryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BatchEntryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BatchEntry{BatchEntryCaller: BatchEntryCaller{contract: contract}, BatchEntryTransactor: BatchEntryTransactor{contract: contract}, BatchEntryFilterer: BatchEntryFilterer{contract: contract}}, nil
}

// BatchEntry is an auto generated Go binding around an Ethereum contract.
type BatchEntry struct {
	BatchEntryCaller     // Read-only binding to the contract
	BatchEntryTransactor // Write-only binding to the contract
	BatchEntryFilterer   // Log filterer for contract events
}

// BatchEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchEntrySession struct {
	Contract     *BatchEntry       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatchEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchEntryCallerSession struct {
	Contract *BatchEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BatchEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchEntryTransactorSession struct {
	Contract     *BatchEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BatchEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchEntryRaw struct {
	Contract *BatchEntry // Generic contract binding to access the raw methods on
}

// BatchEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchEntryCallerRaw struct {
	Contract *BatchEntryCaller // Generic read-only contract binding to access the raw methods on
}

// BatchEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchEntryTransactorRaw struct {
	Contract *BatchEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchEntry creates a new instance of BatchEntry, bound to a specific deployed contract.
func NewBatchEntry(address common.Address, backend bind.ContractBackend) (*BatchEntry, error) {
	contract, err := bindBatchEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchEntry{BatchEntryCaller: BatchEntryCaller{contract: contract}, BatchEntryTransactor: BatchEntryTransactor{contract: contract}, BatchEntryFilterer: BatchEntryFilterer{contract: contract}}, nil
}

// NewBatchEntryCaller creates a new read-only instance of BatchEntry, bound to a specific deployed contract.
func NewBatchEntryCaller(address common.Address, caller bind.ContractCaller) (*BatchEntryCaller, error) {
	contract, err := bindBatchEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchEntryCaller{contract: contract}, nil
}

// NewBatchEntryTransactor creates a new write-only instance of BatchEntry, bound to a specific deployed contract.
func NewBatchEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchEntryTransactor, error) {
	contract, err := bindBatchEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchEntryTransactor{contract: contract}, nil
}

// NewBatchEntryFilterer creates a new log filterer instance of BatchEntry, bound to a specific deployed contract.
func NewBatchEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchEntryFilterer, error) {
	contract, err := bindBatchEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchEntryFilterer{contract: contract}, nil
}

// bindBatchEntry binds a generic wrapper to an already deployed contract.
func bindBatchEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BatchEntryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchEntry *BatchEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchEntry.Contract.BatchEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchEntry *BatchEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchEntry.Contract.BatchEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchEntry *BatchEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchEntry.Contract.BatchEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchEntry *BatchEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchEntry *BatchEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchEntry *BatchEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchEntry.Contract.contract.Transact(opts, method, params...)
}

// BatchProofMetaData contains all meta data concerning the BatchProof contract.
var BatchProofMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220782fbe75a49a5d98b6791ce2a1630d1306a4addf086c1f6fe16761ffedb15a8964736f6c63430008120033",
}

// BatchProofABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchProofMetaData.ABI instead.
var BatchProofABI = BatchProofMetaData.ABI

// BatchProofBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BatchProofMetaData.Bin instead.
var BatchProofBin = BatchProofMetaData.Bin

// DeployBatchProof deploys a new Ethereum contract, binding an instance of BatchProof to it.
func DeployBatchProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BatchProof, error) {
	parsed, err := BatchProofMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BatchProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BatchProof{BatchProofCaller: BatchProofCaller{contract: contract}, BatchProofTransactor: BatchProofTransactor{contract: contract}, BatchProofFilterer: BatchProofFilterer{contract: contract}}, nil
}

// BatchProof is an auto generated Go binding around an Ethereum contract.
type BatchProof struct {
	BatchProofCaller     // Read-only binding to the contract
	BatchProofTransactor // Write-only binding to the contract
	BatchProofFilterer   // Log filterer for contract events
}

// BatchProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchProofSession struct {
	Contract     *BatchProof       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatchProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchProofCallerSession struct {
	Contract *BatchProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BatchProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchProofTransactorSession struct {
	Contract     *BatchProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BatchProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchProofRaw struct {
	Contract *BatchProof // Generic contract binding to access the raw methods on
}

// BatchProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchProofCallerRaw struct {
	Contract *BatchProofCaller // Generic read-only contract binding to access the raw methods on
}

// BatchProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchProofTransactorRaw struct {
	Contract *BatchProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchProof creates a new instance of BatchProof, bound to a specific deployed contract.
func NewBatchProof(address common.Address, backend bind.ContractBackend) (*BatchProof, error) {
	contract, err := bindBatchProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchProof{BatchProofCaller: BatchProofCaller{contract: contract}, BatchProofTransactor: BatchProofTransactor{contract: contract}, BatchProofFilterer: BatchProofFilterer{contract: contract}}, nil
}

// NewBatchProofCaller creates a new read-only instance of BatchProof, bound to a specific deployed contract.
func NewBatchProofCaller(address common.Address, caller bind.ContractCaller) (*BatchProofCaller, error) {
	contract, err := bindBatchProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchProofCaller{contract: contract}, nil
}

// NewBatchProofTransactor creates a new write-only instance of BatchProof, bound to a specific deployed contract.
func NewBatchProofTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchProofTransactor, error) {
	contract, err := bindBatchProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchProofTransactor{contract: contract}, nil
}

// NewBatchProofFilterer creates a new log filterer instance of BatchProof, bound to a specific deployed contract.
func NewBatchProofFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchProofFilterer, error) {
	contract, err := bindBatchProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchProofFilterer{contract: contract}, nil
}

// bindBatchProof binds a generic wrapper to an already deployed contract.
func bindBatchProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BatchProofMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchProof *BatchProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchProof.Contract.BatchProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchProof *BatchProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchProof.Contract.BatchProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchProof *BatchProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchProof.Contract.BatchProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchProof *BatchProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchProof *BatchProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchProof *BatchProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchProof.Contract.contract.Transact(opts, method, params...)
}

// BlockIDMetaData contains all meta data concerning the BlockID contract.
var BlockIDMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea264697066735822122070d44cef53241cc98795bebc4041acfbdac6ef3014ca86f74e31793e3d36a02964736f6c63430008120033",
}

// BlockIDABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockIDMetaData.ABI instead.
var BlockIDABI = BlockIDMetaData.ABI

// BlockIDBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BlockIDMetaData.Bin instead.
var BlockIDBin = BlockIDMetaData.Bin

// DeployBlockID deploys a new Ethereum contract, binding an instance of BlockID to it.
func DeployBlockID(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BlockID, error) {
	parsed, err := BlockIDMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlockIDBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BlockID{BlockIDCaller: BlockIDCaller{contract: contract}, BlockIDTransactor: BlockIDTransactor{contract: contract}, BlockIDFilterer: BlockIDFilterer{contract: contract}}, nil
}

// BlockID is an auto generated Go binding around an Ethereum contract.
type BlockID struct {
	BlockIDCaller     // Read-only binding to the contract
	BlockIDTransactor // Write-only binding to the contract
	BlockIDFilterer   // Log filterer for contract events
}

// BlockIDCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockIDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockIDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockIDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockIDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockIDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockIDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockIDSession struct {
	Contract     *BlockID          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockIDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockIDCallerSession struct {
	Contract *BlockIDCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BlockIDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockIDTransactorSession struct {
	Contract     *BlockIDTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BlockIDRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockIDRaw struct {
	Contract *BlockID // Generic contract binding to access the raw methods on
}

// BlockIDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockIDCallerRaw struct {
	Contract *BlockIDCaller // Generic read-only contract binding to access the raw methods on
}

// BlockIDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockIDTransactorRaw struct {
	Contract *BlockIDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockID creates a new instance of BlockID, bound to a specific deployed contract.
func NewBlockID(address common.Address, backend bind.ContractBackend) (*BlockID, error) {
	contract, err := bindBlockID(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockID{BlockIDCaller: BlockIDCaller{contract: contract}, BlockIDTransactor: BlockIDTransactor{contract: contract}, BlockIDFilterer: BlockIDFilterer{contract: contract}}, nil
}

// NewBlockIDCaller creates a new read-only instance of BlockID, bound to a specific deployed contract.
func NewBlockIDCaller(address common.Address, caller bind.ContractCaller) (*BlockIDCaller, error) {
	contract, err := bindBlockID(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockIDCaller{contract: contract}, nil
}

// NewBlockIDTransactor creates a new write-only instance of BlockID, bound to a specific deployed contract.
func NewBlockIDTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockIDTransactor, error) {
	contract, err := bindBlockID(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockIDTransactor{contract: contract}, nil
}

// NewBlockIDFilterer creates a new log filterer instance of BlockID, bound to a specific deployed contract.
func NewBlockIDFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockIDFilterer, error) {
	contract, err := bindBlockID(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockIDFilterer{contract: contract}, nil
}

// bindBlockID binds a generic wrapper to an already deployed contract.
func bindBlockID(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockIDMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockID *BlockIDRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockID.Contract.BlockIDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockID *BlockIDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockID.Contract.BlockIDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockID *BlockIDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockID.Contract.BlockIDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockID *BlockIDCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockID.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockID *BlockIDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockID.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockID *BlockIDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockID.Contract.contract.Transact(opts, method, params...)
}

// BytesMetaData contains all meta data concerning the Bytes contract.
var BytesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"}],\"name\":\"toBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608080604052346100195760e3908161001f823930815050f35b600080fdfe60806004361015600e57600080fd5b6000803560e01c63326cf61c14602357600080fd5b6020918260031936011260a95760043583820152828152604081019181831067ffffffffffffffff84111760955782604052838352815190816060840152805b8281106082575082820160800152601f01601f19168101036040019150f35b8086918501608083820151910152016063565b634e487b7160e01b81526041600452602490fd5b5080fdfea2646970667358221220681cfa0296f7bef70b42dd9543c2bb032d2b8e5cc72d18e54ca0b504c9a114e064736f6c63430008120033",
}

// BytesABI is the input ABI used to generate the binding from.
// Deprecated: Use BytesMetaData.ABI instead.
var BytesABI = BytesMetaData.ABI

// BytesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BytesMetaData.Bin instead.
var BytesBin = BytesMetaData.Bin

// DeployBytes deploys a new Ethereum contract, binding an instance of Bytes to it.
func DeployBytes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bytes, error) {
	parsed, err := BytesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BytesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bytes{BytesCaller: BytesCaller{contract: contract}, BytesTransactor: BytesTransactor{contract: contract}, BytesFilterer: BytesFilterer{contract: contract}}, nil
}

// Bytes is an auto generated Go binding around an Ethereum contract.
type Bytes struct {
	BytesCaller     // Read-only binding to the contract
	BytesTransactor // Write-only binding to the contract
	BytesFilterer   // Log filterer for contract events
}

// BytesCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesSession struct {
	Contract     *Bytes            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesCallerSession struct {
	Contract *BytesCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BytesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesTransactorSession struct {
	Contract     *BytesTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesRaw struct {
	Contract *Bytes // Generic contract binding to access the raw methods on
}

// BytesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesCallerRaw struct {
	Contract *BytesCaller // Generic read-only contract binding to access the raw methods on
}

// BytesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesTransactorRaw struct {
	Contract *BytesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytes creates a new instance of Bytes, bound to a specific deployed contract.
func NewBytes(address common.Address, backend bind.ContractBackend) (*Bytes, error) {
	contract, err := bindBytes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bytes{BytesCaller: BytesCaller{contract: contract}, BytesTransactor: BytesTransactor{contract: contract}, BytesFilterer: BytesFilterer{contract: contract}}, nil
}

// NewBytesCaller creates a new read-only instance of Bytes, bound to a specific deployed contract.
func NewBytesCaller(address common.Address, caller bind.ContractCaller) (*BytesCaller, error) {
	contract, err := bindBytes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesCaller{contract: contract}, nil
}

// NewBytesTransactor creates a new write-only instance of Bytes, bound to a specific deployed contract.
func NewBytesTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesTransactor, error) {
	contract, err := bindBytes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesTransactor{contract: contract}, nil
}

// NewBytesFilterer creates a new log filterer instance of Bytes, bound to a specific deployed contract.
func NewBytesFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesFilterer, error) {
	contract, err := bindBytes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesFilterer{contract: contract}, nil
}

// bindBytes binds a generic wrapper to an already deployed contract.
func bindBytes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BytesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bytes *BytesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bytes.Contract.BytesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bytes *BytesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bytes.Contract.BytesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bytes *BytesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bytes.Contract.BytesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bytes *BytesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bytes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bytes *BytesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bytes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bytes *BytesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bytes.Contract.contract.Transact(opts, method, params...)
}

// ToBytes is a free data retrieval call binding the contract method 0x326cf61c.
//
// Solidity: function toBytes(bytes32 data) pure returns(bytes)
func (_Bytes *BytesCaller) ToBytes(opts *bind.CallOpts, data [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Bytes.contract.Call(opts, &out, "toBytes", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ToBytes is a free data retrieval call binding the contract method 0x326cf61c.
//
// Solidity: function toBytes(bytes32 data) pure returns(bytes)
func (_Bytes *BytesSession) ToBytes(data [32]byte) ([]byte, error) {
	return _Bytes.Contract.ToBytes(&_Bytes.CallOpts, data)
}

// ToBytes is a free data retrieval call binding the contract method 0x326cf61c.
//
// Solidity: function toBytes(bytes32 data) pure returns(bytes)
func (_Bytes *BytesCallerSession) ToBytes(data [32]byte) ([]byte, error) {
	return _Bytes.Contract.ToBytes(&_Bytes.CallOpts, data)
}

// BytesLibMetaData contains all meta data concerning the BytesLib contract.
var BytesLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212202fae5edda758fe4eb8f499416859616e33037005f70134529732c733b262475464736f6c63430008120033",
}

// BytesLibABI is the input ABI used to generate the binding from.
// Deprecated: Use BytesLibMetaData.ABI instead.
var BytesLibABI = BytesLibMetaData.ABI

// BytesLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BytesLibMetaData.Bin instead.
var BytesLibBin = BytesLibMetaData.Bin

// DeployBytesLib deploys a new Ethereum contract, binding an instance of BytesLib to it.
func DeployBytesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesLib, error) {
	parsed, err := BytesLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BytesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// BytesLib is an auto generated Go binding around an Ethereum contract.
type BytesLib struct {
	BytesLibCaller     // Read-only binding to the contract
	BytesLibTransactor // Write-only binding to the contract
	BytesLibFilterer   // Log filterer for contract events
}

// BytesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesLibSession struct {
	Contract     *BytesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesLibCallerSession struct {
	Contract *BytesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BytesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesLibTransactorSession struct {
	Contract     *BytesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesLibRaw struct {
	Contract *BytesLib // Generic contract binding to access the raw methods on
}

// BytesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesLibCallerRaw struct {
	Contract *BytesLibCaller // Generic read-only contract binding to access the raw methods on
}

// BytesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesLibTransactorRaw struct {
	Contract *BytesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesLib creates a new instance of BytesLib, bound to a specific deployed contract.
func NewBytesLib(address common.Address, backend bind.ContractBackend) (*BytesLib, error) {
	contract, err := bindBytesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// NewBytesLibCaller creates a new read-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibCaller(address common.Address, caller bind.ContractCaller) (*BytesLibCaller, error) {
	contract, err := bindBytesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibCaller{contract: contract}, nil
}

// NewBytesLibTransactor creates a new write-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesLibTransactor, error) {
	contract, err := bindBytesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibTransactor{contract: contract}, nil
}

// NewBytesLibFilterer creates a new log filterer instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesLibFilterer, error) {
	contract, err := bindBytesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesLibFilterer{contract: contract}, nil
}

// bindBytesLib binds a generic wrapper to an already deployed contract.
func bindBytesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BytesLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.BytesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transact(opts, method, params...)
}

// CanonicalBlockIDMetaData contains all meta data concerning the CanonicalBlockID contract.
var CanonicalBlockIDMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220fa69403ec57b0aa98795cf1fd64e9a468b4884e348d8ebdc69d4f52f2b39cd2164736f6c63430008120033",
}

// CanonicalBlockIDABI is the input ABI used to generate the binding from.
// Deprecated: Use CanonicalBlockIDMetaData.ABI instead.
var CanonicalBlockIDABI = CanonicalBlockIDMetaData.ABI

// CanonicalBlockIDBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CanonicalBlockIDMetaData.Bin instead.
var CanonicalBlockIDBin = CanonicalBlockIDMetaData.Bin

// DeployCanonicalBlockID deploys a new Ethereum contract, binding an instance of CanonicalBlockID to it.
func DeployCanonicalBlockID(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CanonicalBlockID, error) {
	parsed, err := CanonicalBlockIDMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CanonicalBlockIDBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CanonicalBlockID{CanonicalBlockIDCaller: CanonicalBlockIDCaller{contract: contract}, CanonicalBlockIDTransactor: CanonicalBlockIDTransactor{contract: contract}, CanonicalBlockIDFilterer: CanonicalBlockIDFilterer{contract: contract}}, nil
}

// CanonicalBlockID is an auto generated Go binding around an Ethereum contract.
type CanonicalBlockID struct {
	CanonicalBlockIDCaller     // Read-only binding to the contract
	CanonicalBlockIDTransactor // Write-only binding to the contract
	CanonicalBlockIDFilterer   // Log filterer for contract events
}

// CanonicalBlockIDCaller is an auto generated read-only Go binding around an Ethereum contract.
type CanonicalBlockIDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanonicalBlockIDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CanonicalBlockIDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanonicalBlockIDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CanonicalBlockIDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanonicalBlockIDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CanonicalBlockIDSession struct {
	Contract     *CanonicalBlockID // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CanonicalBlockIDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CanonicalBlockIDCallerSession struct {
	Contract *CanonicalBlockIDCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CanonicalBlockIDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CanonicalBlockIDTransactorSession struct {
	Contract     *CanonicalBlockIDTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CanonicalBlockIDRaw is an auto generated low-level Go binding around an Ethereum contract.
type CanonicalBlockIDRaw struct {
	Contract *CanonicalBlockID // Generic contract binding to access the raw methods on
}

// CanonicalBlockIDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CanonicalBlockIDCallerRaw struct {
	Contract *CanonicalBlockIDCaller // Generic read-only contract binding to access the raw methods on
}

// CanonicalBlockIDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CanonicalBlockIDTransactorRaw struct {
	Contract *CanonicalBlockIDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCanonicalBlockID creates a new instance of CanonicalBlockID, bound to a specific deployed contract.
func NewCanonicalBlockID(address common.Address, backend bind.ContractBackend) (*CanonicalBlockID, error) {
	contract, err := bindCanonicalBlockID(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CanonicalBlockID{CanonicalBlockIDCaller: CanonicalBlockIDCaller{contract: contract}, CanonicalBlockIDTransactor: CanonicalBlockIDTransactor{contract: contract}, CanonicalBlockIDFilterer: CanonicalBlockIDFilterer{contract: contract}}, nil
}

// NewCanonicalBlockIDCaller creates a new read-only instance of CanonicalBlockID, bound to a specific deployed contract.
func NewCanonicalBlockIDCaller(address common.Address, caller bind.ContractCaller) (*CanonicalBlockIDCaller, error) {
	contract, err := bindCanonicalBlockID(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CanonicalBlockIDCaller{contract: contract}, nil
}

// NewCanonicalBlockIDTransactor creates a new write-only instance of CanonicalBlockID, bound to a specific deployed contract.
func NewCanonicalBlockIDTransactor(address common.Address, transactor bind.ContractTransactor) (*CanonicalBlockIDTransactor, error) {
	contract, err := bindCanonicalBlockID(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CanonicalBlockIDTransactor{contract: contract}, nil
}

// NewCanonicalBlockIDFilterer creates a new log filterer instance of CanonicalBlockID, bound to a specific deployed contract.
func NewCanonicalBlockIDFilterer(address common.Address, filterer bind.ContractFilterer) (*CanonicalBlockIDFilterer, error) {
	contract, err := bindCanonicalBlockID(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CanonicalBlockIDFilterer{contract: contract}, nil
}

// bindCanonicalBlockID binds a generic wrapper to an already deployed contract.
func bindCanonicalBlockID(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CanonicalBlockIDMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CanonicalBlockID *CanonicalBlockIDRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CanonicalBlockID.Contract.CanonicalBlockIDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CanonicalBlockID *CanonicalBlockIDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanonicalBlockID.Contract.CanonicalBlockIDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CanonicalBlockID *CanonicalBlockIDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CanonicalBlockID.Contract.CanonicalBlockIDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CanonicalBlockID *CanonicalBlockIDCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CanonicalBlockID.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CanonicalBlockID *CanonicalBlockIDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanonicalBlockID.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CanonicalBlockID *CanonicalBlockIDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CanonicalBlockID.Contract.contract.Transact(opts, method, params...)
}

// CanonicalPartSetHeaderMetaData contains all meta data concerning the CanonicalPartSetHeader contract.
var CanonicalPartSetHeaderMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212207c0c822cce7e58d3345a2622bf1d8c4e00c04109949697e3d36ff66cab40b4f064736f6c63430008120033",
}

// CanonicalPartSetHeaderABI is the input ABI used to generate the binding from.
// Deprecated: Use CanonicalPartSetHeaderMetaData.ABI instead.
var CanonicalPartSetHeaderABI = CanonicalPartSetHeaderMetaData.ABI

// CanonicalPartSetHeaderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CanonicalPartSetHeaderMetaData.Bin instead.
var CanonicalPartSetHeaderBin = CanonicalPartSetHeaderMetaData.Bin

// DeployCanonicalPartSetHeader deploys a new Ethereum contract, binding an instance of CanonicalPartSetHeader to it.
func DeployCanonicalPartSetHeader(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CanonicalPartSetHeader, error) {
	parsed, err := CanonicalPartSetHeaderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CanonicalPartSetHeaderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CanonicalPartSetHeader{CanonicalPartSetHeaderCaller: CanonicalPartSetHeaderCaller{contract: contract}, CanonicalPartSetHeaderTransactor: CanonicalPartSetHeaderTransactor{contract: contract}, CanonicalPartSetHeaderFilterer: CanonicalPartSetHeaderFilterer{contract: contract}}, nil
}

// CanonicalPartSetHeader is an auto generated Go binding around an Ethereum contract.
type CanonicalPartSetHeader struct {
	CanonicalPartSetHeaderCaller     // Read-only binding to the contract
	CanonicalPartSetHeaderTransactor // Write-only binding to the contract
	CanonicalPartSetHeaderFilterer   // Log filterer for contract events
}

// CanonicalPartSetHeaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type CanonicalPartSetHeaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanonicalPartSetHeaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CanonicalPartSetHeaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanonicalPartSetHeaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CanonicalPartSetHeaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanonicalPartSetHeaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CanonicalPartSetHeaderSession struct {
	Contract     *CanonicalPartSetHeader // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CanonicalPartSetHeaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CanonicalPartSetHeaderCallerSession struct {
	Contract *CanonicalPartSetHeaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// CanonicalPartSetHeaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CanonicalPartSetHeaderTransactorSession struct {
	Contract     *CanonicalPartSetHeaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// CanonicalPartSetHeaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type CanonicalPartSetHeaderRaw struct {
	Contract *CanonicalPartSetHeader // Generic contract binding to access the raw methods on
}

// CanonicalPartSetHeaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CanonicalPartSetHeaderCallerRaw struct {
	Contract *CanonicalPartSetHeaderCaller // Generic read-only contract binding to access the raw methods on
}

// CanonicalPartSetHeaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CanonicalPartSetHeaderTransactorRaw struct {
	Contract *CanonicalPartSetHeaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCanonicalPartSetHeader creates a new instance of CanonicalPartSetHeader, bound to a specific deployed contract.
func NewCanonicalPartSetHeader(address common.Address, backend bind.ContractBackend) (*CanonicalPartSetHeader, error) {
	contract, err := bindCanonicalPartSetHeader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CanonicalPartSetHeader{CanonicalPartSetHeaderCaller: CanonicalPartSetHeaderCaller{contract: contract}, CanonicalPartSetHeaderTransactor: CanonicalPartSetHeaderTransactor{contract: contract}, CanonicalPartSetHeaderFilterer: CanonicalPartSetHeaderFilterer{contract: contract}}, nil
}

// NewCanonicalPartSetHeaderCaller creates a new read-only instance of CanonicalPartSetHeader, bound to a specific deployed contract.
func NewCanonicalPartSetHeaderCaller(address common.Address, caller bind.ContractCaller) (*CanonicalPartSetHeaderCaller, error) {
	contract, err := bindCanonicalPartSetHeader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CanonicalPartSetHeaderCaller{contract: contract}, nil
}

// NewCanonicalPartSetHeaderTransactor creates a new write-only instance of CanonicalPartSetHeader, bound to a specific deployed contract.
func NewCanonicalPartSetHeaderTransactor(address common.Address, transactor bind.ContractTransactor) (*CanonicalPartSetHeaderTransactor, error) {
	contract, err := bindCanonicalPartSetHeader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CanonicalPartSetHeaderTransactor{contract: contract}, nil
}

// NewCanonicalPartSetHeaderFilterer creates a new log filterer instance of CanonicalPartSetHeader, bound to a specific deployed contract.
func NewCanonicalPartSetHeaderFilterer(address common.Address, filterer bind.ContractFilterer) (*CanonicalPartSetHeaderFilterer, error) {
	contract, err := bindCanonicalPartSetHeader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CanonicalPartSetHeaderFilterer{contract: contract}, nil
}

// bindCanonicalPartSetHeader binds a generic wrapper to an already deployed contract.
func bindCanonicalPartSetHeader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CanonicalPartSetHeaderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CanonicalPartSetHeader *CanonicalPartSetHeaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CanonicalPartSetHeader.Contract.CanonicalPartSetHeaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CanonicalPartSetHeader *CanonicalPartSetHeaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanonicalPartSetHeader.Contract.CanonicalPartSetHeaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CanonicalPartSetHeader *CanonicalPartSetHeaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CanonicalPartSetHeader.Contract.CanonicalPartSetHeaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CanonicalPartSetHeader *CanonicalPartSetHeaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CanonicalPartSetHeader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CanonicalPartSetHeader *CanonicalPartSetHeaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanonicalPartSetHeader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CanonicalPartSetHeader *CanonicalPartSetHeaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CanonicalPartSetHeader.Contract.contract.Transact(opts, method, params...)
}

// CanonicalVoteMetaData contains all meta data concerning the CanonicalVote contract.
var CanonicalVoteMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220cf06e0baf440dfa2689739b7eb64facc156b472d936d4acbab965b9e17008ff664736f6c63430008120033",
}

// CanonicalVoteABI is the input ABI used to generate the binding from.
// Deprecated: Use CanonicalVoteMetaData.ABI instead.
var CanonicalVoteABI = CanonicalVoteMetaData.ABI

// CanonicalVoteBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CanonicalVoteMetaData.Bin instead.
var CanonicalVoteBin = CanonicalVoteMetaData.Bin

// DeployCanonicalVote deploys a new Ethereum contract, binding an instance of CanonicalVote to it.
func DeployCanonicalVote(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CanonicalVote, error) {
	parsed, err := CanonicalVoteMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CanonicalVoteBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CanonicalVote{CanonicalVoteCaller: CanonicalVoteCaller{contract: contract}, CanonicalVoteTransactor: CanonicalVoteTransactor{contract: contract}, CanonicalVoteFilterer: CanonicalVoteFilterer{contract: contract}}, nil
}

// CanonicalVote is an auto generated Go binding around an Ethereum contract.
type CanonicalVote struct {
	CanonicalVoteCaller     // Read-only binding to the contract
	CanonicalVoteTransactor // Write-only binding to the contract
	CanonicalVoteFilterer   // Log filterer for contract events
}

// CanonicalVoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type CanonicalVoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanonicalVoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CanonicalVoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanonicalVoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CanonicalVoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanonicalVoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CanonicalVoteSession struct {
	Contract     *CanonicalVote    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CanonicalVoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CanonicalVoteCallerSession struct {
	Contract *CanonicalVoteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CanonicalVoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CanonicalVoteTransactorSession struct {
	Contract     *CanonicalVoteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CanonicalVoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type CanonicalVoteRaw struct {
	Contract *CanonicalVote // Generic contract binding to access the raw methods on
}

// CanonicalVoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CanonicalVoteCallerRaw struct {
	Contract *CanonicalVoteCaller // Generic read-only contract binding to access the raw methods on
}

// CanonicalVoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CanonicalVoteTransactorRaw struct {
	Contract *CanonicalVoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCanonicalVote creates a new instance of CanonicalVote, bound to a specific deployed contract.
func NewCanonicalVote(address common.Address, backend bind.ContractBackend) (*CanonicalVote, error) {
	contract, err := bindCanonicalVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CanonicalVote{CanonicalVoteCaller: CanonicalVoteCaller{contract: contract}, CanonicalVoteTransactor: CanonicalVoteTransactor{contract: contract}, CanonicalVoteFilterer: CanonicalVoteFilterer{contract: contract}}, nil
}

// NewCanonicalVoteCaller creates a new read-only instance of CanonicalVote, bound to a specific deployed contract.
func NewCanonicalVoteCaller(address common.Address, caller bind.ContractCaller) (*CanonicalVoteCaller, error) {
	contract, err := bindCanonicalVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CanonicalVoteCaller{contract: contract}, nil
}

// NewCanonicalVoteTransactor creates a new write-only instance of CanonicalVote, bound to a specific deployed contract.
func NewCanonicalVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*CanonicalVoteTransactor, error) {
	contract, err := bindCanonicalVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CanonicalVoteTransactor{contract: contract}, nil
}

// NewCanonicalVoteFilterer creates a new log filterer instance of CanonicalVote, bound to a specific deployed contract.
func NewCanonicalVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*CanonicalVoteFilterer, error) {
	contract, err := bindCanonicalVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CanonicalVoteFilterer{contract: contract}, nil
}

// bindCanonicalVote binds a generic wrapper to an already deployed contract.
func bindCanonicalVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CanonicalVoteMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CanonicalVote *CanonicalVoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CanonicalVote.Contract.CanonicalVoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CanonicalVote *CanonicalVoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanonicalVote.Contract.CanonicalVoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CanonicalVote *CanonicalVoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CanonicalVote.Contract.CanonicalVoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CanonicalVote *CanonicalVoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CanonicalVote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CanonicalVote *CanonicalVoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanonicalVote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CanonicalVote *CanonicalVoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CanonicalVote.Contract.contract.Transact(opts, method, params...)
}

// ClientStateMetaData contains all meta data concerning the ClientState contract.
var ClientStateMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220527ed53b3f842a15f34a349c27d099ebd88e04d1c5f2e10bef4b8c3387c8f57664736f6c63430008120033",
}

// ClientStateABI is the input ABI used to generate the binding from.
// Deprecated: Use ClientStateMetaData.ABI instead.
var ClientStateABI = ClientStateMetaData.ABI

// ClientStateBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ClientStateMetaData.Bin instead.
var ClientStateBin = ClientStateMetaData.Bin

// DeployClientState deploys a new Ethereum contract, binding an instance of ClientState to it.
func DeployClientState(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ClientState, error) {
	parsed, err := ClientStateMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ClientStateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ClientState{ClientStateCaller: ClientStateCaller{contract: contract}, ClientStateTransactor: ClientStateTransactor{contract: contract}, ClientStateFilterer: ClientStateFilterer{contract: contract}}, nil
}

// ClientState is an auto generated Go binding around an Ethereum contract.
type ClientState struct {
	ClientStateCaller     // Read-only binding to the contract
	ClientStateTransactor // Write-only binding to the contract
	ClientStateFilterer   // Log filterer for contract events
}

// ClientStateCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClientStateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientStateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClientStateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientStateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClientStateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClientStateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClientStateSession struct {
	Contract     *ClientState      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClientStateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClientStateCallerSession struct {
	Contract *ClientStateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ClientStateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClientStateTransactorSession struct {
	Contract     *ClientStateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ClientStateRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClientStateRaw struct {
	Contract *ClientState // Generic contract binding to access the raw methods on
}

// ClientStateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClientStateCallerRaw struct {
	Contract *ClientStateCaller // Generic read-only contract binding to access the raw methods on
}

// ClientStateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClientStateTransactorRaw struct {
	Contract *ClientStateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClientState creates a new instance of ClientState, bound to a specific deployed contract.
func NewClientState(address common.Address, backend bind.ContractBackend) (*ClientState, error) {
	contract, err := bindClientState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClientState{ClientStateCaller: ClientStateCaller{contract: contract}, ClientStateTransactor: ClientStateTransactor{contract: contract}, ClientStateFilterer: ClientStateFilterer{contract: contract}}, nil
}

// NewClientStateCaller creates a new read-only instance of ClientState, bound to a specific deployed contract.
func NewClientStateCaller(address common.Address, caller bind.ContractCaller) (*ClientStateCaller, error) {
	contract, err := bindClientState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClientStateCaller{contract: contract}, nil
}

// NewClientStateTransactor creates a new write-only instance of ClientState, bound to a specific deployed contract.
func NewClientStateTransactor(address common.Address, transactor bind.ContractTransactor) (*ClientStateTransactor, error) {
	contract, err := bindClientState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClientStateTransactor{contract: contract}, nil
}

// NewClientStateFilterer creates a new log filterer instance of ClientState, bound to a specific deployed contract.
func NewClientStateFilterer(address common.Address, filterer bind.ContractFilterer) (*ClientStateFilterer, error) {
	contract, err := bindClientState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClientStateFilterer{contract: contract}, nil
}

// bindClientState binds a generic wrapper to an already deployed contract.
func bindClientState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClientStateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClientState *ClientStateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClientState.Contract.ClientStateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClientState *ClientStateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClientState.Contract.ClientStateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClientState *ClientStateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClientState.Contract.ClientStateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClientState *ClientStateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClientState.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClientState *ClientStateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClientState.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClientState *ClientStateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClientState.Contract.contract.Transact(opts, method, params...)
}

// CmnPkgMetaData contains all meta data concerning the CmnPkg contract.
var CmnPkgMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea264697066735822122048c233865084504cbd81f52c045a98bda9cb32e48264b04eeb5cb5bc1a5d08fd64736f6c63430008120033",
}

// CmnPkgABI is the input ABI used to generate the binding from.
// Deprecated: Use CmnPkgMetaData.ABI instead.
var CmnPkgABI = CmnPkgMetaData.ABI

// CmnPkgBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CmnPkgMetaData.Bin instead.
var CmnPkgBin = CmnPkgMetaData.Bin

// DeployCmnPkg deploys a new Ethereum contract, binding an instance of CmnPkg to it.
func DeployCmnPkg(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CmnPkg, error) {
	parsed, err := CmnPkgMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CmnPkgBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CmnPkg{CmnPkgCaller: CmnPkgCaller{contract: contract}, CmnPkgTransactor: CmnPkgTransactor{contract: contract}, CmnPkgFilterer: CmnPkgFilterer{contract: contract}}, nil
}

// CmnPkg is an auto generated Go binding around an Ethereum contract.
type CmnPkg struct {
	CmnPkgCaller     // Read-only binding to the contract
	CmnPkgTransactor // Write-only binding to the contract
	CmnPkgFilterer   // Log filterer for contract events
}

// CmnPkgCaller is an auto generated read-only Go binding around an Ethereum contract.
type CmnPkgCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CmnPkgTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CmnPkgTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CmnPkgFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CmnPkgFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CmnPkgSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CmnPkgSession struct {
	Contract     *CmnPkg           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CmnPkgCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CmnPkgCallerSession struct {
	Contract *CmnPkgCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CmnPkgTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CmnPkgTransactorSession struct {
	Contract     *CmnPkgTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CmnPkgRaw is an auto generated low-level Go binding around an Ethereum contract.
type CmnPkgRaw struct {
	Contract *CmnPkg // Generic contract binding to access the raw methods on
}

// CmnPkgCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CmnPkgCallerRaw struct {
	Contract *CmnPkgCaller // Generic read-only contract binding to access the raw methods on
}

// CmnPkgTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CmnPkgTransactorRaw struct {
	Contract *CmnPkgTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCmnPkg creates a new instance of CmnPkg, bound to a specific deployed contract.
func NewCmnPkg(address common.Address, backend bind.ContractBackend) (*CmnPkg, error) {
	contract, err := bindCmnPkg(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CmnPkg{CmnPkgCaller: CmnPkgCaller{contract: contract}, CmnPkgTransactor: CmnPkgTransactor{contract: contract}, CmnPkgFilterer: CmnPkgFilterer{contract: contract}}, nil
}

// NewCmnPkgCaller creates a new read-only instance of CmnPkg, bound to a specific deployed contract.
func NewCmnPkgCaller(address common.Address, caller bind.ContractCaller) (*CmnPkgCaller, error) {
	contract, err := bindCmnPkg(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CmnPkgCaller{contract: contract}, nil
}

// NewCmnPkgTransactor creates a new write-only instance of CmnPkg, bound to a specific deployed contract.
func NewCmnPkgTransactor(address common.Address, transactor bind.ContractTransactor) (*CmnPkgTransactor, error) {
	contract, err := bindCmnPkg(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CmnPkgTransactor{contract: contract}, nil
}

// NewCmnPkgFilterer creates a new log filterer instance of CmnPkg, bound to a specific deployed contract.
func NewCmnPkgFilterer(address common.Address, filterer bind.ContractFilterer) (*CmnPkgFilterer, error) {
	contract, err := bindCmnPkg(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CmnPkgFilterer{contract: contract}, nil
}

// bindCmnPkg binds a generic wrapper to an already deployed contract.
func bindCmnPkg(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CmnPkgMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CmnPkg *CmnPkgRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CmnPkg.Contract.CmnPkgCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CmnPkg *CmnPkgRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CmnPkg.Contract.CmnPkgTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CmnPkg *CmnPkgRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CmnPkg.Contract.CmnPkgTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CmnPkg *CmnPkgCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CmnPkg.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CmnPkg *CmnPkgTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CmnPkg.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CmnPkg *CmnPkgTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CmnPkg.Contract.contract.Transact(opts, method, params...)
}

// CommitMetaData contains all meta data concerning the Commit contract.
var CommitMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220763521df57664ee793d08632de791ac45134ed7940226201516e35036468cb6764736f6c63430008120033",
}

// CommitABI is the input ABI used to generate the binding from.
// Deprecated: Use CommitMetaData.ABI instead.
var CommitABI = CommitMetaData.ABI

// CommitBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CommitMetaData.Bin instead.
var CommitBin = CommitMetaData.Bin

// DeployCommit deploys a new Ethereum contract, binding an instance of Commit to it.
func DeployCommit(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Commit, error) {
	parsed, err := CommitMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CommitBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Commit{CommitCaller: CommitCaller{contract: contract}, CommitTransactor: CommitTransactor{contract: contract}, CommitFilterer: CommitFilterer{contract: contract}}, nil
}

// Commit is an auto generated Go binding around an Ethereum contract.
type Commit struct {
	CommitCaller     // Read-only binding to the contract
	CommitTransactor // Write-only binding to the contract
	CommitFilterer   // Log filterer for contract events
}

// CommitCaller is an auto generated read-only Go binding around an Ethereum contract.
type CommitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CommitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CommitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CommitSession struct {
	Contract     *Commit           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CommitCallerSession struct {
	Contract *CommitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CommitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CommitTransactorSession struct {
	Contract     *CommitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommitRaw is an auto generated low-level Go binding around an Ethereum contract.
type CommitRaw struct {
	Contract *Commit // Generic contract binding to access the raw methods on
}

// CommitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CommitCallerRaw struct {
	Contract *CommitCaller // Generic read-only contract binding to access the raw methods on
}

// CommitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CommitTransactorRaw struct {
	Contract *CommitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCommit creates a new instance of Commit, bound to a specific deployed contract.
func NewCommit(address common.Address, backend bind.ContractBackend) (*Commit, error) {
	contract, err := bindCommit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Commit{CommitCaller: CommitCaller{contract: contract}, CommitTransactor: CommitTransactor{contract: contract}, CommitFilterer: CommitFilterer{contract: contract}}, nil
}

// NewCommitCaller creates a new read-only instance of Commit, bound to a specific deployed contract.
func NewCommitCaller(address common.Address, caller bind.ContractCaller) (*CommitCaller, error) {
	contract, err := bindCommit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitCaller{contract: contract}, nil
}

// NewCommitTransactor creates a new write-only instance of Commit, bound to a specific deployed contract.
func NewCommitTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitTransactor, error) {
	contract, err := bindCommit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitTransactor{contract: contract}, nil
}

// NewCommitFilterer creates a new log filterer instance of Commit, bound to a specific deployed contract.
func NewCommitFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitFilterer, error) {
	contract, err := bindCommit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitFilterer{contract: contract}, nil
}

// bindCommit binds a generic wrapper to an already deployed contract.
func bindCommit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommitMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Commit *CommitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Commit.Contract.CommitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Commit *CommitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commit.Contract.CommitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Commit *CommitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Commit.Contract.CommitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Commit *CommitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Commit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Commit *CommitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Commit *CommitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Commit.Contract.contract.Transact(opts, method, params...)
}

// CommitSigMetaData contains all meta data concerning the CommitSig contract.
var CommitSigMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212202c56f5e88588226fc35df2975148cd3b01d5be461a962dc051b6724c3a8308d864736f6c63430008120033",
}

// CommitSigABI is the input ABI used to generate the binding from.
// Deprecated: Use CommitSigMetaData.ABI instead.
var CommitSigABI = CommitSigMetaData.ABI

// CommitSigBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CommitSigMetaData.Bin instead.
var CommitSigBin = CommitSigMetaData.Bin

// DeployCommitSig deploys a new Ethereum contract, binding an instance of CommitSig to it.
func DeployCommitSig(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CommitSig, error) {
	parsed, err := CommitSigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CommitSigBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CommitSig{CommitSigCaller: CommitSigCaller{contract: contract}, CommitSigTransactor: CommitSigTransactor{contract: contract}, CommitSigFilterer: CommitSigFilterer{contract: contract}}, nil
}

// CommitSig is an auto generated Go binding around an Ethereum contract.
type CommitSig struct {
	CommitSigCaller     // Read-only binding to the contract
	CommitSigTransactor // Write-only binding to the contract
	CommitSigFilterer   // Log filterer for contract events
}

// CommitSigCaller is an auto generated read-only Go binding around an Ethereum contract.
type CommitSigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitSigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CommitSigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitSigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CommitSigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitSigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CommitSigSession struct {
	Contract     *CommitSig        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommitSigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CommitSigCallerSession struct {
	Contract *CommitSigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CommitSigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CommitSigTransactorSession struct {
	Contract     *CommitSigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CommitSigRaw is an auto generated low-level Go binding around an Ethereum contract.
type CommitSigRaw struct {
	Contract *CommitSig // Generic contract binding to access the raw methods on
}

// CommitSigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CommitSigCallerRaw struct {
	Contract *CommitSigCaller // Generic read-only contract binding to access the raw methods on
}

// CommitSigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CommitSigTransactorRaw struct {
	Contract *CommitSigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCommitSig creates a new instance of CommitSig, bound to a specific deployed contract.
func NewCommitSig(address common.Address, backend bind.ContractBackend) (*CommitSig, error) {
	contract, err := bindCommitSig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CommitSig{CommitSigCaller: CommitSigCaller{contract: contract}, CommitSigTransactor: CommitSigTransactor{contract: contract}, CommitSigFilterer: CommitSigFilterer{contract: contract}}, nil
}

// NewCommitSigCaller creates a new read-only instance of CommitSig, bound to a specific deployed contract.
func NewCommitSigCaller(address common.Address, caller bind.ContractCaller) (*CommitSigCaller, error) {
	contract, err := bindCommitSig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitSigCaller{contract: contract}, nil
}

// NewCommitSigTransactor creates a new write-only instance of CommitSig, bound to a specific deployed contract.
func NewCommitSigTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitSigTransactor, error) {
	contract, err := bindCommitSig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitSigTransactor{contract: contract}, nil
}

// NewCommitSigFilterer creates a new log filterer instance of CommitSig, bound to a specific deployed contract.
func NewCommitSigFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitSigFilterer, error) {
	contract, err := bindCommitSig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitSigFilterer{contract: contract}, nil
}

// bindCommitSig binds a generic wrapper to an already deployed contract.
func bindCommitSig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommitSigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommitSig *CommitSigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitSig.Contract.CommitSigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommitSig *CommitSigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitSig.Contract.CommitSigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommitSig *CommitSigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitSig.Contract.CommitSigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommitSig *CommitSigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitSig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommitSig *CommitSigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitSig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommitSig *CommitSigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitSig.Contract.contract.Transact(opts, method, params...)
}

// CommitmentProofMetaData contains all meta data concerning the CommitmentProof contract.
var CommitmentProofMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea264697066735822122023174c09af142c6a027d5b361ea5b4ff5d823db8297d0cff31184c2e737ae5dd64736f6c63430008120033",
}

// CommitmentProofABI is the input ABI used to generate the binding from.
// Deprecated: Use CommitmentProofMetaData.ABI instead.
var CommitmentProofABI = CommitmentProofMetaData.ABI

// CommitmentProofBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CommitmentProofMetaData.Bin instead.
var CommitmentProofBin = CommitmentProofMetaData.Bin

// DeployCommitmentProof deploys a new Ethereum contract, binding an instance of CommitmentProof to it.
func DeployCommitmentProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CommitmentProof, error) {
	parsed, err := CommitmentProofMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CommitmentProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CommitmentProof{CommitmentProofCaller: CommitmentProofCaller{contract: contract}, CommitmentProofTransactor: CommitmentProofTransactor{contract: contract}, CommitmentProofFilterer: CommitmentProofFilterer{contract: contract}}, nil
}

// CommitmentProof is an auto generated Go binding around an Ethereum contract.
type CommitmentProof struct {
	CommitmentProofCaller     // Read-only binding to the contract
	CommitmentProofTransactor // Write-only binding to the contract
	CommitmentProofFilterer   // Log filterer for contract events
}

// CommitmentProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type CommitmentProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitmentProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CommitmentProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitmentProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CommitmentProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitmentProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CommitmentProofSession struct {
	Contract     *CommitmentProof  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommitmentProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CommitmentProofCallerSession struct {
	Contract *CommitmentProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CommitmentProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CommitmentProofTransactorSession struct {
	Contract     *CommitmentProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CommitmentProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type CommitmentProofRaw struct {
	Contract *CommitmentProof // Generic contract binding to access the raw methods on
}

// CommitmentProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CommitmentProofCallerRaw struct {
	Contract *CommitmentProofCaller // Generic read-only contract binding to access the raw methods on
}

// CommitmentProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CommitmentProofTransactorRaw struct {
	Contract *CommitmentProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCommitmentProof creates a new instance of CommitmentProof, bound to a specific deployed contract.
func NewCommitmentProof(address common.Address, backend bind.ContractBackend) (*CommitmentProof, error) {
	contract, err := bindCommitmentProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CommitmentProof{CommitmentProofCaller: CommitmentProofCaller{contract: contract}, CommitmentProofTransactor: CommitmentProofTransactor{contract: contract}, CommitmentProofFilterer: CommitmentProofFilterer{contract: contract}}, nil
}

// NewCommitmentProofCaller creates a new read-only instance of CommitmentProof, bound to a specific deployed contract.
func NewCommitmentProofCaller(address common.Address, caller bind.ContractCaller) (*CommitmentProofCaller, error) {
	contract, err := bindCommitmentProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitmentProofCaller{contract: contract}, nil
}

// NewCommitmentProofTransactor creates a new write-only instance of CommitmentProof, bound to a specific deployed contract.
func NewCommitmentProofTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitmentProofTransactor, error) {
	contract, err := bindCommitmentProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitmentProofTransactor{contract: contract}, nil
}

// NewCommitmentProofFilterer creates a new log filterer instance of CommitmentProof, bound to a specific deployed contract.
func NewCommitmentProofFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitmentProofFilterer, error) {
	contract, err := bindCommitmentProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitmentProofFilterer{contract: contract}, nil
}

// bindCommitmentProof binds a generic wrapper to an already deployed contract.
func bindCommitmentProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommitmentProofMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommitmentProof *CommitmentProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitmentProof.Contract.CommitmentProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommitmentProof *CommitmentProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitmentProof.Contract.CommitmentProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommitmentProof *CommitmentProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitmentProof.Contract.CommitmentProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommitmentProof *CommitmentProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitmentProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommitmentProof *CommitmentProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitmentProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommitmentProof *CommitmentProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitmentProof.Contract.contract.Transact(opts, method, params...)
}

// CompressMetaData contains all meta data concerning the Compress contract.
var CompressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea264697066735822122023012f67973b5138acd50bc7ec68999d5102714cc56ab468c8a464d337d25fac64736f6c63430008120033",
}

// CompressABI is the input ABI used to generate the binding from.
// Deprecated: Use CompressMetaData.ABI instead.
var CompressABI = CompressMetaData.ABI

// CompressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CompressMetaData.Bin instead.
var CompressBin = CompressMetaData.Bin

// DeployCompress deploys a new Ethereum contract, binding an instance of Compress to it.
func DeployCompress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Compress, error) {
	parsed, err := CompressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CompressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Compress{CompressCaller: CompressCaller{contract: contract}, CompressTransactor: CompressTransactor{contract: contract}, CompressFilterer: CompressFilterer{contract: contract}}, nil
}

// Compress is an auto generated Go binding around an Ethereum contract.
type Compress struct {
	CompressCaller     // Read-only binding to the contract
	CompressTransactor // Write-only binding to the contract
	CompressFilterer   // Log filterer for contract events
}

// CompressCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompressSession struct {
	Contract     *Compress         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CompressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompressCallerSession struct {
	Contract *CompressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CompressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompressTransactorSession struct {
	Contract     *CompressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CompressRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompressRaw struct {
	Contract *Compress // Generic contract binding to access the raw methods on
}

// CompressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompressCallerRaw struct {
	Contract *CompressCaller // Generic read-only contract binding to access the raw methods on
}

// CompressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompressTransactorRaw struct {
	Contract *CompressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompress creates a new instance of Compress, bound to a specific deployed contract.
func NewCompress(address common.Address, backend bind.ContractBackend) (*Compress, error) {
	contract, err := bindCompress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Compress{CompressCaller: CompressCaller{contract: contract}, CompressTransactor: CompressTransactor{contract: contract}, CompressFilterer: CompressFilterer{contract: contract}}, nil
}

// NewCompressCaller creates a new read-only instance of Compress, bound to a specific deployed contract.
func NewCompressCaller(address common.Address, caller bind.ContractCaller) (*CompressCaller, error) {
	contract, err := bindCompress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompressCaller{contract: contract}, nil
}

// NewCompressTransactor creates a new write-only instance of Compress, bound to a specific deployed contract.
func NewCompressTransactor(address common.Address, transactor bind.ContractTransactor) (*CompressTransactor, error) {
	contract, err := bindCompress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompressTransactor{contract: contract}, nil
}

// NewCompressFilterer creates a new log filterer instance of Compress, bound to a specific deployed contract.
func NewCompressFilterer(address common.Address, filterer bind.ContractFilterer) (*CompressFilterer, error) {
	contract, err := bindCompress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompressFilterer{contract: contract}, nil
}

// bindCompress binds a generic wrapper to an already deployed contract.
func bindCompress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CompressMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compress *CompressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compress.Contract.CompressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compress *CompressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compress.Contract.CompressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compress *CompressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compress.Contract.CompressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compress *CompressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compress.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compress *CompressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compress.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compress *CompressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compress.Contract.contract.Transact(opts, method, params...)
}

// CompressedBatchEntryMetaData contains all meta data concerning the CompressedBatchEntry contract.
var CompressedBatchEntryMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220901e45855b5509698913496f77cc94f156b3860a55af96626a735de9f393fe9464736f6c63430008120033",
}

// CompressedBatchEntryABI is the input ABI used to generate the binding from.
// Deprecated: Use CompressedBatchEntryMetaData.ABI instead.
var CompressedBatchEntryABI = CompressedBatchEntryMetaData.ABI

// CompressedBatchEntryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CompressedBatchEntryMetaData.Bin instead.
var CompressedBatchEntryBin = CompressedBatchEntryMetaData.Bin

// DeployCompressedBatchEntry deploys a new Ethereum contract, binding an instance of CompressedBatchEntry to it.
func DeployCompressedBatchEntry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompressedBatchEntry, error) {
	parsed, err := CompressedBatchEntryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CompressedBatchEntryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CompressedBatchEntry{CompressedBatchEntryCaller: CompressedBatchEntryCaller{contract: contract}, CompressedBatchEntryTransactor: CompressedBatchEntryTransactor{contract: contract}, CompressedBatchEntryFilterer: CompressedBatchEntryFilterer{contract: contract}}, nil
}

// CompressedBatchEntry is an auto generated Go binding around an Ethereum contract.
type CompressedBatchEntry struct {
	CompressedBatchEntryCaller     // Read-only binding to the contract
	CompressedBatchEntryTransactor // Write-only binding to the contract
	CompressedBatchEntryFilterer   // Log filterer for contract events
}

// CompressedBatchEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompressedBatchEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedBatchEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompressedBatchEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedBatchEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompressedBatchEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedBatchEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompressedBatchEntrySession struct {
	Contract     *CompressedBatchEntry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CompressedBatchEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompressedBatchEntryCallerSession struct {
	Contract *CompressedBatchEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CompressedBatchEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompressedBatchEntryTransactorSession struct {
	Contract     *CompressedBatchEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CompressedBatchEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompressedBatchEntryRaw struct {
	Contract *CompressedBatchEntry // Generic contract binding to access the raw methods on
}

// CompressedBatchEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompressedBatchEntryCallerRaw struct {
	Contract *CompressedBatchEntryCaller // Generic read-only contract binding to access the raw methods on
}

// CompressedBatchEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompressedBatchEntryTransactorRaw struct {
	Contract *CompressedBatchEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompressedBatchEntry creates a new instance of CompressedBatchEntry, bound to a specific deployed contract.
func NewCompressedBatchEntry(address common.Address, backend bind.ContractBackend) (*CompressedBatchEntry, error) {
	contract, err := bindCompressedBatchEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompressedBatchEntry{CompressedBatchEntryCaller: CompressedBatchEntryCaller{contract: contract}, CompressedBatchEntryTransactor: CompressedBatchEntryTransactor{contract: contract}, CompressedBatchEntryFilterer: CompressedBatchEntryFilterer{contract: contract}}, nil
}

// NewCompressedBatchEntryCaller creates a new read-only instance of CompressedBatchEntry, bound to a specific deployed contract.
func NewCompressedBatchEntryCaller(address common.Address, caller bind.ContractCaller) (*CompressedBatchEntryCaller, error) {
	contract, err := bindCompressedBatchEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompressedBatchEntryCaller{contract: contract}, nil
}

// NewCompressedBatchEntryTransactor creates a new write-only instance of CompressedBatchEntry, bound to a specific deployed contract.
func NewCompressedBatchEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*CompressedBatchEntryTransactor, error) {
	contract, err := bindCompressedBatchEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompressedBatchEntryTransactor{contract: contract}, nil
}

// NewCompressedBatchEntryFilterer creates a new log filterer instance of CompressedBatchEntry, bound to a specific deployed contract.
func NewCompressedBatchEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*CompressedBatchEntryFilterer, error) {
	contract, err := bindCompressedBatchEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompressedBatchEntryFilterer{contract: contract}, nil
}

// bindCompressedBatchEntry binds a generic wrapper to an already deployed contract.
func bindCompressedBatchEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CompressedBatchEntryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompressedBatchEntry *CompressedBatchEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompressedBatchEntry.Contract.CompressedBatchEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompressedBatchEntry *CompressedBatchEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompressedBatchEntry.Contract.CompressedBatchEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompressedBatchEntry *CompressedBatchEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompressedBatchEntry.Contract.CompressedBatchEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompressedBatchEntry *CompressedBatchEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompressedBatchEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompressedBatchEntry *CompressedBatchEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompressedBatchEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompressedBatchEntry *CompressedBatchEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompressedBatchEntry.Contract.contract.Transact(opts, method, params...)
}

// CompressedBatchProofMetaData contains all meta data concerning the CompressedBatchProof contract.
var CompressedBatchProofMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea264697066735822122040b62834007e2c411ac558ed0cda541ec6d3c8947f4d38c8e143fec8b189f30964736f6c63430008120033",
}

// CompressedBatchProofABI is the input ABI used to generate the binding from.
// Deprecated: Use CompressedBatchProofMetaData.ABI instead.
var CompressedBatchProofABI = CompressedBatchProofMetaData.ABI

// CompressedBatchProofBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CompressedBatchProofMetaData.Bin instead.
var CompressedBatchProofBin = CompressedBatchProofMetaData.Bin

// DeployCompressedBatchProof deploys a new Ethereum contract, binding an instance of CompressedBatchProof to it.
func DeployCompressedBatchProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompressedBatchProof, error) {
	parsed, err := CompressedBatchProofMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CompressedBatchProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CompressedBatchProof{CompressedBatchProofCaller: CompressedBatchProofCaller{contract: contract}, CompressedBatchProofTransactor: CompressedBatchProofTransactor{contract: contract}, CompressedBatchProofFilterer: CompressedBatchProofFilterer{contract: contract}}, nil
}

// CompressedBatchProof is an auto generated Go binding around an Ethereum contract.
type CompressedBatchProof struct {
	CompressedBatchProofCaller     // Read-only binding to the contract
	CompressedBatchProofTransactor // Write-only binding to the contract
	CompressedBatchProofFilterer   // Log filterer for contract events
}

// CompressedBatchProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompressedBatchProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedBatchProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompressedBatchProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedBatchProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompressedBatchProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedBatchProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompressedBatchProofSession struct {
	Contract     *CompressedBatchProof // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CompressedBatchProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompressedBatchProofCallerSession struct {
	Contract *CompressedBatchProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CompressedBatchProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompressedBatchProofTransactorSession struct {
	Contract     *CompressedBatchProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CompressedBatchProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompressedBatchProofRaw struct {
	Contract *CompressedBatchProof // Generic contract binding to access the raw methods on
}

// CompressedBatchProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompressedBatchProofCallerRaw struct {
	Contract *CompressedBatchProofCaller // Generic read-only contract binding to access the raw methods on
}

// CompressedBatchProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompressedBatchProofTransactorRaw struct {
	Contract *CompressedBatchProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompressedBatchProof creates a new instance of CompressedBatchProof, bound to a specific deployed contract.
func NewCompressedBatchProof(address common.Address, backend bind.ContractBackend) (*CompressedBatchProof, error) {
	contract, err := bindCompressedBatchProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompressedBatchProof{CompressedBatchProofCaller: CompressedBatchProofCaller{contract: contract}, CompressedBatchProofTransactor: CompressedBatchProofTransactor{contract: contract}, CompressedBatchProofFilterer: CompressedBatchProofFilterer{contract: contract}}, nil
}

// NewCompressedBatchProofCaller creates a new read-only instance of CompressedBatchProof, bound to a specific deployed contract.
func NewCompressedBatchProofCaller(address common.Address, caller bind.ContractCaller) (*CompressedBatchProofCaller, error) {
	contract, err := bindCompressedBatchProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompressedBatchProofCaller{contract: contract}, nil
}

// NewCompressedBatchProofTransactor creates a new write-only instance of CompressedBatchProof, bound to a specific deployed contract.
func NewCompressedBatchProofTransactor(address common.Address, transactor bind.ContractTransactor) (*CompressedBatchProofTransactor, error) {
	contract, err := bindCompressedBatchProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompressedBatchProofTransactor{contract: contract}, nil
}

// NewCompressedBatchProofFilterer creates a new log filterer instance of CompressedBatchProof, bound to a specific deployed contract.
func NewCompressedBatchProofFilterer(address common.Address, filterer bind.ContractFilterer) (*CompressedBatchProofFilterer, error) {
	contract, err := bindCompressedBatchProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompressedBatchProofFilterer{contract: contract}, nil
}

// bindCompressedBatchProof binds a generic wrapper to an already deployed contract.
func bindCompressedBatchProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CompressedBatchProofMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompressedBatchProof *CompressedBatchProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompressedBatchProof.Contract.CompressedBatchProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompressedBatchProof *CompressedBatchProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompressedBatchProof.Contract.CompressedBatchProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompressedBatchProof *CompressedBatchProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompressedBatchProof.Contract.CompressedBatchProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompressedBatchProof *CompressedBatchProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompressedBatchProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompressedBatchProof *CompressedBatchProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompressedBatchProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompressedBatchProof *CompressedBatchProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompressedBatchProof.Contract.contract.Transact(opts, method, params...)
}

// CompressedExistenceProofMetaData contains all meta data concerning the CompressedExistenceProof contract.
var CompressedExistenceProofMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220a2ec581052c240afbc76d6811519b016a2f2691c1cf704880b10cdc9bdaba4c164736f6c63430008120033",
}

// CompressedExistenceProofABI is the input ABI used to generate the binding from.
// Deprecated: Use CompressedExistenceProofMetaData.ABI instead.
var CompressedExistenceProofABI = CompressedExistenceProofMetaData.ABI

// CompressedExistenceProofBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CompressedExistenceProofMetaData.Bin instead.
var CompressedExistenceProofBin = CompressedExistenceProofMetaData.Bin

// DeployCompressedExistenceProof deploys a new Ethereum contract, binding an instance of CompressedExistenceProof to it.
func DeployCompressedExistenceProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompressedExistenceProof, error) {
	parsed, err := CompressedExistenceProofMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CompressedExistenceProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CompressedExistenceProof{CompressedExistenceProofCaller: CompressedExistenceProofCaller{contract: contract}, CompressedExistenceProofTransactor: CompressedExistenceProofTransactor{contract: contract}, CompressedExistenceProofFilterer: CompressedExistenceProofFilterer{contract: contract}}, nil
}

// CompressedExistenceProof is an auto generated Go binding around an Ethereum contract.
type CompressedExistenceProof struct {
	CompressedExistenceProofCaller     // Read-only binding to the contract
	CompressedExistenceProofTransactor // Write-only binding to the contract
	CompressedExistenceProofFilterer   // Log filterer for contract events
}

// CompressedExistenceProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompressedExistenceProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedExistenceProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompressedExistenceProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedExistenceProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompressedExistenceProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedExistenceProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompressedExistenceProofSession struct {
	Contract     *CompressedExistenceProof // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CompressedExistenceProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompressedExistenceProofCallerSession struct {
	Contract *CompressedExistenceProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// CompressedExistenceProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompressedExistenceProofTransactorSession struct {
	Contract     *CompressedExistenceProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// CompressedExistenceProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompressedExistenceProofRaw struct {
	Contract *CompressedExistenceProof // Generic contract binding to access the raw methods on
}

// CompressedExistenceProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompressedExistenceProofCallerRaw struct {
	Contract *CompressedExistenceProofCaller // Generic read-only contract binding to access the raw methods on
}

// CompressedExistenceProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompressedExistenceProofTransactorRaw struct {
	Contract *CompressedExistenceProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompressedExistenceProof creates a new instance of CompressedExistenceProof, bound to a specific deployed contract.
func NewCompressedExistenceProof(address common.Address, backend bind.ContractBackend) (*CompressedExistenceProof, error) {
	contract, err := bindCompressedExistenceProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompressedExistenceProof{CompressedExistenceProofCaller: CompressedExistenceProofCaller{contract: contract}, CompressedExistenceProofTransactor: CompressedExistenceProofTransactor{contract: contract}, CompressedExistenceProofFilterer: CompressedExistenceProofFilterer{contract: contract}}, nil
}

// NewCompressedExistenceProofCaller creates a new read-only instance of CompressedExistenceProof, bound to a specific deployed contract.
func NewCompressedExistenceProofCaller(address common.Address, caller bind.ContractCaller) (*CompressedExistenceProofCaller, error) {
	contract, err := bindCompressedExistenceProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompressedExistenceProofCaller{contract: contract}, nil
}

// NewCompressedExistenceProofTransactor creates a new write-only instance of CompressedExistenceProof, bound to a specific deployed contract.
func NewCompressedExistenceProofTransactor(address common.Address, transactor bind.ContractTransactor) (*CompressedExistenceProofTransactor, error) {
	contract, err := bindCompressedExistenceProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompressedExistenceProofTransactor{contract: contract}, nil
}

// NewCompressedExistenceProofFilterer creates a new log filterer instance of CompressedExistenceProof, bound to a specific deployed contract.
func NewCompressedExistenceProofFilterer(address common.Address, filterer bind.ContractFilterer) (*CompressedExistenceProofFilterer, error) {
	contract, err := bindCompressedExistenceProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompressedExistenceProofFilterer{contract: contract}, nil
}

// bindCompressedExistenceProof binds a generic wrapper to an already deployed contract.
func bindCompressedExistenceProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CompressedExistenceProofMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompressedExistenceProof *CompressedExistenceProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompressedExistenceProof.Contract.CompressedExistenceProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompressedExistenceProof *CompressedExistenceProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompressedExistenceProof.Contract.CompressedExistenceProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompressedExistenceProof *CompressedExistenceProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompressedExistenceProof.Contract.CompressedExistenceProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompressedExistenceProof *CompressedExistenceProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompressedExistenceProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompressedExistenceProof *CompressedExistenceProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompressedExistenceProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompressedExistenceProof *CompressedExistenceProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompressedExistenceProof.Contract.contract.Transact(opts, method, params...)
}

// CompressedNonExistenceProofMetaData contains all meta data concerning the CompressedNonExistenceProof contract.
var CompressedNonExistenceProofMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea264697066735822122032c2bc4fbcdbad2c7a9446b2d35d836e6ee98130a931a9061bdce1fb3d6b044f64736f6c63430008120033",
}

// CompressedNonExistenceProofABI is the input ABI used to generate the binding from.
// Deprecated: Use CompressedNonExistenceProofMetaData.ABI instead.
var CompressedNonExistenceProofABI = CompressedNonExistenceProofMetaData.ABI

// CompressedNonExistenceProofBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CompressedNonExistenceProofMetaData.Bin instead.
var CompressedNonExistenceProofBin = CompressedNonExistenceProofMetaData.Bin

// DeployCompressedNonExistenceProof deploys a new Ethereum contract, binding an instance of CompressedNonExistenceProof to it.
func DeployCompressedNonExistenceProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompressedNonExistenceProof, error) {
	parsed, err := CompressedNonExistenceProofMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CompressedNonExistenceProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CompressedNonExistenceProof{CompressedNonExistenceProofCaller: CompressedNonExistenceProofCaller{contract: contract}, CompressedNonExistenceProofTransactor: CompressedNonExistenceProofTransactor{contract: contract}, CompressedNonExistenceProofFilterer: CompressedNonExistenceProofFilterer{contract: contract}}, nil
}

// CompressedNonExistenceProof is an auto generated Go binding around an Ethereum contract.
type CompressedNonExistenceProof struct {
	CompressedNonExistenceProofCaller     // Read-only binding to the contract
	CompressedNonExistenceProofTransactor // Write-only binding to the contract
	CompressedNonExistenceProofFilterer   // Log filterer for contract events
}

// CompressedNonExistenceProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompressedNonExistenceProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedNonExistenceProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompressedNonExistenceProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedNonExistenceProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompressedNonExistenceProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedNonExistenceProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompressedNonExistenceProofSession struct {
	Contract     *CompressedNonExistenceProof // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CompressedNonExistenceProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompressedNonExistenceProofCallerSession struct {
	Contract *CompressedNonExistenceProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// CompressedNonExistenceProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompressedNonExistenceProofTransactorSession struct {
	Contract     *CompressedNonExistenceProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// CompressedNonExistenceProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompressedNonExistenceProofRaw struct {
	Contract *CompressedNonExistenceProof // Generic contract binding to access the raw methods on
}

// CompressedNonExistenceProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompressedNonExistenceProofCallerRaw struct {
	Contract *CompressedNonExistenceProofCaller // Generic read-only contract binding to access the raw methods on
}

// CompressedNonExistenceProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompressedNonExistenceProofTransactorRaw struct {
	Contract *CompressedNonExistenceProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompressedNonExistenceProof creates a new instance of CompressedNonExistenceProof, bound to a specific deployed contract.
func NewCompressedNonExistenceProof(address common.Address, backend bind.ContractBackend) (*CompressedNonExistenceProof, error) {
	contract, err := bindCompressedNonExistenceProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompressedNonExistenceProof{CompressedNonExistenceProofCaller: CompressedNonExistenceProofCaller{contract: contract}, CompressedNonExistenceProofTransactor: CompressedNonExistenceProofTransactor{contract: contract}, CompressedNonExistenceProofFilterer: CompressedNonExistenceProofFilterer{contract: contract}}, nil
}

// NewCompressedNonExistenceProofCaller creates a new read-only instance of CompressedNonExistenceProof, bound to a specific deployed contract.
func NewCompressedNonExistenceProofCaller(address common.Address, caller bind.ContractCaller) (*CompressedNonExistenceProofCaller, error) {
	contract, err := bindCompressedNonExistenceProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompressedNonExistenceProofCaller{contract: contract}, nil
}

// NewCompressedNonExistenceProofTransactor creates a new write-only instance of CompressedNonExistenceProof, bound to a specific deployed contract.
func NewCompressedNonExistenceProofTransactor(address common.Address, transactor bind.ContractTransactor) (*CompressedNonExistenceProofTransactor, error) {
	contract, err := bindCompressedNonExistenceProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompressedNonExistenceProofTransactor{contract: contract}, nil
}

// NewCompressedNonExistenceProofFilterer creates a new log filterer instance of CompressedNonExistenceProof, bound to a specific deployed contract.
func NewCompressedNonExistenceProofFilterer(address common.Address, filterer bind.ContractFilterer) (*CompressedNonExistenceProofFilterer, error) {
	contract, err := bindCompressedNonExistenceProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompressedNonExistenceProofFilterer{contract: contract}, nil
}

// bindCompressedNonExistenceProof binds a generic wrapper to an already deployed contract.
func bindCompressedNonExistenceProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CompressedNonExistenceProofMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompressedNonExistenceProof *CompressedNonExistenceProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompressedNonExistenceProof.Contract.CompressedNonExistenceProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompressedNonExistenceProof *CompressedNonExistenceProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompressedNonExistenceProof.Contract.CompressedNonExistenceProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompressedNonExistenceProof *CompressedNonExistenceProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompressedNonExistenceProof.Contract.CompressedNonExistenceProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompressedNonExistenceProof *CompressedNonExistenceProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompressedNonExistenceProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompressedNonExistenceProof *CompressedNonExistenceProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompressedNonExistenceProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompressedNonExistenceProof *CompressedNonExistenceProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompressedNonExistenceProof.Contract.contract.Transact(opts, method, params...)
}

// ConsensusMetaData contains all meta data concerning the Consensus contract.
var ConsensusMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea264697066735822122028f3367e4fe2641a6e23a23b854f591bc8319d02ed7b730c74c021140f4c4d3664736f6c63430008120033",
}

// ConsensusABI is the input ABI used to generate the binding from.
// Deprecated: Use ConsensusMetaData.ABI instead.
var ConsensusABI = ConsensusMetaData.ABI

// ConsensusBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConsensusMetaData.Bin instead.
var ConsensusBin = ConsensusMetaData.Bin

// DeployConsensus deploys a new Ethereum contract, binding an instance of Consensus to it.
func DeployConsensus(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Consensus, error) {
	parsed, err := ConsensusMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConsensusBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Consensus{ConsensusCaller: ConsensusCaller{contract: contract}, ConsensusTransactor: ConsensusTransactor{contract: contract}, ConsensusFilterer: ConsensusFilterer{contract: contract}}, nil
}

// Consensus is an auto generated Go binding around an Ethereum contract.
type Consensus struct {
	ConsensusCaller     // Read-only binding to the contract
	ConsensusTransactor // Write-only binding to the contract
	ConsensusFilterer   // Log filterer for contract events
}

// ConsensusCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConsensusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsensusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConsensusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsensusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConsensusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsensusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConsensusSession struct {
	Contract     *Consensus        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConsensusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConsensusCallerSession struct {
	Contract *ConsensusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ConsensusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConsensusTransactorSession struct {
	Contract     *ConsensusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ConsensusRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConsensusRaw struct {
	Contract *Consensus // Generic contract binding to access the raw methods on
}

// ConsensusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConsensusCallerRaw struct {
	Contract *ConsensusCaller // Generic read-only contract binding to access the raw methods on
}

// ConsensusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConsensusTransactorRaw struct {
	Contract *ConsensusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConsensus creates a new instance of Consensus, bound to a specific deployed contract.
func NewConsensus(address common.Address, backend bind.ContractBackend) (*Consensus, error) {
	contract, err := bindConsensus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Consensus{ConsensusCaller: ConsensusCaller{contract: contract}, ConsensusTransactor: ConsensusTransactor{contract: contract}, ConsensusFilterer: ConsensusFilterer{contract: contract}}, nil
}

// NewConsensusCaller creates a new read-only instance of Consensus, bound to a specific deployed contract.
func NewConsensusCaller(address common.Address, caller bind.ContractCaller) (*ConsensusCaller, error) {
	contract, err := bindConsensus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConsensusCaller{contract: contract}, nil
}

// NewConsensusTransactor creates a new write-only instance of Consensus, bound to a specific deployed contract.
func NewConsensusTransactor(address common.Address, transactor bind.ContractTransactor) (*ConsensusTransactor, error) {
	contract, err := bindConsensus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConsensusTransactor{contract: contract}, nil
}

// NewConsensusFilterer creates a new log filterer instance of Consensus, bound to a specific deployed contract.
func NewConsensusFilterer(address common.Address, filterer bind.ContractFilterer) (*ConsensusFilterer, error) {
	contract, err := bindConsensus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConsensusFilterer{contract: contract}, nil
}

// bindConsensus binds a generic wrapper to an already deployed contract.
func bindConsensus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConsensusMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Consensus *ConsensusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Consensus.Contract.ConsensusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Consensus *ConsensusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Consensus.Contract.ConsensusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Consensus *ConsensusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Consensus.Contract.ConsensusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Consensus *ConsensusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Consensus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Consensus *ConsensusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Consensus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Consensus *ConsensusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Consensus.Contract.contract.Transact(opts, method, params...)
}

// ConsensusStateMetaData contains all meta data concerning the ConsensusState contract.
var ConsensusStateMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220b5d29bfd6a7eecafbccdc376ac7b3eddce165a92db35034fbba0f33d0544235064736f6c63430008120033",
}

// ConsensusStateABI is the input ABI used to generate the binding from.
// Deprecated: Use ConsensusStateMetaData.ABI instead.
var ConsensusStateABI = ConsensusStateMetaData.ABI

// ConsensusStateBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConsensusStateMetaData.Bin instead.
var ConsensusStateBin = ConsensusStateMetaData.Bin

// DeployConsensusState deploys a new Ethereum contract, binding an instance of ConsensusState to it.
func DeployConsensusState(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ConsensusState, error) {
	parsed, err := ConsensusStateMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConsensusStateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ConsensusState{ConsensusStateCaller: ConsensusStateCaller{contract: contract}, ConsensusStateTransactor: ConsensusStateTransactor{contract: contract}, ConsensusStateFilterer: ConsensusStateFilterer{contract: contract}}, nil
}

// ConsensusState is an auto generated Go binding around an Ethereum contract.
type ConsensusState struct {
	ConsensusStateCaller     // Read-only binding to the contract
	ConsensusStateTransactor // Write-only binding to the contract
	ConsensusStateFilterer   // Log filterer for contract events
}

// ConsensusStateCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConsensusStateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsensusStateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConsensusStateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsensusStateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConsensusStateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsensusStateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConsensusStateSession struct {
	Contract     *ConsensusState   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConsensusStateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConsensusStateCallerSession struct {
	Contract *ConsensusStateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ConsensusStateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConsensusStateTransactorSession struct {
	Contract     *ConsensusStateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ConsensusStateRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConsensusStateRaw struct {
	Contract *ConsensusState // Generic contract binding to access the raw methods on
}

// ConsensusStateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConsensusStateCallerRaw struct {
	Contract *ConsensusStateCaller // Generic read-only contract binding to access the raw methods on
}

// ConsensusStateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConsensusStateTransactorRaw struct {
	Contract *ConsensusStateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConsensusState creates a new instance of ConsensusState, bound to a specific deployed contract.
func NewConsensusState(address common.Address, backend bind.ContractBackend) (*ConsensusState, error) {
	contract, err := bindConsensusState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConsensusState{ConsensusStateCaller: ConsensusStateCaller{contract: contract}, ConsensusStateTransactor: ConsensusStateTransactor{contract: contract}, ConsensusStateFilterer: ConsensusStateFilterer{contract: contract}}, nil
}

// NewConsensusStateCaller creates a new read-only instance of ConsensusState, bound to a specific deployed contract.
func NewConsensusStateCaller(address common.Address, caller bind.ContractCaller) (*ConsensusStateCaller, error) {
	contract, err := bindConsensusState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConsensusStateCaller{contract: contract}, nil
}

// NewConsensusStateTransactor creates a new write-only instance of ConsensusState, bound to a specific deployed contract.
func NewConsensusStateTransactor(address common.Address, transactor bind.ContractTransactor) (*ConsensusStateTransactor, error) {
	contract, err := bindConsensusState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConsensusStateTransactor{contract: contract}, nil
}

// NewConsensusStateFilterer creates a new log filterer instance of ConsensusState, bound to a specific deployed contract.
func NewConsensusStateFilterer(address common.Address, filterer bind.ContractFilterer) (*ConsensusStateFilterer, error) {
	contract, err := bindConsensusState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConsensusStateFilterer{contract: contract}, nil
}

// bindConsensusState binds a generic wrapper to an already deployed contract.
func bindConsensusState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConsensusStateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConsensusState *ConsensusStateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConsensusState.Contract.ConsensusStateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConsensusState *ConsensusStateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsensusState.Contract.ConsensusStateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConsensusState *ConsensusStateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConsensusState.Contract.ConsensusStateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConsensusState *ConsensusStateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConsensusState.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConsensusState *ConsensusStateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsensusState.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConsensusState *ConsensusStateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConsensusState.Contract.contract.Transact(opts, method, params...)
}

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextMetaData.ABI instead.
var ContextABI = ContextMetaData.ABI

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContextMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// CrossChainMetaData contains all meta data concerning the CrossChain contract.
var CrossChainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"oracleSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"CrossChainPackage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"packageType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"}],\"name\":\"ReceivedPackage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lowLevelData\",\"type\":\"bytes\"}],\"name\":\"UnexpectedFailureAssertionInPackageHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"UnexpectedRevertInPackageHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"UnsupportedPackage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACK_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN_KEY_PREFIX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FAIL_ACK_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_BATCH_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STORE_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYN_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchSizeForOracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelHandlerContractMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelReceiveSequenceMap\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelSendSequenceMap\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"packageType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"relayFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"encodePayload\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"}],\"name\":\"handlePackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_system\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"internalType\":\"structCrossChain.ChannelInit[]\",\"name\":\"receiveChannelInit\",\"type\":\"tuple[]\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleSequence\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"previousTxHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"registeredContractChannelMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"relayFee\",\"type\":\"uint256\"}],\"name\":\"sendSynPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"txCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608034620006ab5760a081016001600160401b038111828210176200049557604090815260018083526000602084019081528383018281526060850192835283519194928285016001600160401b03811184821017620004955785526006835264014661fc9160d11b602084015260808401928352845195606087016001600160401b038111888210176200049557865260028752602087018636823787511562000695576000905286516001101562000695576001878701528551600081526020810197906001600160401b038911818a1017620004955760e08101906001600160401b0382118a831017620004955760019260c0928a528a526020898201528260608201526080808201528060a081015201528551946080860186811060018060401b038211176200049557875280865260208601978852600087870152600060608701525160078110156200038657600a549151936007851015620003865751926007841015620003865751916009831015620003865760ff61ff0063ff00000062ff00009560181b169660081b1692169063ffffffff191617179160101b161717600a555192835160018060401b0381116200049557600b54600181811c911680156200068a575b60208210146200047457601f811162000637575b506020601f8211600114620005ad578192939495600092620005a1575b50508160011b916000199060031b1c191617600b555b5180518051919291906001600160401b03821162000495576801000000000000000082116200049557602090600c5483600c558084106200054b575b500190600c60005260206000209060005b8160031c81106200050557506007198116808203620004ab575b50505050602082015191600d54908481015160201b6060820151861b6bffffffff0000000000000000169063ffffffff95869467ffffffff00000000978892169060018060601b0319161791161717600d55608081015180519060018060401b0382116200049557600e54600181811c911680156200048a575b60208210146200047457601f811162000426575b50602090601f8311600114620003a85760a0939291600091836200039c575b50508160011b916000199060031b1c191617600e555b01516007811015620003865760ff8019600f5416911617600f558382015192606060105493015160201b1692169060018060401b031916171760105551611ba79081620006ca8239f35b634e487b7160e01b600052602160045260246000fd5b01519050388062000326565b600e600090815260008051602062002271833981519152929190601f198516905b8181106200040d575091600193918560a097969410620003f3575b505050811b01600e556200033c565b015160001960f88460031b161c19169055388080620003e4565b92936020600181928786015181550195019301620003c9565b600e600052620004629060008051602062002271833981519152601f850160051c8101916020861062000469575b601f0160051c0190620006b0565b3862000307565b909150819062000454565b634e487b7160e01b600052602260045260246000fd5b90607f1690620002f3565b634e487b7160e01b600052604160045260246000fd5b9260009360005b8184038110620004ce5750505060031c01553880808062000279565b9091946020620004fa600192885160030b908560021b60031b9163ffffffff809116831b921b19161790565b9601929101620004b2565b6000805b60088110620005205750838201556001016200025f565b855190959160019160209163ffffffff60058a901b81811b199092169216901b179201950162000509565b6200057f90600c60005283600020600780870160031c820192601c8860021b168062000586575b500160031c0190620006b0565b386200024e565b60001990818601918254918a0360031b1c1690553862000572565b015190503880620001fc565b600b6000908152601f198316967f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db992915b8881106200061e5750836001959697981062000604575b505050811b01600b5562000212565b015160001960f88460031b161c19169055388080620005f5565b91926020600181928685015181550194019201620005de565b600b60005262000683907f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db9601f840160051c810191602085106200046957601f0160051c0190620006b0565b38620001df565b90607f1690620001cb565b634e487b7160e01b600052603260045260246000fd5b600080fd5b818110620006bc575050565b60008155600101620006b056fe6080604052600436101561001257600080fd5b60003560e01c806305e682581461013757806314b3023b1461013257806322556cdc1461012d5780632ff32aea14610128578063308325f4146101235780633bdc47a61461011e5780635d6ba4fa146101195780636e47a51a1461011457806374f079b81461010f57806384013b6a1461010a578063863fe4ab146101055780638cc8f56114610100578063b0355f5b146100fb578063c27cdcfb146100f6578063d31f968d146100f1578063d76a8675146100ec578063e3b04805146100e75763f7a251d7146100e257600080fd5b610817565b6107d5565b610791565b61072d565b6106eb565b6106cf565b6106b3565b610694565b610573565b610527565b6104e6565b61041b565b610308565b6101c3565b6101a2565b610186565b610168565b61014c565b600091031261014757565b600080fd5b3461014757600036600319011261014757602060405160008152f35b34610147576000366003190112610147576020600154604051908152f35b3461014757600036600319011261014757602060405160328152f35b3461014757600036600319011261014757602060045460070b604051908152f35b34610147576000366003190112610147576020600254604051908152f35b60ff81160361014757565b634e487b7160e01b600052604160045260246000fd5b6040810190811067ffffffffffffffff82111761021e57604052565b6101ec565b67ffffffffffffffff811161021e57604052565b90601f8019910116810190811067ffffffffffffffff82111761021e57604052565b67ffffffffffffffff811161021e57601f01601f191660200190565b92919261028182610259565b9161028f6040519384610237565b829481845281830111610147578281602093846000960137010152565b60005b8381106102bf5750506000910152565b81810151838201526020016102af565b906020916102e8815180928185528580860191016102ac565b601f01601f1916010190565b9060206103059281815201906102cf565b90565b3461014757606036600319011261014757600435610325816101e1565b60443567ffffffffffffffff8111610147573660238201121561014757610356903690602481600401359101610275565b908151602181018091116103c25761036d81610d26565b9160218301908184116103c2576024358252600184019182116103c257815282116103c2578152604181018082116103c257826103b29160206103be95519201611b05565b604051918291826102f4565b0390f35b610cd7565b6001600160a01b0381160361014757565b6044359067ffffffffffffffff8216820361014757565b6064359067ffffffffffffffff8216820361014757565b359067ffffffffffffffff8216820361014757565b346101475760408060031936011261014757600435610439816103c7565b67ffffffffffffffff90602435828111610147573660238201121561014757806004013592831161021e57835192602090610479828260051b0186610237565b80855260248286019160061b8401019236841161014757602401905b8382106104a8576104a6868661090a565b005b86823603126101475782879182516104bf81610202565b84356104ca816101e1565b81526104d7838601610406565b83820152815201910190610495565b346101475760203660031901126101475760ff600435610505816101e1565b16600052600560205260206001600160a01b0360406000205416604051908152f35b34610147576000366003190112610147576020600354604051908152f35b9181601f840112156101475782359167ffffffffffffffff8311610147576020838186019501011161014757565b346101475760a03660031901126101475767ffffffffffffffff600435818111610147576105a5903690600401610545565b602435838111610147576105bd903690600401610545565b6105c56103d8565b916105ce6103ef565b93608435956105dc876101e1565b6105ea60ff60005416610e28565b60ff871660005260086020528760406000205416809887160361064f5761064a6106166104a699610e74565b61062d8960ff166000526008602052604060002090565b9067ffffffffffffffff1667ffffffffffffffff19825416179055565b610efd565b60405162461bcd60e51b815260206004820152601560248201527f73657175656e6365206e6f7420696e206f7264657200000000000000000000006044820152606490fd5b3461014757600036600319011261014757602060405163010060008152f35b3461014757600036600319011261014757602060405160028152f35b3461014757600036600319011261014757602060405160018152f35b346101475760203660031901126101475760ff60043561070a816101e1565b166000526008602052602067ffffffffffffffff60406000205416604051908152f35b3461014757604036600319011261014757602060ff610785600435610751816103c7565b6001600160a01b0360243591610766836101e1565b166000526006845260406000209060ff16600052602052604060002090565b54166040519015158152f35b34610147576000366003190112610147576103be6040516107b181610202565b600381526269626360e81b60208201526040519182916020835260208301906102cf565b346101475760203660031901126101475760ff6004356107f4816101e1565b166000526007602052602067ffffffffffffffff60406000205416604051908152f35b3461014757606036600319011261014757600435610834816101e1565b60243567ffffffffffffffff811161014757610854903690600401610545565b61086560ff60009493945416610e28565b33600052600660205260ff61088b8360406000209060ff16600052602052604060002090565b54161561089f576104a69260443592611a3e565b60405162461bcd60e51b815260206004820152603160248201527f74686520636f6e747261637420616e64206368616e6e656c2068617665206e6f60448201527f74206265656e20726567697374657265640000000000000000000000000000006064820152608490fd5b9190600060ff815416610c10576009805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039586161790556109616109556009546001600160a01b031690565b6001600160a01b031690565b60405163e89a950160e01b908181526020968782600481875afa8015610b8757610995928691610bf3575b50161515610c76565b6040518181528681600481865afa908115610b87578491610bd6575b506040519087826004816343d549bb60e11b978882525afa918215610b8757610a1b926109f3918791610bb9575b5060ff166000526005602052604060002090565b906001600160a01b031673ffffffffffffffffffffffffffffffffffffffff19825416179055565b610a306109556009546001600160a01b031690565b906040519081528681600481855afa908115610b87578791610a6d918691610b8c575b506001600160a01b03166000526006602052604060002090565b9260046040518094819382525afa8015610b8757610ab092610aa3928592610b58575b509060ff16600052602052604060002090565b805460ff19166001179055565b610aba6032600155565b610ad267ffffffffffffffff80196004541617600455565b610adc6000600255565b610ae66000600355565b8151811015610b405780610b36610b00610b3b9385610cfc565b5161062d610b23610b1b8984015167ffffffffffffffff1690565b925160ff1690565b60ff166000526008602052604060002090565b610ced565b610ae6565b50509050610b56600160ff196000541617600055565b565b610b79919250883d8a11610b80575b610b718183610237565b810190610cc2565b9038610a90565b503d610b67565b610c6a565b610bac9150833d8511610bb2575b610ba48183610237565b810190610c55565b38610a53565b503d610b9a565b610bd091508a3d8c11610b8057610b718183610237565b386109df565b610bed9150873d8911610bb257610ba48183610237565b386109b1565b610c0a9150893d8b11610bb257610ba48183610237565b3861098c565b60405162461bcd60e51b815260206004820152601360248201527f616c726561647920696e697469616c697a6564000000000000000000000000006044820152606490fd5b908160209103126101475751610305816103c7565b6040513d6000823e3d90fd5b15610c7d57565b60405162461bcd60e51b815260206004820152601460248201527f73797374656d20756e696e697469616c697a65640000000000000000000000006044820152606490fd5b908160209103126101475751610305816101e1565b634e487b7160e01b600052601160045260246000fd5b60001981146103c25760010190565b8051821015610d105760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b90610d3082610259565b610d3d6040519182610237565b8281528092610d4e601f1991610259565b0190602036910137565b6020198101919082116103c257565b6000198101919082116103c257565b805190602182018092116103c257610d8d82610d26565b91602183018084116103c25760008152600184019081116103c2576001815283116103c257825260418201908183116103c25761030591602082519201611b05565b805190602182018092116103c257610de682610d26565b91602183018084116103c25760008152600184019081116103c2576002815283116103c257825260418201908183116103c25761030591602082519201611b05565b15610e2f57565b60405162461bcd60e51b815260206004820152600f60248201527f6e6f7420696e697469616c697a656400000000000000000000000000000000006044820152606490fd5b90600167ffffffffffffffff809316019182116103c257565b90816020910312610147575180151581036101475790565b15610eac57565b60405162461bcd60e51b815260206004820152602360248201527f6c6967687420636c69656e74206e6f742073796e632074686520626c6f636b206044820152621e595d60ea1b6064820152608490fd5b95949395610f166109556009546001600160a01b031690565b604051978891632b1d780560e21b835282600460209485935afa988915610b87576001600160a01b03998392610f7c92600092610fd0575b506040516337d7f9c160e21b815267ffffffffffffffff90911660048201529a8b9283919082906024820190565b0392165afa8015610b8757610b5698610f9e92600092610fa3575b5050610ea5565b610fef565b610fc29250803d10610fc9575b610fba8183610237565b810190610e8d565b3880610f97565b503d610fb0565b610fe8919250843d8611610bb257610ba48183610237565b9038610f4e565b949392919060ff851660005260056020526001600160a01b03604060002054161561101d57610b56956111e4565b60405162461bcd60e51b815260206004820152601860248201527f6368616e6e656c206973206e6f7420737570706f7274656400000000000000006044820152606490fd5b1561106957565b60405162461bcd60e51b815260206004820152600b60248201527f6e6f742072656c617965720000000000000000000000000000000000000000006044820152606490fd5b60409060ff610305949316815281602082015201906102cf565b60009060033d116110d557565b905060046000803e60005160e01c90565b600060443d1061030557604051600319913d83016004833e815167ffffffffffffffff918282113d6024840111176111445781840194855193841161114c573d85010160208487010111611144575061030592910160200190610237565b949350505050565b50949350505050565b3d15611180573d9061116682610259565b916111746040519384610237565b82523d6000602084013e565b606090565b6020818303126101475780519067ffffffffffffffff8211610147570181601f820112156101475780516111b881610259565b926111c66040519485610237565b818452602082840101116101475761030591602080850191016102ac565b929394916004906112006109556009546001600160a01b031690565b946020604096875194858092638406c07960e01b82525afa928315610b87576000936117a5575b5061124d9061125593946112466001600160a01b038097163314611062565b3691610275565b953691610275565b5061125f846117cd565b96919290501561175d575083949560ff929394517f48484b8ae53584e6447d0535a274159337a74351c4adf243a6bf94b4c7a16c2e67ffffffffffffffff858816931691806112b8868291909160ff6020820193169052565b0390a3168061150857506112e96112dc8360ff166000526005602052604060002090565b546001600160a01b031690565b16928251631182b87560e01b8152600081806113098587600484016110ae565b038183895af1600091816114e5575b506114aa575060016113286110c8565b6308c379a0146113fa575b61133e575b50505050565b7ffb111707dba108ae503813ed49a361e67f4720ac104ff7f5e9692e32039a1a68926113e46113ee936113ab611372611155565b94826113a561139f6113918360ff166000526007602052604060002090565b5467ffffffffffffffff1690565b92610dcf565b916118b2565b61062d6113d06113cb6113918460ff166000526007602052604060002090565b610e74565b9160ff166000526007602052604060002090565b51918291826102f4565b0390a238808080611338565b6114026110e6565b8061140e575b50611333565b9050847fad48df68de34d2557f1ab71adc0bf7d0a1e23c433408e7b17ff51d588e28b1366114a160009361145f6114556113918960ff166000526007602052604060002090565b886113a589610dcf565b61149661147f6113cb6113918a60ff166000526007602052604060002090565b61062d8960ff166000526007602052604060002090565b8751918291826102f4565b0390a238611408565b935050905081516114ba575b5050565b6113ab610b5692826113a56114df6113918360ff166000526007602052604060002090565b92610d76565b61150191923d8091833e6114f98183610237565b810190611185565b9038611318565b91939160018103611630575061152e6112dc8560ff166000526005602052604060002090565b1692833b1561014757825163831d65d160e01b815291600091839182916115599190600484016110ae565b038183875af19081611617575b506114b65760016115756110c8565b6308c379a0146115c5575b611588575050565b6115c07ffb111707dba108ae503813ed49a361e67f4720ac104ff7f5e9692e32039a1a68916115b5611155565b9051918291826102f4565b0390a2565b6115cd6110e6565b806115d9575b50611580565b9050827fad48df68de34d2557f1ab71adc0bf7d0a1e23c433408e7b17ff51d588e28b13661160e6000938551918291826102f4565b0390a2386115d3565b8061162461162a92610223565b8061013c565b38611566565b60021461163d5750505050565b6116576112dc8560ff166000526005602052604060002090565b1692833b1561014757825163c8509d8160e01b815291600091839182916116829190600484016110ae565b038183875af1908161174a575b5061174357600161169e6110c8565b6308c379a0146116f1575b6116ba575b50505b38808080611338565b6116e77ffb111707dba108ae503813ed49a361e67f4720ac104ff7f5e9692e32039a1a68916115b5611155565b0390a238806116ae565b6116f96110e6565b80611705575b506116a9565b9050827fad48df68de34d2557f1ab71adc0bf7d0a1e23c433408e7b17ff51d588e28b13661173a6000938551918291826102f4565b0390a2386116ff565b50506116b1565b8061162461175792610223565b3861168f565b9450505060ff906117a067ffffffffffffffff7fdee9845a11ea343955bc9858cfa6d6fbcce9c8c1cc4ca22a12997e88b1726b97945193849316961694826102f4565b0390a3565b6112559350906117c561124d9260203d8111610bb257610ba48183610237565b935090611227565b60218151106118365760018101908181116103c257815192602182018093116103c257825192604183018091116103c25782516020198101919082116103c25761182361181c61182e93610d26565b9451610d58565b906020850190611b05565b600193929190565b506040516020810181811067ffffffffffffffff82111761021e57604052600090818152813681378192829190565b60070b677fffffffffffffff81146103c25760010190565b90816020910312610147575161ffff811681036101475790565b60409061ffff610305949316815281602082015201906102cf565b9160049160025443116000146119ff576118f76118d86118d3855460070b90565b611865565b67ffffffffffffffff19600454169067ffffffffffffffff1617600455565b6119016001600355565b61190a43600255565b60206119216109556009546001600160a01b031690565b60405163493279b160e01b815294859182905afa928315610b87576000936119aa575b5060ff7fab63a97d3853c0e724511020661966ef0a1cf04a2f67bf2ccf1015ae390bc7ff916119a561198b61197e61197e60045460070b90565b67ffffffffffffffff1690565b9467ffffffffffffffff6040519485941698169683611897565b0390a4565b7fab63a97d3853c0e724511020661966ef0a1cf04a2f67bf2ccf1015ae390bc7ff9193506119f060ff9160203d81116119f8575b6119e88183610237565b81019061187d565b939150611944565b503d6119de565b611a12611a0d600354610ced565b600355565b600354600154101561190a57611a2f6118d86118d3855460070b90565b611a396001600355565b61190a565b92919260ff811692836000526007602052611a6c67ffffffffffffffff938460406000205416963691610275565b918251602181018091116103c257611a8381610d26565b9260218401908185116103c2578152600184019081116103c2576000815283116103c2578252604182018083116103c25783611ac7916020611acd96519201611b05565b856118b2565b82146103c257610b569160019160005260076020520160406000209067ffffffffffffffff1667ffffffffffffffff19825416179055565b9290925b602093848410611b3e57815181528481018091116103c2579381018091116103c25791601f1981019081116103c25791611b09565b9290919350602003602081116103c257601f81116103c257611b63906101000a610d67565b90518251821691191617905256fea264697066735822122052bee87a56799187406c4e817eb5db5905f7a8a27247b509a6d1f8e7e0e1263364736f6c63430008120033bb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3fd",
}

// CrossChainABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossChainMetaData.ABI instead.
var CrossChainABI = CrossChainMetaData.ABI

// CrossChainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossChainMetaData.Bin instead.
var CrossChainBin = CrossChainMetaData.Bin

// DeployCrossChain deploys a new Ethereum contract, binding an instance of CrossChain to it.
func DeployCrossChain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CrossChain, error) {
	parsed, err := CrossChainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossChainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossChain{CrossChainCaller: CrossChainCaller{contract: contract}, CrossChainTransactor: CrossChainTransactor{contract: contract}, CrossChainFilterer: CrossChainFilterer{contract: contract}}, nil
}

// CrossChain is an auto generated Go binding around an Ethereum contract.
type CrossChain struct {
	CrossChainCaller     // Read-only binding to the contract
	CrossChainTransactor // Write-only binding to the contract
	CrossChainFilterer   // Log filterer for contract events
}

// CrossChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossChainSession struct {
	Contract     *CrossChain       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossChainCallerSession struct {
	Contract *CrossChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CrossChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossChainTransactorSession struct {
	Contract     *CrossChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CrossChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossChainRaw struct {
	Contract *CrossChain // Generic contract binding to access the raw methods on
}

// CrossChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossChainCallerRaw struct {
	Contract *CrossChainCaller // Generic read-only contract binding to access the raw methods on
}

// CrossChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossChainTransactorRaw struct {
	Contract *CrossChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossChain creates a new instance of CrossChain, bound to a specific deployed contract.
func NewCrossChain(address common.Address, backend bind.ContractBackend) (*CrossChain, error) {
	contract, err := bindCrossChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossChain{CrossChainCaller: CrossChainCaller{contract: contract}, CrossChainTransactor: CrossChainTransactor{contract: contract}, CrossChainFilterer: CrossChainFilterer{contract: contract}}, nil
}

// NewCrossChainCaller creates a new read-only instance of CrossChain, bound to a specific deployed contract.
func NewCrossChainCaller(address common.Address, caller bind.ContractCaller) (*CrossChainCaller, error) {
	contract, err := bindCrossChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainCaller{contract: contract}, nil
}

// NewCrossChainTransactor creates a new write-only instance of CrossChain, bound to a specific deployed contract.
func NewCrossChainTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossChainTransactor, error) {
	contract, err := bindCrossChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainTransactor{contract: contract}, nil
}

// NewCrossChainFilterer creates a new log filterer instance of CrossChain, bound to a specific deployed contract.
func NewCrossChainFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossChainFilterer, error) {
	contract, err := bindCrossChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossChainFilterer{contract: contract}, nil
}

// bindCrossChain binds a generic wrapper to an already deployed contract.
func bindCrossChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossChainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossChain *CrossChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChain.Contract.CrossChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossChain *CrossChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChain.Contract.CrossChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossChain *CrossChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChain.Contract.CrossChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossChain *CrossChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossChain *CrossChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossChain *CrossChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChain.Contract.contract.Transact(opts, method, params...)
}

// ACKPACKAGE is a free data retrieval call binding the contract method 0xb0355f5b.
//
// Solidity: function ACK_PACKAGE() view returns(uint8)
func (_CrossChain *CrossChainCaller) ACKPACKAGE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "ACK_PACKAGE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ACKPACKAGE is a free data retrieval call binding the contract method 0xb0355f5b.
//
// Solidity: function ACK_PACKAGE() view returns(uint8)
func (_CrossChain *CrossChainSession) ACKPACKAGE() (uint8, error) {
	return _CrossChain.Contract.ACKPACKAGE(&_CrossChain.CallOpts)
}

// ACKPACKAGE is a free data retrieval call binding the contract method 0xb0355f5b.
//
// Solidity: function ACK_PACKAGE() view returns(uint8)
func (_CrossChain *CrossChainCallerSession) ACKPACKAGE() (uint8, error) {
	return _CrossChain.Contract.ACKPACKAGE(&_CrossChain.CallOpts)
}

// CROSSCHAINKEYPREFIX is a free data retrieval call binding the contract method 0x863fe4ab.
//
// Solidity: function CROSS_CHAIN_KEY_PREFIX() view returns(uint256)
func (_CrossChain *CrossChainCaller) CROSSCHAINKEYPREFIX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "CROSS_CHAIN_KEY_PREFIX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CROSSCHAINKEYPREFIX is a free data retrieval call binding the contract method 0x863fe4ab.
//
// Solidity: function CROSS_CHAIN_KEY_PREFIX() view returns(uint256)
func (_CrossChain *CrossChainSession) CROSSCHAINKEYPREFIX() (*big.Int, error) {
	return _CrossChain.Contract.CROSSCHAINKEYPREFIX(&_CrossChain.CallOpts)
}

// CROSSCHAINKEYPREFIX is a free data retrieval call binding the contract method 0x863fe4ab.
//
// Solidity: function CROSS_CHAIN_KEY_PREFIX() view returns(uint256)
func (_CrossChain *CrossChainCallerSession) CROSSCHAINKEYPREFIX() (*big.Int, error) {
	return _CrossChain.Contract.CROSSCHAINKEYPREFIX(&_CrossChain.CallOpts)
}

// FAILACKPACKAGE is a free data retrieval call binding the contract method 0x8cc8f561.
//
// Solidity: function FAIL_ACK_PACKAGE() view returns(uint8)
func (_CrossChain *CrossChainCaller) FAILACKPACKAGE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "FAIL_ACK_PACKAGE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// FAILACKPACKAGE is a free data retrieval call binding the contract method 0x8cc8f561.
//
// Solidity: function FAIL_ACK_PACKAGE() view returns(uint8)
func (_CrossChain *CrossChainSession) FAILACKPACKAGE() (uint8, error) {
	return _CrossChain.Contract.FAILACKPACKAGE(&_CrossChain.CallOpts)
}

// FAILACKPACKAGE is a free data retrieval call binding the contract method 0x8cc8f561.
//
// Solidity: function FAIL_ACK_PACKAGE() view returns(uint8)
func (_CrossChain *CrossChainCallerSession) FAILACKPACKAGE() (uint8, error) {
	return _CrossChain.Contract.FAILACKPACKAGE(&_CrossChain.CallOpts)
}

// INITBATCHSIZE is a free data retrieval call binding the contract method 0x22556cdc.
//
// Solidity: function INIT_BATCH_SIZE() view returns(uint256)
func (_CrossChain *CrossChainCaller) INITBATCHSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "INIT_BATCH_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITBATCHSIZE is a free data retrieval call binding the contract method 0x22556cdc.
//
// Solidity: function INIT_BATCH_SIZE() view returns(uint256)
func (_CrossChain *CrossChainSession) INITBATCHSIZE() (*big.Int, error) {
	return _CrossChain.Contract.INITBATCHSIZE(&_CrossChain.CallOpts)
}

// INITBATCHSIZE is a free data retrieval call binding the contract method 0x22556cdc.
//
// Solidity: function INIT_BATCH_SIZE() view returns(uint256)
func (_CrossChain *CrossChainCallerSession) INITBATCHSIZE() (*big.Int, error) {
	return _CrossChain.Contract.INITBATCHSIZE(&_CrossChain.CallOpts)
}

// STORENAME is a free data retrieval call binding the contract method 0xd76a8675.
//
// Solidity: function STORE_NAME() view returns(string)
func (_CrossChain *CrossChainCaller) STORENAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "STORE_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// STORENAME is a free data retrieval call binding the contract method 0xd76a8675.
//
// Solidity: function STORE_NAME() view returns(string)
func (_CrossChain *CrossChainSession) STORENAME() (string, error) {
	return _CrossChain.Contract.STORENAME(&_CrossChain.CallOpts)
}

// STORENAME is a free data retrieval call binding the contract method 0xd76a8675.
//
// Solidity: function STORE_NAME() view returns(string)
func (_CrossChain *CrossChainCallerSession) STORENAME() (string, error) {
	return _CrossChain.Contract.STORENAME(&_CrossChain.CallOpts)
}

// SYNPACKAGE is a free data retrieval call binding the contract method 0x05e68258.
//
// Solidity: function SYN_PACKAGE() view returns(uint8)
func (_CrossChain *CrossChainCaller) SYNPACKAGE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "SYN_PACKAGE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SYNPACKAGE is a free data retrieval call binding the contract method 0x05e68258.
//
// Solidity: function SYN_PACKAGE() view returns(uint8)
func (_CrossChain *CrossChainSession) SYNPACKAGE() (uint8, error) {
	return _CrossChain.Contract.SYNPACKAGE(&_CrossChain.CallOpts)
}

// SYNPACKAGE is a free data retrieval call binding the contract method 0x05e68258.
//
// Solidity: function SYN_PACKAGE() view returns(uint8)
func (_CrossChain *CrossChainCallerSession) SYNPACKAGE() (uint8, error) {
	return _CrossChain.Contract.SYNPACKAGE(&_CrossChain.CallOpts)
}

// BatchSizeForOracle is a free data retrieval call binding the contract method 0x14b3023b.
//
// Solidity: function batchSizeForOracle() view returns(uint256)
func (_CrossChain *CrossChainCaller) BatchSizeForOracle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "batchSizeForOracle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchSizeForOracle is a free data retrieval call binding the contract method 0x14b3023b.
//
// Solidity: function batchSizeForOracle() view returns(uint256)
func (_CrossChain *CrossChainSession) BatchSizeForOracle() (*big.Int, error) {
	return _CrossChain.Contract.BatchSizeForOracle(&_CrossChain.CallOpts)
}

// BatchSizeForOracle is a free data retrieval call binding the contract method 0x14b3023b.
//
// Solidity: function batchSizeForOracle() view returns(uint256)
func (_CrossChain *CrossChainCallerSession) BatchSizeForOracle() (*big.Int, error) {
	return _CrossChain.Contract.BatchSizeForOracle(&_CrossChain.CallOpts)
}

// ChannelHandlerContractMap is a free data retrieval call binding the contract method 0x6e47a51a.
//
// Solidity: function channelHandlerContractMap(uint8 ) view returns(address)
func (_CrossChain *CrossChainCaller) ChannelHandlerContractMap(opts *bind.CallOpts, arg0 uint8) (common.Address, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "channelHandlerContractMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChannelHandlerContractMap is a free data retrieval call binding the contract method 0x6e47a51a.
//
// Solidity: function channelHandlerContractMap(uint8 ) view returns(address)
func (_CrossChain *CrossChainSession) ChannelHandlerContractMap(arg0 uint8) (common.Address, error) {
	return _CrossChain.Contract.ChannelHandlerContractMap(&_CrossChain.CallOpts, arg0)
}

// ChannelHandlerContractMap is a free data retrieval call binding the contract method 0x6e47a51a.
//
// Solidity: function channelHandlerContractMap(uint8 ) view returns(address)
func (_CrossChain *CrossChainCallerSession) ChannelHandlerContractMap(arg0 uint8) (common.Address, error) {
	return _CrossChain.Contract.ChannelHandlerContractMap(&_CrossChain.CallOpts, arg0)
}

// ChannelReceiveSequenceMap is a free data retrieval call binding the contract method 0xc27cdcfb.
//
// Solidity: function channelReceiveSequenceMap(uint8 ) view returns(uint64)
func (_CrossChain *CrossChainCaller) ChannelReceiveSequenceMap(opts *bind.CallOpts, arg0 uint8) (uint64, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "channelReceiveSequenceMap", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChannelReceiveSequenceMap is a free data retrieval call binding the contract method 0xc27cdcfb.
//
// Solidity: function channelReceiveSequenceMap(uint8 ) view returns(uint64)
func (_CrossChain *CrossChainSession) ChannelReceiveSequenceMap(arg0 uint8) (uint64, error) {
	return _CrossChain.Contract.ChannelReceiveSequenceMap(&_CrossChain.CallOpts, arg0)
}

// ChannelReceiveSequenceMap is a free data retrieval call binding the contract method 0xc27cdcfb.
//
// Solidity: function channelReceiveSequenceMap(uint8 ) view returns(uint64)
func (_CrossChain *CrossChainCallerSession) ChannelReceiveSequenceMap(arg0 uint8) (uint64, error) {
	return _CrossChain.Contract.ChannelReceiveSequenceMap(&_CrossChain.CallOpts, arg0)
}

// ChannelSendSequenceMap is a free data retrieval call binding the contract method 0xe3b04805.
//
// Solidity: function channelSendSequenceMap(uint8 ) view returns(uint64)
func (_CrossChain *CrossChainCaller) ChannelSendSequenceMap(opts *bind.CallOpts, arg0 uint8) (uint64, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "channelSendSequenceMap", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChannelSendSequenceMap is a free data retrieval call binding the contract method 0xe3b04805.
//
// Solidity: function channelSendSequenceMap(uint8 ) view returns(uint64)
func (_CrossChain *CrossChainSession) ChannelSendSequenceMap(arg0 uint8) (uint64, error) {
	return _CrossChain.Contract.ChannelSendSequenceMap(&_CrossChain.CallOpts, arg0)
}

// ChannelSendSequenceMap is a free data retrieval call binding the contract method 0xe3b04805.
//
// Solidity: function channelSendSequenceMap(uint8 ) view returns(uint64)
func (_CrossChain *CrossChainCallerSession) ChannelSendSequenceMap(arg0 uint8) (uint64, error) {
	return _CrossChain.Contract.ChannelSendSequenceMap(&_CrossChain.CallOpts, arg0)
}

// EncodePayload is a free data retrieval call binding the contract method 0x3bdc47a6.
//
// Solidity: function encodePayload(uint8 packageType, uint256 relayFee, bytes msgBytes) pure returns(bytes)
func (_CrossChain *CrossChainCaller) EncodePayload(opts *bind.CallOpts, packageType uint8, relayFee *big.Int, msgBytes []byte) ([]byte, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "encodePayload", packageType, relayFee, msgBytes)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodePayload is a free data retrieval call binding the contract method 0x3bdc47a6.
//
// Solidity: function encodePayload(uint8 packageType, uint256 relayFee, bytes msgBytes) pure returns(bytes)
func (_CrossChain *CrossChainSession) EncodePayload(packageType uint8, relayFee *big.Int, msgBytes []byte) ([]byte, error) {
	return _CrossChain.Contract.EncodePayload(&_CrossChain.CallOpts, packageType, relayFee, msgBytes)
}

// EncodePayload is a free data retrieval call binding the contract method 0x3bdc47a6.
//
// Solidity: function encodePayload(uint8 packageType, uint256 relayFee, bytes msgBytes) pure returns(bytes)
func (_CrossChain *CrossChainCallerSession) EncodePayload(packageType uint8, relayFee *big.Int, msgBytes []byte) ([]byte, error) {
	return _CrossChain.Contract.EncodePayload(&_CrossChain.CallOpts, packageType, relayFee, msgBytes)
}

// OracleSequence is a free data retrieval call binding the contract method 0x2ff32aea.
//
// Solidity: function oracleSequence() view returns(int64)
func (_CrossChain *CrossChainCaller) OracleSequence(opts *bind.CallOpts) (int64, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "oracleSequence")

	if err != nil {
		return *new(int64), err
	}

	out0 := *abi.ConvertType(out[0], new(int64)).(*int64)

	return out0, err

}

// OracleSequence is a free data retrieval call binding the contract method 0x2ff32aea.
//
// Solidity: function oracleSequence() view returns(int64)
func (_CrossChain *CrossChainSession) OracleSequence() (int64, error) {
	return _CrossChain.Contract.OracleSequence(&_CrossChain.CallOpts)
}

// OracleSequence is a free data retrieval call binding the contract method 0x2ff32aea.
//
// Solidity: function oracleSequence() view returns(int64)
func (_CrossChain *CrossChainCallerSession) OracleSequence() (int64, error) {
	return _CrossChain.Contract.OracleSequence(&_CrossChain.CallOpts)
}

// PreviousTxHeight is a free data retrieval call binding the contract method 0x308325f4.
//
// Solidity: function previousTxHeight() view returns(uint256)
func (_CrossChain *CrossChainCaller) PreviousTxHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "previousTxHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviousTxHeight is a free data retrieval call binding the contract method 0x308325f4.
//
// Solidity: function previousTxHeight() view returns(uint256)
func (_CrossChain *CrossChainSession) PreviousTxHeight() (*big.Int, error) {
	return _CrossChain.Contract.PreviousTxHeight(&_CrossChain.CallOpts)
}

// PreviousTxHeight is a free data retrieval call binding the contract method 0x308325f4.
//
// Solidity: function previousTxHeight() view returns(uint256)
func (_CrossChain *CrossChainCallerSession) PreviousTxHeight() (*big.Int, error) {
	return _CrossChain.Contract.PreviousTxHeight(&_CrossChain.CallOpts)
}

// RegisteredContractChannelMap is a free data retrieval call binding the contract method 0xd31f968d.
//
// Solidity: function registeredContractChannelMap(address , uint8 ) view returns(bool)
func (_CrossChain *CrossChainCaller) RegisteredContractChannelMap(opts *bind.CallOpts, arg0 common.Address, arg1 uint8) (bool, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "registeredContractChannelMap", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredContractChannelMap is a free data retrieval call binding the contract method 0xd31f968d.
//
// Solidity: function registeredContractChannelMap(address , uint8 ) view returns(bool)
func (_CrossChain *CrossChainSession) RegisteredContractChannelMap(arg0 common.Address, arg1 uint8) (bool, error) {
	return _CrossChain.Contract.RegisteredContractChannelMap(&_CrossChain.CallOpts, arg0, arg1)
}

// RegisteredContractChannelMap is a free data retrieval call binding the contract method 0xd31f968d.
//
// Solidity: function registeredContractChannelMap(address , uint8 ) view returns(bool)
func (_CrossChain *CrossChainCallerSession) RegisteredContractChannelMap(arg0 common.Address, arg1 uint8) (bool, error) {
	return _CrossChain.Contract.RegisteredContractChannelMap(&_CrossChain.CallOpts, arg0, arg1)
}

// TxCounter is a free data retrieval call binding the contract method 0x74f079b8.
//
// Solidity: function txCounter() view returns(uint256)
func (_CrossChain *CrossChainCaller) TxCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossChain.contract.Call(opts, &out, "txCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TxCounter is a free data retrieval call binding the contract method 0x74f079b8.
//
// Solidity: function txCounter() view returns(uint256)
func (_CrossChain *CrossChainSession) TxCounter() (*big.Int, error) {
	return _CrossChain.Contract.TxCounter(&_CrossChain.CallOpts)
}

// TxCounter is a free data retrieval call binding the contract method 0x74f079b8.
//
// Solidity: function txCounter() view returns(uint256)
func (_CrossChain *CrossChainCallerSession) TxCounter() (*big.Int, error) {
	return _CrossChain.Contract.TxCounter(&_CrossChain.CallOpts)
}

// HandlePackage is a paid mutator transaction binding the contract method 0x84013b6a.
//
// Solidity: function handlePackage(bytes payload, bytes proof, uint64 height, uint64 packageSequence, uint8 channelId) returns()
func (_CrossChain *CrossChainTransactor) HandlePackage(opts *bind.TransactOpts, payload []byte, proof []byte, height uint64, packageSequence uint64, channelId uint8) (*types.Transaction, error) {
	return _CrossChain.contract.Transact(opts, "handlePackage", payload, proof, height, packageSequence, channelId)
}

// HandlePackage is a paid mutator transaction binding the contract method 0x84013b6a.
//
// Solidity: function handlePackage(bytes payload, bytes proof, uint64 height, uint64 packageSequence, uint8 channelId) returns()
func (_CrossChain *CrossChainSession) HandlePackage(payload []byte, proof []byte, height uint64, packageSequence uint64, channelId uint8) (*types.Transaction, error) {
	return _CrossChain.Contract.HandlePackage(&_CrossChain.TransactOpts, payload, proof, height, packageSequence, channelId)
}

// HandlePackage is a paid mutator transaction binding the contract method 0x84013b6a.
//
// Solidity: function handlePackage(bytes payload, bytes proof, uint64 height, uint64 packageSequence, uint8 channelId) returns()
func (_CrossChain *CrossChainTransactorSession) HandlePackage(payload []byte, proof []byte, height uint64, packageSequence uint64, channelId uint8) (*types.Transaction, error) {
	return _CrossChain.Contract.HandlePackage(&_CrossChain.TransactOpts, payload, proof, height, packageSequence, channelId)
}

// Init is a paid mutator transaction binding the contract method 0x5d6ba4fa.
//
// Solidity: function init(address _system, (uint8,uint64)[] receiveChannelInit) returns()
func (_CrossChain *CrossChainTransactor) Init(opts *bind.TransactOpts, _system common.Address, receiveChannelInit []CrossChainChannelInit) (*types.Transaction, error) {
	return _CrossChain.contract.Transact(opts, "init", _system, receiveChannelInit)
}

// Init is a paid mutator transaction binding the contract method 0x5d6ba4fa.
//
// Solidity: function init(address _system, (uint8,uint64)[] receiveChannelInit) returns()
func (_CrossChain *CrossChainSession) Init(_system common.Address, receiveChannelInit []CrossChainChannelInit) (*types.Transaction, error) {
	return _CrossChain.Contract.Init(&_CrossChain.TransactOpts, _system, receiveChannelInit)
}

// Init is a paid mutator transaction binding the contract method 0x5d6ba4fa.
//
// Solidity: function init(address _system, (uint8,uint64)[] receiveChannelInit) returns()
func (_CrossChain *CrossChainTransactorSession) Init(_system common.Address, receiveChannelInit []CrossChainChannelInit) (*types.Transaction, error) {
	return _CrossChain.Contract.Init(&_CrossChain.TransactOpts, _system, receiveChannelInit)
}

// SendSynPackage is a paid mutator transaction binding the contract method 0xf7a251d7.
//
// Solidity: function sendSynPackage(uint8 channelId, bytes msgBytes, uint256 relayFee) returns()
func (_CrossChain *CrossChainTransactor) SendSynPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte, relayFee *big.Int) (*types.Transaction, error) {
	return _CrossChain.contract.Transact(opts, "sendSynPackage", channelId, msgBytes, relayFee)
}

// SendSynPackage is a paid mutator transaction binding the contract method 0xf7a251d7.
//
// Solidity: function sendSynPackage(uint8 channelId, bytes msgBytes, uint256 relayFee) returns()
func (_CrossChain *CrossChainSession) SendSynPackage(channelId uint8, msgBytes []byte, relayFee *big.Int) (*types.Transaction, error) {
	return _CrossChain.Contract.SendSynPackage(&_CrossChain.TransactOpts, channelId, msgBytes, relayFee)
}

// SendSynPackage is a paid mutator transaction binding the contract method 0xf7a251d7.
//
// Solidity: function sendSynPackage(uint8 channelId, bytes msgBytes, uint256 relayFee) returns()
func (_CrossChain *CrossChainTransactorSession) SendSynPackage(channelId uint8, msgBytes []byte, relayFee *big.Int) (*types.Transaction, error) {
	return _CrossChain.Contract.SendSynPackage(&_CrossChain.TransactOpts, channelId, msgBytes, relayFee)
}

// CrossChainCrossChainPackageIterator is returned from FilterCrossChainPackage and is used to iterate over the raw logs and unpacked data for CrossChainPackage events raised by the CrossChain contract.
type CrossChainCrossChainPackageIterator struct {
	Event *CrossChainCrossChainPackage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossChainCrossChainPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainCrossChainPackage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossChainCrossChainPackage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossChainCrossChainPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainCrossChainPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainCrossChainPackage represents a CrossChainPackage event raised by the CrossChain contract.
type CrossChainCrossChainPackage struct {
	ChainId         uint16
	OracleSequence  uint64
	PackageSequence uint64
	ChannelId       uint8
	Payload         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCrossChainPackage is a free log retrieval operation binding the contract event 0xab63a97d3853c0e724511020661966ef0a1cf04a2f67bf2ccf1015ae390bc7ff.
//
// Solidity: event CrossChainPackage(uint16 chainId, uint64 indexed oracleSequence, uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_CrossChain *CrossChainFilterer) FilterCrossChainPackage(opts *bind.FilterOpts, oracleSequence []uint64, packageSequence []uint64, channelId []uint8) (*CrossChainCrossChainPackageIterator, error) {

	var oracleSequenceRule []interface{}
	for _, oracleSequenceItem := range oracleSequence {
		oracleSequenceRule = append(oracleSequenceRule, oracleSequenceItem)
	}
	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _CrossChain.contract.FilterLogs(opts, "CrossChainPackage", oracleSequenceRule, packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainCrossChainPackageIterator{contract: _CrossChain.contract, event: "CrossChainPackage", logs: logs, sub: sub}, nil
}

// WatchCrossChainPackage is a free log subscription operation binding the contract event 0xab63a97d3853c0e724511020661966ef0a1cf04a2f67bf2ccf1015ae390bc7ff.
//
// Solidity: event CrossChainPackage(uint16 chainId, uint64 indexed oracleSequence, uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_CrossChain *CrossChainFilterer) WatchCrossChainPackage(opts *bind.WatchOpts, sink chan<- *CrossChainCrossChainPackage, oracleSequence []uint64, packageSequence []uint64, channelId []uint8) (event.Subscription, error) {

	var oracleSequenceRule []interface{}
	for _, oracleSequenceItem := range oracleSequence {
		oracleSequenceRule = append(oracleSequenceRule, oracleSequenceItem)
	}
	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _CrossChain.contract.WatchLogs(opts, "CrossChainPackage", oracleSequenceRule, packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainCrossChainPackage)
				if err := _CrossChain.contract.UnpackLog(event, "CrossChainPackage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossChainPackage is a log parse operation binding the contract event 0xab63a97d3853c0e724511020661966ef0a1cf04a2f67bf2ccf1015ae390bc7ff.
//
// Solidity: event CrossChainPackage(uint16 chainId, uint64 indexed oracleSequence, uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_CrossChain *CrossChainFilterer) ParseCrossChainPackage(log types.Log) (*CrossChainCrossChainPackage, error) {
	event := new(CrossChainCrossChainPackage)
	if err := _CrossChain.contract.UnpackLog(event, "CrossChainPackage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainReceivedPackageIterator is returned from FilterReceivedPackage and is used to iterate over the raw logs and unpacked data for ReceivedPackage events raised by the CrossChain contract.
type CrossChainReceivedPackageIterator struct {
	Event *CrossChainReceivedPackage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossChainReceivedPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainReceivedPackage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossChainReceivedPackage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossChainReceivedPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainReceivedPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainReceivedPackage represents a ReceivedPackage event raised by the CrossChain contract.
type CrossChainReceivedPackage struct {
	PackageType     uint8
	PackageSequence uint64
	ChannelId       uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReceivedPackage is a free log retrieval operation binding the contract event 0x48484b8ae53584e6447d0535a274159337a74351c4adf243a6bf94b4c7a16c2e.
//
// Solidity: event ReceivedPackage(uint8 packageType, uint64 indexed packageSequence, uint8 indexed channelId)
func (_CrossChain *CrossChainFilterer) FilterReceivedPackage(opts *bind.FilterOpts, packageSequence []uint64, channelId []uint8) (*CrossChainReceivedPackageIterator, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _CrossChain.contract.FilterLogs(opts, "ReceivedPackage", packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainReceivedPackageIterator{contract: _CrossChain.contract, event: "ReceivedPackage", logs: logs, sub: sub}, nil
}

// WatchReceivedPackage is a free log subscription operation binding the contract event 0x48484b8ae53584e6447d0535a274159337a74351c4adf243a6bf94b4c7a16c2e.
//
// Solidity: event ReceivedPackage(uint8 packageType, uint64 indexed packageSequence, uint8 indexed channelId)
func (_CrossChain *CrossChainFilterer) WatchReceivedPackage(opts *bind.WatchOpts, sink chan<- *CrossChainReceivedPackage, packageSequence []uint64, channelId []uint8) (event.Subscription, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _CrossChain.contract.WatchLogs(opts, "ReceivedPackage", packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainReceivedPackage)
				if err := _CrossChain.contract.UnpackLog(event, "ReceivedPackage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceivedPackage is a log parse operation binding the contract event 0x48484b8ae53584e6447d0535a274159337a74351c4adf243a6bf94b4c7a16c2e.
//
// Solidity: event ReceivedPackage(uint8 packageType, uint64 indexed packageSequence, uint8 indexed channelId)
func (_CrossChain *CrossChainFilterer) ParseReceivedPackage(log types.Log) (*CrossChainReceivedPackage, error) {
	event := new(CrossChainReceivedPackage)
	if err := _CrossChain.contract.UnpackLog(event, "ReceivedPackage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainUnexpectedFailureAssertionInPackageHandlerIterator is returned from FilterUnexpectedFailureAssertionInPackageHandler and is used to iterate over the raw logs and unpacked data for UnexpectedFailureAssertionInPackageHandler events raised by the CrossChain contract.
type CrossChainUnexpectedFailureAssertionInPackageHandlerIterator struct {
	Event *CrossChainUnexpectedFailureAssertionInPackageHandler // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossChainUnexpectedFailureAssertionInPackageHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainUnexpectedFailureAssertionInPackageHandler)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossChainUnexpectedFailureAssertionInPackageHandler)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossChainUnexpectedFailureAssertionInPackageHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainUnexpectedFailureAssertionInPackageHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainUnexpectedFailureAssertionInPackageHandler represents a UnexpectedFailureAssertionInPackageHandler event raised by the CrossChain contract.
type CrossChainUnexpectedFailureAssertionInPackageHandler struct {
	ContractAddr common.Address
	LowLevelData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnexpectedFailureAssertionInPackageHandler is a free log retrieval operation binding the contract event 0xfb111707dba108ae503813ed49a361e67f4720ac104ff7f5e9692e32039a1a68.
//
// Solidity: event UnexpectedFailureAssertionInPackageHandler(address indexed contractAddr, bytes lowLevelData)
func (_CrossChain *CrossChainFilterer) FilterUnexpectedFailureAssertionInPackageHandler(opts *bind.FilterOpts, contractAddr []common.Address) (*CrossChainUnexpectedFailureAssertionInPackageHandlerIterator, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _CrossChain.contract.FilterLogs(opts, "UnexpectedFailureAssertionInPackageHandler", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainUnexpectedFailureAssertionInPackageHandlerIterator{contract: _CrossChain.contract, event: "UnexpectedFailureAssertionInPackageHandler", logs: logs, sub: sub}, nil
}

// WatchUnexpectedFailureAssertionInPackageHandler is a free log subscription operation binding the contract event 0xfb111707dba108ae503813ed49a361e67f4720ac104ff7f5e9692e32039a1a68.
//
// Solidity: event UnexpectedFailureAssertionInPackageHandler(address indexed contractAddr, bytes lowLevelData)
func (_CrossChain *CrossChainFilterer) WatchUnexpectedFailureAssertionInPackageHandler(opts *bind.WatchOpts, sink chan<- *CrossChainUnexpectedFailureAssertionInPackageHandler, contractAddr []common.Address) (event.Subscription, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _CrossChain.contract.WatchLogs(opts, "UnexpectedFailureAssertionInPackageHandler", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainUnexpectedFailureAssertionInPackageHandler)
				if err := _CrossChain.contract.UnpackLog(event, "UnexpectedFailureAssertionInPackageHandler", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnexpectedFailureAssertionInPackageHandler is a log parse operation binding the contract event 0xfb111707dba108ae503813ed49a361e67f4720ac104ff7f5e9692e32039a1a68.
//
// Solidity: event UnexpectedFailureAssertionInPackageHandler(address indexed contractAddr, bytes lowLevelData)
func (_CrossChain *CrossChainFilterer) ParseUnexpectedFailureAssertionInPackageHandler(log types.Log) (*CrossChainUnexpectedFailureAssertionInPackageHandler, error) {
	event := new(CrossChainUnexpectedFailureAssertionInPackageHandler)
	if err := _CrossChain.contract.UnpackLog(event, "UnexpectedFailureAssertionInPackageHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainUnexpectedRevertInPackageHandlerIterator is returned from FilterUnexpectedRevertInPackageHandler and is used to iterate over the raw logs and unpacked data for UnexpectedRevertInPackageHandler events raised by the CrossChain contract.
type CrossChainUnexpectedRevertInPackageHandlerIterator struct {
	Event *CrossChainUnexpectedRevertInPackageHandler // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossChainUnexpectedRevertInPackageHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainUnexpectedRevertInPackageHandler)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossChainUnexpectedRevertInPackageHandler)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossChainUnexpectedRevertInPackageHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainUnexpectedRevertInPackageHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainUnexpectedRevertInPackageHandler represents a UnexpectedRevertInPackageHandler event raised by the CrossChain contract.
type CrossChainUnexpectedRevertInPackageHandler struct {
	ContractAddr common.Address
	Reason       string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnexpectedRevertInPackageHandler is a free log retrieval operation binding the contract event 0xad48df68de34d2557f1ab71adc0bf7d0a1e23c433408e7b17ff51d588e28b136.
//
// Solidity: event UnexpectedRevertInPackageHandler(address indexed contractAddr, string reason)
func (_CrossChain *CrossChainFilterer) FilterUnexpectedRevertInPackageHandler(opts *bind.FilterOpts, contractAddr []common.Address) (*CrossChainUnexpectedRevertInPackageHandlerIterator, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _CrossChain.contract.FilterLogs(opts, "UnexpectedRevertInPackageHandler", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainUnexpectedRevertInPackageHandlerIterator{contract: _CrossChain.contract, event: "UnexpectedRevertInPackageHandler", logs: logs, sub: sub}, nil
}

// WatchUnexpectedRevertInPackageHandler is a free log subscription operation binding the contract event 0xad48df68de34d2557f1ab71adc0bf7d0a1e23c433408e7b17ff51d588e28b136.
//
// Solidity: event UnexpectedRevertInPackageHandler(address indexed contractAddr, string reason)
func (_CrossChain *CrossChainFilterer) WatchUnexpectedRevertInPackageHandler(opts *bind.WatchOpts, sink chan<- *CrossChainUnexpectedRevertInPackageHandler, contractAddr []common.Address) (event.Subscription, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _CrossChain.contract.WatchLogs(opts, "UnexpectedRevertInPackageHandler", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainUnexpectedRevertInPackageHandler)
				if err := _CrossChain.contract.UnpackLog(event, "UnexpectedRevertInPackageHandler", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnexpectedRevertInPackageHandler is a log parse operation binding the contract event 0xad48df68de34d2557f1ab71adc0bf7d0a1e23c433408e7b17ff51d588e28b136.
//
// Solidity: event UnexpectedRevertInPackageHandler(address indexed contractAddr, string reason)
func (_CrossChain *CrossChainFilterer) ParseUnexpectedRevertInPackageHandler(log types.Log) (*CrossChainUnexpectedRevertInPackageHandler, error) {
	event := new(CrossChainUnexpectedRevertInPackageHandler)
	if err := _CrossChain.contract.UnpackLog(event, "UnexpectedRevertInPackageHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainUnsupportedPackageIterator is returned from FilterUnsupportedPackage and is used to iterate over the raw logs and unpacked data for UnsupportedPackage events raised by the CrossChain contract.
type CrossChainUnsupportedPackageIterator struct {
	Event *CrossChainUnsupportedPackage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrossChainUnsupportedPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainUnsupportedPackage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrossChainUnsupportedPackage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrossChainUnsupportedPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainUnsupportedPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainUnsupportedPackage represents a UnsupportedPackage event raised by the CrossChain contract.
type CrossChainUnsupportedPackage struct {
	PackageSequence uint64
	ChannelId       uint8
	Payload         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnsupportedPackage is a free log retrieval operation binding the contract event 0xdee9845a11ea343955bc9858cfa6d6fbcce9c8c1cc4ca22a12997e88b1726b97.
//
// Solidity: event UnsupportedPackage(uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_CrossChain *CrossChainFilterer) FilterUnsupportedPackage(opts *bind.FilterOpts, packageSequence []uint64, channelId []uint8) (*CrossChainUnsupportedPackageIterator, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _CrossChain.contract.FilterLogs(opts, "UnsupportedPackage", packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainUnsupportedPackageIterator{contract: _CrossChain.contract, event: "UnsupportedPackage", logs: logs, sub: sub}, nil
}

// WatchUnsupportedPackage is a free log subscription operation binding the contract event 0xdee9845a11ea343955bc9858cfa6d6fbcce9c8c1cc4ca22a12997e88b1726b97.
//
// Solidity: event UnsupportedPackage(uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_CrossChain *CrossChainFilterer) WatchUnsupportedPackage(opts *bind.WatchOpts, sink chan<- *CrossChainUnsupportedPackage, packageSequence []uint64, channelId []uint8) (event.Subscription, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _CrossChain.contract.WatchLogs(opts, "UnsupportedPackage", packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainUnsupportedPackage)
				if err := _CrossChain.contract.UnpackLog(event, "UnsupportedPackage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnsupportedPackage is a log parse operation binding the contract event 0xdee9845a11ea343955bc9858cfa6d6fbcce9c8c1cc4ca22a12997e88b1726b97.
//
// Solidity: event UnsupportedPackage(uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_CrossChain *CrossChainFilterer) ParseUnsupportedPackage(log types.Log) (*CrossChainUnsupportedPackage, error) {
	event := new(CrossChainUnsupportedPackage)
	if err := _CrossChain.contract.UnpackLog(event, "UnsupportedPackage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DurationMetaData contains all meta data concerning the Duration contract.
var DurationMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220348be9cc4e01bbbae9a12ef5656aec3f2a29a8e5276b7ba3e9ceafaadf2b7d0464736f6c63430008120033",
}

// DurationABI is the input ABI used to generate the binding from.
// Deprecated: Use DurationMetaData.ABI instead.
var DurationABI = DurationMetaData.ABI

// DurationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DurationMetaData.Bin instead.
var DurationBin = DurationMetaData.Bin

// DeployDuration deploys a new Ethereum contract, binding an instance of Duration to it.
func DeployDuration(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Duration, error) {
	parsed, err := DurationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DurationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Duration{DurationCaller: DurationCaller{contract: contract}, DurationTransactor: DurationTransactor{contract: contract}, DurationFilterer: DurationFilterer{contract: contract}}, nil
}

// Duration is an auto generated Go binding around an Ethereum contract.
type Duration struct {
	DurationCaller     // Read-only binding to the contract
	DurationTransactor // Write-only binding to the contract
	DurationFilterer   // Log filterer for contract events
}

// DurationCaller is an auto generated read-only Go binding around an Ethereum contract.
type DurationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DurationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DurationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DurationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DurationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DurationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DurationSession struct {
	Contract     *Duration         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DurationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DurationCallerSession struct {
	Contract *DurationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DurationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DurationTransactorSession struct {
	Contract     *DurationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DurationRaw is an auto generated low-level Go binding around an Ethereum contract.
type DurationRaw struct {
	Contract *Duration // Generic contract binding to access the raw methods on
}

// DurationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DurationCallerRaw struct {
	Contract *DurationCaller // Generic read-only contract binding to access the raw methods on
}

// DurationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DurationTransactorRaw struct {
	Contract *DurationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDuration creates a new instance of Duration, bound to a specific deployed contract.
func NewDuration(address common.Address, backend bind.ContractBackend) (*Duration, error) {
	contract, err := bindDuration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Duration{DurationCaller: DurationCaller{contract: contract}, DurationTransactor: DurationTransactor{contract: contract}, DurationFilterer: DurationFilterer{contract: contract}}, nil
}

// NewDurationCaller creates a new read-only instance of Duration, bound to a specific deployed contract.
func NewDurationCaller(address common.Address, caller bind.ContractCaller) (*DurationCaller, error) {
	contract, err := bindDuration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DurationCaller{contract: contract}, nil
}

// NewDurationTransactor creates a new write-only instance of Duration, bound to a specific deployed contract.
func NewDurationTransactor(address common.Address, transactor bind.ContractTransactor) (*DurationTransactor, error) {
	contract, err := bindDuration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DurationTransactor{contract: contract}, nil
}

// NewDurationFilterer creates a new log filterer instance of Duration, bound to a specific deployed contract.
func NewDurationFilterer(address common.Address, filterer bind.ContractFilterer) (*DurationFilterer, error) {
	contract, err := bindDuration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DurationFilterer{contract: contract}, nil
}

// bindDuration binds a generic wrapper to an already deployed contract.
func bindDuration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DurationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Duration *DurationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Duration.Contract.DurationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Duration *DurationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Duration.Contract.DurationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Duration *DurationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Duration.Contract.DurationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Duration *DurationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Duration.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Duration *DurationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Duration.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Duration *DurationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Duration.Contract.contract.Transact(opts, method, params...)
}

// ECDSAMetaData contains all meta data concerning the ECDSA contract.
var ECDSAMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220674d2eb739903b5819fb956017332674f7b8f989f08bb12077cb498adcd62b7964736f6c63430008120033",
}

// ECDSAABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSAMetaData.ABI instead.
var ECDSAABI = ECDSAMetaData.ABI

// ECDSABin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ECDSAMetaData.Bin instead.
var ECDSABin = ECDSAMetaData.Bin

// DeployECDSA deploys a new Ethereum contract, binding an instance of ECDSA to it.
func DeployECDSA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSA, error) {
	parsed, err := ECDSAMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ECDSABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// ECDSA is an auto generated Go binding around an Ethereum contract.
type ECDSA struct {
	ECDSACaller     // Read-only binding to the contract
	ECDSATransactor // Write-only binding to the contract
	ECDSAFilterer   // Log filterer for contract events
}

// ECDSACaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSASession struct {
	Contract     *ECDSA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSACallerSession struct {
	Contract *ECDSACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ECDSATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSATransactorSession struct {
	Contract     *ECDSATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSARaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSARaw struct {
	Contract *ECDSA // Generic contract binding to access the raw methods on
}

// ECDSACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSACallerRaw struct {
	Contract *ECDSACaller // Generic read-only contract binding to access the raw methods on
}

// ECDSATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSATransactorRaw struct {
	Contract *ECDSATransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSA creates a new instance of ECDSA, bound to a specific deployed contract.
func NewECDSA(address common.Address, backend bind.ContractBackend) (*ECDSA, error) {
	contract, err := bindECDSA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// NewECDSACaller creates a new read-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSACaller(address common.Address, caller bind.ContractCaller) (*ECDSACaller, error) {
	contract, err := bindECDSA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACaller{contract: contract}, nil
}

// NewECDSATransactor creates a new write-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSATransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSATransactor, error) {
	contract, err := bindECDSA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATransactor{contract: contract}, nil
}

// NewECDSAFilterer creates a new log filterer instance of ECDSA, bound to a specific deployed contract.
func NewECDSAFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAFilterer, error) {
	contract, err := bindECDSA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAFilterer{contract: contract}, nil
}

// bindECDSA binds a generic wrapper to an already deployed contract.
func bindECDSA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ECDSAMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.ECDSACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transact(opts, method, params...)
}

// Ed25519VerifierMetaData contains all meta data concerning the Ed25519Verifier contract.
var Ed25519VerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commit\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[57]\",\"name\":\"input\",\"type\":\"uint256[57]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608080604052346100165761238b908161001c8239f35b600080fdfe604060808152600436101561001357600080fd5b6000803560e01c635db27c3d1461002957600080fd5b346100c7576108603660031901126100c75761004436610125565b9036606312156100c7576100566100e0565b908160c4913683116100c757506044905b8282106100ae576100aa86610099878761008036610167565b610089366101a5565b9161009336610225565b93610616565b905190151581529081906020820190565b0390f35b602086916100bc36856101e5565b815201910190610067565b80fd5b634e487b7160e01b600052604160045260246000fd5b604051906040820182811067ffffffffffffffff82111761010057604052565b6100ca565b604051906080820182811067ffffffffffffffff82111761010057604052565b8060231215610162576101366100e0565b90816044918211610162576004905b8282106101525750505090565b8135815260209182019101610145565b600080fd5b8060e31215610162576101786100e0565b90816101049182116101625760c4905b8282106101955750505090565b8135815260209182019101610188565b806101231215610162576101b76100e0565b908161014491821161016257610104905b8282106101d55750505090565b81358152602091820191016101c8565b9080601f83011215610162576101f96100e0565b80926040810192831161016257905b8282106102155750505090565b8135815260209182019101610208565b90610144916108641161016257565b61023c610105565b906080368337565b604051906060820182811067ffffffffffffffff821117610100576040526060368337565b604051906020820182811067ffffffffffffffff821117610100576040526020368337565b6102966100e0565b9061029f6100e0565b604036823782526102ae6100e0565b60403682376020830152565b6102c2610105565b906102cb6100e0565b60009081815281602082015283526102e161028e565b60208401526102ee6100e0565b81815281602082015260408401526103046100e0565b9080825260208201526060830152565b634e487b7160e01b600052603260045260246000fd5b1561033157565b60405162461bcd60e51b815260206004820152601760248201527f76657269666965722d61582d6774652d7072696d652d710000000000000000006044820152606490fd5b1561037d57565b60405162461bcd60e51b815260206004820152601760248201527f76657269666965722d61592d6774652d7072696d652d710000000000000000006044820152606490fd5b156103c957565b60405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258302d6774652d7072696d652d7100000000000000006044820152606490fd5b1561041557565b60405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259302d6774652d7072696d652d7100000000000000006044820152606490fd5b1561046157565b60405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6258312d6774652d7072696d652d7100000000000000006044820152606490fd5b156104ad57565b60405162461bcd60e51b815260206004820152601860248201527f76657269666965722d6259312d6774652d7072696d652d7100000000000000006044820152606490fd5b156104f957565b60405162461bcd60e51b815260206004820152601760248201527f76657269666965722d63582d6774652d7072696d652d710000000000000000006044820152606490fd5b1561054557565b60405162461bcd60e51b815260206004820152601760248201527f76657269666965722d63592d6774652d7072696d652d710000000000000000006044820152606490fd5b634e487b7160e01b600052601160045260246000fd5b60001981146105af5760010190565b61058a565b9060398110156105c55760051b0190565b610314565b156105d157565b60405162461bcd60e51b815260206004820152601f60248201527f76657269666965722d6774652d736e61726b2d7363616c61722d6669656c64006044820152606490fd5b949394929190926106256102ba565b81516020809301516106356100e0565b9182528382015281526106466100e0565b948051518652610657815160200190565b5183870152610676836106686100e0565b920180515183525160200190565b51838201526106836100e0565b9586528286015281810194855281835193015161069e6100e0565b93845282840152604081019283528351938281019485516106bd6100e0565b91825284820152606083019081526107717f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd476106fc818651511061032a565b61070b81878751015110610376565b610719818a515151106103c2565b61072981878b510151511061040e565b610740816107398b515160200190565b511061045a565b61075981610752888c51015160200190565b51106104a6565b61076681885151106104f2565b85875101511061053e565b60005b60398110611c70575090611c49979893929161078e61208c565b966107976100e0565b9260008452611c10848781019860008a5281996107b2610234565b918a6107006107bf610244565b927f195db37264b52fdad37dd63b1b1c610b5c5a471671c8365c3b74c2b29df020dc6107e96100e0565b95600087526000858801527f14eb8e3c7fef63d5ea028a5002f8175cb05814ddd0e2c8e8bc4e2cda415726328952527f05db00ee424ba23a6870e9da82166fc3e8d842609ae5db81c37b81ade701ec2584527f130b90aaf71fc9339e7eda4bfc03468c60bf82ef406ad95844a88086a619b8918385017f1accd6fc21b68d3866c361ccb2903b3dbab8ac4554e6938b44089b09403c51ee81526040860194833586526108978a8a8a8a612308565b7f0f9c9b8610c2855fbe0d5e355a20f1db41775d76a99fe9700efc965da6f10fe587527f1f462a880c0e1408e52ada883ade15bd9a3ee43ad19f1369b1e38e6abbca9721825283013585526108ee89898989612308565b7f122b3ade7f9a0d55e1faa3b56ca7db21f9cb816fb156a0fcfe7ea3753b98067186527f0aa190b82afc48b57b1bb7a3eddce4b88561de39d201bc789bc19fc640291ac881526040830135855261094789898989612308565b7f0d70669e3c09792c003e0d08f76e71a3b26d46bf42c18db8e0e5cacc75392ace86527f1f7382a8da98df4de8b9e3c793b90e3e2876e99b384cd62274a3452a788178c78152606083013585526109a089898989612308565b7f21a970c5ffca6adbf42340020624312695dc7329f7d980ad74689ba893d45b9386527f226bfa0e7dd0852c51ec3c5a90b4d2d94dcf459f2cbdb4496ddc96129d05011d8152608083013585526109f989898989612308565b7f08b7408c9452bb9b467edaea5c590d9644addea75a405196fb97a123f1628eef86527f09261f8ca6de306c599fcdfb1d8cc6a8fe77fd16d6c07211f6669579faf1207b815260a08301358552610a5289898989612308565b7f0199be4cbdf300ccba6675bc195df46fe600c3c7763e94b243d08caef7584a8d86527f2e60f0a6bff6f78d2bcdfca37c611d3b519431f571fbfbf9958db7e38b48aea4815260c08301358552610aab89898989612308565b7f2133a681b4751f70c9e30cf6d6173e09654a676f0af860b4c2967757c1e15a2786527f02cc868c8214f13e5f0965aab8606660c846988afac9b9540d43a1885d161afa815260e08301358552610b0489898989612308565b7f0b8b79c2908ff8d2af7b948bfb469167cc81c7a7c0c0db21318120d4b689bf6686527f1fe55313297bfce728d5d96c4263dcd01969b0bbb334982bfe9f49bfa963d06f81526101008301358552610b5e89898989612308565b7f2d7bf1de380b766154432deb8c5d04bb0a5ff66bd4a21c5ce62d54a76399e4b686527f2a48784d10692e6f81716ec66dcac0850d4c3b3f0254c1f7f63a05daa97aaf6281526101208301358552610bb889898989612308565b7f0171a56c0d9a03591fc0c9c1d8daaad2d87961408066de2bd50fa3bb5412c4b386527f0209128072b2aa0176f4449dedb93536a70217e6daca1dd069db90baed048a8a81526101408301358552610c1289898989612308565b7f2e3ef6f35dbab4b60a4a43fc042ddbdb905d5a5d0906a97eca44d2cfb2eefbff86527f124468405f3e7c05c68f8d6b3228c66a8110fbbeea9b6af1ee9b526031b62eea81526101608301358552610c6c89898989612308565b7f2f45ffefcefab959648438d7f52a7c1d142024a3f1c675e962c656743685a53286527f2ce411430951557660a5d21539a36bda9e9cce625f477c39f8601c58f196d3f681526101808301358552610cc689898989612308565b7f0db853bc60253ad9a15075ba631c2913d0568613a590bbd54f721195868138c986527f205369d4cb13aad77b748657b616885f77dd6f2e971dd57f33246e8becd504e681526101a08301358552610d2089898989612308565b7f22f694ba1c3777f20d8ba1964c35bff2bd1f89ec088660277d824c4e66105be986527f2dcf046680ca92b1e87541e5e915b821f63f494090f8a63c93ce5928c8c1e02381526101c08301358552610d7a89898989612308565b7f1cff1e41e33a67b9387cffc02261bb6c5cb7d712f479024c66f7ee290fdab53186527f02f20f4124852753ebb6dc4d0d1ed74211b1b2b6d0f8d9b8adc90410723a410981526101e08301358552610dd489898989612308565b7f304437e6b3a43fba395ef07e461656188f8dd3b1a68c8a253666bc2800fe802186527f0ee377f37e028bc805d729e52a8d0b918f91afff9ea84abab1b783f2acfcf01d81526102008301358552610e2e89898989612308565b7f2bd404584ec18f2a34af984c4e53342fca08fd2489dd420cf9b6a0d7bd54aa0486527f26917439d6eb592ab4b454d7737189d6f49b9982a12b1a57d23cd40310ddbcec81526102208301358552610e8889898989612308565b7f18c52704bab4fd4b9f8f838fb8ca08b28bf3b0b8e938f226866e49c275c4fcb886527f23b388966e739c3d5ce32d39865925618a8b43d6a5819a47ae356794233cdf8581526102408301358552610ee289898989612308565b7f2f4014ef261fb7dbd305ea2d98b50fd94a957c89d6de5bf715710b195dca05ec86527f2808b9f7f05544339dfec4413ee827bd3ab1793ad0b3d3a3f101fa3bb1c067ac81526102608301358552610f3c89898989612308565b7f0498945d70f982392cfa9422c9d613a97f661c476cfd1078011e1fa3244d51dd86527f2cdb300e6badb5aeaf81d6f56fa21539ce536db04656c5b002cc2f7fa070525081526102808301358552610f9689898989612308565b7f2dbb073fa40e87ef96b7a9a1499eb9f217fb6ec00f5e045053236c7c63bf6a2786527f08a10d4acceae1d856a27714640c631b6018aa2aaf3858bc4534d9efd56f7bee81526102a08301358552610ff089898989612308565b7f1ea1704383792942357529b9f09083364551c12e66dc5afccdf02091275e9b1586527f02dbd34b7937de2538c4da0fdda426c7584173eadeeb5ddf9a0453185218607781526102c0830135855261104a89898989612308565b7f086a387160ccc10a8bdfa884c046357a43bf2b6302d198e002ff06013713226986527f19257222baaff5e36ad4128fd0e5ff8b3462426e789a408bacbd7f759b18bee881526102e083013585526110a489898989612308565b7f068946bd2571e14f50d418b899abcb1d41bb021e28e9283a35abd3315b5c6fdd86527f024acd68f1dc77a4287de542f40c8ee44a39f4de77ec48ce4ba6b74b4782b6ab815261030083013585526110fe89898989612308565b7f07c9efff640383e690716625236390d5a0ca94196c6130078e0df871163eda4c86527f1d043ab52c1a5ec6e9bde294ab128d9fad7ba567fb55f21130afd70ff93206998152610320830135855261115889898989612308565b7f168eb4a4ebdb1440038183f21277b11f4d748c74245e8df71c71372cabdc086786527f2a7e40549727d6ec39286c8aa11705b2a4c11d7f18a6e65d567811282ee17178815261034083013585526111b289898989612308565b7f1e314b1ae648e8b65116480b3de428843c92aca14a3743ae075a9a6250fd09ec86527f293dbaf88138010a710965910a0dedca7a2174a90d7663bd1c024a9c83bd29b98152610360830135855261120c89898989612308565b7f2db15dc5852aae3351b672c7f136a5afb97d6dbea6546667e95ccf3001c58dc486527f167a0c774e634a0ee7931c0fd8c3ae56342df4dc3fc6422b71ac98b120ca5b7a8152610380830135855261126689898989612308565b7f1ff34eeba697ad0bf57ff8c082955d544a5435db79aa253cebdae4bf2cef778086527f22a912043ac284e40b2e755e482adae3978d1f25e3c46ca4bf4eedad04fce87981526103a083013585526112c089898989612308565b7f137d4eca0594c11dfaaaa38f28c46669ee0f6910108e516c7e4539e1ee91ff7486527f1cc32eec8a37e6f4abca3d34e7bd9ec432bc55edd4911a2d03c9294d7447bee681526103c0830135855261131a89898989612308565b7f113b8dfa95300e399ad25ccdf0b3f3855db841716c1fbb631f42340f3b4ae88786527f1e18c9ecdde07d20bc74ea79b4a22c9cea2a4c9624b6728910ae5e225a93a86481526103e0830135855261137489898989612308565b7f1e92f8f16734a7f800605901175ee66d6b1bf27ac4354b8134a4b682c98ef7cf86527f031541443d700d8e2c063c4b3b430d711f6e8c464e2780fada0a179700465f7b815261040083013585526113ce89898989612308565b7f1d4b8187c2cfe78f825efc2bb3dc71cef5a01a6168c24f0f8cd3d115e2be85c286527f0cead51f98c0329afb7cb33dcbc52f02f2a7f581b57a7ccb39614b63816b39e78152610420830135855261142889898989612308565b7f09623e4e1631f799dda9228297d35c7596e024e59557d03ea1304cd47426cbf386527f0c4ee07bd5372f0aad83d6677d9f3aa230288a0e65bbe53ea67794d619e148068152610440830135855261148289898989612308565b7f2468ade79ea8d3a39533e023376e01c85beacf21b21320377e50152d293222b086527f1cc67a6a35c248e119ef792e75e5561e0ee02754cb4226508ad4ce905477e536815261046083013585526114dc89898989612308565b7f03c81c307cdbd755daea8e0436c7e5405a080d74a05fe8f9bbdbc4332b8d0a6086527f0eaef820d031e2abe96f48c098b29741bc8b6a559a99034b6452a87e28fd787c8152610480830135855261153689898989612308565b7f1eac3ab4946bf7f27e5ce971407bde34a52ca00833c136acddbc1353b63a934e86527f06109092d0c3f7ad2514d5a774602486dee5f7031a0886e023718bd5245aff3681526104a0830135855261159089898989612308565b7f0c317f3e48f1b9fbdc3d3b94797cc261d51d69bdb492d6c49ec7304a125957c386527f0f1a1f42fcb133223355dc983815e0dcad482e1a28e5cd147e510203df44a22081526104c083013585526115ea89898989612308565b7f1ab77b70ec2ea962347077937b7f1f83d5e0c4c1eb285a497225fc8630a42c5386527f2bd3f3f3f4b639cd4968e55a8141b1fab2e18f2bafaa5b770473453f742ad50581526104e0830135855261164489898989612308565b7f03c6acc776c139d158ab4acf5eef063e4ab5ab3e84abf4fba0bf3a21b82c872586527f048dd28f86217b675c28c29aedf9e60fb33a170ca4e38f8cc224df7dcc40e09f8152610500830135855261169e89898989612308565b7f0995e900237757017e97af946aae0a8ea2b98820b012657c197e4a6894b0ecd486527f0fd68b6c5b05cfee9beb0bfd7f254125d3e721b7f78ed9a142a91ffb435ded6b815261052083013585526116f889898989612308565b7f29136a3af7471ef1102bba59fe8f95f674ade4d040f97921c007bf9b531f2bdb86527f1063dbb063b9d65d71f6d8cc7d4f17ddfaa309a0e9553ca7f87558bf799be6138152610540830135855261175289898989612308565b7f2377bdc36cd2ed7eb0d08aa65256fc8be51019b4b065c6218ff3dc40b08ab81b86527f14498d19b7545adf2f72ee1ccdf7649300a96c9facd5b58a403da754aca5db75815261056083013585526117ac89898989612308565b7f0a2ece1ea46e366429edcfd56a82edf3b5eb600bfd49c71aa086f79fbdaee83a86527f245503a87b8723a03f8ab337ab0424e393884f16a02ee4114243cc6c7a988e0a8152610580830135855261180689898989612308565b7f2f129f51f8e706e6f249e230ef2cf541b1adfc7424abd898dd0576d78919464386527f1ea599c58c22453bbaa3e9cd49af517dc8ff78a0de84f6c04fb829c9c89f42ea81526105a0830135855261186089898989612308565b7f1e16290678e28ff123e036090b58a6835532ce9b670a7acbcadebfb8c7c219f186527f2f74c1793b858f11089ba108bd5023d5d14870ae79089a9a041e1bbf6428298581526105c083013585526118ba89898989612308565b7f2013ab2bc100fc703e5a72500c6bfee143620f8c1321ebf32cb2f6820351280786527f1ceedec39d9436e9e1716a82e86e134c04f055fb105d30ea5602d0c4d6315eab81526105e0830135855261191489898989612308565b7f2b60c70de94e3f33aadf90c5e5a289ea7941a25769cc24ba5225c242b305273386527f1ddf9a6ca2ba3017934b0ed416d6074e5f62bab181515a582509a7d23239219d8152610600830135855261196e89898989612308565b7f2c594d9dcf90c8da15c8b09bb5cc8ebde70c436a499d894d815f9c7cf250336486527f1aa415195e03d6d51355848081a52370fcf2bbb925994cf41b83be01a24ad5d8815261062083013585526119c889898989612308565b7f1d6b37fbcd0650f69221ea6c572e624debcc7975b30e6dcdc5f33ff908b958f386527f017e4b3cebdb13756d35d20bf2a84c921ae4ebeaee2250f70b0f50f15bb08b5d81526106408301358552611a2289898989612308565b7f1ab2d40157cc4b25b3f5338b75963b24779628d6d753d8cad0915ecd7a2a8ddd86527f2c5b4c201408b7aa4f342b55c7d4aca63099e7a21d61b75d71ecec6d843b748f81526106608301358552611a7c89898989612308565b7f0fa2009caab26e708508e54ccfb9eba091cf30fdd350ded858723bfc927960ce86527f29241cd41494a0cb2456a413bd9c7224250e5c8f55cef9295094256932ec8dba81526106808301358552611ad689898989612308565b7f136daf1a59100e34f70af2b1174f2bf69bbd53912993227babd9200beed16dd986527e9ef60dee5c42ad366a30d096e9a48481898715575521930d19aad99e54a88f81526106a08301358552611b2f89898989612308565b7f26a0151ba75e58ea181e318eb28629edc97fa9ba989d424167461622c48bf90686527f187dc0ed3ed9fa3546f335ecc7d616745e8d37c7668ffa2a0195607b604a048481526106c08301358552611b8989898989612308565b7f299c283d820e5a8365242a9becfac6eaf9c24c2c2ceadc6ef8acfc107ff229cc86527f28ad291c0dce9ed902f42dc97b33039d87b9baabc225bb8930cec7e46219097481526106e08301358552611be389898989612308565b7f18466fb853ddf087603bd9b2c31855e5d5cd6c6c2ad85e1e9ef1d4189770e18b86525201359052612308565b511590811591611c65575b50611c4c575b5050611c2d9051611cb3565b9451908451908501519160606040870151955196015196611ee6565b90565b51919350611c2d91611c5d91611d9a565b929038611c21565b905051151538611c1b565b80611ca97f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001611ca2611cae948e6105b4565b35106105ca565b6105a0565b610774565b60006020611cbf6100e0565b8281520152805190811580611d31575b15611cec575050611cde6100e0565b600081526000602082015290565b602001517f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd479081900681039081116105af57611d266100e0565b918252602082015290565b50602081015115611ccf565b9060048110156105c55760051b0190565b15611d5557565b60405162461bcd60e51b815260206004820152601260248201527f70616972696e672d6164642d6661696c656400000000000000000000000000006044820152606490fd5b60609092919260c0611daa6100e0565b9160008352600060208401526020839681611dc3610105565b936080368637805185520151828401528051604084015201518482015260066107cf195a01fa8015611dfa57611df890611d4e565b565bfe5b60405190610320820182811067ffffffffffffffff8211176101005760405260188252610300366020840137565b906006820291808304600614901517156105af57565b90600182018092116105af57565b90600282018092116105af57565b90600382018092116105af57565b90600482018092116105af57565b90600582018092116105af57565b80518210156105c55760209160051b010190565b15611ea157565b60405162461bcd60e51b815260206004820152601560248201527f70616972696e672d6f70636f64652d6661696c656400000000000000000000006044820152606490fd5b9491959692909396611ef6610105565b95865260209788978888015260408701526060860152611f14610105565b9384528584015260408301526060820152611f2d611dfc565b9160005b60048110611f6957505050610300611f47610269565b9384920160086107cf195a01fa8015611dfa57611f6390611e9a565b51151590565b61203b9192939450611f7a81611e2a565b611f848285611d3d565b5151611f908288611e86565b5286611f9c8386611d3d565b510151611fb1611fab83611e40565b88611e86565b52611fbc8286611d3d565b515151611fcb611fab83611e4e565b52611fe1611fd98387611d3d565b515160200190565b51611fee611fab83611e5c565b5286611ffa8387611d3d565b5101515161200a611fab83611e6a565b5261203561202f6120288961201f868a611d3d565b51015160200190565b5192611e78565b87611e86565b526105a0565b9084939291611f31565b61204d610105565b906120566100e0565b6000815260006020820152825261206b61028e565b602083015261207861028e565b604083015261208561028e565b6060830152565b612094612045565b9061209d6100e0565b7f248fee176e10478ce9565376076110465a92175ec51de6e0c92da5a0b4dea0e281526020907f273ba3729b38054dc06319a540cb2c8062ff8800e9f986d7de16418c84f2816b8282015283526120f26100e0565b7f294d7b428d8916493bb2bc2be1037c6f1f2d97bbaa5bd5c1e2500b4e52b8fdd981527f1af577b195b03b5edc4d753c3d436613dd86d83be91ed668bef16097f8493edf828201526121426100e0565b7f0286bd43af5b7313e07483e3d59820cee25f9d6ff5b8f6971628cd3eef1e8abf81527f2ed186d8c7887e4dd020a7bdffe9ced507d6e0fbfb544f72a8e4b1614e46d299838201526121926100e0565b91825282820152818401526121a56100e0565b7f2bc6004658e78f0e22f3f98dc5fd1f4c42ba601d40e71382b44959a6ea83e12981527f186264fd247bbdeaf41deb352b12b7cccbcd6cda2f96e6ca8588ff6bebca9a4c828201526121f56100e0565b7f20d6bbde3c1db951d6f24a4d9f63805e3f9cbd32e58de737890d402d06336baf81527f031e20e90059227edd2ed33e2f51ca84c40e68cd5384b69ab2f26a28c1fd1bd9838201526122456100e0565b9182528282015260408401526122596100e0565b907f1f069585bc15e9f52a6952370f9e16cb640e1df20ef17a42fae819368165e70882527f03fce6f53c1c85c8a7500f9a4b356c66fc8cc08419b09fa1f13e5d577515b3c4818301526122aa6100e0565b907f2a65cedff74fa1b36a6f5e43ad2c47b6781c3c7cf75a08d0d75336ca09c65b3582527f1dd3d3c12f9f47e41dc610575fd32b17d14a9a3115260255512c793a18fab52a818301526122fb6100e0565b9283528201526060830152565b90929160608460806107cf19946007865a01fa15611dfa57600660c0926020606096865185528187015182860152805160408601520151868401525a01fa8015611dfa57611df890611d4e56fea26469706673582212200ea00be38d2aa12baa86b25589e8bf39b1c7e962906c8b6ae889be08dcaed84e64736f6c63430008120033",
}

// Ed25519VerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use Ed25519VerifierMetaData.ABI instead.
var Ed25519VerifierABI = Ed25519VerifierMetaData.ABI

// Ed25519VerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Ed25519VerifierMetaData.Bin instead.
var Ed25519VerifierBin = Ed25519VerifierMetaData.Bin

// DeployEd25519Verifier deploys a new Ethereum contract, binding an instance of Ed25519Verifier to it.
func DeployEd25519Verifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ed25519Verifier, error) {
	parsed, err := Ed25519VerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Ed25519VerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ed25519Verifier{Ed25519VerifierCaller: Ed25519VerifierCaller{contract: contract}, Ed25519VerifierTransactor: Ed25519VerifierTransactor{contract: contract}, Ed25519VerifierFilterer: Ed25519VerifierFilterer{contract: contract}}, nil
}

// Ed25519Verifier is an auto generated Go binding around an Ethereum contract.
type Ed25519Verifier struct {
	Ed25519VerifierCaller     // Read-only binding to the contract
	Ed25519VerifierTransactor // Write-only binding to the contract
	Ed25519VerifierFilterer   // Log filterer for contract events
}

// Ed25519VerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type Ed25519VerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ed25519VerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Ed25519VerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ed25519VerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ed25519VerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ed25519VerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ed25519VerifierSession struct {
	Contract     *Ed25519Verifier  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ed25519VerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ed25519VerifierCallerSession struct {
	Contract *Ed25519VerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// Ed25519VerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ed25519VerifierTransactorSession struct {
	Contract     *Ed25519VerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Ed25519VerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type Ed25519VerifierRaw struct {
	Contract *Ed25519Verifier // Generic contract binding to access the raw methods on
}

// Ed25519VerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ed25519VerifierCallerRaw struct {
	Contract *Ed25519VerifierCaller // Generic read-only contract binding to access the raw methods on
}

// Ed25519VerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ed25519VerifierTransactorRaw struct {
	Contract *Ed25519VerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEd25519Verifier creates a new instance of Ed25519Verifier, bound to a specific deployed contract.
func NewEd25519Verifier(address common.Address, backend bind.ContractBackend) (*Ed25519Verifier, error) {
	contract, err := bindEd25519Verifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ed25519Verifier{Ed25519VerifierCaller: Ed25519VerifierCaller{contract: contract}, Ed25519VerifierTransactor: Ed25519VerifierTransactor{contract: contract}, Ed25519VerifierFilterer: Ed25519VerifierFilterer{contract: contract}}, nil
}

// NewEd25519VerifierCaller creates a new read-only instance of Ed25519Verifier, bound to a specific deployed contract.
func NewEd25519VerifierCaller(address common.Address, caller bind.ContractCaller) (*Ed25519VerifierCaller, error) {
	contract, err := bindEd25519Verifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ed25519VerifierCaller{contract: contract}, nil
}

// NewEd25519VerifierTransactor creates a new write-only instance of Ed25519Verifier, bound to a specific deployed contract.
func NewEd25519VerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*Ed25519VerifierTransactor, error) {
	contract, err := bindEd25519Verifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ed25519VerifierTransactor{contract: contract}, nil
}

// NewEd25519VerifierFilterer creates a new log filterer instance of Ed25519Verifier, bound to a specific deployed contract.
func NewEd25519VerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*Ed25519VerifierFilterer, error) {
	contract, err := bindEd25519Verifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ed25519VerifierFilterer{contract: contract}, nil
}

// bindEd25519Verifier binds a generic wrapper to an already deployed contract.
func bindEd25519Verifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Ed25519VerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ed25519Verifier *Ed25519VerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ed25519Verifier.Contract.Ed25519VerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ed25519Verifier *Ed25519VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ed25519Verifier.Contract.Ed25519VerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ed25519Verifier *Ed25519VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ed25519Verifier.Contract.Ed25519VerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ed25519Verifier *Ed25519VerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ed25519Verifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ed25519Verifier *Ed25519VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ed25519Verifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ed25519Verifier *Ed25519VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ed25519Verifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x5db27c3d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] commit, uint256[57] input) view returns(bool r)
func (_Ed25519Verifier *Ed25519VerifierCaller) VerifyProof(opts *bind.CallOpts, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, commit [2]*big.Int, input [57]*big.Int) (bool, error) {
	var out []interface{}
	err := _Ed25519Verifier.contract.Call(opts, &out, "verifyProof", a, b, c, commit, input)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x5db27c3d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] commit, uint256[57] input) view returns(bool r)
func (_Ed25519Verifier *Ed25519VerifierSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, commit [2]*big.Int, input [57]*big.Int) (bool, error) {
	return _Ed25519Verifier.Contract.VerifyProof(&_Ed25519Verifier.CallOpts, a, b, c, commit, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x5db27c3d.
//
// Solidity: function verifyProof(uint256[2] a, uint256[2][2] b, uint256[2] c, uint256[2] commit, uint256[57] input) view returns(bool r)
func (_Ed25519Verifier *Ed25519VerifierCallerSession) VerifyProof(a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, commit [2]*big.Int, input [57]*big.Int) (bool, error) {
	return _Ed25519Verifier.Contract.VerifyProof(&_Ed25519Verifier.CallOpts, a, b, c, commit, input)
}

// EncoderMetaData contains all meta data concerning the Encoder contract.
var EncoderMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220cb6b12ab67ea95280ae7aa9b41cc78c3506a0cea5f5fff2daa231ea0c529719264736f6c63430008120033",
}

// EncoderABI is the input ABI used to generate the binding from.
// Deprecated: Use EncoderMetaData.ABI instead.
var EncoderABI = EncoderMetaData.ABI

// EncoderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EncoderMetaData.Bin instead.
var EncoderBin = EncoderMetaData.Bin

// DeployEncoder deploys a new Ethereum contract, binding an instance of Encoder to it.
func DeployEncoder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Encoder, error) {
	parsed, err := EncoderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EncoderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Encoder{EncoderCaller: EncoderCaller{contract: contract}, EncoderTransactor: EncoderTransactor{contract: contract}, EncoderFilterer: EncoderFilterer{contract: contract}}, nil
}

// Encoder is an auto generated Go binding around an Ethereum contract.
type Encoder struct {
	EncoderCaller     // Read-only binding to the contract
	EncoderTransactor // Write-only binding to the contract
	EncoderFilterer   // Log filterer for contract events
}

// EncoderCaller is an auto generated read-only Go binding around an Ethereum contract.
type EncoderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EncoderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EncoderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EncoderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EncoderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EncoderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EncoderSession struct {
	Contract     *Encoder          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EncoderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EncoderCallerSession struct {
	Contract *EncoderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// EncoderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EncoderTransactorSession struct {
	Contract     *EncoderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EncoderRaw is an auto generated low-level Go binding around an Ethereum contract.
type EncoderRaw struct {
	Contract *Encoder // Generic contract binding to access the raw methods on
}

// EncoderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EncoderCallerRaw struct {
	Contract *EncoderCaller // Generic read-only contract binding to access the raw methods on
}

// EncoderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EncoderTransactorRaw struct {
	Contract *EncoderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEncoder creates a new instance of Encoder, bound to a specific deployed contract.
func NewEncoder(address common.Address, backend bind.ContractBackend) (*Encoder, error) {
	contract, err := bindEncoder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Encoder{EncoderCaller: EncoderCaller{contract: contract}, EncoderTransactor: EncoderTransactor{contract: contract}, EncoderFilterer: EncoderFilterer{contract: contract}}, nil
}

// NewEncoderCaller creates a new read-only instance of Encoder, bound to a specific deployed contract.
func NewEncoderCaller(address common.Address, caller bind.ContractCaller) (*EncoderCaller, error) {
	contract, err := bindEncoder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EncoderCaller{contract: contract}, nil
}

// NewEncoderTransactor creates a new write-only instance of Encoder, bound to a specific deployed contract.
func NewEncoderTransactor(address common.Address, transactor bind.ContractTransactor) (*EncoderTransactor, error) {
	contract, err := bindEncoder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EncoderTransactor{contract: contract}, nil
}

// NewEncoderFilterer creates a new log filterer instance of Encoder, bound to a specific deployed contract.
func NewEncoderFilterer(address common.Address, filterer bind.ContractFilterer) (*EncoderFilterer, error) {
	contract, err := bindEncoder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EncoderFilterer{contract: contract}, nil
}

// bindEncoder binds a generic wrapper to an already deployed contract.
func bindEncoder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EncoderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Encoder *EncoderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Encoder.Contract.EncoderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Encoder *EncoderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Encoder.Contract.EncoderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Encoder *EncoderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Encoder.Contract.EncoderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Encoder *EncoderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Encoder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Encoder *EncoderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Encoder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Encoder *EncoderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Encoder.Contract.contract.Transact(opts, method, params...)
}

// ExistenceProofMetaData contains all meta data concerning the ExistenceProof contract.
var ExistenceProofMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220fee51a524ab0710e15730d0c1143a82b69724c3f1490dac92accbe9aac5016d864736f6c63430008120033",
}

// ExistenceProofABI is the input ABI used to generate the binding from.
// Deprecated: Use ExistenceProofMetaData.ABI instead.
var ExistenceProofABI = ExistenceProofMetaData.ABI

// ExistenceProofBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExistenceProofMetaData.Bin instead.
var ExistenceProofBin = ExistenceProofMetaData.Bin

// DeployExistenceProof deploys a new Ethereum contract, binding an instance of ExistenceProof to it.
func DeployExistenceProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExistenceProof, error) {
	parsed, err := ExistenceProofMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExistenceProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExistenceProof{ExistenceProofCaller: ExistenceProofCaller{contract: contract}, ExistenceProofTransactor: ExistenceProofTransactor{contract: contract}, ExistenceProofFilterer: ExistenceProofFilterer{contract: contract}}, nil
}

// ExistenceProof is an auto generated Go binding around an Ethereum contract.
type ExistenceProof struct {
	ExistenceProofCaller     // Read-only binding to the contract
	ExistenceProofTransactor // Write-only binding to the contract
	ExistenceProofFilterer   // Log filterer for contract events
}

// ExistenceProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExistenceProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExistenceProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExistenceProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExistenceProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExistenceProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExistenceProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExistenceProofSession struct {
	Contract     *ExistenceProof   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExistenceProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExistenceProofCallerSession struct {
	Contract *ExistenceProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ExistenceProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExistenceProofTransactorSession struct {
	Contract     *ExistenceProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ExistenceProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExistenceProofRaw struct {
	Contract *ExistenceProof // Generic contract binding to access the raw methods on
}

// ExistenceProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExistenceProofCallerRaw struct {
	Contract *ExistenceProofCaller // Generic read-only contract binding to access the raw methods on
}

// ExistenceProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExistenceProofTransactorRaw struct {
	Contract *ExistenceProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExistenceProof creates a new instance of ExistenceProof, bound to a specific deployed contract.
func NewExistenceProof(address common.Address, backend bind.ContractBackend) (*ExistenceProof, error) {
	contract, err := bindExistenceProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExistenceProof{ExistenceProofCaller: ExistenceProofCaller{contract: contract}, ExistenceProofTransactor: ExistenceProofTransactor{contract: contract}, ExistenceProofFilterer: ExistenceProofFilterer{contract: contract}}, nil
}

// NewExistenceProofCaller creates a new read-only instance of ExistenceProof, bound to a specific deployed contract.
func NewExistenceProofCaller(address common.Address, caller bind.ContractCaller) (*ExistenceProofCaller, error) {
	contract, err := bindExistenceProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExistenceProofCaller{contract: contract}, nil
}

// NewExistenceProofTransactor creates a new write-only instance of ExistenceProof, bound to a specific deployed contract.
func NewExistenceProofTransactor(address common.Address, transactor bind.ContractTransactor) (*ExistenceProofTransactor, error) {
	contract, err := bindExistenceProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExistenceProofTransactor{contract: contract}, nil
}

// NewExistenceProofFilterer creates a new log filterer instance of ExistenceProof, bound to a specific deployed contract.
func NewExistenceProofFilterer(address common.Address, filterer bind.ContractFilterer) (*ExistenceProofFilterer, error) {
	contract, err := bindExistenceProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExistenceProofFilterer{contract: contract}, nil
}

// bindExistenceProof binds a generic wrapper to an already deployed contract.
func bindExistenceProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExistenceProofMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExistenceProof *ExistenceProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExistenceProof.Contract.ExistenceProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExistenceProof *ExistenceProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExistenceProof.Contract.ExistenceProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExistenceProof *ExistenceProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExistenceProof.Contract.ExistenceProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExistenceProof *ExistenceProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExistenceProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExistenceProof *ExistenceProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExistenceProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExistenceProof *ExistenceProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExistenceProof.Contract.contract.Transact(opts, method, params...)
}

// FractionMetaData contains all meta data concerning the Fraction contract.
var FractionMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220c659a68b1cbbf8299fc1635b6a1e0ee2502818af41e4299ffa9963facdf2a30564736f6c63430008120033",
}

// FractionABI is the input ABI used to generate the binding from.
// Deprecated: Use FractionMetaData.ABI instead.
var FractionABI = FractionMetaData.ABI

// FractionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FractionMetaData.Bin instead.
var FractionBin = FractionMetaData.Bin

// DeployFraction deploys a new Ethereum contract, binding an instance of Fraction to it.
func DeployFraction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Fraction, error) {
	parsed, err := FractionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FractionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Fraction{FractionCaller: FractionCaller{contract: contract}, FractionTransactor: FractionTransactor{contract: contract}, FractionFilterer: FractionFilterer{contract: contract}}, nil
}

// Fraction is an auto generated Go binding around an Ethereum contract.
type Fraction struct {
	FractionCaller     // Read-only binding to the contract
	FractionTransactor // Write-only binding to the contract
	FractionFilterer   // Log filterer for contract events
}

// FractionCaller is an auto generated read-only Go binding around an Ethereum contract.
type FractionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FractionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FractionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FractionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FractionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FractionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FractionSession struct {
	Contract     *Fraction         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FractionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FractionCallerSession struct {
	Contract *FractionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FractionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FractionTransactorSession struct {
	Contract     *FractionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FractionRaw is an auto generated low-level Go binding around an Ethereum contract.
type FractionRaw struct {
	Contract *Fraction // Generic contract binding to access the raw methods on
}

// FractionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FractionCallerRaw struct {
	Contract *FractionCaller // Generic read-only contract binding to access the raw methods on
}

// FractionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FractionTransactorRaw struct {
	Contract *FractionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFraction creates a new instance of Fraction, bound to a specific deployed contract.
func NewFraction(address common.Address, backend bind.ContractBackend) (*Fraction, error) {
	contract, err := bindFraction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fraction{FractionCaller: FractionCaller{contract: contract}, FractionTransactor: FractionTransactor{contract: contract}, FractionFilterer: FractionFilterer{contract: contract}}, nil
}

// NewFractionCaller creates a new read-only instance of Fraction, bound to a specific deployed contract.
func NewFractionCaller(address common.Address, caller bind.ContractCaller) (*FractionCaller, error) {
	contract, err := bindFraction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FractionCaller{contract: contract}, nil
}

// NewFractionTransactor creates a new write-only instance of Fraction, bound to a specific deployed contract.
func NewFractionTransactor(address common.Address, transactor bind.ContractTransactor) (*FractionTransactor, error) {
	contract, err := bindFraction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FractionTransactor{contract: contract}, nil
}

// NewFractionFilterer creates a new log filterer instance of Fraction, bound to a specific deployed contract.
func NewFractionFilterer(address common.Address, filterer bind.ContractFilterer) (*FractionFilterer, error) {
	contract, err := bindFraction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FractionFilterer{contract: contract}, nil
}

// bindFraction binds a generic wrapper to an already deployed contract.
func bindFraction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FractionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fraction *FractionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fraction.Contract.FractionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fraction *FractionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fraction.Contract.FractionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fraction *FractionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fraction.Contract.FractionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fraction *FractionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fraction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fraction *FractionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fraction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fraction *FractionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fraction.Contract.contract.Transact(opts, method, params...)
}

// GoogleProtobufAnyMetaData contains all meta data concerning the GoogleProtobufAny contract.
var GoogleProtobufAnyMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220bae104b74fd1557405247405543de03ebfc4d4854e04a1d2f1023bfb7eb98e7a64736f6c63430008120033",
}

// GoogleProtobufAnyABI is the input ABI used to generate the binding from.
// Deprecated: Use GoogleProtobufAnyMetaData.ABI instead.
var GoogleProtobufAnyABI = GoogleProtobufAnyMetaData.ABI

// GoogleProtobufAnyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GoogleProtobufAnyMetaData.Bin instead.
var GoogleProtobufAnyBin = GoogleProtobufAnyMetaData.Bin

// DeployGoogleProtobufAny deploys a new Ethereum contract, binding an instance of GoogleProtobufAny to it.
func DeployGoogleProtobufAny(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GoogleProtobufAny, error) {
	parsed, err := GoogleProtobufAnyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GoogleProtobufAnyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GoogleProtobufAny{GoogleProtobufAnyCaller: GoogleProtobufAnyCaller{contract: contract}, GoogleProtobufAnyTransactor: GoogleProtobufAnyTransactor{contract: contract}, GoogleProtobufAnyFilterer: GoogleProtobufAnyFilterer{contract: contract}}, nil
}

// GoogleProtobufAny is an auto generated Go binding around an Ethereum contract.
type GoogleProtobufAny struct {
	GoogleProtobufAnyCaller     // Read-only binding to the contract
	GoogleProtobufAnyTransactor // Write-only binding to the contract
	GoogleProtobufAnyFilterer   // Log filterer for contract events
}

// GoogleProtobufAnyCaller is an auto generated read-only Go binding around an Ethereum contract.
type GoogleProtobufAnyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoogleProtobufAnyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GoogleProtobufAnyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoogleProtobufAnyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GoogleProtobufAnyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoogleProtobufAnySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GoogleProtobufAnySession struct {
	Contract     *GoogleProtobufAny // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GoogleProtobufAnyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GoogleProtobufAnyCallerSession struct {
	Contract *GoogleProtobufAnyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// GoogleProtobufAnyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GoogleProtobufAnyTransactorSession struct {
	Contract     *GoogleProtobufAnyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// GoogleProtobufAnyRaw is an auto generated low-level Go binding around an Ethereum contract.
type GoogleProtobufAnyRaw struct {
	Contract *GoogleProtobufAny // Generic contract binding to access the raw methods on
}

// GoogleProtobufAnyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GoogleProtobufAnyCallerRaw struct {
	Contract *GoogleProtobufAnyCaller // Generic read-only contract binding to access the raw methods on
}

// GoogleProtobufAnyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GoogleProtobufAnyTransactorRaw struct {
	Contract *GoogleProtobufAnyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGoogleProtobufAny creates a new instance of GoogleProtobufAny, bound to a specific deployed contract.
func NewGoogleProtobufAny(address common.Address, backend bind.ContractBackend) (*GoogleProtobufAny, error) {
	contract, err := bindGoogleProtobufAny(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GoogleProtobufAny{GoogleProtobufAnyCaller: GoogleProtobufAnyCaller{contract: contract}, GoogleProtobufAnyTransactor: GoogleProtobufAnyTransactor{contract: contract}, GoogleProtobufAnyFilterer: GoogleProtobufAnyFilterer{contract: contract}}, nil
}

// NewGoogleProtobufAnyCaller creates a new read-only instance of GoogleProtobufAny, bound to a specific deployed contract.
func NewGoogleProtobufAnyCaller(address common.Address, caller bind.ContractCaller) (*GoogleProtobufAnyCaller, error) {
	contract, err := bindGoogleProtobufAny(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GoogleProtobufAnyCaller{contract: contract}, nil
}

// NewGoogleProtobufAnyTransactor creates a new write-only instance of GoogleProtobufAny, bound to a specific deployed contract.
func NewGoogleProtobufAnyTransactor(address common.Address, transactor bind.ContractTransactor) (*GoogleProtobufAnyTransactor, error) {
	contract, err := bindGoogleProtobufAny(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GoogleProtobufAnyTransactor{contract: contract}, nil
}

// NewGoogleProtobufAnyFilterer creates a new log filterer instance of GoogleProtobufAny, bound to a specific deployed contract.
func NewGoogleProtobufAnyFilterer(address common.Address, filterer bind.ContractFilterer) (*GoogleProtobufAnyFilterer, error) {
	contract, err := bindGoogleProtobufAny(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GoogleProtobufAnyFilterer{contract: contract}, nil
}

// bindGoogleProtobufAny binds a generic wrapper to an already deployed contract.
func bindGoogleProtobufAny(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GoogleProtobufAnyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GoogleProtobufAny *GoogleProtobufAnyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GoogleProtobufAny.Contract.GoogleProtobufAnyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GoogleProtobufAny *GoogleProtobufAnyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoogleProtobufAny.Contract.GoogleProtobufAnyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GoogleProtobufAny *GoogleProtobufAnyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GoogleProtobufAny.Contract.GoogleProtobufAnyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GoogleProtobufAny *GoogleProtobufAnyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GoogleProtobufAny.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GoogleProtobufAny *GoogleProtobufAnyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoogleProtobufAny.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GoogleProtobufAny *GoogleProtobufAnyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GoogleProtobufAny.Contract.contract.Transact(opts, method, params...)
}

// IApplicationMetaData contains all meta data concerning the IApplication contract.
var IApplicationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleAckPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleFailAckPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleSynPackage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"responsePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IApplicationABI is the input ABI used to generate the binding from.
// Deprecated: Use IApplicationMetaData.ABI instead.
var IApplicationABI = IApplicationMetaData.ABI

// IApplication is an auto generated Go binding around an Ethereum contract.
type IApplication struct {
	IApplicationCaller     // Read-only binding to the contract
	IApplicationTransactor // Write-only binding to the contract
	IApplicationFilterer   // Log filterer for contract events
}

// IApplicationCaller is an auto generated read-only Go binding around an Ethereum contract.
type IApplicationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IApplicationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IApplicationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IApplicationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IApplicationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IApplicationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IApplicationSession struct {
	Contract     *IApplication     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IApplicationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IApplicationCallerSession struct {
	Contract *IApplicationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IApplicationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IApplicationTransactorSession struct {
	Contract     *IApplicationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IApplicationRaw is an auto generated low-level Go binding around an Ethereum contract.
type IApplicationRaw struct {
	Contract *IApplication // Generic contract binding to access the raw methods on
}

// IApplicationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IApplicationCallerRaw struct {
	Contract *IApplicationCaller // Generic read-only contract binding to access the raw methods on
}

// IApplicationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IApplicationTransactorRaw struct {
	Contract *IApplicationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIApplication creates a new instance of IApplication, bound to a specific deployed contract.
func NewIApplication(address common.Address, backend bind.ContractBackend) (*IApplication, error) {
	contract, err := bindIApplication(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IApplication{IApplicationCaller: IApplicationCaller{contract: contract}, IApplicationTransactor: IApplicationTransactor{contract: contract}, IApplicationFilterer: IApplicationFilterer{contract: contract}}, nil
}

// NewIApplicationCaller creates a new read-only instance of IApplication, bound to a specific deployed contract.
func NewIApplicationCaller(address common.Address, caller bind.ContractCaller) (*IApplicationCaller, error) {
	contract, err := bindIApplication(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IApplicationCaller{contract: contract}, nil
}

// NewIApplicationTransactor creates a new write-only instance of IApplication, bound to a specific deployed contract.
func NewIApplicationTransactor(address common.Address, transactor bind.ContractTransactor) (*IApplicationTransactor, error) {
	contract, err := bindIApplication(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IApplicationTransactor{contract: contract}, nil
}

// NewIApplicationFilterer creates a new log filterer instance of IApplication, bound to a specific deployed contract.
func NewIApplicationFilterer(address common.Address, filterer bind.ContractFilterer) (*IApplicationFilterer, error) {
	contract, err := bindIApplication(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IApplicationFilterer{contract: contract}, nil
}

// bindIApplication binds a generic wrapper to an already deployed contract.
func bindIApplication(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IApplicationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IApplication *IApplicationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IApplication.Contract.IApplicationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IApplication *IApplicationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IApplication.Contract.IApplicationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IApplication *IApplicationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IApplication.Contract.IApplicationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IApplication *IApplicationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IApplication.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IApplication *IApplicationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IApplication.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IApplication *IApplicationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IApplication.Contract.contract.Transact(opts, method, params...)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_IApplication *IApplicationTransactor) HandleAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.contract.Transact(opts, "handleAckPackage", channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_IApplication *IApplicationSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.Contract.HandleAckPackage(&_IApplication.TransactOpts, channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_IApplication *IApplicationTransactorSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.Contract.HandleAckPackage(&_IApplication.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_IApplication *IApplicationTransactor) HandleFailAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.contract.Transact(opts, "handleFailAckPackage", channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_IApplication *IApplicationSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.Contract.HandleFailAckPackage(&_IApplication.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_IApplication *IApplicationTransactorSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.Contract.HandleFailAckPackage(&_IApplication.TransactOpts, channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 channelId, bytes msgBytes) returns(bytes responsePayload)
func (_IApplication *IApplicationTransactor) HandleSynPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.contract.Transact(opts, "handleSynPackage", channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 channelId, bytes msgBytes) returns(bytes responsePayload)
func (_IApplication *IApplicationSession) HandleSynPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.Contract.HandleSynPackage(&_IApplication.TransactOpts, channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 channelId, bytes msgBytes) returns(bytes responsePayload)
func (_IApplication *IApplicationTransactorSession) HandleSynPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.Contract.HandleSynPackage(&_IApplication.TransactOpts, channelId, msgBytes)
}

// IBSCValidatorSetMetaData contains all meta data concerning the IBSCValidatorSet contract.
var IBSCValidatorSetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"valAddress\",\"type\":\"address\"}],\"name\":\"isCurrentValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IBSCValidatorSetABI is the input ABI used to generate the binding from.
// Deprecated: Use IBSCValidatorSetMetaData.ABI instead.
var IBSCValidatorSetABI = IBSCValidatorSetMetaData.ABI

// IBSCValidatorSet is an auto generated Go binding around an Ethereum contract.
type IBSCValidatorSet struct {
	IBSCValidatorSetCaller     // Read-only binding to the contract
	IBSCValidatorSetTransactor // Write-only binding to the contract
	IBSCValidatorSetFilterer   // Log filterer for contract events
}

// IBSCValidatorSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBSCValidatorSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBSCValidatorSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBSCValidatorSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBSCValidatorSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBSCValidatorSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBSCValidatorSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBSCValidatorSetSession struct {
	Contract     *IBSCValidatorSet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBSCValidatorSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBSCValidatorSetCallerSession struct {
	Contract *IBSCValidatorSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IBSCValidatorSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBSCValidatorSetTransactorSession struct {
	Contract     *IBSCValidatorSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IBSCValidatorSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBSCValidatorSetRaw struct {
	Contract *IBSCValidatorSet // Generic contract binding to access the raw methods on
}

// IBSCValidatorSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBSCValidatorSetCallerRaw struct {
	Contract *IBSCValidatorSetCaller // Generic read-only contract binding to access the raw methods on
}

// IBSCValidatorSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBSCValidatorSetTransactorRaw struct {
	Contract *IBSCValidatorSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBSCValidatorSet creates a new instance of IBSCValidatorSet, bound to a specific deployed contract.
func NewIBSCValidatorSet(address common.Address, backend bind.ContractBackend) (*IBSCValidatorSet, error) {
	contract, err := bindIBSCValidatorSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBSCValidatorSet{IBSCValidatorSetCaller: IBSCValidatorSetCaller{contract: contract}, IBSCValidatorSetTransactor: IBSCValidatorSetTransactor{contract: contract}, IBSCValidatorSetFilterer: IBSCValidatorSetFilterer{contract: contract}}, nil
}

// NewIBSCValidatorSetCaller creates a new read-only instance of IBSCValidatorSet, bound to a specific deployed contract.
func NewIBSCValidatorSetCaller(address common.Address, caller bind.ContractCaller) (*IBSCValidatorSetCaller, error) {
	contract, err := bindIBSCValidatorSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBSCValidatorSetCaller{contract: contract}, nil
}

// NewIBSCValidatorSetTransactor creates a new write-only instance of IBSCValidatorSet, bound to a specific deployed contract.
func NewIBSCValidatorSetTransactor(address common.Address, transactor bind.ContractTransactor) (*IBSCValidatorSetTransactor, error) {
	contract, err := bindIBSCValidatorSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBSCValidatorSetTransactor{contract: contract}, nil
}

// NewIBSCValidatorSetFilterer creates a new log filterer instance of IBSCValidatorSet, bound to a specific deployed contract.
func NewIBSCValidatorSetFilterer(address common.Address, filterer bind.ContractFilterer) (*IBSCValidatorSetFilterer, error) {
	contract, err := bindIBSCValidatorSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBSCValidatorSetFilterer{contract: contract}, nil
}

// bindIBSCValidatorSet binds a generic wrapper to an already deployed contract.
func bindIBSCValidatorSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBSCValidatorSetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBSCValidatorSet *IBSCValidatorSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBSCValidatorSet.Contract.IBSCValidatorSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBSCValidatorSet *IBSCValidatorSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.IBSCValidatorSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBSCValidatorSet *IBSCValidatorSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.IBSCValidatorSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBSCValidatorSet *IBSCValidatorSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBSCValidatorSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBSCValidatorSet *IBSCValidatorSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBSCValidatorSet *IBSCValidatorSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.contract.Transact(opts, method, params...)
}

// IsCurrentValidator is a free data retrieval call binding the contract method 0x55614fcc.
//
// Solidity: function isCurrentValidator(address valAddress) view returns(bool)
func (_IBSCValidatorSet *IBSCValidatorSetCaller) IsCurrentValidator(opts *bind.CallOpts, valAddress common.Address) (bool, error) {
	var out []interface{}
	err := _IBSCValidatorSet.contract.Call(opts, &out, "isCurrentValidator", valAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurrentValidator is a free data retrieval call binding the contract method 0x55614fcc.
//
// Solidity: function isCurrentValidator(address valAddress) view returns(bool)
func (_IBSCValidatorSet *IBSCValidatorSetSession) IsCurrentValidator(valAddress common.Address) (bool, error) {
	return _IBSCValidatorSet.Contract.IsCurrentValidator(&_IBSCValidatorSet.CallOpts, valAddress)
}

// IsCurrentValidator is a free data retrieval call binding the contract method 0x55614fcc.
//
// Solidity: function isCurrentValidator(address valAddress) view returns(bool)
func (_IBSCValidatorSet *IBSCValidatorSetCallerSession) IsCurrentValidator(valAddress common.Address) (bool, error) {
	return _IBSCValidatorSet.Contract.IsCurrentValidator(&_IBSCValidatorSet.CallOpts, valAddress)
}

// ICrossChainMetaData contains all meta data concerning the ICrossChain contract.
var ICrossChainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"relayFee\",\"type\":\"uint256\"}],\"name\":\"sendSynPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ICrossChainABI is the input ABI used to generate the binding from.
// Deprecated: Use ICrossChainMetaData.ABI instead.
var ICrossChainABI = ICrossChainMetaData.ABI

// ICrossChain is an auto generated Go binding around an Ethereum contract.
type ICrossChain struct {
	ICrossChainCaller     // Read-only binding to the contract
	ICrossChainTransactor // Write-only binding to the contract
	ICrossChainFilterer   // Log filterer for contract events
}

// ICrossChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICrossChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrossChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICrossChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrossChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICrossChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrossChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICrossChainSession struct {
	Contract     *ICrossChain      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICrossChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICrossChainCallerSession struct {
	Contract *ICrossChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ICrossChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICrossChainTransactorSession struct {
	Contract     *ICrossChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ICrossChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICrossChainRaw struct {
	Contract *ICrossChain // Generic contract binding to access the raw methods on
}

// ICrossChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICrossChainCallerRaw struct {
	Contract *ICrossChainCaller // Generic read-only contract binding to access the raw methods on
}

// ICrossChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICrossChainTransactorRaw struct {
	Contract *ICrossChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICrossChain creates a new instance of ICrossChain, bound to a specific deployed contract.
func NewICrossChain(address common.Address, backend bind.ContractBackend) (*ICrossChain, error) {
	contract, err := bindICrossChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICrossChain{ICrossChainCaller: ICrossChainCaller{contract: contract}, ICrossChainTransactor: ICrossChainTransactor{contract: contract}, ICrossChainFilterer: ICrossChainFilterer{contract: contract}}, nil
}

// NewICrossChainCaller creates a new read-only instance of ICrossChain, bound to a specific deployed contract.
func NewICrossChainCaller(address common.Address, caller bind.ContractCaller) (*ICrossChainCaller, error) {
	contract, err := bindICrossChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICrossChainCaller{contract: contract}, nil
}

// NewICrossChainTransactor creates a new write-only instance of ICrossChain, bound to a specific deployed contract.
func NewICrossChainTransactor(address common.Address, transactor bind.ContractTransactor) (*ICrossChainTransactor, error) {
	contract, err := bindICrossChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICrossChainTransactor{contract: contract}, nil
}

// NewICrossChainFilterer creates a new log filterer instance of ICrossChain, bound to a specific deployed contract.
func NewICrossChainFilterer(address common.Address, filterer bind.ContractFilterer) (*ICrossChainFilterer, error) {
	contract, err := bindICrossChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICrossChainFilterer{contract: contract}, nil
}

// bindICrossChain binds a generic wrapper to an already deployed contract.
func bindICrossChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICrossChainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICrossChain *ICrossChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICrossChain.Contract.ICrossChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICrossChain *ICrossChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICrossChain.Contract.ICrossChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICrossChain *ICrossChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICrossChain.Contract.ICrossChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICrossChain *ICrossChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICrossChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICrossChain *ICrossChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICrossChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICrossChain *ICrossChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICrossChain.Contract.contract.Transact(opts, method, params...)
}

// SendSynPackage is a paid mutator transaction binding the contract method 0xf7a251d7.
//
// Solidity: function sendSynPackage(uint8 channelId, bytes msgBytes, uint256 relayFee) returns()
func (_ICrossChain *ICrossChainTransactor) SendSynPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte, relayFee *big.Int) (*types.Transaction, error) {
	return _ICrossChain.contract.Transact(opts, "sendSynPackage", channelId, msgBytes, relayFee)
}

// SendSynPackage is a paid mutator transaction binding the contract method 0xf7a251d7.
//
// Solidity: function sendSynPackage(uint8 channelId, bytes msgBytes, uint256 relayFee) returns()
func (_ICrossChain *ICrossChainSession) SendSynPackage(channelId uint8, msgBytes []byte, relayFee *big.Int) (*types.Transaction, error) {
	return _ICrossChain.Contract.SendSynPackage(&_ICrossChain.TransactOpts, channelId, msgBytes, relayFee)
}

// SendSynPackage is a paid mutator transaction binding the contract method 0xf7a251d7.
//
// Solidity: function sendSynPackage(uint8 channelId, bytes msgBytes, uint256 relayFee) returns()
func (_ICrossChain *ICrossChainTransactorSession) SendSynPackage(channelId uint8, msgBytes []byte, relayFee *big.Int) (*types.Transaction, error) {
	return _ICrossChain.Contract.SendSynPackage(&_ICrossChain.TransactOpts, channelId, msgBytes, relayFee)
}

// IEthereumLightClientMetaData contains all meta data concerning the IEthereumLightClient contract.
var IEthereumLightClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"finalizedExecutionStateRootAndSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IEthereumLightClientABI is the input ABI used to generate the binding from.
// Deprecated: Use IEthereumLightClientMetaData.ABI instead.
var IEthereumLightClientABI = IEthereumLightClientMetaData.ABI

// IEthereumLightClient is an auto generated Go binding around an Ethereum contract.
type IEthereumLightClient struct {
	IEthereumLightClientCaller     // Read-only binding to the contract
	IEthereumLightClientTransactor // Write-only binding to the contract
	IEthereumLightClientFilterer   // Log filterer for contract events
}

// IEthereumLightClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthereumLightClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthereumLightClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthereumLightClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthereumLightClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthereumLightClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthereumLightClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthereumLightClientSession struct {
	Contract     *IEthereumLightClient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IEthereumLightClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthereumLightClientCallerSession struct {
	Contract *IEthereumLightClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IEthereumLightClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthereumLightClientTransactorSession struct {
	Contract     *IEthereumLightClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IEthereumLightClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthereumLightClientRaw struct {
	Contract *IEthereumLightClient // Generic contract binding to access the raw methods on
}

// IEthereumLightClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthereumLightClientCallerRaw struct {
	Contract *IEthereumLightClientCaller // Generic read-only contract binding to access the raw methods on
}

// IEthereumLightClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthereumLightClientTransactorRaw struct {
	Contract *IEthereumLightClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthereumLightClient creates a new instance of IEthereumLightClient, bound to a specific deployed contract.
func NewIEthereumLightClient(address common.Address, backend bind.ContractBackend) (*IEthereumLightClient, error) {
	contract, err := bindIEthereumLightClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthereumLightClient{IEthereumLightClientCaller: IEthereumLightClientCaller{contract: contract}, IEthereumLightClientTransactor: IEthereumLightClientTransactor{contract: contract}, IEthereumLightClientFilterer: IEthereumLightClientFilterer{contract: contract}}, nil
}

// NewIEthereumLightClientCaller creates a new read-only instance of IEthereumLightClient, bound to a specific deployed contract.
func NewIEthereumLightClientCaller(address common.Address, caller bind.ContractCaller) (*IEthereumLightClientCaller, error) {
	contract, err := bindIEthereumLightClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthereumLightClientCaller{contract: contract}, nil
}

// NewIEthereumLightClientTransactor creates a new write-only instance of IEthereumLightClient, bound to a specific deployed contract.
func NewIEthereumLightClientTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthereumLightClientTransactor, error) {
	contract, err := bindIEthereumLightClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthereumLightClientTransactor{contract: contract}, nil
}

// NewIEthereumLightClientFilterer creates a new log filterer instance of IEthereumLightClient, bound to a specific deployed contract.
func NewIEthereumLightClientFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthereumLightClientFilterer, error) {
	contract, err := bindIEthereumLightClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthereumLightClientFilterer{contract: contract}, nil
}

// bindIEthereumLightClient binds a generic wrapper to an already deployed contract.
func bindIEthereumLightClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEthereumLightClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthereumLightClient *IEthereumLightClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthereumLightClient.Contract.IEthereumLightClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthereumLightClient *IEthereumLightClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthereumLightClient.Contract.IEthereumLightClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthereumLightClient *IEthereumLightClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthereumLightClient.Contract.IEthereumLightClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthereumLightClient *IEthereumLightClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthereumLightClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthereumLightClient *IEthereumLightClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthereumLightClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthereumLightClient *IEthereumLightClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthereumLightClient.Contract.contract.Transact(opts, method, params...)
}

// FinalizedExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x12420766.
//
// Solidity: function finalizedExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_IEthereumLightClient *IEthereumLightClientCaller) FinalizedExecutionStateRootAndSlot(opts *bind.CallOpts) (struct {
	Root [32]byte
	Slot uint64
}, error) {
	var out []interface{}
	err := _IEthereumLightClient.contract.Call(opts, &out, "finalizedExecutionStateRootAndSlot")

	outstruct := new(struct {
		Root [32]byte
		Slot uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Slot = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// FinalizedExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x12420766.
//
// Solidity: function finalizedExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_IEthereumLightClient *IEthereumLightClientSession) FinalizedExecutionStateRootAndSlot() (struct {
	Root [32]byte
	Slot uint64
}, error) {
	return _IEthereumLightClient.Contract.FinalizedExecutionStateRootAndSlot(&_IEthereumLightClient.CallOpts)
}

// FinalizedExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x12420766.
//
// Solidity: function finalizedExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_IEthereumLightClient *IEthereumLightClientCallerSession) FinalizedExecutionStateRootAndSlot() (struct {
	Root [32]byte
	Slot uint64
}, error) {
	return _IEthereumLightClient.Contract.FinalizedExecutionStateRootAndSlot(&_IEthereumLightClient.CallOpts)
}

// ITendermintLightClientMetaData contains all meta data concerning the ITendermintLightClient contract.
var ITendermintLightClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getAppHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"isHeaderSynced\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ITendermintLightClientABI is the input ABI used to generate the binding from.
// Deprecated: Use ITendermintLightClientMetaData.ABI instead.
var ITendermintLightClientABI = ITendermintLightClientMetaData.ABI

// ITendermintLightClient is an auto generated Go binding around an Ethereum contract.
type ITendermintLightClient struct {
	ITendermintLightClientCaller     // Read-only binding to the contract
	ITendermintLightClientTransactor // Write-only binding to the contract
	ITendermintLightClientFilterer   // Log filterer for contract events
}

// ITendermintLightClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITendermintLightClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITendermintLightClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITendermintLightClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITendermintLightClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITendermintLightClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITendermintLightClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITendermintLightClientSession struct {
	Contract     *ITendermintLightClient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ITendermintLightClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITendermintLightClientCallerSession struct {
	Contract *ITendermintLightClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ITendermintLightClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITendermintLightClientTransactorSession struct {
	Contract     *ITendermintLightClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ITendermintLightClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITendermintLightClientRaw struct {
	Contract *ITendermintLightClient // Generic contract binding to access the raw methods on
}

// ITendermintLightClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITendermintLightClientCallerRaw struct {
	Contract *ITendermintLightClientCaller // Generic read-only contract binding to access the raw methods on
}

// ITendermintLightClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITendermintLightClientTransactorRaw struct {
	Contract *ITendermintLightClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITendermintLightClient creates a new instance of ITendermintLightClient, bound to a specific deployed contract.
func NewITendermintLightClient(address common.Address, backend bind.ContractBackend) (*ITendermintLightClient, error) {
	contract, err := bindITendermintLightClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITendermintLightClient{ITendermintLightClientCaller: ITendermintLightClientCaller{contract: contract}, ITendermintLightClientTransactor: ITendermintLightClientTransactor{contract: contract}, ITendermintLightClientFilterer: ITendermintLightClientFilterer{contract: contract}}, nil
}

// NewITendermintLightClientCaller creates a new read-only instance of ITendermintLightClient, bound to a specific deployed contract.
func NewITendermintLightClientCaller(address common.Address, caller bind.ContractCaller) (*ITendermintLightClientCaller, error) {
	contract, err := bindITendermintLightClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITendermintLightClientCaller{contract: contract}, nil
}

// NewITendermintLightClientTransactor creates a new write-only instance of ITendermintLightClient, bound to a specific deployed contract.
func NewITendermintLightClientTransactor(address common.Address, transactor bind.ContractTransactor) (*ITendermintLightClientTransactor, error) {
	contract, err := bindITendermintLightClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITendermintLightClientTransactor{contract: contract}, nil
}

// NewITendermintLightClientFilterer creates a new log filterer instance of ITendermintLightClient, bound to a specific deployed contract.
func NewITendermintLightClientFilterer(address common.Address, filterer bind.ContractFilterer) (*ITendermintLightClientFilterer, error) {
	contract, err := bindITendermintLightClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITendermintLightClientFilterer{contract: contract}, nil
}

// bindITendermintLightClient binds a generic wrapper to an already deployed contract.
func bindITendermintLightClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ITendermintLightClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITendermintLightClient *ITendermintLightClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITendermintLightClient.Contract.ITendermintLightClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITendermintLightClient *ITendermintLightClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITendermintLightClient.Contract.ITendermintLightClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITendermintLightClient *ITendermintLightClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITendermintLightClient.Contract.ITendermintLightClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITendermintLightClient *ITendermintLightClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITendermintLightClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITendermintLightClient *ITendermintLightClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITendermintLightClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITendermintLightClient *ITendermintLightClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITendermintLightClient.Contract.contract.Transact(opts, method, params...)
}

// GetAppHash is a free data retrieval call binding the contract method 0xcba510a9.
//
// Solidity: function getAppHash(uint64 height) view returns(bytes32)
func (_ITendermintLightClient *ITendermintLightClientCaller) GetAppHash(opts *bind.CallOpts, height uint64) ([32]byte, error) {
	var out []interface{}
	err := _ITendermintLightClient.contract.Call(opts, &out, "getAppHash", height)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAppHash is a free data retrieval call binding the contract method 0xcba510a9.
//
// Solidity: function getAppHash(uint64 height) view returns(bytes32)
func (_ITendermintLightClient *ITendermintLightClientSession) GetAppHash(height uint64) ([32]byte, error) {
	return _ITendermintLightClient.Contract.GetAppHash(&_ITendermintLightClient.CallOpts, height)
}

// GetAppHash is a free data retrieval call binding the contract method 0xcba510a9.
//
// Solidity: function getAppHash(uint64 height) view returns(bytes32)
func (_ITendermintLightClient *ITendermintLightClientCallerSession) GetAppHash(height uint64) ([32]byte, error) {
	return _ITendermintLightClient.Contract.GetAppHash(&_ITendermintLightClient.CallOpts, height)
}

// IsHeaderSynced is a free data retrieval call binding the contract method 0xdf5fe704.
//
// Solidity: function isHeaderSynced(uint64 height) view returns(bool)
func (_ITendermintLightClient *ITendermintLightClientCaller) IsHeaderSynced(opts *bind.CallOpts, height uint64) (bool, error) {
	var out []interface{}
	err := _ITendermintLightClient.contract.Call(opts, &out, "isHeaderSynced", height)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHeaderSynced is a free data retrieval call binding the contract method 0xdf5fe704.
//
// Solidity: function isHeaderSynced(uint64 height) view returns(bool)
func (_ITendermintLightClient *ITendermintLightClientSession) IsHeaderSynced(height uint64) (bool, error) {
	return _ITendermintLightClient.Contract.IsHeaderSynced(&_ITendermintLightClient.CallOpts, height)
}

// IsHeaderSynced is a free data retrieval call binding the contract method 0xdf5fe704.
//
// Solidity: function isHeaderSynced(uint64 height) view returns(bool)
func (_ITendermintLightClient *ITendermintLightClientCallerSession) IsHeaderSynced(height uint64) (bool, error) {
	return _ITendermintLightClient.Contract.IsHeaderSynced(&_ITendermintLightClient.CallOpts, height)
}

// Ics23MetaData contains all meta data concerning the Ics23 contract.
var Ics23MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220fa2cf8f51a238c376a968c3fa492206ebd0950d7f6b3db786e71e2cf2858b2a664736f6c63430008120033",
}

// Ics23ABI is the input ABI used to generate the binding from.
// Deprecated: Use Ics23MetaData.ABI instead.
var Ics23ABI = Ics23MetaData.ABI

// Ics23Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Ics23MetaData.Bin instead.
var Ics23Bin = Ics23MetaData.Bin

// DeployIcs23 deploys a new Ethereum contract, binding an instance of Ics23 to it.
func DeployIcs23(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ics23, error) {
	parsed, err := Ics23MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Ics23Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ics23{Ics23Caller: Ics23Caller{contract: contract}, Ics23Transactor: Ics23Transactor{contract: contract}, Ics23Filterer: Ics23Filterer{contract: contract}}, nil
}

// Ics23 is an auto generated Go binding around an Ethereum contract.
type Ics23 struct {
	Ics23Caller     // Read-only binding to the contract
	Ics23Transactor // Write-only binding to the contract
	Ics23Filterer   // Log filterer for contract events
}

// Ics23Caller is an auto generated read-only Go binding around an Ethereum contract.
type Ics23Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ics23Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Ics23Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ics23Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ics23Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ics23Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ics23Session struct {
	Contract     *Ics23            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ics23CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ics23CallerSession struct {
	Contract *Ics23Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Ics23TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ics23TransactorSession struct {
	Contract     *Ics23Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ics23Raw is an auto generated low-level Go binding around an Ethereum contract.
type Ics23Raw struct {
	Contract *Ics23 // Generic contract binding to access the raw methods on
}

// Ics23CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ics23CallerRaw struct {
	Contract *Ics23Caller // Generic read-only contract binding to access the raw methods on
}

// Ics23TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ics23TransactorRaw struct {
	Contract *Ics23Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIcs23 creates a new instance of Ics23, bound to a specific deployed contract.
func NewIcs23(address common.Address, backend bind.ContractBackend) (*Ics23, error) {
	contract, err := bindIcs23(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ics23{Ics23Caller: Ics23Caller{contract: contract}, Ics23Transactor: Ics23Transactor{contract: contract}, Ics23Filterer: Ics23Filterer{contract: contract}}, nil
}

// NewIcs23Caller creates a new read-only instance of Ics23, bound to a specific deployed contract.
func NewIcs23Caller(address common.Address, caller bind.ContractCaller) (*Ics23Caller, error) {
	contract, err := bindIcs23(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ics23Caller{contract: contract}, nil
}

// NewIcs23Transactor creates a new write-only instance of Ics23, bound to a specific deployed contract.
func NewIcs23Transactor(address common.Address, transactor bind.ContractTransactor) (*Ics23Transactor, error) {
	contract, err := bindIcs23(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ics23Transactor{contract: contract}, nil
}

// NewIcs23Filterer creates a new log filterer instance of Ics23, bound to a specific deployed contract.
func NewIcs23Filterer(address common.Address, filterer bind.ContractFilterer) (*Ics23Filterer, error) {
	contract, err := bindIcs23(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ics23Filterer{contract: contract}, nil
}

// bindIcs23 binds a generic wrapper to an already deployed contract.
func bindIcs23(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Ics23MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ics23 *Ics23Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ics23.Contract.Ics23Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ics23 *Ics23Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ics23.Contract.Ics23Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ics23 *Ics23Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ics23.Contract.Ics23Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ics23 *Ics23CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ics23.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ics23 *Ics23TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ics23.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ics23 *Ics23TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ics23.Contract.contract.Transact(opts, method, params...)
}

// InitializableMetaData contains all meta data concerning the Initializable contract.
var InitializableMetaData = &bind.MetaData{
	ABI: "[]",
}

// InitializableABI is the input ABI used to generate the binding from.
// Deprecated: Use InitializableMetaData.ABI instead.
var InitializableABI = InitializableMetaData.ABI

// Initializable is an auto generated Go binding around an Ethereum contract.
type Initializable struct {
	InitializableCaller     // Read-only binding to the contract
	InitializableTransactor // Write-only binding to the contract
	InitializableFilterer   // Log filterer for contract events
}

// InitializableCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializableSession struct {
	Contract     *Initializable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitializableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializableCallerSession struct {
	Contract *InitializableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InitializableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializableTransactorSession struct {
	Contract     *InitializableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InitializableRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializableRaw struct {
	Contract *Initializable // Generic contract binding to access the raw methods on
}

// InitializableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializableCallerRaw struct {
	Contract *InitializableCaller // Generic read-only contract binding to access the raw methods on
}

// InitializableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializableTransactorRaw struct {
	Contract *InitializableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializable creates a new instance of Initializable, bound to a specific deployed contract.
func NewInitializable(address common.Address, backend bind.ContractBackend) (*Initializable, error) {
	contract, err := bindInitializable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Initializable{InitializableCaller: InitializableCaller{contract: contract}, InitializableTransactor: InitializableTransactor{contract: contract}, InitializableFilterer: InitializableFilterer{contract: contract}}, nil
}

// NewInitializableCaller creates a new read-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableCaller(address common.Address, caller bind.ContractCaller) (*InitializableCaller, error) {
	contract, err := bindInitializable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableCaller{contract: contract}, nil
}

// NewInitializableTransactor creates a new write-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializableTransactor, error) {
	contract, err := bindInitializable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableTransactor{contract: contract}, nil
}

// NewInitializableFilterer creates a new log filterer instance of Initializable, bound to a specific deployed contract.
func NewInitializableFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializableFilterer, error) {
	contract, err := bindInitializable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializableFilterer{contract: contract}, nil
}

// bindInitializable binds a generic wrapper to an already deployed contract.
func bindInitializable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InitializableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.InitializableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transact(opts, method, params...)
}

// InnerOpMetaData contains all meta data concerning the InnerOp contract.
var InnerOpMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212201ddb66bfc756b2493b4962135c8c49855b3d65fcf52feb06f55b7d384065d9f764736f6c63430008120033",
}

// InnerOpABI is the input ABI used to generate the binding from.
// Deprecated: Use InnerOpMetaData.ABI instead.
var InnerOpABI = InnerOpMetaData.ABI

// InnerOpBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InnerOpMetaData.Bin instead.
var InnerOpBin = InnerOpMetaData.Bin

// DeployInnerOp deploys a new Ethereum contract, binding an instance of InnerOp to it.
func DeployInnerOp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InnerOp, error) {
	parsed, err := InnerOpMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InnerOpBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InnerOp{InnerOpCaller: InnerOpCaller{contract: contract}, InnerOpTransactor: InnerOpTransactor{contract: contract}, InnerOpFilterer: InnerOpFilterer{contract: contract}}, nil
}

// InnerOp is an auto generated Go binding around an Ethereum contract.
type InnerOp struct {
	InnerOpCaller     // Read-only binding to the contract
	InnerOpTransactor // Write-only binding to the contract
	InnerOpFilterer   // Log filterer for contract events
}

// InnerOpCaller is an auto generated read-only Go binding around an Ethereum contract.
type InnerOpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InnerOpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InnerOpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InnerOpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InnerOpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InnerOpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InnerOpSession struct {
	Contract     *InnerOp          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InnerOpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InnerOpCallerSession struct {
	Contract *InnerOpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// InnerOpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InnerOpTransactorSession struct {
	Contract     *InnerOpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// InnerOpRaw is an auto generated low-level Go binding around an Ethereum contract.
type InnerOpRaw struct {
	Contract *InnerOp // Generic contract binding to access the raw methods on
}

// InnerOpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InnerOpCallerRaw struct {
	Contract *InnerOpCaller // Generic read-only contract binding to access the raw methods on
}

// InnerOpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InnerOpTransactorRaw struct {
	Contract *InnerOpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInnerOp creates a new instance of InnerOp, bound to a specific deployed contract.
func NewInnerOp(address common.Address, backend bind.ContractBackend) (*InnerOp, error) {
	contract, err := bindInnerOp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InnerOp{InnerOpCaller: InnerOpCaller{contract: contract}, InnerOpTransactor: InnerOpTransactor{contract: contract}, InnerOpFilterer: InnerOpFilterer{contract: contract}}, nil
}

// NewInnerOpCaller creates a new read-only instance of InnerOp, bound to a specific deployed contract.
func NewInnerOpCaller(address common.Address, caller bind.ContractCaller) (*InnerOpCaller, error) {
	contract, err := bindInnerOp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InnerOpCaller{contract: contract}, nil
}

// NewInnerOpTransactor creates a new write-only instance of InnerOp, bound to a specific deployed contract.
func NewInnerOpTransactor(address common.Address, transactor bind.ContractTransactor) (*InnerOpTransactor, error) {
	contract, err := bindInnerOp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InnerOpTransactor{contract: contract}, nil
}

// NewInnerOpFilterer creates a new log filterer instance of InnerOp, bound to a specific deployed contract.
func NewInnerOpFilterer(address common.Address, filterer bind.ContractFilterer) (*InnerOpFilterer, error) {
	contract, err := bindInnerOp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InnerOpFilterer{contract: contract}, nil
}

// bindInnerOp binds a generic wrapper to an already deployed contract.
func bindInnerOp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InnerOpMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InnerOp *InnerOpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InnerOp.Contract.InnerOpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InnerOp *InnerOpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InnerOp.Contract.InnerOpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InnerOp *InnerOpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InnerOp.Contract.InnerOpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InnerOp *InnerOpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InnerOp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InnerOp *InnerOpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InnerOp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InnerOp *InnerOpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InnerOp.Contract.contract.Transact(opts, method, params...)
}

// InnerSpecMetaData contains all meta data concerning the InnerSpec contract.
var InnerSpecMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea264697066735822122038bbb65dc0949dedd39d0fe28dae1d0ecdf0d19574ef5d414c71120e2b55907764736f6c63430008120033",
}

// InnerSpecABI is the input ABI used to generate the binding from.
// Deprecated: Use InnerSpecMetaData.ABI instead.
var InnerSpecABI = InnerSpecMetaData.ABI

// InnerSpecBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InnerSpecMetaData.Bin instead.
var InnerSpecBin = InnerSpecMetaData.Bin

// DeployInnerSpec deploys a new Ethereum contract, binding an instance of InnerSpec to it.
func DeployInnerSpec(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InnerSpec, error) {
	parsed, err := InnerSpecMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InnerSpecBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InnerSpec{InnerSpecCaller: InnerSpecCaller{contract: contract}, InnerSpecTransactor: InnerSpecTransactor{contract: contract}, InnerSpecFilterer: InnerSpecFilterer{contract: contract}}, nil
}

// InnerSpec is an auto generated Go binding around an Ethereum contract.
type InnerSpec struct {
	InnerSpecCaller     // Read-only binding to the contract
	InnerSpecTransactor // Write-only binding to the contract
	InnerSpecFilterer   // Log filterer for contract events
}

// InnerSpecCaller is an auto generated read-only Go binding around an Ethereum contract.
type InnerSpecCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InnerSpecTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InnerSpecTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InnerSpecFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InnerSpecFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InnerSpecSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InnerSpecSession struct {
	Contract     *InnerSpec        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InnerSpecCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InnerSpecCallerSession struct {
	Contract *InnerSpecCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// InnerSpecTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InnerSpecTransactorSession struct {
	Contract     *InnerSpecTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// InnerSpecRaw is an auto generated low-level Go binding around an Ethereum contract.
type InnerSpecRaw struct {
	Contract *InnerSpec // Generic contract binding to access the raw methods on
}

// InnerSpecCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InnerSpecCallerRaw struct {
	Contract *InnerSpecCaller // Generic read-only contract binding to access the raw methods on
}

// InnerSpecTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InnerSpecTransactorRaw struct {
	Contract *InnerSpecTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInnerSpec creates a new instance of InnerSpec, bound to a specific deployed contract.
func NewInnerSpec(address common.Address, backend bind.ContractBackend) (*InnerSpec, error) {
	contract, err := bindInnerSpec(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InnerSpec{InnerSpecCaller: InnerSpecCaller{contract: contract}, InnerSpecTransactor: InnerSpecTransactor{contract: contract}, InnerSpecFilterer: InnerSpecFilterer{contract: contract}}, nil
}

// NewInnerSpecCaller creates a new read-only instance of InnerSpec, bound to a specific deployed contract.
func NewInnerSpecCaller(address common.Address, caller bind.ContractCaller) (*InnerSpecCaller, error) {
	contract, err := bindInnerSpec(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InnerSpecCaller{contract: contract}, nil
}

// NewInnerSpecTransactor creates a new write-only instance of InnerSpec, bound to a specific deployed contract.
func NewInnerSpecTransactor(address common.Address, transactor bind.ContractTransactor) (*InnerSpecTransactor, error) {
	contract, err := bindInnerSpec(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InnerSpecTransactor{contract: contract}, nil
}

// NewInnerSpecFilterer creates a new log filterer instance of InnerSpec, bound to a specific deployed contract.
func NewInnerSpecFilterer(address common.Address, filterer bind.ContractFilterer) (*InnerSpecFilterer, error) {
	contract, err := bindInnerSpec(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InnerSpecFilterer{contract: contract}, nil
}

// bindInnerSpec binds a generic wrapper to an already deployed contract.
func bindInnerSpec(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InnerSpecMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InnerSpec *InnerSpecRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InnerSpec.Contract.InnerSpecCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InnerSpec *InnerSpecRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InnerSpec.Contract.InnerSpecTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InnerSpec *InnerSpecRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InnerSpec.Contract.InnerSpecTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InnerSpec *InnerSpecCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InnerSpec.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InnerSpec *InnerSpecTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InnerSpec.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InnerSpec *InnerSpecTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InnerSpec.Contract.contract.Transact(opts, method, params...)
}

// LeafOpMetaData contains all meta data concerning the LeafOp contract.
var LeafOpMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220f9bb0f3710f8604c57497cbd004618859f7556e9435454e716b423dee98c9c3764736f6c63430008120033",
}

// LeafOpABI is the input ABI used to generate the binding from.
// Deprecated: Use LeafOpMetaData.ABI instead.
var LeafOpABI = LeafOpMetaData.ABI

// LeafOpBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LeafOpMetaData.Bin instead.
var LeafOpBin = LeafOpMetaData.Bin

// DeployLeafOp deploys a new Ethereum contract, binding an instance of LeafOp to it.
func DeployLeafOp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LeafOp, error) {
	parsed, err := LeafOpMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LeafOpBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LeafOp{LeafOpCaller: LeafOpCaller{contract: contract}, LeafOpTransactor: LeafOpTransactor{contract: contract}, LeafOpFilterer: LeafOpFilterer{contract: contract}}, nil
}

// LeafOp is an auto generated Go binding around an Ethereum contract.
type LeafOp struct {
	LeafOpCaller     // Read-only binding to the contract
	LeafOpTransactor // Write-only binding to the contract
	LeafOpFilterer   // Log filterer for contract events
}

// LeafOpCaller is an auto generated read-only Go binding around an Ethereum contract.
type LeafOpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeafOpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LeafOpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeafOpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LeafOpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeafOpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LeafOpSession struct {
	Contract     *LeafOp           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LeafOpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LeafOpCallerSession struct {
	Contract *LeafOpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LeafOpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LeafOpTransactorSession struct {
	Contract     *LeafOpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LeafOpRaw is an auto generated low-level Go binding around an Ethereum contract.
type LeafOpRaw struct {
	Contract *LeafOp // Generic contract binding to access the raw methods on
}

// LeafOpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LeafOpCallerRaw struct {
	Contract *LeafOpCaller // Generic read-only contract binding to access the raw methods on
}

// LeafOpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LeafOpTransactorRaw struct {
	Contract *LeafOpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLeafOp creates a new instance of LeafOp, bound to a specific deployed contract.
func NewLeafOp(address common.Address, backend bind.ContractBackend) (*LeafOp, error) {
	contract, err := bindLeafOp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LeafOp{LeafOpCaller: LeafOpCaller{contract: contract}, LeafOpTransactor: LeafOpTransactor{contract: contract}, LeafOpFilterer: LeafOpFilterer{contract: contract}}, nil
}

// NewLeafOpCaller creates a new read-only instance of LeafOp, bound to a specific deployed contract.
func NewLeafOpCaller(address common.Address, caller bind.ContractCaller) (*LeafOpCaller, error) {
	contract, err := bindLeafOp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LeafOpCaller{contract: contract}, nil
}

// NewLeafOpTransactor creates a new write-only instance of LeafOp, bound to a specific deployed contract.
func NewLeafOpTransactor(address common.Address, transactor bind.ContractTransactor) (*LeafOpTransactor, error) {
	contract, err := bindLeafOp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LeafOpTransactor{contract: contract}, nil
}

// NewLeafOpFilterer creates a new log filterer instance of LeafOp, bound to a specific deployed contract.
func NewLeafOpFilterer(address common.Address, filterer bind.ContractFilterer) (*LeafOpFilterer, error) {
	contract, err := bindLeafOp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LeafOpFilterer{contract: contract}, nil
}

// bindLeafOp binds a generic wrapper to an already deployed contract.
func bindLeafOp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LeafOpMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LeafOp *LeafOpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LeafOp.Contract.LeafOpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LeafOp *LeafOpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LeafOp.Contract.LeafOpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LeafOp *LeafOpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LeafOp.Contract.LeafOpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LeafOp *LeafOpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LeafOp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LeafOp *LeafOpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LeafOp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LeafOp *LeafOpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LeafOp.Contract.contract.Transact(opts, method, params...)
}

// LightHeaderMetaData contains all meta data concerning the LightHeader contract.
var LightHeaderMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220eb3141a3b7db17dd0785481b2916a667f678d57b5b9024d5a6b8f214b86aa49b64736f6c63430008120033",
}

// LightHeaderABI is the input ABI used to generate the binding from.
// Deprecated: Use LightHeaderMetaData.ABI instead.
var LightHeaderABI = LightHeaderMetaData.ABI

// LightHeaderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LightHeaderMetaData.Bin instead.
var LightHeaderBin = LightHeaderMetaData.Bin

// DeployLightHeader deploys a new Ethereum contract, binding an instance of LightHeader to it.
func DeployLightHeader(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LightHeader, error) {
	parsed, err := LightHeaderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LightHeaderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LightHeader{LightHeaderCaller: LightHeaderCaller{contract: contract}, LightHeaderTransactor: LightHeaderTransactor{contract: contract}, LightHeaderFilterer: LightHeaderFilterer{contract: contract}}, nil
}

// LightHeader is an auto generated Go binding around an Ethereum contract.
type LightHeader struct {
	LightHeaderCaller     // Read-only binding to the contract
	LightHeaderTransactor // Write-only binding to the contract
	LightHeaderFilterer   // Log filterer for contract events
}

// LightHeaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type LightHeaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightHeaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LightHeaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightHeaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LightHeaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightHeaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LightHeaderSession struct {
	Contract     *LightHeader      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LightHeaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LightHeaderCallerSession struct {
	Contract *LightHeaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// LightHeaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LightHeaderTransactorSession struct {
	Contract     *LightHeaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LightHeaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type LightHeaderRaw struct {
	Contract *LightHeader // Generic contract binding to access the raw methods on
}

// LightHeaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LightHeaderCallerRaw struct {
	Contract *LightHeaderCaller // Generic read-only contract binding to access the raw methods on
}

// LightHeaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LightHeaderTransactorRaw struct {
	Contract *LightHeaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLightHeader creates a new instance of LightHeader, bound to a specific deployed contract.
func NewLightHeader(address common.Address, backend bind.ContractBackend) (*LightHeader, error) {
	contract, err := bindLightHeader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LightHeader{LightHeaderCaller: LightHeaderCaller{contract: contract}, LightHeaderTransactor: LightHeaderTransactor{contract: contract}, LightHeaderFilterer: LightHeaderFilterer{contract: contract}}, nil
}

// NewLightHeaderCaller creates a new read-only instance of LightHeader, bound to a specific deployed contract.
func NewLightHeaderCaller(address common.Address, caller bind.ContractCaller) (*LightHeaderCaller, error) {
	contract, err := bindLightHeader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LightHeaderCaller{contract: contract}, nil
}

// NewLightHeaderTransactor creates a new write-only instance of LightHeader, bound to a specific deployed contract.
func NewLightHeaderTransactor(address common.Address, transactor bind.ContractTransactor) (*LightHeaderTransactor, error) {
	contract, err := bindLightHeader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LightHeaderTransactor{contract: contract}, nil
}

// NewLightHeaderFilterer creates a new log filterer instance of LightHeader, bound to a specific deployed contract.
func NewLightHeaderFilterer(address common.Address, filterer bind.ContractFilterer) (*LightHeaderFilterer, error) {
	contract, err := bindLightHeader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LightHeaderFilterer{contract: contract}, nil
}

// bindLightHeader binds a generic wrapper to an already deployed contract.
func bindLightHeader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LightHeaderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightHeader *LightHeaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightHeader.Contract.LightHeaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightHeader *LightHeaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightHeader.Contract.LightHeaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightHeader *LightHeaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightHeader.Contract.LightHeaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightHeader *LightHeaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightHeader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightHeader *LightHeaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightHeader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightHeader *LightHeaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightHeader.Contract.contract.Transact(opts, method, params...)
}

// MathMetaData contains all meta data concerning the Math contract.
var MathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212204f67e2502e5333df28d66176a7ac174394ca55e6c1c8cd2da8b448232da5210064736f6c63430008120033",
}

// MathABI is the input ABI used to generate the binding from.
// Deprecated: Use MathMetaData.ABI instead.
var MathABI = MathMetaData.ABI

// MathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MathMetaData.Bin instead.
var MathBin = MathMetaData.Bin

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := MathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// MemoryMetaData contains all meta data concerning the Memory contract.
var MemoryMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea264697066735822122092e6ea845afe82f74b983c0e2528de7713da0bfca6360b5b1537ef8670cd959a64736f6c63430008120033",
}

// MemoryABI is the input ABI used to generate the binding from.
// Deprecated: Use MemoryMetaData.ABI instead.
var MemoryABI = MemoryMetaData.ABI

// MemoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MemoryMetaData.Bin instead.
var MemoryBin = MemoryMetaData.Bin

// DeployMemory deploys a new Ethereum contract, binding an instance of Memory to it.
func DeployMemory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Memory, error) {
	parsed, err := MemoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MemoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Memory{MemoryCaller: MemoryCaller{contract: contract}, MemoryTransactor: MemoryTransactor{contract: contract}, MemoryFilterer: MemoryFilterer{contract: contract}}, nil
}

// Memory is an auto generated Go binding around an Ethereum contract.
type Memory struct {
	MemoryCaller     // Read-only binding to the contract
	MemoryTransactor // Write-only binding to the contract
	MemoryFilterer   // Log filterer for contract events
}

// MemoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MemoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MemoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MemoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MemorySession struct {
	Contract     *Memory           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MemoryCallerSession struct {
	Contract *MemoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MemoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MemoryTransactorSession struct {
	Contract     *MemoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MemoryRaw struct {
	Contract *Memory // Generic contract binding to access the raw methods on
}

// MemoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MemoryCallerRaw struct {
	Contract *MemoryCaller // Generic read-only contract binding to access the raw methods on
}

// MemoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MemoryTransactorRaw struct {
	Contract *MemoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMemory creates a new instance of Memory, bound to a specific deployed contract.
func NewMemory(address common.Address, backend bind.ContractBackend) (*Memory, error) {
	contract, err := bindMemory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Memory{MemoryCaller: MemoryCaller{contract: contract}, MemoryTransactor: MemoryTransactor{contract: contract}, MemoryFilterer: MemoryFilterer{contract: contract}}, nil
}

// NewMemoryCaller creates a new read-only instance of Memory, bound to a specific deployed contract.
func NewMemoryCaller(address common.Address, caller bind.ContractCaller) (*MemoryCaller, error) {
	contract, err := bindMemory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MemoryCaller{contract: contract}, nil
}

// NewMemoryTransactor creates a new write-only instance of Memory, bound to a specific deployed contract.
func NewMemoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MemoryTransactor, error) {
	contract, err := bindMemory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MemoryTransactor{contract: contract}, nil
}

// NewMemoryFilterer creates a new log filterer instance of Memory, bound to a specific deployed contract.
func NewMemoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MemoryFilterer, error) {
	contract, err := bindMemory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MemoryFilterer{contract: contract}, nil
}

// bindMemory binds a generic wrapper to an already deployed contract.
func bindMemory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MemoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Memory *MemoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Memory.Contract.MemoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Memory *MemoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Memory.Contract.MemoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Memory *MemoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Memory.Contract.MemoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Memory *MemoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Memory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Memory *MemoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Memory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Memory *MemoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Memory.Contract.contract.Transact(opts, method, params...)
}

// MerkleRootMetaData contains all meta data concerning the MerkleRoot contract.
var MerkleRootMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220c70b057d35dd4f4564d94fb2ed873044a26f6127393af846c169286342b6ba0464736f6c63430008120033",
}

// MerkleRootABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleRootMetaData.ABI instead.
var MerkleRootABI = MerkleRootMetaData.ABI

// MerkleRootBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleRootMetaData.Bin instead.
var MerkleRootBin = MerkleRootMetaData.Bin

// DeployMerkleRoot deploys a new Ethereum contract, binding an instance of MerkleRoot to it.
func DeployMerkleRoot(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleRoot, error) {
	parsed, err := MerkleRootMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleRootBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleRoot{MerkleRootCaller: MerkleRootCaller{contract: contract}, MerkleRootTransactor: MerkleRootTransactor{contract: contract}, MerkleRootFilterer: MerkleRootFilterer{contract: contract}}, nil
}

// MerkleRoot is an auto generated Go binding around an Ethereum contract.
type MerkleRoot struct {
	MerkleRootCaller     // Read-only binding to the contract
	MerkleRootTransactor // Write-only binding to the contract
	MerkleRootFilterer   // Log filterer for contract events
}

// MerkleRootCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleRootCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleRootTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleRootTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleRootFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleRootFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleRootSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleRootSession struct {
	Contract     *MerkleRoot       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleRootCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleRootCallerSession struct {
	Contract *MerkleRootCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MerkleRootTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleRootTransactorSession struct {
	Contract     *MerkleRootTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MerkleRootRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleRootRaw struct {
	Contract *MerkleRoot // Generic contract binding to access the raw methods on
}

// MerkleRootCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleRootCallerRaw struct {
	Contract *MerkleRootCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleRootTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleRootTransactorRaw struct {
	Contract *MerkleRootTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleRoot creates a new instance of MerkleRoot, bound to a specific deployed contract.
func NewMerkleRoot(address common.Address, backend bind.ContractBackend) (*MerkleRoot, error) {
	contract, err := bindMerkleRoot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleRoot{MerkleRootCaller: MerkleRootCaller{contract: contract}, MerkleRootTransactor: MerkleRootTransactor{contract: contract}, MerkleRootFilterer: MerkleRootFilterer{contract: contract}}, nil
}

// NewMerkleRootCaller creates a new read-only instance of MerkleRoot, bound to a specific deployed contract.
func NewMerkleRootCaller(address common.Address, caller bind.ContractCaller) (*MerkleRootCaller, error) {
	contract, err := bindMerkleRoot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleRootCaller{contract: contract}, nil
}

// NewMerkleRootTransactor creates a new write-only instance of MerkleRoot, bound to a specific deployed contract.
func NewMerkleRootTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleRootTransactor, error) {
	contract, err := bindMerkleRoot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleRootTransactor{contract: contract}, nil
}

// NewMerkleRootFilterer creates a new log filterer instance of MerkleRoot, bound to a specific deployed contract.
func NewMerkleRootFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleRootFilterer, error) {
	contract, err := bindMerkleRoot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleRootFilterer{contract: contract}, nil
}

// bindMerkleRoot binds a generic wrapper to an already deployed contract.
func bindMerkleRoot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MerkleRootMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleRoot *MerkleRootRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleRoot.Contract.MerkleRootCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleRoot *MerkleRootRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleRoot.Contract.MerkleRootTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleRoot *MerkleRootRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleRoot.Contract.MerkleRootTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleRoot *MerkleRootCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleRoot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleRoot *MerkleRootTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleRoot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleRoot *MerkleRootTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleRoot.Contract.contract.Transact(opts, method, params...)
}

// MerkleTreeMetaData contains all meta data concerning the MerkleTree contract.
var MerkleTreeMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea264697066735822122055927365b40d6f9790466d4606a4d2a1c0e90e1233251ddaefcf2b609622213064736f6c63430008120033",
}

// MerkleTreeABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleTreeMetaData.ABI instead.
var MerkleTreeABI = MerkleTreeMetaData.ABI

// MerkleTreeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleTreeMetaData.Bin instead.
var MerkleTreeBin = MerkleTreeMetaData.Bin

// DeployMerkleTree deploys a new Ethereum contract, binding an instance of MerkleTree to it.
func DeployMerkleTree(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleTree, error) {
	parsed, err := MerkleTreeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleTreeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleTree{MerkleTreeCaller: MerkleTreeCaller{contract: contract}, MerkleTreeTransactor: MerkleTreeTransactor{contract: contract}, MerkleTreeFilterer: MerkleTreeFilterer{contract: contract}}, nil
}

// MerkleTree is an auto generated Go binding around an Ethereum contract.
type MerkleTree struct {
	MerkleTreeCaller     // Read-only binding to the contract
	MerkleTreeTransactor // Write-only binding to the contract
	MerkleTreeFilterer   // Log filterer for contract events
}

// MerkleTreeCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleTreeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleTreeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleTreeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleTreeSession struct {
	Contract     *MerkleTree       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleTreeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleTreeCallerSession struct {
	Contract *MerkleTreeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MerkleTreeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleTreeTransactorSession struct {
	Contract     *MerkleTreeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MerkleTreeRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleTreeRaw struct {
	Contract *MerkleTree // Generic contract binding to access the raw methods on
}

// MerkleTreeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleTreeCallerRaw struct {
	Contract *MerkleTreeCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleTreeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleTreeTransactorRaw struct {
	Contract *MerkleTreeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleTree creates a new instance of MerkleTree, bound to a specific deployed contract.
func NewMerkleTree(address common.Address, backend bind.ContractBackend) (*MerkleTree, error) {
	contract, err := bindMerkleTree(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleTree{MerkleTreeCaller: MerkleTreeCaller{contract: contract}, MerkleTreeTransactor: MerkleTreeTransactor{contract: contract}, MerkleTreeFilterer: MerkleTreeFilterer{contract: contract}}, nil
}

// NewMerkleTreeCaller creates a new read-only instance of MerkleTree, bound to a specific deployed contract.
func NewMerkleTreeCaller(address common.Address, caller bind.ContractCaller) (*MerkleTreeCaller, error) {
	contract, err := bindMerkleTree(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeCaller{contract: contract}, nil
}

// NewMerkleTreeTransactor creates a new write-only instance of MerkleTree, bound to a specific deployed contract.
func NewMerkleTreeTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleTreeTransactor, error) {
	contract, err := bindMerkleTree(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeTransactor{contract: contract}, nil
}

// NewMerkleTreeFilterer creates a new log filterer instance of MerkleTree, bound to a specific deployed contract.
func NewMerkleTreeFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleTreeFilterer, error) {
	contract, err := bindMerkleTree(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeFilterer{contract: contract}, nil
}

// bindMerkleTree binds a generic wrapper to an already deployed contract.
func bindMerkleTree(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MerkleTreeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTree *MerkleTreeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleTree.Contract.MerkleTreeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTree *MerkleTreeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTree.Contract.MerkleTreeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTree *MerkleTreeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTree.Contract.MerkleTreeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTree *MerkleTreeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleTree.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTree *MerkleTreeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTree.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTree *MerkleTreeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTree.Contract.contract.Transact(opts, method, params...)
}

// NonExistenceProofMetaData contains all meta data concerning the NonExistenceProof contract.
var NonExistenceProofMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220f07d63b4d63717c4e6ad7ad25f1a467733b0c81d3962f837f612a9e6b719dfb164736f6c63430008120033",
}

// NonExistenceProofABI is the input ABI used to generate the binding from.
// Deprecated: Use NonExistenceProofMetaData.ABI instead.
var NonExistenceProofABI = NonExistenceProofMetaData.ABI

// NonExistenceProofBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NonExistenceProofMetaData.Bin instead.
var NonExistenceProofBin = NonExistenceProofMetaData.Bin

// DeployNonExistenceProof deploys a new Ethereum contract, binding an instance of NonExistenceProof to it.
func DeployNonExistenceProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NonExistenceProof, error) {
	parsed, err := NonExistenceProofMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NonExistenceProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NonExistenceProof{NonExistenceProofCaller: NonExistenceProofCaller{contract: contract}, NonExistenceProofTransactor: NonExistenceProofTransactor{contract: contract}, NonExistenceProofFilterer: NonExistenceProofFilterer{contract: contract}}, nil
}

// NonExistenceProof is an auto generated Go binding around an Ethereum contract.
type NonExistenceProof struct {
	NonExistenceProofCaller     // Read-only binding to the contract
	NonExistenceProofTransactor // Write-only binding to the contract
	NonExistenceProofFilterer   // Log filterer for contract events
}

// NonExistenceProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type NonExistenceProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonExistenceProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NonExistenceProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonExistenceProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NonExistenceProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonExistenceProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NonExistenceProofSession struct {
	Contract     *NonExistenceProof // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NonExistenceProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NonExistenceProofCallerSession struct {
	Contract *NonExistenceProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// NonExistenceProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NonExistenceProofTransactorSession struct {
	Contract     *NonExistenceProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// NonExistenceProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type NonExistenceProofRaw struct {
	Contract *NonExistenceProof // Generic contract binding to access the raw methods on
}

// NonExistenceProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NonExistenceProofCallerRaw struct {
	Contract *NonExistenceProofCaller // Generic read-only contract binding to access the raw methods on
}

// NonExistenceProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NonExistenceProofTransactorRaw struct {
	Contract *NonExistenceProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNonExistenceProof creates a new instance of NonExistenceProof, bound to a specific deployed contract.
func NewNonExistenceProof(address common.Address, backend bind.ContractBackend) (*NonExistenceProof, error) {
	contract, err := bindNonExistenceProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NonExistenceProof{NonExistenceProofCaller: NonExistenceProofCaller{contract: contract}, NonExistenceProofTransactor: NonExistenceProofTransactor{contract: contract}, NonExistenceProofFilterer: NonExistenceProofFilterer{contract: contract}}, nil
}

// NewNonExistenceProofCaller creates a new read-only instance of NonExistenceProof, bound to a specific deployed contract.
func NewNonExistenceProofCaller(address common.Address, caller bind.ContractCaller) (*NonExistenceProofCaller, error) {
	contract, err := bindNonExistenceProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NonExistenceProofCaller{contract: contract}, nil
}

// NewNonExistenceProofTransactor creates a new write-only instance of NonExistenceProof, bound to a specific deployed contract.
func NewNonExistenceProofTransactor(address common.Address, transactor bind.ContractTransactor) (*NonExistenceProofTransactor, error) {
	contract, err := bindNonExistenceProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NonExistenceProofTransactor{contract: contract}, nil
}

// NewNonExistenceProofFilterer creates a new log filterer instance of NonExistenceProof, bound to a specific deployed contract.
func NewNonExistenceProofFilterer(address common.Address, filterer bind.ContractFilterer) (*NonExistenceProofFilterer, error) {
	contract, err := bindNonExistenceProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NonExistenceProofFilterer{contract: contract}, nil
}

// bindNonExistenceProof binds a generic wrapper to an already deployed contract.
func bindNonExistenceProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NonExistenceProofMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonExistenceProof *NonExistenceProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NonExistenceProof.Contract.NonExistenceProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonExistenceProof *NonExistenceProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonExistenceProof.Contract.NonExistenceProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonExistenceProof *NonExistenceProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonExistenceProof.Contract.NonExistenceProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonExistenceProof *NonExistenceProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NonExistenceProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonExistenceProof *NonExistenceProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonExistenceProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonExistenceProof *NonExistenceProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonExistenceProof.Contract.contract.Transact(opts, method, params...)
}

// OpsMetaData contains all meta data concerning the Ops contract.
var OpsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220760d1162199714a08694121c534c4417fdcaa26ab8e4e226a640493735b01d7f64736f6c63430008120033",
}

// OpsABI is the input ABI used to generate the binding from.
// Deprecated: Use OpsMetaData.ABI instead.
var OpsABI = OpsMetaData.ABI

// OpsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OpsMetaData.Bin instead.
var OpsBin = OpsMetaData.Bin

// DeployOps deploys a new Ethereum contract, binding an instance of Ops to it.
func DeployOps(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ops, error) {
	parsed, err := OpsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OpsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ops{OpsCaller: OpsCaller{contract: contract}, OpsTransactor: OpsTransactor{contract: contract}, OpsFilterer: OpsFilterer{contract: contract}}, nil
}

// Ops is an auto generated Go binding around an Ethereum contract.
type Ops struct {
	OpsCaller     // Read-only binding to the contract
	OpsTransactor // Write-only binding to the contract
	OpsFilterer   // Log filterer for contract events
}

// OpsCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpsSession struct {
	Contract     *Ops              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpsCallerSession struct {
	Contract *OpsCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OpsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpsTransactorSession struct {
	Contract     *OpsTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpsRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpsRaw struct {
	Contract *Ops // Generic contract binding to access the raw methods on
}

// OpsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpsCallerRaw struct {
	Contract *OpsCaller // Generic read-only contract binding to access the raw methods on
}

// OpsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpsTransactorRaw struct {
	Contract *OpsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOps creates a new instance of Ops, bound to a specific deployed contract.
func NewOps(address common.Address, backend bind.ContractBackend) (*Ops, error) {
	contract, err := bindOps(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ops{OpsCaller: OpsCaller{contract: contract}, OpsTransactor: OpsTransactor{contract: contract}, OpsFilterer: OpsFilterer{contract: contract}}, nil
}

// NewOpsCaller creates a new read-only instance of Ops, bound to a specific deployed contract.
func NewOpsCaller(address common.Address, caller bind.ContractCaller) (*OpsCaller, error) {
	contract, err := bindOps(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpsCaller{contract: contract}, nil
}

// NewOpsTransactor creates a new write-only instance of Ops, bound to a specific deployed contract.
func NewOpsTransactor(address common.Address, transactor bind.ContractTransactor) (*OpsTransactor, error) {
	contract, err := bindOps(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpsTransactor{contract: contract}, nil
}

// NewOpsFilterer creates a new log filterer instance of Ops, bound to a specific deployed contract.
func NewOpsFilterer(address common.Address, filterer bind.ContractFilterer) (*OpsFilterer, error) {
	contract, err := bindOps(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpsFilterer{contract: contract}, nil
}

// bindOps binds a generic wrapper to an already deployed contract.
func bindOps(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OpsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ops *OpsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ops.Contract.OpsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ops *OpsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ops.Contract.OpsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ops *OpsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ops.Contract.OpsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ops *OpsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ops.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ops *OpsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ops.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ops *OpsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ops.Contract.contract.Transact(opts, method, params...)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PROOFSPROTOGLOBALENUMSMetaData contains all meta data concerning the PROOFSPROTOGLOBALENUMS contract.
var PROOFSPROTOGLOBALENUMSMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220f4571f0355a5935b8e78597430e352dfa44ff011a7228eaa095824b68ce91a0364736f6c63430008120033",
}

// PROOFSPROTOGLOBALENUMSABI is the input ABI used to generate the binding from.
// Deprecated: Use PROOFSPROTOGLOBALENUMSMetaData.ABI instead.
var PROOFSPROTOGLOBALENUMSABI = PROOFSPROTOGLOBALENUMSMetaData.ABI

// PROOFSPROTOGLOBALENUMSBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PROOFSPROTOGLOBALENUMSMetaData.Bin instead.
var PROOFSPROTOGLOBALENUMSBin = PROOFSPROTOGLOBALENUMSMetaData.Bin

// DeployPROOFSPROTOGLOBALENUMS deploys a new Ethereum contract, binding an instance of PROOFSPROTOGLOBALENUMS to it.
func DeployPROOFSPROTOGLOBALENUMS(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PROOFSPROTOGLOBALENUMS, error) {
	parsed, err := PROOFSPROTOGLOBALENUMSMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PROOFSPROTOGLOBALENUMSBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PROOFSPROTOGLOBALENUMS{PROOFSPROTOGLOBALENUMSCaller: PROOFSPROTOGLOBALENUMSCaller{contract: contract}, PROOFSPROTOGLOBALENUMSTransactor: PROOFSPROTOGLOBALENUMSTransactor{contract: contract}, PROOFSPROTOGLOBALENUMSFilterer: PROOFSPROTOGLOBALENUMSFilterer{contract: contract}}, nil
}

// PROOFSPROTOGLOBALENUMS is an auto generated Go binding around an Ethereum contract.
type PROOFSPROTOGLOBALENUMS struct {
	PROOFSPROTOGLOBALENUMSCaller     // Read-only binding to the contract
	PROOFSPROTOGLOBALENUMSTransactor // Write-only binding to the contract
	PROOFSPROTOGLOBALENUMSFilterer   // Log filterer for contract events
}

// PROOFSPROTOGLOBALENUMSCaller is an auto generated read-only Go binding around an Ethereum contract.
type PROOFSPROTOGLOBALENUMSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROOFSPROTOGLOBALENUMSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PROOFSPROTOGLOBALENUMSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROOFSPROTOGLOBALENUMSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PROOFSPROTOGLOBALENUMSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROOFSPROTOGLOBALENUMSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PROOFSPROTOGLOBALENUMSSession struct {
	Contract     *PROOFSPROTOGLOBALENUMS // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PROOFSPROTOGLOBALENUMSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PROOFSPROTOGLOBALENUMSCallerSession struct {
	Contract *PROOFSPROTOGLOBALENUMSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// PROOFSPROTOGLOBALENUMSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PROOFSPROTOGLOBALENUMSTransactorSession struct {
	Contract     *PROOFSPROTOGLOBALENUMSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// PROOFSPROTOGLOBALENUMSRaw is an auto generated low-level Go binding around an Ethereum contract.
type PROOFSPROTOGLOBALENUMSRaw struct {
	Contract *PROOFSPROTOGLOBALENUMS // Generic contract binding to access the raw methods on
}

// PROOFSPROTOGLOBALENUMSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PROOFSPROTOGLOBALENUMSCallerRaw struct {
	Contract *PROOFSPROTOGLOBALENUMSCaller // Generic read-only contract binding to access the raw methods on
}

// PROOFSPROTOGLOBALENUMSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PROOFSPROTOGLOBALENUMSTransactorRaw struct {
	Contract *PROOFSPROTOGLOBALENUMSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPROOFSPROTOGLOBALENUMS creates a new instance of PROOFSPROTOGLOBALENUMS, bound to a specific deployed contract.
func NewPROOFSPROTOGLOBALENUMS(address common.Address, backend bind.ContractBackend) (*PROOFSPROTOGLOBALENUMS, error) {
	contract, err := bindPROOFSPROTOGLOBALENUMS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PROOFSPROTOGLOBALENUMS{PROOFSPROTOGLOBALENUMSCaller: PROOFSPROTOGLOBALENUMSCaller{contract: contract}, PROOFSPROTOGLOBALENUMSTransactor: PROOFSPROTOGLOBALENUMSTransactor{contract: contract}, PROOFSPROTOGLOBALENUMSFilterer: PROOFSPROTOGLOBALENUMSFilterer{contract: contract}}, nil
}

// NewPROOFSPROTOGLOBALENUMSCaller creates a new read-only instance of PROOFSPROTOGLOBALENUMS, bound to a specific deployed contract.
func NewPROOFSPROTOGLOBALENUMSCaller(address common.Address, caller bind.ContractCaller) (*PROOFSPROTOGLOBALENUMSCaller, error) {
	contract, err := bindPROOFSPROTOGLOBALENUMS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PROOFSPROTOGLOBALENUMSCaller{contract: contract}, nil
}

// NewPROOFSPROTOGLOBALENUMSTransactor creates a new write-only instance of PROOFSPROTOGLOBALENUMS, bound to a specific deployed contract.
func NewPROOFSPROTOGLOBALENUMSTransactor(address common.Address, transactor bind.ContractTransactor) (*PROOFSPROTOGLOBALENUMSTransactor, error) {
	contract, err := bindPROOFSPROTOGLOBALENUMS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PROOFSPROTOGLOBALENUMSTransactor{contract: contract}, nil
}

// NewPROOFSPROTOGLOBALENUMSFilterer creates a new log filterer instance of PROOFSPROTOGLOBALENUMS, bound to a specific deployed contract.
func NewPROOFSPROTOGLOBALENUMSFilterer(address common.Address, filterer bind.ContractFilterer) (*PROOFSPROTOGLOBALENUMSFilterer, error) {
	contract, err := bindPROOFSPROTOGLOBALENUMS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PROOFSPROTOGLOBALENUMSFilterer{contract: contract}, nil
}

// bindPROOFSPROTOGLOBALENUMS binds a generic wrapper to an already deployed contract.
func bindPROOFSPROTOGLOBALENUMS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PROOFSPROTOGLOBALENUMSMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROOFSPROTOGLOBALENUMS *PROOFSPROTOGLOBALENUMSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROOFSPROTOGLOBALENUMS.Contract.PROOFSPROTOGLOBALENUMSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROOFSPROTOGLOBALENUMS *PROOFSPROTOGLOBALENUMSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROOFSPROTOGLOBALENUMS.Contract.PROOFSPROTOGLOBALENUMSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROOFSPROTOGLOBALENUMS *PROOFSPROTOGLOBALENUMSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROOFSPROTOGLOBALENUMS.Contract.PROOFSPROTOGLOBALENUMSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROOFSPROTOGLOBALENUMS *PROOFSPROTOGLOBALENUMSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROOFSPROTOGLOBALENUMS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROOFSPROTOGLOBALENUMS *PROOFSPROTOGLOBALENUMSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROOFSPROTOGLOBALENUMS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROOFSPROTOGLOBALENUMS *PROOFSPROTOGLOBALENUMSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROOFSPROTOGLOBALENUMS.Contract.contract.Transact(opts, method, params...)
}

// PairingMetaData contains all meta data concerning the Pairing contract.
var PairingMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220997b5b28a1676d4c98ee143848b2f6c3ca740e808199a2f762a47ff5be11ea6064736f6c63430008120033",
}

// PairingABI is the input ABI used to generate the binding from.
// Deprecated: Use PairingMetaData.ABI instead.
var PairingABI = PairingMetaData.ABI

// PairingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PairingMetaData.Bin instead.
var PairingBin = PairingMetaData.Bin

// DeployPairing deploys a new Ethereum contract, binding an instance of Pairing to it.
func DeployPairing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pairing, error) {
	parsed, err := PairingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PairingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pairing{PairingCaller: PairingCaller{contract: contract}, PairingTransactor: PairingTransactor{contract: contract}, PairingFilterer: PairingFilterer{contract: contract}}, nil
}

// Pairing is an auto generated Go binding around an Ethereum contract.
type Pairing struct {
	PairingCaller     // Read-only binding to the contract
	PairingTransactor // Write-only binding to the contract
	PairingFilterer   // Log filterer for contract events
}

// PairingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairingSession struct {
	Contract     *Pairing          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairingCallerSession struct {
	Contract *PairingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PairingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairingTransactorSession struct {
	Contract     *PairingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PairingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairingRaw struct {
	Contract *Pairing // Generic contract binding to access the raw methods on
}

// PairingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairingCallerRaw struct {
	Contract *PairingCaller // Generic read-only contract binding to access the raw methods on
}

// PairingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairingTransactorRaw struct {
	Contract *PairingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPairing creates a new instance of Pairing, bound to a specific deployed contract.
func NewPairing(address common.Address, backend bind.ContractBackend) (*Pairing, error) {
	contract, err := bindPairing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pairing{PairingCaller: PairingCaller{contract: contract}, PairingTransactor: PairingTransactor{contract: contract}, PairingFilterer: PairingFilterer{contract: contract}}, nil
}

// NewPairingCaller creates a new read-only instance of Pairing, bound to a specific deployed contract.
func NewPairingCaller(address common.Address, caller bind.ContractCaller) (*PairingCaller, error) {
	contract, err := bindPairing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairingCaller{contract: contract}, nil
}

// NewPairingTransactor creates a new write-only instance of Pairing, bound to a specific deployed contract.
func NewPairingTransactor(address common.Address, transactor bind.ContractTransactor) (*PairingTransactor, error) {
	contract, err := bindPairing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairingTransactor{contract: contract}, nil
}

// NewPairingFilterer creates a new log filterer instance of Pairing, bound to a specific deployed contract.
func NewPairingFilterer(address common.Address, filterer bind.ContractFilterer) (*PairingFilterer, error) {
	contract, err := bindPairing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairingFilterer{contract: contract}, nil
}

// bindPairing binds a generic wrapper to an already deployed contract.
func bindPairing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PairingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairing *PairingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pairing.Contract.PairingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairing *PairingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairing.Contract.PairingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairing *PairingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairing.Contract.PairingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pairing *PairingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pairing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pairing *PairingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pairing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pairing *PairingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pairing.Contract.contract.Transact(opts, method, params...)
}

// PartSetHeaderMetaData contains all meta data concerning the PartSetHeader contract.
var PartSetHeaderMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220f337fc2d0b8832e551cd8240f12bbd59a2a15d621055d6354d60dad4d8ee59ef64736f6c63430008120033",
}

// PartSetHeaderABI is the input ABI used to generate the binding from.
// Deprecated: Use PartSetHeaderMetaData.ABI instead.
var PartSetHeaderABI = PartSetHeaderMetaData.ABI

// PartSetHeaderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PartSetHeaderMetaData.Bin instead.
var PartSetHeaderBin = PartSetHeaderMetaData.Bin

// DeployPartSetHeader deploys a new Ethereum contract, binding an instance of PartSetHeader to it.
func DeployPartSetHeader(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PartSetHeader, error) {
	parsed, err := PartSetHeaderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PartSetHeaderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PartSetHeader{PartSetHeaderCaller: PartSetHeaderCaller{contract: contract}, PartSetHeaderTransactor: PartSetHeaderTransactor{contract: contract}, PartSetHeaderFilterer: PartSetHeaderFilterer{contract: contract}}, nil
}

// PartSetHeader is an auto generated Go binding around an Ethereum contract.
type PartSetHeader struct {
	PartSetHeaderCaller     // Read-only binding to the contract
	PartSetHeaderTransactor // Write-only binding to the contract
	PartSetHeaderFilterer   // Log filterer for contract events
}

// PartSetHeaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type PartSetHeaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartSetHeaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PartSetHeaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartSetHeaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PartSetHeaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PartSetHeaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PartSetHeaderSession struct {
	Contract     *PartSetHeader    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PartSetHeaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PartSetHeaderCallerSession struct {
	Contract *PartSetHeaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PartSetHeaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PartSetHeaderTransactorSession struct {
	Contract     *PartSetHeaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PartSetHeaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type PartSetHeaderRaw struct {
	Contract *PartSetHeader // Generic contract binding to access the raw methods on
}

// PartSetHeaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PartSetHeaderCallerRaw struct {
	Contract *PartSetHeaderCaller // Generic read-only contract binding to access the raw methods on
}

// PartSetHeaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PartSetHeaderTransactorRaw struct {
	Contract *PartSetHeaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPartSetHeader creates a new instance of PartSetHeader, bound to a specific deployed contract.
func NewPartSetHeader(address common.Address, backend bind.ContractBackend) (*PartSetHeader, error) {
	contract, err := bindPartSetHeader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PartSetHeader{PartSetHeaderCaller: PartSetHeaderCaller{contract: contract}, PartSetHeaderTransactor: PartSetHeaderTransactor{contract: contract}, PartSetHeaderFilterer: PartSetHeaderFilterer{contract: contract}}, nil
}

// NewPartSetHeaderCaller creates a new read-only instance of PartSetHeader, bound to a specific deployed contract.
func NewPartSetHeaderCaller(address common.Address, caller bind.ContractCaller) (*PartSetHeaderCaller, error) {
	contract, err := bindPartSetHeader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PartSetHeaderCaller{contract: contract}, nil
}

// NewPartSetHeaderTransactor creates a new write-only instance of PartSetHeader, bound to a specific deployed contract.
func NewPartSetHeaderTransactor(address common.Address, transactor bind.ContractTransactor) (*PartSetHeaderTransactor, error) {
	contract, err := bindPartSetHeader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PartSetHeaderTransactor{contract: contract}, nil
}

// NewPartSetHeaderFilterer creates a new log filterer instance of PartSetHeader, bound to a specific deployed contract.
func NewPartSetHeaderFilterer(address common.Address, filterer bind.ContractFilterer) (*PartSetHeaderFilterer, error) {
	contract, err := bindPartSetHeader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PartSetHeaderFilterer{contract: contract}, nil
}

// bindPartSetHeader binds a generic wrapper to an already deployed contract.
func bindPartSetHeader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PartSetHeaderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PartSetHeader *PartSetHeaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PartSetHeader.Contract.PartSetHeaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PartSetHeader *PartSetHeaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartSetHeader.Contract.PartSetHeaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PartSetHeader *PartSetHeaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PartSetHeader.Contract.PartSetHeaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PartSetHeader *PartSetHeaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PartSetHeader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PartSetHeader *PartSetHeaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PartSetHeader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PartSetHeader *PartSetHeaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PartSetHeader.Contract.contract.Transact(opts, method, params...)
}

// PoALightClientMetaData contains all meta data concerning the PoALightClient contract.
var PoALightClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bscValidatorSet\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"bscValidatorSet\",\"outputs\":[{\"internalType\":\"contractIBSCValidatorSet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizedExecutionStateRootAndSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimisticExecutionStateRootAndSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sha3Uncles\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"logsBloom\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasUsed\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"mixHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes8\",\"name\":\"nonce\",\"type\":\"bytes8\"}],\"internalType\":\"structPoALightClient.BNBHeaderInfo\",\"name\":\"header\",\"type\":\"tuple\"}],\"name\":\"updateHeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60803461007557601f61172638819003918201601f19168301916001600160401b0383118484101761007a5780849260209460405283398101031261007557516001600160a01b038116908190036100755763ffffffff60e01b6000541617600055600060015560405161169590816100918239f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b6000803560e01c90816312420766146101a457816343a6c5a6146101a45781636afa002e14610050575063e89a95011461004b57600080fd5b6101db565b346101a15760031960203682011261019d5760043567ffffffffffffffff91828211610199576101e090823603011261019557806101046001600160a01b039201359260206100e286546100aa848260a01c168811610202565b6100b6856004016105f2565b60405163155853f360e21b81526001600160a01b03909116600482015295869283919082906024820190565b0392165afa9081156101905761015f94610109610158936064968991610162575b506102f3565b167fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff00000000000000000000000000000000000000006000549260a01b16911617600055565b0135600155565b80f35b610183915060203d8111610189575b61017b81836102ad565b8101906102cf565b89610103565b503d610171565b6102e7565b8280fd5b8380fd5b5080fd5b80fd5b346101d65760003660031901126101d657604060015467ffffffffffffffff60005460a01c1682519182526020820152f35b600080fd5b346101d65760003660031901126101d65760206001600160a01b0360005416604051908152f35b1561020957565b60405162461bcd60e51b8152602060048201526024808201527f506f414c69676874436c69656e743a20696e76616c696420626c6f636b206e7560448201526336b132b960e11b6064820152608490fd5b634e487b7160e01b600052604160045260246000fd5b6040810190811067ffffffffffffffff82111761028c57604052565b61025a565b6020810190811067ffffffffffffffff82111761028c57604052565b90601f8019910116810190811067ffffffffffffffff82111761028c57604052565b908160209103126101d6575180151581036101d65790565b6040513d6000823e3d90fd5b156102fa57565b60405162461bcd60e51b815260206004820152602660248201527f506f414c69676874436c69656e743a20696e76616c6964207369676e6572206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608490fd5b903590601e19813603018212156101d6570180359067ffffffffffffffff82116101d6576020019181360383136101d657565b604051906101e0820182811067ffffffffffffffff82111761028c57604052565b67ffffffffffffffff811161028c57601f01601f191660200190565b9291926103e1826103b9565b916103ef60405193846102ad565b8294818452818301116101d6578281602093846000960137010152565b1561041357565b60405162461bcd60e51b815260206004820152602d60248201527f506f414c69676874436c69656e743a20696e76616c696420657874726120646160448201527f746120666f722076616e697479000000000000000000000000000000000000006064820152608490fd5b634e487b7160e01b600052601160045260246000fd5b90600182018092116104a257565b61047e565b919082018092116104a257565b156104bb57565b60405162461bcd60e51b815260206004820152603060248201527f506f414c69676874436c69656e743a20696e76616c696420657874726120646160448201527f746120666f72207369676e6174757265000000000000000000000000000000006064820152608490fd5b6040198101919082116104a257565b919082039182116104a257565b1561054957565b60405162461bcd60e51b815260206004820152602a60248201527f506f414c69676874436c69656e743a207369676e61747572652072657472696560448201527f76616c206661696c6564000000000000000000000000000000000000000000006064820152608490fd5b3567ffffffffffffffff811681036101d65790565b356001600160a01b03811681036101d65790565b356001600160c01b0319811681036101d65790565b61018061060182820183610365565b369061060c926103d5565b91825160201061061b9061040c565b82516061111561062a906104b4565b825161063590610526565b61063f9084610f83565b92805161064b90610526565b81519061065792611031565b92835160411461066690610542565b610120916106758184016105b4565b9161014094858301610686906105b4565b61069360c0850185610365565b916106a0604087016105c9565b986101c0976106b0888a016105dd565b99610160946106c08a87016105b4565b966106c9610398565b8b3581526020808d0135908201526001600160a01b03909e1660408f01528d60608c013590606001528d60808c013590608001528d60a08c01359060a001523690610713926103d5565b60c08d015260e089013560e08d0152610100808a0135908d01528b0190610743919067ffffffffffffffff169052565b67ffffffffffffffff9091169089015267ffffffffffffffff909116908701528501526101a090810135908401526001600160c01b03199091169082015261078a90610911565b9061079491610c4e565b90565b60405190610220820182811067ffffffffffffffff82111761028c57604052601082528160005b61020081106107cb575050565b8060606020809385010152016107be565b634e487b7160e01b600052603260045260246000fd5b8051156107ff5760200190565b6107dc565b8051600110156107ff5760400190565b8051600210156107ff5760600190565b8051600310156107ff5760800190565b8051600410156107ff5760a00190565b8051600510156107ff5760c00190565b8051600610156107ff5760e00190565b8051600710156107ff576101000190565b8051600810156107ff576101200190565b8051600910156107ff576101400190565b8051600a10156107ff576101600190565b8051600b10156107ff576101800190565b8051600c10156107ff576101a00190565b8051600d10156107ff576101c00190565b8051600e10156107ff576101e00190565b8051600f10156107ff576102000190565b80518210156107ff5760209160051b010190565b610c4661091c610797565b61092c610927611458565b611146565b610935826107f2565b5261093f816107f2565b50610c2d610c216109278551956040610c066101c08251936109818561096d60209d8e830160209181520190565b0395610927601f19978881018352826102ad565b61098a8a610804565b5261099489610804565b506109bf8b8201516109278d6109b38851938492830160209181520190565b038881018352826102ad565b6109c88a610814565b526109d289610814565b506109ee6109e9858301516001600160a01b031690565b6113b9565b6109f78a610824565b52610a0189610824565b50610a2160608201516109278d6109b38851938492830160209181520190565b610a2a8a610834565b52610a3489610834565b50610a5460808201516109278d6109b38851938492830160209181520190565b610a5d8a610844565b52610a6789610844565b50610a8760a08201516109278d6109b38851938492830160209181520190565b610a908a610854565b52610a9a89610854565b50610aa860c0820151611146565b610ab18a610864565b52610abb89610864565b50610ac960e08201516113e6565b610ad28a610875565b52610adc89610875565b50610aeb6101008201516113e6565b610af48a610886565b52610afe89610886565b50610b2d610b28610b1b61012084015167ffffffffffffffff1690565b67ffffffffffffffff1690565b6113e6565b610b368a610897565b52610b4089610897565b50610b5d610b28610b1b61014084015167ffffffffffffffff1690565b610b668a6108a8565b52610b70896108a8565b50610b8d610b28610b1b61016084015167ffffffffffffffff1690565b610b968a6108b9565b52610ba0896108b9565b50610baf610180820151611146565b610bb88a6108ca565b52610bc2896108ca565b50610be36101a08201516109278d6109b38851938492830160209181520190565b610bec8a6108db565b52610bf6896108db565b5001516001600160c01b03191690565b90519384918983016001600160c01b03196008921681520190565b039081018352826102ad565b610c36826108ec565b52610c40816108ec565b506112bd565b805191012090565b6041825103610d3e57816000916020809401519160606040820151910151841a92610c9b7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610d83565b83158015610d34575b610d1d575b90610ced91610cc760ff8616601b8114908115610d12575b50610e30565b604051948594859094939260ff6060936080840197845216602083015260408201520152565b838052039060015afa15610190576000516107946001600160a01b0382161515610e87565b601c91501438610cc1565b9290610d2b610ced92610dda565b93909150610ca9565b5060018414610ca4565b60405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606490fd5b15610d8a57565b60405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608490fd5b60ff601b9116019060ff82116104a257565b60ff166080019060ff82116104a257565b60ff1660c0019060ff82116104a257565b60ff1660b7019060ff82116104a257565b60ff1660f7019060ff82116104a257565b15610e3757565b60405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608490fd5b15610e8e57565b60405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606490fd5b15610eda57565b60405162461bcd60e51b815260206004820152601860248201527f4d656d6f72793a20746f206f7574206f6620626f756e647300000000000000006044820152606490fd5b60405190610f2c82610270565b6001825260203681840137565b604051610f4581610291565b60008152906000368137565b90610f5b826103b9565b610f6860405191826102ad565b8281528092610f79601f19916103b9565b0190602036910137565b90801561101c578151151580611014575b15610fcf5761079490825181111580610fc7575b610fb190610ed3565b6020610fbc82610f51565b9301602084016110b2565b506001610fa8565b60405162461bcd60e51b815260206004820152601a60248201527f4d656d6f72793a2066726f6d206f7574206f6620626f756e64730000000000006044820152606490fd5b506001610f94565b505060405161102a81610291565b6000815290565b91808210156110a357825182108061109b575b15610fcf5761108961079492845183111580611093575b61106490610ed3565b61108381602061107c6110778388610535565b610f51565b97016104a7565b92610535565b90602084016110b2565b50600161105b565b506001611044565b50505060405161102a81610291565b92905b6020938484106110ea57815181528481018091116104a2579381018091116104a25791601f1981019081116104a257916110b5565b9290919350600019906020036101000a0190811990511690825116179052565b9081518110156107ff570160200190565b9081519160005b838110611133575050016000815290565b8060208092840101518185015201611122565b9060009180519260019384811490816112a5575b501561116557509150565b81519360388510156111d75750926111c9610794926111c394956111a661119660ff61118f610f1f565b9616610dec565b60f81b6001600160f81b03191690565b901a6111b1846107f2565b535b604051948593602085019061111b565b9061111b565b03601f1981018352826102ad565b9190808380805b611271575b50506111f161107783610494565b9361120161119660ff8516610e0e565b821a61120c866107f2565b535b8281111561122c575050506111c3929350906111c9610794926111b3565b8061125a61119661125461125461124e61124961126c978a610535565b611438565b8c6113f2565b60ff1690565b831a611266828861110a565b53611429565b61120e565b61127e81899593956113f2565b1561129d5761128f61129591611412565b91611429565b9290816111de565b9290926111e3565b9050156107ff576080602083015160f81c103861115a565b6112c6906115c7565b8051906000603883101561130357506111c3916111c9610794926112f761119660ff6112f0610f1f565b9516610dfd565b60001a6111b1846107f2565b6001928380805b61138b575b505061131d61107783610494565b9361132d61119660ff8516610e1f565b60001a611339866107f2565b535b82811115611356575050506111c3916111c9610794926111b3565b80611379611196611254611254611373611249611386978a610535565b876113f2565b60001a611266828861110a565b61133b565b61139881849593956113f2565b156113b15761128f6113a991611412565b92908161130a565b92909261130f565b61079490604051906bffffffffffffffffffffffff199060601b1660208201526014815261092781610270565b610927610794916114f0565b81156113fc570490565b634e487b7160e01b600052601260045260246000fd5b908160081b9180830461010014901517156104a257565b60001981146104a25760010190565b601f81116104a2576101000a90565b60ff166020039060ff82116104a257565b6040516020808201916061835281815261147181610270565b60005b60ff90818116848110156114e5576001600160f81b031990611496908561110a565b51166114a55760010116611474565b92919390505b816114c060ff6114ba86611447565b16610f51565b9460ff851601018091116104a2576114dd61125461079494611447565b9184016110b2565b5092919390506114ab565b80156115af576040519160209161152084611512858201938460209181520190565b03601f1981018652856102ad565b60005b60ff908181168581101561159e5761154f61154161155c928961110a565b516001600160f81b03191690565b6001600160f81b03191690565b6115695760010116611523565b90506107949294509290925b6114dd61125461159861158d61107761125486611447565b9660ff8516906104a7565b92611447565b509050610794929450929092611575565b506115b8610f1f565b60006115c3826107f2565b5390565b8051156116565790600091825b81518410156115fd576115f56001916115ed86856108fd565b5151906104a7565b9301926115d4565b61160991929350610f51565b906020808301936000945b835186101561164e5761164660019161163c61163089886108fd565b518681519101836110b2565b6115ed88876108fd565b950194611614565b509350505090565b50610794610f3956fea264697066735822122027864f18204125fb3295ef8cd7dd425e3f04385a05e41adba4139bf93fe86f7964736f6c63430008120033",
}

// PoALightClientABI is the input ABI used to generate the binding from.
// Deprecated: Use PoALightClientMetaData.ABI instead.
var PoALightClientABI = PoALightClientMetaData.ABI

// PoALightClientBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoALightClientMetaData.Bin instead.
var PoALightClientBin = PoALightClientMetaData.Bin

// DeployPoALightClient deploys a new Ethereum contract, binding an instance of PoALightClient to it.
func DeployPoALightClient(auth *bind.TransactOpts, backend bind.ContractBackend, _bscValidatorSet common.Address) (common.Address, *types.Transaction, *PoALightClient, error) {
	parsed, err := PoALightClientMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoALightClientBin), backend, _bscValidatorSet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PoALightClient{PoALightClientCaller: PoALightClientCaller{contract: contract}, PoALightClientTransactor: PoALightClientTransactor{contract: contract}, PoALightClientFilterer: PoALightClientFilterer{contract: contract}}, nil
}

// PoALightClient is an auto generated Go binding around an Ethereum contract.
type PoALightClient struct {
	PoALightClientCaller     // Read-only binding to the contract
	PoALightClientTransactor // Write-only binding to the contract
	PoALightClientFilterer   // Log filterer for contract events
}

// PoALightClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoALightClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoALightClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoALightClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoALightClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoALightClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoALightClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoALightClientSession struct {
	Contract     *PoALightClient   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoALightClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoALightClientCallerSession struct {
	Contract *PoALightClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PoALightClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoALightClientTransactorSession struct {
	Contract     *PoALightClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PoALightClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoALightClientRaw struct {
	Contract *PoALightClient // Generic contract binding to access the raw methods on
}

// PoALightClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoALightClientCallerRaw struct {
	Contract *PoALightClientCaller // Generic read-only contract binding to access the raw methods on
}

// PoALightClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoALightClientTransactorRaw struct {
	Contract *PoALightClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoALightClient creates a new instance of PoALightClient, bound to a specific deployed contract.
func NewPoALightClient(address common.Address, backend bind.ContractBackend) (*PoALightClient, error) {
	contract, err := bindPoALightClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoALightClient{PoALightClientCaller: PoALightClientCaller{contract: contract}, PoALightClientTransactor: PoALightClientTransactor{contract: contract}, PoALightClientFilterer: PoALightClientFilterer{contract: contract}}, nil
}

// NewPoALightClientCaller creates a new read-only instance of PoALightClient, bound to a specific deployed contract.
func NewPoALightClientCaller(address common.Address, caller bind.ContractCaller) (*PoALightClientCaller, error) {
	contract, err := bindPoALightClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoALightClientCaller{contract: contract}, nil
}

// NewPoALightClientTransactor creates a new write-only instance of PoALightClient, bound to a specific deployed contract.
func NewPoALightClientTransactor(address common.Address, transactor bind.ContractTransactor) (*PoALightClientTransactor, error) {
	contract, err := bindPoALightClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoALightClientTransactor{contract: contract}, nil
}

// NewPoALightClientFilterer creates a new log filterer instance of PoALightClient, bound to a specific deployed contract.
func NewPoALightClientFilterer(address common.Address, filterer bind.ContractFilterer) (*PoALightClientFilterer, error) {
	contract, err := bindPoALightClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoALightClientFilterer{contract: contract}, nil
}

// bindPoALightClient binds a generic wrapper to an already deployed contract.
func bindPoALightClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoALightClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoALightClient *PoALightClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoALightClient.Contract.PoALightClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoALightClient *PoALightClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoALightClient.Contract.PoALightClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoALightClient *PoALightClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoALightClient.Contract.PoALightClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoALightClient *PoALightClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoALightClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoALightClient *PoALightClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoALightClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoALightClient *PoALightClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoALightClient.Contract.contract.Transact(opts, method, params...)
}

// BscValidatorSet is a free data retrieval call binding the contract method 0xe89a9501.
//
// Solidity: function bscValidatorSet() view returns(address)
func (_PoALightClient *PoALightClientCaller) BscValidatorSet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoALightClient.contract.Call(opts, &out, "bscValidatorSet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BscValidatorSet is a free data retrieval call binding the contract method 0xe89a9501.
//
// Solidity: function bscValidatorSet() view returns(address)
func (_PoALightClient *PoALightClientSession) BscValidatorSet() (common.Address, error) {
	return _PoALightClient.Contract.BscValidatorSet(&_PoALightClient.CallOpts)
}

// BscValidatorSet is a free data retrieval call binding the contract method 0xe89a9501.
//
// Solidity: function bscValidatorSet() view returns(address)
func (_PoALightClient *PoALightClientCallerSession) BscValidatorSet() (common.Address, error) {
	return _PoALightClient.Contract.BscValidatorSet(&_PoALightClient.CallOpts)
}

// FinalizedExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x12420766.
//
// Solidity: function finalizedExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_PoALightClient *PoALightClientCaller) FinalizedExecutionStateRootAndSlot(opts *bind.CallOpts) (struct {
	Root [32]byte
	Slot uint64
}, error) {
	var out []interface{}
	err := _PoALightClient.contract.Call(opts, &out, "finalizedExecutionStateRootAndSlot")

	outstruct := new(struct {
		Root [32]byte
		Slot uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Slot = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// FinalizedExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x12420766.
//
// Solidity: function finalizedExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_PoALightClient *PoALightClientSession) FinalizedExecutionStateRootAndSlot() (struct {
	Root [32]byte
	Slot uint64
}, error) {
	return _PoALightClient.Contract.FinalizedExecutionStateRootAndSlot(&_PoALightClient.CallOpts)
}

// FinalizedExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x12420766.
//
// Solidity: function finalizedExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_PoALightClient *PoALightClientCallerSession) FinalizedExecutionStateRootAndSlot() (struct {
	Root [32]byte
	Slot uint64
}, error) {
	return _PoALightClient.Contract.FinalizedExecutionStateRootAndSlot(&_PoALightClient.CallOpts)
}

// OptimisticExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x43a6c5a6.
//
// Solidity: function optimisticExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_PoALightClient *PoALightClientCaller) OptimisticExecutionStateRootAndSlot(opts *bind.CallOpts) (struct {
	Root [32]byte
	Slot uint64
}, error) {
	var out []interface{}
	err := _PoALightClient.contract.Call(opts, &out, "optimisticExecutionStateRootAndSlot")

	outstruct := new(struct {
		Root [32]byte
		Slot uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Slot = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OptimisticExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x43a6c5a6.
//
// Solidity: function optimisticExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_PoALightClient *PoALightClientSession) OptimisticExecutionStateRootAndSlot() (struct {
	Root [32]byte
	Slot uint64
}, error) {
	return _PoALightClient.Contract.OptimisticExecutionStateRootAndSlot(&_PoALightClient.CallOpts)
}

// OptimisticExecutionStateRootAndSlot is a free data retrieval call binding the contract method 0x43a6c5a6.
//
// Solidity: function optimisticExecutionStateRootAndSlot() view returns(bytes32 root, uint64 slot)
func (_PoALightClient *PoALightClientCallerSession) OptimisticExecutionStateRootAndSlot() (struct {
	Root [32]byte
	Slot uint64
}, error) {
	return _PoALightClient.Contract.OptimisticExecutionStateRootAndSlot(&_PoALightClient.CallOpts)
}

// UpdateHeader is a paid mutator transaction binding the contract method 0x6afa002e.
//
// Solidity: function updateHeader((bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes,uint256,uint256,uint64,uint64,uint64,bytes,bytes32,bytes8) header) returns()
func (_PoALightClient *PoALightClientTransactor) UpdateHeader(opts *bind.TransactOpts, header PoALightClientBNBHeaderInfo) (*types.Transaction, error) {
	return _PoALightClient.contract.Transact(opts, "updateHeader", header)
}

// UpdateHeader is a paid mutator transaction binding the contract method 0x6afa002e.
//
// Solidity: function updateHeader((bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes,uint256,uint256,uint64,uint64,uint64,bytes,bytes32,bytes8) header) returns()
func (_PoALightClient *PoALightClientSession) UpdateHeader(header PoALightClientBNBHeaderInfo) (*types.Transaction, error) {
	return _PoALightClient.Contract.UpdateHeader(&_PoALightClient.TransactOpts, header)
}

// UpdateHeader is a paid mutator transaction binding the contract method 0x6afa002e.
//
// Solidity: function updateHeader((bytes32,bytes32,address,bytes32,bytes32,bytes32,bytes,uint256,uint256,uint64,uint64,uint64,bytes,bytes32,bytes8) header) returns()
func (_PoALightClient *PoALightClientTransactorSession) UpdateHeader(header PoALightClientBNBHeaderInfo) (*types.Transaction, error) {
	return _PoALightClient.Contract.UpdateHeader(&_PoALightClient.TransactOpts, header)
}

// ProofMetaData contains all meta data concerning the Proof contract.
var ProofMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220d9c2ea657805d357f5ff614affd887136d9ee6efa5100ababf4216ae001acb4164736f6c63430008120033",
}

// ProofABI is the input ABI used to generate the binding from.
// Deprecated: Use ProofMetaData.ABI instead.
var ProofABI = ProofMetaData.ABI

// ProofBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProofMetaData.Bin instead.
var ProofBin = ProofMetaData.Bin

// DeployProof deploys a new Ethereum contract, binding an instance of Proof to it.
func DeployProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Proof, error) {
	parsed, err := ProofMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Proof{ProofCaller: ProofCaller{contract: contract}, ProofTransactor: ProofTransactor{contract: contract}, ProofFilterer: ProofFilterer{contract: contract}}, nil
}

// Proof is an auto generated Go binding around an Ethereum contract.
type Proof struct {
	ProofCaller     // Read-only binding to the contract
	ProofTransactor // Write-only binding to the contract
	ProofFilterer   // Log filterer for contract events
}

// ProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProofSession struct {
	Contract     *Proof            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProofCallerSession struct {
	Contract *ProofCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProofTransactorSession struct {
	Contract     *ProofTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProofRaw struct {
	Contract *Proof // Generic contract binding to access the raw methods on
}

// ProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProofCallerRaw struct {
	Contract *ProofCaller // Generic read-only contract binding to access the raw methods on
}

// ProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProofTransactorRaw struct {
	Contract *ProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProof creates a new instance of Proof, bound to a specific deployed contract.
func NewProof(address common.Address, backend bind.ContractBackend) (*Proof, error) {
	contract, err := bindProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proof{ProofCaller: ProofCaller{contract: contract}, ProofTransactor: ProofTransactor{contract: contract}, ProofFilterer: ProofFilterer{contract: contract}}, nil
}

// NewProofCaller creates a new read-only instance of Proof, bound to a specific deployed contract.
func NewProofCaller(address common.Address, caller bind.ContractCaller) (*ProofCaller, error) {
	contract, err := bindProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProofCaller{contract: contract}, nil
}

// NewProofTransactor creates a new write-only instance of Proof, bound to a specific deployed contract.
func NewProofTransactor(address common.Address, transactor bind.ContractTransactor) (*ProofTransactor, error) {
	contract, err := bindProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProofTransactor{contract: contract}, nil
}

// NewProofFilterer creates a new log filterer instance of Proof, bound to a specific deployed contract.
func NewProofFilterer(address common.Address, filterer bind.ContractFilterer) (*ProofFilterer, error) {
	contract, err := bindProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProofFilterer{contract: contract}, nil
}

// bindProof binds a generic wrapper to an already deployed contract.
func bindProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProofMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proof *ProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proof.Contract.ProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proof *ProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proof.Contract.ProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proof *ProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proof.Contract.ProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proof *ProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proof *ProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proof *ProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proof.Contract.contract.Transact(opts, method, params...)
}

// ProofSpecMetaData contains all meta data concerning the ProofSpec contract.
var ProofSpecMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220acc45d92a6b2f71a117eeeaab9c7e073993799c994dfc7bf7bca89dadb2bdd2964736f6c63430008120033",
}

// ProofSpecABI is the input ABI used to generate the binding from.
// Deprecated: Use ProofSpecMetaData.ABI instead.
var ProofSpecABI = ProofSpecMetaData.ABI

// ProofSpecBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProofSpecMetaData.Bin instead.
var ProofSpecBin = ProofSpecMetaData.Bin

// DeployProofSpec deploys a new Ethereum contract, binding an instance of ProofSpec to it.
func DeployProofSpec(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProofSpec, error) {
	parsed, err := ProofSpecMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProofSpecBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProofSpec{ProofSpecCaller: ProofSpecCaller{contract: contract}, ProofSpecTransactor: ProofSpecTransactor{contract: contract}, ProofSpecFilterer: ProofSpecFilterer{contract: contract}}, nil
}

// ProofSpec is an auto generated Go binding around an Ethereum contract.
type ProofSpec struct {
	ProofSpecCaller     // Read-only binding to the contract
	ProofSpecTransactor // Write-only binding to the contract
	ProofSpecFilterer   // Log filterer for contract events
}

// ProofSpecCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProofSpecCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofSpecTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProofSpecTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofSpecFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProofSpecFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofSpecSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProofSpecSession struct {
	Contract     *ProofSpec        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProofSpecCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProofSpecCallerSession struct {
	Contract *ProofSpecCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ProofSpecTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProofSpecTransactorSession struct {
	Contract     *ProofSpecTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ProofSpecRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProofSpecRaw struct {
	Contract *ProofSpec // Generic contract binding to access the raw methods on
}

// ProofSpecCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProofSpecCallerRaw struct {
	Contract *ProofSpecCaller // Generic read-only contract binding to access the raw methods on
}

// ProofSpecTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProofSpecTransactorRaw struct {
	Contract *ProofSpecTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProofSpec creates a new instance of ProofSpec, bound to a specific deployed contract.
func NewProofSpec(address common.Address, backend bind.ContractBackend) (*ProofSpec, error) {
	contract, err := bindProofSpec(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProofSpec{ProofSpecCaller: ProofSpecCaller{contract: contract}, ProofSpecTransactor: ProofSpecTransactor{contract: contract}, ProofSpecFilterer: ProofSpecFilterer{contract: contract}}, nil
}

// NewProofSpecCaller creates a new read-only instance of ProofSpec, bound to a specific deployed contract.
func NewProofSpecCaller(address common.Address, caller bind.ContractCaller) (*ProofSpecCaller, error) {
	contract, err := bindProofSpec(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProofSpecCaller{contract: contract}, nil
}

// NewProofSpecTransactor creates a new write-only instance of ProofSpec, bound to a specific deployed contract.
func NewProofSpecTransactor(address common.Address, transactor bind.ContractTransactor) (*ProofSpecTransactor, error) {
	contract, err := bindProofSpec(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProofSpecTransactor{contract: contract}, nil
}

// NewProofSpecFilterer creates a new log filterer instance of ProofSpec, bound to a specific deployed contract.
func NewProofSpecFilterer(address common.Address, filterer bind.ContractFilterer) (*ProofSpecFilterer, error) {
	contract, err := bindProofSpec(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProofSpecFilterer{contract: contract}, nil
}

// bindProofSpec binds a generic wrapper to an already deployed contract.
func bindProofSpec(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProofSpecMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProofSpec *ProofSpecRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProofSpec.Contract.ProofSpecCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProofSpec *ProofSpecRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProofSpec.Contract.ProofSpecTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProofSpec *ProofSpecRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProofSpec.Contract.ProofSpecTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProofSpec *ProofSpecCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProofSpec.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProofSpec *ProofSpecTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProofSpec.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProofSpec *ProofSpecTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProofSpec.Contract.contract.Transact(opts, method, params...)
}

// ProtoBufRuntimeMetaData contains all meta data concerning the ProtoBufRuntime contract.
var ProtoBufRuntimeMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea264697066735822122042aa0374c3106ab5fe6a5c370ef6c4ed37537f3eb17dd1b3098b794dd885f57264736f6c63430008120033",
}

// ProtoBufRuntimeABI is the input ABI used to generate the binding from.
// Deprecated: Use ProtoBufRuntimeMetaData.ABI instead.
var ProtoBufRuntimeABI = ProtoBufRuntimeMetaData.ABI

// ProtoBufRuntimeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProtoBufRuntimeMetaData.Bin instead.
var ProtoBufRuntimeBin = ProtoBufRuntimeMetaData.Bin

// DeployProtoBufRuntime deploys a new Ethereum contract, binding an instance of ProtoBufRuntime to it.
func DeployProtoBufRuntime(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProtoBufRuntime, error) {
	parsed, err := ProtoBufRuntimeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProtoBufRuntimeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProtoBufRuntime{ProtoBufRuntimeCaller: ProtoBufRuntimeCaller{contract: contract}, ProtoBufRuntimeTransactor: ProtoBufRuntimeTransactor{contract: contract}, ProtoBufRuntimeFilterer: ProtoBufRuntimeFilterer{contract: contract}}, nil
}

// ProtoBufRuntime is an auto generated Go binding around an Ethereum contract.
type ProtoBufRuntime struct {
	ProtoBufRuntimeCaller     // Read-only binding to the contract
	ProtoBufRuntimeTransactor // Write-only binding to the contract
	ProtoBufRuntimeFilterer   // Log filterer for contract events
}

// ProtoBufRuntimeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtoBufRuntimeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtoBufRuntimeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtoBufRuntimeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtoBufRuntimeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtoBufRuntimeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtoBufRuntimeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtoBufRuntimeSession struct {
	Contract     *ProtoBufRuntime  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProtoBufRuntimeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtoBufRuntimeCallerSession struct {
	Contract *ProtoBufRuntimeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ProtoBufRuntimeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtoBufRuntimeTransactorSession struct {
	Contract     *ProtoBufRuntimeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ProtoBufRuntimeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtoBufRuntimeRaw struct {
	Contract *ProtoBufRuntime // Generic contract binding to access the raw methods on
}

// ProtoBufRuntimeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtoBufRuntimeCallerRaw struct {
	Contract *ProtoBufRuntimeCaller // Generic read-only contract binding to access the raw methods on
}

// ProtoBufRuntimeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtoBufRuntimeTransactorRaw struct {
	Contract *ProtoBufRuntimeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtoBufRuntime creates a new instance of ProtoBufRuntime, bound to a specific deployed contract.
func NewProtoBufRuntime(address common.Address, backend bind.ContractBackend) (*ProtoBufRuntime, error) {
	contract, err := bindProtoBufRuntime(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProtoBufRuntime{ProtoBufRuntimeCaller: ProtoBufRuntimeCaller{contract: contract}, ProtoBufRuntimeTransactor: ProtoBufRuntimeTransactor{contract: contract}, ProtoBufRuntimeFilterer: ProtoBufRuntimeFilterer{contract: contract}}, nil
}

// NewProtoBufRuntimeCaller creates a new read-only instance of ProtoBufRuntime, bound to a specific deployed contract.
func NewProtoBufRuntimeCaller(address common.Address, caller bind.ContractCaller) (*ProtoBufRuntimeCaller, error) {
	contract, err := bindProtoBufRuntime(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtoBufRuntimeCaller{contract: contract}, nil
}

// NewProtoBufRuntimeTransactor creates a new write-only instance of ProtoBufRuntime, bound to a specific deployed contract.
func NewProtoBufRuntimeTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtoBufRuntimeTransactor, error) {
	contract, err := bindProtoBufRuntime(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtoBufRuntimeTransactor{contract: contract}, nil
}

// NewProtoBufRuntimeFilterer creates a new log filterer instance of ProtoBufRuntime, bound to a specific deployed contract.
func NewProtoBufRuntimeFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtoBufRuntimeFilterer, error) {
	contract, err := bindProtoBufRuntime(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtoBufRuntimeFilterer{contract: contract}, nil
}

// bindProtoBufRuntime binds a generic wrapper to an already deployed contract.
func bindProtoBufRuntime(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProtoBufRuntimeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtoBufRuntime *ProtoBufRuntimeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtoBufRuntime.Contract.ProtoBufRuntimeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtoBufRuntime *ProtoBufRuntimeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtoBufRuntime.Contract.ProtoBufRuntimeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtoBufRuntime *ProtoBufRuntimeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtoBufRuntime.Contract.ProtoBufRuntimeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtoBufRuntime *ProtoBufRuntimeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtoBufRuntime.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtoBufRuntime *ProtoBufRuntimeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtoBufRuntime.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtoBufRuntime *ProtoBufRuntimeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtoBufRuntime.Contract.contract.Transact(opts, method, params...)
}

// PublicKeyMetaData contains all meta data concerning the PublicKey contract.
var PublicKeyMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212207edb354d77e41df76cea902d0573ed29f3ed24f95b187a2882264a8ae559b5fb64736f6c63430008120033",
}

// PublicKeyABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicKeyMetaData.ABI instead.
var PublicKeyABI = PublicKeyMetaData.ABI

// PublicKeyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublicKeyMetaData.Bin instead.
var PublicKeyBin = PublicKeyMetaData.Bin

// DeployPublicKey deploys a new Ethereum contract, binding an instance of PublicKey to it.
func DeployPublicKey(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PublicKey, error) {
	parsed, err := PublicKeyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublicKeyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicKey{PublicKeyCaller: PublicKeyCaller{contract: contract}, PublicKeyTransactor: PublicKeyTransactor{contract: contract}, PublicKeyFilterer: PublicKeyFilterer{contract: contract}}, nil
}

// PublicKey is an auto generated Go binding around an Ethereum contract.
type PublicKey struct {
	PublicKeyCaller     // Read-only binding to the contract
	PublicKeyTransactor // Write-only binding to the contract
	PublicKeyFilterer   // Log filterer for contract events
}

// PublicKeyCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicKeyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicKeyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicKeyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicKeyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicKeyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicKeySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicKeySession struct {
	Contract     *PublicKey        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PublicKeyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicKeyCallerSession struct {
	Contract *PublicKeyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PublicKeyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicKeyTransactorSession struct {
	Contract     *PublicKeyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PublicKeyRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicKeyRaw struct {
	Contract *PublicKey // Generic contract binding to access the raw methods on
}

// PublicKeyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicKeyCallerRaw struct {
	Contract *PublicKeyCaller // Generic read-only contract binding to access the raw methods on
}

// PublicKeyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicKeyTransactorRaw struct {
	Contract *PublicKeyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicKey creates a new instance of PublicKey, bound to a specific deployed contract.
func NewPublicKey(address common.Address, backend bind.ContractBackend) (*PublicKey, error) {
	contract, err := bindPublicKey(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicKey{PublicKeyCaller: PublicKeyCaller{contract: contract}, PublicKeyTransactor: PublicKeyTransactor{contract: contract}, PublicKeyFilterer: PublicKeyFilterer{contract: contract}}, nil
}

// NewPublicKeyCaller creates a new read-only instance of PublicKey, bound to a specific deployed contract.
func NewPublicKeyCaller(address common.Address, caller bind.ContractCaller) (*PublicKeyCaller, error) {
	contract, err := bindPublicKey(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicKeyCaller{contract: contract}, nil
}

// NewPublicKeyTransactor creates a new write-only instance of PublicKey, bound to a specific deployed contract.
func NewPublicKeyTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicKeyTransactor, error) {
	contract, err := bindPublicKey(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicKeyTransactor{contract: contract}, nil
}

// NewPublicKeyFilterer creates a new log filterer instance of PublicKey, bound to a specific deployed contract.
func NewPublicKeyFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicKeyFilterer, error) {
	contract, err := bindPublicKey(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicKeyFilterer{contract: contract}, nil
}

// bindPublicKey binds a generic wrapper to an already deployed contract.
func bindPublicKey(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicKeyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicKey *PublicKeyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicKey.Contract.PublicKeyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicKey *PublicKeyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicKey.Contract.PublicKeyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicKey *PublicKeyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicKey.Contract.PublicKeyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicKey *PublicKeyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicKey.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicKey *PublicKeyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicKey.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicKey *PublicKeyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicKey.Contract.contract.Transact(opts, method, params...)
}

// RLPDecodeMetaData contains all meta data concerning the RLPDecode contract.
var RLPDecodeMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220181416ee13608c2ada7b8f53c21b8a0f4c6ccf21f449aa20a7f3defeb6bc155c64736f6c63430008120033",
}

// RLPDecodeABI is the input ABI used to generate the binding from.
// Deprecated: Use RLPDecodeMetaData.ABI instead.
var RLPDecodeABI = RLPDecodeMetaData.ABI

// RLPDecodeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RLPDecodeMetaData.Bin instead.
var RLPDecodeBin = RLPDecodeMetaData.Bin

// DeployRLPDecode deploys a new Ethereum contract, binding an instance of RLPDecode to it.
func DeployRLPDecode(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RLPDecode, error) {
	parsed, err := RLPDecodeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RLPDecodeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RLPDecode{RLPDecodeCaller: RLPDecodeCaller{contract: contract}, RLPDecodeTransactor: RLPDecodeTransactor{contract: contract}, RLPDecodeFilterer: RLPDecodeFilterer{contract: contract}}, nil
}

// RLPDecode is an auto generated Go binding around an Ethereum contract.
type RLPDecode struct {
	RLPDecodeCaller     // Read-only binding to the contract
	RLPDecodeTransactor // Write-only binding to the contract
	RLPDecodeFilterer   // Log filterer for contract events
}

// RLPDecodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPDecodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPDecodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPDecodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPDecodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPDecodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPDecodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPDecodeSession struct {
	Contract     *RLPDecode        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPDecodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPDecodeCallerSession struct {
	Contract *RLPDecodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RLPDecodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPDecodeTransactorSession struct {
	Contract     *RLPDecodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RLPDecodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPDecodeRaw struct {
	Contract *RLPDecode // Generic contract binding to access the raw methods on
}

// RLPDecodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPDecodeCallerRaw struct {
	Contract *RLPDecodeCaller // Generic read-only contract binding to access the raw methods on
}

// RLPDecodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPDecodeTransactorRaw struct {
	Contract *RLPDecodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLPDecode creates a new instance of RLPDecode, bound to a specific deployed contract.
func NewRLPDecode(address common.Address, backend bind.ContractBackend) (*RLPDecode, error) {
	contract, err := bindRLPDecode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLPDecode{RLPDecodeCaller: RLPDecodeCaller{contract: contract}, RLPDecodeTransactor: RLPDecodeTransactor{contract: contract}, RLPDecodeFilterer: RLPDecodeFilterer{contract: contract}}, nil
}

// NewRLPDecodeCaller creates a new read-only instance of RLPDecode, bound to a specific deployed contract.
func NewRLPDecodeCaller(address common.Address, caller bind.ContractCaller) (*RLPDecodeCaller, error) {
	contract, err := bindRLPDecode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPDecodeCaller{contract: contract}, nil
}

// NewRLPDecodeTransactor creates a new write-only instance of RLPDecode, bound to a specific deployed contract.
func NewRLPDecodeTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPDecodeTransactor, error) {
	contract, err := bindRLPDecode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPDecodeTransactor{contract: contract}, nil
}

// NewRLPDecodeFilterer creates a new log filterer instance of RLPDecode, bound to a specific deployed contract.
func NewRLPDecodeFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPDecodeFilterer, error) {
	contract, err := bindRLPDecode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPDecodeFilterer{contract: contract}, nil
}

// bindRLPDecode binds a generic wrapper to an already deployed contract.
func bindRLPDecode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RLPDecodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPDecode *RLPDecodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPDecode.Contract.RLPDecodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPDecode *RLPDecodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPDecode.Contract.RLPDecodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPDecode *RLPDecodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPDecode.Contract.RLPDecodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPDecode *RLPDecodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPDecode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPDecode *RLPDecodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPDecode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPDecode *RLPDecodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPDecode.Contract.contract.Transact(opts, method, params...)
}

// RLPEncodeMetaData contains all meta data concerning the RLPEncode contract.
var RLPEncodeMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212204ac8fd5371e0cdaef97f5bdc4211515adb967a8fa87a6c133657941b1805b67364736f6c63430008120033",
}

// RLPEncodeABI is the input ABI used to generate the binding from.
// Deprecated: Use RLPEncodeMetaData.ABI instead.
var RLPEncodeABI = RLPEncodeMetaData.ABI

// RLPEncodeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RLPEncodeMetaData.Bin instead.
var RLPEncodeBin = RLPEncodeMetaData.Bin

// DeployRLPEncode deploys a new Ethereum contract, binding an instance of RLPEncode to it.
func DeployRLPEncode(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RLPEncode, error) {
	parsed, err := RLPEncodeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RLPEncodeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RLPEncode{RLPEncodeCaller: RLPEncodeCaller{contract: contract}, RLPEncodeTransactor: RLPEncodeTransactor{contract: contract}, RLPEncodeFilterer: RLPEncodeFilterer{contract: contract}}, nil
}

// RLPEncode is an auto generated Go binding around an Ethereum contract.
type RLPEncode struct {
	RLPEncodeCaller     // Read-only binding to the contract
	RLPEncodeTransactor // Write-only binding to the contract
	RLPEncodeFilterer   // Log filterer for contract events
}

// RLPEncodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPEncodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPEncodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPEncodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPEncodeSession struct {
	Contract     *RLPEncode        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPEncodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPEncodeCallerSession struct {
	Contract *RLPEncodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RLPEncodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPEncodeTransactorSession struct {
	Contract     *RLPEncodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RLPEncodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPEncodeRaw struct {
	Contract *RLPEncode // Generic contract binding to access the raw methods on
}

// RLPEncodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPEncodeCallerRaw struct {
	Contract *RLPEncodeCaller // Generic read-only contract binding to access the raw methods on
}

// RLPEncodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPEncodeTransactorRaw struct {
	Contract *RLPEncodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLPEncode creates a new instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncode(address common.Address, backend bind.ContractBackend) (*RLPEncode, error) {
	contract, err := bindRLPEncode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLPEncode{RLPEncodeCaller: RLPEncodeCaller{contract: contract}, RLPEncodeTransactor: RLPEncodeTransactor{contract: contract}, RLPEncodeFilterer: RLPEncodeFilterer{contract: contract}}, nil
}

// NewRLPEncodeCaller creates a new read-only instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeCaller(address common.Address, caller bind.ContractCaller) (*RLPEncodeCaller, error) {
	contract, err := bindRLPEncode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeCaller{contract: contract}, nil
}

// NewRLPEncodeTransactor creates a new write-only instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPEncodeTransactor, error) {
	contract, err := bindRLPEncode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeTransactor{contract: contract}, nil
}

// NewRLPEncodeFilterer creates a new log filterer instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPEncodeFilterer, error) {
	contract, err := bindRLPEncode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeFilterer{contract: contract}, nil
}

// bindRLPEncode binds a generic wrapper to an already deployed contract.
func bindRLPEncode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RLPEncodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPEncode *RLPEncodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPEncode.Contract.RLPEncodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPEncode *RLPEncodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPEncode.Contract.RLPEncodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPEncode *RLPEncodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPEncode.Contract.RLPEncodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPEncode *RLPEncodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPEncode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPEncode *RLPEncodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPEncode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPEncode *RLPEncodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPEncode.Contract.contract.Transact(opts, method, params...)
}

// RLPWriterMetaData contains all meta data concerning the RLPWriter contract.
var RLPWriterMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220d0641fedb8f1fa3f374641d4b19445efda3b4e2928d737886add0a7bb017504c64736f6c63430008120033",
}

// RLPWriterABI is the input ABI used to generate the binding from.
// Deprecated: Use RLPWriterMetaData.ABI instead.
var RLPWriterABI = RLPWriterMetaData.ABI

// RLPWriterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RLPWriterMetaData.Bin instead.
var RLPWriterBin = RLPWriterMetaData.Bin

// DeployRLPWriter deploys a new Ethereum contract, binding an instance of RLPWriter to it.
func DeployRLPWriter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RLPWriter, error) {
	parsed, err := RLPWriterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RLPWriterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RLPWriter{RLPWriterCaller: RLPWriterCaller{contract: contract}, RLPWriterTransactor: RLPWriterTransactor{contract: contract}, RLPWriterFilterer: RLPWriterFilterer{contract: contract}}, nil
}

// RLPWriter is an auto generated Go binding around an Ethereum contract.
type RLPWriter struct {
	RLPWriterCaller     // Read-only binding to the contract
	RLPWriterTransactor // Write-only binding to the contract
	RLPWriterFilterer   // Log filterer for contract events
}

// RLPWriterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPWriterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPWriterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPWriterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPWriterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPWriterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPWriterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPWriterSession struct {
	Contract     *RLPWriter        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPWriterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPWriterCallerSession struct {
	Contract *RLPWriterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RLPWriterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPWriterTransactorSession struct {
	Contract     *RLPWriterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RLPWriterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPWriterRaw struct {
	Contract *RLPWriter // Generic contract binding to access the raw methods on
}

// RLPWriterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPWriterCallerRaw struct {
	Contract *RLPWriterCaller // Generic read-only contract binding to access the raw methods on
}

// RLPWriterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPWriterTransactorRaw struct {
	Contract *RLPWriterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLPWriter creates a new instance of RLPWriter, bound to a specific deployed contract.
func NewRLPWriter(address common.Address, backend bind.ContractBackend) (*RLPWriter, error) {
	contract, err := bindRLPWriter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLPWriter{RLPWriterCaller: RLPWriterCaller{contract: contract}, RLPWriterTransactor: RLPWriterTransactor{contract: contract}, RLPWriterFilterer: RLPWriterFilterer{contract: contract}}, nil
}

// NewRLPWriterCaller creates a new read-only instance of RLPWriter, bound to a specific deployed contract.
func NewRLPWriterCaller(address common.Address, caller bind.ContractCaller) (*RLPWriterCaller, error) {
	contract, err := bindRLPWriter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPWriterCaller{contract: contract}, nil
}

// NewRLPWriterTransactor creates a new write-only instance of RLPWriter, bound to a specific deployed contract.
func NewRLPWriterTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPWriterTransactor, error) {
	contract, err := bindRLPWriter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPWriterTransactor{contract: contract}, nil
}

// NewRLPWriterFilterer creates a new log filterer instance of RLPWriter, bound to a specific deployed contract.
func NewRLPWriterFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPWriterFilterer, error) {
	contract, err := bindRLPWriter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPWriterFilterer{contract: contract}, nil
}

// bindRLPWriter binds a generic wrapper to an already deployed contract.
func bindRLPWriter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RLPWriterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPWriter *RLPWriterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPWriter.Contract.RLPWriterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPWriter *RLPWriterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPWriter.Contract.RLPWriterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPWriter *RLPWriterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPWriter.Contract.RLPWriterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPWriter *RLPWriterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPWriter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPWriter *RLPWriterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPWriter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPWriter *RLPWriterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPWriter.Contract.contract.Transact(opts, method, params...)
}

// SafeCastMetaData contains all meta data concerning the SafeCast contract.
var SafeCastMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220a960fcbf0a68ac2f0e8176a42a57029ac757dfe2553a349b47a1b6d85891366364736f6c63430008120033",
}

// SafeCastABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeCastMetaData.ABI instead.
var SafeCastABI = SafeCastMetaData.ABI

// SafeCastBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeCastMetaData.Bin instead.
var SafeCastBin = SafeCastMetaData.Bin

// DeploySafeCast deploys a new Ethereum contract, binding an instance of SafeCast to it.
func DeploySafeCast(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeCast, error) {
	parsed, err := SafeCastMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeCastBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// SafeCast is an auto generated Go binding around an Ethereum contract.
type SafeCast struct {
	SafeCastCaller     // Read-only binding to the contract
	SafeCastTransactor // Write-only binding to the contract
	SafeCastFilterer   // Log filterer for contract events
}

// SafeCastCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeCastCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeCastTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeCastFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeCastSession struct {
	Contract     *SafeCast         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeCastCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeCastCallerSession struct {
	Contract *SafeCastCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeCastTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeCastTransactorSession struct {
	Contract     *SafeCastTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeCastRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeCastRaw struct {
	Contract *SafeCast // Generic contract binding to access the raw methods on
}

// SafeCastCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeCastCallerRaw struct {
	Contract *SafeCastCaller // Generic read-only contract binding to access the raw methods on
}

// SafeCastTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeCastTransactorRaw struct {
	Contract *SafeCastTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeCast creates a new instance of SafeCast, bound to a specific deployed contract.
func NewSafeCast(address common.Address, backend bind.ContractBackend) (*SafeCast, error) {
	contract, err := bindSafeCast(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// NewSafeCastCaller creates a new read-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastCaller(address common.Address, caller bind.ContractCaller) (*SafeCastCaller, error) {
	contract, err := bindSafeCast(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastCaller{contract: contract}, nil
}

// NewSafeCastTransactor creates a new write-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeCastTransactor, error) {
	contract, err := bindSafeCast(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastTransactor{contract: contract}, nil
}

// NewSafeCastFilterer creates a new log filterer instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeCastFilterer, error) {
	contract, err := bindSafeCast(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeCastFilterer{contract: contract}, nil
}

// bindSafeCast binds a generic wrapper to an already deployed contract.
func bindSafeCast(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeCastMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.SafeCastCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transact(opts, method, params...)
}

// SignedHeaderMetaData contains all meta data concerning the SignedHeader contract.
var SignedHeaderMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea264697066735822122026c10513e891753571eaa6addaa98f75b955dcf81bb589bf81501bc62f8c74d864736f6c63430008120033",
}

// SignedHeaderABI is the input ABI used to generate the binding from.
// Deprecated: Use SignedHeaderMetaData.ABI instead.
var SignedHeaderABI = SignedHeaderMetaData.ABI

// SignedHeaderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SignedHeaderMetaData.Bin instead.
var SignedHeaderBin = SignedHeaderMetaData.Bin

// DeploySignedHeader deploys a new Ethereum contract, binding an instance of SignedHeader to it.
func DeploySignedHeader(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SignedHeader, error) {
	parsed, err := SignedHeaderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SignedHeaderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SignedHeader{SignedHeaderCaller: SignedHeaderCaller{contract: contract}, SignedHeaderTransactor: SignedHeaderTransactor{contract: contract}, SignedHeaderFilterer: SignedHeaderFilterer{contract: contract}}, nil
}

// SignedHeader is an auto generated Go binding around an Ethereum contract.
type SignedHeader struct {
	SignedHeaderCaller     // Read-only binding to the contract
	SignedHeaderTransactor // Write-only binding to the contract
	SignedHeaderFilterer   // Log filterer for contract events
}

// SignedHeaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignedHeaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedHeaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignedHeaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedHeaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignedHeaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedHeaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignedHeaderSession struct {
	Contract     *SignedHeader     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignedHeaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignedHeaderCallerSession struct {
	Contract *SignedHeaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SignedHeaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignedHeaderTransactorSession struct {
	Contract     *SignedHeaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SignedHeaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignedHeaderRaw struct {
	Contract *SignedHeader // Generic contract binding to access the raw methods on
}

// SignedHeaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignedHeaderCallerRaw struct {
	Contract *SignedHeaderCaller // Generic read-only contract binding to access the raw methods on
}

// SignedHeaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignedHeaderTransactorRaw struct {
	Contract *SignedHeaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignedHeader creates a new instance of SignedHeader, bound to a specific deployed contract.
func NewSignedHeader(address common.Address, backend bind.ContractBackend) (*SignedHeader, error) {
	contract, err := bindSignedHeader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignedHeader{SignedHeaderCaller: SignedHeaderCaller{contract: contract}, SignedHeaderTransactor: SignedHeaderTransactor{contract: contract}, SignedHeaderFilterer: SignedHeaderFilterer{contract: contract}}, nil
}

// NewSignedHeaderCaller creates a new read-only instance of SignedHeader, bound to a specific deployed contract.
func NewSignedHeaderCaller(address common.Address, caller bind.ContractCaller) (*SignedHeaderCaller, error) {
	contract, err := bindSignedHeader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignedHeaderCaller{contract: contract}, nil
}

// NewSignedHeaderTransactor creates a new write-only instance of SignedHeader, bound to a specific deployed contract.
func NewSignedHeaderTransactor(address common.Address, transactor bind.ContractTransactor) (*SignedHeaderTransactor, error) {
	contract, err := bindSignedHeader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignedHeaderTransactor{contract: contract}, nil
}

// NewSignedHeaderFilterer creates a new log filterer instance of SignedHeader, bound to a specific deployed contract.
func NewSignedHeaderFilterer(address common.Address, filterer bind.ContractFilterer) (*SignedHeaderFilterer, error) {
	contract, err := bindSignedHeader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignedHeaderFilterer{contract: contract}, nil
}

// bindSignedHeader binds a generic wrapper to an already deployed contract.
func bindSignedHeader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SignedHeaderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignedHeader *SignedHeaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignedHeader.Contract.SignedHeaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignedHeader *SignedHeaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedHeader.Contract.SignedHeaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignedHeader *SignedHeaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignedHeader.Contract.SignedHeaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignedHeader *SignedHeaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignedHeader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignedHeader *SignedHeaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedHeader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignedHeader *SignedHeaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignedHeader.Contract.contract.Transact(opts, method, params...)
}

// SimpleValidatorMetaData contains all meta data concerning the SimpleValidator contract.
var SimpleValidatorMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212203c3e1fbe853e904c070f1b1c37b2cdf99ec8639cb2f54bfc9c1d6f6f9910ccd264736f6c63430008120033",
}

// SimpleValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleValidatorMetaData.ABI instead.
var SimpleValidatorABI = SimpleValidatorMetaData.ABI

// SimpleValidatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleValidatorMetaData.Bin instead.
var SimpleValidatorBin = SimpleValidatorMetaData.Bin

// DeploySimpleValidator deploys a new Ethereum contract, binding an instance of SimpleValidator to it.
func DeploySimpleValidator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimpleValidator, error) {
	parsed, err := SimpleValidatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleValidatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleValidator{SimpleValidatorCaller: SimpleValidatorCaller{contract: contract}, SimpleValidatorTransactor: SimpleValidatorTransactor{contract: contract}, SimpleValidatorFilterer: SimpleValidatorFilterer{contract: contract}}, nil
}

// SimpleValidator is an auto generated Go binding around an Ethereum contract.
type SimpleValidator struct {
	SimpleValidatorCaller     // Read-only binding to the contract
	SimpleValidatorTransactor // Write-only binding to the contract
	SimpleValidatorFilterer   // Log filterer for contract events
}

// SimpleValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleValidatorSession struct {
	Contract     *SimpleValidator  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleValidatorCallerSession struct {
	Contract *SimpleValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SimpleValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleValidatorTransactorSession struct {
	Contract     *SimpleValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SimpleValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleValidatorRaw struct {
	Contract *SimpleValidator // Generic contract binding to access the raw methods on
}

// SimpleValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleValidatorCallerRaw struct {
	Contract *SimpleValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleValidatorTransactorRaw struct {
	Contract *SimpleValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleValidator creates a new instance of SimpleValidator, bound to a specific deployed contract.
func NewSimpleValidator(address common.Address, backend bind.ContractBackend) (*SimpleValidator, error) {
	contract, err := bindSimpleValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleValidator{SimpleValidatorCaller: SimpleValidatorCaller{contract: contract}, SimpleValidatorTransactor: SimpleValidatorTransactor{contract: contract}, SimpleValidatorFilterer: SimpleValidatorFilterer{contract: contract}}, nil
}

// NewSimpleValidatorCaller creates a new read-only instance of SimpleValidator, bound to a specific deployed contract.
func NewSimpleValidatorCaller(address common.Address, caller bind.ContractCaller) (*SimpleValidatorCaller, error) {
	contract, err := bindSimpleValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleValidatorCaller{contract: contract}, nil
}

// NewSimpleValidatorTransactor creates a new write-only instance of SimpleValidator, bound to a specific deployed contract.
func NewSimpleValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleValidatorTransactor, error) {
	contract, err := bindSimpleValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleValidatorTransactor{contract: contract}, nil
}

// NewSimpleValidatorFilterer creates a new log filterer instance of SimpleValidator, bound to a specific deployed contract.
func NewSimpleValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleValidatorFilterer, error) {
	contract, err := bindSimpleValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleValidatorFilterer{contract: contract}, nil
}

// bindSimpleValidator binds a generic wrapper to an already deployed contract.
func bindSimpleValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleValidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleValidator *SimpleValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleValidator.Contract.SimpleValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleValidator *SimpleValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleValidator.Contract.SimpleValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleValidator *SimpleValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleValidator.Contract.SimpleValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleValidator *SimpleValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleValidator *SimpleValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleValidator *SimpleValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleValidator.Contract.contract.Transact(opts, method, params...)
}

// SystemMetaData contains all meta data concerning the System contract.
var SystemMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CODE_OK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_FAIL_DECODE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bscChainID\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bscValidatorSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossChain\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_bscChainID\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bscValidatorSet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tmLightClient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_crossChain\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"setRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tmLightClient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461005b5760008054336001600160a01b0319821681178355916001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a361059890816100618239f35b600080fdfe6080604081815260048036101561001557600080fd5b600092833560e01c908163018042c814610379575080630bee7a671461035d5780634917db6614610335578063493279b11461030f5780636548e9bc146102ca578063715018a61461026d5780638406c0791461024357806387aa9376146102275780638da5cb5b14610201578063ab51bb96146101e6578063ac75e014146101be578063e89a9501146101925763f2fde38b146100b257600080fd5b3461018e57602036600319011261018e576100cb6104fc565b908354906001600160a01b03808316936100e6338614610517565b169384156101255750506001600160a01b031916821783557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b906020608492519162461bcd60e51b8352820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152fd5b8280fd5b5050346101ba57816003193601126101ba576020906001600160a01b03600154169051908152f35b5080fd5b5050346101ba57816003193601126101ba576020906001600160a01b03600254169051908152f35b5050346101ba57816003193601126101ba5751908152602090f35b5050346101ba57816003193601126101ba576001600160a01b0360209254169051908152f35b5050346101ba57816003193601126101ba576020905160088152f35b50913461026a578060031936011261026a57506001600160a01b0360209254169051908152f35b80fd5b833461026a578060031936011261026a578080546001600160a01b03196001600160a01b038216916102a0338414610517565b1682557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b8382346101ba5760203660031901126101ba576102e56104fc565b6001600160a01b03906102fc828554163314610517565b166001600160a01b031982541617905580f35b5050346101ba57816003193601126101ba5760209061ffff60035460a01c169051908152f35b5050346101ba57816003193601126101ba576020906001600160a01b03600354169051908152f35b5050346101ba57816003193601126101ba576020905160648152f35b9250503461018e5760a036600319011261018e5780359061ffff821682036104f857602435916001600160a01b03928381168091036104f457604435908482168092036104f057606435918583168093036104ec57608435958087168097036104e85788549760ff8960a01c166104a7575093740100000000000000000000000000000000000000009795937fffffffffffffffffffff000000000000000000000000000000000000000000009375ffff00000000000000000000000000000000000000009361046f7fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff9a9833908b1614610517565b600354966001600160a01b03199283825416179055816001541617600155600254161760025560a01b16911617176003551617815580f35b62461bcd60e51b8152602087820152601360248201527f616c726561647920696e697469616c697a6564000000000000000000000000006044820152606490fd5b8880fd5b8780fd5b8680fd5b8580fd5b8380fd5b600435906001600160a01b038216820361051257565b600080fd5b1561051e57565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fdfea264697066735822122042ebeec571c6a7ab0b7efe7aab3d5ca94bed1005ed7785735b7406a0f8dbbacc64736f6c63430008120033",
}

// SystemABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemMetaData.ABI instead.
var SystemABI = SystemMetaData.ABI

// SystemBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SystemMetaData.Bin instead.
var SystemBin = SystemMetaData.Bin

// DeploySystem deploys a new Ethereum contract, binding an instance of System to it.
func DeploySystem(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *System, error) {
	parsed, err := SystemMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SystemBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &System{SystemCaller: SystemCaller{contract: contract}, SystemTransactor: SystemTransactor{contract: contract}, SystemFilterer: SystemFilterer{contract: contract}}, nil
}

// System is an auto generated Go binding around an Ethereum contract.
type System struct {
	SystemCaller     // Read-only binding to the contract
	SystemTransactor // Write-only binding to the contract
	SystemFilterer   // Log filterer for contract events
}

// SystemCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemSession struct {
	Contract     *System           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemCallerSession struct {
	Contract *SystemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SystemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemTransactorSession struct {
	Contract     *SystemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemRaw struct {
	Contract *System // Generic contract binding to access the raw methods on
}

// SystemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemCallerRaw struct {
	Contract *SystemCaller // Generic read-only contract binding to access the raw methods on
}

// SystemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemTransactorRaw struct {
	Contract *SystemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystem creates a new instance of System, bound to a specific deployed contract.
func NewSystem(address common.Address, backend bind.ContractBackend) (*System, error) {
	contract, err := bindSystem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &System{SystemCaller: SystemCaller{contract: contract}, SystemTransactor: SystemTransactor{contract: contract}, SystemFilterer: SystemFilterer{contract: contract}}, nil
}

// NewSystemCaller creates a new read-only instance of System, bound to a specific deployed contract.
func NewSystemCaller(address common.Address, caller bind.ContractCaller) (*SystemCaller, error) {
	contract, err := bindSystem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemCaller{contract: contract}, nil
}

// NewSystemTransactor creates a new write-only instance of System, bound to a specific deployed contract.
func NewSystemTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemTransactor, error) {
	contract, err := bindSystem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemTransactor{contract: contract}, nil
}

// NewSystemFilterer creates a new log filterer instance of System, bound to a specific deployed contract.
func NewSystemFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemFilterer, error) {
	contract, err := bindSystem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemFilterer{contract: contract}, nil
}

// bindSystem binds a generic wrapper to an already deployed contract.
func bindSystem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SystemMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_System *SystemRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _System.Contract.SystemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_System *SystemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _System.Contract.SystemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_System *SystemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _System.Contract.SystemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_System *SystemCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _System.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_System *SystemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _System.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_System *SystemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _System.Contract.contract.Transact(opts, method, params...)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_System *SystemCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "CODE_OK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_System *SystemSession) CODEOK() (uint32, error) {
	return _System.Contract.CODEOK(&_System.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_System *SystemCallerSession) CODEOK() (uint32, error) {
	return _System.Contract.CODEOK(&_System.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_System *SystemCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "ERROR_FAIL_DECODE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_System *SystemSession) ERRORFAILDECODE() (uint32, error) {
	return _System.Contract.ERRORFAILDECODE(&_System.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_System *SystemCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _System.Contract.ERRORFAILDECODE(&_System.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x87aa9376.
//
// Solidity: function STAKING_CHANNEL_ID() view returns(uint8)
func (_System *SystemCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "STAKING_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x87aa9376.
//
// Solidity: function STAKING_CHANNEL_ID() view returns(uint8)
func (_System *SystemSession) STAKINGCHANNELID() (uint8, error) {
	return _System.Contract.STAKINGCHANNELID(&_System.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x87aa9376.
//
// Solidity: function STAKING_CHANNEL_ID() view returns(uint8)
func (_System *SystemCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _System.Contract.STAKINGCHANNELID(&_System.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_System *SystemCaller) BscChainID(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "bscChainID")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_System *SystemSession) BscChainID() (uint16, error) {
	return _System.Contract.BscChainID(&_System.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_System *SystemCallerSession) BscChainID() (uint16, error) {
	return _System.Contract.BscChainID(&_System.CallOpts)
}

// BscValidatorSet is a free data retrieval call binding the contract method 0xe89a9501.
//
// Solidity: function bscValidatorSet() view returns(address)
func (_System *SystemCaller) BscValidatorSet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "bscValidatorSet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BscValidatorSet is a free data retrieval call binding the contract method 0xe89a9501.
//
// Solidity: function bscValidatorSet() view returns(address)
func (_System *SystemSession) BscValidatorSet() (common.Address, error) {
	return _System.Contract.BscValidatorSet(&_System.CallOpts)
}

// BscValidatorSet is a free data retrieval call binding the contract method 0xe89a9501.
//
// Solidity: function bscValidatorSet() view returns(address)
func (_System *SystemCallerSession) BscValidatorSet() (common.Address, error) {
	return _System.Contract.BscValidatorSet(&_System.CallOpts)
}

// CrossChain is a free data retrieval call binding the contract method 0x4917db66.
//
// Solidity: function crossChain() view returns(address)
func (_System *SystemCaller) CrossChain(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "crossChain")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossChain is a free data retrieval call binding the contract method 0x4917db66.
//
// Solidity: function crossChain() view returns(address)
func (_System *SystemSession) CrossChain() (common.Address, error) {
	return _System.Contract.CrossChain(&_System.CallOpts)
}

// CrossChain is a free data retrieval call binding the contract method 0x4917db66.
//
// Solidity: function crossChain() view returns(address)
func (_System *SystemCallerSession) CrossChain() (common.Address, error) {
	return _System.Contract.CrossChain(&_System.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_System *SystemCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_System *SystemSession) Owner() (common.Address, error) {
	return _System.Contract.Owner(&_System.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_System *SystemCallerSession) Owner() (common.Address, error) {
	return _System.Contract.Owner(&_System.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_System *SystemCaller) Relayer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "relayer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_System *SystemSession) Relayer() (common.Address, error) {
	return _System.Contract.Relayer(&_System.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_System *SystemCallerSession) Relayer() (common.Address, error) {
	return _System.Contract.Relayer(&_System.CallOpts)
}

// TmLightClient is a free data retrieval call binding the contract method 0xac75e014.
//
// Solidity: function tmLightClient() view returns(address)
func (_System *SystemCaller) TmLightClient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "tmLightClient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TmLightClient is a free data retrieval call binding the contract method 0xac75e014.
//
// Solidity: function tmLightClient() view returns(address)
func (_System *SystemSession) TmLightClient() (common.Address, error) {
	return _System.Contract.TmLightClient(&_System.CallOpts)
}

// TmLightClient is a free data retrieval call binding the contract method 0xac75e014.
//
// Solidity: function tmLightClient() view returns(address)
func (_System *SystemCallerSession) TmLightClient() (common.Address, error) {
	return _System.Contract.TmLightClient(&_System.CallOpts)
}

// Init is a paid mutator transaction binding the contract method 0x018042c8.
//
// Solidity: function init(uint16 _bscChainID, address _relayer, address _bscValidatorSet, address _tmLightClient, address _crossChain) returns()
func (_System *SystemTransactor) Init(opts *bind.TransactOpts, _bscChainID uint16, _relayer common.Address, _bscValidatorSet common.Address, _tmLightClient common.Address, _crossChain common.Address) (*types.Transaction, error) {
	return _System.contract.Transact(opts, "init", _bscChainID, _relayer, _bscValidatorSet, _tmLightClient, _crossChain)
}

// Init is a paid mutator transaction binding the contract method 0x018042c8.
//
// Solidity: function init(uint16 _bscChainID, address _relayer, address _bscValidatorSet, address _tmLightClient, address _crossChain) returns()
func (_System *SystemSession) Init(_bscChainID uint16, _relayer common.Address, _bscValidatorSet common.Address, _tmLightClient common.Address, _crossChain common.Address) (*types.Transaction, error) {
	return _System.Contract.Init(&_System.TransactOpts, _bscChainID, _relayer, _bscValidatorSet, _tmLightClient, _crossChain)
}

// Init is a paid mutator transaction binding the contract method 0x018042c8.
//
// Solidity: function init(uint16 _bscChainID, address _relayer, address _bscValidatorSet, address _tmLightClient, address _crossChain) returns()
func (_System *SystemTransactorSession) Init(_bscChainID uint16, _relayer common.Address, _bscValidatorSet common.Address, _tmLightClient common.Address, _crossChain common.Address) (*types.Transaction, error) {
	return _System.Contract.Init(&_System.TransactOpts, _bscChainID, _relayer, _bscValidatorSet, _tmLightClient, _crossChain)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_System *SystemTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _System.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_System *SystemSession) RenounceOwnership() (*types.Transaction, error) {
	return _System.Contract.RenounceOwnership(&_System.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_System *SystemTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _System.Contract.RenounceOwnership(&_System.TransactOpts)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _relayer) returns()
func (_System *SystemTransactor) SetRelayer(opts *bind.TransactOpts, _relayer common.Address) (*types.Transaction, error) {
	return _System.contract.Transact(opts, "setRelayer", _relayer)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _relayer) returns()
func (_System *SystemSession) SetRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _System.Contract.SetRelayer(&_System.TransactOpts, _relayer)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _relayer) returns()
func (_System *SystemTransactorSession) SetRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _System.Contract.SetRelayer(&_System.TransactOpts, _relayer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_System *SystemTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _System.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_System *SystemSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _System.Contract.TransferOwnership(&_System.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_System *SystemTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _System.Contract.TransferOwnership(&_System.TransactOpts, newOwner)
}

// SystemOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the System contract.
type SystemOwnershipTransferredIterator struct {
	Event *SystemOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SystemOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SystemOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SystemOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemOwnershipTransferred represents a OwnershipTransferred event raised by the System contract.
type SystemOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_System *SystemFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SystemOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _System.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SystemOwnershipTransferredIterator{contract: _System.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_System *SystemFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SystemOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _System.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemOwnershipTransferred)
				if err := _System.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_System *SystemFilterer) ParseOwnershipTransferred(log types.Log) (*SystemOwnershipTransferred, error) {
	event := new(SystemOwnershipTransferred)
	if err := _System.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TENDERMINTLIGHTPROTOGLOBALENUMSMetaData contains all meta data concerning the TENDERMINTLIGHTPROTOGLOBALENUMS contract.
var TENDERMINTLIGHTPROTOGLOBALENUMSMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220c908bcf71aaa0c268bc98e2a889e034cf4135cb478091f8907793e17e25d429f64736f6c63430008120033",
}

// TENDERMINTLIGHTPROTOGLOBALENUMSABI is the input ABI used to generate the binding from.
// Deprecated: Use TENDERMINTLIGHTPROTOGLOBALENUMSMetaData.ABI instead.
var TENDERMINTLIGHTPROTOGLOBALENUMSABI = TENDERMINTLIGHTPROTOGLOBALENUMSMetaData.ABI

// TENDERMINTLIGHTPROTOGLOBALENUMSBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TENDERMINTLIGHTPROTOGLOBALENUMSMetaData.Bin instead.
var TENDERMINTLIGHTPROTOGLOBALENUMSBin = TENDERMINTLIGHTPROTOGLOBALENUMSMetaData.Bin

// DeployTENDERMINTLIGHTPROTOGLOBALENUMS deploys a new Ethereum contract, binding an instance of TENDERMINTLIGHTPROTOGLOBALENUMS to it.
func DeployTENDERMINTLIGHTPROTOGLOBALENUMS(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TENDERMINTLIGHTPROTOGLOBALENUMS, error) {
	parsed, err := TENDERMINTLIGHTPROTOGLOBALENUMSMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TENDERMINTLIGHTPROTOGLOBALENUMSBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TENDERMINTLIGHTPROTOGLOBALENUMS{TENDERMINTLIGHTPROTOGLOBALENUMSCaller: TENDERMINTLIGHTPROTOGLOBALENUMSCaller{contract: contract}, TENDERMINTLIGHTPROTOGLOBALENUMSTransactor: TENDERMINTLIGHTPROTOGLOBALENUMSTransactor{contract: contract}, TENDERMINTLIGHTPROTOGLOBALENUMSFilterer: TENDERMINTLIGHTPROTOGLOBALENUMSFilterer{contract: contract}}, nil
}

// TENDERMINTLIGHTPROTOGLOBALENUMS is an auto generated Go binding around an Ethereum contract.
type TENDERMINTLIGHTPROTOGLOBALENUMS struct {
	TENDERMINTLIGHTPROTOGLOBALENUMSCaller     // Read-only binding to the contract
	TENDERMINTLIGHTPROTOGLOBALENUMSTransactor // Write-only binding to the contract
	TENDERMINTLIGHTPROTOGLOBALENUMSFilterer   // Log filterer for contract events
}

// TENDERMINTLIGHTPROTOGLOBALENUMSCaller is an auto generated read-only Go binding around an Ethereum contract.
type TENDERMINTLIGHTPROTOGLOBALENUMSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TENDERMINTLIGHTPROTOGLOBALENUMSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TENDERMINTLIGHTPROTOGLOBALENUMSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TENDERMINTLIGHTPROTOGLOBALENUMSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TENDERMINTLIGHTPROTOGLOBALENUMSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TENDERMINTLIGHTPROTOGLOBALENUMSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TENDERMINTLIGHTPROTOGLOBALENUMSSession struct {
	Contract     *TENDERMINTLIGHTPROTOGLOBALENUMS // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// TENDERMINTLIGHTPROTOGLOBALENUMSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TENDERMINTLIGHTPROTOGLOBALENUMSCallerSession struct {
	Contract *TENDERMINTLIGHTPROTOGLOBALENUMSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// TENDERMINTLIGHTPROTOGLOBALENUMSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TENDERMINTLIGHTPROTOGLOBALENUMSTransactorSession struct {
	Contract     *TENDERMINTLIGHTPROTOGLOBALENUMSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// TENDERMINTLIGHTPROTOGLOBALENUMSRaw is an auto generated low-level Go binding around an Ethereum contract.
type TENDERMINTLIGHTPROTOGLOBALENUMSRaw struct {
	Contract *TENDERMINTLIGHTPROTOGLOBALENUMS // Generic contract binding to access the raw methods on
}

// TENDERMINTLIGHTPROTOGLOBALENUMSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TENDERMINTLIGHTPROTOGLOBALENUMSCallerRaw struct {
	Contract *TENDERMINTLIGHTPROTOGLOBALENUMSCaller // Generic read-only contract binding to access the raw methods on
}

// TENDERMINTLIGHTPROTOGLOBALENUMSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TENDERMINTLIGHTPROTOGLOBALENUMSTransactorRaw struct {
	Contract *TENDERMINTLIGHTPROTOGLOBALENUMSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTENDERMINTLIGHTPROTOGLOBALENUMS creates a new instance of TENDERMINTLIGHTPROTOGLOBALENUMS, bound to a specific deployed contract.
func NewTENDERMINTLIGHTPROTOGLOBALENUMS(address common.Address, backend bind.ContractBackend) (*TENDERMINTLIGHTPROTOGLOBALENUMS, error) {
	contract, err := bindTENDERMINTLIGHTPROTOGLOBALENUMS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TENDERMINTLIGHTPROTOGLOBALENUMS{TENDERMINTLIGHTPROTOGLOBALENUMSCaller: TENDERMINTLIGHTPROTOGLOBALENUMSCaller{contract: contract}, TENDERMINTLIGHTPROTOGLOBALENUMSTransactor: TENDERMINTLIGHTPROTOGLOBALENUMSTransactor{contract: contract}, TENDERMINTLIGHTPROTOGLOBALENUMSFilterer: TENDERMINTLIGHTPROTOGLOBALENUMSFilterer{contract: contract}}, nil
}

// NewTENDERMINTLIGHTPROTOGLOBALENUMSCaller creates a new read-only instance of TENDERMINTLIGHTPROTOGLOBALENUMS, bound to a specific deployed contract.
func NewTENDERMINTLIGHTPROTOGLOBALENUMSCaller(address common.Address, caller bind.ContractCaller) (*TENDERMINTLIGHTPROTOGLOBALENUMSCaller, error) {
	contract, err := bindTENDERMINTLIGHTPROTOGLOBALENUMS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TENDERMINTLIGHTPROTOGLOBALENUMSCaller{contract: contract}, nil
}

// NewTENDERMINTLIGHTPROTOGLOBALENUMSTransactor creates a new write-only instance of TENDERMINTLIGHTPROTOGLOBALENUMS, bound to a specific deployed contract.
func NewTENDERMINTLIGHTPROTOGLOBALENUMSTransactor(address common.Address, transactor bind.ContractTransactor) (*TENDERMINTLIGHTPROTOGLOBALENUMSTransactor, error) {
	contract, err := bindTENDERMINTLIGHTPROTOGLOBALENUMS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TENDERMINTLIGHTPROTOGLOBALENUMSTransactor{contract: contract}, nil
}

// NewTENDERMINTLIGHTPROTOGLOBALENUMSFilterer creates a new log filterer instance of TENDERMINTLIGHTPROTOGLOBALENUMS, bound to a specific deployed contract.
func NewTENDERMINTLIGHTPROTOGLOBALENUMSFilterer(address common.Address, filterer bind.ContractFilterer) (*TENDERMINTLIGHTPROTOGLOBALENUMSFilterer, error) {
	contract, err := bindTENDERMINTLIGHTPROTOGLOBALENUMS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TENDERMINTLIGHTPROTOGLOBALENUMSFilterer{contract: contract}, nil
}

// bindTENDERMINTLIGHTPROTOGLOBALENUMS binds a generic wrapper to an already deployed contract.
func bindTENDERMINTLIGHTPROTOGLOBALENUMS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TENDERMINTLIGHTPROTOGLOBALENUMSMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TENDERMINTLIGHTPROTOGLOBALENUMS *TENDERMINTLIGHTPROTOGLOBALENUMSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TENDERMINTLIGHTPROTOGLOBALENUMS.Contract.TENDERMINTLIGHTPROTOGLOBALENUMSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TENDERMINTLIGHTPROTOGLOBALENUMS *TENDERMINTLIGHTPROTOGLOBALENUMSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TENDERMINTLIGHTPROTOGLOBALENUMS.Contract.TENDERMINTLIGHTPROTOGLOBALENUMSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TENDERMINTLIGHTPROTOGLOBALENUMS *TENDERMINTLIGHTPROTOGLOBALENUMSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TENDERMINTLIGHTPROTOGLOBALENUMS.Contract.TENDERMINTLIGHTPROTOGLOBALENUMSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TENDERMINTLIGHTPROTOGLOBALENUMS *TENDERMINTLIGHTPROTOGLOBALENUMSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TENDERMINTLIGHTPROTOGLOBALENUMS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TENDERMINTLIGHTPROTOGLOBALENUMS *TENDERMINTLIGHTPROTOGLOBALENUMSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TENDERMINTLIGHTPROTOGLOBALENUMS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TENDERMINTLIGHTPROTOGLOBALENUMS *TENDERMINTLIGHTPROTOGLOBALENUMSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TENDERMINTLIGHTPROTOGLOBALENUMS.Contract.contract.Transact(opts, method, params...)
}

// TendermintMetaData contains all meta data concerning the Tendermint contract.
var TendermintMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea264697066735822122040d151af8f5a46256f15a08bc7864e7ecd0d340b0f11bd129020576a6c48db9564736f6c63430008120033",
}

// TendermintABI is the input ABI used to generate the binding from.
// Deprecated: Use TendermintMetaData.ABI instead.
var TendermintABI = TendermintMetaData.ABI

// TendermintBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TendermintMetaData.Bin instead.
var TendermintBin = TendermintMetaData.Bin

// DeployTendermint deploys a new Ethereum contract, binding an instance of Tendermint to it.
func DeployTendermint(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tendermint, error) {
	parsed, err := TendermintMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TendermintBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tendermint{TendermintCaller: TendermintCaller{contract: contract}, TendermintTransactor: TendermintTransactor{contract: contract}, TendermintFilterer: TendermintFilterer{contract: contract}}, nil
}

// Tendermint is an auto generated Go binding around an Ethereum contract.
type Tendermint struct {
	TendermintCaller     // Read-only binding to the contract
	TendermintTransactor // Write-only binding to the contract
	TendermintFilterer   // Log filterer for contract events
}

// TendermintCaller is an auto generated read-only Go binding around an Ethereum contract.
type TendermintCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TendermintTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TendermintTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TendermintFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TendermintFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TendermintSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TendermintSession struct {
	Contract     *Tendermint       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TendermintCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TendermintCallerSession struct {
	Contract *TendermintCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TendermintTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TendermintTransactorSession struct {
	Contract     *TendermintTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TendermintRaw is an auto generated low-level Go binding around an Ethereum contract.
type TendermintRaw struct {
	Contract *Tendermint // Generic contract binding to access the raw methods on
}

// TendermintCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TendermintCallerRaw struct {
	Contract *TendermintCaller // Generic read-only contract binding to access the raw methods on
}

// TendermintTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TendermintTransactorRaw struct {
	Contract *TendermintTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTendermint creates a new instance of Tendermint, bound to a specific deployed contract.
func NewTendermint(address common.Address, backend bind.ContractBackend) (*Tendermint, error) {
	contract, err := bindTendermint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tendermint{TendermintCaller: TendermintCaller{contract: contract}, TendermintTransactor: TendermintTransactor{contract: contract}, TendermintFilterer: TendermintFilterer{contract: contract}}, nil
}

// NewTendermintCaller creates a new read-only instance of Tendermint, bound to a specific deployed contract.
func NewTendermintCaller(address common.Address, caller bind.ContractCaller) (*TendermintCaller, error) {
	contract, err := bindTendermint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TendermintCaller{contract: contract}, nil
}

// NewTendermintTransactor creates a new write-only instance of Tendermint, bound to a specific deployed contract.
func NewTendermintTransactor(address common.Address, transactor bind.ContractTransactor) (*TendermintTransactor, error) {
	contract, err := bindTendermint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TendermintTransactor{contract: contract}, nil
}

// NewTendermintFilterer creates a new log filterer instance of Tendermint, bound to a specific deployed contract.
func NewTendermintFilterer(address common.Address, filterer bind.ContractFilterer) (*TendermintFilterer, error) {
	contract, err := bindTendermint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TendermintFilterer{contract: contract}, nil
}

// bindTendermint binds a generic wrapper to an already deployed contract.
func bindTendermint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TendermintMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tendermint *TendermintRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tendermint.Contract.TendermintCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tendermint *TendermintRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tendermint.Contract.TendermintTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tendermint *TendermintRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tendermint.Contract.TendermintTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tendermint *TendermintCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tendermint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tendermint *TendermintTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tendermint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tendermint *TendermintTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tendermint.Contract.contract.Transact(opts, method, params...)
}

// TendermintHelperMetaData contains all meta data concerning the TendermintHelper contract.
var TendermintHelperMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea264697066735822122066c71c2fb8c2728501c256e352b883dc7821778bd5b79c5a4716c4238637592b64736f6c63430008120033",
}

// TendermintHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use TendermintHelperMetaData.ABI instead.
var TendermintHelperABI = TendermintHelperMetaData.ABI

// TendermintHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TendermintHelperMetaData.Bin instead.
var TendermintHelperBin = TendermintHelperMetaData.Bin

// DeployTendermintHelper deploys a new Ethereum contract, binding an instance of TendermintHelper to it.
func DeployTendermintHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TendermintHelper, error) {
	parsed, err := TendermintHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TendermintHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TendermintHelper{TendermintHelperCaller: TendermintHelperCaller{contract: contract}, TendermintHelperTransactor: TendermintHelperTransactor{contract: contract}, TendermintHelperFilterer: TendermintHelperFilterer{contract: contract}}, nil
}

// TendermintHelper is an auto generated Go binding around an Ethereum contract.
type TendermintHelper struct {
	TendermintHelperCaller     // Read-only binding to the contract
	TendermintHelperTransactor // Write-only binding to the contract
	TendermintHelperFilterer   // Log filterer for contract events
}

// TendermintHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type TendermintHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TendermintHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TendermintHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TendermintHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TendermintHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TendermintHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TendermintHelperSession struct {
	Contract     *TendermintHelper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TendermintHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TendermintHelperCallerSession struct {
	Contract *TendermintHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TendermintHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TendermintHelperTransactorSession struct {
	Contract     *TendermintHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TendermintHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type TendermintHelperRaw struct {
	Contract *TendermintHelper // Generic contract binding to access the raw methods on
}

// TendermintHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TendermintHelperCallerRaw struct {
	Contract *TendermintHelperCaller // Generic read-only contract binding to access the raw methods on
}

// TendermintHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TendermintHelperTransactorRaw struct {
	Contract *TendermintHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTendermintHelper creates a new instance of TendermintHelper, bound to a specific deployed contract.
func NewTendermintHelper(address common.Address, backend bind.ContractBackend) (*TendermintHelper, error) {
	contract, err := bindTendermintHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TendermintHelper{TendermintHelperCaller: TendermintHelperCaller{contract: contract}, TendermintHelperTransactor: TendermintHelperTransactor{contract: contract}, TendermintHelperFilterer: TendermintHelperFilterer{contract: contract}}, nil
}

// NewTendermintHelperCaller creates a new read-only instance of TendermintHelper, bound to a specific deployed contract.
func NewTendermintHelperCaller(address common.Address, caller bind.ContractCaller) (*TendermintHelperCaller, error) {
	contract, err := bindTendermintHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TendermintHelperCaller{contract: contract}, nil
}

// NewTendermintHelperTransactor creates a new write-only instance of TendermintHelper, bound to a specific deployed contract.
func NewTendermintHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*TendermintHelperTransactor, error) {
	contract, err := bindTendermintHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TendermintHelperTransactor{contract: contract}, nil
}

// NewTendermintHelperFilterer creates a new log filterer instance of TendermintHelper, bound to a specific deployed contract.
func NewTendermintHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*TendermintHelperFilterer, error) {
	contract, err := bindTendermintHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TendermintHelperFilterer{contract: contract}, nil
}

// bindTendermintHelper binds a generic wrapper to an already deployed contract.
func bindTendermintHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TendermintHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TendermintHelper *TendermintHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TendermintHelper.Contract.TendermintHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TendermintHelper *TendermintHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TendermintHelper.Contract.TendermintHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TendermintHelper *TendermintHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TendermintHelper.Contract.TendermintHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TendermintHelper *TendermintHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TendermintHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TendermintHelper *TendermintHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TendermintHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TendermintHelper *TendermintHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TendermintHelper.Contract.contract.Transact(opts, method, params...)
}

// TendermintLightClientMetaData contains all meta data concerning the TendermintLightClient contract.
var TendermintLightClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ed25519Verifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"initialHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"appHash\",\"type\":\"bytes32\"}],\"name\":\"ConsensusStateInit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"appHash\",\"type\":\"bytes32\"}],\"name\":\"ConsensusStateSynced\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"consensusStates\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"Seconds\",\"type\":\"int64\"},{\"internalType\":\"int32\",\"name\":\"nanos\",\"type\":\"int32\"}],\"internalType\":\"structTimestamp.Data\",\"name\":\"timestamp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleRoot.Data\",\"name\":\"root\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"next_validators_hash\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getAppHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_system\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_initHeader\",\"type\":\"bytes\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"isHeaderSynced\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"proofA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"proofB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"proofC\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"proofCommit\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"proofCommitPub\",\"type\":\"uint256\"}],\"name\":\"syncTendermintHeader\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"synced\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080346200007a57601f62005b6038819003918201601f19168301916001600160401b038311848410176200007f578084926020946040528339810103126200007a57516001600160a01b038116908190036200007a57600780546001600160a01b031916919091179055604051615aca9081620000968239f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe604060808152600436101561001357600080fd5b6000803560e01c918263c0d91eaf1461009c5750508063cba510a914610097578063df5fe70414610092578063e08322d51461008d578063e2761af014610088578063e405bbc314610083578063ead3cbe71461007e5763facb618d1461007957600080fd5b610a0b565b610977565b6107f8565b6107d1565b610740565b61058f565b61051b565b346103655780600319360112610365576004356100b881610369565b6024356001600160401b0381116103615736602382011215610361576100e89036906024816004013591016104c2565b9060ff84541661031d5761021583926102106001600160a01b0361021d946101db7fa004c198518af96775c486676e922051a150cf22b054780d1223d312127b72f19851602081018161015f826020907f2f74656e6465726d696e742e74797065732e436f6e73656e737573537461746581520190565b0391610173601f199384810183528261043d565b5190209089516101bc6020820192826101b085601a907f2f74656e6465726d696e742e74797065732e546d48656164657200000000000081520190565b0390810183528261043d565b5190206101c761045e565b828152602001819052600191909155600255565b166001600160a01b03167fffffffffffffffffffffffff00000000000000000000000000000000000000006006541617600655565b61136c565b919091610a3f565b6102c161024761024161023585855151015160070b90565b6001600160401b031690565b92614a17565b9161026e83610269836001600160401b03166000526003602052604060002090565b610bad565b61028e816001600160401b03166001600160401b03196005541617600555565b6fffffffffffffffff00000000000000006005549160401b16906fffffffffffffffff0000000000000000191617600555565b6103056102e460206102db6005546001600160401b031690565b93015151610cdf565b925192839283602090939291936001600160401b0360408201951681520152565b0390a161031a600160ff196000541617600055565b80f35b825162461bcd60e51b815260206004820152601360248201527f616c726561647920696e697469616c697a6564000000000000000000000000006044820152606490fd5b8380fd5b5080fd5b6001600160a01b0381160361037a57565b600080fd5b634e487b7160e01b600052604160045260246000fd5b604081019081106001600160401b038211176103b057604052565b61037f565b602081019081106001600160401b038211176103b057604052565b606081019081106001600160401b038211176103b057604052565b608081019081106001600160401b038211176103b057604052565b60c081019081106001600160401b038211176103b057604052565b61010081019081106001600160401b038211176103b057604052565b90601f801991011681019081106001600160401b038211176103b057604052565b6040519061046b82610395565b565b6040519061046b82610421565b6040519061046b82610406565b604051906101c082018281106001600160401b038211176103b057604052565b6001600160401b0381116103b057601f01601f191660200190565b9291926104ce826104a7565b916104dc604051938461043d565b82948184528183011161037a578281602093846000960137010152565b602090600319011261037a576004356001600160401b038116810361037a5790565b3461037a576001600160401b03610531366104f9565b1660005260036020526020600160406000200160009061055181546105ff565b9181601f8411610582575b50505490828110610571575b50604051908152f35b60001990830360031b1b1638610568565b849250815220388061055c565b3461037a5760206001600160401b03806105a8366104f9565b16806000526004835260ff604060002054169182156105ce575b50506040519015158152f35b6005541614905038806105c2565b906040516105e981610395565b60208193548060070b835260401c60030b910152565b90600182811c9216801561062f575b602083101461061957565b634e487b7160e01b600052602260045260246000fd5b91607f169161060e565b906040519182600082549261064d846105ff565b9081845260019485811690816000146106ba5750600114610677575b505061046b9250038361043d565b9093915060005260209081600020936000915b8183106106a257505061046b93508201013880610669565b8554888401850152948501948794509183019161068a565b91505061046b94506020925060ff191682840152151560051b8201013880610669565b906040516106ea816103b5565b6106f48193610639565b9052565b60005b83811061070b5750506000910152565b81810151838201526020016106fb565b90602091610734815180928185528580860191016106f8565b601f01601f1916010190565b3461037a576001600160401b03610756366104f9565b166000526003602052602060406000206107cd610772826105dc565b916107bf61078e6002610787600185016106dd565b9301610639565b91604051958587965160070b8752015160030b602086015260806040860152516020608086015260a085019061071b565b90838203606085015261071b565b0390f35b3461037a57600036600319011261037a5760206001600160401b0360055416604051908152f35b3461037a57600036600319011261037a5760206005546001600160401b036040519160401c168152f35b806043121561037a576040519061083882610395565b81606491821161037a576024905b8282106108535750505090565b8135815260209182019101610846565b80610103121561037a576040519061087a82610395565b8161012491821161037a5760e4905b8282106108965750505090565b8135815260209182019101610889565b80610143121561037a57604051906108bd82610395565b8161016491821161037a57610124905b8282106108da5750505090565b81358152602091820191016108cd565b806083121561037a5760409081519161090283610395565b60e48383821161037a576064905b82821061091f57505050505090565b84601f8301121561037a57835161093581610395565b8085840187811161037a5791869285949294905b80821061096157505081529201916020019050610910565b8135865260209586019589955090910190610949565b3461037a5761018036600319011261037a576004356001600160401b0380821161037a573660238301121561037a57816004013590811161037a57366024828401011161037a576107cd916109f9916109cf36610822565b6109d8366108ea565b6109e136610863565b916109eb366108a6565b936024610164359601610e7a565b60405190151581529081906020820190565b3461037a576001600160401b03610a21366104f9565b166000526004602052602060ff604060002054166040519015158152f35b15610a4657565b60405162461bcd60e51b815260206004820152601a60248201527f4c433a206c6967687420626c6f636b20697320696e76616c69640000000000006044820152606490fd5b90601f8111610a9957505050565b600091825260208220906020601f850160051c83019410610ad5575b601f0160051c01915b828110610aca57505050565b818155600101610abe565b9092508290610ab5565b91909182516001600160401b0381116103b057610b0681610b0084546105ff565b84610a8b565b602080601f8311600114610b42575081929394600092610b37575b50508160011b916000199060031b1c1916179055565b015190503880610b21565b90601f19831695610b5885600052602060002090565b926000905b888210610b9557505083600195969710610b7c575b505050811b019055565b015160001960f88460031b161c19169055388080610b72565b80600185968294968601518155019501930190610b5d565b81518051825460209283015160401b6bffffffff0000000000000000166001600160401b039283166bffffffffffffffffffffffff19909216919091171783558184015151805190949260018086019390919083116103b057610c1a83610c1486546105ff565b86610a8b565b80601f8411600114610c645750918080604095936002979561046b9a600094610c59575b50501b916000199060031b1c19161790555b01519101610adf565b015192503880610c3e565b9193949596601f198416610c7d87600052602060002090565b936000905b828210610cc85750509260409592859261046b9a99966002999610610caf575b505050811b019055610c50565b015160001960f88460031b161c19169055388080610ca2565b808886978294978701518155019601940190610c82565b602081519101519060208110610cf3575090565b6000199060200360031b1b1690565b9081602091031261037a5751610d1781610369565b90565b6040513d6000823e3d90fd5b15610d2d57565b60405162461bcd60e51b815260206004820152600b60248201527f6e6f742072656c617965720000000000000000000000000000000000000000006044820152606490fd5b15610d7957565b60405162461bcd60e51b815260206004820152601c60248201527f63616e27742073796e63206475706c69636174656420686561646572000000006044820152606490fd5b15610dc557565b60405162461bcd60e51b815260206004820152603760248201527f4c433a2068656164657220686569676874206e6f74206e65776572207468616e60448201527f20636f6e73656e737573207374617465206865696768740000000000000000006064820152608490fd5b90604051610e3d816103d0565b6040610e7560028395610e4f816105dc565b85528351610e5c816103b5565b610e6860018301610639565b8152602086015201610639565b910152565b9060049294969593956020610ea6610e9a6006546001600160a01b031690565b6001600160a01b031690565b604051638406c07960e01b815295869182905afa948515611074577f82639a39e1f497d8456613d012ef21caf479c800899e97b68647432032a41ed998610f629861102598610f23610210610f2b98610fca98610f1c6001600160a01b0360209e610ffc9e600091611046575b50163314610d26565b36916104c2565b979097610a3f565b610f72610f6d610f69610f4761023560408c5151015160070b90565b9d8e6001600160401b03166000526004602052604060002090565b5460ff1690565b1590565b610d72565b86610fc5610fc08d610fa7610f936005546001600160401b039060401c1690565b916001600160401b03808416911611610dbe565b6001600160401b03166000526003602052604060002090565b610e30565b6111f1565b610ff7610fea866001600160401b03166000526004602052604060002090565b805460ff19166001179055565b614a17565b61101d81610269866001600160401b03166000526003602052604060002090565b015151610cdf565b604080516001600160401b03939093168352602083019190915290a1600190565b602061106792503d811161106d575b61105f818361043d565b810190610d02565b38610f13565b503d611055565b610d1a565b6040519061108682610395565b60006020838281520152565b6040519061109f82610395565b6060602083600081520152565b604051906110b982610395565b81606081526020610e75611092565b604051906101c082018281106001600160401b038211176103b057604052816110ef611079565b81526101a06060918260208201526000604082015261110c611079565b838201526111186110ac565b60808201528260a08201528260c08201528260e082015282610100820152826101208201528261014082015282610160820152826101808201520152565b60405190611163826103eb565b60608083600081526000602082015261117a6110ac565b60408201520152565b6040519061119082610395565b816111996110c8565b81526020610e75611156565b156111ac57565b60405162461bcd60e51b815260206004820152601b60248201527f4c433a206661696c656420746f207665726966792068656164657200000000006044820152606490fd5b9161046b9694926112a896949260406112086110c8565b926020835151015160208501526001600160401b03600554831c1660070b82850152015191610100928382015261123d611183565b9081526020808351930151916112856001600160a01b0360075416956112648386886116fe565b61127f61127560e088510151611595565b9184510151611595565b14611518565b510151916020810151926112a26040808601519351015160070b90565b92611a91565b6111a5565b604051906112ba826103d0565b60606040838281528260208201520152565b604051906112d9826103eb565b60006060838181526112e96112ad565b60208201528260408201520152565b60405190611305826103d0565b6000604083606081526113166112cc565b60208201520152565b6040519061132c826103eb565b81611335611183565b815261133f6112f8565b6020820152600060408201526060610e756112f8565b90611368602092828151948592016106f8565b0190565b61137461131f565b61137c612151565b508151611387612151565b50611390612151565b90611399611c52565b9260205b6113a683611d6a565b8110156114c757806113bb876113c49361514c565b90939192611de8565b9160018281036113f4575050506113ea816113e487878a6113a69661219f565b90611de8565b925b92905061139d565b600292830361141957505050611413816113e487878a6113a6966121e4565b926113ec565b6114258296949661216a565b81146114b2575b6114358161216a565b6005811461149d575b6114478161216a565b8015611488575b6114578161216a565b14611466575b6113a6906113ec565b9161148081611478886113a694615357565b919050611de8565b92905061145d565b9380611478896114979361516e565b9361144e565b9380611478896114ac93615269565b9361143e565b9380611478896114c1936152ed565b9361142c565b50925050915081516040516114f2816114e4602082018095611355565b03601f19810183528261043d565b519020600254036115115750602061150b9101516127a7565b90600190565b9160009150565b1561151f57565b60405162461bcd60e51b815260206004820152604260248201527f6578706563746564206f6c6420686561646572206e6578742076616c6964617460448201527f6f727320746f206d617463682074686f73652066726f6d206e6577206865616460648201526132b960f11b608482015260a490fd5b60208151036115a5576020015190565b60405162461bcd60e51b815260206004820152601d60248201527f42797465733a20746f4279746573333220696e76616c69642073697a650000006044820152606490fd5b156115f157565b60405162461bcd60e51b815260206004820152601f60248201527f6865616465722062656c6f6e677320746f20616e6f7468657220636861696e006044820152606490fd5b1561163d57565b60405162461bcd60e51b815260206004820152602160248201527f68656164657220616e6420636f6d6d697420686569676874206d69736d6174636044820152600d60fb1b6064820152608490fd5b1561169357565b60405162461bcd60e51b815260206004820152603f60248201527f6578706563746564206e6577206865616465722068656967687420746f20626560448201527f2067726561746572207468616e206f6e65206f66206f6c6420686561646572006064820152608490fd5b6117c9926117ba61178d60406117b0816117c49661176a6020825101518351908161172d602082018093611355565b0391611741601f199384810183528261043d565b5190209060208a51015190611761865191826101b0602082018096611355565b519020146115ea565b61179d61177c60208301515160070b90565b61179361178d858551015160070b90565b60070b90565b9060070b14611636565b6117a681614bc6565b5051015160070b90565b9351015160070b90565b9060070b1361168c565b6150e4565b50565b156117d357565b60405162461bcd60e51b815260206004820152601960248201527f696e76616c696420636f6d6d6974207369676e617475726573000000000000006044820152606490fd5b1561181f57565b60405162461bcd60e51b815260206004820152601760248201527f696e73756666696369656e74207369676e6174757265730000000000000000006044820152606490fd5b1561186b57565b60405162461bcd60e51b815260206004820152601560248201527f696e76616c696420636f6d6d69742068656967687400000000000000000000006044820152606490fd5b156118b757565b606460405162461bcd60e51b815260206004820152602060248201527f696e76616c696420636f6d6d6974202d2d2077726f6e6720626c6f636b2049446044820152fd5b6040519061190882610421565b8160005b6101008110611919575050565b60608282015260200161190c565b634e487b7160e01b600052601160045260246000fd5b600019811461194c5760010190565b611927565b634e487b7160e01b600052603260045260246000fd5b805182101561197b5760209160051b010190565b611951565b634e487b7160e01b600052602160045260246000fd5b600411156119a057565b611980565b51610d1781611996565b90600881101561197b5760051b0190565b9081602091031261037a5751801515810361037a5790565b6000915b600283106119e957505050565b6001908251815260208091019201920191906119dc565b929095949391611a15846108608101986119d8565b600060408581015b60028310611a7457505050508291611a3f611a4a9260c06101409601906119d8565b6101008301906119d8565b01906000915b60398310611a5d57505050565b600190825181526020809101920192019190611a50565b60208282611a8560019488516119d8565b01940192019192611a1d565b929091611ab89a949297969597611af2611aed86515193606087019e8f95865151146117cc565b611ae3600896611acb8888515111611818565b611ad961178d8a5160070b90565b9060070b14611864565b6040870151614a86565b6118b0565b611afa6118fb565b93611b036118fb565b95600080945b518051861015611c2f57611b28611b2287600293611967565b516119a5565b611b3181611996565b03611c2557611b8f906020611b47878551611967565b51015151611b55828a6119af565b52611b6081896119af565b50611b74611b6f8787876120a2565b614813565b611b7e828b6119af565b52611b89818a6119af565b5061193d565b93858514611ba757611ba09061193d565b938e611b09565b50505050505060209798506001600160a01b039392611be692611bc992611e65565b604051635db27c3d60e01b81529889978896879560048701611a00565b0392165afa90811561107457600091611bfd575090565b610d17915060203d8111611c1e575b611c16818361043d565b8101906119c0565b503d611c0c565b93611ba09061193d565b5050505050505060209798506001600160a01b039392611be692611bc992611e65565b60405190611c5f826103d0565b6060368337565b604051906107208083018381106001600160401b038211176103b057604052368337565b6040519060a082018281106001600160401b038211176103b05760405260a0368337565b604051906101e08083018381106001600160401b038211176103b057604052368337565b60405190611cdf826103eb565b6080368337565b906020825192015166ffffffffffffff19908181169360198110611d0957505050565b60190360031b82901b16169150565b908160050291600583040361194c57565b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81160361194c57565b90603981101561197b5760051b0190565b602001908160201161194c57565b906001820180921161194c57565b906002820180921161194c57565b906003820180921161194c57565b906004820180921161194c57565b906028820180921161194c57565b906009820180921161194c57565b600101908160011161194c57565b906020820180921161194c57565b9190820180921161194c57565b906020825192015169ffffffffffffffffffff19908181169360168110611e1b57505050565b60160360031b82901b16169150565b90602082519201516fffffffffffffffffffffffffffffffff19908181169360108110611e5657505050565b60100360031b82901b16169150565b9190611e6f611c66565b9260005b60088110611e875750505061070082015290565b611e9181846119af565b51611e9b90614600565b611ea582856119af565b51611eaf9061464b565b611eb983866119af565b51611ec390614698565b611ecd84876119af565b51611ed7906146e4565b91611ee285886119af565b51611eec90614730565b93611ef690611ce6565b60381c611f0286611d18565b611f0c908b611d59565b52611f1690611ce6565b60381c611f2285611d18565b611f2b90611d78565b611f35908a611d59565b52611f3f90611ce6565b60381c611f4b84611d18565b611f5490611d86565b611f5e9089611d59565b52611f6890611ce6565b60381c611f7483611d18565b611f7d90611d94565b611f879088611d59565b52611f9190611df5565b60501c611f9d82611d18565b611fa690611da2565b611fb09087611d59565b52611fbb81836119af565b51611fc59061477c565b611fcf82846119af565b51611fd9906147c7565b90611fe390611e2a565b60801c611fef83611d29565b611ff890611db0565b6120029088611d59565b5261200c90611e2a565b60801c61201882611d29565b61202190611d78565b61202a90611db0565b6120349087611d59565b5261203e9061193d565b611e73565b6040519061010082018281106001600160401b038211176103b057604052606060e0836000815260006020820152600060408201526120806110ac565b8382015261208c611079565b60808201528260a0820152600060c08201520152565b90610d179261214c926120b3612043565b506120bc612043565b506121426120ce836060840151611967565b51825160070b9360406120e5602086015160030b90565b94015160408301519061212960606121016020870151956155ce565b9501519661211f61211061046d565b600281529960070b60208b0152565b60030b6040890152565b6060870152608086015260a085015260030b60c0840152565b60e0820152614913565b612251565b6040519061215e82610395565b60606020838281520152565b600611156119a057565b60001981019190821161194c57565b601f1981019190821161194c57565b9190820391821161194c57565b60209293916121ad91615357565b93806121c5575050016121c08151611d78565b905290565b5201805190816121d457505090565b600019820191821161194c575290565b60409293916121f291615357565b9380612205575050016121c08151611d78565b602001520180518061221657505090565b6121c090612174565b90612229826104a7565b612236604051918261043d565b8281528092612247601f19916104a7565b0190602036910137565b61226f612267825161226281611996565b614568565b60030b6155a8565b6001018060011161194c5761228661228b91611dbe565b611dbe565b6122a86122a361229e606085015161263e565b61554f565b611dcc565b810180911161194c576122c46122a361229e608085015161277d565b810180911161194c576122ed6122e86122f4926113e46122a360a08701515161554f565b61221f565b80926122f9565b815290565b610d17916020825161230a81611996565b61231381611996565b61240f575b61238a61237b61237161236260a0948660208901612337815160070b90565b60070b6123ef575b506040890161235261178d825160070b90565b6123c0575b506113e4908261540e565b6113e4868260608a0151612455565b6113e4858261541b565b6113e484826080880151612666565b92019081515161239c575b5050612183565b806123b1846113e46113e4946123b997615428565b8093516154a7565b3880612395565b816113e4916123d9856113e46123e3966123e898615401565b9384915160070b90565b615518565b8638612357565b816113e4916123d9856113e46123e396612408986153f4565b863861233f565b5060a061238a61237b61237161236261244a61243261242d886153bc565b611d6a565b6113e488826124458c5161226281611996565b615504565b945050505050612318565b91906124636122e88461263e565b90602093805151612516575b6000939260129290868101845b60071c94851561249e576080178153600196870196607f86169591019061247c565b9060019397929695509793975382010180911161194c5761250c83612506936124e16124dc856113e461251199612500986020610d179e0151612547565b612183565b9485926124f86124f2848b8761546b565b8a611de8565b958691611dda565b92611de8565b906150f3565b611de8565b612192565b9361252083615435565b908181019182821161194c57612538858489516154a7565b010180911161194c579361246f565b916125546122e8846125f9565b9060208451516125ca575b61250c83612506936124e161250094612511976020610d179b0163ffffffff61258c825163ffffffff1690565b16612598575050612183565b816113e4916125b1856113e46125be966125c398615444565b9384915163ffffffff1690565b6154f3565b8338612395565b6125d383615435565b908181019182821161194c576125eb858489516154a7565b010180911161194c5761255f565b61260c63ffffffff602083015116615564565b80600101918260011161194c5761262490515161554f565b8060010160011161194c57600291010180911161194c5790565b61264981515161554f565b80600101918260011161194c5761229e60206126249201516125f9565b9161267d6126766122e88561277d565b80946126c8565b9061268981848461546b565b83019384841161194c576020810180911161194c5784820180921161194c576126b39183916150f3565b820180921161194c57810390811161194c5790565b6020908051600790810b612713575b50916020610d179301805160030b6126f0575050612183565b816113e491612709856113e4612445966123b998615444565b9384915160030b90565b600092600892858201845b841c94851561273f576080178153600195860195607f86169591019061271e565b9296945090949250536021810180941161194c5760219061276b838686516001600160401b031661546b565b010180931161194c57919060206126d7565b61278a815160070b6155a8565b80600101918260011161194c57602061262491015160030b615585565b906127b061131f565b508151906127bc61131f565b506127c561131f565b6127cd611c8a565b9160205b6127da85611d6a565b811015612933576127f86127ee878361514c565b9092919293611de8565b91826001838103612824575050505061281a816113e486868a6127da9661293c565b945b9490506127d1565b60029380850361284c575050505050612846816113e486868a6127da96612ae4565b9461281c565b6003810361286c575050505050612846816113e486868a6127da96612b17565b60040361288a5750505050612846816113e486868a6127da96612b4f565b6128968399959961216a565b821461291c575b506128a78161216a565b60058114612907575b6128b98161216a565b80156128f2575b6128c98161216a565b146128d8575b6127da9061281c565b936128ea81611478886127da94615357565b9490506128cf565b9580611478896129019361516e565b956128c0565b95806114788961291693615269565b956128b0565b61292c91975061147889826152ed565b953861289d565b50935091505090565b92919093612948611183565b50612953858561516e565b92909383860180961161194c57612968611183565b50612971611183565b9361297a611c52565b95875b612987828a611de8565b811015612aad578188888c6129a961299f828761514c565b9092919297611de8565b9560018281036129d4575050506129c99285926113e49261298797612b82565b995b9991505061297d565b9295506002945092508303612a00575050506129fa826113e48b8b879561298797612e90565b996129cb565b90919250612a10829c949c61216a565b8114612a98575b612a208161216a565b60058114612a83575b612a328161216a565b8015612a6e575b612a428161216a565b14612a52575b81612987916129cb565b9781612a65826114788d61298795615357565b99915050612a48565b99806114788d612a7d9361516e565b99612a39565b99806114788d612a9293615269565b99612a29565b99806114788d612aa7936152ed565b99612a17565b50909194939750612ac392965060209550611de8565b9380612ad6575050016121c08151611d78565b520180518061221657505090565b6040929391612af291613e3c565b9380612b05575050016121c08151611d78565b6020015201805190816121d457505090565b6060929391612b259161516e565b939081612b39575050016121c08151611d78565b60409060070b9101520180518061221657505090565b6080929391612b5d91613e3c565b9380612b70575050016121c08151611d78565b6060015201805190816121d457505090565b92919092612b8e6110c8565b50612b99848261516e565b92909383830180931161194c57612bae6110c8565b50612bb76110c8565b93612bc0611cae565b95845b612bcd8287611de8565b811015612e7b578188888b84612be682612bef9761514c565b90979192611de8565b95600191808303612c1b57505050612c109285926113e492612bcd976131c1565b965b96915050612bc3565b600292818403612c425750505050612c3c9285926113e492612bcd976121e4565b96612c12565b60038203612c615750505050612c3c9285926113e492612bcd97612b17565b60048203612c805750505050612c3c9285926113e492612bcd97613337565b600591808303612ca2575050505050612c3c9285926113e492612bcd97613369565b60068103612cc2575050505050612c3c9285926113e492612bcd9761339b565b60078103612ce2575050505050612c3c9285926113e492612bcd976133cd565b60088103612d02575050505050612c3c9285926113e492612bcd976133ff565b60098103612d22575050505050612c3c9285926113e492612bcd97613432565b600a8103612d42575050505050612c3c9285926113e492612bcd97613466565b600b8103612d62575050505050612c3c9285926113e492612bcd9761349a565b600c8103612d82575050505050612c3c9285926113e492612bcd976134ce565b600d8103612da2575050505050612c3c9285926113e492612bcd97613502565b600e91929394959697985014600014612dcd5750505050926113e48593612c3c93612bcd9684613536565b92955092509250612de0839a959a61216a565b8214612e66575b612df08261216a565b8114612e51575b612e008161216a565b8015612e3c575b612e108161216a565b14612e20575b81612bcd91612c12565b9481612e33826114788c612bcd95615357565b96915050612e16565b96806114788c612e4b9361516e565b96612e07565b96806114788c612e6093615269565b96612df7565b97806114788d612e75936152ed565b97612de7565b509091949350612ac392965060209550611de8565b91929092612e9c611156565b50612ea7848461516e565b93909184820180921161194c57612ebc611156565b50612ec5611156565b93612ece611c8a565b95835b612edb8686611de8565b81101561300857806113bb8a612ef09361514c565b916001828103612f1557505050806113e489898c612f0d95613b5b565b945b94612ed1565b600292808403612f3a5750505050806113e489898c612f339561378a565b945b612f0f565b60038103612f565750505050806113e489898c612f3395613b90565b600403612f7757505050612f33906113e489612f70613dd0565b8c84613bc3565b612f838299949961216a565b8114612ff3575b612f938161216a565b60058114612fde575b612fa58161216a565b8015612fc9575b612fb58161216a565b03612f355794806114788a612f3393615357565b96806114788c612fd89361516e565b96612fac565b96806114788c612fed93615269565b96612f9c565b96806114788c613002936152ed565b96612f8a565b50909192938461301b6080890151613b0b565b60608801525b61302b8287611de8565b8110156131ac578188888b84612be6826130449761514c565b9560018281036130795750505061302b94915061306e926113e491613067613dd0565b9084613b5b565b965b96915050613021565b6002928084036130aa575050505061302b9491506130a4926113e49161309d613dd0565b908461378a565b96613070565b600381036130d3575050505061302b9491506130a4926113e4916130cc613dd0565b9084613b90565b60049192939495969750146000146130fc57505050926113e485936130a49361302b9684613bc3565b925061310f91945080935099949961216a565b8114613197575b61311f8161216a565b60058114613182575b6131318161216a565b801561316d575b6131418161216a565b14613151575b8161302b91613070565b9481613164826114788c61302b95615357565b96915050613147565b96806114788c61317c9361516e565b96613138565b96806114788c61319193615269565b96613128565b96806114788c6131a6936152ed565b96613116565b5090919493506121f292965060409550611de8565b929190926131cd611079565b506131d8848261516e565b92909383830180931161194c576131ed611079565b506131f6611079565b936131ff611c52565b95845b61320c8287611de8565b811015612e7b578188888b84612be6826132259761514c565b9586600183810361325257505050506132479285926113e49261320c9761356a565b965b96915050613202565b909192939495965060028094146000146132845750505050926113e4859361327e9361320c96846135a4565b96613249565b92955092509250613297839a959a61216a565b8214613320575b506132a88161216a565b6005811461330b575b6132ba8161216a565b80156132f6575b6132ca8161216a565b146132da575b8161320c91613249565b94816132ed826114788c61320c95615357565b969150506132d0565b96806114788c6133059361516e565b966132c1565b96806114788c61331a93615269565b966132b1565b6133309198506114788c826152ed565b963861329e565b6080929391613345916135e2565b9380613358575050016121c08151611d78565b606001520180518061221657505090565b60a0929391613377916137c2565b938061338a575050016121c08151611d78565b608001520180518061221657505090565b60c09293916133a991615357565b93806133bc575050016121c08151611d78565b60a001520180518061221657505090565b60e09293916133db91615357565b93806133ee575050016121c08151611d78565b60c001520180518061221657505090565b61010092939161340e91615357565b9380613421575050016121c08151611d78565b60e001520180518061221657505090565b61012092939161344191615357565b9380613454575050016121c08151611d78565b61010001520180518061221657505090565b61014092939161347591615357565b9380613488575050016121c08151611d78565b61012001520180518061221657505090565b6101609293916134a991615357565b93806134bc575050016121c08151611d78565b61014001520180518061221657505090565b6101809293916134dd91615357565b93806134f0575050016121c08151611d78565b61016001520180518061221657505090565b6101a092939161351191615357565b9380613524575050016121c08151611d78565b61018001520180518061221657505090565b6101c092939161354591615357565b9380613558575050016121c08151611d78565b6101a001520180518061221657505090565b60209293916135789161516e565b93908161358c575050016121c08151611d78565b6001600160401b031690520180518061221657505090565b60409293916135b29161516e565b9390816135c6575050016121c08151611d78565b6001600160401b03602091169101520180518061221657505090565b916135eb611079565b506135f6828461516e565b93909184820180921161194c5761360b611079565b50613614611079565b9461361d611c52565b93835b61362a8286611de8565b811015613746578181613640896136499461514c565b90949192611de8565b92600182810361367457505050613669826113e4898c8c61362a97613756565b955b95915050613620565b600293509091830361369d5750505081613697826113e4898c8c61362a9761378a565b9561366b565b6136a98298949861216a565b8114613731575b6136b98161216a565b6005811461371c575b6136cb8161216a565b8015613707575b6136db8161216a565b146136eb575b8161362a9161366b565b93816136fe826114788a61362a95615357565b959150506136e1565b95806114788a6137169361516e565b956136d2565b95806114788a61372b93615269565b956136c2565b95806114788a613740936152ed565b956136b0565b509194509150610d179250611de8565b60209293916137649161516e565b939081613778575050016121c08151611d78565b60070b90520180518061221657505090565b60409293916137989161516e565b9390816137ac575050016121c08151611d78565b60209060030b9101520180518061221657505090565b916137cb6110ac565b506137d6828461516e565b93909184820180921161194c576137eb6110ac565b506137f46110ac565b946137fd611c52565b93835b61380a8286611de8565b811015613746578181613640896138209461514c565b92600182810361384b57505050613840826113e4898c8c61380a9761219f565b955b95915050613800565b6002935090918303613874575050508161386e826113e4898c8c61380a9761391d565b95613842565b6138808298949861216a565b8114613908575b6138908161216a565b600581146138f3575b6138a28161216a565b80156138de575b6138b28161216a565b146138c2575b8161380a91613842565b93816138d5826114788a61380a95615357565b959150506138b8565b95806114788a6138ed9361516e565b956138a9565b95806114788a61390293615269565b95613899565b95806114788a613917936152ed565b95613887565b92919093613929611092565b50613934858561516e565b92909383860180961161194c57613949611092565b50613952611092565b9361395b611c52565b95875b613968828a611de8565b811015613a84578188888c61398061299f828761514c565b9560018281036139ab575050506139a09285926113e49261396897613a9a565b995b9991505061395e565b92955060029450925083036139d7575050506139d1826113e48b8b8795613968976121e4565b996139a2565b909192506139e7829c949c61216a565b8114613a6f575b6139f78161216a565b60058114613a5a575b613a098161216a565b8015613a45575b613a198161216a565b14613a29575b81613968916139a2565b9781613a3c826114788d61396895615357565b99915050613a1f565b99806114788d613a549361516e565b99613a10565b99806114788d613a6993615269565b99613a00565b99806114788d613a7e936152ed565b996139ee565b509091949397506121f292965060409550611de8565b6020929391613aa89161516e565b939081613abc575050016121c08151611d78565b63ffffffff1690520180518061221657505090565b6001600160401b0381116103b05760051b60200190565b60405190613af5826103eb565b606080836000815281602082015261117a611079565b90613b1582613ad1565b613b22604051918261043d565b8281528092613b33601f1991613ad1565b019060005b828110613b4457505050565b602090613b4f613ae8565b82828501015201613b38565b6020929391613b699161516e565b939081613b7d575050016121c08151611d78565b60070b905201805190816121d457505090565b6060929391613b9e916137c2565b9380613bb1575050016121c08151611d78565b6040015201805190816121d457505090565b92919092613bcf613ae8565b50613bda848261516e565b92909383830180931161194c57613bef613ae8565b50613bf8613ae8565b93613c01611c8a565b95845b613c0e8287611de8565b811015613d71578188888b84612be682613c279761514c565b956001828103613c5257505050613c479285926113e492613c0e97613dde565b965b96915050613c04565b600292808403613c795750505050613c739285926113e492613c0e97613e12565b96613c49565b60038103613c985750505050613c739285926113e492613c0e97613e20565b6004919293949596975014600014613cc157505050926113e48593613c7393613c0e9684613e2e565b9250613cd491945080935099949961216a565b8114613d5c575b613ce48161216a565b60058114613d47575b613cf68161216a565b8015613d32575b613d068161216a565b14613d16575b81613c0e91613c49565b9481613d29826114788c613c0e95615357565b96915050613d0c565b96806114788c613d419361516e565b96613cfd565b96806114788c613d5693615269565b96613ced565b96806114788c613d6b936152ed565b96613cdb565b50919550919550613d829250611de8565b9280613d975750506080016121c08151611d78565b906060613dc3920151613db260808251950194855190612192565b91613dbd8383611967565b52611967565b5080518061221657505090565b613dd8611156565b50600090565b91613def613df6949260209461516e565b9490614532565b9080613e09575050016121c08151611d78565b6121c582611996565b6040929391612af291615357565b6060929391613b9e916135e2565b6080929391612b5d91615357565b613e446112f8565b50613e4f828261516e565b9182810180911161194c57613e626112f8565b50613e6b6112f8565b94613e74611cd2565b93825b613e818585611de8565b811015613f9257806113bb88613e969361514c565b916001828103613ec257505050613eba906113e487613eb3614337565b8a8461414e565b935b93613e77565b600292808403613ee75750505050806113e4878a8a613ee095614192565b935b613ebc565b600303613f0157505050806113e4878a8a613ee095612b17565b613f0d8298949861216a565b8114613f7d575b613f1d8161216a565b60058114613f68575b613f2f8161216a565b8015613f53575b613f3f8161216a565b03613ee257938061147888613ee093615357565b95806114788a613f629361516e565b95613f36565b95806114788a613f7793615269565b95613f26565b95806114788a613f8c936152ed565b95613f14565b5090919283613fa460208701516140fe565b88525b613fb18286611de8565b81101561374657818161364089613fc79461514c565b92836001838103613ff45750505050613fe9826113e4898c8c613fb19761414e565b955b95915050613fa7565b60029380850361402357505050505061401d613fb1926113e489614016614337565b8c84614192565b95613feb565b60039192939495501460001461405257505050508161401d613fb1926113e48961404b614337565b8c84612b17565b61405e8399959961216a565b82146140e7575b5061406f8161216a565b600581146140d2575b6140818161216a565b80156140bd575b6140918161216a565b146140a1575b81613fb191613feb565b93816140b4826114788a613fb195615357565b95915050614097565b95806114788a6140cc9361516e565b95614088565b95806114788a6140e193615269565b95614078565b6140f79197506114788a826152ed565b9538614065565b9061410882613ad1565b614115604051918261043d565b8281528092614126601f1991613ad1565b019060005b82811061413757505050565b6020906141426112cc565b8282850101520161412b565b9061415b9193929361419c565b92806141705750506020016121c08151611d78565b5191602083519101928351820391821161194c57613dc392613dbd8383611967565b60409293916121f2915b916141a56112cc565b506141b0828461516e565b93909184820180921161194c576141c56112cc565b506141ce6112cc565b946141d7611c8a565b93835b6141e48286611de8565b811015613746578181613640896141fa9461514c565b9260018281036142255750505061421a826113e4898c8c6141e49761219f565b955b959150506141da565b60029280840361424c5750505050614246826113e4898c8c6141e49761433f565b9561421c565b6003810361426b5750505050614246826113e4898c8c6141e497612b17565b929350909160040361428e5750505081614246826113e4898c8c6141e4976144c8565b61429a8298949861216a565b8114614322575b6142aa8161216a565b6005811461430d575b6142bc8161216a565b80156142f8575b6142cc8161216a565b146142dc575b816141e49161421c565b93816142ef826114788a6141e495615357565b959150506142d2565b95806114788a6143079361516e565b956142c3565b95806114788a61431c93615269565b956142b3565b95806114788a614331936152ed565b956142a1565b613dd86112f8565b9291909361434b6112ad565b50614356858561516e565b92909383860180961161194c5761436b6112ad565b506143746112ad565b9361437d611cd2565b95875b61438a828a611de8565b811015613a84578188888c6143a261299f828761514c565b9560018281036143cd575050506143c29285926113e49261438a9761219f565b995b99915050614380565b6002928084036143f457505050506143ee9285926113e49261438a976121e4565b996143c4565b93965091945092509060030361441b575050506143ee826113e48b8b879561438a97614500565b9091925061442b829c949c61216a565b81146144b3575b61443b8161216a565b6005811461449e575b61444d8161216a565b8015614489575b61445d8161216a565b1461446d575b8161438a916143c4565b9781614480826114788d61438a95615357565b99915050614463565b99806114788d6144989361516e565b99614454565b99806114788d6144ad93615269565b99614444565b99806114788d6144c2936152ed565b99614432565b60809293916144d69161516e565b9390816144ea575050016121c08151611d78565b60609060070b9101520180518061221657505090565b606092939161450e91615357565b9380614521575050016121c08151611d78565b604001520180518061221657505090565b60070b8015613dd85760018114614562576002811461455c5760031461455757600080fd5b600390565b50600290565b50600190565b61457181611996565b8015613dd85761458081611996565b600181146145625761459181611996565b6002811461455c57806145a5600392611996565b146145af57600080fd5b602090565b156145bb57565b60405162461bcd60e51b815260206004820152601160248201527f736c6963655f6f75744f66426f756e64730000000000000000000000000000006044820152606490fd5b61460e6019825110156145b4565b6040519060198083019101603283015b80831061463857505060198252601f01601f191660405290565b909182518152602080910192019061461e565b6146596032825110156145b4565b60405190601982019060328084019101905b80831061468557505060198252601f01601f191660405290565b909182518152602080910192019061466b565b6146a6604b825110156145b4565b604051906019820190604b01603283015b8083106146d157505060198252601f01601f191660405290565b90918251815260208091019201906146b7565b6146f26064825110156145b4565b604051906019820190606401603283015b80831061471d57505060198252601f01601f191660405290565b9091825181526020809101920190614703565b61473e607a825110156145b4565b604051906016820190607a01602c83015b80831061476957505060168252601f01601f191660405290565b909182518152602080910192019061474f565b61478a6010825110156145b4565b604051906010808301910160208084015b8084106147b65750506010835250601f01601f191660405290565b82518452928101929181019161479b565b60206147d681835110156145b4565b604051916010830191808085019201915b8084106148025750506010835250601f01601f191660405290565b8251845292810192918101916147e7565b6001600160401b039182825110156148945781519283166148366122e882615564565b9060009260208301607f809716925b60071c92831561486657608017815360019485019483881693910190614845565b61488e969297506114e4935094610d1795600192530183526040519485936020850190611355565b90611355565b60405162461bcd60e51b815260206004820152601f60248201527f456e636f6465723a206f7574206f6620626f756e6473202875696e74363429006044820152606490fd5b604051906148e682610395565b606082526040516020836148f983610395565b606083526000828401520152565b61491082611996565b52565b90606060a060405161492481610406565b60008152600060208201526000604082015261493e6148d9565b8382015261494a611079565b6080820152015281519161495d83611996565b602081015160070b906149bc61498061497a604084015160030b90565b60030b90565b6149b2608061499260608601516149ce565b940151946149a86149a161047a565b9889614907565b60070b6020880152565b60070b6040860152565b6060840152608083015260a082015290565b6149d66148d9565b5060208151910151602063ffffffff82511691015190604051916149f983610395565b8252602082015260405191614a0d83610395565b8252602082015290565b6040906060828051614a28816103d0565b614a30611079565b81528151614a3d816103b5565b8381526020820152015280515191610100610140606085015194015192825193614a66856103b5565b84525151015191815193614a79856103d0565b8452602084015282015290565b8051906040519060209282614a9e8582018093611355565b0392614ab2601f199485810183528261043d565b5190208451604051614ad781614acb8882018095611355565b0386810183528261043d565b51902003614b715782015182614af1825163ffffffff1690565b94019363ffffffff614b15614b0c875163ffffffff90511690565b63ffffffff1690565b911603614b71576101b0918380614b5f930151604051614b4881614b3c8582018095611355565b0385810183528261043d565b519020955101519360405193849182018096611355565b51902003614b6c57600190565b600090565b50505050600090565b15614b8157565b60405162461bcd60e51b815260206004820152601f60248201527f54656e6465726d696e743a20686173682063616e277420626520656d707479006044820152606490fd5b600e6000610d1792614bdf60e082510151511515614b7a565b614bea815151614d02565b90614bf9606082510151614fb6565b614c07608083510151614dee565b90614c10610487565b938452614c21602084510151614fcc565b6020850152614c3d614c3860408551015160070b90565b615036565b604085015260608401526080830152614c5a60a082510151614fcc565b60a0830152614c6d60c082510151614fcc565b60c0830152614c8060e082510151614fcc565b60e0830152610100614c958183510151614fcc565b90830152610120614ca98183510151614fcc565b90830152610140614cbd8183510151614fcc565b90830152610160614cd18183510151614fcc565b90830152610180614ce58183510151614fcc565b90830152614cf96101a08092510151614fcc565b90820152615a30565b6001600160401b0390614d1782825116615564565b91826001018060011161194c576020830193614d3583865116615564565b8060010160011161194c57600291010180911161194c57614d559061221f565b8093602092845116614db6575b6122f49350614d7b61023582516001600160401b031690565b614d86575050612183565b816113e491614d9f856113e461244596614daf98615444565b938491516001600160401b031690565b8238612395565b9190614dc1906153bc565b810180911161194c57614de8816113e48680946124456122f498516001600160401b031690565b91614d62565b614df981515161554f565b90816001018060011161194c576020820192614e1861229e8551614f8b565b8060010160011161194c57600291010180911161194c57614e389061221f565b91602091805151614e6b575b50614e4f8383615451565b820180921161194c57816113e4846122f4946124dc9451614e9d565b9091614e7684615435565b8082019283831161194c578584614e8d92516154a7565b010180911161194c579038614e44565b929192614eac6122e882614f8b565b9060209063ffffffff80825116614f56575b50602001805151614ee6575b5061250061250c61251193610d1796976124e161250695612183565b9491929093601291600094868101845b60071c948515614f18576080178153600196870196607f861695910190614ef6565b6125069550614f48612500959a6113e4856123b1610d179d600161250c9a9e6124e1986125119f9a530190611de8565b955050979650935050614eca565b9091614f61846153bc565b8082019283831161194c578584614f7a9287511661546b565b010180911161194c57906020614ebe565b614f9b63ffffffff825116615564565b80600101918260011161194c5760206126249101515161554f565b6122f4614fc56122e88361277d565b80926126c8565b614fd6815161554f565b6001018060011161194c57614fea9061221f565b906020908051615007575b50601f19810190811161194c57815290565b9061501183615435565b8082019283831161194c578484615027926154a7565b010180911161194c5738614ff5565b61503f816155a8565b600190810180821161194c576150549061221f565b91602091600782810b615076575b505050601f19810190811161194c57815290565b600894929493919392600092808201855b841c9586156150a95787918291608017815301940194607f8116959490615087565b9350919694509450536021810180931161194c576150d384846001600160401b036021951661546b565b010180911161194c57388080615062565b610d1790516000815191615661565b9290925b60209384841061512c578151815284810180911161194c5793810180911161194c5791601f19810190811161194c57916150f7565b9290919350600019906020036101000a0190811990511690825116179052565b906151569161516e565b9091600783169260068410156119a05760031c929190565b91909160009081936151808151611dda565b916080908201925b6080821660801461519b57505050509190565b80838598959796939497031015615239578694969391929351926000945b60209081871060808216608014161561522857508288880310156151fa575083851a9081607f16846007021b179260018080920196019601959492906151b9565b62461bcd60e51b600052600452600f6024526e6c656e677468206f766572666c6f7760881b60445260536000fd5b939897919695509093506151889050565b62461bcd60e51b6000526020600452600f6024526e6c656e677468206f766572666c6f7760881b60445260536000fd5b91906000908051936020850180951161194c576004948582019081831161194c576152989194939411156152d0565b8192019182515b8582106152b457505063ffffffff9150169190565b909181831a8360031b1b1792600180910192019092919261529f565b156152d757565b634e487b7160e01b600052600160045260246000fd5b91906000908051936020850180951161194c576008948582019081831161194c5761531c9194939411156152d0565b8192019182515b85821061533b5750506001600160401b039150169190565b909181831a8360031b1b17926001809101920190929192615323565b9190615363818461516e565b9361536d8261221f565b928051906020820180921161194c578683019182841161194c5784830180931161194c576153af936153a38992879511156152d0565b602087019201016150f3565b830180931161194c579190565b600890602060009101825b60071c9283156153ea57608017815360018091019101607f8316929190916153c7565b9060019350530190565b60009182910160116153ea565b60009182910160196153ea565b60009182910160226153ea565b600091829101602a6153ea565b60009182910160326153ea565b6000908190602001600a6153ea565b60009182910160106153ea565b60009182910160126153ea565b600091829101601a6153ea565b607f9392600092858316929101905b60071c91821561549b5760801781536001928301928583169291019061547a565b91506001939450530190565b908151916154b684838561546b565b93602060009186839501019201915b8484106154dd575050509050810180911161194c5790565b8251821a815360019384019392830192016154c5565b9063ffffffff610d1793921661546b565b906001600160401b03610d1793921661546b565b6001600160401b0316910160080190600891825b61553857505050600890565b6000198091019282816020031a845301918261552c565b61555881615564565b810180911161194c5790565b60018091600790811c805b6155795750505090565b92820192811c8061556f565b60008160030b126000146155995750600a90565b63ffffffff610d179116615564565b60008160070b126000146155bc5750600a90565b6001600160401b03610d179116615564565b637fffffff1981121580615653575b156155e85760030b90565b60405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203360448201527f32206269747300000000000000000000000000000000000000000000000000006064820152608490fd5b50637fffffff8113156155dd565b90918061567357505050610d176158df565b600181036156d4575061568c610d17926156cf92611967565b516000602060405161569d81610395565b6156a56112ad565b815201526040602082015191015160070b604051916156c383610395565b8252602082015261570f565b61591b565b6156dd816159a2565b926156e9848285615661565b9380820180921161194c57820391821161194c5761570692615661565b610d1791615966565b61571c61229e825161587b565b80600101908160011161194c576020918284019161573d835160070b6155a8565b8060010160011161194c57600291010180911161194c5761575d9061221f565b9261576784615435565b83019283811161194c57836113e46122f4956125116157e8946157bc8a9751916157936122e88461587b565b9281815151615858575b810184815151615837575b50506040018381515161581b575050612183565b9061250c826157de6157d86157d28c898561546b565b88611de8565b93611dda565b612506848c611de8565b916157f4815160070b90565b60070b615802575050612183565b816113e4916123d9856113e461244596614daf98615444565b806123b1846113e46113e4946158309761545e565b3883612395565b906113e4826123b1866113e46158509660409899615451565b9190846157a8565b915061587561586961242d86615435565b6113e4868285516154a7565b9161579d565b61588681515161554f565b9060018281019283821161194c576158a260208401515161554f565b90818301831161194c5701916002830180941161194c5760406158c79101515161554f565b908181011061194c57600391010180911161194c5790565b602060006159086040518281526158f5816103b5565b83604051928284809451938492016106f8565b8101039060025afa156110745760005190565b600061590860209260405161595560218287810194878652615945815180928b86860191016106f8565b810103600181018452018261043d565b6040519283928392519283916106f8565b6159086000916020936040519085820192600160f81b84526021830152604182015260418152615955816103eb565b801561194c576000190190565b6001808211156159eb578091600019810190811161194c575b8181116159c757505090565b6159d081612174565b8116156159e5576159e090615995565b6159bb565b91505090565b60405162461bcd60e51b815260206004820152601960248201527f4d65726b6c65547265653a20696e76616c696420696e707574000000000000006044820152606490fd5b909180615a4257505050610d176158df565b60018103615a625750600e82101561197b57610d179160051b015161591b565b615a6b816159a2565b92615a77848285615a30565b9380820180921161194c57820391821161194c5761570692615a3056fea2646970667358221220bd3a69dc8547430d59d67b38cfff9ac2172097c68198ccb0d563cd9ede3a3ffd64736f6c63430008120033",
}

// TendermintLightClientABI is the input ABI used to generate the binding from.
// Deprecated: Use TendermintLightClientMetaData.ABI instead.
var TendermintLightClientABI = TendermintLightClientMetaData.ABI

// TendermintLightClientBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TendermintLightClientMetaData.Bin instead.
var TendermintLightClientBin = TendermintLightClientMetaData.Bin

// DeployTendermintLightClient deploys a new Ethereum contract, binding an instance of TendermintLightClient to it.
func DeployTendermintLightClient(auth *bind.TransactOpts, backend bind.ContractBackend, _ed25519Verifier common.Address) (common.Address, *types.Transaction, *TendermintLightClient, error) {
	parsed, err := TendermintLightClientMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TendermintLightClientBin), backend, _ed25519Verifier)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TendermintLightClient{TendermintLightClientCaller: TendermintLightClientCaller{contract: contract}, TendermintLightClientTransactor: TendermintLightClientTransactor{contract: contract}, TendermintLightClientFilterer: TendermintLightClientFilterer{contract: contract}}, nil
}

// TendermintLightClient is an auto generated Go binding around an Ethereum contract.
type TendermintLightClient struct {
	TendermintLightClientCaller     // Read-only binding to the contract
	TendermintLightClientTransactor // Write-only binding to the contract
	TendermintLightClientFilterer   // Log filterer for contract events
}

// TendermintLightClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type TendermintLightClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TendermintLightClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TendermintLightClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TendermintLightClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TendermintLightClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TendermintLightClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TendermintLightClientSession struct {
	Contract     *TendermintLightClient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TendermintLightClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TendermintLightClientCallerSession struct {
	Contract *TendermintLightClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// TendermintLightClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TendermintLightClientTransactorSession struct {
	Contract     *TendermintLightClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// TendermintLightClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type TendermintLightClientRaw struct {
	Contract *TendermintLightClient // Generic contract binding to access the raw methods on
}

// TendermintLightClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TendermintLightClientCallerRaw struct {
	Contract *TendermintLightClientCaller // Generic read-only contract binding to access the raw methods on
}

// TendermintLightClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TendermintLightClientTransactorRaw struct {
	Contract *TendermintLightClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTendermintLightClient creates a new instance of TendermintLightClient, bound to a specific deployed contract.
func NewTendermintLightClient(address common.Address, backend bind.ContractBackend) (*TendermintLightClient, error) {
	contract, err := bindTendermintLightClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TendermintLightClient{TendermintLightClientCaller: TendermintLightClientCaller{contract: contract}, TendermintLightClientTransactor: TendermintLightClientTransactor{contract: contract}, TendermintLightClientFilterer: TendermintLightClientFilterer{contract: contract}}, nil
}

// NewTendermintLightClientCaller creates a new read-only instance of TendermintLightClient, bound to a specific deployed contract.
func NewTendermintLightClientCaller(address common.Address, caller bind.ContractCaller) (*TendermintLightClientCaller, error) {
	contract, err := bindTendermintLightClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TendermintLightClientCaller{contract: contract}, nil
}

// NewTendermintLightClientTransactor creates a new write-only instance of TendermintLightClient, bound to a specific deployed contract.
func NewTendermintLightClientTransactor(address common.Address, transactor bind.ContractTransactor) (*TendermintLightClientTransactor, error) {
	contract, err := bindTendermintLightClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TendermintLightClientTransactor{contract: contract}, nil
}

// NewTendermintLightClientFilterer creates a new log filterer instance of TendermintLightClient, bound to a specific deployed contract.
func NewTendermintLightClientFilterer(address common.Address, filterer bind.ContractFilterer) (*TendermintLightClientFilterer, error) {
	contract, err := bindTendermintLightClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TendermintLightClientFilterer{contract: contract}, nil
}

// bindTendermintLightClient binds a generic wrapper to an already deployed contract.
func bindTendermintLightClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TendermintLightClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TendermintLightClient *TendermintLightClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TendermintLightClient.Contract.TendermintLightClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TendermintLightClient *TendermintLightClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TendermintLightClient.Contract.TendermintLightClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TendermintLightClient *TendermintLightClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TendermintLightClient.Contract.TendermintLightClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TendermintLightClient *TendermintLightClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TendermintLightClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TendermintLightClient *TendermintLightClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TendermintLightClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TendermintLightClient *TendermintLightClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TendermintLightClient.Contract.contract.Transact(opts, method, params...)
}

// ConsensusStates is a free data retrieval call binding the contract method 0xe08322d5.
//
// Solidity: function consensusStates(uint64 ) view returns((int64,int32) timestamp, (bytes) root, bytes next_validators_hash)
func (_TendermintLightClient *TendermintLightClientCaller) ConsensusStates(opts *bind.CallOpts, arg0 uint64) (struct {
	Timestamp          TimestampData
	Root               MerkleRootData
	NextValidatorsHash []byte
}, error) {
	var out []interface{}
	err := _TendermintLightClient.contract.Call(opts, &out, "consensusStates", arg0)

	outstruct := new(struct {
		Timestamp          TimestampData
		Root               MerkleRootData
		NextValidatorsHash []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(TimestampData)).(*TimestampData)
	outstruct.Root = *abi.ConvertType(out[1], new(MerkleRootData)).(*MerkleRootData)
	outstruct.NextValidatorsHash = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// ConsensusStates is a free data retrieval call binding the contract method 0xe08322d5.
//
// Solidity: function consensusStates(uint64 ) view returns((int64,int32) timestamp, (bytes) root, bytes next_validators_hash)
func (_TendermintLightClient *TendermintLightClientSession) ConsensusStates(arg0 uint64) (struct {
	Timestamp          TimestampData
	Root               MerkleRootData
	NextValidatorsHash []byte
}, error) {
	return _TendermintLightClient.Contract.ConsensusStates(&_TendermintLightClient.CallOpts, arg0)
}

// ConsensusStates is a free data retrieval call binding the contract method 0xe08322d5.
//
// Solidity: function consensusStates(uint64 ) view returns((int64,int32) timestamp, (bytes) root, bytes next_validators_hash)
func (_TendermintLightClient *TendermintLightClientCallerSession) ConsensusStates(arg0 uint64) (struct {
	Timestamp          TimestampData
	Root               MerkleRootData
	NextValidatorsHash []byte
}, error) {
	return _TendermintLightClient.Contract.ConsensusStates(&_TendermintLightClient.CallOpts, arg0)
}

// GetAppHash is a free data retrieval call binding the contract method 0xcba510a9.
//
// Solidity: function getAppHash(uint64 height) view returns(bytes32)
func (_TendermintLightClient *TendermintLightClientCaller) GetAppHash(opts *bind.CallOpts, height uint64) ([32]byte, error) {
	var out []interface{}
	err := _TendermintLightClient.contract.Call(opts, &out, "getAppHash", height)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAppHash is a free data retrieval call binding the contract method 0xcba510a9.
//
// Solidity: function getAppHash(uint64 height) view returns(bytes32)
func (_TendermintLightClient *TendermintLightClientSession) GetAppHash(height uint64) ([32]byte, error) {
	return _TendermintLightClient.Contract.GetAppHash(&_TendermintLightClient.CallOpts, height)
}

// GetAppHash is a free data retrieval call binding the contract method 0xcba510a9.
//
// Solidity: function getAppHash(uint64 height) view returns(bytes32)
func (_TendermintLightClient *TendermintLightClientCallerSession) GetAppHash(height uint64) ([32]byte, error) {
	return _TendermintLightClient.Contract.GetAppHash(&_TendermintLightClient.CallOpts, height)
}

// InitialHeight is a free data retrieval call binding the contract method 0xe2761af0.
//
// Solidity: function initialHeight() view returns(uint64)
func (_TendermintLightClient *TendermintLightClientCaller) InitialHeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TendermintLightClient.contract.Call(opts, &out, "initialHeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InitialHeight is a free data retrieval call binding the contract method 0xe2761af0.
//
// Solidity: function initialHeight() view returns(uint64)
func (_TendermintLightClient *TendermintLightClientSession) InitialHeight() (uint64, error) {
	return _TendermintLightClient.Contract.InitialHeight(&_TendermintLightClient.CallOpts)
}

// InitialHeight is a free data retrieval call binding the contract method 0xe2761af0.
//
// Solidity: function initialHeight() view returns(uint64)
func (_TendermintLightClient *TendermintLightClientCallerSession) InitialHeight() (uint64, error) {
	return _TendermintLightClient.Contract.InitialHeight(&_TendermintLightClient.CallOpts)
}

// IsHeaderSynced is a free data retrieval call binding the contract method 0xdf5fe704.
//
// Solidity: function isHeaderSynced(uint64 height) view returns(bool)
func (_TendermintLightClient *TendermintLightClientCaller) IsHeaderSynced(opts *bind.CallOpts, height uint64) (bool, error) {
	var out []interface{}
	err := _TendermintLightClient.contract.Call(opts, &out, "isHeaderSynced", height)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHeaderSynced is a free data retrieval call binding the contract method 0xdf5fe704.
//
// Solidity: function isHeaderSynced(uint64 height) view returns(bool)
func (_TendermintLightClient *TendermintLightClientSession) IsHeaderSynced(height uint64) (bool, error) {
	return _TendermintLightClient.Contract.IsHeaderSynced(&_TendermintLightClient.CallOpts, height)
}

// IsHeaderSynced is a free data retrieval call binding the contract method 0xdf5fe704.
//
// Solidity: function isHeaderSynced(uint64 height) view returns(bool)
func (_TendermintLightClient *TendermintLightClientCallerSession) IsHeaderSynced(height uint64) (bool, error) {
	return _TendermintLightClient.Contract.IsHeaderSynced(&_TendermintLightClient.CallOpts, height)
}

// LatestHeight is a free data retrieval call binding the contract method 0xe405bbc3.
//
// Solidity: function latestHeight() view returns(uint64)
func (_TendermintLightClient *TendermintLightClientCaller) LatestHeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TendermintLightClient.contract.Call(opts, &out, "latestHeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LatestHeight is a free data retrieval call binding the contract method 0xe405bbc3.
//
// Solidity: function latestHeight() view returns(uint64)
func (_TendermintLightClient *TendermintLightClientSession) LatestHeight() (uint64, error) {
	return _TendermintLightClient.Contract.LatestHeight(&_TendermintLightClient.CallOpts)
}

// LatestHeight is a free data retrieval call binding the contract method 0xe405bbc3.
//
// Solidity: function latestHeight() view returns(uint64)
func (_TendermintLightClient *TendermintLightClientCallerSession) LatestHeight() (uint64, error) {
	return _TendermintLightClient.Contract.LatestHeight(&_TendermintLightClient.CallOpts)
}

// Synced is a free data retrieval call binding the contract method 0xfacb618d.
//
// Solidity: function synced(uint64 ) view returns(bool)
func (_TendermintLightClient *TendermintLightClientCaller) Synced(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _TendermintLightClient.contract.Call(opts, &out, "synced", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Synced is a free data retrieval call binding the contract method 0xfacb618d.
//
// Solidity: function synced(uint64 ) view returns(bool)
func (_TendermintLightClient *TendermintLightClientSession) Synced(arg0 uint64) (bool, error) {
	return _TendermintLightClient.Contract.Synced(&_TendermintLightClient.CallOpts, arg0)
}

// Synced is a free data retrieval call binding the contract method 0xfacb618d.
//
// Solidity: function synced(uint64 ) view returns(bool)
func (_TendermintLightClient *TendermintLightClientCallerSession) Synced(arg0 uint64) (bool, error) {
	return _TendermintLightClient.Contract.Synced(&_TendermintLightClient.CallOpts, arg0)
}

// Init is a paid mutator transaction binding the contract method 0xc0d91eaf.
//
// Solidity: function init(address _system, bytes _initHeader) returns()
func (_TendermintLightClient *TendermintLightClientTransactor) Init(opts *bind.TransactOpts, _system common.Address, _initHeader []byte) (*types.Transaction, error) {
	return _TendermintLightClient.contract.Transact(opts, "init", _system, _initHeader)
}

// Init is a paid mutator transaction binding the contract method 0xc0d91eaf.
//
// Solidity: function init(address _system, bytes _initHeader) returns()
func (_TendermintLightClient *TendermintLightClientSession) Init(_system common.Address, _initHeader []byte) (*types.Transaction, error) {
	return _TendermintLightClient.Contract.Init(&_TendermintLightClient.TransactOpts, _system, _initHeader)
}

// Init is a paid mutator transaction binding the contract method 0xc0d91eaf.
//
// Solidity: function init(address _system, bytes _initHeader) returns()
func (_TendermintLightClient *TendermintLightClientTransactorSession) Init(_system common.Address, _initHeader []byte) (*types.Transaction, error) {
	return _TendermintLightClient.Contract.Init(&_TendermintLightClient.TransactOpts, _system, _initHeader)
}

// SyncTendermintHeader is a paid mutator transaction binding the contract method 0xead3cbe7.
//
// Solidity: function syncTendermintHeader(bytes header, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC, uint256[2] proofCommit, uint256 proofCommitPub) returns(bool)
func (_TendermintLightClient *TendermintLightClientTransactor) SyncTendermintHeader(opts *bind.TransactOpts, header []byte, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int, proofCommit [2]*big.Int, proofCommitPub *big.Int) (*types.Transaction, error) {
	return _TendermintLightClient.contract.Transact(opts, "syncTendermintHeader", header, proofA, proofB, proofC, proofCommit, proofCommitPub)
}

// SyncTendermintHeader is a paid mutator transaction binding the contract method 0xead3cbe7.
//
// Solidity: function syncTendermintHeader(bytes header, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC, uint256[2] proofCommit, uint256 proofCommitPub) returns(bool)
func (_TendermintLightClient *TendermintLightClientSession) SyncTendermintHeader(header []byte, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int, proofCommit [2]*big.Int, proofCommitPub *big.Int) (*types.Transaction, error) {
	return _TendermintLightClient.Contract.SyncTendermintHeader(&_TendermintLightClient.TransactOpts, header, proofA, proofB, proofC, proofCommit, proofCommitPub)
}

// SyncTendermintHeader is a paid mutator transaction binding the contract method 0xead3cbe7.
//
// Solidity: function syncTendermintHeader(bytes header, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC, uint256[2] proofCommit, uint256 proofCommitPub) returns(bool)
func (_TendermintLightClient *TendermintLightClientTransactorSession) SyncTendermintHeader(header []byte, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int, proofCommit [2]*big.Int, proofCommitPub *big.Int) (*types.Transaction, error) {
	return _TendermintLightClient.Contract.SyncTendermintHeader(&_TendermintLightClient.TransactOpts, header, proofA, proofB, proofC, proofCommit, proofCommitPub)
}

// TendermintLightClientConsensusStateInitIterator is returned from FilterConsensusStateInit and is used to iterate over the raw logs and unpacked data for ConsensusStateInit events raised by the TendermintLightClient contract.
type TendermintLightClientConsensusStateInitIterator struct {
	Event *TendermintLightClientConsensusStateInit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TendermintLightClientConsensusStateInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TendermintLightClientConsensusStateInit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TendermintLightClientConsensusStateInit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TendermintLightClientConsensusStateInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TendermintLightClientConsensusStateInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TendermintLightClientConsensusStateInit represents a ConsensusStateInit event raised by the TendermintLightClient contract.
type TendermintLightClientConsensusStateInit struct {
	InitialHeight uint64
	AppHash       [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConsensusStateInit is a free log retrieval operation binding the contract event 0xa004c198518af96775c486676e922051a150cf22b054780d1223d312127b72f1.
//
// Solidity: event ConsensusStateInit(uint64 initialHeight, bytes32 appHash)
func (_TendermintLightClient *TendermintLightClientFilterer) FilterConsensusStateInit(opts *bind.FilterOpts) (*TendermintLightClientConsensusStateInitIterator, error) {

	logs, sub, err := _TendermintLightClient.contract.FilterLogs(opts, "ConsensusStateInit")
	if err != nil {
		return nil, err
	}
	return &TendermintLightClientConsensusStateInitIterator{contract: _TendermintLightClient.contract, event: "ConsensusStateInit", logs: logs, sub: sub}, nil
}

// WatchConsensusStateInit is a free log subscription operation binding the contract event 0xa004c198518af96775c486676e922051a150cf22b054780d1223d312127b72f1.
//
// Solidity: event ConsensusStateInit(uint64 initialHeight, bytes32 appHash)
func (_TendermintLightClient *TendermintLightClientFilterer) WatchConsensusStateInit(opts *bind.WatchOpts, sink chan<- *TendermintLightClientConsensusStateInit) (event.Subscription, error) {

	logs, sub, err := _TendermintLightClient.contract.WatchLogs(opts, "ConsensusStateInit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TendermintLightClientConsensusStateInit)
				if err := _TendermintLightClient.contract.UnpackLog(event, "ConsensusStateInit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConsensusStateInit is a log parse operation binding the contract event 0xa004c198518af96775c486676e922051a150cf22b054780d1223d312127b72f1.
//
// Solidity: event ConsensusStateInit(uint64 initialHeight, bytes32 appHash)
func (_TendermintLightClient *TendermintLightClientFilterer) ParseConsensusStateInit(log types.Log) (*TendermintLightClientConsensusStateInit, error) {
	event := new(TendermintLightClientConsensusStateInit)
	if err := _TendermintLightClient.contract.UnpackLog(event, "ConsensusStateInit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TendermintLightClientConsensusStateSyncedIterator is returned from FilterConsensusStateSynced and is used to iterate over the raw logs and unpacked data for ConsensusStateSynced events raised by the TendermintLightClient contract.
type TendermintLightClientConsensusStateSyncedIterator struct {
	Event *TendermintLightClientConsensusStateSynced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TendermintLightClientConsensusStateSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TendermintLightClientConsensusStateSynced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TendermintLightClientConsensusStateSynced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TendermintLightClientConsensusStateSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TendermintLightClientConsensusStateSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TendermintLightClientConsensusStateSynced represents a ConsensusStateSynced event raised by the TendermintLightClient contract.
type TendermintLightClientConsensusStateSynced struct {
	Height  uint64
	AppHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConsensusStateSynced is a free log retrieval operation binding the contract event 0x82639a39e1f497d8456613d012ef21caf479c800899e97b68647432032a41ed9.
//
// Solidity: event ConsensusStateSynced(uint64 height, bytes32 appHash)
func (_TendermintLightClient *TendermintLightClientFilterer) FilterConsensusStateSynced(opts *bind.FilterOpts) (*TendermintLightClientConsensusStateSyncedIterator, error) {

	logs, sub, err := _TendermintLightClient.contract.FilterLogs(opts, "ConsensusStateSynced")
	if err != nil {
		return nil, err
	}
	return &TendermintLightClientConsensusStateSyncedIterator{contract: _TendermintLightClient.contract, event: "ConsensusStateSynced", logs: logs, sub: sub}, nil
}

// WatchConsensusStateSynced is a free log subscription operation binding the contract event 0x82639a39e1f497d8456613d012ef21caf479c800899e97b68647432032a41ed9.
//
// Solidity: event ConsensusStateSynced(uint64 height, bytes32 appHash)
func (_TendermintLightClient *TendermintLightClientFilterer) WatchConsensusStateSynced(opts *bind.WatchOpts, sink chan<- *TendermintLightClientConsensusStateSynced) (event.Subscription, error) {

	logs, sub, err := _TendermintLightClient.contract.WatchLogs(opts, "ConsensusStateSynced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TendermintLightClientConsensusStateSynced)
				if err := _TendermintLightClient.contract.UnpackLog(event, "ConsensusStateSynced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConsensusStateSynced is a log parse operation binding the contract event 0x82639a39e1f497d8456613d012ef21caf479c800899e97b68647432032a41ed9.
//
// Solidity: event ConsensusStateSynced(uint64 height, bytes32 appHash)
func (_TendermintLightClient *TendermintLightClientFilterer) ParseConsensusStateSynced(log types.Log) (*TendermintLightClientConsensusStateSynced, error) {
	event := new(TendermintLightClientConsensusStateSynced)
	if err := _TendermintLightClient.contract.UnpackLog(event, "ConsensusStateSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimestampMetaData contains all meta data concerning the Timestamp contract.
var TimestampMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212201bf357151a7a5be980dfaa72ac34f6f9d73067dd922864ff0bb2f8187379de3464736f6c63430008120033",
}

// TimestampABI is the input ABI used to generate the binding from.
// Deprecated: Use TimestampMetaData.ABI instead.
var TimestampABI = TimestampMetaData.ABI

// TimestampBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TimestampMetaData.Bin instead.
var TimestampBin = TimestampMetaData.Bin

// DeployTimestamp deploys a new Ethereum contract, binding an instance of Timestamp to it.
func DeployTimestamp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Timestamp, error) {
	parsed, err := TimestampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TimestampBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Timestamp{TimestampCaller: TimestampCaller{contract: contract}, TimestampTransactor: TimestampTransactor{contract: contract}, TimestampFilterer: TimestampFilterer{contract: contract}}, nil
}

// Timestamp is an auto generated Go binding around an Ethereum contract.
type Timestamp struct {
	TimestampCaller     // Read-only binding to the contract
	TimestampTransactor // Write-only binding to the contract
	TimestampFilterer   // Log filterer for contract events
}

// TimestampCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimestampCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimestampTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimestampTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimestampFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimestampFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimestampSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimestampSession struct {
	Contract     *Timestamp        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimestampCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimestampCallerSession struct {
	Contract *TimestampCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TimestampTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimestampTransactorSession struct {
	Contract     *TimestampTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TimestampRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimestampRaw struct {
	Contract *Timestamp // Generic contract binding to access the raw methods on
}

// TimestampCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimestampCallerRaw struct {
	Contract *TimestampCaller // Generic read-only contract binding to access the raw methods on
}

// TimestampTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimestampTransactorRaw struct {
	Contract *TimestampTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimestamp creates a new instance of Timestamp, bound to a specific deployed contract.
func NewTimestamp(address common.Address, backend bind.ContractBackend) (*Timestamp, error) {
	contract, err := bindTimestamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Timestamp{TimestampCaller: TimestampCaller{contract: contract}, TimestampTransactor: TimestampTransactor{contract: contract}, TimestampFilterer: TimestampFilterer{contract: contract}}, nil
}

// NewTimestampCaller creates a new read-only instance of Timestamp, bound to a specific deployed contract.
func NewTimestampCaller(address common.Address, caller bind.ContractCaller) (*TimestampCaller, error) {
	contract, err := bindTimestamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimestampCaller{contract: contract}, nil
}

// NewTimestampTransactor creates a new write-only instance of Timestamp, bound to a specific deployed contract.
func NewTimestampTransactor(address common.Address, transactor bind.ContractTransactor) (*TimestampTransactor, error) {
	contract, err := bindTimestamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimestampTransactor{contract: contract}, nil
}

// NewTimestampFilterer creates a new log filterer instance of Timestamp, bound to a specific deployed contract.
func NewTimestampFilterer(address common.Address, filterer bind.ContractFilterer) (*TimestampFilterer, error) {
	contract, err := bindTimestamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimestampFilterer{contract: contract}, nil
}

// bindTimestamp binds a generic wrapper to an already deployed contract.
func bindTimestamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TimestampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Timestamp *TimestampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Timestamp.Contract.TimestampCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Timestamp *TimestampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Timestamp.Contract.TimestampTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Timestamp *TimestampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Timestamp.Contract.TimestampTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Timestamp *TimestampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Timestamp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Timestamp *TimestampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Timestamp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Timestamp *TimestampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Timestamp.Contract.contract.Transact(opts, method, params...)
}

// TmHeaderMetaData contains all meta data concerning the TmHeader contract.
var TmHeaderMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212204ae09af4dfb78939812edd30711a0098fea5c775a14bea39e9ff431c535de77364736f6c63430008120033",
}

// TmHeaderABI is the input ABI used to generate the binding from.
// Deprecated: Use TmHeaderMetaData.ABI instead.
var TmHeaderABI = TmHeaderMetaData.ABI

// TmHeaderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TmHeaderMetaData.Bin instead.
var TmHeaderBin = TmHeaderMetaData.Bin

// DeployTmHeader deploys a new Ethereum contract, binding an instance of TmHeader to it.
func DeployTmHeader(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TmHeader, error) {
	parsed, err := TmHeaderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TmHeaderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TmHeader{TmHeaderCaller: TmHeaderCaller{contract: contract}, TmHeaderTransactor: TmHeaderTransactor{contract: contract}, TmHeaderFilterer: TmHeaderFilterer{contract: contract}}, nil
}

// TmHeader is an auto generated Go binding around an Ethereum contract.
type TmHeader struct {
	TmHeaderCaller     // Read-only binding to the contract
	TmHeaderTransactor // Write-only binding to the contract
	TmHeaderFilterer   // Log filterer for contract events
}

// TmHeaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type TmHeaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TmHeaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TmHeaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TmHeaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TmHeaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TmHeaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TmHeaderSession struct {
	Contract     *TmHeader         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TmHeaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TmHeaderCallerSession struct {
	Contract *TmHeaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TmHeaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TmHeaderTransactorSession struct {
	Contract     *TmHeaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TmHeaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type TmHeaderRaw struct {
	Contract *TmHeader // Generic contract binding to access the raw methods on
}

// TmHeaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TmHeaderCallerRaw struct {
	Contract *TmHeaderCaller // Generic read-only contract binding to access the raw methods on
}

// TmHeaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TmHeaderTransactorRaw struct {
	Contract *TmHeaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTmHeader creates a new instance of TmHeader, bound to a specific deployed contract.
func NewTmHeader(address common.Address, backend bind.ContractBackend) (*TmHeader, error) {
	contract, err := bindTmHeader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TmHeader{TmHeaderCaller: TmHeaderCaller{contract: contract}, TmHeaderTransactor: TmHeaderTransactor{contract: contract}, TmHeaderFilterer: TmHeaderFilterer{contract: contract}}, nil
}

// NewTmHeaderCaller creates a new read-only instance of TmHeader, bound to a specific deployed contract.
func NewTmHeaderCaller(address common.Address, caller bind.ContractCaller) (*TmHeaderCaller, error) {
	contract, err := bindTmHeader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TmHeaderCaller{contract: contract}, nil
}

// NewTmHeaderTransactor creates a new write-only instance of TmHeader, bound to a specific deployed contract.
func NewTmHeaderTransactor(address common.Address, transactor bind.ContractTransactor) (*TmHeaderTransactor, error) {
	contract, err := bindTmHeader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TmHeaderTransactor{contract: contract}, nil
}

// NewTmHeaderFilterer creates a new log filterer instance of TmHeader, bound to a specific deployed contract.
func NewTmHeaderFilterer(address common.Address, filterer bind.ContractFilterer) (*TmHeaderFilterer, error) {
	contract, err := bindTmHeader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TmHeaderFilterer{contract: contract}, nil
}

// bindTmHeader binds a generic wrapper to an already deployed contract.
func bindTmHeader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TmHeaderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TmHeader *TmHeaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TmHeader.Contract.TmHeaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TmHeader *TmHeaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TmHeader.Contract.TmHeaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TmHeader *TmHeaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TmHeader.Contract.TmHeaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TmHeader *TmHeaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TmHeader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TmHeader *TmHeaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TmHeader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TmHeader *TmHeaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TmHeader.Contract.contract.Transact(opts, method, params...)
}

// ValidatorMetaData contains all meta data concerning the Validator contract.
var ValidatorMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220c3caf3efea5c2ffc259ac676ee4edc75690f05a2a2dc69f8296e1fdd22468a3564736f6c63430008120033",
}

// ValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorMetaData.ABI instead.
var ValidatorABI = ValidatorMetaData.ABI

// ValidatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorMetaData.Bin instead.
var ValidatorBin = ValidatorMetaData.Bin

// DeployValidator deploys a new Ethereum contract, binding an instance of Validator to it.
func DeployValidator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Validator, error) {
	parsed, err := ValidatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// Validator is an auto generated Go binding around an Ethereum contract.
type Validator struct {
	ValidatorCaller     // Read-only binding to the contract
	ValidatorTransactor // Write-only binding to the contract
	ValidatorFilterer   // Log filterer for contract events
}

// ValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSession struct {
	Contract     *Validator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorCallerSession struct {
	Contract *ValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorTransactorSession struct {
	Contract     *ValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorRaw struct {
	Contract *Validator // Generic contract binding to access the raw methods on
}

// ValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorCallerRaw struct {
	Contract *ValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorTransactorRaw struct {
	Contract *ValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidator creates a new instance of Validator, bound to a specific deployed contract.
func NewValidator(address common.Address, backend bind.ContractBackend) (*Validator, error) {
	contract, err := bindValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// NewValidatorCaller creates a new read-only instance of Validator, bound to a specific deployed contract.
func NewValidatorCaller(address common.Address, caller bind.ContractCaller) (*ValidatorCaller, error) {
	contract, err := bindValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorCaller{contract: contract}, nil
}

// NewValidatorTransactor creates a new write-only instance of Validator, bound to a specific deployed contract.
func NewValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorTransactor, error) {
	contract, err := bindValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorTransactor{contract: contract}, nil
}

// NewValidatorFilterer creates a new log filterer instance of Validator, bound to a specific deployed contract.
func NewValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorFilterer, error) {
	contract, err := bindValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorFilterer{contract: contract}, nil
}

// bindValidator binds a generic wrapper to an already deployed contract.
func bindValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.ValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transact(opts, method, params...)
}

// ValidatorSetMetaData contains all meta data concerning the ValidatorSet contract.
var ValidatorSetMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212201214b4952e84589314147c1de1d743eb8b06cac95d82c3f1546d6b188514bdb664736f6c63430008120033",
}

// ValidatorSetABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorSetMetaData.ABI instead.
var ValidatorSetABI = ValidatorSetMetaData.ABI

// ValidatorSetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorSetMetaData.Bin instead.
var ValidatorSetBin = ValidatorSetMetaData.Bin

// DeployValidatorSet deploys a new Ethereum contract, binding an instance of ValidatorSet to it.
func DeployValidatorSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValidatorSet, error) {
	parsed, err := ValidatorSetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorSet{ValidatorSetCaller: ValidatorSetCaller{contract: contract}, ValidatorSetTransactor: ValidatorSetTransactor{contract: contract}, ValidatorSetFilterer: ValidatorSetFilterer{contract: contract}}, nil
}

// ValidatorSet is an auto generated Go binding around an Ethereum contract.
type ValidatorSet struct {
	ValidatorSetCaller     // Read-only binding to the contract
	ValidatorSetTransactor // Write-only binding to the contract
	ValidatorSetFilterer   // Log filterer for contract events
}

// ValidatorSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSetSession struct {
	Contract     *ValidatorSet     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorSetCallerSession struct {
	Contract *ValidatorSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ValidatorSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorSetTransactorSession struct {
	Contract     *ValidatorSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ValidatorSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorSetRaw struct {
	Contract *ValidatorSet // Generic contract binding to access the raw methods on
}

// ValidatorSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorSetCallerRaw struct {
	Contract *ValidatorSetCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorSetTransactorRaw struct {
	Contract *ValidatorSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorSet creates a new instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSet(address common.Address, backend bind.ContractBackend) (*ValidatorSet, error) {
	contract, err := bindValidatorSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorSet{ValidatorSetCaller: ValidatorSetCaller{contract: contract}, ValidatorSetTransactor: ValidatorSetTransactor{contract: contract}, ValidatorSetFilterer: ValidatorSetFilterer{contract: contract}}, nil
}

// NewValidatorSetCaller creates a new read-only instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSetCaller(address common.Address, caller bind.ContractCaller) (*ValidatorSetCaller, error) {
	contract, err := bindValidatorSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetCaller{contract: contract}, nil
}

// NewValidatorSetTransactor creates a new write-only instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSetTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorSetTransactor, error) {
	contract, err := bindValidatorSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetTransactor{contract: contract}, nil
}

// NewValidatorSetFilterer creates a new log filterer instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSetFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorSetFilterer, error) {
	contract, err := bindValidatorSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetFilterer{contract: contract}, nil
}

// bindValidatorSet binds a generic wrapper to an already deployed contract.
func bindValidatorSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorSetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorSet *ValidatorSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorSet.Contract.ValidatorSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorSet *ValidatorSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSet.Contract.ValidatorSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorSet *ValidatorSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorSet.Contract.ValidatorSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorSet *ValidatorSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorSet *ValidatorSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorSet *ValidatorSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorSet.Contract.contract.Transact(opts, method, params...)
}

// VoteMetaData contains all meta data concerning the Vote contract.
var VoteMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220c8a1e49bd52e12a52abfa945d93c4df9b3a242df88b3e33b633c801037155fea64736f6c63430008120033",
}

// VoteABI is the input ABI used to generate the binding from.
// Deprecated: Use VoteMetaData.ABI instead.
var VoteABI = VoteMetaData.ABI

// VoteBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VoteMetaData.Bin instead.
var VoteBin = VoteMetaData.Bin

// DeployVote deploys a new Ethereum contract, binding an instance of Vote to it.
func DeployVote(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vote, error) {
	parsed, err := VoteMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VoteBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// Vote is an auto generated Go binding around an Ethereum contract.
type Vote struct {
	VoteCaller     // Read-only binding to the contract
	VoteTransactor // Write-only binding to the contract
	VoteFilterer   // Log filterer for contract events
}

// VoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoteSession struct {
	Contract     *Vote             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoteCallerSession struct {
	Contract *VoteCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoteTransactorSession struct {
	Contract     *VoteTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoteRaw struct {
	Contract *Vote // Generic contract binding to access the raw methods on
}

// VoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoteCallerRaw struct {
	Contract *VoteCaller // Generic read-only contract binding to access the raw methods on
}

// VoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoteTransactorRaw struct {
	Contract *VoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVote creates a new instance of Vote, bound to a specific deployed contract.
func NewVote(address common.Address, backend bind.ContractBackend) (*Vote, error) {
	contract, err := bindVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// NewVoteCaller creates a new read-only instance of Vote, bound to a specific deployed contract.
func NewVoteCaller(address common.Address, caller bind.ContractCaller) (*VoteCaller, error) {
	contract, err := bindVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoteCaller{contract: contract}, nil
}

// NewVoteTransactor creates a new write-only instance of Vote, bound to a specific deployed contract.
func NewVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*VoteTransactor, error) {
	contract, err := bindVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoteTransactor{contract: contract}, nil
}

// NewVoteFilterer creates a new log filterer instance of Vote, bound to a specific deployed contract.
func NewVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*VoteFilterer, error) {
	contract, err := bindVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoteFilterer{contract: contract}, nil
}

// bindVote binds a generic wrapper to an already deployed contract.
func bindVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VoteMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.VoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transact(opts, method, params...)
}
