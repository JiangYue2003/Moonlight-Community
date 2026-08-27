# Author Tier Repository Evidence

- `FeedWriter.OnPostPublished` selects Hybrid push when follower count is at or
  below 1,000 and pull otherwise.
- `PublishLogic` maps Counter failure to follower count zero before calling the
  writer.
- `FeedReader.classifyFollowingsResult` independently calls batch Counter and
  maps Counter failure to all-normal with `cacheable=false`.
- RouteSnapshot caches successful following/BigV classification for at most
  five seconds and intentionally does not cache dependency-degraded results.
- Reader Inbox is not filtered by author tier; tier only determines which
  additional BigV Outboxes are merged. Monotonic promotion therefore needs no
  historical Inbox migration.
- The repository uses numbered golang-migrate SQL files and custom narrow model
  interfaces over sqlx/go-zero generated models.
- `services/relation/shared/model.FollowerModel.CountActive` already counts
  authoritative active `follower` rows with `to_user_id=? AND rel_status=1`;
  Tier read resolution additionally needs a bounded batch form or equivalent
  narrow adapter.
- Existing local and Docker Feed evolution options default off and are validated
  in service-context/config tests.

## Planning Decisions

- The tier remains `normal -> bigv` only. A permanent promotion requires an
  accepted Counter or MySQL count above 1,000.
- Counter failure or missing data falls back to the MySQL active-follower
  count. Dual-source failure returns a retryable tier-resolution error without
  delivery dispatch, promotion, or RouteSnapshot caching.
- Tier `enforce` remains blocked until the transactional Outbox compensation
  path is active and verified.
- The implementation preserves forced strategy, Cursor, PageCache,
  RouteSnapshot, and Combined Pipeline behavior and must not reuse generic CRUD
  that permits BigV demotion.
