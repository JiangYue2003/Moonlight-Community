package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	readerCardinalityHot         = "hot"
	readerCardinalityDistributed = "distributed"
	readerCardinalityHigh        = "high"
	readerCardinalityCursorDeep  = "cursor-deep"
)

type topologyConfig struct {
	FollowerUsers       int
	NormalAuthors       int
	BigVAuthors         int
	Readers             int
	NormalFollowerCount int
	BigVFollowerCount   int
	Seed                int64
}

func (c topologyConfig) RequiredUsers() int {
	return c.NormalAuthors + c.BigVAuthors + c.FollowerUsers
}

func applyReaderCardinalityPreset(name string, cfg *topologyConfig) error {
	if cfg == nil {
		return fmt.Errorf("reader cardinality target is required")
	}
	seed := cfg.Seed
	switch strings.ToLower(strings.TrimSpace(name)) {
	case readerCardinalityHot:
		*cfg = topologyConfig{
			FollowerUsers: 1200, NormalAuthors: 20, BigVAuthors: 5, Readers: 20,
			NormalFollowerCount: 100, BigVFollowerCount: 1100, Seed: seed,
		}
	case readerCardinalityDistributed:
		*cfg = topologyConfig{
			FollowerUsers: 1200, NormalAuthors: 20, BigVAuthors: 5, Readers: 1200,
			NormalFollowerCount: 100, BigVFollowerCount: 1100, Seed: seed,
		}
	case readerCardinalityHigh:
		*cfg = topologyConfig{
			FollowerUsers: 20000, NormalAuthors: 20, BigVAuthors: 5, Readers: 20000,
			NormalFollowerCount: 100, BigVFollowerCount: 19000, Seed: seed,
		}
	case readerCardinalityCursorDeep:
		*cfg = topologyConfig{
			FollowerUsers: 1200, NormalAuthors: 20, BigVAuthors: 5, Readers: 20,
			NormalFollowerCount: 100, BigVFollowerCount: 1100, Seed: seed,
		}
	default:
		return fmt.Errorf("reader cardinality %q must be hot, distributed, high, or cursor-deep", name)
	}
	return nil
}

type topologyPlan struct {
	NormalAuthors   []int64
	BigVAuthors     []int64
	Readers         []int64
	FollowingByUser map[int64][]int64
	FollowerCounts  map[int64]int
}

func buildTopology(userIDs []int64, cfg topologyConfig) (*topologyPlan, error) {
	if cfg.NormalAuthors < 1 || cfg.BigVAuthors < 1 {
		return nil, fmt.Errorf("normal and big-v author counts must both be positive")
	}
	if cfg.FollowerUsers < 1 || len(userIDs) < cfg.RequiredUsers() {
		return nil, fmt.Errorf("users: got %d, need %d", len(userIDs), cfg.RequiredUsers())
	}
	if cfg.BigVFollowerCount > cfg.FollowerUsers {
		return nil, fmt.Errorf("big-v follower count %d exceeds follower users %d", cfg.BigVFollowerCount, cfg.FollowerUsers)
	}
	if cfg.NormalFollowerCount > cfg.FollowerUsers {
		return nil, fmt.Errorf("normal follower count %d exceeds follower users %d", cfg.NormalFollowerCount, cfg.FollowerUsers)
	}
	if cfg.Readers < 1 || cfg.Readers > cfg.FollowerUsers {
		return nil, fmt.Errorf("readers must fit within follower users")
	}

	normalEnd := cfg.NormalAuthors
	bigVEnd := normalEnd + cfg.BigVAuthors
	audienceEnd := bigVEnd + cfg.FollowerUsers
	plan := &topologyPlan{
		NormalAuthors:   append([]int64(nil), userIDs[:normalEnd]...),
		BigVAuthors:     append([]int64(nil), userIDs[normalEnd:bigVEnd]...),
		Readers:         append([]int64(nil), userIDs[bigVEnd:bigVEnd+cfg.Readers]...),
		FollowingByUser: make(map[int64][]int64),
		FollowerCounts:  make(map[int64]int),
	}
	audience := append([]int64(nil), userIDs[bigVEnd:audienceEnd]...)
	rng := rand.New(rand.NewSource(cfg.Seed))

	authorOrdinal := 0
	addFollowers := func(author int64, count int) {
		readerCount := minInt(count, len(plan.Readers))
		followers := make([]int64, 0, count)
		if readerCount == len(plan.Readers) {
			followers = append(followers, plan.Readers...)
		} else {
			start := authorOrdinal * count % len(plan.Readers)
			for offset := 0; offset < readerCount; offset++ {
				followers = append(followers, plan.Readers[(start+offset)%len(plan.Readers)])
			}
		}
		if len(followers) < count {
			remaining := append([]int64(nil), audience[cfg.Readers:]...)
			rng.Shuffle(len(remaining), func(i, j int) { remaining[i], remaining[j] = remaining[j], remaining[i] })
			followers = append(followers, remaining[:count-len(followers)]...)
		}
		for _, follower := range followers {
			plan.FollowingByUser[follower] = append(plan.FollowingByUser[follower], author)
		}
		plan.FollowerCounts[author] = len(followers)
		authorOrdinal++
	}

	for _, author := range plan.NormalAuthors {
		addFollowers(author, cfg.NormalFollowerCount)
	}
	for _, author := range plan.BigVAuthors {
		addFollowers(author, cfg.BigVFollowerCount)
	}
	return plan, nil
}
