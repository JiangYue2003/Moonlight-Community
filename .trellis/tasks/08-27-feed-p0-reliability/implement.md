# Feed P0 Reliability Execution Plan

## Task Order

1. Implement `08-27-feed-author-delivery-tier`, but deploy it only in `off` or
   `shadow`; Tier `enforce` remains blocked.
2. Implement `08-27-feed-outbox-dispatch` against the Tier resolver contract
   and prove the transactional compensation path.
3. Unlock and integration-test Tier `enforce` only after step 2 is healthy.
4. Implement `08-27-feed-fanout-recovery`.
5. Implement `08-27-feed-timeline-rebuild`.
6. Run the parent integration review across all four contracts.

The parent owns no direct production implementation. Each child is started,
implemented, checked, and committed independently. A child dependency must be
green before its dependent child starts.

## Cross-Child Gates

- The tier store and resolver are the only Hybrid author-classification owner.
- Permanent BigV promotion requires a Counter or MySQL active-follower count
  above 1,000. Counter failure alone is never promotion evidence.
- Counter failure falls back to MySQL; dual-source failure returns a retryable
  `tier-resolution` error without dispatch, promotion, or RouteSnapshot cache.
- Tier `enforce` is blocked until durable Outbox compensation is active and
  verified; `direct` derivation alone does not satisfy this gate. The Outbox
  child owns cross-configuration validation because it introduces the durable
  derivation mode and consumer enablement settings.
- Publish time and event identity survive Outbox -> Canal -> Feed fanout.
- DLQ replay uses the same production decoder and handler as live consumption.
- Rebuild uses the same tier constants, Redis projection writer, ordering, and
  visibility predicate as live delivery.
- Existing page/Cursor and performance-evolution changes remain intact.

## Integration Validation

1. Apply migrations to an isolated test database and verify up/down behavior.
2. Start with all new modes off and run existing Feed tests.
3. Run tier shadow tests and backfill verification while keeping `enforce`
   disabled. Cover Counter success, MySQL fallback, and dual-source failure.
4. Inject a failure after publish commit and prove Outbox replay projects the
   post exactly once in logical membership.
5. Enable Tier `enforce` only with the verified durable consumer. Prove a
   `tier-resolution` error does not commit the Outbox offset, emit push fanout,
   persist promotion, or poison RouteSnapshot, and succeeds after dependency
   recovery.
6. Inject partial fanout failure, retry exhaustion, DLQ, and replay.
7. Delete and corrupt selected Inbox/Outbox keys; run dry-run, apply, resume,
   and post-write fingerprint verification.
8. Run page/Cursor correctness suites and relevant integration-tag tests.
9. Run focused `go test`, repeated concurrency tests, `go vet`, and Trellis
   check for each child; run the combined Feed scope again at parent review.

No new long-duration performance WP is part of this plan. Use focused
correctness, integration, and short regression guards only.

## Rollback Points

- Tier: set mode to `off`; additive evidence-backed promotion rows remain
  inert. Never stop the durable consumer while Tier remains in `enforce`.
- Derivation: first disable Tier `enforce`, then return derivation to `direct`
  and stop the dedicated consumer group.
- Fanout: stop the new consumer and restore the previous worker only before
  producing events that depend on new validation fields.
- Rebuild: stop the command; the last verified checkpoint identifies the next
  safe batch. Live writers continue normally.

## Dirty Worktree Protection

The repository already contains extensive uncommitted Feed evolution work,
including files these children must touch. Before every edit and commit:

- record the pre-child diff for overlapping files;
- make surgical changes without reverting existing hunks;
- stage only child-owned files or verified child-owned hunks;
- do not use blanket `git add -A`, reset, checkout, or cleanup commands;
- stop for user direction if a task-specific commit cannot be isolated from
  pre-existing user changes.
