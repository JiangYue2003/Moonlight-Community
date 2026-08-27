# Recoverable Fanout Design

Use `pkg/kafkax.RunConsumer` for `feed-fanout`, with manual commit, finite
`MaxRetries`, and `feed-fanout-dlq`. The existing handler remains the single
fanout implementation, but its decoder validates positive post/creator/time
fields and returns errors for poison input.

Run the reliable consumer through KnowPost's context-supervised listener model.
A temporary `legacy|reliable` mode provides rollback without running two
consumers from the same group in one process.

Whole-event replay intentionally restarts from follower zero. At most ten
batches are repeated; ZADD is idempotent and subsequent trim/expire converges.
A numeric batch checkpoint is rejected because the follower list is not a
durable snapshot and offset reuse could skip recipients after relationship
changes.

The replay command reads a bounded number of DLQ messages, supports key/offset
filters and dry-run, republishes explicitly selected records to the original
topic, and appends replay run/time/source headers. It commits DLQ offsets only
after successful republish.
