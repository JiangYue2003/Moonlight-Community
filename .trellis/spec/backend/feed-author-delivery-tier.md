# Feed Author Delivery Tier

## Scenario: Monotonic BigV Delivery Routing

### 1. Scope / Trigger

- Applies when hybrid Feed routing classifies an author as push (`normal`) or
  pull (`bigv`) from persisted tier state and follower-count evidence.
- The contract spans MySQL schema/model code, Counter and MySQL evidence,
  FeedReader/FeedWriter routing, rollout configuration, and the backfill CLI.
- It prevents transient dependency failures from causing high-amplification
  push delivery or an unsupported permanent promotion.

### 2. Signatures

```go
type FeedAuthorDeliveryTierModel interface {
	IsBigV(ctx context.Context, authorID int64) (bool, error)
	BatchBigVs(ctx context.Context, authorIDs []int64) (map[int64]struct{}, error)
	PromoteBigV(ctx context.Context, authorID int64, reason string, observedFollowers int64) error
}

type ActiveFollowerCounter interface {
	BatchCountActive(ctx context.Context, authorIDs []int64) (map[int64]int64, error)
}

func (r *AuthorTierResolver) Resolve(ctx context.Context, authorID int64) (TierResolution, error)
func (r *AuthorTierResolver) ResolveBatch(ctx context.Context, authorIDs []int64) (map[int64]TierResolution, error)
```

Database contract:

```sql
feed_author_delivery_tier(
  author_id BIGINT UNSIGNED PRIMARY KEY,
  tier TINYINT UNSIGNED CHECK (tier = 1),
  promotion_reason VARCHAR(32),
  observed_followers BIGINT UNSIGNED CHECK (observed_followers > 1000),
  promoted_at DATETIME(3),
  updated_at DATETIME(3)
)
```

Backfill command:

```text
go run ./services/knowpost/rpc/cmd/feed-tier-backfill \
  -datasource <mysql-dsn> [-after-author-id N] [-batch-size 1..10000] \
  [-max-pages N] [-page-interval DURATION] [-apply]
```

### 3. Contracts

- Tier is monotonic: absence means `normal`; a stored `tier=1` row means
  `bigv`. Runtime code never demotes or rewrites the original promotion
  evidence on duplicate promotion.
- A persisted BigV bypasses Counter and follower-count queries.
- An unpromoted author uses Counter evidence first. Counter failure or a
  missing author result falls back to authoritative MySQL active followers,
  defined as `follower.rel_status=1`.
- Only an authoritative count strictly greater than `1000` may persist a
  BigV promotion. Counts at or below `1000` remain normal.
- Forced `push` and `pull` strategies bypass Tier resolution. Cursor,
  PageCache, RouteSnapshot, and Combined Pipeline behavior is unchanged.
- `Feed.AuthorTier.Mode` accepts `off`, `shadow`, or `enforce` and defaults to
  `off`. Shadow records evidence/fallback/mismatch but preserves legacy
  routing.
- Shipped configuration must remain `off` until the transactional Outbox
  sibling provides verified `dual|outbox` compensation and retry behavior.
- Backfill is bounded and dry-run by default. It scans by ascending author ID
  and commits `NextAfterID` only after every promotion in the page succeeds.

### 4. Validation & Error Matrix

| Condition | Required result |
| --- | --- |
| Persisted BigV row exists | Pull classification; no count dependency call |
| Counter returns `>1000` | Persist BigV, then pull |
| Counter returns `<=1000` | Normal classification; no promotion |
| Counter fails or omits author; MySQL returns `>1000` | Persist BigV, then pull |
| Counter fails or omits author; MySQL returns `<=1000` | Normal classification; no promotion |
| Tier read, both evidence paths, or promotion fails | Retryable `tier-resolution` error |
| Enforce resolution fails | No push/pull dispatch, promotion, or RouteSnapshot cache write |
| Tier resolver missing in enforce mode | Retryable configuration-stage `tier-resolution` error |
| Invalid mode | Configuration validation error |
| Backfill page has any promotion failure | Report failure and do not advance the committed page cursor |

### 5. Good/Base/Bad Cases

- Good: Counter is unavailable, MySQL reports `1001`, the author is promoted
  once, and later requests use persisted pull routing.
- Base: Counter reports `1000`; the author remains normal without a tier row.
- Bad: Counter and MySQL fail and hybrid enforce continues with push. The
  correct result is a retryable error before any dispatch or route caching.

### 6. Tests Required

- Model tests assert missing rows are normal, batch lookup is complete, and
  duplicate promotion is monotonic and preserves original evidence.
- Resolver tests cover persisted, above/equal/below threshold, missing Counter
  entries, dual-source failure, tier read failure, and promotion failure.
- Writer/Reader integration tests assert shared persisted classification,
  forced strategy bypass, shadow non-enforcement, zero dispatch on unresolved
  enforce, and no degraded RouteSnapshot cache entry.
- Configuration tests assert the default is `off`, valid modes parse, and
  unknown modes fail startup validation.
- Backfill tests assert dry-run default, bounds, threshold behavior, resume
  cursor stability, and that a failed page does not advance `NextAfterID`.

### 7. Wrong vs Correct

#### Wrong

```go
if counterErr != nil {
	// Treating unknown as normal can trigger high-amplification push.
	return TierResolution{BigV: false}, nil
}
```

#### Correct

```go
if counterErr != nil {
	counts, mysqlErr := followers.BatchCountActive(ctx, []int64{authorID})
	if mysqlErr != nil {
		return TierResolution{}, newTierResolutionError("evidence", errors.Join(counterErr, mysqlErr))
	}
	// Promote only when counts[authorID] > BIGV_THRESHOLD.
}
```
