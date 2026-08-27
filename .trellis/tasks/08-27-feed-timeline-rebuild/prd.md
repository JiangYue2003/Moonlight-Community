# Rebuild Feed Timeline projections

## Goal

Provide an auditable and resumable way to detect and repair missing or drifted
Redis Inbox/BigV Outbox projections from authoritative MySQL state.

## Requirements

- Support targeted reader, targeted author, and explicit bulk runs.
- Default to dry-run; require an explicit apply mode and run ID for mutation.
- Page bulk work by stable IDs, rate-limit it, persist checkpoints, and resume.
- Rebuild canonical order, deduplication, capacities, visibility, and retention.
- For a current normal-author follow, include only posts whose publish time is
  at or after `following.updated_at`; do not backfill earlier history.
- Atomically replace one complete Redis projection key and verify it before
  advancing the checkpoint.
- Never run automatically from the RPC process or a periodic timer in v1.

## Acceptance Criteria

- [ ] Dry-run reports missing, extra, order-drifted, and matching projections
      without changing Redis.
- [ ] Apply restores deliberately deleted/corrupted Inbox and BigV Outbox keys
      with the expected fingerprint and TTL.
- [ ] An interrupted bulk run resumes after the last verified batch and does
      not skip or repeat logical targets incorrectly.
- [ ] Re-running a completed target is idempotent.
- [ ] Current follows, monotonic author tiers, current post status/visibility,
      self posts, ordering tie-breaks, and key limits are covered by tests.
- [ ] Invalid selectors, reused incompatible run IDs, and concurrent ownership
      fail closed.

## Dependencies And Scope

- Depends on `08-27-feed-author-delivery-tier` and reuses live projection
  constants/writers rather than duplicating them.
- Does not implement automatic periodic repair, cold storage, Feed Head, or
  cross-region orchestration.
