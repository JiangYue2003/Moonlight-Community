# Durable Outbox Dispatch Design

Add a `KnowPostPublished` processor to KnowPost's supervised listener set. It
parses Canal outbox messages strictly enough that malformed relevant rows return
an error, loads the committed post by ID, checks status and Feed visibility,
resolves the durable author tier, and calls a time-explicit FeedWriter method.

The processor is at-least-once. BigV ZADD is idempotent; normal delivery emits a
post-keyed fanout event whose consumer is also idempotent. A crash before offset
commit can repeat work without changing the final projection.

Configuration adds a dedicated group ID, retry limit, DLQ topic, and derivation
mode. In `dual`, the existing post-commit call remains for low latency and the
durable consumer provides compensation. In target `outbox`, PublishLogic stops
calling FeedWriter directly.

Cross-configuration validation rejects Tier `enforce` while derivation remains
`direct` or the durable consumer is disabled. `dual` or `outbox` is necessary
but becomes operationally eligible only after failure injection proves that a
retryable tier-resolution error is not routed, is not prematurely committed,
and succeeds after the evidence dependency recovers.

The consumer starts at the earliest available offset for a new group. It uses
the database row's `publish_time`, so historical replay cannot reorder old posts
as new. Kafka retention is not treated as permanent recovery; the sibling
Timeline rebuild covers loss beyond retained events.
