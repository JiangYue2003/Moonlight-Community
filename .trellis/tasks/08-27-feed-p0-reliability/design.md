# Feed P0 Reliability Design

## Architecture Boundary

MySQL remains the source of truth for posts, relationships, durable publish
events, author delivery tier, and rebuild progress. Redis Inbox and BigV Outbox
ZSETs remain bounded derived projections. Kafka transports at-least-once events;
duplicates are expected and must be harmless.

The existing read architecture is preserved:

```text
PageCache miss / Cursor
  -> RouteSnapshot
  -> Inbox + BigV Outbox Pipeline
  -> chronological Top-N and dedup
  -> FeedItem hydrate
  -> current relationship and visibility filter
```

## 1. Author Delivery Tier

Add a MySQL `feed_author_delivery_tier` table owned by the KnowPost Feed
domain. Only promotions need durable rows; absence means `normal`, while a row
with tier `bigv` is permanent. Promotion records the reason, observed follower
count when available, and timestamps.

Hybrid resolution is shared by writer and reader:

- Forced Push and Pull strategies retain their current explicit behavior.
- A persisted BigV always uses pull delivery.
- For an unpromoted author, a successful Counter result over 1,000 atomically
  promotes BigV with Counter evidence; a result at or below 1,000 remains
  normal without a tier row.
- Counter failure or a missing required Counter result falls back to an
  authoritative MySQL count of `follower` rows where `rel_status=1`. MySQL
  evidence over 1,000 atomically promotes BigV; a count at or below 1,000
  remains normal.
- If the tier store cannot be read or promoted, or Counter and MySQL both fail,
  resolution returns a typed retryable `tier-resolution` error. No delivery
  route or failure-driven promotion is emitted.
- Read resolution performs tier and fallback count lookups in batches. A
  failed batch does not create or cache a partial RouteSnapshot; successful
  route data retains the existing cache contract.

Rollout uses `off -> shadow -> enforce`, defaulting to `off`. A bounded backfill
promotes existing authors whose authoritative active follower rows exceed the
threshold before `enforce` is enabled. Shadow mode compares persisted and legacy
Counter classifications without changing delivery and records fallback/error
outcomes.

Tier capability is implemented before durable dispatch because the Outbox
consumer consumes the resolver contract. Activation is two-stage: the Tier task
may deploy schema, backfill, and shadow behavior, but `enforce` remains blocked
until the transactional `KnowPostPublished` consumer runs in a compensating
mode (`dual` or `outbox`) and a post-commit failure test proves that a retryable
tier-resolution failure leaves the event uncommitted for later retry.

Promotion does not require demotion migration: older normal-author content
already remains in the reader Inbox, while future content is added to the
author Outbox. The existing short final-consistency window for a newly promoted
author is retained; permanent disappearance caused by later demotion is
eliminated.

## 2. Durable Publish-to-Feed Derivation

Reuse the existing transactional `KnowPostPublished` Outbox row. Add a dedicated
Canal-outbox consumer group that validates the event, loads the committed post
when optional event fields are absent, resolves the durable author tier, and
invokes FeedWriter with the authoritative publish time.

Rollout uses `direct -> dual -> outbox`:

- `direct`: current post-commit best-effort dispatch.
- `dual`: direct dispatch plus durable consumer; duplicate ZADD and fanout
  events are safe.
- `outbox`: publish returns after the database transaction and cache/counter
  work; Feed derivation is exclusively driven by the durable Outbox consumer.

The consumer starts from the earliest available group offset. Old, deleted, or
non-Feed-visible posts are skipped after authoritative loading. A consumer
offset is committed only after the Redis Outbox write or downstream fanout
event publication succeeds.

## 3. Recoverable Fanout

Move Feed fanout consumption onto `pkg/kafkax.RunConsumer` so it uses manual
offset commits, exponential backoff, configurable finite retries, and a Feed
DLQ. Invalid JSON and invalid required fields return errors and therefore reach
the same terminal recovery path as retry-exhausted dependency failures.

Do not add a per-batch progress checkpoint. Hybrid push authors have at most
1,000 followers, so a full retry is bounded to at most ten 100-recipient
batches. Re-reading a changing follower list from a numeric batch offset could
skip recipients; replaying the full event keeps the safer contract, and ZADD,
trim, and expiry operations are idempotent for the same post.

Provide an explicit replay command that copies selected DLQ records back to the
original topic while preserving the original key/value and adding replay audit
headers. Replay never edits Kafka offsets or Redis by hand.

## 4. Timeline Reconciliation and Rebuild

Add an operator command with `dry-run` as the default and explicit `apply`.
Selectors support one reader, one author, and a bulk run. Bulk runs page by
stable numeric IDs, rate-limit batches, and persist run/checkpoint state in
MySQL so another process can resume the same run ID.

Rebuild inputs:

- Inbox: the reader's own Feed-visible published posts plus posts from current
  normal-tier followings, ordered by `(publish_time DESC, post_id DESC)`, capped
  at 1,000.
- BigV Outbox: one promoted author's Feed-visible published posts, capped at
  100.

Each Redis key is replaced atomically from a complete desired ZSET, receives a
TTL derived from the newest included publish time and the existing retention
policy, and is read back for count/order/fingerprint verification before the
checkpoint advances. Empty desired projections deliberately delete drifted
keys in apply mode.

Bulk mode scans reader IDs and promoted author IDs separately. It is never
started automatically by the RPC service or a timer in the first delivery.

## Compatibility And Rollback

- All new runtime behavior is behind conservative flags.
- Database migrations are additive; rollback first returns modes to `off` or
  `direct`, then old binaries can run while new tables remain unused.
- Existing Redis key formats, PageCache keys, page API, Cursor codec, ordering,
  Inbox/Outbox limits, and TTL constants are preserved.
- New event payload fields are optional; consumers can load the post for old
  events.
- Tier `enforce` is an integration rollout gate, not a prerequisite for
  deploying the Tier schema or shadow resolver. Rollback disables `enforce`
  before stopping the durable Outbox consumer.
- No task may revert or reformat unrelated changes in the existing dirty worktree.

## Key Trade-offs

- MySQL active-follower fallback adds database work only when Counter evidence
  is unavailable; batching and dependency metrics bound and expose that cost.
- Failing closed when neither source is available can delay Feed projection or
  a cold read, but prevents both high-amplification push and irreversible
  promotion without evidence. The durable Outbox chain absorbs publish-side
  retries once `enforce` is enabled.
- Tier-store failure can fail a cold Feed read instead of silently returning an
  incomplete result.
- At-least-once delivery may do duplicate work, but it does not lose committed
  Feed mutations.
- Bulk rebuild is operator-triggered and resumable, not a continuously running
  self-healing service.
