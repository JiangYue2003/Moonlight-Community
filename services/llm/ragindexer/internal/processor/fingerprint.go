package processor

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/zhiguang/zhiguang-go/pkg/esx"
)

// fingerprint 给定文本的稳定 sha256 hex。
func fingerprint(s string) string {
	h := sha256.Sum256([]byte(s))
	return hex.EncodeToString(h[:])
}

// alreadyIndexed 在 ES 里查找 metadata.post_id == postId 的任意文档，
// 取其 content_sha256 与给定 sha256 比较；命中则视为已索引（跳过）。
func alreadyIndexed(ctx context.Context, es *esx.Client, index string, postId int64, sha string) (bool, error) {
	q := json.RawMessage(fmt.Sprintf(`{
        "size": 1,
        "_source": ["content_sha256"],
        "query": {"bool":{"filter":[
            {"term":{"post_id": %d}},
            {"term":{"content_sha256": %q}}
        ]}}
    }`, postId, sha))
	res, err := es.Search(ctx, index, q)
	if err != nil {
		return false, err
	}
	return len(res.Hits.Hits) > 0, nil
}

// deleteAllChunks 把同 postId 下的所有 chunk 删干净。
func deleteAllChunks(ctx context.Context, es *esx.Client, index string, postId int64) error {
	q := json.RawMessage(fmt.Sprintf(`{"query":{"term":{"post_id": %d}}}`, postId))
	_, err := es.DeleteByQuery(ctx, index, q)
	return err
}
