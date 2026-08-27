package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFeedKeysForManifestAreExactAndDeduplicated(t *testing.T) {
	manifest := datasetManifest{
		Users:         []benchmarkUser{{ID: 2}, {ID: 1}, {ID: 2}},
		NormalAuthors: []int64{1},
		BigVAuthors:   []int64{3, 2},
		Posts:         []int64{11, 10, 11},
	}

	keys := feedKeysForManifest(manifest)

	require.Equal(t, []string{
		"feed:bigv:1",
		"feed:bigv:2",
		"feed:bigv:3",
		"feed:fanout:processing:10",
		"feed:fanout:processing:11",
		"feed:inbox:1",
		"feed:inbox:2",
		"feed:inbox:3",
	}, keys)
}
