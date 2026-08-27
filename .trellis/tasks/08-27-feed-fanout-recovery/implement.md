# Recoverable Fanout Implementation Plan

1. Add failing consumer/handler tests for malformed input, bounded retries, DLQ
   success/failure, cancellation, duplicate delivery, and partial batch replay.
2. Refactor fanout startup from the legacy `kq` queue wrapper to a context-runner
   backed by `pkg/kafkax`, retaining a rollback mode.
3. Add validated retry/DLQ configuration and operational metrics/log fields.
4. Implement the bounded DLQ replay command using the shared event decoder.
5. Run real Kafka integration tests for failure -> DLQ -> replay -> Inbox and
   verify lag/offset behavior.
6. Run focused repeated tests, app shutdown tests, `go vet`, and Trellis check.

Do not widen the shared `kafkax` behavior unless a shared regression test proves
the change is compatible with its existing consumers.
