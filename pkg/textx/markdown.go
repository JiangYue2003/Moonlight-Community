// Package textx 提供 Markdown 切块与中文文本清洗工具。
//
// 范围：
//   - SplitByHeader：按 Markdown 标题（# / ## / ###...）切段，每段挂最近一级父标题
//   - Chunk：按 rune（不是 byte）做滑窗切块，避免中文截半
//   - 中文清洗：NFKC + 折叠空白 + 去包裹引号 + 去末尾标点 + 截断到 N codepoints
package textx

import (
	"bufio"
	"strings"
)

// Section 一段带标题的内容。
//
// Title 是这段当前最近的 H1/H2/H3...（取最大字号）；为空表示整篇没有标题。
type Section struct {
	Title string
	Body  string
}

// SplitByHeader 把 markdown 按 ATX 风格标题（行首 1~6 个 # 后跟空格）切段。
//
//	如果整篇没有任何标题，返回单元素 [{Title:"", Body: 整文}]。
func SplitByHeader(md string) []Section {
	if md == "" {
		return nil
	}
	sc := bufio.NewScanner(strings.NewReader(md))
	sc.Buffer(make([]byte, 0, 64*1024), 4*1024*1024)
	var sections []Section
	cur := Section{}
	flush := func() {
		body := strings.TrimSpace(cur.Body)
		if cur.Title == "" && body == "" {
			return
		}
		sections = append(sections, Section{Title: cur.Title, Body: body})
	}
	for sc.Scan() {
		line := sc.Text()
		if t, ok := parseHeader(line); ok {
			flush()
			cur = Section{Title: t}
			continue
		}
		cur.Body += line + "\n"
	}
	flush()
	return sections
}

// parseHeader 识别 ATX header（# 后必须跟空白或行尾）。代码块内的 # 不识别（简化：不处理 code fence）。
func parseHeader(line string) (string, bool) {
	s := strings.TrimLeft(line, " \t")
	if !strings.HasPrefix(s, "#") {
		return "", false
	}
	i := 0
	for i < len(s) && s[i] == '#' && i < 6 {
		i++
	}
	if i == 0 {
		return "", false
	}
	if i >= len(s) {
		return "", true
	}
	if s[i] != ' ' && s[i] != '\t' {
		return "", false
	}
	return strings.TrimSpace(s[i+1:]), true
}
