package proof

import (
	"math/big"

	"github.com/brevis-network/zk-utils/common/eth"
	"github.com/brevis-network/zk-utils/common/utils"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	S3_UPDATE_PREFIX             = "ethereum/update"
	S3_UPDATE_COMMITTEE_PREFIX   = S3_UPDATE_PREFIX + "/committee"
	S3_UPDATE_HEADER_PREFIX      = S3_UPDATE_PREFIX + "/header"
	S3_UPDATE_CHUNK_PREFIX       = S3_UPDATE_PREFIX + "/chunk"
	S3_UPDATE_SMT_PREFIX         = S3_UPDATE_PREFIX + "/smt"
	S3_UPDATE_STORAGE_PREFIX     = S3_UPDATE_PREFIX + "/storage"
	S3_UPDATE_TRANSACTION_PREFIX = S3_UPDATE_PREFIX + "/transaction"
	S3_UPDATE_RECEIPT_PREFIX     = S3_UPDATE_PREFIX + "/receipt"
	S3_UPDATE_UNI_SUM_PREFIX     = S3_UPDATE_PREFIX + "/uniswapvolsum"
	S3_UPDATE_ZKSDK_PREFIX       = S3_UPDATE_PREFIX + "/query"

	S3_PROOF_PREFIX             = "ethereum/proof"
	S3_PROOF_HEADER_PREFIX      = S3_PROOF_PREFIX + "/sig"
	S3_PROOF_COMMITTEE_PREFIX   = S3_PROOF_PREFIX + "/committee"
	S3_PROOF_CHUNK_PREFIX       = S3_PROOF_PREFIX + "/chunk"
	S3_PROOF_SMT_PREFIX         = S3_PROOF_PREFIX + "/smt"
	S3_PROOF_STORAGE_PREFIX     = S3_PROOF_PREFIX + "/storage"
	S3_PROOF_TRANSACTION_PREFIX = S3_PROOF_PREFIX + "/transaction"
	S3_PROOF_RECEIPT_PREFIX     = S3_PROOF_PREFIX + "/receipt"
	S3_PROOF_UNI_VOL_SUM_PREFIX = S3_PROOF_PREFIX + "/uniswapvolsum"
	S3_PROOF_UNI_RECEIPT_PREFIX = S3_PROOF_PREFIX + "/uniswap_receipt"
	S3_PROOF_ZKSDK_PREFIX       = S3_PROOF_PREFIX + "/query"

	S3_ASSIGNMENT_PREFIX     = S3_UPDATE_PREFIX + "_assignment"
	S3_ASSIGNMENT_SMT_PREFIX = S3_ASSIGNMENT_PREFIX + "/smt"
)

type ChunkUpdate struct {
	StartBlockNum uint64 `json:"start_block_num"`
	Count         uint64 `json:"count"`
	ChainId       uint64 `json:"chain_id"`
	TargetChainId uint64 `json:"target_chain_id"`
}

type StorageProvingInfo struct {
	BlockNum  uint64 `json:"start_block_num"`
	Address   string `json:"address"`
	Key       string `json:"key"`
	SlotValue string `json:"slot_value"`
}

type TransactionProvingInfo struct {
	TransactionHash string `json:"transaction_hash"`
	ChainId         uint64 `json:"chain_id"`
}

type ReceiptProvingInfo struct {
	TransactionHash string `json:"transaction_hash"`
	ChainId         uint64 `json:"chain_id"`
}

type UniswapSumProvingInfo struct {
	TxHashes []string `json:"tx_hashes"`
	ChainId  uint64   `json:"chain_id"`
}

func (queries *SDKQueryProvingInfo) GetSubProofS3Path() []string {
	var ids []string
	for _, receiptInfo := range queries.ReceiptInfos {
		ids = append(ids, receiptInfo.TransactionHash)
	}
	for _, storageInfo := range queries.StorageSlotInfos {
		ids = append(ids, storageInfo.StorageHash)
	}

	for _, transactionInfo := range queries.TransactionInfos {
		ids = append(ids, transactionInfo.TransactionHash)
	}

	return ids
}

