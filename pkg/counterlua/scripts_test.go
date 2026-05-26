package counterlua

import (
	"context"
	"encoding/binary"
	"strconv"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

const (
	schemaLen = 5
	fieldSize = 4
	idxLike   = 1
	idxFav    = 2
)

func newRedis(t *testing.T) (*miniredis.Miniredis, *redis.Client) {
	t.Helper()
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("miniredis start: %v", err)
	}
	t.Cleanup(mr.Close)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })
	return mr, rdb
}

// readFieldUint32 读 SDS 中 idx 位置的 uint32（BigEndian）。
func readFieldUint32(t *testing.T, raw []byte, idx int) uint32 {
	t.Helper()
	off := idx * fieldSize
	if off+fieldSize > len(raw) {
		t.Fatalf("buffer too short: len=%d need %d", len(raw), off+fieldSize)
	}
	return binary.BigEndian.Uint32(raw[off : off+fieldSize])
}

// ---------- toggle.lua ----------

func TestToggle_AddSetsBitAndReturns1(t *testing.T) {
	_, rdb := newRedis(t)
	ctx := context.Background()
	script := redis.NewScript(Toggle)

	got, err := script.Run(ctx, rdb, []string{"bm:like:knowpost:1:0"}, 42, "add").Int64()
	if err != nil {
		t.Fatal(err)
	}
	if got != 1 {
		t.Fatalf("first add should return 1 (changed), got %d", got)
	}
	bit, _ := rdb.GetBit(ctx, "bm:like:knowpost:1:0", 42).Result()
	if bit != 1 {
		t.Fatalf("bit not set after add")
	}
}

func TestToggle_AddIsIdempotent(t *testing.T) {
	_, rdb := newRedis(t)
	ctx := context.Background()
	script := redis.NewScript(Toggle)
	key := "bm:like:knowpost:1:0"

	// 第一次 add → 1
	first, _ := script.Run(ctx, rdb, []string{key}, 42, "add").Int64()
	// 第二次 add → 0（已置位，幂等）
	second, _ := script.Run(ctx, rdb, []string{key}, 42, "add").Int64()
	if first != 1 || second != 0 {
		t.Fatalf("idempotency violated: first=%d second=%d", first, second)
	}
}

func TestToggle_RemoveClearsBitAndReturns1(t *testing.T) {
	_, rdb := newRedis(t)
	ctx := context.Background()
	script := redis.NewScript(Toggle)
	key := "bm:like:knowpost:1:0"
	_, _ = script.Run(ctx, rdb, []string{key}, 42, "add").Int64()

	got, err := script.Run(ctx, rdb, []string{key}, 42, "remove").Int64()
	if err != nil {
		t.Fatal(err)
	}
	if got != 1 {
		t.Fatalf("remove of set bit should return 1, got %d", got)
	}
	bit, _ := rdb.GetBit(ctx, key, 42).Result()
	if bit != 0 {
		t.Fatalf("bit not cleared after remove")
	}
}

func TestToggle_RemoveOnUnsetIsNoop(t *testing.T) {
	_, rdb := newRedis(t)
	ctx := context.Background()
	script := redis.NewScript(Toggle)
	got, err := script.Run(ctx, rdb, []string{"bm:like:knowpost:9:0"}, 100, "remove").Int64()
	if err != nil {
		t.Fatal(err)
	}
	if got != 0 {
		t.Fatalf("remove on unset bit should be noop (0), got %d", got)
	}
}

func TestToggle_DoesNotInterfereWithNeighbouringBits(t *testing.T) {
	_, rdb := newRedis(t)
	ctx := context.Background()
	script := redis.NewScript(Toggle)
	key := "bm:like:knowpost:1:0"

	for _, off := range []int{0, 1, 7, 100, 32767} {
		_, _ = script.Run(ctx, rdb, []string{key}, off, "add").Int64()
	}
	// remove 一个不应误伤其他
	_, _ = script.Run(ctx, rdb, []string{key}, 7, "remove").Int64()

	for _, off := range []int{0, 1, 100, 32767} {
		bit, _ := rdb.GetBit(ctx, key, int64(off)).Result()
		if bit != 1 {
			t.Errorf("offset %d should remain set", off)
		}
	}
	bit, _ := rdb.GetBit(ctx, key, 7).Result()
	if bit != 0 {
		t.Errorf("offset 7 should be cleared")
	}
}

