package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash"
	"sort"
	"strconv"
	"strings"
)

const mutationPostColumns = "id, creator_id, tag_id, tags, title, description, content_url, content_object_key, content_etag, content_size, content_sha256, is_top, type, visible, img_urls, video_url, status, create_time, update_time, publish_time"

const mutationOutboxColumns = "id, aggregate_type, aggregate_id, type, payload, created_at"

type mutationRowsFingerprint struct {
	ids         []int64
	fingerprint string
}

type mutationMySQLRestoreResult struct {
	PostIDs   []int64
	OutboxIDs []int64
}

func captureMutationMySQLSnapshot(
	ctx context.Context,
	db *sql.DB,
	manifest datasetManifest,
) (mutationMySQLSnapshot, error) {
	if db == nil {
		return mutationMySQLSnapshot{}, fmt.Errorf("mutation checkpoint MySQL connection is required")
	}
	authors := sortedUniqueInt64s(append(append([]int64(nil), manifest.NormalAuthors...), manifest.BigVAuthors...))
	if len(authors) == 0 {
		return mutationMySQLSnapshot{}, fmt.Errorf("mutation checkpoint requires at least one manifest author")
	}
	tx, err := db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return mutationMySQLSnapshot{}, fmt.Errorf("begin mutation checkpoint MySQL transaction: %w", err)
	}
	defer func() { _ = tx.Rollback() }()
	if err := verifyManifestUsers(ctx, tx, manifest, manifestUserIDs(manifest)); err != nil {
		return mutationMySQLSnapshot{}, err
	}

	postQuery := "SELECT " + mutationPostColumns + " FROM know_posts WHERE creator_id IN (" + sqlPlaceholders(len(authors)) + ") ORDER BY id"
	postRows, err := tx.QueryContext(ctx, postQuery, int64Args(authors)...)
	if err != nil {
		return mutationMySQLSnapshot{}, fmt.Errorf("query mutation checkpoint posts: %w", err)
	}
	posts, err := fingerprintMutationRows(postRows, "id")
	closeErr := postRows.Close()
	if err != nil {
		return mutationMySQLSnapshot{}, fmt.Errorf("fingerprint mutation checkpoint posts: %w", err)
	}
	if closeErr != nil {
		return mutationMySQLSnapshot{}, fmt.Errorf("close mutation checkpoint posts: %w", closeErr)
	}

	outbox := mutationRowsFingerprint{fingerprint: emptyMutationRowsFingerprint(strings.Split(mutationOutboxColumns, ", "))}
	if len(posts.ids) > 0 {
		outboxQuery := "SELECT " + mutationOutboxColumns + " FROM outbox WHERE aggregate_type = ? AND aggregate_id IN (" + sqlPlaceholders(len(posts.ids)) + ") ORDER BY id"
		args := append([]any{"knowpost"}, int64Args(posts.ids)...)
		outboxRows, queryErr := tx.QueryContext(ctx, outboxQuery, args...)
		if queryErr != nil {
			return mutationMySQLSnapshot{}, fmt.Errorf("query mutation checkpoint outbox: %w", queryErr)
		}
		outbox, err = fingerprintMutationRows(outboxRows, "id")
		closeErr = outboxRows.Close()
		if err != nil {
			return mutationMySQLSnapshot{}, fmt.Errorf("fingerprint mutation checkpoint outbox: %w", err)
		}
		if closeErr != nil {
			return mutationMySQLSnapshot{}, fmt.Errorf("close mutation checkpoint outbox: %w", closeErr)
		}
	}
	if err := tx.Rollback(); err != nil {
		return mutationMySQLSnapshot{}, fmt.Errorf("finish mutation checkpoint MySQL snapshot: %w", err)
	}
	return mutationMySQLSnapshot{
		AuthorIDs: authors, PostIDs: posts.ids, PostFingerprint: posts.fingerprint,
		OutboxIDs: outbox.ids, OutboxFingerprint: outbox.fingerprint,
	}, nil
}

