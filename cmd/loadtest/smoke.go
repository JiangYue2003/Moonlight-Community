package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type smokeRoutingObservation struct {
	NormalFanout int   `json:"normal_fanout"`
	BigVFanout   int   `json:"bigv_fanout"`
	NormalOutbox bool  `json:"normal_outbox"`
	BigVOutbox   bool  `json:"bigv_outbox"`
	KafkaEvents  int64 `json:"kafka_events"`
}

type smokeEvidence struct {
	RunID        string                  `json:"run_id"`
	Strategy     string                  `json:"strategy"`
	ValidatedAt  time.Time               `json:"validated_at"`
	ReaderID     int64                   `json:"reader_id"`
	NormalAuthor int64                   `json:"normal_author"`
	BigVAuthor   int64                   `json:"bigv_author"`
	NormalPost   int64                   `json:"normal_post"`
	BigVPost     int64                   `json:"bigv_post"`
	KafkaBefore  map[string]float64      `json:"kafka_before"`
	KafkaAfter   map[string]float64      `json:"kafka_after"`
	Routing      smokeRoutingObservation `json:"routing"`
	Checks       map[string]bool         `json:"checks"`
	Complete     bool                    `json:"complete"`
}

func runSmokeDataset(ctx context.Context, cfg benchmarkConfig, clients rpcClients, manifest datasetManifest, strategy string) error {
	if len(manifest.Readers) == 0 || len(manifest.NormalAuthors) == 0 || len(manifest.BigVAuthors) == 0 {
		return fmt.Errorf("manifest lacks readers or authors")
	}
	strategy = strings.ToLower(strings.TrimSpace(strategy))
	if strategy != "push" && strategy != "pull" && strategy != "hybrid" {
		return fmt.Errorf("invalid smoke strategy %q", strategy)
	}

	kafkaBefore, err := waitForKafkaZero(ctx, cfg.Monitor, 30*time.Second)
	if err != nil {
		return fmt.Errorf("wait for Kafka before smoke: %w", err)
	}
	authors := []int64{manifest.NormalAuthors[0], manifest.BigVAuthors[0]}
	posts, err := publishAuthorsOnce(ctx, clients.knowpost, authors, manifest.RunID+"-smoke-"+strategy, cfg.Load.RequestTimeout)
	if err != nil {
		return err
	}

	reader := newRPCFeedClient(clients.knowpost)
	wanted := map[int64][]int64{
		manifest.Readers[0]: {posts[0], posts[1]},
		authors[0]:          {posts[0]},
		authors[1]:          {posts[1]},
	}
	if err := waitForFeedVisibility(ctx, reader, wanted, 30*time.Second); err != nil {
		return fmt.Errorf("smoke feed visibility: %w", err)
	}
	kafkaAfter, err := waitForKafkaZero(ctx, cfg.Monitor, 30*time.Second)
	if err != nil {
		return fmt.Errorf("wait for Kafka after smoke: %w", err)
	}

	observation, err := observeSmokeRouting(ctx, cfg.RedisAddr, manifest, authors, posts)
	if err != nil {
		return err
	}
	observation.KafkaEvents = int64(kafkaAfter["log_end_offset_total"] - kafkaBefore["log_end_offset_total"])
	if err := validateSmokeRouting(
		strategy,
		manifest.FollowerCounts[authors[0]],
		manifest.FollowerCounts[authors[1]],
		observation,
	); err != nil {
		return err
	}

	evidence := smokeEvidence{
		RunID: manifest.RunID, Strategy: strategy, ValidatedAt: time.Now(),
		ReaderID: manifest.Readers[0], NormalAuthor: authors[0], BigVAuthor: authors[1],
		NormalPost: posts[0], BigVPost: posts[1], KafkaBefore: kafkaBefore, KafkaAfter: kafkaAfter,
		Routing: observation,
		Checks: map[string]bool{
			"reader_visibility": true, "normal_author_visibility": true,
			"bigv_author_visibility": true, "no_duplicate_items": true, "kafka_drained": true,
		},
		Complete: true,
	}
	path, err := writeSmokeEvidence(cfg.ReportDir, evidence)
	if err != nil {
		return err
	}
	fmt.Printf("smoke: strategy=%s reader=%d normal_fanout=%d bigv_fanout=%d kafka_events=%d evidence=%s\n",
		strategy, manifest.Readers[0], observation.NormalFanout, observation.BigVFanout, observation.KafkaEvents, path)
	return nil
}

func waitForFeedVisibility(ctx context.Context, reader feedReader, wanted map[int64][]int64, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	var lastErr error
	for time.Now().Before(deadline) {
		allVisible := true
		for userID, postIDs := range wanted {
			page, err := reader.GetUserFeed(ctx, readerIdentity{UserID: userID}, feedReadRequest{Page: 1, Size: 100})
			if err == nil {
				err = validateFeedPage(page, loadSpec{})
			}
			if err != nil {
				lastErr = err
				allVisible = false
				break
			}
			postSet := make(map[string]struct{}, len(postIDs))
			for _, postID := range postIDs {
				postSet[strconv.FormatInt(postID, 10)] = struct{}{}
			}
			if !feedContainsAll(page, postSet) {
				allVisible = false
				break
			}
		}
		if allVisible {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(250 * time.Millisecond):
		}
	}
	return fmt.Errorf("timed out: last_error=%v", lastErr)
}

