# Harden Feed P0 reliability

## Goal

Eliminate the four known P0 correctness and recovery gaps in the Hybrid
following Feed while preserving its existing chronological ordering, page and
Cursor API compatibility, final-consistency model, and feature-flag rollback
boundaries.

## Background

The current Hybrid Feed routes authors by a live follower-count threshold:
authors with at most 1,000 followers fan out through Kafka to follower Inbox
ZSETs, while authors above 1,000 followers write a BigV Outbox ZSET that is
merged at read time. Redis Inbox and BigV Outbox data are derived indexes;
MySQL posts and relationship data remain authoritative.

The following P0 gaps are confirmed in the current implementation:

- Writer and reader classify authors independently from live Counter data.
  Threshold changes can make previously written Outbox content disappear, and
  Counter failures currently degrade authors to push mode without checking the
  authoritative active-follower rows in MySQL.
- Post state and the generic business Outbox commit atomically, but Feed
  derivation is invoked directly after commit. A FeedWriter or Kafka failure is
  logged without a confirmed durable replay path.
- Fanout retries valid events indefinitely and writes idempotently, but has no
  batch progress checkpoint. Malformed JSON is logged and acknowledged instead
  of entering an explicit dead-letter or repair path.
- Redis Feed indexes have bounded size and TTL but no production recovery path
  that reconciles or rebuilds them from authoritative data after loss or drift.

## Requirements

### R1. Stable author delivery routing

- Writer and reader must consume one authoritative delivery-tier decision.
- Hybrid delivery tier is monotonic: authors start as `normal`, may be promoted
  to `bigv`, and are never automatically demoted because follower counts fall.
- A persisted `bigv` tier is sufficient to select pull delivery. For an
  unpromoted author, a successful Counter result above 1,000 or an authoritative
  MySQL count above 1,000 is required before permanent promotion.
- When Counter lookup fails or omits a required author, resolution must fall
  back to the authoritative count of active `follower` rows in MySQL. A count at
  or below 1,000 keeps the author `normal`; it does not create a tier row.
- When both Counter and MySQL resolution fail, Hybrid `enforce` routing must
  return a retryable `tier-resolution` error. It must neither select
  high-amplification push fanout nor persist a failure-driven promotion.
- Promotion, demotion, and migration behavior must not create a window in which
  committed posts disappear solely because writer and reader chose different
  routes.
- Existing forced Push and forced Pull strategies must retain their explicit
  test semantics; the new tier contract applies to Hybrid routing.

### R2. Durable Feed derivation

- A successfully committed publish must leave durable evidence from which its
  Feed mutation can be retried after process exit, Kafka failure, or transient
  dependency failure.
- Feed event production and consumption must be idempotent and safe under
  duplicate and out-of-order delivery.
- Publishing must retain its current business success contract: failure to
  update a derived Feed index after commit must not roll back an already
  committed post.

### R3. Recoverable fanout processing

- Poison or malformed events must be observable and recoverable instead of
  being silently acknowledged.
- Transient failures must use bounded, operationally visible retry behavior and
  an explicit dead-letter or equivalent terminal recovery path.
- Partial batch completion must be replay-safe and must not corrupt Inbox
  ordering, capacity, or TTL semantics.

### R4. Timeline reconciliation and rebuild

- Operators must be able to detect drift between authoritative post/relation
  data and Redis Inbox/BigV Outbox projections.
- Redis Feed indexes must have an auditable, idempotent recovery mechanism for
  a bounded target scope without hand-editing Redis keys.
- Recovery must preserve the canonical `(sort_time DESC, post_id DESC)` order,
  current visibility rules, Inbox/Outbox capacity, and retention policy.
- Inbox reconstruction for a current normal-author relationship includes only
  posts published since that relationship's latest activation time
  (`following.updated_at`); it does not introduce historical backfill.
- Recovery execution must expose progress, failures, and a safe retry path.
- The first delivery must expose one bounded rebuild core with targeted
  reader/author selectors and an explicit bulk mode. Bulk execution must be
  rate-limited, checkpointed, and resumable; dry-run is the default.
- Automatic periodic global rebuild is not part of the first delivery.

### R5. Compatibility and safety

- Keep the existing Hybrid push/pull data model, chronological ordering,
  RouteSnapshot, Combined Pipeline, standard-first-page PageCache, and Cursor
  seek behavior.
- Do not return posts from unfollowed authors or posts that are deleted,
  private, or otherwise no longer Feed-visible.
- New persistent state and events must be additive and support controlled
  rollout and rollback with mixed old/new data during migration.
- Tier `enforce` mode must remain disabled until the transactional
  `KnowPostPublished` Outbox compensation consumer is deployed, healthy, and
  proven to retry `tier-resolution` failures without republishing the post.
- Do not overwrite or discard the existing uncommitted Feed performance work.

## Acceptance Criteria

- [ ] For an unpromoted Hybrid author, Counter success above 1,000 or MySQL
      fallback success above 1,000 atomically promotes BigV; a result at or
      below 1,000 remains normal without creating a promotion row.
- [ ] When Counter and MySQL both fail, Hybrid enforce writer and reader return
      a retryable `tier-resolution` error, do not dispatch push fanout, do not
      persist a promotion, and do not cache a degraded RouteSnapshot.
- [ ] Crossing the 1,000-follower boundary does not make committed Feed items
      disappear during the supported migration window.
- [ ] Tier `enforce` configuration is not activated until the transactional
      Outbox compensation path passes a post-commit failure-and-retry test.
- [ ] A publish committed immediately before Feed dispatch failure is
      eventually projected after replay without republishing the post.
- [ ] Duplicate Feed events and a retry after partial fanout produce the same
      Inbox/Outbox membership and ordering as one successful execution.
- [ ] Malformed and retry-exhausted events reach an observable terminal recovery
      path and can be deliberately replayed after correction.
- [ ] A bounded reconciliation/rebuild run restores deliberately removed or
      drifted Redis Feed data from authoritative state and verifies the result.
- [ ] Existing page and Cursor correctness tests remain green, including ties,
      deduplication, deletion, visibility changes, and unfollow behavior.
- [ ] New rollout flags or compatibility modes default conservatively and have
      tests for rollback/fallback behavior.
- [ ] Targeted unit and integration tests, relevant `go vet`, and the Trellis
      quality review complete before any production-code commit.

## Out of Scope

- Recommendation/ranking Feed behavior or changing chronological ordering.
- Active-reader Inbox sizing, hot-reader PageCache admission, or cold Timeline
  storage tiers.
- BigV hot-key replication/sharding and a dedicated Timeline KV service.
- Expanding full-page caching beyond Hybrid `page=1,size=20`.
- Replacing Cursor seek pagination or running a new long-duration performance
  WP as part of the correctness fix.

## Delivery Map

1. Persistent monotonic author delivery tier and conservative Hybrid routing,
   implemented with `off`/`shadow` defaults while `enforce` remains blocked.
2. Durable `KnowPostPublished` Outbox-to-Feed derivation, dependent on item 1;
   its verified compensation path unlocks Tier `enforce` activation.
3. Bounded retry, DLQ, replay, and fanout recovery semantics, dependent on the
   durable event path from item 2.
4. Targeted and resumable bulk Timeline reconciliation/rebuild, dependent on
   the authoritative tier contract from item 1.
