# Durable Dispatch Repository Evidence

- Publish updates `know_posts` and inserts `KnowPostPublished` into `outbox` in
  one transaction through `OutboxModel.InsertInTx`.
- The event payload contains type, post ID, and author. The authoritative
  `know_posts.publish_time`, status, and visibility remain queryable.
- Canal emits outbox table mutations to `canal-outbox`. KnowPost already runs
  independent relation-epoch and content-safety consumer groups through
  `pkg/kafkax` and supervises them with the RPC process.
- Content-safety intentionally skips Published events, so a dedicated Feed
  projection consumer group does not conflict with it.
- Current Feed dispatch runs only after transaction commit and logs errors; no
  confirmed consumer bridges `KnowPostPublished` back into FeedWriter.
- ZADD, post-keyed fanout, merge, and dedup make at-least-once replay logically
  idempotent.

The processor must use the database publish time during historical replay and
must not treat a newly consumed old event as newly published now.

## Planning Decisions

- This child owns the cross-configuration guard that rejects Tier `enforce`
  with `direct` derivation or a disabled durable consumer.
- A retryable tier-resolution error must emit no Feed route, remain recoverable
  through retry/DLQ, and succeed after Counter or MySQL evidence recovers.
- The transactional compensation path must pass failure injection before the
  parent rollout may activate Tier `enforce`.
