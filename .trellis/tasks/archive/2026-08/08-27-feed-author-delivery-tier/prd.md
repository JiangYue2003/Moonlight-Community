# Persist Feed author delivery tier

## Goal

Give Hybrid Feed reads and writes one durable, monotonic author-classification
contract so Counter failures and follower-count changes cannot select unsafe or
incompatible routes.

## Background

The current writer maps Counter failure to zero followers, while the reader
maps batch Counter failure to all-normal. MySQL `follower` rows with
`rel_status=1` provide the existing authoritative active-follower count. Redis
Feed projections remain derived state, and a persisted BigV promotion is
permanent.

## Requirements

- Add an authoritative MySQL delivery-tier projection owned by KnowPost Feed.
- Treat missing rows as normal and persisted BigV promotion as permanent.
- Share one resolver contract between FeedWriter and FeedReader.
- For an unpromoted author, promote before choosing pull delivery only when a
  successful Counter result is above 1,000.
- If Counter lookup fails or omits a required author, query the authoritative
  MySQL count of active `follower` rows. Promote only when that count is above
  1,000; a count at or below 1,000 remains normal without a tier row.
- If Counter and MySQL both fail, Hybrid `enforce` returns a retryable
  `tier-resolution` error. Do not emit high-amplification push fanout and do not
  permanently promote the author because of dependency failure.
- Tier-store read or promotion failure must also fail closed as a retryable
  tier-resolution error rather than silently selecting normal.
- Preserve forced Push/Pull semantics and provide `off`, `shadow`, and `enforce`
  rollout modes, defaulting to `off`.
- Backfill existing authors over the threshold from authoritative active
  follower rows before enforce mode is used.
- Keep `enforce` disabled until the transactional `KnowPostPublished` Outbox
  compensation consumer is deployed, healthy, and verified to retry a failed
  tier resolution after publish commit.
- Preserve existing Cursor, PageCache, RouteSnapshot, and Combined Pipeline
  behavior. Only successfully resolved Hybrid routes may enter RouteSnapshot.

## Acceptance Criteria

- [ ] Atomic promotion is idempotent and cannot overwrite BigV with normal.
- [ ] Writer and reader use the same persisted BigV set in enforce mode.
- [ ] Counter success above 1,000 promotes with Counter evidence; Counter
      failure followed by MySQL success above 1,000 promotes with MySQL
      evidence; either source at or below 1,000 leaves no promotion row.
- [ ] For an unpromoted author, Counter and MySQL dual failure makes the Hybrid
      enforce resolver return a retryable `tier-resolution` error without push
      dispatch, pull dispatch, promotion, or RouteSnapshot caching; retry
      succeeds once an evidence source recovers.
- [ ] Authors remain BigV after follower counts fall below the threshold.
- [ ] Shadow mode records mismatches without changing current routing.
- [ ] Backfill is bounded, repeatable, and reports promoted/already-promoted/
      failed counts.
- [ ] Tier configuration defaults to `off`; no shipped runtime configuration
      enables `enforce` before the Outbox child supplies and verifies the
      cross-configuration rollout guard.
- [ ] Mode-off behavior and all existing Feed strategy tests remain unchanged.
- [ ] Forced Push/Pull, Cursor, PageCache, RouteSnapshot, and Combined Pipeline
      regression tests remain unchanged and green.

## Out of Scope

- Automatic demotion, hysteresis, or tier deletion.
- Activity-based reader tiers or dynamic cost-model thresholds.
- Feed event durability, fanout DLQ, and Timeline rebuild implementation; those
  are sibling tasks that consume this contract.

## Dependency And Activation Boundary

The Tier schema, resolver, backfill, and shadow mode are implemented and checked
first so the Outbox child can consume their contract. Production `enforce`
activation is a parent integration gate owned by `08-27-feed-outbox-dispatch`,
which introduces the durable derivation mode and can validate that its consumer
is enabled. This child remains independently deliverable by keeping `enforce`
off until that later gate passes.
