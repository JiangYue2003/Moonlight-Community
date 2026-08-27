# Author Delivery Tier Implementation Plan

## Steps

1. Add failing model tests for batch lookup and monotonic idempotent promotion.
2. Add migration `000007` up/down and keep `db/schema.sql` synchronized.
3. Implement the narrow tier model and test it with the repository's SQL test
   pattern.
4. Add a narrow single/batch MySQL active-follower count interface backed by
   `follower.rel_status=1`; reuse existing relation-model semantics and add SQL
   tests for active/inactive rows, threshold equality, and batch completeness.
5. Add failing resolver tests for persisted BigV, Counter above/at/below the
   threshold, Counter failure with MySQL above/at/below the threshold, missing
   Counter entries, dual-source failure, tier-store read/promotion failure,
   concurrent promotion, forced strategies, and no demotion.
6. Implement the single/batch resolver and typed retryable `tier-resolution`
   error. Record evidence-source, fallback, mismatch, and dependency outcomes;
   never dispatch or cache an unresolved result.
7. Add `off|shadow|enforce` configuration validation with default `off`. Keep
   all shipped configurations at `off` or `shadow`; the later Outbox child owns
   the cross-configuration guard because it introduces `dual|outbox` and the
   durable-consumer enablement setting.
8. Inject the resolver into FeedWriter and FeedReader without changing forced
   Push/Pull, PageCache, Cursor, RouteSnapshot, or Combined Pipeline behavior.
9. Add the bounded tier-backfill command and dry-run/apply tests; persist only
   MySQL counts above 1,000 and report evidence-based results.
10. After the Outbox child is available, add the failure-injection integration
    test proving that tier-resolution failure leaves durable work retryable and
    emits no push/promotion before dependency recovery.
11. Run focused Feed/model/config/service-context tests repeatedly, relevant
    integration-tag tests, `go vet`, and Trellis check. Do not run a new
    long-duration performance WP.

## Likely Files

- `db/migrations/000007_feed_author_delivery_tier.{up,down}.sql`
- `db/schema.sql`
- `services/knowpost/shared/model/feed_author_delivery_tier*.go`
- `services/knowpost/rpc/internal/feed/{tier_resolver,writer,reader}.go`
- `services/relation/shared/model/followermodel*.go` or an equivalent narrow
  MySQL adapter that supports batch active-follower counts
- `services/knowpost/rpc/internal/config/config.go`
- `services/knowpost/rpc/internal/svc/servicecontext.go`
- KnowPost local/Docker configuration and focused tests
- a bounded backfill command under the KnowPost/Feed ownership boundary

## Rollback Gate

Do not enable enforce until the backfill is verified and the transactional
Outbox compensation consumer has passed its post-commit retry test. If shadow
shows unexpected missing BigV rows, fallback load, or tier-store errors, leave
mode off and stop. Before disabling the durable consumer, first return Tier mode
to `off` or `shadow`.

## Dirty Worktree Gate

Capture the pre-task diff for every overlapping file, make surgical edits,
stage only task-owned files or isolated hunks, and never reset, checkout, clean,
or reformat unrelated work from the existing dirty worktree.
