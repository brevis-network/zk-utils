package eth

import (
	"fmt"
	"strings"
)

func (u LightClientUpdate) String() string {
	return fmt.Sprintf(`
attestedHeader {%s}, 
finalizedHeader {%s}, 
finalityBranch [%s], 
nextSyncCommitteeRoot %x, 
nextSyncCommitteeBranch [%s], 
nextSyncCommitteePoseidonRoot %x, 
nextSyncCommitteeRootMappingProof %v, 
syncAggregate %s, 
signatureSlot %d`,
		u.AttestedHeader,
		u.FinalizedHeader,
		b2s(u.FinalityBranch),
		u.NextSyncCommitteeRoot,
		b2s(u.NextSyncCommitteeBranch),
		u.NextSyncCommitteePoseidonRoot,
		u.NextSyncCommitteeRootMappingProof,
		u.SyncAggregate,
		u.SignatureSlot)
}

func b2s(in [][32]byte) string {
	ret := []string{}
	for _, item := range in {
		ret = append(ret, fmt.Sprintf("%x", item[:]))
	}
	return strings.Join(ret, ", ")
}

func (h BeaconBlockHeader) String() string {
	return fmt.Sprintf("slot %d, proposerIndex %d, parentRoot %x, stateRoot %x, bodyRoot %x",
		h.Slot, h.ProposerIndex, h.ParentRoot, h.StateRoot, h.BodyRoot)
}

func (h HeaderWithExecution) String() string {
	return fmt.Sprintf("\n\tbeacon %s\n\texecution %s\n\texecutionRoot %s\n", h.Beacon, h.Execution, h.ExecutionRoot)
}

func (e ExecutionPayload) String() string {
	return fmt.Sprintf("\n\tstateRoot {%s}\n\tblockHash {%s}\n\tblockNumber {%s}\n", e.StateRoot, e.BlockHash, e.BlockNumber)
}

func (l LeafWithBranch) String() string {
	return fmt.Sprintf("leaf %x, branch %s", l.Leaf, b2s(l.Branch))
}

func (s SyncAggregate) String() string {
	return fmt.Sprintf("participation %d, poseidonRoot %x, commitment %d, proof %v", s.Participation, s.PoseidonRoot, s.Commitment, s.Proof)
}

func (u ISMTSmtUpdate) String() string {
	return fmt.Sprintf("newSmtRoot %x, endBlockHash %x, endBlockNum %d, nextChunkMerkleRoot %x, proof %s, commit: %s, Pok: %s", u.NewSmtRoot, u.EndBlockHash, u.EndBlockNum, u.NextChunkMerkleRoot, u.Proof, u.Commit, u.KnowledgeProof)
}

func (p IBeaconVerifierProof) String() string {
	return fmt.Sprintf("a %s, b %s, c %s, commitment %s", p.A, p.B, p.C, p.Commitment)
}
