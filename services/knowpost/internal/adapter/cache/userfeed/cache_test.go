package userfeed

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/zhiguang/zhiguang-go/pkg/cachex"
	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

type fakeL1 struct {
	mu       sync.Mutex
	values   map[string]any
	gets     int
	sets     int
	lastKey  string
	lastVal  any
	lastCost int64
	lastTTL  time.Duration
}

func (f *fakeL1) Get(key string) (any, bool) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.gets++
	value, ok := f.values[key]
	return value, ok
}

func (f *fakeL1) SetWithTTL(key string, value any, cost int64, ttl time.Duration) bool {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.sets++
	f.lastKey, f.lastVal, f.lastCost, f.lastTTL = key, value, cost, ttl
	if f.values == nil {
		f.values = make(map[string]any)
	}
	f.values[key] = value
	return true
}

type fakeL2 struct {
	mu      sync.Mutex
	values  map[string][]byte
	getErr  error
	setErr  error
	gets    int
	sets    int
	lastKey string
	lastVal []byte
	lastTTL time.Duration
}

func (f *fakeL2) Get(_ context.Context, key string) ([]byte, bool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.gets++
	if f.getErr != nil {
		return nil, false, f.getErr
	}
	value, ok := f.values[key]
	return append([]byte(nil), value...), ok, nil
}

func (f *fakeL2) Set(_ context.Context, key string, value []byte, ttl time.Duration) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.sets++
	f.lastKey, f.lastVal, f.lastTTL = key, append([]byte(nil), value...), ttl
	if f.setErr != nil {
		return f.setErr
	}
	if f.values == nil {
		f.values = make(map[string][]byte)
	}
	f.values[key] = append([]byte(nil), value...)
	return nil
}

func TestKeyContainsEveryInvalidationDimension(t *testing.T) {
	req := Request{
		UserID:          42,
		StrategyVersion: "hybrid-v1",
		RelationEpoch:   7,
		SafetyEpoch:     9,
		Page:            1,
		Size:            20,
	}
	require.Equal(t, "feed:page:v1:hybrid-v1:42:r7:s9:p1:n20", Key("feed", req))
	require.NotEqual(t, Key("feed", req), Key("feed", withRequest(req, func(r *Request) { r.UserID++ })))
	require.NotEqual(t, Key("feed", req), Key("feed", withRequest(req, func(r *Request) { r.StrategyVersion = "hybrid-v2" })))
	require.NotEqual(t, Key("feed", req), Key("feed", withRequest(req, func(r *Request) { r.RelationEpoch++ })))
	require.NotEqual(t, Key("feed", req), Key("feed", withRequest(req, func(r *Request) { r.SafetyEpoch++ })))
	require.NotEqual(t, Key("feed", req), Key("feed", withRequest(req, func(r *Request) { r.Page++ })))
	require.NotEqual(t, Key("feed", req), Key("feed", withRequest(req, func(r *Request) { r.Size++ })))
}

func TestGetOrLoadBypassesIneligibleRequests(t *testing.T) {
	for _, tc := range []struct {
		name string
		req  Request
	}{
		{name: "non hybrid", req: cacheableRequest(func(r *Request) { r.StrategyVersion = "pull-v1" })},
		{name: "other page", req: cacheableRequest(func(r *Request) { r.Page = 2 })},
		{name: "other size", req: cacheableRequest(func(r *Request) { r.Size = 10 })},
	} {
		t.Run(tc.name, func(t *testing.T) {
			l1, l2 := &fakeL1{}, &fakeL2{}
			cache := mustCache(t, Config{Mode: ModeL1L2, KeyPrefix: "feed", L1FreshTTL: time.Second, L2FreshTTL: 5 * time.Second}, l1, l2)
			loads := 0
			got, source, err := cache.GetOrLoad(context.Background(), tc.req, func(context.Context) (*pb.FeedPage, error) {
				loads++
				return testPage(1), nil
			})
			require.NoError(t, err)
			require.Equal(t, SourceBypass, source)
			require.Equal(t, int32(1), got.Page)
			require.Equal(t, 1, loads)
			require.Zero(t, l1.gets+l1.sets+l2.gets+l2.sets)
		})
	}
}

func TestGetOrLoadL1FreshHitSkipsL2AndLoader(t *testing.T) {
	req := cacheableRequest(nil)
	key := Key("feed", req)
	want := testPage(1)
	now := time.Now()
	l1, l2 := &fakeL1{values: map[string]any{key: &cacheRecord{
		page: want, freshUntil: now.Add(time.Second), refreshAt: now.Add(800 * time.Millisecond), staleUntil: now.Add(11 * time.Second),
	}}}, &fakeL2{}
	cache := mustCache(t, Config{Mode: ModeL1L2, KeyPrefix: "feed", L1FreshTTL: time.Second, L2FreshTTL: 5 * time.Second}, l1, l2)

	got, source, err := cache.GetOrLoad(context.Background(), req, loaderMustNotRun(t))
	require.NoError(t, err)
	require.Equal(t, SourceL1Fresh, source)
	require.Same(t, want, got)
	require.Equal(t, 1, l1.gets)
	require.Zero(t, l2.gets+l2.sets)
}

