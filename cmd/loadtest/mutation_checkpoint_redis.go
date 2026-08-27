package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

const mutationRedisBatchSize = 128

func captureMutationRedisSnapshot(
	ctx context.Context,
	client *redis.Client,
	manifest datasetManifest,
) (mutationRedisSnapshot, error) {
	if client == nil {
		return mutationRedisSnapshot{}, fmt.Errorf("mutation checkpoint Redis connection is required")
	}
	identity, err := mutationRedisIdentity(ctx, client)
	if err != nil {
		return mutationRedisSnapshot{}, err
	}
	keys, err := captureMutationRedisKeys(ctx, client, mutationFeedKeys(manifest))
	if err != nil {
		return mutationRedisSnapshot{}, err
	}
	return finalizeMutationRedisSnapshot(mutationRedisSnapshot{Identity: identity, Keys: keys})
}

func captureMutationRedisKeys(
	ctx context.Context,
	client *redis.Client,
	keys []string,
) ([]mutationRedisKeySnapshot, error) {
	result := make([]mutationRedisKeySnapshot, 0, len(keys))
	for start := 0; start < len(keys); start += mutationRedisBatchSize {
		end := minInt(start+mutationRedisBatchSize, len(keys))
		pipe := client.Pipeline()
		type commands struct {
			key  string
			dump *redis.StringCmd
			pttl *redis.DurationCmd
		}
		batch := make([]commands, 0, end-start)
		for _, key := range keys[start:end] {
			batch = append(batch, commands{key: key, dump: pipe.Dump(ctx, key), pttl: pipe.PTTL(ctx, key)})
		}
		_, execErr := pipe.Exec(ctx)
		if execErr != nil && !errors.Is(execErr, redis.Nil) {
			return nil, fmt.Errorf("capture mutation Redis keys [%d:%d]: %w", start, end, execErr)
		}
		for _, command := range batch {
			dump, dumpErr := command.dump.Result()
			ttl, ttlErr := command.pttl.Result()
			if ttlErr != nil {
				return nil, fmt.Errorf("capture mutation Redis PTTL for %q: %w", command.key, ttlErr)
			}
			ttlMillis, err := mutationPTTLMillis(ttl)
			if err != nil {
				return nil, fmt.Errorf("capture mutation Redis PTTL for %q: %w", command.key, err)
			}
			switch {
			case errors.Is(dumpErr, redis.Nil):
				if ttlMillis != -2 {
					return nil, fmt.Errorf("mutation Redis key %q disappeared inconsistently: pttl=%d", command.key, ttlMillis)
				}
				result = append(result, mutationRedisKeySnapshot{Key: command.key, PTTLMillis: -2})
			case dumpErr != nil:
				return nil, fmt.Errorf("capture mutation Redis DUMP for %q: %w", command.key, dumpErr)
			default:
				if ttlMillis == -2 {
					return nil, fmt.Errorf("mutation Redis key %q expired during checkpoint capture", command.key)
				}
				result = append(result, mutationRedisKeySnapshot{
					Key: command.key, Exists: true, PTTLMillis: ttlMillis, Dump: []byte(dump),
				})
			}
		}
	}
	return result, nil
}

func restoreMutationRedisSnapshot(
	ctx context.Context,
	client *redis.Client,
	snapshot mutationRedisSnapshot,
) (int64, error) {
	if err := validateStoredMutationRedisSnapshot(snapshot); err != nil {
		return 0, err
	}
	identity, err := mutationRedisIdentity(ctx, client)
	if err != nil {
		return 0, err
	}
	if identity != snapshot.Identity {
		return 0, fmt.Errorf("mutation checkpoint Redis identity changed: %q != %q", identity, snapshot.Identity)
	}
	keys := make([]string, len(snapshot.Keys))
	for index := range snapshot.Keys {
		keys[index] = snapshot.Keys[index].Key
	}
	for start := 0; start < len(keys); start += mutationRedisBatchSize {
		end := minInt(start+mutationRedisBatchSize, len(keys))
		if err := client.Del(ctx, keys[start:end]...).Err(); err != nil {
			return 0, fmt.Errorf("clear mutation Redis keys [%d:%d]: %w", start, end, err)
		}
	}
	for start := 0; start < len(snapshot.Keys); start += mutationRedisBatchSize {
		end := minInt(start+mutationRedisBatchSize, len(snapshot.Keys))
		pipe := client.Pipeline()
		commands := 0
		for _, key := range snapshot.Keys[start:end] {
			if !key.Exists {
				continue
			}
			ttl, err := mutationRestoreTTL(key.PTTLMillis)
			if err != nil {
				return 0, fmt.Errorf("restore mutation Redis key %q: %w", key.Key, err)
			}
			pipe.RestoreReplace(ctx, key.Key, ttl, string(key.Dump))
			commands++
		}
		if commands > 0 {
			if _, err := pipe.Exec(ctx); err != nil {
				return 0, fmt.Errorf("restore mutation Redis keys [%d:%d]: %w", start, end, err)
			}
		}
	}
	if err := verifyMutationRedisSnapshot(ctx, client, snapshot, 5*time.Second); err != nil {
		return 0, err
	}
	epoch, err := client.Incr(ctx, "feed:content:safety:epoch").Result()
	if err != nil {
		return 0, fmt.Errorf("bump Feed safety epoch after mutation restore: %w", err)
	}
	return epoch, nil
}

