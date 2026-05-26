// Package mine 我的 Feed 整页缓存（L1+L2 都缓存完整 FeedPage JSON）。
package mine

import (
	"encoding/json"
	"time"

	"github.com/zhiguang/zhiguang-go/pkg/cachex"
	"github.com/zhiguang/zhiguang-go/pkg/hotkey"
	cachekeys "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/cache"
	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

func New(l1 *cachex.L1, l2 *cachex.L2, hot *hotkey.Detector) cachex.Cache[*pb.FeedPage] {
	return cachex.New[*pb.FeedPage](cachex.Options[*pb.FeedPage]{
		L1: l1, L2: l2,
		BaseTTL:       cachekeys.FeedMineBaseTTL,
		JitterMax:     cachekeys.FeedMineJitterMax,
		NullTTL:       cachekeys.NullTTL,
		NullJitterMax: cachekeys.NullJitterMax,
		Hot:           hot,
		TTLExtender: func(base time.Duration, level hotkey.Level) time.Duration {
			return hotkey.TTLForMine(base, level)
		},
		Marshal: func(v *pb.FeedPage) ([]byte, error) { return json.Marshal(v) },
		Unmarshal: func(b []byte, v **pb.FeedPage) error {
			var x pb.FeedPage
			if err := json.Unmarshal(b, &x); err != nil {
				return err
			}
			*v = &x
			return nil
		},
	})
}

// 让编译器看到 cachekeys 用法（mine 暂未直接使用 keys 包，留作将来扩展）。
var _ = cachekeys.FeedMineBaseTTL