func TestGetOrLoadL2FreshHitDecodesAndFillsL1(t *testing.T) {
	req := cacheableRequest(nil)
	key := Key("feed", req)
	want := testPage(8)
	now := time.Now()
	raw, err := encodeRecord(cacheRecord{
		page: want, freshUntil: now.Add(5 * time.Second), refreshAt: now.Add(4 * time.Second), staleUntil: now.Add(15 * time.Second),
	})
	require.NoError(t, err)
	l1, l2 := &fakeL1{}, &fakeL2{values: map[string][]byte{key: raw}}
	cache := mustCache(t, Config{Mode: ModeL1L2, KeyPrefix: "feed", L1FreshTTL: time.Second, L2FreshTTL: 5 * time.Second}, l1, l2)

	got, source, err := cache.GetOrLoad(context.Background(), req, loaderMustNotRun(t))
	require.NoError(t, err)
	require.Equal(t, SourceL2Fresh, source)
	require.True(t, proto.Equal(want, got))
	require.Equal(t, 1, l2.gets)
	require.Equal(t, 1, l1.sets)
	require.Equal(t, int64(2*proto.Size(got)+len(key)+256), l1.lastCost)
	require.Equal(t, 800*time.Millisecond, l1.lastTTL)
}

func TestGetOrLoadMissLoadsAndFillsL2ThenL1(t *testing.T) {
	req := cacheableRequest(nil)
	want := testPage(3)
	l1, l2 := &fakeL1{}, &fakeL2{}
	cache := mustCache(t, Config{Mode: ModeL1L2, KeyPrefix: "feed", L1FreshTTL: time.Second, L2FreshTTL: 5 * time.Second}, l1, l2)
	loads := 0

	got, source, err := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
		loads++
		return want, nil
	})
	require.NoError(t, err)
	require.Equal(t, SourceLoad, source)
	require.Same(t, want, got)
	require.Equal(t, 1, loads)
	require.Equal(t, 1, l2.sets)
	require.Equal(t, 14*time.Second, l2.lastTTL)
	require.Equal(t, 1, l1.sets)
	decoded, decodeErr := decodeRecord(l2.lastVal)
	require.NoError(t, decodeErr)
	require.True(t, proto.Equal(want, decoded.page))
}

func TestGetOrLoadDoesNotPolluteCacheOnLoaderFailureOrNilPage(t *testing.T) {
	for _, tc := range []struct {
		name   string
		loader Loader
	}{
		{name: "error", loader: func(context.Context) (*pb.FeedPage, error) { return nil, errors.New("load") }},
		{name: "nil page", loader: func(context.Context) (*pb.FeedPage, error) { return nil, nil }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			l1, l2 := &fakeL1{}, &fakeL2{}
			cache := mustCache(t, Config{Mode: ModeL1L2, KeyPrefix: "feed", L1FreshTTL: time.Second, L2FreshTTL: 5 * time.Second}, l1, l2)
			got, source, err := cache.GetOrLoad(context.Background(), cacheableRequest(nil), tc.loader)
			require.Error(t, err)
			require.Nil(t, got)
			require.Equal(t, SourceLoad, source)
			require.Zero(t, l1.sets+l2.sets)
		})
	}
}

func TestGetOrLoadCorruptL2FallsBackToLoaderAndOverwritesIt(t *testing.T) {
	req := cacheableRequest(nil)
	key := Key("feed", req)
	l1, l2 := &fakeL1{}, &fakeL2{values: map[string][]byte{key: {0xff, 0x00}}}
	cache := mustCache(t, Config{Mode: ModeL1L2, KeyPrefix: "feed", L1FreshTTL: time.Second, L2FreshTTL: 5 * time.Second}, l1, l2)

	got, source, err := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
		return testPage(4), nil
	})
	require.NoError(t, err)
	require.Equal(t, SourceLoad, source)
	require.Equal(t, int32(4), got.Page)
	require.Equal(t, 1, l2.sets)
	require.Equal(t, 1, l1.sets)
}

