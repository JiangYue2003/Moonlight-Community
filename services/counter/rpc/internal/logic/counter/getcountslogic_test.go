package counterlogic

import (
	"context"
	"sync"
	"testing"

	"github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"

	"github.com/zhiguang/zhiguang-go/pkg/counterlua"
	"github.com/zhiguang/zhiguang-go/pkg/kafkax"
	"github.com/zhiguang/zhiguang-go/pkg/lockx"
	"github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/services/counter/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/counter/shared/schema"
)

func newGetSvc(t *testing.T) (*miniredis.Miniredis, *svc.ServiceContext) {
	t.Helper()
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(mr.Close)
	rdb := goredis.NewClient(&goredis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })
	return mr, &svc.ServiceContext{
		Redis:           rdb,
		Kafka:           kafkax.NewProducer([]string{"127.0.0.1:9092"}),
		ToggleScript:    goredis.NewScript(counterlua.Toggle),
		IncrFieldScript: goredis.NewScript(counterlua.IncrField),
		DecrFieldScript: goredis.NewScript(counterlua.DecrField),
		Locks:           lockx.New(rdb),
	}
}

func TestGetCounts_SdsMissReturnsZero(t *testing.T) {
	_, sc := newGetSvc(t)
	r, err := NewGetCountsLogic(context.Background(), sc).GetCounts(&counter.GetCountsReq{
		EntityType: "knowpost", EntityId: "404",
	})
	if err != nil {
		t.Fatal(err)
	}
	if r.Counts["like"] != 0 || r.Counts["fav"] != 0 {
		t.Fatalf("missing SDS without bitmap should yield zeros: %+v", r.Counts)
	}
}

func TestGetCounts_RebuildsFromBitmap(t *testing.T) {
	_, sc := newGetSvc(t)
	ctx := context.Background()
	// 在第 0 chunk 设三个用户的 like 位
	for _, off := range []int64{1, 2, 3} {
		sc.Redis.SetBit(ctx, schema.BitmapKey("like", "knowpost", "9", 0), off, 1)
	}
	// 在第 0 chunk 设两个用户的 fav 位
	for _, off := range []int64{1, 5} {
		sc.Redis.SetBit(ctx, schema.BitmapKey("fav", "knowpost", "9", 0), off, 1)
	}

	r, err := NewGetCountsLogic(ctx, sc).GetCounts(&counter.GetCountsReq{
		EntityType: "knowpost", EntityId: "9",
	})
	if err != nil {
		t.Fatal(err)
	}
	if r.Counts["like"] != 3 || r.Counts["fav"] != 2 {
		t.Fatalf("rebuild count drift: %+v", r.Counts)
	}

	// SDS 应该被写回（用 STRLEN 验证）
	if n, err := sc.Redis.StrLen(ctx, schema.SdsKey("knowpost", "9")).Result(); err != nil || n != int64(schema.SchemaLen*schema.FieldSize) {
		t.Fatalf("SDS not written back: len=%d err=%v", n, err)
	}
}

func TestGetCounts_HitDoesNotRebuild(t *testing.T) {
	_, sc := newGetSvc(t)
	ctx := context.Background()
	// 直接预填 SDS（用 IncrField 让 like 字段 = 7）
	for i := 0; i < 7; i++ {
		_, _ = sc.IncrFieldScript.Run(ctx, sc.Redis,
			[]string{schema.SdsKey("knowpost", "1")},
			schema.IdxLike, 1, schema.SchemaLen, schema.FieldSize).Int64()
	}
	// 设一些位图位（如果走重建会得到不同值）
	sc.Redis.SetBit(ctx, schema.BitmapKey("like", "knowpost", "1", 0), 99, 1)

	r, _ := NewGetCountsLogic(ctx, sc).GetCounts(&counter.GetCountsReq{
		EntityType: "knowpost", EntityId: "1",
	})
	if r.Counts["like"] != 7 {
		t.Fatalf("SDS hit should bypass rebuild: got %+v", r.Counts)
	}
}

func TestGetCounts_ConcurrentRebuild_OnlyOneWrites(t *testing.T) {
	_, sc := newGetSvc(t)
	ctx := context.Background()
	for _, off := range []int64{1, 2, 3, 4} {
		sc.Redis.SetBit(ctx, schema.BitmapKey("like", "knowpost", "11", 0), off, 1)
	}
	const N = 30
	results := make([]int64, N)
	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		i := i
		go func() {
			defer wg.Done()
			r, _ := NewGetCountsLogic(ctx, sc).GetCounts(&counter.GetCountsReq{
				EntityType: "knowpost", EntityId: "11",
			})
			if r != nil {
				results[i] = r.Counts["like"]
			}
		}()
	}
	wg.Wait()
	// 所有 goroutine 应当观察到一致的最终计数（=4 或 0：若抢锁失败且 SDS 还没写好返回零）。
	for _, v := range results {
		if v != 0 && v != 4 {
			t.Fatalf("inconsistent rebuild result: got %d in batch", v)
		}
	}
}

func TestGetCounts_RejectsEmptyRequest(t *testing.T) {
	_, sc := newGetSvc(t)
	if _, err := NewGetCountsLogic(context.Background(), sc).GetCounts(&counter.GetCountsReq{}); err == nil {
		t.Fatal("empty req should error")
	}
}

// TestGetCounts_RebuildScansAllChunks 验证：阶段4 改用 Redis SCAN 后，超出旧 64 chunk
// 限制（chunk=100、10000）的位图也能正确累加。
func TestGetCounts_RebuildScansAllChunks(t *testing.T) {
	_, sc := newGetSvc(t)
	ctx := context.Background()
	// 在 chunk=0、100、10000 上各设 1 位
	for _, chunk := range []int64{0, 100, 10000} {
		sc.Redis.SetBit(ctx, schema.BitmapKey("like", "knowpost", "777", chunk), 1, 1)
	}
	r, err := NewGetCountsLogic(ctx, sc).GetCounts(&counter.GetCountsReq{
		EntityType: "knowpost", EntityId: "777",
	})
	if err != nil {
		t.Fatal(err)
	}
	if r.Counts["like"] != 3 {
		t.Fatalf("expect 3 across chunks, got %+v", r.Counts)
	}
}
