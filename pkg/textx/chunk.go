package textx

// Chunk 按 rune 滑窗切块。
//
//	size:    每块最大 rune 数
//	overlap: 与上一块尾部重叠的 rune 数（用于上下文连贯）
//
// 退化情形：
//   - 文本 rune 数 ≤ size：返回单元素 [text]
//   - size ≤ 0：返回单元素 [text]
//   - overlap >= size：overlap 自动收敛为 size/2，避免死循环
func Chunk(text string, size, overlap int) []string {
	if text == "" || size <= 0 {
		if text == "" {
			return nil
		}
		return []string{text}
	}
	if overlap >= size {
		overlap = size / 2
	}
	runes := []rune(text)
	if len(runes) <= size {
		return []string{text}
	}
	step := size - overlap
	if step <= 0 {
		step = size
	}
	var out []string
	for i := 0; i < len(runes); i += step {
		end := i + size
		if end > len(runes) {
			end = len(runes)
		}
		out = append(out, string(runes[i:end]))
		if end == len(runes) {
			break
		}
	}
	return out
}
