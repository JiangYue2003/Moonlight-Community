package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	cacheStateNatural        = "natural"
	cacheStateCold           = "cold"
	cacheStateL1Warm         = "l1-warm"
	cacheStateL2Warm         = "l2-warm"
	cacheStateExpireTogether = "expire-together"

	epochL1SettleDelay = 1100 * time.Millisecond
	l1ExpiryDelay      = 1100 * time.Millisecond
	// The page-cache implementation clamps the effective jittered L2 Fresh TTL
	// to at most five seconds. Keep a margin so a scheduler/timer boundary can
	// never leave an "expire-together" trial accidentally Fresh.
	l2FreshExpiryDelay = 5200 * time.Millisecond
	maxL1WarmReaders   = 20
	maxL1WarmDuration  = 400 * time.Millisecond
	maxL2WarmDuration  = 1700 * time.Millisecond
	// L2 Fresh is clamped to at least 3s and Stale is 10s, so the earliest
	// physical expiry is 13s after an individual warm. Reserve one second for
	// scheduler and clock-boundary variance and fail closed above 12s.
	maxExpireTogetherPreparation = 12 * time.Second
)

type cacheStateRuntime interface {
	BumpSafety(context.Context) error
	Warm(context.Context) (cacheWarmResult, error)
	Wait(context.Context, time.Duration) error
}

type cacheWarmResult struct {
	Readers  int
	Duration time.Duration
}

type cacheStateObservation struct {
	State            string
	WarmReaders      int
	WarmDuration     time.Duration
	BoundaryDuration time.Duration
}

func validateCacheStateMeasurementBoundary(observation cacheStateObservation, boundaryDuration time.Duration) error {
	observation.BoundaryDuration = boundaryDuration
	switch observation.State {
	case cacheStateL1Warm:
		if observation.WarmDuration+boundaryDuration > maxL1WarmDuration {
			return fmt.Errorf(
				"l1-warm measurement boundary is not deterministic: warm %s + boundary %s exceeds %s",
				observation.WarmDuration, boundaryDuration, maxL1WarmDuration,
			)
		}
	case cacheStateL2Warm:
		if observation.WarmDuration+boundaryDuration > maxL2WarmDuration {
			return fmt.Errorf(
				"l2-warm measurement boundary is not deterministic: warm %s + boundary %s exceeds %s",
				observation.WarmDuration, boundaryDuration, maxL2WarmDuration,
			)
		}
	case cacheStateExpireTogether:
		if observation.WarmDuration+l2FreshExpiryDelay+boundaryDuration > maxExpireTogetherPreparation {
			return fmt.Errorf(
				"expire-together measurement boundary exceeds stale retention: warm %s + expiry wait %s + boundary %s exceeds %s",
				observation.WarmDuration, l2FreshExpiryDelay, boundaryDuration, maxExpireTogetherPreparation,
			)
		}
	}
	return nil
}

func normalizeCacheState(value string) (string, error) {
	value = strings.ToLower(strings.TrimSpace(value))
	if value == "" {
		value = cacheStateNatural
	}
	switch value {
	case cacheStateNatural, cacheStateCold, cacheStateL1Warm, cacheStateL2Warm, cacheStateExpireTogether:
		return value, nil
	default:
		return "", fmt.Errorf("cache state %q must be natural, cold, l1-warm, l2-warm, or expire-together", value)
	}
}

