package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

type localProcessState struct {
	Topology string               `json:"topology"`
	Services []localProcessRecord `json:"services"`
}

type localProcessRecord struct {
	Name               string `json:"name"`
	PID                int32  `json:"pid"`
	Executable         string `json:"executable"`
	ProcessStartUnixMS int64  `json:"process_start_unix_ms"`
}

func collectLocalProcessStats(statePath string, timestamp time.Time) ([]resourceSample, error) {
	statePath = strings.TrimSpace(statePath)
	if statePath == "" {
		return nil, fmt.Errorf("local process state file is not configured")
	}
	body, err := os.ReadFile(statePath)
	if err != nil {
		return nil, fmt.Errorf("read local process state %s: %w", statePath, err)
	}
	var state localProcessState
	if err := json.Unmarshal(body, &state); err != nil {
		return nil, fmt.Errorf("decode local process state %s: %w", statePath, err)
	}
	if state.Topology != "compose-middleware-local-services" {
		return nil, fmt.Errorf("local process state %s has topology %q", statePath, state.Topology)
	}
	if len(state.Services) == 0 {
		return nil, fmt.Errorf("local process state %s has no services", statePath)
	}

	samples := make([]resourceSample, 0, len(state.Services))
	seen := make(map[string]struct{}, len(state.Services))
	for _, record := range state.Services {
		name := strings.ToLower(strings.TrimSpace(record.Name))
		if name == "" || record.PID <= 0 || strings.TrimSpace(record.Executable) == "" || record.ProcessStartUnixMS <= 0 {
			return nil, fmt.Errorf("local process state %s has incomplete service record %q", statePath, record.Name)
		}
		if _, exists := seen[name]; exists {
			return nil, fmt.Errorf("local process state %s contains duplicate service %q", statePath, name)
		}
		seen[name] = struct{}{}

		current, err := process.NewProcess(record.PID)
		if err != nil {
			return nil, fmt.Errorf("open local service %s PID %d: %w", name, record.PID, err)
		}
		running, err := current.IsRunning()
		if err != nil || !running {
			return nil, fmt.Errorf("local service %s PID %d is not running: %w", name, record.PID, err)
		}
		executable, err := current.Exe()
		if err != nil {
			return nil, fmt.Errorf("read local service %s executable: %w", name, err)
		}
		if !sameExecutablePath(executable, record.Executable) {
			return nil, fmt.Errorf("local service %s PID %d executable changed: got %s, expected %s", name, record.PID, executable, record.Executable)
		}
		createdAt, err := current.CreateTime()
		if err != nil {
			return nil, fmt.Errorf("read local service %s start time: %w", name, err)
		}
		if createdAt != record.ProcessStartUnixMS {
			return nil, fmt.Errorf("local service %s PID %d start time changed: got %d, expected %d", name, record.PID, createdAt, record.ProcessStartUnixMS)
		}
		cpuTimes, err := current.Times()
		if err != nil {
			return nil, fmt.Errorf("read local service %s CPU time: %w", name, err)
		}
		memory, err := current.MemoryInfo()
		if err != nil {
			return nil, fmt.Errorf("read local service %s memory: %w", name, err)
		}
		identity := fmt.Sprintf("%s|%d", filepath.Clean(executable), createdAt)
		samples = append(samples, resourceSample{
			Timestamp: timestamp,
			Source:    "process:" + name,
			Identity:  identity,
			Values: map[string]float64{
				"running":           1,
				"pid":               float64(record.PID),
				"process_start_ms":  float64(createdAt),
				"cpu_seconds_total": cpuTimes.Total(),
				"rss_bytes":         float64(memory.RSS),
			},
		})
	}
	sort.Slice(samples, func(i, j int) bool { return samples[i].Source < samples[j].Source })
	return samples, nil
}

func sameExecutablePath(left, right string) bool {
	leftAbsolute, leftErr := filepath.Abs(left)
	rightAbsolute, rightErr := filepath.Abs(right)
	if leftErr == nil {
		left = leftAbsolute
	}
	if rightErr == nil {
		right = rightAbsolute
	}
	return strings.EqualFold(filepath.Clean(left), filepath.Clean(right))
}

func addProcessCPUPercent(samples []resourceSample) []string {
	bySource := make(map[string][]int)
	for index, sample := range samples {
		if strings.HasPrefix(sample.Source, "process:") {
			bySource[sample.Source] = append(bySource[sample.Source], index)
		}
	}
	missing := make([]string, 0, 2)
	add := func(value string) {
		if !containsString(missing, value) {
			missing = append(missing, value)
		}
	}
	for _, indices := range bySource {
		if len(indices) < 2 {
			add("local_process_boundary")
			continue
		}
		for position := 1; position < len(indices); position++ {
			previous := samples[indices[position-1]]
			current := &samples[indices[position]]
			if previous.Identity == "" || current.Identity == "" || previous.Identity != current.Identity {
				add("local_process_changed")
				continue
			}
			if previous.Values["running"] < 1 || current.Values["running"] < 1 {
				add("local_process_changed")
				continue
			}
			elapsed := current.Timestamp.Sub(previous.Timestamp).Seconds()
			cpuDelta := current.Values["cpu_seconds_total"] - previous.Values["cpu_seconds_total"]
			if elapsed <= 0 || cpuDelta < 0 {
				add("local_process_changed")
				continue
			}
			current.Values["cpu_percent"] = cpuDelta / elapsed * 100
		}
	}
	return missing
}
