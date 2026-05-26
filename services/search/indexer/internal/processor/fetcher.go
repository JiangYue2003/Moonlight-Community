package processor

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/zhiguang/zhiguang-go/pkg/textx"
)

// fetchContent HTTP GET contentUrl 然后 NFKC 归一化 + 截断到 maxRunes。
//
// 失败返回空串与 error，调用方决定是否跳过这条事件。
func fetchContent(ctx context.Context, hc *http.Client, url string, maxRunes int) (string, error) {
	if url == "" {
		return "", nil
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	resp, err := hc.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("fetch %s: %d", url, resp.StatusCode)
	}
	buf, err := io.ReadAll(io.LimitReader(resp.Body, 4<<20)) // 安全上限 4MB
	if err != nil {
		return "", err
	}
	s := textx.NormalizeNFKC(string(buf))
	return textx.TruncateRunes(s, maxRunes), nil
}