type ExportUniswapVolPublicInputs struct {
	Recipient []byte
	Volume    *big.Int
	SmtRoot0  []byte
	SmtRoot1  []byte
	Length    *big.Int
	VkHash    []byte
}

type ReceiptDecoderAggPubInputs struct {
	CommitHash []byte
	Length     *big.Int
	VkHash     []byte
}

type UpdateProof struct {
	SszProof         *CommitteeSszMappingProof `json:"committee_map_proof"`
	SigProof         *SigProof                 `json:"bls_sig_proof"`
	ChunkProof       *ChunkProof               `json:"chunk_proof"`
	StorageProof     *StorageProof             `json:"storage_proof"`
	TransactionProof *TransactionProof         `json:"transaction_proof"`
	ReceiptProof     *ReceiptProof             `json:"receipt_proof"`
	SMTUpdateProof   *SMTUpdateProof           `json:"smt_update_proof"`
	DurationSeconds  uint64                    `json:"duration_seconds"`
}

type CommitteeSszMappingProof struct {
	Sha256SszRoot   string       `json:"sha256_ssz_root"`
	PoseidonSszRoot string       `json:"poseidon_ssz_root"`
	Proof           Groth16Proof `json:"proof"`
}

func (p *CommitteeSszMappingProof) GetProof() eth.IBeaconVerifierProof {
	if p == nil {
		z := big.NewInt(0)
		// populate empty values otherwise contract binding would complain
		return eth.IBeaconVerifierProof{
			A:          [2]*big.Int{z, z},
			B:          [2][2]*big.Int{{z, z}, {z, z}},
			C:          [2]*big.Int{z, z},
			Commitment: [2]*big.Int{z, z},
		}
	}
	return p.Proof.BeaconProof()
}

type SigProof struct {
	Participation uint64       `json:"participation"`
	PoseidonRoot  string       `json:"poseidon_root"`
	CommitmentPub string       `json:"commitment_pub"`
	Proof         Groth16Proof `json:"proof"`
}

func (p *SigProof) GetCommitment() *big.Int {
	return utils.Hex2BigInt(p.CommitmentPub)
}

func (p *SigProof) GetPoseidonRoot() common.Hash {
	return utils.Hex2Hash(p.PoseidonRoot)
}

type ChunkProof struct {
	ChunkRoot     string       `json:"chunk_root,omitempty"`
	PrevHash      string       `json:"prev_hash,omitempty"`
	EndHash       string       `json:"end_hash,omitempty"`
	StartBlockNum uint64       `json:"start_block_num,omitempty"`
	EndBlockNum   uint64       `json:"end_block_num,omitempty"`
	Proof         Groth16Proof `json:"proof,omitempty"`
}

func (p *ChunkProof) EncodeABI() []byte {
	var args [][]byte

	args = append(args, utils.Hex2Bytes(p.Proof.A[0]), utils.Hex2Bytes(p.Proof.A[1]))
	args = append(args, utils.Hex2Bytes(p.Proof.B[0][0]), utils.Hex2Bytes(p.Proof.B[0][1]))
	args = append(args, utils.Hex2Bytes(p.Proof.B[1][0]), utils.Hex2Bytes(p.Proof.B[1][1]))
	args = append(args, utils.Hex2Bytes(p.Proof.C[0]), utils.Hex2Bytes(p.Proof.C[1]))
	args = append(args, []byte{}, []byte{}) // field commitment should always be empty

	chunkRoot := utils.Hex2Hash(p.ChunkRoot)
	args = append(args, chunkRoot[0:16], chunkRoot[16:])
	prevHash := utils.Hex2Hash(p.PrevHash)
	args = append(args, prevHash[0:16], prevHash[16:])
	endHash := utils.Hex2Hash(p.EndHash)
	args = append(args, endHash[0:16], endHash[16:])
	startBlockNum := big.NewInt(int64(p.StartBlockNum)).Bytes()
	args = append(args, startBlockNum)
	endBlockNum := big.NewInt(int64(p.EndBlockNum)).Bytes()
	args = append(args, endBlockNum)

	var ret []byte
	for _, arg := range args {
		ret = append(ret, utils.Pad32Bytes(arg)...)
	}
	return ret
}

