// Package userfeed caches the fully materialized first page of a Hybrid Feed.
package userfeed

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/protobuf/proto"

	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

const (
	minL1EffectiveFreshTTL  = 500 * time.Millisecond
	maxL1EffectiveFreshTTL  = time.Second
	minL2EffectiveFreshTTL  = 3 * time.Second
	maxL2EffectiveFreshTTL  = 5 * time.Second
	maxStaleTTL             = 10 * time.Second
	defaultStaleTTL         = 10 * time.Second
	defaultJitterPercent    = 20
	maxJitterPercent        = 20
	defaultRefreshWorkers   = 32
	defaultRefreshQueue     = 1024
	defaultLoaderTimeout    = 2 * time.Second
	defaultErrorLogInterval = 10 * time.Second
)

type Mode string

const (
	ModeOff  Mode = "off"
	ModeL2   Mode = "l2"
	ModeL1L2 Mode = "l1-l2"
)

type Source string

const (
	SourceBypass  Source = "bypass"
	SourceL1Fresh Source = "l1_fresh"
	SourceL2Fresh Source = "l2_fresh"
	SourceL1Stale Source = "l1_stale"
	SourceL2Stale Source = "l2_stale"
	SourceLoad    Source = "miss"
)

type Loader func(ctx context.Context) (*pb.FeedPage, error)

type Operation string

const (
	OperationL2Get              Operation = "l2_get"
	OperationL2Decode           Operation = "l2_decode"
	OperationL2Encode           Operation = "l2_encode"
	OperationL2Set              Operation = "l2_set"
	OperationSingleflightShared Operation = "singleflight_shared"
	OperationRefreshEnqueue     Operation = "refresh_enqueue"
	OperationRefreshLoad        Operation = "refresh_load"
)

type ObserveFunc func(operation Operation, err error)
type ErrorReporter func(operation Operation, err error)

type RefreshState struct {
	QueueDepth    int
	ActiveWorkers int
	PendingKeys   int
}

type Cache interface {
	GetOrLoad(ctx context.Context, req Request, loader Loader) (*pb.FeedPage, Source, error)
	Close() error
}

type L1 interface {
	Get(key string) (any, bool)
	SetWithTTL(key string, value any, cost int64, ttl time.Duration) bool
}

type L2 interface {
	Get(ctx context.Context, key string) ([]byte, bool, error)
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
}

type Config struct {
	Mode                Mode
	KeyPrefix           string
	L1FreshTTL          time.Duration
	L2FreshTTL          time.Duration
	StaleTTL            time.Duration
	JitterPercent       int
	RefreshWorkers      int
	RefreshQueue        int
	LoaderTimeout       time.Duration
	Clock               Clock
	RandomFloat64       func() float64
	Observe             ObserveFunc
	ObserveRefreshState func(state RefreshState)
	ReportError         ErrorReporter
	ErrorLogInterval    time.Duration
}

type loadResult struct {
	page         *pb.FeedPage
	source       Source
	needsRefresh bool
}

type loadCall struct {
	done   chan struct{}
	result loadResult
	err    error
}

type pageCache struct {
	cfg      Config
	l1       L1
	l2       L2
	clock    Clock
	ctx      context.Context
	cancel   context.CancelFunc
	errorLog [6]atomic.Int64

	loadMu sync.Mutex
	loads  map[string]*loadCall
	loadWG sync.WaitGroup
	closed bool

	refreshPending      map[string]*refreshPromise
	refreshActive       atomic.Int64
	refreshPendingCount atomic.Int64
	refreshQueue        chan refreshJob
	refreshWG           sync.WaitGroup
}

