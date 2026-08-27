package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildTopologyCreatesRealBigVAndMixedReaders(t *testing.T) {
	cfg := topologyConfig{
		FollowerUsers:       1100,
		NormalAuthors:       2,
		BigVAuthors:         1,
		Readers:             10,
		NormalFollowerCount: 100,
		BigVFollowerCount:   1001,
		Seed:                42,
	}
	userIDs := make([]int64, cfg.RequiredUsers())
	for i := range userIDs {
		userIDs[i] = int64(i + 1)
	}

	plan, err := buildTopology(userIDs, cfg)

	require.NoError(t, err)
	require.Len(t, plan.NormalAuthors, 2)
	require.Len(t, plan.BigVAuthors, 1)
	require.Len(t, plan.Readers, 10)
	require.Equal(t, 100, plan.FollowerCounts[plan.NormalAuthors[0]])
	require.Equal(t, 1001, plan.FollowerCounts[plan.BigVAuthors[0]])
	for _, reader := range plan.Readers {
		require.Contains(t, plan.FollowingByUser[reader], plan.NormalAuthors[0])
		require.Contains(t, plan.FollowingByUser[reader], plan.BigVAuthors[0])
	}
}

func TestBuildTopologyRejectsImpossibleFollowerCounts(t *testing.T) {
	cfg := topologyConfig{
		FollowerUsers:       10,
		NormalAuthors:       1,
		BigVAuthors:         1,
		Readers:             2,
		NormalFollowerCount: 5,
		BigVFollowerCount:   11,
	}

	_, err := buildTopology(make([]int64, cfg.RequiredUsers()), cfg)

	require.ErrorContains(t, err, "big-v follower count")
}

func TestReaderCardinalityPresetsBuildCoveredReaderPopulations(t *testing.T) {
	for _, testCase := range []struct {
		name        string
		wantReaders int
	}{
		{name: "hot", wantReaders: 20},
		{name: "distributed", wantReaders: 1200},
		{name: "high", wantReaders: 20000},
		{name: "cursor-deep", wantReaders: 20},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			cfg := topologyConfig{Seed: 42}
			require.NoError(t, applyReaderCardinalityPreset(testCase.name, &cfg))
			require.Equal(t, testCase.wantReaders, cfg.Readers)

			userIDs := make([]int64, cfg.RequiredUsers())
			for i := range userIDs {
				userIDs[i] = int64(i + 1)
			}
			plan, err := buildTopology(userIDs, cfg)
			require.NoError(t, err)
			require.Len(t, plan.Readers, testCase.wantReaders)
			for _, readerID := range plan.Readers {
				require.NotEmpty(t, plan.FollowingByUser[readerID], "reader %d must follow at least one benchmark author", readerID)
			}
			for _, authorID := range plan.NormalAuthors {
				require.Equal(t, cfg.NormalFollowerCount, plan.FollowerCounts[authorID])
			}
			for _, authorID := range plan.BigVAuthors {
				require.Equal(t, cfg.BigVFollowerCount, plan.FollowerCounts[authorID])
			}
		})
	}
}

func TestReaderCardinalityPresetRejectsUnknownName(t *testing.T) {
	err := applyReaderCardinalityPreset("planet-scale", &topologyConfig{})
	require.ErrorContains(t, err, "reader cardinality")
}