func TestGetOrLoadL2FailuresFallBackWithoutFailingRequest(t *testing.T) {
	for _, tc := range []struct {
		name   string
		getErr error
		setErr error
	}{
		{name: "read", getErr: errors.New("redis read")},
		{name: "write", setErr: errors.New("redis write")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			l1, l2 := &fakeL1{}, &fakeL2{getErr: tc.getErr, setErr: tc.setErr}
			cache := mustCache(t, Config{Mode: ModeL1L2, KeyPrefix: "feed", L1FreshTTL: time.Second, L2FreshTTL: 5 * time.Second}, l1, l2)
			got, source, err := cache.GetOrLoad(context.Background(), cacheableRequest(nil), func(context.Context) (*pb.FeedPage, error) {
				return testPage(6), nil
			})
			require.NoError(t, err)
			require.Equal(t, SourceLoad, source)
			require.Equal(t, int32(6), got.Page)
			require.Equal(t, 1, l1.sets)
		})
	}
}

func TestGetOrLoadObservesL2GetDecodeAndSetFailures(t *testing.T) {
	type event struct {
		operation Operation
		err       bool
	}
	for _, tc := range []struct {
		name      string
		l2        *fakeL2
		wantError Operation
	}{
		{name: "get", l2: &fakeL2{getErr: errors.New("get")}, wantError: OperationL2Get},
		{name: "decode", l2: &fakeL2{values: map[string][]byte{Key("feed", cacheableRequest(nil)): {0xff}}}, wantError: OperationL2Decode},
		{name: "set", l2: &fakeL2{setErr: errors.New("set")}, wantError: OperationL2Set},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var events []event
			var reports []Operation
			cache := mustCache(t, Config{
				Mode: ModeL1L2, KeyPrefix: "feed", L1FreshTTL: time.Second, L2FreshTTL: 5 * time.Second,
				Observe: func(operation Operation, err error) {
					events = append(events, event{operation: operation, err: err != nil})
				},
				ReportError: func(operation Operation, _ error) {
					reports = append(reports, operation)
				},
			}, &fakeL1{}, tc.l2)

			page, source, err := cache.GetOrLoad(context.Background(), cacheableRequest(nil), func(context.Context) (*pb.FeedPage, error) {
				return testPage(1), nil
			})
			require.NoError(t, err)
			require.NotNil(t, page)
			require.Equal(t, SourceLoad, source)
			require.Contains(t, events, event{operation: tc.wantError, err: true})
			require.Equal(t, []Operation{tc.wantError}, reports)
		})
	}
}

func TestPageCacheErrorReporterIsRateLimitedPerOperation(t *testing.T) {
	reports := 0
	cache := mustCache(t, Config{
		Mode: ModeL2, KeyPrefix: "feed", L2FreshTTL: 5 * time.Second, ErrorLogInterval: time.Hour,
		ReportError: func(Operation, error) { reports++ },
	}, nil, &fakeL2{getErr: errors.New("redis")})

	for i := 0; i < 10; i++ {
		_, _, err := cache.GetOrLoad(context.Background(), cacheableRequest(nil), func(context.Context) (*pb.FeedPage, error) {
			return testPage(1), nil
		})
		require.NoError(t, err)
	}
	require.Equal(t, 1, reports)
}

func TestModeL2NeverTouchesL1(t *testing.T) {
	l1, l2 := &fakeL1{}, &fakeL2{}
	cache := mustCache(t, Config{Mode: ModeL2, KeyPrefix: "feed", L2FreshTTL: 5 * time.Second}, l1, l2)
	_, source, err := cache.GetOrLoad(context.Background(), cacheableRequest(nil), func(context.Context) (*pb.FeedPage, error) {
		return testPage(1), nil
	})
	require.NoError(t, err)
	require.Equal(t, SourceLoad, source)
	require.Zero(t, l1.gets+l1.sets)
	require.Equal(t, 1, l2.gets)
	require.Equal(t, 1, l2.sets)
}

func TestNewRejectsInvalidConfiguration(t *testing.T) {
	for _, cfg := range []Config{
		{Mode: "unknown", L1FreshTTL: time.Second, L2FreshTTL: 5 * time.Second},
		{Mode: ModeL2, L2FreshTTL: 0},
		{Mode: ModeL2, L2FreshTTL: 5*time.Second + time.Nanosecond},
		{Mode: ModeL2, L2FreshTTL: 5 * time.Second},
		{Mode: ModeL2, L2FreshTTL: 2 * time.Second},
		{Mode: ModeL1L2, L1FreshTTL: 0, L2FreshTTL: 5 * time.Second},
		{Mode: ModeL1L2, L1FreshTTL: time.Second + time.Nanosecond, L2FreshTTL: 5 * time.Second},
		{Mode: ModeL1L2, L1FreshTTL: time.Second, L2FreshTTL: 4 * time.Second},
		{Mode: ModeL1L2, L1FreshTTL: 400 * time.Millisecond, L2FreshTTL: 4 * time.Second},
		{Mode: ModeL2, L2FreshTTL: 5 * time.Second, StaleTTL: -time.Second},
		{Mode: ModeL2, L2FreshTTL: 5 * time.Second, StaleTTL: 10*time.Second + time.Nanosecond},
		{Mode: ModeL2, L2FreshTTL: 5 * time.Second, JitterPercent: -1},
		{Mode: ModeL2, L2FreshTTL: 4 * time.Second, JitterPercent: 21},
		{Mode: ModeL2, L2FreshTTL: 5 * time.Second, RefreshWorkers: -1},
		{Mode: ModeL2, L2FreshTTL: 5 * time.Second, RefreshQueue: -1},
		{Mode: ModeL2, L2FreshTTL: 5 * time.Second, LoaderTimeout: -time.Second},
	} {
		_, err := New(cfg, &fakeL1{}, &fakeL2{})
		require.Error(t, err, "cfg=%+v", cfg)
	}
}

