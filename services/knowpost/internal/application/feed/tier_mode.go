package feed

import (
	"fmt"
	"strings"
)

type AuthorTierMode string

const (
	AuthorTierModeOff     AuthorTierMode = "off"
	AuthorTierModeShadow  AuthorTierMode = "shadow"
	AuthorTierModeEnforce AuthorTierMode = "enforce"
)

func ParseAuthorTierMode(raw string) (AuthorTierMode, error) {
	mode := AuthorTierMode(strings.ToLower(strings.TrimSpace(raw)))
	if mode == "" {
		return AuthorTierModeOff, nil
	}
	switch mode {
	case AuthorTierModeOff, AuthorTierModeShadow, AuthorTierModeEnforce:
		return mode, nil
	default:
		return "", fmt.Errorf("invalid feed author tier mode %q: expected off, shadow, or enforce", raw)
	}
}
