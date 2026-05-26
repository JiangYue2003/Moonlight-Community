package security

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	numberLikeRe = regexp.MustCompile(`\b[A-Za-z]{1,6}[-_]?\d{2,}\b`)
	urlRe        = regexp.MustCompile(`https?://`)
)

// ValidateQueryInput 对用户查询做基础防护，避免超长输入和明显注入片段。
func ValidateQueryInput(q string, maxLen int) error {
	q = strings.TrimSpace(q)
	if q == "" {
		return fmt.Errorf("question required")
	}
	if maxLen <= 0 {
		maxLen = 2000
	}
	if len([]rune(q)) > maxLen {
		return fmt.Errorf("question too long")
	}
	lower := strings.ToLower(q)
	blocked := []string{"ignore previous instructions", "system prompt", "api_key", "token="}
	for _, b := range blocked {
		if strings.Contains(lower, b) {
			return fmt.Errorf("unsafe query")
		}
	}
	return nil
}

// ClampTopK 限制检索 topK，防止资源滥用。
func ClampTopK(v, def, min, max int) int {
	if def <= 0 {
		def = 10
	}
	if min <= 0 {
		min = 1
	}
	if max < min {
		max = min
	}
	if v <= 0 {
		v = def
	}
	if v < min {
		v = min
	}
	if v > max {
		v = max
	}
	return v
}

// GuessIntent 规则式意图判断，用于通道路由权重。
func GuessIntent(q string) string {
	text := strings.TrimSpace(strings.ToLower(q))
	if text == "" {
		return "semantic"
	}
	if numberLikeRe.MatchString(text) || strings.Contains(text, "编号") || strings.Contains(text, "代码") {
		return "exact"
	}
	if strings.Contains(text, "谁") && (strings.Contains(text, "负责") || strings.Contains(text, "关系")) {
		return "relation"
	}
	if urlRe.MatchString(text) {
		return "exact"
	}
	return "semantic"
}

// EnsureUserScope 强制 user_id 合法，防止工具越权。
func EnsureUserScope(userID int64) error {
	if userID <= 0 {
		return fmt.Errorf("unauthorized")
	}
	return nil
}