type Groth16Proof struct {
	A          [2]string    `json:"a"`
	B          [2][2]string `json:"b"`
	C          [2]string    `json:"c"`
	Commitment [2]string    `json:"commitment"`
}

func (p Groth16Proof) AsBigInts() (a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, commitment [2]*big.Int) {
	a = [2]*big.Int{utils.Hex2BigInt(p.A[0]), utils.Hex2BigInt(p.A[1])}
	b = [2][2]*big.Int{{utils.Hex2BigInt(p.B[0][0]), utils.Hex2BigInt(p.B[0][1])}, {utils.Hex2BigInt(p.B[1][0]), utils.Hex2BigInt(p.B[1][1])}}
	c = [2]*big.Int{utils.Hex2BigInt(p.C[0]), utils.Hex2BigInt(p.C[1])}
	commitment = [2]*big.Int{utils.Hex2BigInt(p.Commitment[0]), utils.Hex2BigInt(p.Commitment[1])}
	return
}

func (p Groth16Proof) OnChainProof() eth.IVerifierProof {
	a, b, c, commitment := p.AsBigInts()
	proof := eth.IVerifierProof{
		A:          a,
		B:          b,
		C:          c,
		Commitment: commitment,
	}
	return proof
}

func (p Groth16Proof) BeaconProof() eth.IBeaconVerifierProof {
	a, b, c, commitment := p.AsBigInts()
	proof := eth.IBeaconVerifierProof{
		A:          a,
		B:          b,
		C:          c,
		Commitment: commitment,
	}
	return proof
}

type StorageProof struct {
	BlockHash       string       `json:"block_hash,omitempty"`
	AddressProofKey string       `json:"address_proof_key,omitempty"`
	Slot            string       `json:"slot,omitempty"`
	SlotValue       string       `json:"slot_value,omitempty"`
	BlockNumber     uint64       `json:"block_num,omitempty"`
	Proof           Groth16Proof `json:"proof,omitempty"`
}

type TransactionProof struct {
	TransactionHash string       `json:"address_proof_key,omitempty"`
	BlockHash       string       `json:"block_hash,omitempty"`
	BlockNumber     uint64       `json:"block_num,omitempty"`
	BlockTime       uint64       `json:"block_time,omitempty"`
	LeafHash        string       `json:"leaf_hash,omitempty"`
	LeafRlpPrefix   string       `json:"leaf_rlp_prefix,omitempty"`
	Proof           Groth16Proof `json:"proof,omitempty"`
}

type ReceiptProof struct {
	TransactionHash string       `json:"transaction_hash,omitempty"`
	BlockHash       string       `json:"block_hash,omitempty"`
	BlockNumber     uint64       `json:"block_num,omitempty"`
	BlockTime       uint64       `json:"block_time,omitempty"`
	LeafHash        string       `json:"leaf_hash,omitempty"`
	LeafRlpPrefix   string       `json:"leaf_rlp_prefix,omitempty"`
	Proof           Groth16Proof `json:"proof,omitempty"`
}

type SMTUpdateProof struct {
	OldSmtRoot          string       `json:"transaction_hash,omitempty"`
	NewSmtRoot          string       `json:"block_hash,omitempty"`
	EndBlockHash        string       `json:"block_num,omitempty"`
	EndBlockNum         uint64       `json:"block_time,omitempty"`
	NextChunkMerkleRoot string       `json:"leaf_hash,omitempty"`
	Proof               Groth16Proof `json:"proof,omitempty"`
	CommitmentPub       string       `json:"commitment_pub,omitempty"`
}

