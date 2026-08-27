package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/shirou/gopsutil/v3/process"
	"github.com/stretchr/testify/require"
)

func TestCollectLocalProcessStatsVerifiesRecordedProcessIdentity(t *testing.T) {
	current, err := process.NewProcess(int32(os.Getpid()))
	require.NoError(t, err)
	executable, err := current.Exe()
	require.NoError(t, err)
	createdAt, err := current.CreateTime()
	require.NoError(t, err)

	state := localProcessState{Topology: "compose-middleware-local-services"}
	state.Services = append(state.Services, localProcessRecord{
		Name: "knowpost", PID: int32(os.Getpid()), Executable: executable, ProcessStartUnixMS: createdAt,
	})
	body, err := json.Marshal(state)
	require.NoError(t, err)
	path := filepath.Join(t.TempDir(), "state.json")
	require.NoError(t, os.WriteFile(path, body, 0o600))

	samples, err := collectLocalProcessStats(path, time.Unix(10, 0))

	require.NoError(t, err)
	require.Len(t, samples, 1)
	require.Equal(t, "process:knowpost", samples[0].Source)
	require.Contains(t, samples[0].Identity, executable)
	require.Equal(t, 1.0, samples[0].Values["running"])
	require.Positive(t, samples[0].Values["rss_bytes"])
}

func TestAddProcessCPUPercentUsesCumulativeCPUAndWallTime(t *testing.T) {
	samples := []resourceSample{
		{Timestamp: time.Unix(10, 0), Source: "process:knowpost", Identity: "same", Values: map[string]float64{"cpu_seconds_total": 100, "running": 1}},
		{Timestamp: time.Unix(12, 0), Source: "process:knowpost", Identity: "same", Values: map[string]float64{"cpu_seconds_total": 106, "running": 1}},
	}

	missing := addProcessCPUPercent(samples)

	require.Empty(t, missing)
	require.Equal(t, 300.0, samples[1].Values["cpu_percent"])
}

func TestAddProcessCPUPercentRejectsIdentityChangeAndMissingBoundary(t *testing.T) {
	samples := []resourceSample{
		{Timestamp: time.Unix(10, 0), Source: "process:knowpost", Identity: "old", Values: map[string]float64{"cpu_seconds_total": 100, "running": 1}},
		{Timestamp: time.Unix(12, 0), Source: "process:knowpost", Identity: "new", Values: map[string]float64{"cpu_seconds_total": 101, "running": 1}},
		{Timestamp: time.Unix(10, 0), Source: "process:relation", Identity: "same", Values: map[string]float64{"cpu_seconds_total": 20, "running": 1}},
	}

	missing := addProcessCPUPercent(samples)

	require.Contains(t, missing, "local_process_changed")
	require.Contains(t, missing, "local_process_boundary")
}
