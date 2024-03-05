package mpt

import (
	"github.com/brevis-network/zk-utils/circuits/gadgets/rlp"
	"github.com/consensys/gnark/std/selector"

	"github.com/celer-network/goutils/log"
	"github.com/consensys/gnark/frontend"
)

// max rlp len 6 + (per rlp prefix len 2 + per rlp len 64) + 0x80
const MaxBranchNodeRlpHexLen = 6 + (64+2)*16 + 2

type MPTTwoItemsNodeCircuit struct {
	api frontend.API
	In  []frontend.Variable
}

type MPTCheckResult struct {
	output         frontend.Variable
	rlpTotalLength frontend.Variable
}

func (mr MPTCheckResult) GetOutput() frontend.Variable {
	return mr.output
}

type MPTLeafCheck struct {
	maxKeyLength            int
	maxValueLength          int
	maxRLPLength            int
	maxRLPArrayPrefixLength int
}

type MPTLeafCheckResult struct {
	result      MPTCheckResult
	valueLength frontend.Variable
}

func (m *MPTLeafCheckResult) GetResult() MPTCheckResult {
	return m.result
}

func NewMPTLeafCheck(maxKeyLength int, maxValueLength int) MPTLeafCheck {
	leafCheck := &MPTLeafCheck{}
	leafCheck.maxKeyLength = maxKeyLength
	leafCheck.maxValueLength = maxValueLength
	leafCheck.maxRLPLength = 4 + (maxKeyLength + 2) + 4 + maxValueLength
	leafBits := rlp.LogCeil(leafCheck.maxRLPLength)
	leafCheck.maxRLPArrayPrefixLength = 2 * int(leafBits/8+1)
	return *leafCheck
}

func (leafCheck *MPTLeafCheck) CheckLeaf(
	api frontend.API,
	keyNibbleLen frontend.Variable,
	keyNibbles []frontend.Variable,
	values []frontend.Variable,
	leafRlp []frontend.Variable,
	leafPathPrefixLength frontend.Variable,
) MPTLeafCheckResult {
	api.AssertIsEqual(len(keyNibbles), leafCheck.maxKeyLength)
	api.AssertIsEqual(len(values), leafCheck.maxValueLength)

	valuePrefix := values[:2]
	isSingleValue := rlp.ArrayEqual(api, valuePrefix[:], []frontend.Variable{8,0}, 2,2)
	isZeroArr := rlp.ArrayEqual(api, values[2:], []frontend.Variable{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}, 64, 64)
	isZeroValue := api.And(isSingleValue, isZeroArr)

	arrayCheck := &rlp.ArrayCheck{}
	arrayCheck.MaxHexLen = leafCheck.maxRLPLength
	arrayCheck.MaxFields = 2
	arrayCheck.ArrayPrefixMaxHexLen = leafCheck.maxRLPArrayPrefixLength
	arrayCheck.FieldMinHexLen = []int{0, 0}
	arrayCheck.FieldMaxHexLen = []int{leafCheck.maxKeyLength + 2, leafCheck.maxValueLength}

	rlpout, totalRlpLength, fieldsLength, fields := arrayCheck.RlpArrayCheck(api, leafRlp)

	oddLengthPrefix := rlp.Equal(api, leafPathPrefixLength, 1)
	oddPath := rlp.Equal(api, fields[0][0], 3)

	evenLengthPrefix := rlp.Equal(api, leafPathPrefixLength, 2)
	firstPrefixForEven := rlp.Equal(api, fields[0][0], 2)
	secondPrefixForEven := rlp.Equal(api, fields[0][1], 0)

	prefixCheck := api.Add(api.Mul(oddLengthPrefix, oddPath), api.Mul(evenLengthPrefix, firstPrefixForEven, secondPrefixForEven))

	keyNibblesFromRLP := rlp.ShiftLeft(api, leafCheck.maxRLPLength, 0, 2, fields[0], leafPathPrefixLength)

	keyNibblesMatched := rlp.ArrayEqual(api, keyNibbles, keyNibblesFromRLP, leafCheck.maxKeyLength, api.Sub(fieldsLength[0], leafPathPrefixLength))
	isNonExistence := api.And(isZeroValue, api.Sub(1, keyNibblesMatched))
	keyNibblesLengthMatched := rlp.Equal(api, keyNibbleLen, api.Sub(fieldsLength[0], leafPathPrefixLength))

	keyPass := api.Mul(keyNibblesMatched, keyNibblesLengthMatched)
	valueMatched := rlp.ArrayEqual(api, values, fields[1], leafCheck.maxValueLength, fieldsLength[1])
	rlpCheckOut := api.Add(rlpout, prefixCheck)
	checkOut := api.Add(rlpCheckOut, keyPass, valueMatched)
	checkOut = api.Select(isNonExistence, api.Add(rlpCheckOut, 2), checkOut)
	return MPTLeafCheckResult{
		result: MPTCheckResult{
			output:         checkOut,
			rlpTotalLength: totalRlpLength,
		},
		valueLength: fieldsLength[1],
	}
}