func (p *SMTUpdateProof) GetOldSmtRoot() [32]byte   { return utils.Hex2Hash(p.OldSmtRoot) }
func (p *SMTUpdateProof) GetNewSmtRoot() [32]byte   { return utils.Hex2Hash(p.NewSmtRoot) }
func (p *SMTUpdateProof) GetEndBlockHash() [32]byte { return utils.Hex2Hash(p.EndBlockHash) }
func (p *SMTUpdateProof) GetNextChunkMerkleRoot() [32]byte {
	return utils.Hex2Hash(p.NextChunkMerkleRoot)
}
func (p *SMTUpdateProof) GetCommitmentPub() [32]byte { return utils.Hex2Hash(p.CommitmentPub) }

type SMTUpdateAssignmentData struct {
	OldSmtRoot          []byte            `json:"old_smt_root,omitempty"`
	NewSmtRoot          []byte            `json:"new_smt_root,omitempty"`
	EndBlockHash        []byte            `json:"end_block_hash,omitempty"`
	EndBlockNum         uint64            `json:"end_block_number,omitempty"`
	NextChunkMerkleRoot []byte            `json:"next_chunk_merkle_root,omitempty"`
	InsertionMerklePath [][]byte          `json:"insertion_merkle_path,omitempty"`
	NextChunkMerklePath [][]byte          `json:"next_chunk_merkle_path,omitempty"`
	Headers             []ethtypes.Header `json:"headers,omitempty"`
	UpdatedChunkIndexes []uint64          `json:"updated_chunk_indexes,omitempty"`
	NewChunkIndex       uint64            `json:"new_chunk_index,omitempty"`
	NewChunkMerkleRoot  []byte            `json:"new_chunk_merkle_root,omitempty"`
	NewChunkRoot        []byte            `json:"new_chunk_root,omitempty"`
}

type SDKQueryProvingInfo struct {
	Hash             string                               `json:"hash"`
	ChainId          uint64                               `json:"chain_id"`
	ReceiptInfos     []*SDKQueryProvingInfoForReceipt     `json:"receipt_infos"`
	StorageSlotInfos []*SDKQueryProvingInfoForStorageSlot `json:"storage_slot_infos"`
	TransactionInfos []*SDKQueryProvingInfoForTransaction `json:"transaction_infos"`
	SMTRoot          string                               `json:"smt_root"`
	BatchSize        uint64                               `json:"batch_size"`
	AppCircuitInfo   *SDKQueryProvingInfoForAppCircuit    `json:"app_circuit_info"`
}

type SDKQueryProvingInfoForReceipt struct {
	TransactionHash         string                    `json:"transaction_hash"`
	LogExtractInfos         []*SDKQueryLogExtractInfo `json:"log_extract_infos"`
	BlkNum                  uint64                    `json:"blk_num"`
	ReceiptIndex            uint64                    `json:"receipt_index"`
	ReceiptRlp              string                    `json:"receipt_rlp"`
	BlockHash               string                    `json:"block_hash"`
	BlockTime               uint64                    `json:"block_time"`
	MPTKey                  string                    `json:"mpt_key"`
	MPTProofs               []string                  `json:"mpt_proofs"`
	BlockRlp                string                    `json:"block_rlp"`
	BlockFieldsNum          int                       `json:"block_fields_num"`
	SMTRoot                 string                    `json:"smt_root"`
	ChunkProofs             []string                  `json:"chunk_proofs"`
	SmtStartBlockNumber     uint64                    `json:"smt_start_block_number"`
	SmtStartBlockParentHash string                    `json:"smt_start_block_parent_hash"`
	SMTInitialBlockNum      uint64                    `json:"smt_initial_block_number"`
	SmtProofs               []string                  `json:"smt_proofs"`
	CommitHash              string                    `json:"commit_hash"`
	TransactionType         uint8                     `json:"transaction_type"`
	QueryRaw                []string                  `json:"query_raw"`
}