func restoreMutationMySQLSnapshot(
	ctx context.Context,
	db *sql.DB,
	manifest datasetManifest,
	expected mutationMySQLSnapshot,
) (mutationMySQLRestoreResult, error) {
	if db == nil {
		return mutationMySQLRestoreResult{}, fmt.Errorf("mutation restore MySQL connection is required")
	}
	authors := sortedUniqueInt64s(append(append([]int64(nil), manifest.NormalAuthors...), manifest.BigVAuthors...))
	if !equalInt64s(authors, expected.AuthorIDs) {
		return mutationMySQLRestoreResult{}, fmt.Errorf("mutation checkpoint author set does not match manifest")
	}
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return mutationMySQLRestoreResult{}, fmt.Errorf("begin mutation restore MySQL transaction: %w", err)
	}
	defer func() { _ = tx.Rollback() }()
	if err := verifyManifestUsers(ctx, tx, manifest, manifestUserIDs(manifest)); err != nil {
		return mutationMySQLRestoreResult{}, err
	}

	currentPostIDs, err := queryInt64Column(ctx, tx,
		"SELECT id FROM know_posts WHERE creator_id IN ("+sqlPlaceholders(len(authors))+") ORDER BY id",
		int64Args(authors),
	)
	if err != nil {
		return mutationMySQLRestoreResult{}, fmt.Errorf("query current mutation posts: %w", err)
	}
	mutationPostIDs, err := mutationIDsAfterBaseline(currentPostIDs, expected.PostIDs)
	if err != nil {
		return mutationMySQLRestoreResult{}, err
	}
	if err := verifyMutationMySQLRows(ctx, tx, "know_posts", mutationPostColumns, expected.PostIDs, expected.PostFingerprint); err != nil {
		return mutationMySQLRestoreResult{}, fmt.Errorf("verify mutation checkpoint posts: %w", err)
	}
	if err := verifyMutationMySQLRows(ctx, tx, "outbox", mutationOutboxColumns, expected.OutboxIDs, expected.OutboxFingerprint); err != nil {
		return mutationMySQLRestoreResult{}, fmt.Errorf("verify mutation checkpoint outbox: %w", err)
	}

	mutationOutboxIDs := make([]int64, 0)
	if len(mutationPostIDs) > 0 {
		args := append([]any{"knowpost"}, int64Args(mutationPostIDs)...)
		mutationOutboxIDs, err = queryInt64Column(ctx, tx,
			"SELECT id FROM outbox WHERE aggregate_type = ? AND aggregate_id IN ("+sqlPlaceholders(len(mutationPostIDs))+") ORDER BY id",
			args,
		)
		if err != nil {
			return mutationMySQLRestoreResult{}, fmt.Errorf("query mutation outbox rows: %w", err)
		}
	}
	if _, err := deleteMutationRowsByIDs(ctx, tx, "outbox", mutationOutboxIDs); err != nil {
		return mutationMySQLRestoreResult{}, err
	}
	if _, err := deleteMutationRowsByIDs(ctx, tx, "know_posts", mutationPostIDs); err != nil {
		return mutationMySQLRestoreResult{}, err
	}
	if err := tx.Commit(); err != nil {
		return mutationMySQLRestoreResult{}, fmt.Errorf("commit mutation MySQL restore: %w", err)
	}
	return mutationMySQLRestoreResult{PostIDs: mutationPostIDs, OutboxIDs: mutationOutboxIDs}, nil
}

func verifyMutationMySQLSnapshot(
	ctx context.Context,
	db *sql.DB,
	manifest datasetManifest,
	expected mutationMySQLSnapshot,
) error {
	actual, err := captureMutationMySQLSnapshot(ctx, db, manifest)
	if err != nil {
		return err
	}
	if !equalInt64s(actual.AuthorIDs, expected.AuthorIDs) || !equalInt64s(actual.PostIDs, expected.PostIDs) ||
		!equalInt64s(actual.OutboxIDs, expected.OutboxIDs) || actual.PostFingerprint != expected.PostFingerprint ||
		actual.OutboxFingerprint != expected.OutboxFingerprint {
		return fmt.Errorf("mutation MySQL state does not match checkpoint baseline")
	}
	return nil
}

func verifyMutationMySQLRows(
	ctx context.Context,
	tx *sql.Tx,
	table, columns string,
	ids []int64,
	expectedFingerprint string,
) error {
	if len(ids) == 0 {
		if expectedFingerprint != emptyMutationRowsFingerprint(strings.Split(columns, ", ")) {
			return fmt.Errorf("empty %s baseline fingerprint is invalid", table)
		}
		return nil
	}
	rows, err := tx.QueryContext(ctx,
		"SELECT "+columns+" FROM "+table+" WHERE id IN ("+sqlPlaceholders(len(ids))+") ORDER BY id",
		int64Args(ids)...,
	)
	if err != nil {
		return err
	}
	actual, fingerprintErr := fingerprintMutationRows(rows, "id")
	closeErr := rows.Close()
	if fingerprintErr != nil {
		return fingerprintErr
	}
	if closeErr != nil {
		return closeErr
	}
	if !equalInt64s(actual.ids, ids) {
		return fmt.Errorf("%s baseline row IDs changed", table)
	}
	if actual.fingerprint != expectedFingerprint {
		return fmt.Errorf("%s baseline row fingerprint changed", table)
	}
	return nil
}