func New(cfg Config, l1 L1, l2 L2) (Cache, error) {
	if cfg.Mode == "" {
		cfg.Mode = ModeOff
	}
	if cfg.ErrorLogInterval <= 0 {
		cfg.ErrorLogInterval = defaultErrorLogInterval
	}
	if cfg.Clock == nil {
		cfg.Clock = wallClock{}
	}
	if cfg.RandomFloat64 == nil {
		cfg.RandomFloat64 = rand.Float64
	}
	if cfg.StaleTTL == 0 {
		cfg.StaleTTL = defaultStaleTTL
	}
	if cfg.JitterPercent == 0 {
		cfg.JitterPercent = defaultJitterPercent
	}
	if cfg.RefreshWorkers == 0 {
		cfg.RefreshWorkers = defaultRefreshWorkers
	}
	if cfg.RefreshQueue == 0 {
		cfg.RefreshQueue = defaultRefreshQueue
	}
	if cfg.LoaderTimeout == 0 {
		cfg.LoaderTimeout = defaultLoaderTimeout
	}

	switch cfg.Mode {
	case ModeOff:
		ctx, cancel := context.WithCancel(context.Background())
		return &pageCache{cfg: cfg, clock: cfg.Clock, ctx: ctx, cancel: cancel}, nil
	case ModeL2:
		if l2 == nil {
			return nil, errors.New("user feed page L2 is required in l2 mode")
		}
	case ModeL1L2:
		if l1 == nil {
			return nil, errors.New("user feed page L1 is required in l1-l2 mode")
		}
		if l2 == nil {
			return nil, errors.New("user feed page L2 is required in l1-l2 mode")
		}
	default:
		return nil, fmt.Errorf("unsupported user feed page cache mode %q", cfg.Mode)
	}
	if cfg.L2FreshTTL <= 0 {
		return nil, fmt.Errorf("user feed page L2 Fresh TTL %s must be positive", cfg.L2FreshTTL)
	}
	if cfg.Mode == ModeL1L2 && cfg.L1FreshTTL <= 0 {
		return nil, fmt.Errorf("user feed page L1 Fresh TTL %s must be positive", cfg.L1FreshTTL)
	}
	if cfg.StaleTTL <= 0 || cfg.StaleTTL > maxStaleTTL {
		return nil, fmt.Errorf("user feed page Stale TTL %s must be between 0 and %s", cfg.StaleTTL, maxStaleTTL)
	}
	if cfg.JitterPercent < 0 || cfg.JitterPercent > maxJitterPercent {
		return nil, fmt.Errorf("user feed page JitterPercent %d must be between 0 and %d", cfg.JitterPercent, maxJitterPercent)
	}
	if err := validateEffectiveFreshTTL("L2", cfg.L2FreshTTL, cfg.JitterPercent, minL2EffectiveFreshTTL, maxL2EffectiveFreshTTL); err != nil {
		return nil, err
	}
	if cfg.Mode == ModeL1L2 {
		if err := validateEffectiveFreshTTL("L1", cfg.L1FreshTTL, cfg.JitterPercent, minL1EffectiveFreshTTL, maxL1EffectiveFreshTTL); err != nil {
			return nil, err
		}
	}
	if cfg.RefreshWorkers <= 0 {
		return nil, fmt.Errorf("user feed page RefreshWorkers %d must be positive", cfg.RefreshWorkers)
	}
	if cfg.RefreshQueue <= 0 {
		return nil, fmt.Errorf("user feed page RefreshQueue %d must be positive", cfg.RefreshQueue)
	}
	if cfg.LoaderTimeout <= 0 {
		return nil, fmt.Errorf("user feed page LoaderTimeout %s must be positive", cfg.LoaderTimeout)
	}

	ctx, cancel := context.WithCancel(context.Background())
	c := &pageCache{
		cfg:            cfg,
		l1:             l1,
		l2:             l2,
		clock:          cfg.Clock,
		ctx:            ctx,
		cancel:         cancel,
		loads:          make(map[string]*loadCall),
		refreshPending: make(map[string]*refreshPromise),
		refreshQueue:   make(chan refreshJob, cfg.RefreshQueue),
	}
	c.startRefreshWorkers()
	return c, nil
}

func (c *pageCache) GetOrLoad(
	ctx context.Context,
	req Request,
	loader Loader,
) (*pb.FeedPage, Source, error) {
	if loader == nil {
		return nil, SourceBypass, errors.New("user feed page loader is required")
	}
	if err := ctx.Err(); err != nil {
		return nil, SourceBypass, err
	}
	if c.cfg.Mode == ModeOff || !cacheable(req) {
		page, err := loader(ctx)
		if page == nil && err == nil {
			err = errors.New("user feed page loader returned nil")
		}
		return page, SourceBypass, err
	}

	key := Key(c.cfg.KeyPrefix, req)
	now := c.clock.Now()
	if c.cfg.Mode == ModeL1L2 {
		if value, ok := c.l1.Get(key); ok {
			if record, ok := value.(*cacheRecord); ok && record != nil && record.page != nil {
				source, live := classifyRecord(*record, now, SourceL1Fresh, SourceL1Stale)
				if live {
					if !now.Before(record.refreshAt) {
						c.scheduleRefresh(refreshJob{key: key, loader: loader})
					}
					return record.page, source, nil
				}
			}
		}
	}
	result, err := c.coalescedLoad(ctx, key, func(loadCtx context.Context) (loadResult, error) {
		return c.loadFromL2OrSource(loadCtx, key, loader)
	})
	if err != nil {
		return nil, SourceLoad, err
	}
	if result.needsRefresh {
		c.scheduleRefresh(refreshJob{key: key, loader: loader})
	}
	return result.page, result.source, nil
}