type SDKQueryLogExtractInfo struct {
	ContractAddress string `json:"contract_address,omitempty"`
	LogIndex        uint64 `json:"log_index,omitempty"`
	LogTopic0       string `json:"log_topic0,omitempty"`
	ValueFromTopic  bool   `json:"value_from_topic,omitempty"`
	ValueIndex      uint64 `json:"value_index,omitempty"`
	Value           string `json:"value,omitempty"`
}

type SDKQueryProvingInfoForStorageSlot struct {
	AccountAddress          string   `json:"account_address"`
	Slot                    string   `json:"slot"`
	SlotValue               string   `json:"slot_value"`
	AccountProofs           []string `json:"account_proofs"`
	AccountRlp              string   `json:"account_rlp"`
	StorageProofs           []string `json:"storage_proofs"`
	StorageHash             string   `json:"storage_hash"`
	BlockHash               string   `json:"block_hash"`
	BlockNumber             uint64   `json:"block_number"`
	BlockRlp                string   `json:"block_rlp"`
	BlockFieldsNum          int      `json:"block_fields_num"`
	SMTRoot                 string   `json:"smt_root"`
	ChunkProofs             []string `json:"chunk_proofs"`
	SmtStartBlockNumber     uint64   `json:"smt_start_block_number"`
	SmtStartBlockParentHash string   `json:"smt_start_block_parent_hash"`
	SMTInitialBlockNum      uint64   `json:"smt_initial_block_number"`
	SmtProofs               []string `json:"smt_proofs"`
	CommitHash              string   `json:"commit_hash"`
	QueryRaw                []string `json:"query_raw"`
}

type SDKQueryProvingInfoForTransaction struct {
	TransactionHash         string               `json:"transaction_hash"`
	SMTRoot                 string               `json:"smt_root"`
	LeafHash                string               `json:"leaf_hash"`
	BlockHash               string               `json:"block_hash"`
	BlockNumber             uint64               `json:"blk_num"`
	BlockTime               uint64               `json:"block_time"`
	MPTKey                  string               `json:"mpt_key"`
	MPTProofs               []string             `json:"mpt_proofs"`
	MPTLeafRlpPrefix        string               `json:"mpt_leaf_rlp_prefix"`
	BlockRlp                string               `json:"block_rlp"`
	BlockFieldsNum          int                  `json:"block_fields_num"`
	ChunkProofs             []string             `json:"chunk_proofs"`
	SMTStartBlockNumber     uint64               `json:"smt_start_block_number"`
	SMTStartBlockParentHash string               `json:"smt_start_block_parent_hash"`
	SMTInitialBlockNum      uint64               `json:"smt_initial_block_number"`
	SMTProofs               []string             `json:"smt_proofs"`
	UnsignedTxRlp           string               `json:"unsigned_tx_rlp"`
	ExtraInfo               TransactionExtraInfo `json:"extra_info"`
	CommitHash              string               `json:"commit_hash"`
	TransactionType         uint8                `json:"transaction_type"`
	QueryRaw                []string             `json:"query_raw"`
}

type TransactionExtraInfo struct {
	ChainId              uint64 `json:"chain_id"`
	Nonce                uint64 `json:"nonce"`
	MaxPriorityFeePerGas string `json:"max_priority_fee_per_gas"`
	MaxFeePerGas         string `json:"max_fee_per_gas"`
	GasLimit             uint64 `json:"gas_limit"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	Value                string `json:"value"`
}

type SDKQueryProvingInfoForAppCircuit struct {
	Proof             string   `json:"proof"`
	VerifyingKey      string   `json:"verifying_key"`
	InputCommitments  []string `json:"input_commitments"`
	OutputCommitment  string   `json:"output_commitment"`
	TogglesCommitment string   `json:"toggles_commitment"`
	Toggles           []int    `json:"toggles"`
}
