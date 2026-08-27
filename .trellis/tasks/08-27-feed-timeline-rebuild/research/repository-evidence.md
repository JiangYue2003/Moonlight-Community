# Timeline Rebuild Repository Evidence

- Redis keys are `feed:inbox:{reader}` and `feed:bigv:{author}` ZSETs with post
  ID members and second-resolution time scores.
- Live limits/TTLs are 1,000/7 days for Inbox and 100/24 hours for BigV Outbox.
- `know_posts` has `(creator_id,status,publish_time)` indexing and stores current
  publication, visibility, author, and publish time.
- `following`/`follower` rows with `rel_status=1` are the relationship truth.
- Existing counter reconciler demonstrates periodic SCAN and bounded batches,
  while search backfill demonstrates bounded source pagination; neither has a
  resumable durable checkpoint suitable for Feed recovery.
- The load-test tooling already uses explicit run IDs, checkpoint validation,
  fingerprints, and fail-closed mutation confirmation. Rebuild should follow
  those safety principles without depending on load-test code.
- Current read return filtering accepts only published `public` or `followers`
  posts from currently allowed creators.

The rebuild implementation must share live Feed key/limit/order helpers and
must verify a key before advancing durable progress.