type MPTExtensionCheck struct {
	maxKeyLength            int
	maxNodeRefLength        int
	maxRLPLength            int
	maxRLPArrayPrefixLength int
}

func NewMPTExtensionCheck(maxKeyLength int, maxNodeRefLength int) MPTExtensionCheck {
	extensionCheck := &MPTExtensionCheck{}
	extensionCheck.maxKeyLength = maxKeyLength
	extensionCheck.maxNodeRefLength = maxNodeRefLength
	extensionCheck.maxRLPLength = 4 + (maxKeyLength + 2) + 2 + maxNodeRefLength
	leafBits := rlp.LogCeil(extensionCheck.maxRLPLength)
	extensionCheck.maxRLPArrayPrefixLength = 2 * int(leafBits/8+1)
	return *extensionCheck
}

func (extensionCheck *MPTExtensionCheck) CheckExtension(
	api frontend.API,
	keyNibbleLen frontend.Variable,
	keyNibbles []frontend.Variable,
	nodeRefLength frontend.Variable,
	nodeRefs []frontend.Variable,
	nodeRlp []frontend.Variable,
	nodePathPrefixLength frontend.Variable,
) MPTCheckResult {

	log.Debug("extension check params: ",
		extensionCheck.maxKeyLength,
		extensionCheck.maxNodeRefLength,
		extensionCheck.maxRLPArrayPrefixLength,
		extensionCheck.maxRLPLength,
		nodeRefLength,
		nodePathPrefixLength)

	arrayCheck := &rlp.ArrayCheck{}
	arrayCheck.MaxHexLen = extensionCheck.maxRLPLength
	arrayCheck.MaxFields = 2
	arrayCheck.ArrayPrefixMaxHexLen = extensionCheck.maxRLPArrayPrefixLength
	arrayCheck.FieldMinHexLen = []int{0, 0}
	arrayCheck.FieldMaxHexLen = []int{extensionCheck.maxKeyLength + 2, extensionCheck.maxNodeRefLength}

	rlpout, totalRlpLength, fieldsLength, fields := arrayCheck.RlpArrayCheck(api, nodeRlp)

	log.Debug(rlpout, totalRlpLength, fieldsLength, fields)

	log.Debug("nodePathPrefixLength:", nodePathPrefixLength)

	/// Extension
	/// If odd --> rlpFields[0][0] = 1
	/// If even ---> rlpFields[0][0] = 0 rlpFields[0][1] = 0
	oddLengthPrefix := rlp.Equal(api, nodePathPrefixLength, 1)
	oddPath := rlp.Equal(api, fields[0][0], 1)

	evenLengthPrefix := rlp.Equal(api, nodePathPrefixLength, 2)
	firstPrefixForEven := rlp.Equal(api, fields[0][0], 0)
	secondPrefixForEven := rlp.Equal(api, fields[0][1], 0)

	prefixCheck := api.Add(api.Mul(oddLengthPrefix, oddPath), api.Mul(evenLengthPrefix, firstPrefixForEven, secondPrefixForEven))

	log.Debug(prefixCheck)

	keyNibblesFromRLP := rlp.ShiftLeft(api, extensionCheck.maxRLPLength, 0, 2, fields[0], nodePathPrefixLength)

	log.Debug(keyNibblesFromRLP)

	keyNibblesMatched := rlp.ArrayEqual(api, keyNibbles, keyNibblesFromRLP, extensionCheck.maxKeyLength, api.Sub(fieldsLength[0], nodePathPrefixLength))
	keyNibblesLengthMatched := rlp.Equal(api, keyNibbleLen, api.Sub(fieldsLength[0], nodePathPrefixLength))
	keyPass := api.Mul(keyNibblesMatched, keyNibblesLengthMatched)

	nodeRefMatched := rlp.ArrayEqual(api, nodeRefs, fields[1], extensionCheck.maxNodeRefLength, fieldsLength[1])
	nodeRefLengthMatched := rlp.Equal(api, nodeRefLength, fieldsLength[1])

	log.Debug(nodeRefs, fields[1], extensionCheck.maxNodeRefLength, fieldsLength[1], nodeRefMatched)
	log.Debug(nodeRefLength, fieldsLength[1], nodeRefLengthMatched)

	nodeRefPass := api.Mul(nodeRefMatched, nodeRefLengthMatched)

	log.Debug("Extension check result: ", rlpout, prefixCheck, keyPass, nodeRefPass)
	return MPTCheckResult{
		output:         api.Add(rlpout, prefixCheck, keyPass, nodeRefPass),
		rlpTotalLength: totalRlpLength,
	}
}

