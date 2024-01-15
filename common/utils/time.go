package utils

import (
	"time"

	"github.com/celer-network/goutils/log"
)

func TsMilli(t time.Time) uint64 {
	ts := uint64(t.UnixNano())
	return ts / uint64(time.Millisecond)
}

func TsMilliToTime(ts uint64) time.Time {
	return time.Unix(0, int64(ts*1000000))
}

func TsSecToTime(ts uint64) time.Time {
	return time.Unix(int64(ts), 0)
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Infof("%s took %s", name, elapsed)
}
