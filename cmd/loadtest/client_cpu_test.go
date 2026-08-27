package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCalculateClientCPUPercentNormalizesAcrossLogicalCPUs(t *testing.T) {
	total, normalized, err := calculateClientCPUPercent(10, 14, 2*time.Second, 8)

	require.NoError(t, err)
	require.Equal(t, 200.0, total)
	require.Equal(t, 25.0, normalized)
}

func TestCalculateClientCPUPercentRejectsInvalidInputs(t *testing.T) {
	_, _, err := calculateClientCPUPercent(10, 9, time.Second, 8)
	require.ErrorContains(t, err, "went backwards")

	_, _, err = calculateClientCPUPercent(10, 11, 0, 8)
	require.ErrorContains(t, err, "elapsed")

	_, _, err = calculateClientCPUPercent(10, 11, time.Second, 0)
	require.ErrorContains(t, err, "logical CPU")
}
