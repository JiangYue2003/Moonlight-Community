package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

type clientCPUTimer struct {
	process    *process.Process
	startedCPU float64
	startedAt  time.Time
}

func startClientCPUTimer() (*clientCPUTimer, error) {
	current, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		return nil, fmt.Errorf("open load-generator process: %w", err)
	}
	times, err := current.Times()
	if err != nil {
		return nil, fmt.Errorf("read initial load-generator CPU time: %w", err)
	}
	return &clientCPUTimer{
		process: current, startedCPU: times.Total(), startedAt: time.Now(),
	}, nil
}

func (t *clientCPUTimer) Stop() (resourceSample, error) {
	if t == nil || t.process == nil {
		return resourceSample{}, fmt.Errorf("load-generator CPU timer is not initialized")
	}
	endedAt := time.Now()
	times, err := t.process.Times()
	if err != nil {
		return resourceSample{}, fmt.Errorf("read final load-generator CPU time: %w", err)
	}
	total, normalized, err := calculateClientCPUPercent(
		t.startedCPU, times.Total(), endedAt.Sub(t.startedAt), runtime.NumCPU(),
	)
	if err != nil {
		return resourceSample{}, err
	}
	return resourceSample{
		Timestamp: endedAt,
		Source:    "client:loadtest",
		Identity:  fmt.Sprintf("pid:%d", os.Getpid()),
		Values: map[string]float64{
			"cpu_percent_total":      total,
			"cpu_percent_normalized": normalized,
			"logical_cpus":           float64(runtime.NumCPU()),
		},
	}, nil
}

func calculateClientCPUPercent(startCPU, endCPU float64, elapsed time.Duration, logicalCPUs int) (float64, float64, error) {
	if endCPU < startCPU {
		return 0, 0, fmt.Errorf("load-generator CPU time went backwards: start=%.6f end=%.6f", startCPU, endCPU)
	}
	if elapsed <= 0 {
		return 0, 0, fmt.Errorf("load-generator elapsed duration must be positive: %s", elapsed)
	}
	if logicalCPUs <= 0 {
		return 0, 0, fmt.Errorf("load-generator logical CPU count must be positive: %d", logicalCPUs)
	}
	total := (endCPU - startCPU) / elapsed.Seconds() * 100
	return total, total / float64(logicalCPUs), nil
}
