package userfeed

import (
	"fmt"
	"strings"
)

const StrategyHybridV1 = "hybrid-v1"

// Request contains every dimension that can change the meaning of a cached
// personal-feed page. Epoch changes create a new key instead of mutating or
// scanning old page keys.
type Request struct {
	UserID          int64
	StrategyVersion string
	RelationEpoch   uint64
	SafetyEpoch     uint64
	Page            int32
	Size            int32
}

func Key(prefix string, req Request) string {
	prefix = strings.Trim(strings.TrimSpace(prefix), ":")
	if prefix == "" {
		prefix = "feed"
	}
	return fmt.Sprintf("%s:page:v1:%s:%d:r%d:s%d:p%d:n%d",
		prefix,
		req.StrategyVersion,
		req.UserID,
		req.RelationEpoch,
		req.SafetyEpoch,
		req.Page,
		req.Size,
	)
}

func cacheable(req Request) bool {
	return req.UserID > 0 &&
		req.StrategyVersion == StrategyHybridV1 &&
		req.Page == 1 &&
		req.Size == 20
}
