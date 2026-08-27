package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type cleanupRelationRow struct {
	ID         int64
	FromUserID int64
	ToUserID   int64
}

type cleanupOutboxRow struct {
	ID   int64
	Type string
}

type cleanupDatabaseRecords struct {
	PostIDs   []int64
	Following []cleanupRelationRow
	Follower  []cleanupRelationRow
	Outbox    []cleanupOutboxRow
}

type cleanupResult struct {
	RunID          string           `json:"run_id"`
	CompletedAt    time.Time        `json:"completed_at"`
	DryRun         bool             `json:"dry_run"`
	DatabaseRows   map[string]int64 `json:"database_rows"`
	RedisKeys      int64            `json:"redis_keys"`
	RedisKeyPlan   int              `json:"redis_keys_planned"`
	KnownPostIDs   int              `json:"known_post_ids"`
	KnownOutboxIDs int              `json:"known_outbox_ids"`
	Notes          []string         `json:"notes"`
}

func cleanupDataset(ctx context.Context, cfg benchmarkConfig, manifest datasetManifest, confirmation string, dryRun bool) (cleanupResult, error) {
	if err := validateCleanupManifest(manifest, confirmation); err != nil {
		return cleanupResult{}, err
	}
	if strings.TrimSpace(cfg.Monitor.MySQLDSN) == "" {
		return cleanupResult{}, fmt.Errorf("cleanup requires Monitor.MySQLDSN")
	}
	if _, err := waitForKafkaZero(ctx, cfg.Monitor, 2*time.Minute); err != nil {
		return cleanupResult{}, fmt.Errorf("cleanup requires drained feed Kafka group: %w", err)
	}

	db, err := sql.Open("mysql", cfg.Monitor.MySQLDSN)
	if err != nil {
		return cleanupResult{}, fmt.Errorf("open cleanup MySQL: %w", err)
	}
	defer db.Close()
	if err := db.PingContext(ctx); err != nil {
		return cleanupResult{}, fmt.Errorf("connect cleanup MySQL: %w", err)
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return cleanupResult{}, fmt.Errorf("begin cleanup transaction: %w", err)
	}
	defer func() { _ = tx.Rollback() }()

	userIDs := manifestUserIDs(manifest)
	if err := verifyManifestUsers(ctx, tx, manifest, userIDs); err != nil {
		return cleanupResult{}, err
	}
	records, err := loadCleanupDatabaseRecords(ctx, tx, userIDs)
	if err != nil {
		return cleanupResult{}, err
	}
	redisKeys := cleanupRedisKeys(manifest, records)
	if dryRun {
		loginLogs, err := countLoginLogs(ctx, tx, manifest, userIDs)
		if err != nil {
			return cleanupResult{}, err
		}
		result := cleanupResult{
			RunID: manifest.RunID, CompletedAt: time.Now(), DryRun: true,
			DatabaseRows: map[string]int64{
				"outbox": int64(len(records.Outbox)), "login_logs": loginLogs,
				"follower": int64(len(records.Follower)), "following": int64(len(records.Following)),
				"know_posts": int64(len(records.PostIDs)), "users": int64(len(userIDs)),
			},
			RedisKeyPlan: len(redisKeys), KnownPostIDs: len(records.PostIDs), KnownOutboxIDs: len(records.Outbox),
			Notes: []string{
				"Dry run only: ownership, SQL selection, Kafka drain, and Redis key derivation were verified; no data was deleted.",
			},
		}
		if _, err := writeCleanupResult(cfg.ReportDir, result); err != nil {
			return cleanupResult{}, err
		}
		return result, nil
	}
	counts := make(map[string]int64)
	if counts["outbox"], err = deleteRowsByIDs(ctx, tx, "outbox", outboxIDs(records.Outbox)); err != nil {
		return cleanupResult{}, err
	}
	if counts["login_logs"], err = deleteLoginLogs(ctx, tx, manifest, userIDs); err != nil {
		return cleanupResult{}, err
	}
	if counts["follower"], err = deleteRelationsForUsers(ctx, tx, "follower", userIDs); err != nil {
		return cleanupResult{}, err
	}
	if counts["following"], err = deleteRelationsForUsers(ctx, tx, "following", userIDs); err != nil {
		return cleanupResult{}, err
	}
	if counts["know_posts"], err = deleteRowsByIDs(ctx, tx, "know_posts", records.PostIDs); err != nil {
		return cleanupResult{}, err
	}
	if counts["users"], err = deleteRowsByIDs(ctx, tx, "users", userIDs); err != nil {
		return cleanupResult{}, err
	}
	if err := tx.Commit(); err != nil {
		return cleanupResult{}, fmt.Errorf("commit cleanup transaction: %w", err)
	}

	redisDeleted, err := deleteCleanupRedisKeys(ctx, cfg.RedisAddr, redisKeys)
	if err != nil {
		return cleanupResult{}, err
	}
	result := cleanupResult{
		RunID: manifest.RunID, CompletedAt: time.Now(), DatabaseRows: counts,
		RedisKeys: redisDeleted, RedisKeyPlan: len(redisKeys), KnownPostIDs: len(records.PostIDs), KnownOutboxIDs: len(records.Outbox),
		Notes: []string{
			"Cleanup is manifest-scoped and preserves benchmark reports and the manifest.",
			"Elasticsearch documents and immutable Kafka history are not removed.",
			"Restart application containers after cleanup to discard process-local L1 entries.",
		},
	}
	if _, err := writeCleanupResult(cfg.ReportDir, result); err != nil {
		return cleanupResult{}, err
	}
	return result, nil
}

