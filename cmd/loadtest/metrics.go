package main

import (
	"context"
	"errors"
	"math"
	"sort"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type latencySummary struct {
	Min  time.Duration `json:"min"`
	Mean time.Duration `json:"mean"`
	P50  time.Duration `json:"p50"`
	P90  time.Duration `json:"p90"`
	P95  time.Duration `json:"p95"`
	P99  time.Duration `json:"p99"`
	Max  time.Duration `json:"max"`
}

type stageResult struct {
	Name       string                 `json:"name"`
	Total      int64                  `json:"total"`
	Success    int64                  `json:"success"`
	Failed     int64                  `json:"failed"`
	Timeouts   int64                  `json:"timeouts"`
	SuccessQPS float64                `json:"success_qps"`
	Latency    latencySummary         `json:"latency"`
	Errors     map[string]int64       `json:"errors,omitempty"`
	Pagination *paginationObservation `json:"pagination,omitempty"`
}

type stageRecorder struct {
	name string

	mu        sync.Mutex
	total     int64
	success   int64
	failed    int64
	timeouts  int64
	latencies []time.Duration
	errors    map[string]int64
}

func newStageRecorder(name string) *stageRecorder {
	return &stageRecorder{name: name, errors: make(map[string]int64)}
}

func (r *stageRecorder) Record(latency time.Duration, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.total++
	if err == nil {
		r.success++
		r.latencies = append(r.latencies, latency)
		return
	}

	r.failed++
	if isTimeout(err) {
		r.timeouts++
	}
	r.errors[errorClass(err)]++
}

func (r *stageRecorder) Result(elapsed time.Duration) stageResult {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := stageResult{
		Name:     r.name,
		Total:    r.total,
		Success:  r.success,
		Failed:   r.failed,
		Timeouts: r.timeouts,
		Latency:  summarizeLatencies(r.latencies),
		Errors:   make(map[string]int64, len(r.errors)),
	}
	if elapsed > 0 {
		result.SuccessQPS = float64(r.success) / elapsed.Seconds()
	}
	for key, count := range r.errors {
		result.Errors[key] = count
	}
	return result
}

func errorClass(err error) string {
	switch {
	case isTimeout(err):
		return "timeout"
	case isCanceled(err):
		return "canceled"
	case strings.Contains(strings.ToLower(err.Error()), "circuit breaker is open"):
		return "circuit_breaker"
	default:
		return "request_error"
	}
}

func isTimeout(err error) bool {
	return errors.Is(err, context.DeadlineExceeded) || status.Code(err) == codes.DeadlineExceeded
}

func isCanceled(err error) bool {
	return errors.Is(err, context.Canceled) || status.Code(err) == codes.Canceled
}

func summarizeLatencies(samples []time.Duration) latencySummary {
	if len(samples) == 0 {
		return latencySummary{}
	}

	sorted := append([]time.Duration(nil), samples...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	var total time.Duration
	for _, sample := range sorted {
		total += sample
	}

	return latencySummary{
		Min:  sorted[0],
		Mean: total / time.Duration(len(sorted)),
		P50:  percentile(sorted, 0.50),
		P90:  percentile(sorted, 0.90),
		P95:  percentile(sorted, 0.95),
		P99:  percentile(sorted, 0.99),
		Max:  sorted[len(sorted)-1],
	}
}

func percentile(sorted []time.Duration, p float64) time.Duration {
	if len(sorted) == 0 {
		return 0
	}
	index := int(math.Ceil(p*float64(len(sorted)))) - 1
	if index < 0 {
		index = 0
	}
	if index >= len(sorted) {
		index = len(sorted) - 1
	}
	return sorted[index]
}
