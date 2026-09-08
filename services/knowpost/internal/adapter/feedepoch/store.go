// Package feedepoch stores monotonic cache invalidation generations for the
// personal Feed path. Redis is authoritative; the local cache is only a short
// read-through optimization and is never used to fabricate a version after a
// Redis failure.
package feedepoch

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/zhiguang/zhiguang-go/pkg/cachex"
	"golang.org/x/sync/singleflight"
)

const (
	defaultRelationL1TTL  = time.Second
	defaultSafetyL1TTL    = time.Second
	epochCacheCost        = int64(32)
	relationLookupTimeout = 2 * time.Second
	relationLockShards    = 256
)

// Store exposes the invalidation generations needed by route and page caches.
type Store interface {
	Relation(ctx context.Context, userID int64) (uint64, error)
	BumpRelation(ctx context.Context, userID int64) (uint64, error)
	Safety(ctx context.Context) (uint64, error)
	BumpSafety(ctx context.Context) (uint64, error)
	MarkSafetyPending()
	SafetyPending() bool
	FlushPendingSafety(ctx context.Context) error
}

// RedisClient is the narrow Redis surface used by the epoch store.
type RedisClient interface {
	Get(ctx context.Context, key string) *goredis.StringCmd
	Incr(ctx context.Context, key string) *goredis.IntCmd
}

// Config keeps epoch keys test-scoped when desired and gives relation/safety
// versions independent, deliberately short local-cache lifetimes.
type Config struct {
	KeyPrefix     string
	RelationL1TTL time.Duration
	SafetyL1TTL   time.Duration
}

type redisStore struct {
	redis RedisClient
	l1    *cachex.L1
	cfg   Config

	// A shard lock orders a user's cold fill and bump without creating a
	// cross-user global critical section. Singleflight collapses same-key cold
	// misses before they enter Redis.
	relationLocks [relationLockShards]sync.Mutex
	relationReads singleflight.Group
	safetyMu      sync.Mutex
	pending       atomic.Uint64
}

func NewStore(client RedisClient, l1 *cachex.L1, cfg Config) (Store, error) {
	if client == nil {
		return nil, errors.New("feed epoch Redis client is required")
	}
	if l1 == nil {
		return nil, errors.New("feed epoch L1 cache is required")
	}
	if cfg.RelationL1TTL <= 0 {
		cfg.RelationL1TTL = defaultRelationL1TTL
	}
	if cfg.SafetyL1TTL <= 0 {
		cfg.SafetyL1TTL = defaultSafetyL1TTL
	}
	cfg.KeyPrefix = normalizeKeyPrefix(cfg.KeyPrefix)
	return &redisStore{redis: client, l1: l1, cfg: cfg}, nil
}

func (s *redisStore) Relation(ctx context.Context, userID int64) (uint64, error) {
	if err := validateUserID(userID); err != nil {
		return 0, err
	}
	key := relationKey(s.cfg.KeyPrefix, userID)
	if value, ok := cachedEpoch(s.l1, key); ok {
		return value, nil
	}

	resultC := s.relationReads.DoChan(key, func() (any, error) {
		lookupCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), relationLookupTimeout)
		defer cancel()
		lock := s.relationLock(userID)
		lock.Lock()
		defer lock.Unlock()
		if err := lookupCtx.Err(); err != nil {
			return uint64(0), err
		}
		if value, ok := cachedEpoch(s.l1, key); ok {
			return value, nil
		}
		value, err := readEpoch(lookupCtx, s.redis, key)
		if err != nil {
			return uint64(0), fmt.Errorf("read relation epoch for userID %d: %w", userID, err)
		}
		s.l1.SetWithTTL(key, value, epochCacheCost, s.cfg.RelationL1TTL)
		return value, nil
	})
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case result := <-resultC:
		if result.Err != nil {
			return 0, result.Err
		}
		value, ok := result.Val.(uint64)
		if !ok {
			return 0, fmt.Errorf("relation epoch lookup returned %T", result.Val)
		}
		return value, nil
	}
}