func mutationIDsAfterBaseline(current, baseline []int64) ([]int64, error) {
	current = sortedUniqueInt64s(current)
	baseline = sortedUniqueInt64s(baseline)
	currentSet := make(map[int64]struct{}, len(current))
	for _, id := range current {
		currentSet[id] = struct{}{}
	}
	for _, id := range baseline {
		if _, ok := currentSet[id]; !ok {
			return nil, fmt.Errorf("mutation MySQL is missing baseline post %d", id)
		}
	}
	baselineSet := make(map[int64]struct{}, len(baseline))
	for _, id := range baseline {
		baselineSet[id] = struct{}{}
	}
	result := make([]int64, 0, len(current)-len(baseline))
	for _, id := range current {
		if _, ok := baselineSet[id]; !ok {
			result = append(result, id)
		}
	}
	return result, nil
}

func deleteMutationRowsByIDs(ctx context.Context, tx *sql.Tx, table string, ids []int64) (int64, error) {
	var total int64
	for start := 0; start < len(ids); start += redisDeleteBatchSize {
		end := minInt(start+redisDeleteBatchSize, len(ids))
		deleted, err := deleteRowsByIDs(ctx, tx, table, ids[start:end])
		if err != nil {
			return total, err
		}
		total += deleted
	}
	return total, nil
}

func equalInt64s(left, right []int64) bool {
	left = sortedUniqueInt64s(left)
	right = sortedUniqueInt64s(right)
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}

func fingerprintMutationRows(rows *sql.Rows, idColumn string) (mutationRowsFingerprint, error) {
	columns, err := rows.Columns()
	if err != nil {
		return mutationRowsFingerprint{}, err
	}
	idIndex := -1
	for index, column := range columns {
		if column == idColumn {
			idIndex = index
			break
		}
	}
	if idIndex < 0 {
		return mutationRowsFingerprint{}, fmt.Errorf("rows do not contain id column %q", idColumn)
	}
	digest := sha256.New()
	writeMutationFingerprintStrings(digest, columns)
	ids := make([]int64, 0)
	for rows.Next() {
		values := make([]sql.RawBytes, len(columns))
		destinations := make([]any, len(columns))
		for index := range values {
			destinations[index] = &values[index]
		}
		if err := rows.Scan(destinations...); err != nil {
			return mutationRowsFingerprint{}, err
		}
		for _, value := range values {
			writeMutationFingerprintBytes(digest, value)
		}
		id, err := strconv.ParseInt(string(values[idIndex]), 10, 64)
		if err != nil {
			return mutationRowsFingerprint{}, fmt.Errorf("parse row id %q: %w", values[idIndex], err)
		}
		if id <= 0 {
			return mutationRowsFingerprint{}, fmt.Errorf("row id %d must be positive", id)
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return mutationRowsFingerprint{}, err
	}
	if !sort.SliceIsSorted(ids, func(i, j int) bool { return ids[i] < ids[j] }) {
		return mutationRowsFingerprint{}, fmt.Errorf("rows are not ordered by id")
	}
	for index := 1; index < len(ids); index++ {
		if ids[index] == ids[index-1] {
			return mutationRowsFingerprint{}, fmt.Errorf("rows contain duplicate id %d", ids[index])
		}
	}
	return mutationRowsFingerprint{ids: ids, fingerprint: hex.EncodeToString(digest.Sum(nil))}, nil
}

func emptyMutationRowsFingerprint(columns []string) string {
	digest := sha256.New()
	writeMutationFingerprintStrings(digest, columns)
	return hex.EncodeToString(digest.Sum(nil))
}

func writeMutationFingerprintStrings(digest hash.Hash, values []string) {
	for _, value := range values {
		writeMutationFingerprintBytes(digest, []byte(value))
	}
}

func writeMutationFingerprintBytes(digest hash.Hash, value []byte) {
	var size [8]byte
	if value == nil {
		binary.BigEndian.PutUint64(size[:], ^uint64(0))
		_, _ = digest.Write(size[:])
		return
	}
	binary.BigEndian.PutUint64(size[:], uint64(len(value)))
	_, _ = digest.Write(size[:])
	_, _ = digest.Write(value)
}
