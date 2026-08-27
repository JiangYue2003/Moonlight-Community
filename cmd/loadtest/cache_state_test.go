package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type recordingCacheStateRuntime struct {
	actions    []string
	warmResult cacheWarmResult
}

func TestCacheStateFeedClientAlwaysUsesDirectRPCForGatewayPreparation(t *testing.T) {
	direct := &recordingFeedReader{}

	client, err := cacheStateFeedClient("gateway", direct)

	require.NoError(t, err)
	require.Same(t, direct, client)
}

func (r *recordingCacheStateRuntime) BumpSafety(context.Context) error {
	r.actions = append(r.actions, "bump")
	return nil
}

func (r *recordingCacheStateRuntime) Warm(context.Context) (cacheWarmResult, error) {
	r.actions = append(r.actions, "warm")
	result := r.warmResult
	if result.Readers == 0 {
		result.Readers = 1
	}
	return result, nil
}

func (r *recordingCacheStateRuntime) Wait(_ context.Context, duration time.Duration) error {
	r.actions = append(r.actions, duration.String())
	return nil
}

func TestPrepareCacheStateExecutesAuditableStateTransitions(t *testing.T) {
	tests := []struct {
		state string
		want  []string
	}{
		{state: "natural"},
		{state: "cold", want: []string{"bump", "1.1s"}},
		{state: "l1-warm", want: []string{"bump", "1.1s", "warm"}},
		{state: "l2-warm", want: []string{"bump", "1.1s", "warm", "1.1s"}},
		{state: "expire-together", want: []string{"bump", "1.1s", "warm", "5.2s"}},
	}
	for _, testCase := range tests {
		t.Run(testCase.state, func(t *testing.T) {
			runtime := &recordingCacheStateRuntime{}
			_, err := prepareCacheState(context.Background(), testCase.state, runtime)
			require.NoError(t, err)
			require.Equal(t, testCase.want, runtime.actions)
		})
	}
}

func TestPrepareCacheStateRejectsUnknownState(t *testing.T) {
	_, err := prepareCacheState(context.Background(), "half-warm", &recordingCacheStateRuntime{})
	require.ErrorContains(t, err, "cache state")
}

func TestPrepareCacheStateRejectsIndeterminateWarmWindows(t *testing.T) {
	_, err := prepareCacheState(context.Background(), "l1-warm", &recordingCacheStateRuntime{
		warmResult: cacheWarmResult{Readers: 21, Duration: 10 * time.Millisecond},
	})
	require.ErrorContains(t, err, "l1-warm")
	require.ErrorContains(t, err, "21 readers")

	_, err = prepareCacheState(context.Background(), "l1-warm", &recordingCacheStateRuntime{
		warmResult: cacheWarmResult{Readers: 1, Duration: 450 * time.Millisecond},
	})
	require.ErrorContains(t, err, "warm duration")

	_, err = prepareCacheState(context.Background(), "l2-warm", &recordingCacheStateRuntime{
		warmResult: cacheWarmResult{Readers: 1200, Duration: 1800 * time.Millisecond},
	})
	require.ErrorContains(t, err, "L2 Fresh")

	_, err = prepareCacheState(context.Background(), "expire-together", &recordingCacheStateRuntime{
		warmResult: cacheWarmResult{Readers: 20000, Duration: 7 * time.Second},
	})
	require.ErrorContains(t, err, "stale retention")
}

func TestValidateCacheStateMeasurementBoundaryIncludesCollectorDelay(t *testing.T) {
	require.NoError(t, validateCacheStateMeasurementBoundary(cacheStateObservation{
		State: "l1-warm", WarmReaders: 1, WarmDuration: 100 * time.Millisecond,
	}, 100*time.Millisecond))

	err := validateCacheStateMeasurementBoundary(cacheStateObservation{
		State: "l1-warm", WarmReaders: 1, WarmDuration: 300 * time.Millisecond,
	}, 150*time.Millisecond)
	require.ErrorContains(t, err, "measurement boundary")

	err = validateCacheStateMeasurementBoundary(cacheStateObservation{
		State: "l2-warm", WarmReaders: 1200, WarmDuration: 1600 * time.Millisecond,
	}, 200*time.Millisecond)
	require.ErrorContains(t, err, "measurement boundary")

	err = validateCacheStateMeasurementBoundary(cacheStateObservation{
		State: "expire-together", WarmReaders: 20000, WarmDuration: 6*time.Second + 700*time.Millisecond,
	}, 150*time.Millisecond)
	require.ErrorContains(t, err, "stale retention")
}