func TestConcreteL1L2AdaptersFillAndServeWithoutRedisOnL1Hit(t *testing.T) {
	redisServer := miniredis.RunT(t)
	redisClient := redis.NewClient(&redis.Options{Addr: redisServer.Addr()})
	t.Cleanup(func() { _ = redisClient.Close() })
	l1, err := cachex.NewL1(cachex.L1Config{NumCounters: 1000, MaxCost: 1 << 20})
	require.NoError(t, err)
	l2 := cachex.NewL2(redisClient)
	cache := mustCache(t, Config{
		Mode: ModeL1L2, KeyPrefix: "test", L1FreshTTL: time.Second, L2FreshTTL: 5 * time.Second,
	}, l1, l2)
	req := cacheableRequest(nil)
	loads := 0

	first, source, err := cache.GetOrLoad(context.Background(), req, func(context.Context) (*pb.FeedPage, error) {
		loads++
		return testPage(1), nil
	})
	require.NoError(t, err)
	require.Equal(t, SourceLoad, source)
	require.NotNil(t, first)
	l1.Wait()
	require.NoError(t, redisClient.Close())

	second, source, err := cache.GetOrLoad(context.Background(), req, loaderMustNotRun(t))
	require.NoError(t, err)
	require.Equal(t, SourceL1Fresh, source)
	require.Same(t, first, second)
	require.Equal(t, 1, loads)
}

func TestCachedPageSupportsConcurrentReadOnlyMarshal(t *testing.T) {
	req := cacheableRequest(nil)
	want := testPage(1)
	now := time.Now()
	l1 := &fakeL1{values: map[string]any{Key("feed", req): &cacheRecord{
		page: want, freshUntil: now.Add(time.Second), refreshAt: now.Add(800 * time.Millisecond), staleUntil: now.Add(11 * time.Second),
	}}}
	cache := mustCache(t, Config{
		Mode: ModeL1L2, KeyPrefix: "feed", L1FreshTTL: time.Second, L2FreshTTL: 5 * time.Second,
	}, l1, &fakeL2{})

	const readers = 64
	var wg sync.WaitGroup
	errs := make(chan error, readers)
	wg.Add(readers)
	for i := 0; i < readers; i++ {
		go func() {
			defer wg.Done()
			page, source, err := cache.GetOrLoad(context.Background(), req, loaderMustNotRun(t))
			if err != nil {
				errs <- err
				return
			}
			if source != SourceL1Fresh || page != want {
				errs <- errors.New("unexpected concurrent cache result")
				return
			}
			if _, err := proto.Marshal(page); err != nil {
				errs <- err
			}
		}()
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		require.NoError(t, err)
	}
}

func mustCache(t *testing.T, cfg Config, l1 L1, l2 L2) Cache {
	t.Helper()
	// Older WP7 tests specify the pre-jitter hard ceilings. WP9 uses nominal
	// bases whose ±20% effective windows remain inside those ceilings.
	if cfg.L1FreshTTL == time.Second {
		cfg.L1FreshTTL = 800 * time.Millisecond
	}
	if cfg.L2FreshTTL == 5*time.Second {
		cfg.L2FreshTTL = 4 * time.Second
	}
	if cfg.RandomFloat64 == nil {
		cfg.RandomFloat64 = func() float64 { return 0.5 }
	}
	cache, err := New(cfg, l1, l2)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, cache.Close()) })
	return cache
}

func cacheableRequest(mutate func(*Request)) Request {
	req := Request{UserID: 42, StrategyVersion: "hybrid-v1", RelationEpoch: 7, SafetyEpoch: 9, Page: 1, Size: 20}
	if mutate != nil {
		mutate(&req)
	}
	return req
}

func withRequest(req Request, mutate func(*Request)) Request {
	mutate(&req)
	return req
}

func testPage(page int32) *pb.FeedPage {
	return &pb.FeedPage{Page: page, Size: 20, Items: []*pb.FeedItem{{Id: "1", Title: "cached"}}}
}

func loaderMustNotRun(t *testing.T) Loader {
	t.Helper()
	return func(context.Context) (*pb.FeedPage, error) {
		t.Fatal("loader must not run")
		return nil, nil
	}
}
