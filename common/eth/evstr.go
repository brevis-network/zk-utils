package eth

import (
	"fmt"
)

func (s *BrevisRequestRequestSent) String() string {
	return fmt.Sprintf("< %s>", fmt.Sprintf("<requestId %x nonce %d>", s.ProofId, s.Nonce))
}

func (ev *BrevisRequestRequestSent) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("RequestSent-%d: %s", onchid, ev)
}

func (s *BrevisRequestOpRequestsFulfilled) String() string {
	var out string
	for i, id := range s.ProofIds {
		out += fmt.Sprintf("<requestId %x nonce %d> ", id, s.Nonces[i])
	}
	return fmt.Sprintf("< %s>", out)
}

func (ev *BrevisRequestOpRequestsFulfilled) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("OpRequestsFulfilled-%d: %s", onchid, ev)
}