func validateCleanupManifest(manifest datasetManifest, confirmation string) error {
	if manifest.RunID == "" || confirmation != manifest.RunID {
		return fmt.Errorf("cleanup confirmation must exactly match manifest run id %q", manifest.RunID)
	}
	if len(manifest.Users) == 0 {
		return fmt.Errorf("cleanup manifest contains no users")
	}
	prefix := "feed-loadtest+" + manifest.RunID + "-"
	for _, user := range manifest.Users {
		if user.ID <= 0 || !strings.HasPrefix(user.Email, prefix) || !strings.HasSuffix(user.Email, "@example.invalid") {
			return fmt.Errorf("manifest user %d email %q is not owned by run %s", user.ID, user.Email, manifest.RunID)
		}
	}
	return nil
}

func manifestUserIDs(manifest datasetManifest) []int64 {
	ids := make([]int64, 0, len(manifest.Users))
	for _, user := range manifest.Users {
		ids = append(ids, user.ID)
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	return ids
}

func verifyManifestUsers(ctx context.Context, tx *sql.Tx, manifest datasetManifest, userIDs []int64) error {
	query := "SELECT id, email FROM users WHERE id IN (" + sqlPlaceholders(len(userIDs)) + ")"
	rows, err := tx.QueryContext(ctx, query, int64Args(userIDs)...)
	if err != nil {
		return fmt.Errorf("verify manifest users: %w", err)
	}
	defer rows.Close()
	found := make(map[int64]string, len(manifest.Users))
	for rows.Next() {
		var id int64
		var email sql.NullString
		if err := rows.Scan(&id, &email); err != nil {
			return fmt.Errorf("scan manifest user: %w", err)
		}
		if !email.Valid {
			return fmt.Errorf("database user %d has no email", id)
		}
		found[id] = email.String
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("iterate manifest users: %w", err)
	}
	return validateVerifiedManifestUsers(manifest, found)
}

func validateVerifiedManifestUsers(manifest datasetManifest, found map[int64]string) error {
	for _, user := range manifest.Users {
		email, ok := found[user.ID]
		if !ok {
			return fmt.Errorf("database is missing manifest user %d", user.ID)
		}
		if email != user.Email {
			return fmt.Errorf("database user %d email %q does not match manifest ownership", user.ID, email)
		}
	}
	if len(found) != len(manifest.Users) {
		return fmt.Errorf("database returned %d manifest users, expected %d", len(found), len(manifest.Users))
	}
	return nil
}

func loadCleanupDatabaseRecords(ctx context.Context, tx *sql.Tx, userIDs []int64) (cleanupDatabaseRecords, error) {
	posts, err := queryInt64Column(ctx, tx,
		"SELECT id FROM know_posts WHERE creator_id IN ("+sqlPlaceholders(len(userIDs))+")",
		int64Args(userIDs),
	)
	if err != nil {
		return cleanupDatabaseRecords{}, fmt.Errorf("load cleanup posts: %w", err)
	}
	following, err := queryCleanupRelations(ctx, tx, "following", userIDs)
	if err != nil {
		return cleanupDatabaseRecords{}, err
	}
	follower, err := queryCleanupRelations(ctx, tx, "follower", userIDs)
	if err != nil {
		return cleanupDatabaseRecords{}, err
	}
	outbox, err := queryCleanupOutbox(ctx, tx, posts, userIDs)
	if err != nil {
		return cleanupDatabaseRecords{}, err
	}
	return cleanupDatabaseRecords{PostIDs: posts, Following: following, Follower: follower, Outbox: outbox}, nil
}

func queryCleanupRelations(ctx context.Context, tx *sql.Tx, table string, userIDs []int64) ([]cleanupRelationRow, error) {
	placeholders := sqlPlaceholders(len(userIDs))
	args := append(int64Args(userIDs), int64Args(userIDs)...)
	query := fmt.Sprintf("SELECT id, from_user_id, to_user_id FROM %s WHERE from_user_id IN (%s) OR to_user_id IN (%s)", table, placeholders, placeholders)
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("load cleanup %s: %w", table, err)
	}
	defer rows.Close()
	result := make([]cleanupRelationRow, 0)
	for rows.Next() {
		var row cleanupRelationRow
		if err := rows.Scan(&row.ID, &row.FromUserID, &row.ToUserID); err != nil {
			return nil, fmt.Errorf("scan cleanup %s: %w", table, err)
		}
		result = append(result, row)
	}
	return result, rows.Err()
}