// ---------- incr_field.lua ----------

func TestIncrField_InitializesEmptyKeyToZeroBuffer(t *testing.T) {
	_, rdb := newRedis(t)
	ctx := context.Background()
	script := redis.NewScript(IncrField)
	key := "cnt:v1:knowpost:1"

	val, err := script.Run(ctx, rdb,
		[]string{key}, idxLike, 1, schemaLen, fieldSize).Int64()
	if err != nil {
		t.Fatal(err)
	}
	if val != 1 {
		t.Fatalf("first incr should return 1, got %d", val)
	}
	raw, err := rdb.Get(ctx, key).Bytes()
	if err != nil {
		t.Fatal(err)
	}
	if len(raw) != schemaLen*fieldSize {
		t.Fatalf("buffer size drift: want %d, got %d", schemaLen*fieldSize, len(raw))
	}
	if v := readFieldUint32(t, raw, idxLike); v != 1 {
		t.Fatalf("idxLike should be 1, got %d", v)
	}
	// 其它字段必须保持 0
	for _, idx := range []int{0, 2, 3, 4} {
		if v := readFieldUint32(t, raw, idx); v != 0 {
			t.Errorf("field %d should remain 0, got %d", idx, v)
		}
	}
}

func TestIncrField_AccumulatesAcrossCalls(t *testing.T) {
	_, rdb := newRedis(t)
	ctx := context.Background()
	script := redis.NewScript(IncrField)
	key := "cnt:v1:knowpost:1"

	for i := 0; i < 7; i++ {
		_, _ = script.Run(ctx, rdb, []string{key}, idxLike, 3, schemaLen, fieldSize).Int64()
	}
	raw, _ := rdb.Get(ctx, key).Bytes()
	if v := readFieldUint32(t, raw, idxLike); v != 21 {
		t.Fatalf("expected 21 after 7×3, got %d", v)
	}
}

func TestIncrField_HandlesDifferentIndexesIndependently(t *testing.T) {
	_, rdb := newRedis(t)
	ctx := context.Background()
	script := redis.NewScript(IncrField)
	key := "cnt:v1:knowpost:1"

	_, _ = script.Run(ctx, rdb, []string{key}, idxLike, 5, schemaLen, fieldSize).Int64()
	_, _ = script.Run(ctx, rdb, []string{key}, idxFav, 8, schemaLen, fieldSize).Int64()

	raw, _ := rdb.Get(ctx, key).Bytes()
	if v := readFieldUint32(t, raw, idxLike); v != 5 {
		t.Errorf("idxLike: expected 5, got %d", v)
	}
	if v := readFieldUint32(t, raw, idxFav); v != 8 {
		t.Errorf("idxFav: expected 8, got %d", v)
	}
}

func TestIncrField_NegativeDeltaSaturatesAtZero(t *testing.T) {
	// SDS 的字段在 Lua 内有 saturate 0 的逻辑：不允许出现负值。
	_, rdb := newRedis(t)
	ctx := context.Background()
	script := redis.NewScript(IncrField)
	key := "cnt:v1:knowpost:1"

	_, _ = script.Run(ctx, rdb, []string{key}, idxLike, 2, schemaLen, fieldSize).Int64()
	val, err := script.Run(ctx, rdb,
		[]string{key}, idxLike, -10, schemaLen, fieldSize).Int64()
	if err != nil {
		t.Fatal(err)
	}
	if val != 0 {
		t.Fatalf("saturated decrement should yield 0, got %d", val)
	}
	raw, _ := rdb.Get(ctx, key).Bytes()
	if v := readFieldUint32(t, raw, idxLike); v != 0 {
		t.Fatalf("stored field should be 0, got %d", v)
	}
}

func TestIncrField_BigEndianRoundTrip(t *testing.T) {
	_, rdb := newRedis(t)
	ctx := context.Background()
	script := redis.NewScript(IncrField)
	key := "cnt:v1:knowpost:1"

	// 写入 0x01020304 = 16909060
	_, _ = script.Run(ctx, rdb, []string{key}, 0, 16909060, schemaLen, fieldSize).Int64()
	raw, _ := rdb.Get(ctx, key).Bytes()
	want := []byte{0x01, 0x02, 0x03, 0x04}
	for i, b := range want {
		if raw[i] != b {
			t.Fatalf("byte %d: want 0x%02x got 0x%02x (full=%x)", i, b, raw[i], raw)
		}
	}
}

