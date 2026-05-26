// Package detail 提供 knowpost 详情的三级缓存读取入口。
//
// 这是 cachex.Cache[*KnowPostDetail] 的最薄包装：
//   - GetOrLoad：L1 → L2 → loader（业务侧）
//   - Invalidate：双删 L2/L1
//   - PutNullSentinel：DB 不存在时写哨兵
//   - SetWithExtension：命中后按 hot level 延长 TTL
package detail

import (
	"encoding/json"
	"time"

	"github.com/zhiguang/zhiguang-go/pkg/cachex"
	"github.com/zhiguang/zhiguang-go/pkg/hotkey"
	cachekeys "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache"
	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

// New 构造详情专用 Cache。L1 / L2 由 svc 注入；hot 可选（nil 时不延长 TTL）。
func New(l1 *cachex.L1, l2 *cachex.L2, hot *hotkey.Detector) cachex.Cache[*pb.KnowPostDetail] {
	return cachex.New[*pb.KnowPostDetail](cachex.Options[*pb.KnowPostDetail]{
		L1:            l1,
		L2:            l2,
		BaseTTL:       cachekeys.DetailBaseTTL,
		JitterMax:     cachekeys.DetailJitterMax,
		NullTTL:       cachekeys.NullTTL,
		NullJitterMax: cachekeys.NullJitterMax,
		Hot:           hot,
		HotKeyFor:     func(k string) string { return k },
		TTLExtender: func(base time.Duration, level hotkey.Level) time.Duration {
			return hotkey.TTLForPublic(base, level)
		},
		Marshal: func(v *pb.KnowPostDetail) ([]byte, error) { return json.Marshal(v) },
		Unmarshal: func(b []byte, v **pb.KnowPostDetail) error {
			var x pb.KnowPostDetail
			if err := json.Unmarshal(b, &x); err != nil {
				return err
			}
			*v = &x
			return nil
		},
		CostFn: func(_ *pb.KnowPostDetail, raw []byte) int64 {
			if len(raw) == 0 {
				return 1
			}
			return int64(len(raw))
		},
	})
}
