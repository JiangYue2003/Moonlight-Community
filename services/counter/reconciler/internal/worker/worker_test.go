package worker

import (
	"context"
	"encoding/binary"
	"testing"

	miniredis "github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9"

	"github.com/zhiguang/zhiguang-go/services/counter/reconciler/internal/config"
	"github.com/zhiguang/zhiguang-go/services/counter/shared/schema"
)

func newTestWorker(t *testing.T, mr *miniredis.Miniredis) (*Worker, goredis.UniversalClient) {
	t.Helper()
	rdb := goredis.NewUniversalClient(&goredis.UniversalOptions{
		Addrs: []string{mr.Addr()},
	})
	t.Cleanup(func() { _ = rdb.Close() })
	w := New(config.Config{
		Scan: config.ScanConf{
			BatchSize:         256,
			BatchIntervalMs:   0,
			ThresholdAbsolute: 100,
			ThresholdPercent:  1,
		},
	}, rdb)
	return w, rdb
}

func writeSDS(t *testing.T, rdb goredis.UniversalClient, key string, vals [schema.SchemaLen]int64) {
	t.Helper()
	buf := make([]byte, schema.SchemaLen*schema.FieldSize)
	for i, v := range vals {
		binary.BigEndian.PutUint32(buf[i*schema.FieldSize:(i+1)*schema.FieldSize], uint32(v))
	}
	if err := rdb.Set(context.Background(), key, buf, 0).Err(); err != nil {
		t.Fatalf("writeSDS: %v", err)
	}
}

func writeBitmap(t *testing.T, rdb goredis.UniversalClient, key string, bits ...int) {
	t.Helper()
	ctx := context.Background()
	for _, b := range bits {
		if err := rdb.SetBit(ctx, key, int64(b), 1).Err(); err != nil {
			t.Fatalf("SetBit key=%s bit=%d: %v", key, b, err)
		}
	}
}

func TestReconcile_NoDeviation_NoRebuild(t *testing.T) {
	mr, _ := miniredis.Run()
	defer mr.Close()
	w, rdb := newTestWorker(t, mr)

	var vals [schema.SchemaLen]int64
	vals[schema.IdxLike] = 3
	writeSDS(t, rdb, "cnt:v1:post:1", vals)
	writeBitmap(t, rdb, "bm:like:post:1:0", 0, 1, 2)

	if err := w.reconcile(context.Background()); err != nil {
		t.Fatal(err)
	}
	raw, _ := rdb.Get(context.Background(), "cnt:v1:post:1").Bytes()
	got := int64(binary.BigEndian.Uint32(raw[schema.IdxLike*schema.FieldSize:]))
	if got != 3 {
		t.Errorf("expected sds=3, got %d", got)
	}
}

func TestReconcile_LargeDeviation_Rebuilds(t *testing.T) {
	mr, _ := miniredis.Run()
	defer mr.Close()
	w, rdb := newTestWorker(t, mr)

	var vals [schema.SchemaLen]int64
	vals[schema.IdxLike] = 1000
	writeSDS(t, rdb, "cnt:v1:post:2", vals)
	writeBitmap(t, rdb, "bm:like:post:2:0", 0)

	if err := w.reconcile(context.Background()); err != nil {
		t.Fatal(err)
	}
	raw, _ := rdb.Get(context.Background(), "cnt:v1:post:2").Bytes()
	got := int64(binary.BigEndian.Uint32(raw[schema.IdxLike*schema.FieldSize:]))
	if got != 1 {
		t.Errorf("expected rebuilt sds=1, got %d", got)
	}
}

func TestReconcile_SmallDeviation_NoRebuild(t *testing.T) {
	mr, _ := miniredis.Run()
	defer mr.Close()
	w, rdb := newTestWorker(t, mr)

	// SDS like=100, bitmap like=99 → 偏差 1 < 阈值 100 且 1% < 1% → 不重建
	var vals [schema.SchemaLen]int64
	vals[schema.IdxLike] = 100
	writeSDS(t, rdb, "cnt:v1:post:3", vals)
	for i := 0; i < 99; i++ {
		writeBitmap(t, rdb, "bm:like:post:3:0", i)
	}

	if err := w.reconcile(context.Background()); err != nil {
		t.Fatal(err)
	}
	raw, _ := rdb.Get(context.Background(), "cnt:v1:post:3").Bytes()
	got := int64(binary.BigEndian.Uint32(raw[schema.IdxLike*schema.FieldSize:]))
	if got != 100 {
		t.Errorf("expected sds unchanged=100, got %d", got)
	}
}

func TestReconcile_NoKeys_NoError(t *testing.T) {
	mr, _ := miniredis.Run()
	defer mr.Close()
	w, _ := newTestWorker(t, mr)

	if err := w.reconcile(context.Background()); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