func waitForKafkaZero(ctx context.Context, cfg monitorConfig, timeout time.Duration) (map[string]float64, error) {
	deadline := time.Now().Add(timeout)
	var last map[string]float64
	var lastErr error
	for time.Now().Before(deadline) {
		values, err := collectKafkaLag(ctx, cfg)
		if err == nil {
			last = values
			if values["lag_total"] == 0 {
				return values, nil
			}
		} else {
			lastErr = err
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(250 * time.Millisecond):
		}
	}
	return nil, fmt.Errorf("Kafka lag did not reach zero: last=%v last_error=%v", last, lastErr)
}

func observeSmokeRouting(
	ctx context.Context,
	redisAddr string,
	manifest datasetManifest,
	authors, posts []int64,
) (smokeRoutingObservation, error) {
	client := redis.NewClient(&redis.Options{Addr: redisAddr})
	defer client.Close()
	if err := client.Ping(ctx).Err(); err != nil {
		return smokeRoutingObservation{}, fmt.Errorf("connect feed redis %s: %w", redisAddr, err)
	}

	normalFanout, err := countPostInManifestInboxes(ctx, client, manifest.Users, authors[0], posts[0])
	if err != nil {
		return smokeRoutingObservation{}, err
	}
	bigVFanout, err := countPostInManifestInboxes(ctx, client, manifest.Users, authors[1], posts[1])
	if err != nil {
		return smokeRoutingObservation{}, err
	}
	normalOutbox, err := postInZSet(ctx, client, fmt.Sprintf("feed:bigv:%d", authors[0]), posts[0])
	if err != nil {
		return smokeRoutingObservation{}, err
	}
	bigVOutbox, err := postInZSet(ctx, client, fmt.Sprintf("feed:bigv:%d", authors[1]), posts[1])
	if err != nil {
		return smokeRoutingObservation{}, err
	}
	return smokeRoutingObservation{
		NormalFanout: normalFanout, BigVFanout: bigVFanout,
		NormalOutbox: normalOutbox, BigVOutbox: bigVOutbox,
	}, nil
}

func countPostInManifestInboxes(
	ctx context.Context,
	client *redis.Client,
	users []benchmarkUser,
	creatorID, postID int64,
) (int, error) {
	pipe := client.Pipeline()
	type scoreCommand struct {
		userID int64
		cmd    *redis.FloatCmd
	}
	commands := make([]scoreCommand, 0, len(users))
	member := strconv.FormatInt(postID, 10)
	for _, user := range users {
		commands = append(commands, scoreCommand{
			userID: user.ID,
			cmd:    pipe.ZScore(ctx, fmt.Sprintf("feed:inbox:%d", user.ID), member),
		})
	}
	_, execErr := pipe.Exec(ctx)
	if execErr != nil && execErr != redis.Nil {
		return 0, fmt.Errorf("query manifest inboxes for post %d: %w", postID, execErr)
	}
	count := 0
	for _, command := range commands {
		_, err := command.cmd.Result()
		switch {
		case err == nil && command.userID != creatorID:
			count++
		case err == nil, err == redis.Nil:
		case err != nil:
			return 0, fmt.Errorf("query inbox user %d post %d: %w", command.userID, postID, err)
		}
	}
	return count, nil
}

func postInZSet(ctx context.Context, client *redis.Client, key string, postID int64) (bool, error) {
	_, err := client.ZScore(ctx, key, strconv.FormatInt(postID, 10)).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("query %s post %d: %w", key, postID, err)
	}
	return true, nil
}

func validateSmokeRouting(strategy string, normalFollowers, bigVFollowers int, got smokeRoutingObservation) error {
	expected := smokeRoutingObservation{}
	switch strategy {
	case "push":
		expected.NormalFanout = normalFollowers
		expected.BigVFanout = bigVFollowers
		expected.KafkaEvents = 2
	case "pull":
		expected.NormalOutbox = true
		expected.BigVOutbox = true
	case "hybrid":
		expected.NormalFanout = normalFollowers
		expected.BigVOutbox = true
		expected.KafkaEvents = 1
	default:
		return fmt.Errorf("invalid smoke strategy %q", strategy)
	}
	if got.NormalFanout != expected.NormalFanout {
		return fmt.Errorf("%s normal fanout=%d, want %d", strategy, got.NormalFanout, expected.NormalFanout)
	}
	if got.BigVFanout != expected.BigVFanout {
		return fmt.Errorf("%s bigv fanout=%d, want %d", strategy, got.BigVFanout, expected.BigVFanout)
	}
	if got.NormalOutbox != expected.NormalOutbox {
		return fmt.Errorf("%s normal author outbox=%t, want %t", strategy, got.NormalOutbox, expected.NormalOutbox)
	}
	if got.BigVOutbox != expected.BigVOutbox {
		return fmt.Errorf("%s bigv author outbox=%t, want %t", strategy, got.BigVOutbox, expected.BigVOutbox)
	}
	if got.KafkaEvents != expected.KafkaEvents {
		return fmt.Errorf("%s Kafka events=%d, want %d", strategy, got.KafkaEvents, expected.KafkaEvents)
	}
	return nil
}

func writeSmokeEvidence(root string, evidence smokeEvidence) (string, error) {
	dir := filepath.Join(root, safePathSegment(evidence.RunID))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", fmt.Errorf("create smoke evidence directory: %w", err)
	}
	path := filepath.Join(dir, "smoke-"+safePathSegment(evidence.Strategy)+".json")
	body, err := json.MarshalIndent(evidence, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshal smoke evidence: %w", err)
	}
	if err := os.WriteFile(path, append(body, '\n'), 0o644); err != nil {
		return "", fmt.Errorf("write smoke evidence: %w", err)
	}
	return path, nil
}