func (c *pageCache) loadFromL2OrSource(ctx context.Context, key string, loader Loader) (loadResult, error) {
	// Capture an already queued/running refresh before the L2 read. The cold
	// load call is registered in c.loads before this function runs, so no new
	// refresh for this key can start after the capture. Live records still
	// return immediately; only an expired/missing record waits for the older
	// refresh, preventing overlapping write-back.
	promise := c.pendingRefresh(key)
	if result, live := c.readL2(ctx, key); live {
		return result, nil
	}
	if promise != nil {
		if c.takeOverQueuedRefresh(key, promise) {
			return c.loadAndStore(ctx, key, loader)
		}
		result, err := waitForRefresh(c.ctx, promise)
		if err == nil {
			return result, nil
		}
		// A failed background refresh is an optimization failure, not a Feed
		// availability failure. Give the original cold path its own bounded
		// attempt after the refresh has completed, so their writes cannot race.
		fallbackCtx, cancel := context.WithTimeout(c.ctx, c.cfg.LoaderTimeout)
		defer cancel()
		return c.loadAndStore(fallbackCtx, key, loader)
	}
	return c.loadAndStore(ctx, key, loader)
}

func (c *pageCache) pendingRefresh(key string) *refreshPromise {
	c.loadMu.Lock()
	defer c.loadMu.Unlock()
	return c.refreshPending[key]
}

func (c *pageCache) readL2(ctx context.Context, key string) (loadResult, bool) {
	raw, hit, l2Err := c.l2.Get(ctx, key)
	c.observe(OperationL2Get, l2Err)
	if l2Err != nil {
		c.reportError(OperationL2Get, l2Err)
	}
	if l2Err == nil && hit {
		record, decodeErr := decodeRecord(raw)
		c.observe(OperationL2Decode, decodeErr)
		if decodeErr != nil {
			c.reportError(OperationL2Decode, decodeErr)
		} else {
			now := c.clock.Now()
			source, live := classifyRecord(record, now, SourceL2Fresh, SourceL2Stale)
			if live {
				c.fillL1(key, record, now)
				return loadResult{
					page:         record.page,
					source:       source,
					needsRefresh: !now.Before(record.refreshAt),
				}, true
			}
		}
	}
	return loadResult{}, false
}

func (c *pageCache) loadAndStore(ctx context.Context, key string, loader Loader) (loadResult, error) {
	page, err := loader(ctx)
	if err != nil {
		return loadResult{}, err
	}
	if page == nil {
		return loadResult{}, errors.New("user feed page loader returned nil")
	}
	now := c.clock.Now()
	freshTTL := c.jitter(c.cfg.L2FreshTTL)
	record := cacheRecord{
		page:       page,
		freshUntil: now.Add(freshTTL),
		refreshAt:  now.Add(freshTTL * 4 / 5),
		staleUntil: now.Add(freshTTL + c.cfg.StaleTTL),
	}
	c.store(ctx, key, record, now)
	return loadResult{page: page, source: SourceLoad}, nil
}

func (c *pageCache) store(ctx context.Context, key string, record cacheRecord, now time.Time) {
	raw, encodeErr := encodeRecord(record)
	c.observe(OperationL2Encode, encodeErr)
	if encodeErr != nil {
		c.reportError(OperationL2Encode, encodeErr)
	}
	if encodeErr == nil {
		ttl := record.staleUntil.Sub(now)
		if ttl > 0 {
			setErr := c.l2.Set(ctx, key, raw, ttl)
			c.observe(OperationL2Set, setErr)
			if setErr != nil {
				c.reportError(OperationL2Set, setErr)
			}
		}
	}
	c.fillL1(key, record, now)
}

func (c *pageCache) fillL1(key string, record cacheRecord, now time.Time) {
	if c.cfg.Mode != ModeL1L2 || c.l1 == nil || record.page == nil {
		return
	}
	remaining := record.staleUntil.Sub(now)
	if remaining <= 0 {
		return
	}
	ttl := c.jitter(c.cfg.L1FreshTTL)
	if ttl > remaining {
		ttl = remaining
	}
	local := record
	cost := int64(2*proto.Size(local.page) + len(key) + 256)
	c.l1.SetWithTTL(key, &local, cost, ttl)
}

func classifyRecord(record cacheRecord, now time.Time, fresh, stale Source) (Source, bool) {
	if now.Before(record.freshUntil) {
		return fresh, true
	}
	if now.Before(record.staleUntil) {
		return stale, true
	}
	return SourceLoad, false
}

