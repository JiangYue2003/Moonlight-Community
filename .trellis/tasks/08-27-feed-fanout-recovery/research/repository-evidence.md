# Fanout Recovery Repository Evidence

- The active fanout worker is built with go-zero `kq`, not the repository's
  shared `pkg/kafkax` consumer.
- It fetches the complete follower list, writes batches of 100 through one
  Redis Pipeline per batch, and sets a seven-day processed marker only after all
  batches succeed.
- A failed batch returns an error and the current wrapper retries forever with a
  fixed one-second delay. Earlier successful batches are replayed, which is safe
  because ZADD/trim/expire converge.
- Malformed JSON logs an error and returns nil, advancing the message.
- `pkg/kafkax.RunConsumer` provides manual commits, exponential backoff up to 30
  seconds, finite retries, optional DLQ writing, and does not commit when DLQ
  writing fails.
- Hybrid push authors are capped at 1,000 followers, so whole-event replay is
  bounded to at most ten production batches. The follower list is not a durable
  snapshot, making a numeric batch offset unsafe under concurrent follow changes.

The reliable consumer and replay command must share one strict FeedEvent decoder.