func queryCleanupOutbox(ctx context.Context, tx *sql.Tx, postIDs, userIDs []int64) ([]cleanupOutboxRow, error) {
	clauses := []string{"(aggregate_type = ? AND aggregate_id IN (" + sqlPlaceholders(len(userIDs)) + "))"}
	args := []any{"following"}
	args = append(args, int64Args(userIDs)...)
	if len(postIDs) > 0 {
		clauses = append(clauses, "(aggregate_type = ? AND aggregate_id IN ("+sqlPlaceholders(len(postIDs))+"))")
		args = append(args, "knowpost")
		args = append(args, int64Args(postIDs)...)
	}
	rows, err := tx.QueryContext(ctx, "SELECT id, type FROM outbox WHERE "+strings.Join(clauses, " OR "), args...)
	if err != nil {
		return nil, fmt.Errorf("load cleanup outbox: %w", err)
	}
	defer rows.Close()
	result := make([]cleanupOutboxRow, 0)
	for rows.Next() {
		var row cleanupOutboxRow
		if err := rows.Scan(&row.ID, &row.Type); err != nil {
			return nil, fmt.Errorf("scan cleanup outbox: %w", err)
		}
		result = append(result, row)
	}
	return result, rows.Err()
}

func queryInt64Column(ctx context.Context, tx *sql.Tx, query string, args []any) ([]int64, error) {
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]int64, 0)
	for rows.Next() {
		var value int64
		if err := rows.Scan(&value); err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, rows.Err()
}

func deleteLoginLogs(ctx context.Context, tx *sql.Tx, manifest datasetManifest, userIDs []int64) (int64, error) {
	emails := make([]any, 0, len(manifest.Users))
	for _, user := range manifest.Users {
		emails = append(emails, user.Email)
	}
	query := "DELETE FROM login_logs WHERE user_id IN (" + sqlPlaceholders(len(userIDs)) + ") OR identifier IN (" + sqlPlaceholders(len(emails)) + ")"
	args := append(int64Args(userIDs), emails...)
	return execCount(ctx, tx, "delete login_logs", query, args...)
}

func countLoginLogs(ctx context.Context, tx *sql.Tx, manifest datasetManifest, userIDs []int64) (int64, error) {
	emails := make([]any, 0, len(manifest.Users))
	for _, user := range manifest.Users {
		emails = append(emails, user.Email)
	}
	query := "SELECT COUNT(*) FROM login_logs WHERE user_id IN (" + sqlPlaceholders(len(userIDs)) + ") OR identifier IN (" + sqlPlaceholders(len(emails)) + ")"
	args := append(int64Args(userIDs), emails...)
	var count int64
	if err := tx.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("count cleanup login_logs: %w", err)
	}
	return count, nil
}

func deleteRelationsForUsers(ctx context.Context, tx *sql.Tx, table string, userIDs []int64) (int64, error) {
	placeholders := sqlPlaceholders(len(userIDs))
	args := append(int64Args(userIDs), int64Args(userIDs)...)
	query := fmt.Sprintf("DELETE FROM %s WHERE from_user_id IN (%s) OR to_user_id IN (%s)", table, placeholders, placeholders)
	return execCount(ctx, tx, "delete "+table, query, args...)
}

func deleteRowsByIDs(ctx context.Context, tx *sql.Tx, table string, ids []int64) (int64, error) {
	if len(ids) == 0 {
		return 0, nil
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s)", table, sqlPlaceholders(len(ids)))
	return execCount(ctx, tx, "delete "+table, query, int64Args(ids)...)
}

