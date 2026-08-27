package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateSmokeRoutingForEveryStrategy(t *testing.T) {
	tests := []struct {
		strategy    string
		observation smokeRoutingObservation
	}{
		{strategy: "push", observation: smokeRoutingObservation{
			NormalFanout: 100, BigVFanout: 1100, KafkaEvents: 2,
		}},
		{strategy: "pull", observation: smokeRoutingObservation{
			NormalOutbox: true, BigVOutbox: true,
		}},
		{strategy: "hybrid", observation: smokeRoutingObservation{
			NormalFanout: 100, BigVOutbox: true, KafkaEvents: 1,
		}},
	}

	for _, test := range tests {
		t.Run(test.strategy, func(t *testing.T) {
			err := validateSmokeRouting(test.strategy, 100, 1100, test.observation)
			require.NoError(t, err)
		})
	}
}

func TestValidateSmokeRoutingRejectsAWriteToBothPaths(t *testing.T) {
	err := validateSmokeRouting("hybrid", 100, 1100, smokeRoutingObservation{
		NormalFanout: 100,
		NormalOutbox: true,
		BigVOutbox:   true,
		KafkaEvents:  1,
	})

	require.ErrorContains(t, err, "normal author outbox")
}