func verifyMutationRedisSnapshot(
	ctx context.Context,
	client *redis.Client,
	expected mutationRedisSnapshot,
	ttlTolerance time.Duration,
) error {
	return verifyMutationRedisSnapshotWithTTLMode(ctx, client, expected, ttlTolerance, false)
}

func verifyAgedMutationRedisSnapshot(
	ctx context.Context,
	client *redis.Client,
	expected mutationRedisSnapshot,
	ttlTolerance time.Duration,
) error {
	return verifyMutationRedisSnapshotWithTTLMode(ctx, client, expected, ttlTolerance, true)
}

func verifyMutationRedisSnapshotWithTTLMode(
	ctx context.Context,
	client *redis.Client,
	expected mutationRedisSnapshot,
	ttlTolerance time.Duration,
	allowAgedTTL bool,
) error {
	if err := validateStoredMutationRedisSnapshot(expected); err != nil {
		return err
	}
	identity, err := mutationRedisIdentity(ctx, client)
	if err != nil {
		return err
	}
	if identity != expected.Identity {
		return fmt.Errorf("mutation checkpoint Redis identity changed: %q != %q", identity, expected.Identity)
	}
	keys := make([]string, len(expected.Keys))
	for index := range expected.Keys {
		keys[index] = expected.Keys[index].Key
	}
	actualKeys, err := captureMutationRedisKeys(ctx, client, keys)
	if err != nil {
		return err
	}
	actual, err := finalizeMutationRedisSnapshot(mutationRedisSnapshot{Identity: identity, Keys: actualKeys})
	if err != nil {
		return err
	}
	if len(actual.Keys) != len(expected.Keys) {
		return fmt.Errorf("mutation Redis key count changed: %d != %d", len(actual.Keys), len(expected.Keys))
	}
	toleranceMillis := ttlTolerance.Milliseconds()
	for index := range expected.Keys {
		want, got := expected.Keys[index], actual.Keys[index]
		if got.Key != want.Key || got.Exists != want.Exists || got.DumpSHA256 != want.DumpSHA256 {
			return fmt.Errorf("mutation Redis key %q content does not match checkpoint", want.Key)
		}
		if !mutationRedisTTLMatches(want.PTTLMillis, got.PTTLMillis, toleranceMillis, allowAgedTTL) {
			return fmt.Errorf("mutation Redis key %q TTL is outside tolerance: got=%d want=%d tolerance=%d", want.Key, got.PTTLMillis, want.PTTLMillis, toleranceMillis)
		}
	}
	return nil
}

func mutationRedisTTLMatches(want, got, toleranceMillis int64, allowAged bool) bool {
	if want <= 0 {
		return got == want
	}
	if got <= 0 {
		return false
	}
	if allowAged {
		return got <= want+toleranceMillis
	}
	return got <= want+toleranceMillis && want-got <= toleranceMillis
}

func validateStoredMutationRedisSnapshot(snapshot mutationRedisSnapshot) error {
	finalized, err := finalizeMutationRedisSnapshot(snapshot)
	if err != nil {
		return err
	}
	if snapshot.Identity == "" || snapshot.Fingerprint == "" || snapshot.Fingerprint != finalized.Fingerprint {
		return fmt.Errorf("mutation Redis checkpoint fingerprint is invalid")
	}
	if len(snapshot.Keys) != len(finalized.Keys) {
		return fmt.Errorf("mutation Redis checkpoint key count is invalid")
	}
	for index := range snapshot.Keys {
		if snapshot.Keys[index].Key != finalized.Keys[index].Key || snapshot.Keys[index].DumpSHA256 != finalized.Keys[index].DumpSHA256 {
			return fmt.Errorf("mutation Redis checkpoint key %q is not canonical", snapshot.Keys[index].Key)
		}
	}
	return nil
}

func mutationRedisIdentity(ctx context.Context, client *redis.Client) (string, error) {
	info, err := client.Info(ctx).Result()
	if err != nil {
		return "", fmt.Errorf("read mutation checkpoint Redis identity: %w", err)
	}
	identity, err := parseRedisRunID(info)
	if err != nil {
		return "", fmt.Errorf("parse mutation checkpoint Redis identity: %w", err)
	}
	return identity, nil
}

func mutationPTTLMillis(value time.Duration) (int64, error) {
	switch value {
	case -1:
		return -1, nil
	case -2:
		return -2, nil
	default:
		if value <= 0 {
			return 0, fmt.Errorf("invalid Redis PTTL %s", value)
		}
		return value.Milliseconds(), nil
	}
}

func mutationRestoreTTL(milliseconds int64) (time.Duration, error) {
	switch {
	case milliseconds == -1:
		return 0, nil
	case milliseconds > 0:
		return time.Duration(milliseconds) * time.Millisecond, nil
	default:
		return 0, fmt.Errorf("invalid checkpoint PTTL %d", milliseconds)
	}
}

func readRedisInt64OrZero(ctx context.Context, client *redis.Client, key string) (int64, error) {
	raw, err := client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	value, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse Redis integer key %q: %w", key, err)
	}
	return value, nil
}
