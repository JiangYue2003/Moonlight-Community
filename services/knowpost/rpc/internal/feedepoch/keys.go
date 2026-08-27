package feedepoch

import (
	"fmt"
	"strings"
)

const defaultKeyPrefix = "feed"

func normalizeKeyPrefix(prefix string) string {
	prefix = strings.Trim(strings.TrimSpace(prefix), ":")
	if prefix == "" {
		return defaultKeyPrefix
	}
	return prefix
}

func relationKey(prefix string, userID int64) string {
	return fmt.Sprintf("%s:relation:epoch:%d", normalizeKeyPrefix(prefix), userID)
}

func safetyKey(prefix string) string {
	return normalizeKeyPrefix(prefix) + ":content:safety:epoch"
}