func execCount(ctx context.Context, tx *sql.Tx, operation, query string, args ...any) (int64, error) {
	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", operation, err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("%s rows affected: %w", operation, err)
	}
	return count, nil
}

func cleanupRedisKeys(manifest datasetManifest, records cleanupDatabaseRecords) []string {
	keys := make(map[string]struct{})
	for _, key := range feedKeysForManifest(manifest) {
		keys[key] = struct{}{}
	}
	for _, user := range manifest.Users {
		for _, key := range []string{
			fmt.Sprintf("cache:users:id:%d", user.ID),
			"cache:users:email:" + user.Email,
			fmt.Sprintf("uf:flws:%d", user.ID),
			fmt.Sprintf("uf:fans:%d", user.ID),
			fmt.Sprintf("ucnt:%d", user.ID),
			fmt.Sprintf("rl:follow:{%d}", user.ID),
		} {
			keys[key] = struct{}{}
		}
	}
	for _, postID := range records.PostIDs {
		for _, key := range []string{
			fmt.Sprintf("cache:knowPosts:id:%d", postID),
			fmt.Sprintf("feed:fanout:processing:%d", postID),
			fmt.Sprintf("cnt:v1:knowpost:%d", postID),
			fmt.Sprintf("agg:v1:knowpost:%d", postID),
		} {
			keys[key] = struct{}{}
		}
	}
	for _, row := range records.Following {
		keys[fmt.Sprintf("cache:following:id:%d", row.ID)] = struct{}{}
		keys[fmt.Sprintf("cache:following:fromUserId:toUserId:%d:%d", row.FromUserID, row.ToUserID)] = struct{}{}
	}
	for _, row := range records.Follower {
		keys[fmt.Sprintf("cache:follower:id:%d", row.ID)] = struct{}{}
		keys[fmt.Sprintf("cache:follower:toUserId:fromUserId:%d:%d", row.ToUserID, row.FromUserID)] = struct{}{}
	}
	for _, row := range records.Outbox {
		keys[fmt.Sprintf("cache:outbox:id:%d", row.ID)] = struct{}{}
		keys[fmt.Sprintf("dedup:rel:%s:%d", row.Type, row.ID)] = struct{}{}
		keys[fmt.Sprintf("dedup:idx:%s:%d", row.Type, row.ID)] = struct{}{}
	}
	result := make([]string, 0, len(keys))
	for key := range keys {
		result = append(result, key)
	}
	sort.Strings(result)
	return result
}

func deleteCleanupRedisKeys(ctx context.Context, addr string, keys []string) (int64, error) {
	client := redis.NewClient(&redis.Options{Addr: addr})
	defer client.Close()
	if err := client.Ping(ctx).Err(); err != nil {
		return 0, fmt.Errorf("connect cleanup Redis %s: %w", addr, err)
	}
	var total int64
	for start := 0; start < len(keys); start += redisDeleteBatchSize {
		end := minInt(start+redisDeleteBatchSize, len(keys))
		deleted, err := client.Del(ctx, keys[start:end]...).Result()
		if err != nil {
			return total, fmt.Errorf("delete cleanup Redis keys [%d:%d]: %w", start, end, err)
		}
		total += deleted
	}
	return total, nil
}

func outboxIDs(rows []cleanupOutboxRow) []int64 {
	ids := make([]int64, 0, len(rows))
	for _, row := range rows {
		ids = append(ids, row.ID)
	}
	return ids
}

func sqlPlaceholders(count int) string {
	if count <= 0 {
		return "NULL"
	}
	return strings.TrimSuffix(strings.Repeat("?,", count), ",")
}

func int64Args(values []int64) []any {
	args := make([]any, len(values))
	for index, value := range values {
		args[index] = value
	}
	return args
}

func writeCleanupResult(root string, result cleanupResult) (string, error) {
	dir := filepath.Join(root, safePathSegment(result.RunID))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", fmt.Errorf("create cleanup report directory: %w", err)
	}
	body, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshal cleanup report: %w", err)
	}
	name := "cleanup.json"
	if result.DryRun {
		name = "cleanup-preview.json"
	}
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, append(body, '\n'), 0o644); err != nil {
		return "", fmt.Errorf("write cleanup report: %w", err)
	}
	return path, nil
}

func cleanupSummary(result cleanupResult) string {
	parts := make([]string, 0, len(result.DatabaseRows))
	for table, count := range result.DatabaseRows {
		parts = append(parts, table+"="+strconv.FormatInt(count, 10))
	}
	sort.Strings(parts)
	return strings.Join(parts, " ")
}