func prepareCacheState(ctx context.Context, state string, runtime cacheStateRuntime) (cacheStateObservation, error) {
	state, err := normalizeCacheState(state)
	if err != nil {
		return cacheStateObservation{}, err
	}
	observation := cacheStateObservation{State: state}
	if state == cacheStateNatural {
		return observation, nil
	}
	if runtime == nil {
		return observation, fmt.Errorf("cache state runtime is required for %s", state)
	}
	if err := runtime.BumpSafety(ctx); err != nil {
		return observation, fmt.Errorf("prepare %s cache state: bump safety epoch: %w", state, err)
	}
	if err := runtime.Wait(ctx, epochL1SettleDelay); err != nil {
		return observation, fmt.Errorf("prepare %s cache state: wait for epoch L1: %w", state, err)
	}
	if state == cacheStateCold {
		return observation, nil
	}
	warmResult, err := runtime.Warm(ctx)
	observation.WarmReaders = warmResult.Readers
	observation.WarmDuration = warmResult.Duration
	if err != nil {
		return observation, fmt.Errorf("prepare %s cache state: warm pages: %w", state, err)
	}
	switch state {
	case cacheStateL1Warm:
		if warmResult.Readers > maxL1WarmReaders {
			return observation, fmt.Errorf(
				"prepare l1-warm cache state: %d readers exceed deterministic L1 working set %d",
				warmResult.Readers, maxL1WarmReaders,
			)
		}
		if warmResult.Duration > maxL1WarmDuration {
			return observation, fmt.Errorf(
				"prepare l1-warm cache state: warm duration %s exceeds %s freshness budget",
				warmResult.Duration, maxL1WarmDuration,
			)
		}
		return observation, nil
	case cacheStateL2Warm:
		if warmResult.Duration > maxL2WarmDuration {
			return observation, fmt.Errorf(
				"prepare l2-warm cache state: warm duration %s leaves no deterministic L2 Fresh window (max %s)",
				warmResult.Duration, maxL2WarmDuration,
			)
		}
		return observation, runtime.Wait(ctx, l1ExpiryDelay)
	case cacheStateExpireTogether:
		if warmResult.Duration+l2FreshExpiryDelay > maxExpireTogetherPreparation {
			return observation, fmt.Errorf(
				"prepare expire-together cache state: warm %s + expiry wait %s exceeds %s stale retention budget",
				warmResult.Duration, l2FreshExpiryDelay, maxExpireTogetherPreparation,
			)
		}
		return observation, runtime.Wait(ctx, l2FreshExpiryDelay)
	default:
		return observation, nil
	}
}

type liveCacheStateRuntime struct {
	redisAddr      string
	client         feedReader
	readers        []readerIdentity
	allowed        map[int64]struct{}
	requestTimeout time.Duration
	concurrency    int
}

// cacheStateFeedClient keeps cache-state preparation independent from the
// measured ingress. Gateway trials still warm up and measure through HTTP, but
// the auditable Feed cache state itself is prepared directly at KnowPost RPC.
// This avoids mixing authentication/HTTP latency into an L1/L2 freshness
// boundary and gives RPC and Gateway trials the same backend starting state.
func cacheStateFeedClient(entry string, directRPC feedReader) (feedReader, error) {
	if directRPC == nil {
		return nil, fmt.Errorf("direct RPC Feed client is required for cache-state preparation")
	}
	switch entry {
	case "rpc", "gateway":
		return directRPC, nil
	default:
		return nil, fmt.Errorf("invalid entry %q: expected rpc or gateway", entry)
	}
}

func (r *liveCacheStateRuntime) BumpSafety(ctx context.Context) error {
	client := redis.NewClient(&redis.Options{Addr: r.redisAddr})
	defer client.Close()
	return client.Incr(ctx, "feed:content:safety:epoch").Err()
}

func (r *liveCacheStateRuntime) Warm(ctx context.Context) (cacheWarmResult, error) {
	started := time.Now()
	result := cacheWarmResult{Readers: len(r.readers)}
	if r.client == nil || len(r.readers) == 0 {
		result.Duration = time.Since(started)
		return result, fmt.Errorf("cache warm requires a Feed client and readers")
	}
	concurrency := r.concurrency
	if concurrency <= 0 {
		concurrency = 1
	}
	if concurrency > len(r.readers) {
		concurrency = len(r.readers)
	}
	warmCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	jobs := make(chan readerIdentity)
	var workers sync.WaitGroup
	var errMu sync.Mutex
	var firstErr error
	workers.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer workers.Done()
			for reader := range jobs {
				requestCtx, requestCancel := context.WithTimeout(warmCtx, r.requestTimeout)
				page, err := r.client.GetUserFeed(requestCtx, reader, feedReadRequest{Page: 1, Size: 20})
				requestCancel()
				if err == nil {
					err = validateFeedPage(page, loadSpec{RequireItems: true, AllowedCreators: r.allowed})
				}
				if err != nil {
					errMu.Lock()
					if firstErr == nil {
						firstErr = fmt.Errorf("warm reader %d: %w", reader.UserID, err)
						cancel()
					}
					errMu.Unlock()
				}
			}
		}()
	}
sendReaders:
	for _, reader := range r.readers {
		select {
		case <-warmCtx.Done():
			break sendReaders
		case jobs <- reader:
		}
	}
	close(jobs)
	workers.Wait()
	result.Duration = time.Since(started)
	return result, firstErr
}

func (r *liveCacheStateRuntime) Wait(ctx context.Context, duration time.Duration) error {
	timer := time.NewTimer(duration)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func cacheStateReaders(scenario string, readers []readerIdentity) []readerIdentity {
	if scenario == "hot-read" && len(readers) > 0 {
		return readers[:1]
	}
	return readers
}