func (s *redisStore) BumpRelation(ctx context.Context, userID int64) (uint64, error) {
	if err := validateUserID(userID); err != nil {
		return 0, err
	}
	lock := s.relationLock(userID)
	lock.Lock()
	defer lock.Unlock()
	if err := ctx.Err(); err != nil {
		return 0, err
	}
	key := relationKey(s.cfg.KeyPrefix, userID)
	value, err := incrementEpoch(ctx, s.redis, key)
	if err != nil {
		return 0, fmt.Errorf("bump relation epoch for userID %d: %w", userID, err)
	}
	s.cacheImmediately(key, value, s.cfg.RelationL1TTL)
	return value, nil
}

func (s *redisStore) relationLock(userID int64) *sync.Mutex {
	return &s.relationLocks[uint64(userID)%uint64(len(s.relationLocks))]
}

func (s *redisStore) Safety(ctx context.Context) (uint64, error) {
	key := safetyKey(s.cfg.KeyPrefix)
	if value, ok := cachedEpoch(s.l1, key); ok {
		return value, nil
	}
	s.safetyMu.Lock()
	defer s.safetyMu.Unlock()
	if value, ok := cachedEpoch(s.l1, key); ok {
		return value, nil
	}
	value, err := readEpoch(ctx, s.redis, key)
	if err != nil {
		return 0, fmt.Errorf("read content safety epoch: %w", err)
	}
	s.l1.SetWithTTL(key, value, epochCacheCost, s.cfg.SafetyL1TTL)
	return value, nil
}

func (s *redisStore) BumpSafety(ctx context.Context) (uint64, error) {
	s.safetyMu.Lock()
	defer s.safetyMu.Unlock()
	pendingBefore := s.pending.Load()
	value, err := s.bumpSafetyLocked(ctx)
	if err != nil {
		s.pending.Add(1)
		return 0, err
	}
	s.pending.CompareAndSwap(pendingBefore, 0)
	return value, nil
}

func (s *redisStore) MarkSafetyPending() {
	s.pending.Add(1)
}

func (s *redisStore) SafetyPending() bool {
	return s.pending.Load() > 0
}

func (s *redisStore) FlushPendingSafety(ctx context.Context) error {
	s.safetyMu.Lock()
	defer s.safetyMu.Unlock()
	pendingBefore := s.pending.Load()
	if pendingBefore == 0 {
		return nil
	}
	if _, err := s.bumpSafetyLocked(ctx); err != nil {
		return err
	}
	s.pending.CompareAndSwap(pendingBefore, 0)
	return nil
}

func (s *redisStore) bumpSafetyLocked(ctx context.Context) (uint64, error) {
	key := safetyKey(s.cfg.KeyPrefix)
	value, err := incrementEpoch(ctx, s.redis, key)
	if err != nil {
		return 0, fmt.Errorf("bump content safety epoch: %w", err)
	}
	s.cacheImmediately(key, value, s.cfg.SafetyL1TTL)
	return value, nil
}

func (s *redisStore) cacheImmediately(key string, value uint64, ttl time.Duration) {
	// Ristretto admissions are asynchronous. Bumps are low-frequency invalidation
	// events, so waiting here is preferable to letting the current process serve
	// a stale generation after Bump returns. Ordinary read-through fills do not
	// pay this synchronization cost.
	s.l1.Del(key)
	s.l1.SetWithTTL(key, value, epochCacheCost, ttl)
	s.l1.Wait()
}

func cachedEpoch(l1 *cachex.L1, key string) (uint64, bool) {
	raw, ok := l1.Get(key)
	if !ok {
		return 0, false
	}
	value, ok := raw.(uint64)
	return value, ok
}

func readEpoch(ctx context.Context, client RedisClient, key string) (uint64, error) {
	value, err := client.Get(ctx, key).Uint64()
	if errors.Is(err, goredis.Nil) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return value, nil
}

func incrementEpoch(ctx context.Context, client RedisClient, key string) (uint64, error) {
	value, err := client.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	if value < 0 {
		return 0, fmt.Errorf("Redis returned negative epoch %d", value)
	}
	return uint64(value), nil
}

func validateUserID(userID int64) error {
	if userID <= 0 {
		return fmt.Errorf("userID must be positive: %d", userID)
	}
	return nil
}