type MPTBranchCheck struct {
	maxNodeRefLength        int
	maxRLPLength            int
	maxRLPArrayPrefixLength int
}

func NewMPTBranchCheck(maxNodeRefLength int) MPTBranchCheck {
	branchCheck := &MPTBranchCheck{}
	branchCheck.maxNodeRefLength = maxNodeRefLength
	branchCheck.maxRLPLength = MaxBranchNodeRlpHexLen
	branchCheck.maxRLPArrayPrefixLength = 8
	return *branchCheck
}

// CheckBranchForNonExistence prove the location of nibble key point to a data is 0x80, meaning that's an empty, length of this rlp is 0
func (branchCheck *MPTBranchCheck) CheckBranchForNonExistence(
	api frontend.API,
	keyNibble frontend.Variable,
	nodeRLP []frontend.Variable,
) frontend.Variable {
	arrayCheck := &rlp.ArrayCheck{}
	arrayCheck.MaxHexLen = branchCheck.maxRLPLength
	arrayCheck.MaxFields = 17
	arrayCheck.ArrayPrefixMaxHexLen = 8
	arrayCheck.FieldMinHexLen = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	arrayCheck.FieldMaxHexLen = []int{64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 0}

	valid, _, fieldsLength, _ := arrayCheck.RlpArrayCheck(api, nodeRLP)
	isEmpty := api.IsZero(selector.Mux(api, keyNibble, fieldsLength...))
	out := api.Add(valid, isEmpty)
	return api.IsZero(api.Sub(out, 2))
}

func (branchCheck *MPTBranchCheck) CheckBranch(
	api frontend.API,
	keyNibble frontend.Variable,
	nodeRefLength frontend.Variable,
	nodeRefs []frontend.Variable,
	nodeRLP []frontend.Variable,
) MPTCheckResult {
	arrayCheck := &rlp.ArrayCheck{}
	arrayCheck.MaxHexLen = branchCheck.maxRLPLength
	arrayCheck.MaxFields = 17
	arrayCheck.ArrayPrefixMaxHexLen = 8
	arrayCheck.FieldMinHexLen = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	arrayCheck.FieldMaxHexLen = []int{64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 0}

	rlpout, totalRlpLength, fieldsLength, fields := arrayCheck.RlpArrayCheck(api, nodeRLP)

	var fieldsMultiplexerInput [][]frontend.Variable
	var fieldsLengthMultiplexerInput [][]frontend.Variable

	fieldsLengthMultiplexerInput = append(fieldsLengthMultiplexerInput, []frontend.Variable{})

	for i := 0; i < 16; i++ {
		for j := 0; j < 64; j++ {
			if len(fieldsMultiplexerInput) == j {
				fieldsMultiplexerInput = append(fieldsMultiplexerInput, []frontend.Variable{})
			}
			fieldsMultiplexerInput[j] = append(fieldsMultiplexerInput[j], fields[i][j])
		}
		fieldsLengthMultiplexerInput[0] = append(fieldsLengthMultiplexerInput[0], fieldsLength[i])
	}

	fieldsMultiplexer := rlp.Multiplexer(api, keyNibble, 64, 16, fieldsMultiplexerInput)
	nodeRefCheck := rlp.ArrayEqual(api, nodeRefs, fieldsMultiplexer, branchCheck.maxNodeRefLength, nodeRefLength)

	rlpFieldLength := rlp.Multiplexer(api, keyNibble, 1, 16, fieldsLengthMultiplexerInput)[0]
	nodeRefLengthCheck := rlp.Equal(api, rlpFieldLength, nodeRefLength)

	return MPTCheckResult{
		output:         api.Add(rlpout, nodeRefCheck, nodeRefLengthCheck, 1),
		rlpTotalLength: totalRlpLength,
	}
}
