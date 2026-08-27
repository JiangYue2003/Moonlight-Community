# Timeline Rebuild Design

## Command And Run State

Add an operator-only Feed rebuild command with mutually exclusive reader,
author, and bulk selectors. `dry-run` is default. Apply requires a non-empty
run ID and confirmation token. A MySQL run table records selector/config
fingerprint, mode, reader cursor, author cursor, status, counters, owner lease,
and last error.

The same run ID may resume only with the same immutable fingerprint. A batch
checkpoint advances after every target key in that batch has been written and
verified.

## Desired Projections

Reader Inbox query combines the reader's own Feed-visible published posts with
posts from current normal-tier followings published at or after the current
relationship's `following.updated_at`. This preserves the live no-history-
backfill behavior for a new or reactivated follow. BigV Outbox query selects one
promoted author's Feed-visible published posts. Both order by `publish_time
DESC, id DESC`, deduplicate by post ID, and apply existing limits.

The repository exposes shared projection functions so live writes and rebuild
use the same key format, limits, and TTL policy. TTL is derived from the newest
included publish time; an empty desired projection deletes the key in apply
mode.

## Atomicity And Verification

A Lua script atomically replaces one ZSET and applies its TTL. Dry-run and
post-write verification read the actual ZSET, normalize numeric IDs/scores, and
compare count plus a deterministic SHA-256 fingerprint. Checkpoints never move
past a failed or unverifiable target.

Bulk scans users/readers and persisted BigV authors by numeric ID. Configured
batch size, inter-batch delay, context cancellation, and an owner lease bound
resource use and allow safe handoff.
