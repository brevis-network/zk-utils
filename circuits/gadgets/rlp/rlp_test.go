package rlp

import (
	"fmt"
	"testing"

	"github.com/celer-network/zk-utils/circuits/gadgets/conv"
	emu "github.com/celer-network/zk-utils/circuits/gadgets/emulated"
	"github.com/celer-network/zk-utils/circuits/gadgets/keccak"
	"github.com/celer-network/zk-utils/circuits/gadgets/utils"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/emulated/sw_emulated"
	"github.com/consensys/gnark/std/evmprecompiles"
	"github.com/consensys/gnark/std/math/emulated"
	"github.com/consensys/gnark/test"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
)

type LegacyCircuit struct {
	Out frontend.Variable `gnark:",public"`
}

func (c *LegacyCircuit) Define(api frontend.API) error {
	raw := "0xf9034b57851f3305c0d383015fe694090d4613473dee047c3f2706764f49e0821d256e80b902e42e7ba6ef000000000000000000000000000000000000000000000000000000000003d289000000000000000000000000fcaaf32bde89a81c71a8c2fb2c7bc8358c9e3696000000000000000000000000000000000000000000000015af1d78b58c4000000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000001297c61438bfe2221b7921320373b938d75afd5ffaad3c2c1c8600f821d1d2c2c4c5a63314f65054e549380d124e4ae43943aac8e1270c5ba75f51c65609c0c86e795e3009b789ea2c7cfafe90e0f6824e905c2c7579546e9d9804d1d6ce0f145d1bd962a0772320698558459716665860eb6ae2de865b0dd20b3fb86a24f48cd53e7ac5b3b0ee8f655d76985cd283fe2ccffb4c577752370503536afdc489a198ae6bb8589e780eb5066050f0d044450b79591fd399fc6ed18614b828da53e439798ad925730eed5b27309c7077a8b8ed81c132e98df8aaff09f5607848353a58c796657b78487c195c7c8d89a1c717682ae2c26ae4f91914e56db784a9d15ee9b981b8e9f05a09f0d1f2bd32f9dbd332e9a8d9615ff9fea154613e97523d59edc9513bb2603e9c6954d7f096c4b5dfc30e29824196f3643e16cf152badbfc813125e9ee5400a5ce7921109d8d9631a2e73515d25bf34fd01f266465389e4f212b8ea0a0e5d463659621029622a8b24638adb1d2c7fe784b720196075cb06ac21cccb879f5a8c8aebc9df8499c3503c073d660e49a72404baa353821b2823c122c2b0846e905cb2cdabdabf7037f329ce0530018e46f9c5b8aefec559f0f6401f3f66e730374ea6086730a93083d1e54df98ceaa7f0643d6d450c346b57e58ba8f8289956c37c63820ac0e16243f542006665a6fe9f794d07b4f5e397939648ae26ecbdae89fc55ef1d8ef396bfd7ed36d54fc821c8f283fdbb5dcebec5f739dfdeab88d76ca5389f96456da4fb9f1b2d7a54cb9ed14aab85dadcd9ab1d8e9cba25a016dd65ad3860fdefc7e41a4668917e3fe9e7ef055cdf2e2316ae62018a901484a04026fbacd8e4621613e9e5c2491b85e106b169b73a466b9075443ca56f8f964c"
	unsignedRaw := "0xf9030b57851f3305c0d383015fe694090d4613473dee047c3f2706764f49e0821d256e80b902e42e7ba6ef000000000000000000000000000000000000000000000000000000000003d289000000000000000000000000fcaaf32bde89a81c71a8c2fb2c7bc8358c9e3696000000000000000000000000000000000000000000000015af1d78b58c4000000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000001297c61438bfe2221b7921320373b938d75afd5ffaad3c2c1c8600f821d1d2c2c4c5a63314f65054e549380d124e4ae43943aac8e1270c5ba75f51c65609c0c86e795e3009b789ea2c7cfafe90e0f6824e905c2c7579546e9d9804d1d6ce0f145d1bd962a0772320698558459716665860eb6ae2de865b0dd20b3fb86a24f48cd53e7ac5b3b0ee8f655d76985cd283fe2ccffb4c577752370503536afdc489a198ae6bb8589e780eb5066050f0d044450b79591fd399fc6ed18614b828da53e439798ad925730eed5b27309c7077a8b8ed81c132e98df8aaff09f5607848353a58c796657b78487c195c7c8d89a1c717682ae2c26ae4f91914e56db784a9d15ee9b981b8e9f05a09f0d1f2bd32f9dbd332e9a8d9615ff9fea154613e97523d59edc9513bb2603e9c6954d7f096c4b5dfc30e29824196f3643e16cf152badbfc813125e9ee5400a5ce7921109d8d9631a2e73515d25bf34fd01f266465389e4f212b8ea0a0e5d463659621029622a8b24638adb1d2c7fe784b720196075cb06ac21cccb879f5a8c8aebc9df8499c3503c073d660e49a72404baa353821b2823c122c2b0846e905cb2cdabdabf7037f329ce0530018e46f9c5b8aefec559f0f6401f3f66e730374ea6086730a93083d1e54df98ceaa7f0643d6d450c346b57e58ba8f8289956c37c63820ac0e16243f542006665a6fe9f794d07b4f5e397939648ae26ecbdae89fc55ef1d8ef396bfd7ed36d54fc821c8f283fdbb5dcebec5f739dfdeab88d76ca5389f96456da4fb9f1b2d7a54cb9ed14aab85dadcd9ab1d8e9cba018080"

	leafBytes := hexutil.MustDecode(raw)
	paddedLeafBytes := keccak.Pad101Bytes(leafBytes)
	var leafRlpHex [3264]frontend.Variable
	for i, b := range paddedLeafBytes {
		n1 := b >> 4
		n2 := b & 0x0F
		leafRlpHex[i*2] = n1
		leafRlpHex[i*2+1] = n2
	}

	for i := len(paddedLeafBytes) * 2; i < 3264; i++ {
		leafRlpHex[i] = 0
	}

	unsignedTxBytes := hexutil.MustDecode(unsignedRaw)
	paddedUnsignedTxBytes := keccak.Pad101Bytes(unsignedTxBytes)
	var unsignedRlpHex [3264]frontend.Variable
	for i, b := range paddedUnsignedTxBytes {
		n1 := b >> 4
		n2 := b & 0x0F
		unsignedRlpHex[i*2] = n1
		unsignedRlpHex[i*2+1] = n2
	}
	for i := len(paddedUnsignedTxBytes) * 2; i < 3264; i++ {
		unsignedRlpHex[i] = 0
	}

	leafFields, _, LeafFieldsLen := DecodeLegacyTxLeafRlp(api, leafRlpHex[:])

	nonce := HexToDecimal(api, leafFields.Nonce, 16, LeafFieldsLen[0])
	gasPrice := HexToDecimal(api, leafFields.GasPrice, 16, LeafFieldsLen[1])
	gasLimit := HexToDecimal(api, leafFields.GasLimit, 8, LeafFieldsLen[2])
	to := HexToDecimal(api, leafFields.To, 40, LeafFieldsLen[3])
	valueHexLenShort := LessThan(api, LeafFieldsLen[4], 32)
	value1HexLen := api.Select(valueHexLenShort, LeafFieldsLen[4], 32)
	value2HexLen := api.Select(valueHexLenShort, 0, api.Sub(LeafFieldsLen[4], 32))
	value1 := HexToDecimal(api, leafFields.Value[:32], 32, value1HexLen)
	value2 := HexToDecimal(api, leafFields.Value[32:], 32, value2HexLen)
	api.Println(nonce, gasPrice, gasLimit, to, value1, value2)

	// check unsigned tx rlp inclusion in leaf rlp (signed tx raw)

	unsignedLeafBytes := hexutil.MustDecode(unsignedRaw)
	unsignedPaddedLeafBytes := keccak.Pad101Bytes(unsignedLeafBytes)
	var unsignedLeafRlpHex [3264]frontend.Variable
	for i, b := range unsignedPaddedLeafBytes {
		n1 := b >> 4
		n2 := b & 0x0F
		unsignedLeafRlpHex[i*2] = n1
		unsignedLeafRlpHex[i*2+1] = n2
	}

	for i := len(unsignedPaddedLeafBytes) * 2; i < 3264; i++ {
		unsignedLeafRlpHex[i] = 0
	}

	unsignedLeafFields, unsignedFieldsHexLen := DecodeLegacyUnsignedTxLeafRlp(api, unsignedLeafRlpHex[:])
	chainId := HexToDecimal(api, unsignedLeafFields.ChainId, 16, unsignedFieldsHexLen[0])
	api.AssertIsEqual(chainId, 1)

	// keccak256 hash unsigned tx raw to get the recover msg hash
	var unsignedTxBits []frontend.Variable
	for i := 0; i < len(unsignedRlpHex)/2; i++ {
		unsignedTxBits = append(unsignedTxBits, api.ToBinary(unsignedRlpHex[2*i+1], 4)...)
		unsignedTxBits = append(unsignedTxBits, api.ToBinary(unsignedRlpHex[2*i], 4)...)
	}
	digestH := keccak.Keccak256Bits(api, 12, 5, unsignedTxBits)
	digest := utils.Flip(utils.Flip(digestH[:]))

	fp, _ := emulated.NewField[emulated.Secp256k1Fr](api)
	var rBits []frontend.Variable
	rBits = append(rBits, api.ToBinary(leafFields.R[0], 248)...)
	rBits = append(rBits, api.ToBinary(leafFields.R[1], 32*8-248)...)
	rBytes := conv.Bits2Bytes(api, rBits)
	rEmulated := emu.ToElement[emulated.Secp256k1Fr](api, utils.Flip(rBytes), 8)
	rEmulated = fp.Reduce(rEmulated)

	var sBits []frontend.Variable
	sBits = append(sBits, api.ToBinary(leafFields.S[0], 248)...)
	sBits = append(sBits, api.ToBinary(leafFields.S[1], 32*8-248)...)
	sBytes := conv.Bits2Bytes(api, sBits)
	sEmulated := emu.ToElement[emulated.Secp256k1Fr](api, utils.Flip(sBytes), 8)

	digestBytes := conv.Bits2Bytes(api, digest)
	msgEmulated := emu.ToElement[emulated.Secp256k1Fr](api, digestBytes, 8)

	curve, _ := sw_emulated.New[emulated.Secp256k1Fp, emulated.Secp256k1Fr](api, sw_emulated.GetSecp256k1Params())

	from := evmprecompiles.ECRecover(api, *msgEmulated, 27, *rEmulated, *sEmulated, 1)
	point := curve.Generator()
	point.X = from.X
	point.Y = from.Y
	g1MarshalBits := curve.MarshalG1WithoutFlag(*point)
	g1MarshalBits = utils.FlipSubSlice(g1MarshalBits[:], 8)
	paddedG1MarshalBits := keccak.PadBits101(api, g1MarshalBits, 1)
	hashBits := keccak.Keccak256Bits(api, 1, keccak.GetRoundIndex(len(g1MarshalBits)), paddedG1MarshalBits)
	fromBits := utils.FlipSubSlice(hashBits[:], 8)[12*8:]
	fromFr := api.FromBinary(utils.Flip(fromBits)...)
	fmt.Printf("decoded from:%x\n", fromFr)
	return nil
}

func Test_Legacy_Decode(t *testing.T) {
	assert := test.NewAssert(t)

	witness := &LegacyCircuit{
		Out: 1,
	}

	err := test.IsSolved(&LegacyCircuit{}, witness, ecc.BLS12_377.ScalarField())
	assert.NoError(err)
}

// rlpHash encodes x and hashes the encoded bytes.
func rlpHash(x interface{}) ([]byte, error) {
	data, err := rlp.EncodeToBytes(x)
	if err != nil {
		fmt.Errorf("encode failed:%v", err)
	}
	return data, err
}
