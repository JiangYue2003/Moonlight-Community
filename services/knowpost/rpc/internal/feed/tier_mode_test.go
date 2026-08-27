package feed

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseAuthorTierMode(t *testing.T) {
	for _, test := range []struct {
		raw  string
		want AuthorTierMode
	}{
		{raw: "", want: AuthorTierModeOff},
		{raw: "off", want: AuthorTierModeOff},
		{raw: " SHADOW ", want: AuthorTierModeShadow},
		{raw: "enforce", want: AuthorTierModeEnforce},
	} {
		got, err := ParseAuthorTierMode(test.raw)
		require.NoError(t, err)
		require.Equal(t, test.want, got)
	}
}

func TestParseAuthorTierModeRejectsUnknownValue(t *testing.T) {
	_, err := ParseAuthorTierMode("automatic")
	require.ErrorContains(t, err, "expected off, shadow, or enforce")
}
