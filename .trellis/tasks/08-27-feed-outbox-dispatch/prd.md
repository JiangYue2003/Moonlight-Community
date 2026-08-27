# Derive Feed through durable Outbox

## Goal

Ensure every committed publication leaves replayable evidence that eventually
updates the Feed projection even if the publishing process or Kafka dispatch
fails immediately after the database commit.

## Requirements

- Reuse the transactional `KnowPostPublished` Outbox event and Canal topic.
- Add a dedicated Feed projection consumer group with finite retry and DLQ.
- Resolve delivery through the authoritative tier contract from
  `feed-author-delivery-tier`.
- Use the committed post publish time as the Redis score and validate current
  status/visibility before projection.
- Support `direct`, `dual`, and `outbox` rollout modes, defaulting to `direct`.
- Commit the Canal offset only after downstream Feed dispatch succeeds.
- Own the cross-configuration activation guard for author Tier `enforce`:
  `direct` is invalid with Tier `enforce`; the durable consumer must be enabled
  in `dual` or `outbox` mode.

## Acceptance Criteria

- [ ] A post committed before direct dispatch failure is projected by the
      durable consumer without republishing the post.
- [ ] Reprocessing the same Outbox row does not change logical membership or
      order and does not make publish fail retroactively.
- [ ] Old event payloads without new optional metadata are handled by loading
      the authoritative post.
- [ ] Deleted, unpublished, or non-Feed-visible posts are not newly projected.
- [ ] Dual mode is duplicate-safe; outbox mode performs no direct Feed dispatch.
- [ ] Mode changes are configuration validated and reversible.
- [ ] Tier `enforce` plus `direct` derivation or a disabled durable consumer is
      rejected; `dual|outbox` becomes eligible only after the post-commit
      tier-resolution failure-and-retry test passes.

## Dependencies And Scope

- Depends on `08-27-feed-author-delivery-tier`; this child consumes the Tier
  resolver and unlocks the parent rollout gate for Tier `enforce`.
- Does not replace the downstream fanout consumer or implement Timeline bulk
  rebuild; those are later sibling tasks.
