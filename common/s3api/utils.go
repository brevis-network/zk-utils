package s3api

import "fmt"

func ObjectKey(prefix string, slot uint64) string {
	return fmt.Sprintf("%s/%d.json", prefix, slot)
}
