# Feed P0 Repository Evidence

## Confirmed Current Behavior

- Hybrid threshold is 1,000; Inbox is 1,000 entries/7 days; BigV Outbox is
  100 entries/24 hours.
- Publish commits the post and a `KnowPostPublished` Outbox row in one MySQL
  transaction, then calls FeedWriter directly. FeedWriter failure is logged and
  does not change publish success.
- Writer Counter failure uses follower count zero. Reader Counter batch failure
  classifies all followings as normal. Both are unsafe for actual BigV authors.
- Reader always reads the user's Inbox and additionally reads Outboxes for the
  authors classified BigV. Promotion therefore preserves old Inbox content;
  demotion can hide old Outbox-only content.
- Fanout uses 100-recipient Redis pipelines and marks the post processed only
  after all batches. Full retry is idempotent, but malformed JSON currently
  returns success and valid handler errors retry forever.
- `pkg/kafkax.RunConsumer` already supports manual commit, exponential backoff,
  finite retries, and DLQ handoff, but Feed fanout does not use it.
- MySQL posts and following/follower rows are authoritative. Redis Feed ZSETs
  have no production rebuild path.

## Contracts To Preserve

- Chronological `(sort_time DESC, post_id DESC)` order and deduplication.
- Current relationship and visibility filtering before returning Feed items.
- RouteSnapshot and Combined Pipeline cold-path improvements.
- PageCache only for Hybrid `page=1,size=20`; Cursor bypasses full-page cache
  and retains seek semantics.
- Forced Push/Pull strategy tests and default-off evolution flags.

## Performance Evidence Boundary

- Cold Hybrid RPC reached 10,661 QPS/P95 18.112ms after RouteSnapshot and
  Combined Pipeline.
- Cursor page50 reached 7,621.7 QPS/P95 3.36ms versus page-number 592.4 QPS/
  P95 35.66ms on the dedicated deep dataset.
- Correctness work must not claim new capacity without a new formal benchmark.

## User Decisions

- Author delivery tier is monotonic: promote only, never automatic demotion.
- Counter failure falls back to the authoritative MySQL count of active
  `follower` rows. Only Counter or MySQL evidence above 1,000 may permanently
  promote BigV; dual-source failure is a retryable tier-resolution error with
  no dispatch or promotion.
- Tier `enforce` waits for the transactional Outbox compensation path to be
  deployed and verified.
- Timeline recovery supports targeted repair and explicit resumable bulk mode,
  defaults to dry-run, and is not an automatic periodic global job in v1.
- Forced Push/Pull, Cursor, PageCache, RouteSnapshot, and Combined Pipeline
  behavior remains unchanged. No new long-duration performance WP is planned.
