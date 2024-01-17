package circuits

const (
	ChunkDepth         = 7
	SMTDepth           = 18
	HeaderSizePerChunk = uint64(1 << ChunkDepth)
	ChunkSizePerSMT    = uint64(1 << SMTDepth)
)

func ChunkIndexFound(targetIndex uint64, indexes []uint64) bool {
	for _, value := range indexes {
		if value == targetIndex {
			return true
		}
	}
	return false
}
