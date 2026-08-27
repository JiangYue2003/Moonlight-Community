# Make Feed fanout recoverable

## Goal

Replace silent poison-message acknowledgement and unbounded fanout retry with a
bounded, observable, replayable at-least-once path.

## Requirements

- Use the shared manual-commit Kafka consumer with exponential backoff.
- Configure finite retries and a dedicated Feed fanout DLQ.
- Treat malformed JSON and invalid required fields as handler failures.
- Preserve full-event replay and idempotent Redis writes after partial success.
- Provide an explicit audited command to replay selected DLQ events.
- Keep the existing fanout group/key compatibility during rollout.

## Acceptance Criteria

- [ ] Offsets commit only after successful fanout or successful DLQ handoff.
- [ ] A malformed event and a retry-exhausted dependency failure reach the DLQ.
- [ ] Failure to write the DLQ does not commit the original event.
- [ ] Partial success followed by full replay produces exactly the expected
      logical Inbox membership, ordering, trim, and TTL behavior.
- [ ] Replay preserves original key/value, records audit headers, and uses the
      same production decoder/handler after returning to the source topic.
- [ ] Shutdown cancels retry promptly without advancing the failing offset.

## Dependencies And Scope

- Follows `08-27-feed-outbox-dispatch` so the complete durable chain can be
  tested together.
- No per-batch numeric checkpoint is added: a normal Hybrid author has at most
  1,000 followers, making whole-event retry bounded and safer under relationship
  list changes.
