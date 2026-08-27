package feed

import (
	"fmt"
	"strings"
)

type Strategy string

const (
	StrategyPush   Strategy = "push"
	StrategyPull   Strategy = "pull"
	StrategyHybrid Strategy = "hybrid"
)

func ParseStrategy(raw string) (Strategy, error) {
	strategy := Strategy(strings.ToLower(strings.TrimSpace(raw)))
	if strategy == "" {
		return StrategyHybrid, nil
	}

	switch strategy {
	case StrategyPush, StrategyPull, StrategyHybrid:
		return strategy, nil
	default:
		return "", fmt.Errorf("invalid feed strategy %q: expected push, pull, or hybrid", raw)
	}
}
