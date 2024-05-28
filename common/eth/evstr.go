package eth

import (
	"fmt"
)

func (s *BrevisRequestRequestSent) String() string {
	return fmt.Sprintf("< %s>", fmt.Sprintf("<requestId %x >", s.RequestId))
}

func (ev *BrevisRequestRequestSent) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("RequestSent-%d: %s", onchid, ev)
}

func (s *BrevisRequestOpRequestsFulfilled) String() string {
	var out string
	for _, id := range s.RequestIds {
		out += fmt.Sprintf("<requestId %x > ", id)
	}
	return fmt.Sprintf("< %s>", out)
}

func (ev *BrevisRequestOpRequestsFulfilled) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("OpRequestsFulfilled-%d: %s", onchid, ev)
}
