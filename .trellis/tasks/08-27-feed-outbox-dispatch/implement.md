# Durable Outbox Dispatch Implementation Plan

1. Add failing processor tests for valid publish, malformed relevant payload,
   old payload, deleted/private post, tier failure, writer failure, and replay.
2. Add a time-explicit/idempotent FeedWriter entry point while preserving the
   current wrapper for compatibility.
3. Implement the Canal Outbox Feed processor and `kafkax` consumer config.
4. Add `direct|dual|outbox` configuration, defaults, and cross-validation that
   rejects Tier `enforce` with `direct` derivation or a disabled durable
   consumer.
5. Register the listener under the existing app supervision model and update
   PublishLogic dispatch selection.
6. Add integration tests that commit a publish, inject direct dispatch failure
   and retryable tier-resolution failure, then consume/retry the Outbox event
   and verify projection without high-amplification push or failure-driven
   promotion.
7. Run listener, PublishLogic, FeedWriter, service-context, config, integration,
   `go vet`, and Trellis checks.

Rollback to `direct` before stopping the durable consumer. Do not remove the
transactional Outbox row or alter existing event type names.