// ---------- decr_field.lua ----------

func TestDecrField_DecrementsAndReturnsRemaining(t *testing.T) {
	_, rdb := newRedis(t)
	ctx := context.Background()
	script := redis.NewScript(DecrField)
	key := "agg:v1:knowpost:1"

	rdb.HSet(ctx, key, "1", "10")
	left, err := script.Run(ctx, rdb, []string{key}, "1", 3).Int64()
	if err != nil {
		t.Fatal(err)
	}
	if left != 7 {
		t.Fatalf("HINCRBY -3 from 10 should leave 7, got %d", left)
	}
}

func TestDecrField_DeletesFieldWhenReachingZero(t *testing.T) {
	_, rdb := newRedis(t)
	ctx := context.Background()
	script := redis.NewScript(DecrField)
	key := "agg:v1:knowpost:1"

	rdb.HSet(ctx, key, "1", "5")
	left, err := script.Run(ctx, rdb, []string{key}, "1", 5).Int64()
	if err != nil {
		t.Fatal(err)
	}
	if left != 0 {
		t.Fatalf("expected 0 after full drain, got %d", left)
	}
	exists, _ := rdb.HExists(ctx, key, "1").Result()
	if exists {
		t.Fatal("field should be HDEL'd when drained to 0")
	}
}

func TestDecrField_OnAbsentFieldYieldsNegativeAndDoesNotPanic(t *testing.T) {
	// 没有此字段时 HINCRBY 会从 0 开始减；脚本应该返回 -delta 且不出错。
	_, rdb := newRedis(t)
	ctx := context.Background()
	script := redis.NewScript(DecrField)
	key := "agg:v1:knowpost:1"

	left, err := script.Run(ctx, rdb, []string{key}, "1", 4).Int64()
	if err != nil {
		t.Fatal(err)
	}
	if left != -4 {
		t.Fatalf("absent field decr -4 should yield -4, got %d", left)
	}
}

// ---------- 综合：toggle + incr_field 模拟一次点赞流水 ----------

func TestEndToEnd_ToggleThenIncrFieldFlow(t *testing.T) {
	_, rdb := newRedis(t)
	ctx := context.Background()
	toggle := redis.NewScript(Toggle)
	incr := redis.NewScript(IncrField)

	bm := "bm:like:knowpost:1:0"
	sds := "cnt:v1:knowpost:1"

	// 模拟 Toggle 返回 changed=1 后，aggregator 把 +1 直接折算到 SDS 的简化路径。
	for _, uid := range []int{1, 2, 3} {
		changed, _ := toggle.Run(ctx, rdb, []string{bm}, uid, "add").Int64()
		if changed == 1 {
			_, _ = incr.Run(ctx, rdb, []string{sds}, idxLike, 1, schemaLen, fieldSize).Int64()
		}
	}
	// 再次重复 add：幂等不计数
	for _, uid := range []int{1, 2, 3} {
		changed, _ := toggle.Run(ctx, rdb, []string{bm}, uid, "add").Int64()
		if changed != 0 {
			t.Fatalf("uid %d second add should be idempotent (0), got %d", uid, changed)
		}
	}

	raw, _ := rdb.Get(ctx, sds).Bytes()
	if v := readFieldUint32(t, raw, idxLike); v != 3 {
		t.Fatalf("expected 3 distinct likes, got %d", v)
	}
	cnt, _ := rdb.BitCount(ctx, bm, nil).Result()
	if cnt != 3 {
		t.Fatalf("BITCOUNT should equal SDS like field, got %d (SDS=%d)", cnt, 3)
	}

	// 取消两个再次切换，SDS 由 aggregator 折算（脚本本身只负责加减）
	_, _ = toggle.Run(ctx, rdb, []string{bm}, 1, "remove").Int64()
	_, _ = toggle.Run(ctx, rdb, []string{bm}, 2, "remove").Int64()
	_, _ = incr.Run(ctx, rdb, []string{sds}, idxLike, -2, schemaLen, fieldSize).Int64()
	raw, _ = rdb.Get(ctx, sds).Bytes()
	if v := readFieldUint32(t, raw, idxLike); v != 1 {
		t.Fatalf("expected 1 after 2 removes, got %d", v)
	}
}

// 防止 strconv 被静态优化为未使用导入（用于 future 测试拓展）
var _ = strconv.Itoa
