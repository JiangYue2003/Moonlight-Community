package textx

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// NormalizeNFKC 兼容/字形归一（半角全角 → 半角；连字 → 标准字符等）。
func NormalizeNFKC(s string) string { return norm.NFKC.String(s) }

// CollapseWhitespace 把任何连续 unicode 空白（含全角空格、零宽、制表/换行）压成单个 ASCII 空格，并 trim。
func CollapseWhitespace(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	prev := false
	for _, r := range s {
		if unicode.IsSpace(r) || r == ' ' || r == '​' || r == '　' {
			if !prev {
				b.WriteByte(' ')
				prev = true
			}
			continue
		}
		prev = false
		b.WriteRune(r)
	}
	return strings.TrimSpace(b.String())
}

// StripWrappingQuotes 去除首尾成对的引号（中英）。
func StripWrappingQuotes(s string) string {
	pairs := [][2]rune{
		{'"', '"'}, {'\'', '\''},
		{'“', '”'}, // “ ”
		{'‘', '’'}, // ‘ ’
		{'「', '」'}, // 「 」
		{'『', '』'}, // 『 』
		{'《', '》'}, // 《 》
	}
	for {
		r := []rune(s)
		if len(r) < 2 {
			return s
		}
		matched := false
		for _, p := range pairs {
			if r[0] == p[0] && r[len(r)-1] == p[1] {
				s = string(r[1 : len(r)-1])
				s = strings.TrimSpace(s)
				matched = true
				break
			}
		}
		if !matched {
			return s
		}
	}
}

// 末尾标点（中英）。
var trailingPunct = "。.!！?？,，;；:：、~～—–-…"

// StripTrailingPunct 去除末尾的中英标点。
func StripTrailingPunct(s string) string {
	for {
		r := []rune(s)
		if len(r) == 0 {
			return s
		}
		last := r[len(r)-1]
		if !strings.ContainsRune(trailingPunct, last) && !unicode.IsSpace(last) {
			return s
		}
		s = string(r[:len(r)-1])
	}
}

// TruncateRunes 按 rune 截断到至多 n 个码点；n<=0 返回原串。
func TruncateRunes(s string, n int) string {
	if n <= 0 {
		return s
	}
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[:n])
}

// DescriptionPostProcess 是描述生成的通用清洗管线：
// NFKC → 折空白 → 去引号 → 取首段 → 去末标点 → 截断
type DescriptionPostProcess struct {
	MaxRunes int // 0 时不截断
}

// Apply 跑全套管线，返回清洗后的文本。
func (p DescriptionPostProcess) Apply(s string) string {
	s = NormalizeNFKC(s)
	// 取首段（双换行视作段落分隔；单换行也按段处理为更稳妥）
	if i := strings.IndexAny(s, "\n\r"); i >= 0 {
		s = s[:i]
	}
	s = CollapseWhitespace(s)
	s = StripWrappingQuotes(s)
	s = StripTrailingPunct(s)
	s = TruncateRunes(s, p.MaxRunes)
	return s
}
