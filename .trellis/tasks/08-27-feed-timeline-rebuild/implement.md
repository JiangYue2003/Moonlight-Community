# Timeline Rebuild Implementation Plan

1. Add failing repository tests for reader Inbox and BigV Outbox source queries,
   including ties, visibility, follows, self posts, tier split, limits, and
   stable ID scans.
2. Add migration `000008` for rebuild run/checkpoint state and synchronize the
   development schema.
3. Extract/reuse atomic projection replacement and deterministic fingerprint
   helpers; test empty, trim, TTL, and idempotent replacement with real Redis.
4. Implement dry-run and targeted apply with strict selectors/confirmation.
5. Implement bulk paging, lease, checkpoint, resume, fingerprint compatibility,
   rate limiting, cancellation, and progress counters.
6. Add failure-injection integration tests that delete/corrupt keys, interrupt
   after a verified batch, resume, and compare final projections.
7. Run focused command/model/Feed tests, real Redis/MySQL integration tests,
   relevant page/Cursor regressions, `go vet`, and Trellis check.

Stop and leave the checkpoint unchanged on any source-query, Redis-write, TTL,
or verification error. Never provide a force-skip option in v1.
