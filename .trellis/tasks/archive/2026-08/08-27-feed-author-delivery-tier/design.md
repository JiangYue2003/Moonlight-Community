# Author Delivery Tier Design

## Storage

Migration `000007` adds `feed_author_delivery_tier`:

- `author_id BIGINT UNSIGNED PRIMARY KEY`
- `tier TINYINT UNSIGNED NOT NULL` (`1` = BigV; no normal rows required)
- `promotion_reason VARCHAR(32) NOT NULL`
- `observed_followers BIGINT UNSIGNED NOT NULL`
- `promoted_at DATETIME(3) NOT NULL`
- `updated_at DATETIME(3) NOT NULL`

Promotion uses insert-or-monotonic-update semantics. No API can delete or lower
the tier. The model exposes `IsBigV`, `BatchBigVs`, and `PromoteBigV` rather than
generic CRUD that would permit demotion.

`promotion_reason` records evidence, not dependency failure. Expected reasons
include Counter threshold, MySQL fallback threshold, and MySQL backfill
threshold. `observed_followers` is always greater than 1,000 for a promotion.

## Resolver

`AuthorTierResolver` owns Hybrid classification. It depends on the tier model,
Counter adapter, and a narrow MySQL active-follower counter backed by
`follower.rel_status=1`. It exposes single-author publish resolution plus
batched read resolution. The batched MySQL path avoids one query per following.
FeedWriter and FeedReader no longer own independent threshold fallbacks in
enforce mode.

Publish resolution:

1. Return BigV when already promoted.
2. Query Counter for an unpromoted author.
3. When Counter succeeds, atomically promote and return BigV only when the
   count is over 1,000; otherwise return normal without persisting a row.
4. On Counter error or missing required data, query the MySQL active-follower
   count.
5. When MySQL succeeds, atomically promote and return BigV only when the count
   is over 1,000; otherwise return normal without persisting a row.
6. When both sources fail, return a typed retryable `tier-resolution` error and
   choose no delivery route.

Tier-store read or promotion failure is also a retryable tier-resolution error:
the resolver cannot safely assume normal or route on an unpersisted BigV
decision. Dual-source failure does not dispatch push or pull, does not create a
tier row, and is not RouteSnapshot-cacheable.

Read resolution batch-loads persisted promotions, then batch-resolves only the
remaining authors. Counter batch failure falls back to a batched MySQL active-
follower count; missing Counter entries fall back only for those authors. Any
unresolved author fails the RouteSnapshot as a whole, preserving its current
all-or-error cache behavior. Shadow mode records resolved tier, evidence source,
mismatch, and retryable errors, but returns the legacy route.

## Backfill And Rollout

A bounded command scans active `follower` rows grouped by `to_user_id` using a
stable author-ID cursor and promotes counts above 1,000. It supports dry-run and
apply, rate limiting, and repeat execution. Rollout is:

1. Apply additive migration with mode off.
2. Run dry-run and apply backfill; verify counts.
3. Enable shadow and observe mismatches/errors.
4. Deploy the transactional `KnowPostPublished` Outbox consumer in `dual` or
   `outbox` mode and prove post-commit tier-resolution retry.
5. Enable enforce only after that compensation gate passes.

`direct` derivation alone does not satisfy the gate. Rollback disables Tier
enforce before stopping the durable consumer; promotion rows remain unused and
are not deleted.

## Compatibility

Forced Push and Pull bypass Hybrid tier resolution exactly as they do today.
The resolver changes only the Hybrid classification source; Cursor seek,
standard-first-page PageCache, RouteSnapshot lifetime, Combined Pipeline,
Redis key formats, ordering, limits, and TTL behavior remain unchanged.