func (c *pageCache) jitter(base time.Duration) time.Duration {
	value := c.cfg.RandomFloat64()
	if value < 0 {
		value = 0
	}
	if value > 1 {
		value = 1
	}
	spread := float64(c.cfg.JitterPercent) / 100
	factor := 1 - spread + 2*spread*value
	return time.Duration(float64(base) * factor)
}

func validateEffectiveFreshTTL(
	name string,
	base time.Duration,
	jitterPercent int,
	minimum time.Duration,
	maximum time.Duration,
) error {
	lower := base * time.Duration(100-jitterPercent) / 100
	upper := base * time.Duration(100+jitterPercent) / 100
	if lower < minimum || upper > maximum {
		return fmt.Errorf(
			"user feed page %s effective Fresh TTL [%s,%s] must stay within [%s,%s]",
			name, lower, upper, minimum, maximum,
		)
	}
	return nil
}

func (c *pageCache) coalescedLoad(
	waiterCtx context.Context,
	key string,
	fn func(context.Context) (loadResult, error),
) (loadResult, error) {
	c.loadMu.Lock()
	if c.closed {
		c.loadMu.Unlock()
		return loadResult{}, errors.New("user feed page cache is closed")
	}
	if call := c.loads[key]; call != nil {
		c.loadMu.Unlock()
		c.observe(OperationSingleflightShared, nil)
		return waitForLoad(waiterCtx, call)
	}
	call := &loadCall{done: make(chan struct{})}
	c.loads[key] = call
	c.loadWG.Add(1)
	go c.runLoad(key, call, fn)
	c.loadMu.Unlock()
	return waitForLoad(waiterCtx, call)
}

func waitForLoad(waiterCtx context.Context, call *loadCall) (loadResult, error) {
	select {
	case <-waiterCtx.Done():
		return loadResult{}, waiterCtx.Err()
	case <-call.done:
		return call.result, call.err
	}
}

func waitForRefresh(waiterCtx context.Context, promise *refreshPromise) (loadResult, error) {
	select {
	case <-waiterCtx.Done():
		return loadResult{}, waiterCtx.Err()
	case <-promise.done:
		return promise.result, promise.err
	}
}

func (c *pageCache) runLoad(
	key string,
	call *loadCall,
	fn func(context.Context) (loadResult, error),
) {
	defer c.loadWG.Done()
	defer func() {
		c.loadMu.Lock()
		delete(c.loads, key)
		close(call.done)
		c.loadMu.Unlock()
	}()
	loadCtx, cancel := context.WithTimeout(c.ctx, c.cfg.LoaderTimeout)
	defer cancel()
	call.result, call.err = fn(loadCtx)
}

func (c *pageCache) Close() error {
	c.loadMu.Lock()
	if c.closed {
		c.loadMu.Unlock()
		return nil
	}
	c.closed = true
	c.cancel()
	c.loadMu.Unlock()
	c.refreshWG.Wait()
	for {
		select {
		case job := <-c.refreshQueue:
			c.completeRefresh(job, loadResult{}, context.Canceled)
		default:
			c.observeRefreshState()
			c.loadWG.Wait()
			return nil
		}
	}
}

func (c *pageCache) observe(operation Operation, err error) {
	if c.cfg.Observe == nil {
		return
	}
	defer func() { _ = recover() }()
	c.cfg.Observe(operation, err)
}

func (c *pageCache) observeRefreshState() {
	if c.cfg.ObserveRefreshState == nil || c.refreshQueue == nil {
		return
	}
	defer func() { _ = recover() }()
	c.cfg.ObserveRefreshState(RefreshState{
		QueueDepth:    len(c.refreshQueue),
		ActiveWorkers: int(c.refreshActive.Load()),
		PendingKeys:   int(c.refreshPendingCount.Load()),
	})
}

func (c *pageCache) reportError(operation Operation, err error) {
	if err == nil || c.cfg.ReportError == nil {
		return
	}
	index := operationIndex(operation)
	now := time.Now().UnixNano()
	for {
		last := c.errorLog[index].Load()
		if last != 0 && time.Duration(now-last) < c.cfg.ErrorLogInterval {
			return
		}
		if c.errorLog[index].CompareAndSwap(last, now) {
			defer func() { _ = recover() }()
			c.cfg.ReportError(operation, err)
			return
		}
	}
}

func operationIndex(operation Operation) int {
	switch operation {
	case OperationL2Get:
		return 0
	case OperationL2Decode:
		return 1
	case OperationL2Encode:
		return 2
	case OperationL2Set:
		return 3
	case OperationRefreshEnqueue:
		return 4
	case OperationRefreshLoad:
		return 5
	default:
		return 0
	}
}
